// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsByApiProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedApps(v *DescribeAppsByApiProductResponseBodyAuthorizedApps) *DescribeAppsByApiProductResponseBody
	GetAuthorizedApps() *DescribeAppsByApiProductResponseBodyAuthorizedApps
	SetPageNumber(v int32) *DescribeAppsByApiProductResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAppsByApiProductResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAppsByApiProductResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAppsByApiProductResponseBody
	GetTotalCount() *int32
}

type DescribeAppsByApiProductResponseBody struct {
	// The information about authorized applications.
	AuthorizedApps *DescribeAppsByApiProductResponseBodyAuthorizedApps `json:"AuthorizedApps,omitempty" xml:"AuthorizedApps,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC866798-62D3-52F4-8AB5-CA149A53984F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppsByApiProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsByApiProductResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiProductResponseBody) GetAuthorizedApps() *DescribeAppsByApiProductResponseBodyAuthorizedApps {
	return s.AuthorizedApps
}

func (s *DescribeAppsByApiProductResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAppsByApiProductResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppsByApiProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppsByApiProductResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAppsByApiProductResponseBody) SetAuthorizedApps(v *DescribeAppsByApiProductResponseBodyAuthorizedApps) *DescribeAppsByApiProductResponseBody {
	s.AuthorizedApps = v
	return s
}

func (s *DescribeAppsByApiProductResponseBody) SetPageNumber(v int32) *DescribeAppsByApiProductResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBody) SetPageSize(v int32) *DescribeAppsByApiProductResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBody) SetRequestId(v string) *DescribeAppsByApiProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBody) SetTotalCount(v int32) *DescribeAppsByApiProductResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsByApiProductResponseBodyAuthorizedApps struct {
	AuthorizedApp []*DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp `json:"AuthorizedApp,omitempty" xml:"AuthorizedApp,omitempty" type:"Repeated"`
}

func (s DescribeAppsByApiProductResponseBodyAuthorizedApps) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsByApiProductResponseBodyAuthorizedApps) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedApps) GetAuthorizedApp() []*DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp {
	return s.AuthorizedApp
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedApps) SetAuthorizedApp(v []*DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) *DescribeAppsByApiProductResponseBodyAuthorizedApps {
	s.AuthorizedApp = v
	return s
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedApps) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp struct {
	// The application ID.
	//
	// example:
	//
	// 110982419
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// APP_02580_DEV
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The expiration time of the authorization. The time is in GMT. An empty value indicates that the authorization does not expire.
	//
	// example:
	//
	// 2023-06-17T03:41:53Z
	AuthValidTime *string `json:"AuthValidTime,omitempty" xml:"AuthValidTime,omitempty"`
	// The time when the authorization was created. The time is in GMT.
	//
	// example:
	//
	// 2016-07-21T06:17:20Z
	AuthorizedTime *string `json:"AuthorizedTime,omitempty" xml:"AuthorizedTime,omitempty"`
	// The authorization description.
	//
	// example:
	//
	// Test share with nsc qiming
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information.
	//
	// example:
	//
	// extra info
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
}

func (s DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) GetAuthValidTime() *string {
	return s.AuthValidTime
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) GetAuthorizedTime() *string {
	return s.AuthorizedTime
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) GetExtend() *string {
	return s.Extend
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) SetAppId(v int64) *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp {
	s.AppId = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) SetAppName(v string) *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp {
	s.AppName = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) SetAuthValidTime(v string) *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp {
	s.AuthValidTime = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) SetAuthorizedTime(v string) *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp {
	s.AuthorizedTime = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) SetDescription(v string) *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp {
	s.Description = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) SetExtend(v string) *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp {
	s.Extend = &v
	return s
}

func (s *DescribeAppsByApiProductResponseBodyAuthorizedAppsAuthorizedApp) Validate() error {
	return dara.Validate(s)
}
