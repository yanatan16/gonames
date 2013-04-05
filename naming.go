// Package gonames is very simple. It creates a human readable name. It has about 10,000 permutations. It uses names from Star Trek starships.
package gonames

import (
	"fmt"
	"math/rand"
)

// Names taken from http://en.wikipedia.org/wiki/List_of_Starfleet_starships_ordered_by_class
var basenames []string = []string{"Achilles", "Adelphi", "Agamemnon", "Agincourt", "Ahwahnee", "Ajax", "Akagi", "Akira", "Al-Batani", "Amalthea",
	"Ambassador", "Andromeda", "Antares", "Apollo", "Appalachia", "Archon", "Arcos", "Aries", "Armstrong", "Aventine", "Bellerephon", "Bellerophon", "Berkeley",
	"Berlin", "Biko", "Bonchune", "Bonestell", "Bozeman", "Bradbury", "Brattain", "Brigand", "Budapest", "Buran", "Cairo", "Callisto", "Capitoline", "Carolina",
	"Centaur", "Challenger", "Charleston", "Charon", "Chekov", "Cheyenne", "Chimera", "Clement", "Cochraine", "Concorde", "Constantinople", "Constellation",
	"Constitution", "Copernicus", "Cortéz", "Crazy Horse", "Crockett", "Curie", "Curry", "Daedalus", "Dakota", "Danube", "Dauntless", "da Vinci", "Defiant",
	"Demeter", "Deneva", "Denver", "Devore", "Drake", "Endeavour", "Endurance", "Enterprise", "Enterprise-D", "Enterprise-E", "Equinox", "Esquline",
	"Essex", "Europa", "Everett", "Excalibur", "Excelsior", "Exeter", "Farragut", "Fearless", "Firebrand", "Fleming", "Frederick", "Fredrickson",
	"Freedom", "Gage", "Galatea", "Galaxy", "Galen", "Gander", "Gandhi", "Ganges", "Ganymede", "Geronimo", "Gettysburg", "Goddard", "Gorkon", "Grissom",
	"Gryphon", "Hathaway", "Havana", "Hawking", "Helin", "Hera", "Hermes", "Hokule'a", "Honshū", "Hood", "Horatio", "Horizon", "Hornet", "Intrepid", "Io",
	"Istanbul", "Jenolen", "Kearsarge", "Keaton", "Kelvin", "Khitomer", "Korolev", "Kyushu", "Lakota", "Lalo", "Lantree", "LaSalle", "Leeds", "Lexington",
	"Leyte Gulf", "London", "Luna", "Magellan", "Majestic", "Malinche", "Maltby", "Maryland", "Mediterranean", "Mekong", "Melbourne", "Merced", "Merian",
	"Merrimack", "Miranda", "Monitor", "Montgolfier", "Mulciber", "Nash", "Nautilus", "Nebula", "New Orleans", "Nightingale", "Nobel", "Norway", "Nova",
	"Oberon", "Oberth", "Odyssey", "Olympic", "Orinoco", "Pasteur", "Pathfinder", "Pegasus", "Phoenix", "Planck", "Portland", "Potemkin", "Princeton",
	"Prokofiev", "Prometheus", "Proxima", "Quirinal", "Rabin", "Raman", "Ranger", "Relativity", "Reliant", "Renaissance", "Renegade", "Republic",
	"Repulse", "Rhea", "Rhode Island", "Rigel", "Rio Grande", "Robinea", "Roosevelt", "Rotherham", "Rubicon", "Rutledge", "Saber", "Samuel B. Roberts",
	"San Francisco", "Sarajevo", "Saratoga", "Sarek", "Scarab", "Sentinel", "Sequoia", "Sheffield", "Shenandoah", "ShirKahr", "Sitak", "Sovereign", "Soyuz",
	"Spitfire", "Springfield", "Stargazer", "Steamrunner", "Summit", "Surak", "Sutherland", "Sydney", "Tardis", "Theophrastus", "Thomas Calhoun",
	"Thomas Paine", "Thunderchild", "Tian An Men", "Titan", "T'Kumbra", "Tolstoy", "Trial", "Trident", "Trieste", "Tripoli", "Triton", "Tsiolkovsky",
	"Tuscany", "Ulysses", "Valdemar", "Valiant", "Venture", "Vesta", "Victory", "Volga", "Voyager", "Wambundu", "Wellington", "Wells", "Wyoming",
	"Yamaguchi", "Yamato", "Yangtzee Kiang", "Yeager", "Yellowstone", "Yorkshire", "Yorktown", "Yosemite", "Yukon", "Zapata", "Zeus", "Zhukov", "Zodiac"}

var lbase int = len(basenames)

// Create a random readable name.
func Rand() string {
	return fmt.Sprintf("%s_%d", basenames[rand.Int()%lbase], rand.Int()%1000)
}
