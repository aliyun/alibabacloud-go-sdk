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
	// The page number. The value starts from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The dimension by which to filter job templates. Valid values:
	//
	// - **appName**: The name of the job template.
	//
	// - **appIds**: The ID of the job template.
	//
	// example:
	//
	// appName
	FieldType *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	// The name or ID of the target job template. This value corresponds to the dimension specified by **FieldType**.
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
	// The sorting method for the job templates. Valid values:
	//
	// - **running**: Sorts by the number of running instances.
	//
	// - **instances**: Sorts by the number of destination instances.
	//
	// example:
	//
	// running
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page. Valid values: 0 to 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to sort the results in ascending or descending order based on the field specified by the **OrderBy*	- parameter. Valid values:
	//
	// - **true**: ascending order.
	//
	// - **false**: descending order.
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// A list of tags. This is a JSON string. The value consists of the following parts:
	//
	// - **key**: The tag key.
	//
	// - **value**: The tag value.
	//
	// example:
	//
	// [{"key":"key","value":"value"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The workload. Set the value to `job`.
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
