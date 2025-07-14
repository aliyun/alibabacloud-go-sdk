// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListApplicationsRequest
	GetAppName() *string
	SetAppSource(v string) *ListApplicationsRequest
	GetAppSource() *string
	SetCurrentPage(v int32) *ListApplicationsRequest
	GetCurrentPage() *int32
	SetFieldType(v string) *ListApplicationsRequest
	GetFieldType() *string
	SetFieldValue(v string) *ListApplicationsRequest
	GetFieldValue() *string
	SetIsStateful(v string) *ListApplicationsRequest
	GetIsStateful() *string
	SetNamespaceId(v string) *ListApplicationsRequest
	GetNamespaceId() *string
	SetOrderBy(v string) *ListApplicationsRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListApplicationsRequest
	GetPageSize() *int32
	SetReverse(v bool) *ListApplicationsRequest
	GetReverse() *bool
	SetTags(v string) *ListApplicationsRequest
	GetTags() *string
}

type ListApplicationsRequest struct {
	// The application name.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The SAE application type. Valid values:
	//
	// - **micro_service**
	//
	// - **web**
	//
	// - **job**
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Set the filtering criteria for applications. The value options are as follows:
	//
	// - appName: Application name.
	//
	// - appIds: Application IDs.
	//
	// - slbIps: SLB IP addresses.
	//
	// - instanceIps: Instance IP addresses.
	//
	// example:
	//
	// appName
	FieldType *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	// The name, ID, SLB IP, or instance IP of the target application.
	//
	// example:
	//
	// demo-app
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	IsStateful *string `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
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
	// The number of records in each page. Value range: [0,10000]
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Sort by the running status of application instances. If the statuses are the same, sort by instance ID. The value options are as follows:
	//
	// - true: Sort in ascending order. Instances are arranged according to the startup process, for example: to ultimately reach the running state, an instance must first go through steps such as starting containers, pulling images, and initializing the instance.
	//
	// - false: Sort in descending order.
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The tag in the format of a key-value pair.
	//
	// 	- **key**: the tag key. It cannot exceed 128 characters in length.
	//
	// 	- **value**: the tag value. It cannot exceed 128 characters in length.
	//
	// Tag keys and tag values are case-sensitive. If you specify multiple tags, the system adds all the tags to the specified resources. Each tag key on a resource can have only one tag value. If you create a tag that has the same key as an existing tag, the value of the existing tag is overwritten.
	//
	// Tag keys and tag values cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// [{"key":"key","value":"value"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationsRequest) GetAppSource() *string {
	return s.AppSource
}

func (s *ListApplicationsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListApplicationsRequest) GetFieldType() *string {
	return s.FieldType
}

func (s *ListApplicationsRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *ListApplicationsRequest) GetIsStateful() *string {
	return s.IsStateful
}

func (s *ListApplicationsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListApplicationsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListApplicationsRequest) GetTags() *string {
	return s.Tags
}

func (s *ListApplicationsRequest) SetAppName(v string) *ListApplicationsRequest {
	s.AppName = &v
	return s
}

func (s *ListApplicationsRequest) SetAppSource(v string) *ListApplicationsRequest {
	s.AppSource = &v
	return s
}

func (s *ListApplicationsRequest) SetCurrentPage(v int32) *ListApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListApplicationsRequest) SetFieldType(v string) *ListApplicationsRequest {
	s.FieldType = &v
	return s
}

func (s *ListApplicationsRequest) SetFieldValue(v string) *ListApplicationsRequest {
	s.FieldValue = &v
	return s
}

func (s *ListApplicationsRequest) SetIsStateful(v string) *ListApplicationsRequest {
	s.IsStateful = &v
	return s
}

func (s *ListApplicationsRequest) SetNamespaceId(v string) *ListApplicationsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListApplicationsRequest) SetOrderBy(v string) *ListApplicationsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v int32) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) SetReverse(v bool) *ListApplicationsRequest {
	s.Reverse = &v
	return s
}

func (s *ListApplicationsRequest) SetTags(v string) *ListApplicationsRequest {
	s.Tags = &v
	return s
}

func (s *ListApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
