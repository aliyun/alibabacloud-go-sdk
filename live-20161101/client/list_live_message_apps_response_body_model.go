// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppList(v []*ListLiveMessageAppsResponseBodyAppList) *ListLiveMessageAppsResponseBody
	GetAppList() []*ListLiveMessageAppsResponseBodyAppList
	SetHasMore(v bool) *ListLiveMessageAppsResponseBody
	GetHasMore() *bool
	SetNextPageToken(v int64) *ListLiveMessageAppsResponseBody
	GetNextPageToken() *int64
	SetRequestId(v string) *ListLiveMessageAppsResponseBody
	GetRequestId() *string
}

type ListLiveMessageAppsResponseBody struct {
	// The interactive messaging applications.
	AppList []*ListLiveMessageAppsResponseBodyAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	// Indicates whether the current page is followed by a page.
	//
	// example:
	//
	// true
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The starting page number for the next query. This parameter is returned only if the value of HasMore is true.
	//
	// example:
	//
	// 1
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B8EB4994-1368-1458-B9F3-5B88D76D734C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLiveMessageAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveMessageAppsResponseBody) GetAppList() []*ListLiveMessageAppsResponseBodyAppList {
	return s.AppList
}

func (s *ListLiveMessageAppsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListLiveMessageAppsResponseBody) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListLiveMessageAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveMessageAppsResponseBody) SetAppList(v []*ListLiveMessageAppsResponseBodyAppList) *ListLiveMessageAppsResponseBody {
	s.AppList = v
	return s
}

func (s *ListLiveMessageAppsResponseBody) SetHasMore(v bool) *ListLiveMessageAppsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListLiveMessageAppsResponseBody) SetNextPageToken(v int64) *ListLiveMessageAppsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListLiveMessageAppsResponseBody) SetRequestId(v string) *ListLiveMessageAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveMessageAppsResponseBody) Validate() error {
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

type ListLiveMessageAppsResponseBodyAppList struct {
	// The ID of the interactive messaging application queried.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The AppKey of the interactive messaging application. It is used to authorize operations related to the application ID.
	//
	// example:
	//
	// **********************************
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The signature of the interactive messaging application. It is required by the interactive messaging SDK.
	//
	// example:
	//
	// **************************************************************************
	AppSign *string `json:"AppSign,omitempty" xml:"AppSign,omitempty"`
	// The time when the application was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1698305471
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The live center.
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// Indicates whether the interactive messaging application is disabled.
	//
	// example:
	//
	// false
	Disable *string `json:"Disable,omitempty" xml:"Disable,omitempty"`
	// The time when the application was last modified. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1698305471
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The retention period of group messages in the application. Valid values:
	//
	// 	- 0 (default): 30 days
	//
	// 	- 1: 90 days
	//
	// 	- 2: 180 days
	//
	// example:
	//
	// 1
	MsgLifeCycle *int32 `json:"MsgLifeCycle,omitempty" xml:"MsgLifeCycle,omitempty"`
}

func (s ListLiveMessageAppsResponseBodyAppList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageAppsResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *ListLiveMessageAppsResponseBodyAppList) GetAppId() *string {
	return s.AppId
}

func (s *ListLiveMessageAppsResponseBodyAppList) GetAppKey() *string {
	return s.AppKey
}

func (s *ListLiveMessageAppsResponseBodyAppList) GetAppName() *string {
	return s.AppName
}

func (s *ListLiveMessageAppsResponseBodyAppList) GetAppSign() *string {
	return s.AppSign
}

func (s *ListLiveMessageAppsResponseBodyAppList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListLiveMessageAppsResponseBodyAppList) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ListLiveMessageAppsResponseBodyAppList) GetDisable() *string {
	return s.Disable
}

func (s *ListLiveMessageAppsResponseBodyAppList) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListLiveMessageAppsResponseBodyAppList) GetMsgLifeCycle() *int32 {
	return s.MsgLifeCycle
}

func (s *ListLiveMessageAppsResponseBodyAppList) SetAppId(v string) *ListLiveMessageAppsResponseBodyAppList {
	s.AppId = &v
	return s
}

func (s *ListLiveMessageAppsResponseBodyAppList) SetAppKey(v string) *ListLiveMessageAppsResponseBodyAppList {
	s.AppKey = &v
	return s
}

func (s *ListLiveMessageAppsResponseBodyAppList) SetAppName(v string) *ListLiveMessageAppsResponseBodyAppList {
	s.AppName = &v
	return s
}

func (s *ListLiveMessageAppsResponseBodyAppList) SetAppSign(v string) *ListLiveMessageAppsResponseBodyAppList {
	s.AppSign = &v
	return s
}

func (s *ListLiveMessageAppsResponseBodyAppList) SetCreateTime(v int64) *ListLiveMessageAppsResponseBodyAppList {
	s.CreateTime = &v
	return s
}

func (s *ListLiveMessageAppsResponseBodyAppList) SetDataCenter(v string) *ListLiveMessageAppsResponseBodyAppList {
	s.DataCenter = &v
	return s
}

func (s *ListLiveMessageAppsResponseBodyAppList) SetDisable(v string) *ListLiveMessageAppsResponseBodyAppList {
	s.Disable = &v
	return s
}

func (s *ListLiveMessageAppsResponseBodyAppList) SetModifyTime(v int64) *ListLiveMessageAppsResponseBodyAppList {
	s.ModifyTime = &v
	return s
}

func (s *ListLiveMessageAppsResponseBodyAppList) SetMsgLifeCycle(v int32) *ListLiveMessageAppsResponseBodyAppList {
	s.MsgLifeCycle = &v
	return s
}

func (s *ListLiveMessageAppsResponseBodyAppList) Validate() error {
	return dara.Validate(s)
}
