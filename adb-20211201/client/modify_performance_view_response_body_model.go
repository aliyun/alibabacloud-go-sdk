// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPerformanceViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyPerformanceViewResponseBody
	GetAccessDeniedDetail() *string
	SetModifyStatus(v string) *ModifyPerformanceViewResponseBody
	GetModifyStatus() *string
	SetRequestId(v string) *ModifyPerformanceViewResponseBody
	GetRequestId() *string
}

type ModifyPerformanceViewResponseBody struct {
	// The details about the access denial. This parameter is returned only if Resource Access Management (RAM) permission verification failed.
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
	// The modification result. Valid values:
	//
	// 	- **SUCCESS**
	//
	// 	- **FAILED**
	//
	// example:
	//
	// SUCCESS
	ModifyStatus *string `json:"ModifyStatus,omitempty" xml:"ModifyStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7EDB8E4-9769-4233-88C7-DCA4C9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPerformanceViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPerformanceViewResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPerformanceViewResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyPerformanceViewResponseBody) GetModifyStatus() *string {
	return s.ModifyStatus
}

func (s *ModifyPerformanceViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPerformanceViewResponseBody) SetAccessDeniedDetail(v string) *ModifyPerformanceViewResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyPerformanceViewResponseBody) SetModifyStatus(v string) *ModifyPerformanceViewResponseBody {
	s.ModifyStatus = &v
	return s
}

func (s *ModifyPerformanceViewResponseBody) SetRequestId(v string) *ModifyPerformanceViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPerformanceViewResponseBody) Validate() error {
	return dara.Validate(s)
}
