package thesaurus
tyep Thesaurus interface {
	Synonyms(term string) ([]string, error)
}