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
	SetNewSaeVersion(v string) *ListApplicationsRequest
	GetNewSaeVersion() *string
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
	// The type of the SAE application.
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
	// The field to filter applications by. Valid values:
	//
	// - **appName**: The application name.
	//
	// - **appIds**: The application ID.
	//
	// - **slbIps**: The SLB IP address.
	//
	// - **instanceIps**: The instance IP address.
	//
	// example:
	//
	// appName
	FieldType *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	// The value for the field specified by `FieldType`. This can be an application name, application ID, SLB IP address, or instance IP address.
	//
	// example:
	//
	// demo-app
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// Filters applications by whether they are stateful. Set this parameter to `true` to return only stateful applications, or to `false` to return only stateless applications.
	IsStateful *string `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:demo
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The edition of the application:
	//
	// - `lite`: Lite
	//
	// - `std`: Standard
	//
	// - `pro`: Pro
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The field to sort the applications by. Valid values:
	//
	// - **runnings**: Sorts the applications by the current instance count.
	//
	// - **instances**: Sorts the applications by the target instance count.
	//
	// example:
	//
	// runnings
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries to return per page. Valid values: 0 to 10000.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort order. Valid values:
	//
	// - **true**: Sorts the results in ascending order.
	//
	// - **false**: Sorts the results in descending order.
	//
	// 1. ****
	//
	// 2. ****
	//
	// 3. ****
	//
	// 4. ****
	//
	// 5. ****
	//
	// 6. ****
	//
	// 7. ****
	//
	// 8. ****
	//
	// 9. ****
	//
	// 10. ****
	//
	// 11. ****
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// Filters applications by tags. The tags are specified as a JSON string that contains an array of key-value pairs.
	//
	// - **key**: The tag key, which can be 1 to 128 characters in length.
	//
	// - **value**: The tag value, which can be 1 to 128 characters in length.
	//
	// This parameter is case-sensitive. An application is returned only if it matches all specified tags. On a resource, a tag key can have only one tag value.
	//
	// The tag key and tag value cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
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

func (s *ListApplicationsRequest) GetNewSaeVersion() *string {
	return s.NewSaeVersion
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

func (s *ListApplicationsRequest) SetNewSaeVersion(v string) *ListApplicationsRequest {
	s.NewSaeVersion = &v
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
