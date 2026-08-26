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
	// Request ID.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
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
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessageAppResponseBodyResult struct {
	// Application list.
	AppList []*ListMessageAppResponseBodyResultAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
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
	// Total number of interactive message applications.
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

type ListMessageAppResponseBodyResultAppList struct {
	// Application configuration.
	AppConfig map[string]*string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// Interactive message application ID.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Interactive message application name.
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
	// Interactive message application status. A value of **1*	- indicates that the application status is Normal.
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
