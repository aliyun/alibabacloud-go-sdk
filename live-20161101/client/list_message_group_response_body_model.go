// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMessageGroupResponseBody
	GetRequestId() *string
	SetResult(v *ListMessageGroupResponseBodyResult) *ListMessageGroupResponseBody
	GetResult() *ListMessageGroupResponseBodyResult
}

type ListMessageGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *ListMessageGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessageGroupResponseBody) GetResult() *ListMessageGroupResponseBodyResult {
	return s.Result
}

func (s *ListMessageGroupResponseBody) SetRequestId(v string) *ListMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageGroupResponseBody) SetResult(v *ListMessageGroupResponseBodyResult) *ListMessageGroupResponseBody {
	s.Result = v
	return s
}

func (s *ListMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMessageGroupResponseBodyResult struct {
	// The list of message groups.
	GroupList []*ListMessageGroupResponseBodyResultGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
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
	// The total number of message groups.
	//
	// example:
	//
	// 5
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMessageGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMessageGroupResponseBodyResult) GetGroupList() []*ListMessageGroupResponseBodyResultGroupList {
	return s.GroupList
}

func (s *ListMessageGroupResponseBodyResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMessageGroupResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *ListMessageGroupResponseBodyResult) SetGroupList(v []*ListMessageGroupResponseBodyResultGroupList) *ListMessageGroupResponseBodyResult {
	s.GroupList = v
	return s
}

func (s *ListMessageGroupResponseBodyResult) SetHasMore(v bool) *ListMessageGroupResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListMessageGroupResponseBodyResult) SetTotal(v int32) *ListMessageGroupResponseBodyResult {
	s.Total = &v
	return s
}

func (s *ListMessageGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListMessageGroupResponseBodyResultGroupList struct {
	// The ID of the interactive messaging application.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the message group was created. The time is displayed in UTC.
	//
	// example:
	//
	// 1502280113
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// as****hs
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The extended field.
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The ID of the message group.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The status of the message group. The default value is **1**, which indicates that the status of the message group is normal.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMessageGroupResponseBodyResultGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupResponseBodyResultGroupList) GoString() string {
	return s.String()
}

func (s *ListMessageGroupResponseBodyResultGroupList) GetAppId() *string {
	return s.AppId
}

func (s *ListMessageGroupResponseBodyResultGroupList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListMessageGroupResponseBodyResultGroupList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListMessageGroupResponseBodyResultGroupList) GetExtension() map[string]*string {
	return s.Extension
}

func (s *ListMessageGroupResponseBodyResultGroupList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListMessageGroupResponseBodyResultGroupList) GetStatus() *int32 {
	return s.Status
}

func (s *ListMessageGroupResponseBodyResultGroupList) SetAppId(v string) *ListMessageGroupResponseBodyResultGroupList {
	s.AppId = &v
	return s
}

func (s *ListMessageGroupResponseBodyResultGroupList) SetCreateTime(v int64) *ListMessageGroupResponseBodyResultGroupList {
	s.CreateTime = &v
	return s
}

func (s *ListMessageGroupResponseBodyResultGroupList) SetCreatorId(v string) *ListMessageGroupResponseBodyResultGroupList {
	s.CreatorId = &v
	return s
}

func (s *ListMessageGroupResponseBodyResultGroupList) SetExtension(v map[string]*string) *ListMessageGroupResponseBodyResultGroupList {
	s.Extension = v
	return s
}

func (s *ListMessageGroupResponseBodyResultGroupList) SetGroupId(v string) *ListMessageGroupResponseBodyResultGroupList {
	s.GroupId = &v
	return s
}

func (s *ListMessageGroupResponseBodyResultGroupList) SetStatus(v int32) *ListMessageGroupResponseBodyResultGroupList {
	s.Status = &v
	return s
}

func (s *ListMessageGroupResponseBodyResultGroupList) Validate() error {
	return dara.Validate(s)
}
