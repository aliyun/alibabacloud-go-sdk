// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListJobsRequest
	GetAppName() *string
	SetCurrentPage(v int32) *ListJobsRequest
	GetCurrentPage() *int32
	SetFieldType(v string) *ListJobsRequest
	GetFieldType() *string
	SetFieldValue(v string) *ListJobsRequest
	GetFieldValue() *string
	SetNamespaceId(v string) *ListJobsRequest
	GetNamespaceId() *string
	SetOrderBy(v string) *ListJobsRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListJobsRequest
	GetPageSize() *int32
	SetReverse(v bool) *ListJobsRequest
	GetReverse() *bool
	SetTags(v string) *ListJobsRequest
	GetTags() *string
	SetWorkload(v string) *ListJobsRequest
	GetWorkload() *string
}

type ListJobsRequest struct {
	// The name of the job template.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of the page to return. The parameter value is a positive integer that is greater than or equal to 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The dimension by which applications are filtered. Valid values:
	//
	// 	- **appName**: Applications are filtered by job template name.
	//
	// 	- **appIds**: Applications are filtered by job template ID.
	//
	// example:
	//
	// appName
	FieldType *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	// Enter the name and ID of the job template.
	//
	// example:
	//
	// demo-app
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:demo
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// Specifies how applications are sorted. Valid values:
	//
	// 	- **running**: The applications are sorted based on the number of running instances.
	//
	// 	- **instances**: The applications are sorted based on the number of destination instances.
	//
	// example:
	//
	// running
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries to return on each page. Valid value: 0 to 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to sort the field names that are passed by **OrderBy*	- in ascending order. Valid values:
	//
	// 	- **true**: in ascending order
	//
	// 	- **false**: in descending order
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The tags that are displayed in a JSON string. Valid values:
	//
	// 	- **key**: the tag key
	//
	// 	- **value**: the tag value
	//
	// example:
	//
	// [{"key":"key","value":"value"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Set the value to `job`.
	//
	// example:
	//
	// job
	Workload *string `json:"Workload,omitempty" xml:"Workload,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListJobsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListJobsRequest) GetFieldType() *string {
	return s.FieldType
}

func (s *ListJobsRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *ListJobsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListJobsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListJobsRequest) GetTags() *string {
	return s.Tags
}

func (s *ListJobsRequest) GetWorkload() *string {
	return s.Workload
}

func (s *ListJobsRequest) SetAppName(v string) *ListJobsRequest {
	s.AppName = &v
	return s
}

func (s *ListJobsRequest) SetCurrentPage(v int32) *ListJobsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListJobsRequest) SetFieldType(v string) *ListJobsRequest {
	s.FieldType = &v
	return s
}

func (s *ListJobsRequest) SetFieldValue(v string) *ListJobsRequest {
	s.FieldValue = &v
	return s
}

func (s *ListJobsRequest) SetNamespaceId(v string) *ListJobsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListJobsRequest) SetOrderBy(v string) *ListJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetReverse(v bool) *ListJobsRequest {
	s.Reverse = &v
	return s
}

func (s *ListJobsRequest) SetTags(v string) *ListJobsRequest {
	s.Tags = &v
	return s
}

func (s *ListJobsRequest) SetWorkload(v string) *ListJobsRequest {
	s.Workload = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
