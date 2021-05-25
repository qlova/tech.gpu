package abc_test

import (
	"testing"

	"qlova.org/should"
	"qlova.tech/abc"
)

var plurals = [][2]string{
	{"", ""},
	{"human", "humans"},
	{"movie", "movies"},
	{"ox", "oxen"},
	{"user", "users"},
	{"cat", "cats"},
	{"truss", "trusses"},
	{"bus", "buses"},
	{"marsh", "marshes"},
	{"lunch", "lunches"},
	{"tax", "taxes"},
	{"blitz", "blitzes"},
	{"fez", "fezzes"},
	{"wolf", "wolves"},
	{"roof", "roofs"},
	{"belief", "beliefs"},
	{"chef", "chefs"},
	{"chief", "chiefs"},
	{"city", "cities"},
	{"puppy", "puppies"},
	{"ray", "rays"},
	{"boy", "boys"},
	{"potato", "potatoes"},
	{"tomato", "tomatoes"},
	{"photo", "photos"},
	{"piano", "pianos"},
	{"halo", "halos"},
	{"cactus", "cacti"},
	{"focus", "foci"},
	{"datum", "data"},
	{"analysis", "analyses"},
	{"ellipsis", "ellipses"},
	{"phenomenon", "phenomena"},
	{"criterion", "criteria"},
	{"sheep", "sheep"},
	{"series", "series"},
	{"species", "species"},
	{"dear", "dears"},
	{"deer", "deer"},
	{"child", "children"},
	{"goose", "geese"},
	{"woman", "women"},
	{"tooth", "teeth"},
	{"foot", "feet"},
	{"mouse", "mice"},
	{"person", "people"},
	{"search", "searches"},
	{"switch", "switches"},
	{"fix", "fixes"},
	{"ovum", "ova"},
	{"box", "boxes"},
	{"fox", "foxes"},
	{"process", "processes"},
	{"address", "addresses"},
	{"case", "cases"},
	{"stack", "stacks"},
	{"wish", "wishes"},
	{"fish", "fish"},
	{"jeans", "jeans"},
	{"funky jeans", "funky jeans"},
	{"category", "categories"},
	{"query", "queries"},
	{"ability", "abilities"},
	{"agency", "agencies"},
	{"movie", "movies"},
	{"archive", "archives"},
	{"index", "indices"},
	{"wife", "wives"},
	{"safe", "saves"},
	{"half", "halves"},
	{"move", "moves"},
	{"salesperson", "salespeople"},
	{"person", "people"},
	{"spokesman", "spokesmen"},
	{"basis", "bases"},
	{"diagnosis", "diagnoses"},
	{"diagnosis_a", "diagnosis_as"},
	{"datum", "data"},
	{"stadium", "stadia"},
	{"analysis", "analyses"},
	{"node_child", "node_children"},
	{"child", "children"},
	{"experience", "experiences"},
	{"day", "days"},
	{"comment", "comments"},
	{"foobar", "foobars"},
	{"newsletter", "newsletters"},
	{"news", "news"},
	{"series", "series"},
	{"species", "species"},
	{"quiz", "quizzes"},
	{"perspective", "perspectives"},
	{"ox", "oxen"},
	{"photo", "photos"},
	{"buffalo", "buffaloes"},
	{"tomato", "tomatoes"},
	{"dwarf", "dwarves"},
	{"elf", "elves"},
	{"information", "information"},
	{"equipment", "equipment"},
	{"bus", "buses"},
	{"status", "statuses"},
	{"Status", "Statuses"},
	{"status_code", "status_codes"},
	{"mouse", "mice"},
	{"louse", "lice"},
	{"house", "houses"},
	{"spouse", "spouses"},
	{"octopus", "octopi"},
	{"virus", "viri"},
	{"alias", "aliases"},
	{"portfolio", "portfolios"},
	{"matrix", "matrices"},
	{"axis", "axes"},
	{"testis", "testes"},
	{"crisis", "crises"},
	{"rice", "rice"},
	{"shoe", "shoes"},
	{"horse", "horses"},
	{"prize", "prizes"},
	{"edge", "edges"},
	{"database", "databases"},
	{"circus", "circuses"},
	{"plus", "pluses"},
	{"fuse", "fuses"},
	{"prometheus", "prometheuses"},
	{"vedalia", "vedalias"},
	{"media", "media"},
	{"field", "fields"},
	{"custom_field", "custom_fields"},
	//{"sportsEquipment", "sportsEquipment"},
	//{"payment_information", "payment_information"},
	//{"great_person", "great_people"},
}

func TestPlurals(t *testing.T) {
	for _, plural := range plurals {
		should.Be(plural[1])(abc.ToPlural(plural[0])).Test(t)
	}
}