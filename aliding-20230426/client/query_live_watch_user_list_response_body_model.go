// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchUserListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrgUsesList(v []*QueryLiveWatchUserListResponseBodyOrgUsesList) *QueryLiveWatchUserListResponseBody
	GetOrgUsesList() []*QueryLiveWatchUserListResponseBodyOrgUsesList
	SetOutOrgUserList(v []*QueryLiveWatchUserListResponseBodyOutOrgUserList) *QueryLiveWatchUserListResponseBody
	GetOutOrgUserList() []*QueryLiveWatchUserListResponseBodyOutOrgUserList
	SetRequestId(v string) *QueryLiveWatchUserListResponseBody
	GetRequestId() *string
}

type QueryLiveWatchUserListResponseBody struct {
	OrgUsesList    []*QueryLiveWatchUserListResponseBodyOrgUsesList    `json:"orgUsesList,omitempty" xml:"orgUsesList,omitempty" type:"Repeated"`
	OutOrgUserList []*QueryLiveWatchUserListResponseBodyOutOrgUserList `json:"outOrgUserList,omitempty" xml:"outOrgUserList,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryLiveWatchUserListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListResponseBody) GetOrgUsesList() []*QueryLiveWatchUserListResponseBodyOrgUsesList {
	return s.OrgUsesList
}

func (s *QueryLiveWatchUserListResponseBody) GetOutOrgUserList() []*QueryLiveWatchUserListResponseBodyOutOrgUserList {
	return s.OutOrgUserList
}

func (s *QueryLiveWatchUserListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLiveWatchUserListResponseBody) SetOrgUsesList(v []*QueryLiveWatchUserListResponseBodyOrgUsesList) *QueryLiveWatchUserListResponseBody {
	s.OrgUsesList = v
	return s
}

func (s *QueryLiveWatchUserListResponseBody) SetOutOrgUserList(v []*QueryLiveWatchUserListResponseBodyOutOrgUserList) *QueryLiveWatchUserListResponseBody {
	s.OutOrgUserList = v
	return s
}

func (s *QueryLiveWatchUserListResponseBody) SetRequestId(v string) *QueryLiveWatchUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryLiveWatchUserListResponseBodyOrgUsesList struct {
	DeptName *string `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1234
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 19999
	WatchLiveTime *int64 `json:"WatchLiveTime,omitempty" xml:"WatchLiveTime,omitempty"`
	// example:
	//
	// 131312312
	WatchPlaybackTime *int64 `json:"WatchPlaybackTime,omitempty" xml:"WatchPlaybackTime,omitempty"`
	// example:
	//
	// 1323132
	WatchProgressMs *int64 `json:"WatchProgressMs,omitempty" xml:"WatchProgressMs,omitempty"`
}

func (s QueryLiveWatchUserListResponseBodyOrgUsesList) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListResponseBodyOrgUsesList) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) GetDeptName() *string {
	return s.DeptName
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) GetName() *string {
	return s.Name
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) GetUserId() *string {
	return s.UserId
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) GetWatchLiveTime() *int64 {
	return s.WatchLiveTime
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) GetWatchPlaybackTime() *int64 {
	return s.WatchPlaybackTime
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) GetWatchProgressMs() *int64 {
	return s.WatchProgressMs
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) SetDeptName(v string) *QueryLiveWatchUserListResponseBodyOrgUsesList {
	s.DeptName = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) SetName(v string) *QueryLiveWatchUserListResponseBodyOrgUsesList {
	s.Name = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) SetUserId(v string) *QueryLiveWatchUserListResponseBodyOrgUsesList {
	s.UserId = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) SetWatchLiveTime(v int64) *QueryLiveWatchUserListResponseBodyOrgUsesList {
	s.WatchLiveTime = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) SetWatchPlaybackTime(v int64) *QueryLiveWatchUserListResponseBodyOrgUsesList {
	s.WatchPlaybackTime = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) SetWatchProgressMs(v int64) *QueryLiveWatchUserListResponseBodyOrgUsesList {
	s.WatchProgressMs = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOrgUsesList) Validate() error {
	return dara.Validate(s)
}

type QueryLiveWatchUserListResponseBodyOutOrgUserList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12312312
	WatchLiveTime *int64 `json:"WatchLiveTime,omitempty" xml:"WatchLiveTime,omitempty"`
	// example:
	//
	// 21313131
	WatchPlaybackTime *int64 `json:"WatchPlaybackTime,omitempty" xml:"WatchPlaybackTime,omitempty"`
	// example:
	//
	// 123131
	WatchProgressMs *int64 `json:"WatchProgressMs,omitempty" xml:"WatchProgressMs,omitempty"`
}

func (s QueryLiveWatchUserListResponseBodyOutOrgUserList) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListResponseBodyOutOrgUserList) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListResponseBodyOutOrgUserList) GetName() *string {
	return s.Name
}

func (s *QueryLiveWatchUserListResponseBodyOutOrgUserList) GetWatchLiveTime() *int64 {
	return s.WatchLiveTime
}

func (s *QueryLiveWatchUserListResponseBodyOutOrgUserList) GetWatchPlaybackTime() *int64 {
	return s.WatchPlaybackTime
}

func (s *QueryLiveWatchUserListResponseBodyOutOrgUserList) GetWatchProgressMs() *int64 {
	return s.WatchProgressMs
}

func (s *QueryLiveWatchUserListResponseBodyOutOrgUserList) SetName(v string) *QueryLiveWatchUserListResponseBodyOutOrgUserList {
	s.Name = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOutOrgUserList) SetWatchLiveTime(v int64) *QueryLiveWatchUserListResponseBodyOutOrgUserList {
	s.WatchLiveTime = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOutOrgUserList) SetWatchPlaybackTime(v int64) *QueryLiveWatchUserListResponseBodyOutOrgUserList {
	s.WatchPlaybackTime = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOutOrgUserList) SetWatchProgressMs(v int64) *QueryLiveWatchUserListResponseBodyOutOrgUserList {
	s.WatchProgressMs = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyOutOrgUserList) Validate() error {
	return dara.Validate(s)
}
