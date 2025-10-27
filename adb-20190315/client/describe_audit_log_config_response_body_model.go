// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeAuditLogConfigResponseBody
	GetAccessDeniedDetail() *string
	SetAuditLogStatus(v string) *DescribeAuditLogConfigResponseBody
	GetAuditLogStatus() *string
	SetDBClusterId(v string) *DescribeAuditLogConfigResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DescribeAuditLogConfigResponseBody
	GetRequestId() *string
}

type DescribeAuditLogConfigResponseBody struct {
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
	// The status of SQL audit. Valid values:
	//
	// 	- **on**: SQL audit is enabled.
	//
	// 	- **off**: SQL audit is disabled.
	//
	// example:
	//
	// on
	AuditLogStatus *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// am-t4nj8619bz2w3****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F0983B43-B2EC-536A-8791-142B5CF1E9B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAuditLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeAuditLogConfigResponseBody) GetAuditLogStatus() *string {
	return s.AuditLogStatus
}

func (s *DescribeAuditLogConfigResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAuditLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuditLogConfigResponseBody) SetAccessDeniedDetail(v string) *DescribeAuditLogConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetAuditLogStatus(v string) *DescribeAuditLogConfigResponseBody {
	s.AuditLogStatus = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetDBClusterId(v string) *DescribeAuditLogConfigResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetRequestId(v string) *DescribeAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
