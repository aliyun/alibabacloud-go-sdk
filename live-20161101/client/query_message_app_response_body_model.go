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
	// Request ID.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
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
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMessageAppResponseBodyResult struct {
	// Interactive Messages application list.
	AppList []*QueryMessageAppResponseBodyResultAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	// Indicates whether there is a next page. Valid values:
	//
	// - true: There is a next page.
	//
	// - false: There is no next page.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// Total number of query results.
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
	if s.AppList != nil {
		for _, item := range s.AppList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMessageAppResponseBodyResultAppList struct {
	// Application configuration.
	AppConfig map[string]*string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// Interactive Messages application ID.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Interactive Messages application name.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// UTC timestamp when the application was created.
	//
	// example:
	//
	// 502280113
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Extension field.
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Interactive Messages application status. A value of **1*	- indicates that the application status is Normal.
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
