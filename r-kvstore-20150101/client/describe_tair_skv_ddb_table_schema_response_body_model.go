// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairSkvDdbTableSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTairSkvDdbTableSchemaResponseBody
	GetRequestId() *string
	SetSchema(v string) *DescribeTairSkvDdbTableSchemaResponseBody
	GetSchema() *string
	SetTtlSpec(v string) *DescribeTairSkvDdbTableSchemaResponseBody
	GetTtlSpec() *string
}

type DescribeTairSkvDdbTableSchemaResponseBody struct {
	// example:
	//
	// A1604E1B-6825-1577-BBDA-2A64E8D5F126
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {"attributeDefinitions":[{"attributeType":"S","attributeName":"pk"},{"attributeType":"S","attributeName":"sk"}],"keySchema":[{"attributeName":"pk","keyType":"HASH"},{"attributeName":"sk","keyType":"RANGE"}]}
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// example:
	//
	// {"attributeName":"Expiretime","enabled":true}
	TtlSpec *string `json:"TtlSpec,omitempty" xml:"TtlSpec,omitempty"`
}

func (s DescribeTairSkvDdbTableSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairSkvDdbTableSchemaResponseBody) GetSchema() *string {
	return s.Schema
}

func (s *DescribeTairSkvDdbTableSchemaResponseBody) GetTtlSpec() *string {
	return s.TtlSpec
}

func (s *DescribeTairSkvDdbTableSchemaResponseBody) SetRequestId(v string) *DescribeTairSkvDdbTableSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaResponseBody) SetSchema(v string) *DescribeTairSkvDdbTableSchemaResponseBody {
	s.Schema = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaResponseBody) SetTtlSpec(v string) *DescribeTairSkvDdbTableSchemaResponseBody {
	s.TtlSpec = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
