// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMessageAppResponseBody
	GetRequestId() *string
	SetResult(v *ListMessageAppResponseBodyResult) *ListMessageAppResponseBody
	GetResult() *ListMessageAppResponseBodyResult
}

type ListMessageAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *ListMessageAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListMessageAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessageAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessageAppResponseBody) GetResult() *ListMessageAppResponseBodyResult {
	return s.Result
}

func (s *ListMessageAppResponseBody) SetRequestId(v string) *ListMessageAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageAppResponseBody) SetResult(v *ListMessageAppResponseBodyResult) *ListMessageAppResponseBody {
	s.Result = v
	return s
}

func (s *ListMessageAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMessageAppResponseBodyResult struct {
	// Details about the applications.
	AppList []*ListMessageAppResponseBodyResultAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
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
	// The total number of interactive message applications.
	//
	// example:
	//
	// 15
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMessageAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMessageAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMessageAppResponseBodyResult) GetAppList() []*ListMessageAppResponseBodyResultAppList {
	return s.AppList
}

func (s *ListMessageAppResponseBodyResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMessageAppResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *ListMessageAppResponseBodyResult) SetAppList(v []*ListMessageAppResponseBodyResultAppList) *ListMessageAppResponseBodyResult {
	s.AppList = v
	return s
}

func (s *ListMessageAppResponseBodyResult) SetHasMore(v bool) *ListMessageAppResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListMessageAppResponseBodyResult) SetTotal(v int32) *ListMessageAppResponseBodyResult {
	s.Total = &v
	return s
}

func (s *ListMessageAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListMessageAppResponseBodyResultAppList struct {
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

func (s ListMessageAppResponseBodyResultAppList) String() string {
	return dara.Prettify(s)
}

func (s ListMessageAppResponseBodyResultAppList) GoString() string {
	return s.String()
}

func (s *ListMessageAppResponseBodyResultAppList) GetAppConfig() map[string]*string {
	return s.AppConfig
}

func (s *ListMessageAppResponseBodyResultAppList) GetAppId() *string {
	return s.AppId
}

func (s *ListMessageAppResponseBodyResultAppList) GetAppName() *string {
	return s.AppName
}

func (s *ListMessageAppResponseBodyResultAppList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListMessageAppResponseBodyResultAppList) GetExtension() map[string]*string {
	return s.Extension
}

func (s *ListMessageAppResponseBodyResultAppList) GetStatus() *int32 {
	return s.Status
}

func (s *ListMessageAppResponseBodyResultAppList) SetAppConfig(v map[string]*string) *ListMessageAppResponseBodyResultAppList {
	s.AppConfig = v
	return s
}

func (s *ListMessageAppResponseBodyResultAppList) SetAppId(v string) *ListMessageAppResponseBodyResultAppList {
	s.AppId = &v
	return s
}

func (s *ListMessageAppResponseBodyResultAppList) SetAppName(v string) *ListMessageAppResponseBodyResultAppList {
	s.AppName = &v
	return s
}

func (s *ListMessageAppResponseBodyResultAppList) SetCreateTime(v int64) *ListMessageAppResponseBodyResultAppList {
	s.CreateTime = &v
	return s
}

func (s *ListMessageAppResponseBodyResultAppList) SetExtension(v map[string]*string) *ListMessageAppResponseBodyResultAppList {
	s.Extension = v
	return s
}

func (s *ListMessageAppResponseBodyResultAppList) SetStatus(v int32) *ListMessageAppResponseBodyResultAppList {
	s.Status = &v
	return s
}

func (s *ListMessageAppResponseBodyResultAppList) Validate() error {
	return dara.Validate(s)
}
