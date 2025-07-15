// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMessageAppResponseBody
	GetRequestId() *string
	SetResult(v []*QueryMessageAppResponseBodyResult) *QueryMessageAppResponseBody
	GetResult() []*QueryMessageAppResponseBodyResult
}

type QueryMessageAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*QueryMessageAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryMessageAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageAppResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMessageAppResponseBody) GetResult() []*QueryMessageAppResponseBodyResult {
	return s.Result
}

func (s *QueryMessageAppResponseBody) SetRequestId(v string) *QueryMessageAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessageAppResponseBody) SetResult(v []*QueryMessageAppResponseBodyResult) *QueryMessageAppResponseBody {
	s.Result = v
	return s
}

func (s *QueryMessageAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMessageAppResponseBodyResult struct {
	// Details about the interactive messaging applications.
	AppList []*QueryMessageAppResponseBodyResultAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	// Indicates whether the current page is followed by another page. Valid values:
	//
	// 	- true: The current page is followed by another page.
	//
	// 	- false: The current page is not followed by another page.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The total number of applications returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryMessageAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMessageAppResponseBodyResult) GetAppList() []*QueryMessageAppResponseBodyResultAppList {
	return s.AppList
}

func (s *QueryMessageAppResponseBodyResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *QueryMessageAppResponseBodyResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryMessageAppResponseBodyResult) SetAppList(v []*QueryMessageAppResponseBodyResultAppList) *QueryMessageAppResponseBodyResult {
	s.AppList = v
	return s
}

func (s *QueryMessageAppResponseBodyResult) SetHasMore(v bool) *QueryMessageAppResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryMessageAppResponseBodyResult) SetTotalCount(v int32) *QueryMessageAppResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *QueryMessageAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type QueryMessageAppResponseBodyResultAppList struct {
	// The configurations of the application.
	AppConfig map[string]*string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The ID of the interactive messaging application.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the interactive messaging application.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the interactive messaging application was created. The time is displayed in UTC.
	//
	// example:
	//
	// 502280113
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The extended field.
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The status of the interactive message application. A value of **1*	- indicates that the application is normal.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryMessageAppResponseBodyResultAppList) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageAppResponseBodyResultAppList) GoString() string {
	return s.String()
}

func (s *QueryMessageAppResponseBodyResultAppList) GetAppConfig() map[string]*string {
	return s.AppConfig
}

func (s *QueryMessageAppResponseBodyResultAppList) GetAppId() *string {
	return s.AppId
}

func (s *QueryMessageAppResponseBodyResultAppList) GetAppName() *string {
	return s.AppName
}

func (s *QueryMessageAppResponseBodyResultAppList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryMessageAppResponseBodyResultAppList) GetExtension() map[string]*string {
	return s.Extension
}

func (s *QueryMessageAppResponseBodyResultAppList) GetStatus() *int32 {
	return s.Status
}

func (s *QueryMessageAppResponseBodyResultAppList) SetAppConfig(v map[string]*string) *QueryMessageAppResponseBodyResultAppList {
	s.AppConfig = v
	return s
}

func (s *QueryMessageAppResponseBodyResultAppList) SetAppId(v string) *QueryMessageAppResponseBodyResultAppList {
	s.AppId = &v
	return s
}

func (s *QueryMessageAppResponseBodyResultAppList) SetAppName(v string) *QueryMessageAppResponseBodyResultAppList {
	s.AppName = &v
	return s
}

func (s *QueryMessageAppResponseBodyResultAppList) SetCreateTime(v int64) *QueryMessageAppResponseBodyResultAppList {
	s.CreateTime = &v
	return s
}

func (s *QueryMessageAppResponseBodyResultAppList) SetExtension(v map[string]*string) *QueryMessageAppResponseBodyResultAppList {
	s.Extension = v
	return s
}

func (s *QueryMessageAppResponseBodyResultAppList) SetStatus(v int32) *QueryMessageAppResponseBodyResultAppList {
	s.Status = &v
	return s
}

func (s *QueryMessageAppResponseBodyResultAppList) Validate() error {
	return dara.Validate(s)
}
