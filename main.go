package solr
import (
  "fmt"
  "net/http"
)

func main() {
  client := new solrClient{url: "http://localhost:8983/solr"}
  var query queryBuilder
  query.Query()
  query.Filter()
}
