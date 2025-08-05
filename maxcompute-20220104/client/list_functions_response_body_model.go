// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListFunctionsResponseBodyData) *ListFunctionsResponseBody
	GetData() *ListFunctionsResponseBodyData
	SetRequestId(v string) *ListFunctionsResponseBody
	GetRequestId() *string
}

type ListFunctionsResponseBody struct {
	// The returned data.
	Data *ListFunctionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0be3e0b716671885050924814e3623
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) GetData() *ListFunctionsResponseBodyData {
	return s.Data
}

func (s *ListFunctionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFunctionsResponseBody) SetData(v *ListFunctionsResponseBodyData) *ListFunctionsResponseBody {
	s.Data = v
	return s
}

func (s *ListFunctionsResponseBody) SetRequestId(v string) *ListFunctionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFunctionsResponseBodyData struct {
	// The information about each function.
	Functions []*ListFunctionsResponseBodyDataFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Repeated"`
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
}

func (s ListFunctionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyData) GetFunctions() []*ListFunctionsResponseBodyDataFunctions {
	return s.Functions
}

func (s *ListFunctionsResponseBodyData) GetMarker() *string {
	return s.Marker
}

func (s *ListFunctionsResponseBodyData) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListFunctionsResponseBodyData) SetFunctions(v []*ListFunctionsResponseBodyDataFunctions) *ListFunctionsResponseBodyData {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBodyData) SetMarker(v string) *ListFunctionsResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListFunctionsResponseBodyData) SetMaxItem(v int32) *ListFunctionsResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListFunctionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListFunctionsResponseBodyDataFunctions struct {
	// The class in which the function was defined.
	//
	// example:
	//
	// abc
	Class *string `json:"class,omitempty" xml:"class,omitempty"`
	// The time when the function was created. Unit: milliseconds.
	//
	// example:
	//
	// 1664505167000
	CreationTime *int64 `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// The display name of the function.
	//
	// example:
	//
	// getdate
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the function.
	//
	// example:
	//
	// getdate
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the function.
	//
	// example:
	//
	// odpsowner
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The name of the resource that was associated with the function.
	//
	// example:
	//
	// abc
	Resources *string `json:"resources,omitempty" xml:"resources,omitempty"`
	// The schema of the function.
	//
	// example:
	//
	// abc
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s ListFunctionsResponseBodyDataFunctions) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBodyDataFunctions) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyDataFunctions) GetClass() *string {
	return s.Class
}

func (s *ListFunctionsResponseBodyDataFunctions) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ListFunctionsResponseBodyDataFunctions) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListFunctionsResponseBodyDataFunctions) GetName() *string {
	return s.Name
}

func (s *ListFunctionsResponseBodyDataFunctions) GetOwner() *string {
	return s.Owner
}

func (s *ListFunctionsResponseBodyDataFunctions) GetResources() *string {
	return s.Resources
}

func (s *ListFunctionsResponseBodyDataFunctions) GetSchema() *string {
	return s.Schema
}

func (s *ListFunctionsResponseBodyDataFunctions) SetClass(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Class = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetCreationTime(v int64) *ListFunctionsResponseBodyDataFunctions {
	s.CreationTime = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetDisplayName(v string) *ListFunctionsResponseBodyDataFunctions {
	s.DisplayName = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetName(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetOwner(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Owner = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetResources(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Resources = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetSchema(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Schema = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) Validate() error {
	return dara.Validate(s)
}
