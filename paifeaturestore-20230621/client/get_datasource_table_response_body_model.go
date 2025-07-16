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
	Fields []*GetDatasourceTableResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// D7B2F8C4-49C7-5CFA-8075-9D715A114873
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type GetDatasourceTableResponseBodyFields struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// field1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
