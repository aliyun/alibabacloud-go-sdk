// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeSchemasResponseBodyItems) *DescribeSchemasResponseBody
	GetItems() *DescribeSchemasResponseBodyItems
	SetRequestId(v string) *DescribeSchemasResponseBody
	GetRequestId() *string
}

type DescribeSchemasResponseBody struct {
	// The information about the databases of the cluster.
	Items *DescribeSchemasResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBody) GetItems() *DescribeSchemasResponseBodyItems {
	return s.Items
}

func (s *DescribeSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSchemasResponseBody) SetItems(v *DescribeSchemasResponseBodyItems) *DescribeSchemasResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSchemasResponseBody) SetRequestId(v string) *DescribeSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSchemasResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSchemasResponseBodyItems struct {
	Schema []*DescribeSchemasResponseBodyItemsSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeSchemasResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchemasResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBodyItems) GetSchema() []*DescribeSchemasResponseBodyItemsSchema {
	return s.Schema
}

func (s *DescribeSchemasResponseBodyItems) SetSchema(v []*DescribeSchemasResponseBodyItemsSchema) *DescribeSchemasResponseBodyItems {
	s.Schema = v
	return s
}

func (s *DescribeSchemasResponseBodyItems) Validate() error {
	if s.Schema != nil {
		for _, item := range s.Schema {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSchemasResponseBodyItemsSchema struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeSchemasResponseBodyItemsSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchemasResponseBodyItemsSchema) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBodyItemsSchema) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSchemasResponseBodyItemsSchema) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeSchemasResponseBodyItemsSchema) SetDBClusterId(v string) *DescribeSchemasResponseBodyItemsSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSchemasResponseBodyItemsSchema) SetSchemaName(v string) *DescribeSchemasResponseBodyItemsSchema {
	s.SchemaName = &v
	return s
}

func (s *DescribeSchemasResponseBodyItemsSchema) Validate() error {
	return dara.Validate(s)
}
