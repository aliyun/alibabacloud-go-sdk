// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePerformanceViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeletePerformanceViewResponseBody
	GetAccessDeniedDetail() *string
	SetDeleteStatus(v bool) *DeletePerformanceViewResponseBody
	GetDeleteStatus() *bool
	SetRequestId(v string) *DeletePerformanceViewResponseBody
	GetRequestId() *string
}

type DeletePerformanceViewResponseBody struct {
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
	// The delete status.
	//
	// example:
	//
	// SUCCESS
	DeleteStatus *bool `json:"DeleteStatus,omitempty" xml:"DeleteStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePerformanceViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePerformanceViewResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePerformanceViewResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeletePerformanceViewResponseBody) GetDeleteStatus() *bool {
	return s.DeleteStatus
}

func (s *DeletePerformanceViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePerformanceViewResponseBody) SetAccessDeniedDetail(v string) *DeletePerformanceViewResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeletePerformanceViewResponseBody) SetDeleteStatus(v bool) *DeletePerformanceViewResponseBody {
	s.DeleteStatus = &v
	return s
}

func (s *DeletePerformanceViewResponseBody) SetRequestId(v string) *DeletePerformanceViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePerformanceViewResponseBody) Validate() error {
	return dara.Validate(s)
}
