package solr

type queryBuilder struct {
}

// q
func (query *queryBuilder) Query() queryBuilder {
}

// fq
func (query *queryBuilder) Filter() queryBuilder {
}

// sort
func (query *queryBuilder) Sort() queryBuilder {
}

// start
func (query *queryBuilder) Start() queryBuilder {
}

// rows
func (query *queryBuilder) Rows() queryBuilder {
}

// select
func (query *queryBuilder) Select() queryBuilder {
}
