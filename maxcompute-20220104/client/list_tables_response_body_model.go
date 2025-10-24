// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListTablesResponseBodyData) *ListTablesResponseBody
	GetData() *ListTablesResponseBodyData
	SetRequestId(v string) *ListTablesResponseBody
	GetRequestId() *string
}

type ListTablesResponseBody struct {
	// The returned data.
	Data *ListTablesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0a06dd4516687375802853481ec9fd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetData() *ListTablesResponseBodyData {
	return s.Data
}

func (s *ListTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesResponseBody) SetData(v *ListTablesResponseBodyData) *ListTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTablesResponseBodyData struct {
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The information about tables.
	Tables []*ListTablesResponseBodyDataTables `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyData) GetMarker() *string {
	return s.Marker
}

func (s *ListTablesResponseBodyData) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListTablesResponseBodyData) GetTables() []*ListTablesResponseBodyDataTables {
	return s.Tables
}

func (s *ListTablesResponseBodyData) SetMarker(v string) *ListTablesResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListTablesResponseBodyData) SetMaxItem(v int32) *ListTablesResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListTablesResponseBodyData) SetTables(v []*ListTablesResponseBodyDataTables) *ListTablesResponseBodyData {
	s.Tables = v
	return s
}

func (s *ListTablesResponseBodyData) Validate() error {
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTablesResponseBodyDataTables struct {
	// The time when the table was created.
	//
	// example:
	//
	// 2022-01-17T07:07:47Z
	CreationTime *int64 `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// The display name of the table.
	//
	// example:
	//
	// sale_detail
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// dim_odps
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the table.
	//
	// example:
	//
	// 1887853961230110
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The schema to which the table belongs.
	//
	// example:
	//
	// default
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The type of the table.
	//
	// example:
	//
	// internal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTablesResponseBodyDataTables) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyDataTables) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ListTablesResponseBodyDataTables) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListTablesResponseBodyDataTables) GetName() *string {
	return s.Name
}

func (s *ListTablesResponseBodyDataTables) GetOwner() *string {
	return s.Owner
}

func (s *ListTablesResponseBodyDataTables) GetSchema() *string {
	return s.Schema
}

func (s *ListTablesResponseBodyDataTables) GetType() *string {
	return s.Type
}

func (s *ListTablesResponseBodyDataTables) SetCreationTime(v int64) *ListTablesResponseBodyDataTables {
	s.CreationTime = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetDisplayName(v string) *ListTablesResponseBodyDataTables {
	s.DisplayName = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetName(v string) *ListTablesResponseBodyDataTables {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetOwner(v string) *ListTablesResponseBodyDataTables {
	s.Owner = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetSchema(v string) *ListTablesResponseBodyDataTables {
	s.Schema = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetType(v string) *ListTablesResponseBodyDataTables {
	s.Type = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) Validate() error {
	return dara.Validate(s)
}
