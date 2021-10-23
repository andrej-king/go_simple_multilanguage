package main

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func main() {
	multilanguageByJson("fr")
	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: "hello", //1
	}
	localizationUsingJson, _ := localizer.Localize(&localizeConfigWelcome) //2
	fmt.Println(localizationUsingJson)
}

func multilanguageByJson(lang string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("resources/en.json")
	bundle.LoadMessageFile("resources/fr.json")

	// The sequence of the displayed language
	if lang == "fr" {
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String())
	} else {
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.French.String())
	}

}

// multilanguageByHardcoreText parse lang by text
func multilanguageByText() {
	messageEn := i18n.Message{
		ID:    "hello",
		Other: "Hello!",
	}
	messageFr := i18n.Message{
		ID:    "hello",
		Other: "Bonjour!",
	}

	bundle := i18n.NewBundle(language.English)
	bundle.AddMessages(language.English, &messageEn)
	bundle.AddMessages(language.French, &messageFr)
	localizer := i18n.NewLocalizer(bundle,
		language.French.String(),
		language.English.String())
	localizeConfig := i18n.LocalizeConfig{
		MessageID: "hello",
	}
	localization, _ := localizer.Localize(&localizeConfig)

	fmt.Println(localization)
}
