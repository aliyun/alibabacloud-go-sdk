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
	SetProgrammingLanguage(v string) *ListApplicationsRequest
	GetProgrammingLanguage() *string
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
	// The Serverless App Engine (SAE) application type.
	//
	// - **micro_service.**
	//
	// - **web.**
	//
	// - **job.**
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
	// The dimension by which to filter applications. Valid values:
	//
	// - **appName**: application name.
	//
	// - **appIds**: application ID.
	//
	// - **slbIps**: SLB IP address.
	//
	// - **instanceIps**: instance IP address.
	//
	// example:
	//
	// appName
	FieldType *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	// The application name, application ID, SLB IP address, or instance IP address of the target application.
	//
	// example:
	//
	// demo-app
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// Specifies whether the application is stateful.
	IsStateful *string `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:demo
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The application version. Valid values:
	//
	// - lite: Lite Edition
	//
	// - std: Standard Edition
	//
	// - pro: Professional Edition
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The field by which to sort applications. Valid values:
	//
	// - **runnings**: sorts by the current number target instances.
	//
	// - **instances**: sorts by the target number target instances.
	//
	// example:
	//
	// runnings
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page in a paging query. Valid values: [0,10000].
	//
	// example:
	//
	// 20
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// Specifies whether to sort application instances by running status. If instances have the same status, they are sorted by instance ID. Valid values:
	//
	//   - **true**: sorts in ascending order. Instances are arranged based on the startup sequence. For example, to reach the running state, an instance must go through steps such as starting the container, pulling the image, and initializing the instance.
	//
	//   - **false**: sorts in descending order.
	//
	// The ascending order of instances is as follows:
	//
	// 1. **Error**: an error occurred during instance startup.
	//
	// 2. **CrashLoopBackOff**: the container failed to start, encountered an error during startup, and encountered an error again after restart.
	//
	// 3. **ErrImagePull**: an error occurred while pulling the container image for the instance.
	//
	// 4. **ImagePullBackOff**: the container image cannot be obtained.
	//
	// 5. **Pending**: the instance is waiting to be scheduled.
	//
	// 6. **Unknown**: an unknown exception occurred.
	//
	// 7. **Terminating**: the instance is being terminated.
	//
	// 8. **NotFound**: the instance cannot be found.
	//
	// 9. **PodInitializing**: the instance is being initialized.
	//
	// 10. **Init:0/1**: the instance is initializing.
	//
	// 11. **Running**: the instance is running.
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The tag key-value pairs. Valid values:
	//
	// - **key**: the tag key. The length must be in the range of [1,128].
	//
	// - **value**: the tag value. The length must be in the range of [1,128].
	//
	// Tags are case-sensitive. If you specify multiple tags, all specified tags are created and attached to the resource. Each tag key on the same resource can have only one tag value. If you add a tag key that already exists, the corresponding tag value is updated to the new value.
	//
	// Tags cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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

func (s *ListApplicationsRequest) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
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

func (s *ListApplicationsRequest) SetProgrammingLanguage(v string) *ListApplicationsRequest {
	s.ProgrammingLanguage = &v
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
