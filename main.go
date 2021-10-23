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
	messageEn := i18n.Message{ //1
		ID:    "hello",
		Other: "Hello!",
	}
	messageFr := i18n.Message{ //2
		ID:    "hello",
		Other: "Bonjour!",
	}

	bundle := i18n.NewBundle(language.English)       //1
	bundle.AddMessages(language.English, &messageEn) //2
	bundle.AddMessages(language.French, &messageFr)  //3
	localizer := i18n.NewLocalizer(bundle,           //4
		language.French.String(),
		language.English.String())
	localizeConfig := i18n.LocalizeConfig{ //5
		MessageID: "hello",
	}
	localization, _ := localizer.Localize(&localizeConfig) //6

	fmt.Println(localization)
}
