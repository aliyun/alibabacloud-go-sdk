// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) *DescribeOpenSearchResourceUsageResponseBody
	GetAccessDeniedDetail() *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail
	SetData(v *DescribeOpenSearchResourceUsageResponseBodyData) *DescribeOpenSearchResourceUsageResponseBody
	GetData() *DescribeOpenSearchResourceUsageResponseBodyData
	SetRequestId(v string) *DescribeOpenSearchResourceUsageResponseBody
	GetRequestId() *string
}

type DescribeOpenSearchResourceUsageResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned data.
	Data *DescribeOpenSearchResourceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenSearchResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchResourceUsageResponseBody) GetAccessDeniedDetail() *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeOpenSearchResourceUsageResponseBody) GetData() *DescribeOpenSearchResourceUsageResponseBodyData {
	return s.Data
}

func (s *DescribeOpenSearchResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenSearchResourceUsageResponseBody) SetAccessDeniedDetail(v *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) *DescribeOpenSearchResourceUsageResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBody) SetData(v *DescribeOpenSearchResourceUsageResponseBodyData) *DescribeOpenSearchResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBody) SetRequestId(v string) *DescribeOpenSearchResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authentication principal.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The diagnostic information.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// NoPermissionType
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchResourceUsageResponseBodyData struct {
	// The total number of documents in the cluster.
	//
	// example:
	//
	// 1000000
	DocCount *int64 `json:"DocCount,omitempty" xml:"DocCount,omitempty"`
	// The number of indexes. This is a filter condition for the number of indexes that the missing index table currently has. The input format is `operator + separator "" + index count`, for example, `>=100`.
	//
	// example:
	//
	// >=
	IndexCount *int32 `json:"IndexCount,omitempty" xml:"IndexCount,omitempty"`
	// The used storage space, in bytes.
	//
	// example:
	//
	// 107374182400
	StorageSizeInBytes *int64 `json:"StorageSizeInBytes,omitempty" xml:"StorageSizeInBytes,omitempty"`
	// The total storage capacity, in bytes.
	//
	// example:
	//
	// 536870912000
	StorageTotalInBytes *int64 `json:"StorageTotalInBytes,omitempty" xml:"StorageTotalInBytes,omitempty"`
	// The storage space usage.
	//
	// example:
	//
	// 20.0
	StorageUsagePercent *float64 `json:"StorageUsagePercent,omitempty" xml:"StorageUsagePercent,omitempty"`
}

func (s DescribeOpenSearchResourceUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) GetDocCount() *int64 {
	return s.DocCount
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) GetIndexCount() *int32 {
	return s.IndexCount
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) GetStorageSizeInBytes() *int64 {
	return s.StorageSizeInBytes
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) GetStorageTotalInBytes() *int64 {
	return s.StorageTotalInBytes
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) GetStorageUsagePercent() *float64 {
	return s.StorageUsagePercent
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) SetDocCount(v int64) *DescribeOpenSearchResourceUsageResponseBodyData {
	s.DocCount = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) SetIndexCount(v int32) *DescribeOpenSearchResourceUsageResponseBodyData {
	s.IndexCount = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) SetStorageSizeInBytes(v int64) *DescribeOpenSearchResourceUsageResponseBodyData {
	s.StorageSizeInBytes = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) SetStorageTotalInBytes(v int64) *DescribeOpenSearchResourceUsageResponseBodyData {
	s.StorageTotalInBytes = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) SetStorageUsagePercent(v float64) *DescribeOpenSearchResourceUsageResponseBodyData {
	s.StorageUsagePercent = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
