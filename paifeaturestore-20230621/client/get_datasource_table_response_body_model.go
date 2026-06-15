// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasourceTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v []*GetDatasourceTableResponseBodyFields) *GetDatasourceTableResponseBody
	GetFields() []*GetDatasourceTableResponseBodyFields
	SetRequestId(v string) *GetDatasourceTableResponseBody
	GetRequestId() *string
	SetTableName(v string) *GetDatasourceTableResponseBody
	GetTableName() *string
}

type GetDatasourceTableResponseBody struct {
	// The list of fields.
	Fields []*GetDatasourceTableResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D7B2F8C4-49C7-5CFA-8075-9D715A114873
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the data table.
	//
	// example:
	//
	// table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetDatasourceTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasourceTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasourceTableResponseBody) GetFields() []*GetDatasourceTableResponseBodyFields {
	return s.Fields
}

func (s *GetDatasourceTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasourceTableResponseBody) GetTableName() *string {
	return s.TableName
}

func (s *GetDatasourceTableResponseBody) SetFields(v []*GetDatasourceTableResponseBodyFields) *GetDatasourceTableResponseBody {
	s.Fields = v
	return s
}

func (s *GetDatasourceTableResponseBody) SetRequestId(v string) *GetDatasourceTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasourceTableResponseBody) SetTableName(v string) *GetDatasourceTableResponseBody {
	s.TableName = &v
	return s
}

func (s *GetDatasourceTableResponseBody) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDatasourceTableResponseBodyFields struct {
	// The attributes of the field. Valid values:
	//
	// ● Partition: indicates that the field is a partition field.
	//
	// ● EventTime: indicates that the field is an event time field.
	//
	// ● PrimaryKey: indicates that the field is a primary key field.
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The name of the field.
	//
	// example:
	//
	// field1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type of the field. Valid values:
	//
	// ● INT32
	//
	// ● INT64
	//
	// ● FLOAT
	//
	// ● DOUBLE
	//
	// ● STRING
	//
	// ● BOOLEAN
	//
	// ● TIMESTAMP
	//
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDatasourceTableResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s GetDatasourceTableResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetDatasourceTableResponseBodyFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *GetDatasourceTableResponseBodyFields) GetName() *string {
	return s.Name
}

func (s *GetDatasourceTableResponseBodyFields) GetType() *string {
	return s.Type
}

func (s *GetDatasourceTableResponseBodyFields) SetAttributes(v []*string) *GetDatasourceTableResponseBodyFields {
	s.Attributes = v
	return s
}

func (s *GetDatasourceTableResponseBodyFields) SetName(v string) *GetDatasourceTableResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetDatasourceTableResponseBodyFields) SetType(v string) *GetDatasourceTableResponseBodyFields {
	s.Type = &v
	return s
}

func (s *GetDatasourceTableResponseBodyFields) Validate() error {
	return dara.Validate(s)
}
