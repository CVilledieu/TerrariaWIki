package server

type Boss struct {
	Icon  string
	Name  string
	Stage int
}

type Character struct {
	Helm  Item
	Armor Item
	Legs  Item
	Boots Item
	Acc   [5]Item
}

type Item struct {
	Id   int
	Name string
	//? seems like best option, but may have issues when it comes to displaying the Stats.
	Stats map[string]int
}

func GetBosses() []Boss {
	return []Boss{
		{
			Icon:  "Icon",
			Name:  "Wall of Flesh",
			Stage: 2,
		},
		{
			Icon:  "Icon",
			Name:  "Skeletron",
			Stage: 2,
		},
		{
			Icon:  "Icon",
			Name:  "Eater Of Worlds",
			Stage: 2,
		},
	}
}
