package solr

type queryParam [string]string

type queryBuilder struct {
  q queryParam
  fq queryParam
  sort [string]
  start int
  rows int
  fl [string]
}

// q
func (query *queryBuilder) Query(q queryParam) queryBuilder {
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
