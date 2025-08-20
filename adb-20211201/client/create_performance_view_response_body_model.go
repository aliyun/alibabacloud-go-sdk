// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePerformanceViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreatePerformanceViewResponseBody
	GetAccessDeniedDetail() *string
	SetCreateStatus(v string) *CreatePerformanceViewResponseBody
	GetCreateStatus() *string
	SetRequestId(v string) *CreatePerformanceViewResponseBody
	GetRequestId() *string
}

type CreatePerformanceViewResponseBody struct {
	// The details about the access denial.
	//
	// >  This parameter is returned only if Resource Access Management (RAM) permission verification failed.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1*****************7",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcKwyhk62IeYly4hQ+5IpXjkh1GQXuDRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "2***************9",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The creation result. Valid values:
	//
	// 	- **SUCCESS**
	//
	// 	- **FAILED**
	//
	// example:
	//
	// SUCCESS
	CreateStatus *string `json:"CreateStatus,omitempty" xml:"CreateStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E031AABF-BD56-5966-A063-4283EF18DB45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePerformanceViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePerformanceViewResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePerformanceViewResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreatePerformanceViewResponseBody) GetCreateStatus() *string {
	return s.CreateStatus
}

func (s *CreatePerformanceViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePerformanceViewResponseBody) SetAccessDeniedDetail(v string) *CreatePerformanceViewResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreatePerformanceViewResponseBody) SetCreateStatus(v string) *CreatePerformanceViewResponseBody {
	s.CreateStatus = &v
	return s
}

func (s *CreatePerformanceViewResponseBody) SetRequestId(v string) *CreatePerformanceViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePerformanceViewResponseBody) Validate() error {
	return dara.Validate(s)
}
