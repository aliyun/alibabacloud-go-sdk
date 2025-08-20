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
	// The queried databases.
	Items *DescribeSchemasResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 25B56BC7-4978-40B3-9E48-4B7067******
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeSchemasResponseBodyItemsSchema struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
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
