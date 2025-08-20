// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkAppTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeSparkAppTypeResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *DescribeSparkAppTypeResponseBody
	GetRequestId() *string
	SetType(v string) *DescribeSparkAppTypeResponseBody
	GetType() *string
}

type DescribeSparkAppTypeResponseBody struct {
	// The detailed reason why the access was denied.
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
	// The request ID.
	//
	// example:
	//
	// 596AF63B-8798-501E-BA06-CD2184D48A35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the Spark application. Valid values:
	//
	// 	- BATCH
	//
	// 	- SQLENGINE
	//
	// example:
	//
	// Batch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSparkAppTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAppTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkAppTypeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeSparkAppTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSparkAppTypeResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeSparkAppTypeResponseBody) SetAccessDeniedDetail(v string) *DescribeSparkAppTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeSparkAppTypeResponseBody) SetRequestId(v string) *DescribeSparkAppTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkAppTypeResponseBody) SetType(v string) *DescribeSparkAppTypeResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeSparkAppTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
