package solr

type queryParams [string]string

type queryBuilder struct {
  q queryParams
  fq queryParams
  sort [string]
  start int
  rows int
  fl [string]
}

// q
func (query *queryBuilder) Query(q queryParams) queryBuilder {
  if q.queryParams ==  nil {
    q.queryParams = q
  } else {
    // merge with current
    for k, v := range q {
      q.queryParams[k] = v
    }
  }
  return query
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
func (query *queryBuilder) Limit() queryBuilder {
}

// select
func (query *queryBuilder) Select() queryBuilder {
}
