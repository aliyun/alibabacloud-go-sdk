// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResourceTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v []*GetInstanceResourceTableResponseBodyFields) *GetInstanceResourceTableResponseBody
	GetFields() []*GetInstanceResourceTableResponseBodyFields
	SetRequestId(v string) *GetInstanceResourceTableResponseBody
	GetRequestId() *string
	SetTableName(v string) *GetInstanceResourceTableResponseBody
	GetTableName() *string
}

type GetInstanceResourceTableResponseBody struct {
	Fields []*GetInstanceResourceTableResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetInstanceResourceTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResourceTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceTableResponseBody) GetFields() []*GetInstanceResourceTableResponseBodyFields {
	return s.Fields
}

func (s *GetInstanceResourceTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResourceTableResponseBody) GetTableName() *string {
	return s.TableName
}

func (s *GetInstanceResourceTableResponseBody) SetFields(v []*GetInstanceResourceTableResponseBodyFields) *GetInstanceResourceTableResponseBody {
	s.Fields = v
	return s
}

func (s *GetInstanceResourceTableResponseBody) SetRequestId(v string) *GetInstanceResourceTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResourceTableResponseBody) SetTableName(v string) *GetInstanceResourceTableResponseBody {
	s.TableName = &v
	return s
}

func (s *GetInstanceResourceTableResponseBody) Validate() error {
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

type GetInstanceResourceTableResponseBodyFields struct {
	// example:
	//
	// false
	IsDimensionField *bool `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	IsPartitionField *bool `json:"IsPartitionField,omitempty" xml:"IsPartitionField,omitempty"`
	// example:
	//
	// ""
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BIGINT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResourceTableResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResourceTableResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceTableResponseBodyFields) GetIsDimensionField() *bool {
	return s.IsDimensionField
}

func (s *GetInstanceResourceTableResponseBodyFields) GetIsPartitionField() *bool {
	return s.IsPartitionField
}

func (s *GetInstanceResourceTableResponseBodyFields) GetMeaning() *string {
	return s.Meaning
}

func (s *GetInstanceResourceTableResponseBodyFields) GetName() *string {
	return s.Name
}

func (s *GetInstanceResourceTableResponseBodyFields) GetType() *string {
	return s.Type
}

func (s *GetInstanceResourceTableResponseBodyFields) SetIsDimensionField(v bool) *GetInstanceResourceTableResponseBodyFields {
	s.IsDimensionField = &v
	return s
}

func (s *GetInstanceResourceTableResponseBodyFields) SetIsPartitionField(v bool) *GetInstanceResourceTableResponseBodyFields {
	s.IsPartitionField = &v
	return s
}

func (s *GetInstanceResourceTableResponseBodyFields) SetMeaning(v string) *GetInstanceResourceTableResponseBodyFields {
	s.Meaning = &v
	return s
}

func (s *GetInstanceResourceTableResponseBodyFields) SetName(v string) *GetInstanceResourceTableResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetInstanceResourceTableResponseBodyFields) SetType(v string) *GetInstanceResourceTableResponseBodyFields {
	s.Type = &v
	return s
}

func (s *GetInstanceResourceTableResponseBodyFields) Validate() error {
	return dara.Validate(s)
}
