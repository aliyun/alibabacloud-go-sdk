// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePerformanceViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribePerformanceViewsResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *DescribePerformanceViewsResponseBody
	GetRequestId() *string
	SetViews(v []*DescribePerformanceViewsResponseBodyViews) *DescribePerformanceViewsResponseBody
	GetViews() []*DescribePerformanceViewsResponseBodyViews
}

type DescribePerformanceViewsResponseBody struct {
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
	// The request ID.
	//
	// example:
	//
	// 3A8F6106-6AFD-5A34-9C80-8DE2C42D06E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// the list of view.
	Views []*DescribePerformanceViewsResponseBodyViews `json:"Views,omitempty" xml:"Views,omitempty" type:"Repeated"`
}

func (s DescribePerformanceViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribePerformanceViewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePerformanceViewsResponseBody) GetViews() []*DescribePerformanceViewsResponseBodyViews {
	return s.Views
}

func (s *DescribePerformanceViewsResponseBody) SetAccessDeniedDetail(v string) *DescribePerformanceViewsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePerformanceViewsResponseBody) SetRequestId(v string) *DescribePerformanceViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePerformanceViewsResponseBody) SetViews(v []*DescribePerformanceViewsResponseBodyViews) *DescribePerformanceViewsResponseBody {
	s.Views = v
	return s
}

func (s *DescribePerformanceViewsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePerformanceViewsResponseBodyViews struct {
	// The time when created.
	//
	// example:
	//
	// 2024-06-18T07:06:53.000+00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when updated.
	//
	// example:
	//
	// 2024-06-18T07:07:32.000+00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The name of the view.
	//
	// example:
	//
	// Basic
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s DescribePerformanceViewsResponseBodyViews) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewsResponseBodyViews) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewsResponseBodyViews) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePerformanceViewsResponseBodyViews) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribePerformanceViewsResponseBodyViews) GetViewName() *string {
	return s.ViewName
}

func (s *DescribePerformanceViewsResponseBodyViews) SetCreateTime(v string) *DescribePerformanceViewsResponseBodyViews {
	s.CreateTime = &v
	return s
}

func (s *DescribePerformanceViewsResponseBodyViews) SetUpdateTime(v string) *DescribePerformanceViewsResponseBodyViews {
	s.UpdateTime = &v
	return s
}

func (s *DescribePerformanceViewsResponseBodyViews) SetViewName(v string) *DescribePerformanceViewsResponseBodyViews {
	s.ViewName = &v
	return s
}

func (s *DescribePerformanceViewsResponseBodyViews) Validate() error {
	return dara.Validate(s)
}
