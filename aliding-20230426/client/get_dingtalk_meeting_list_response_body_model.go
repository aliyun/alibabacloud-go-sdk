// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetDingtalkMeetingListResponseBody
	GetCurrentPage() *int32
	SetData(v []*GetDingtalkMeetingListResponseBodyData) *GetDingtalkMeetingListResponseBody
	GetData() []*GetDingtalkMeetingListResponseBodyData
	SetRequestId(v string) *GetDingtalkMeetingListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetDingtalkMeetingListResponseBody
	GetTotalCount() *int32
	SetVendorRequestId(v string) *GetDingtalkMeetingListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDingtalkMeetingListResponseBody
	GetVendorType() *string
}

type GetDingtalkMeetingListResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// []
	Data []*GetDingtalkMeetingListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetDingtalkMeetingListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDingtalkMeetingListResponseBody) GetData() []*GetDingtalkMeetingListResponseBodyData {
	return s.Data
}

func (s *GetDingtalkMeetingListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDingtalkMeetingListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetDingtalkMeetingListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDingtalkMeetingListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDingtalkMeetingListResponseBody) SetCurrentPage(v int32) *GetDingtalkMeetingListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBody) SetData(v []*GetDingtalkMeetingListResponseBodyData) *GetDingtalkMeetingListResponseBody {
	s.Data = v
	return s
}

func (s *GetDingtalkMeetingListResponseBody) SetRequestId(v string) *GetDingtalkMeetingListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBody) SetTotalCount(v int32) *GetDingtalkMeetingListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBody) SetVendorRequestId(v string) *GetDingtalkMeetingListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBody) SetVendorType(v string) *GetDingtalkMeetingListResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDingtalkMeetingListResponseBodyData struct {
	// example:
	//
	// cluster-1
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 012345
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 张三
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// example:
	//
	// 3423423
	CreatorWorkNo *string `json:"creatorWorkNo,omitempty" xml:"creatorWorkNo,omitempty"`
	// example:
	//
	// 技术部
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// true
	EnableQualityMonitor *bool `json:"enableQualityMonitor,omitempty" xml:"enableQualityMonitor,omitempty"`
	// example:
	//
	// 1638363600000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// free
	FreeType *string `json:"freeType,omitempty" xml:"freeType,omitempty"`
	// example:
	//
	// meeting
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 1638360000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 3600
	TimeLength *int64 `json:"timeLength,omitempty" xml:"timeLength,omitempty"`
	// example:
	//
	// 技术分享会
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 10
	UserCount *int32 `json:"userCount,omitempty" xml:"userCount,omitempty"`
}

func (s GetDingtalkMeetingListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingListResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetDingtalkMeetingListResponseBodyData) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingListResponseBodyData) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetDingtalkMeetingListResponseBodyData) GetCreatorNick() *string {
	return s.CreatorNick
}

func (s *GetDingtalkMeetingListResponseBodyData) GetCreatorWorkNo() *string {
	return s.CreatorWorkNo
}

func (s *GetDingtalkMeetingListResponseBodyData) GetDeptName() *string {
	return s.DeptName
}

func (s *GetDingtalkMeetingListResponseBodyData) GetEnableQualityMonitor() *bool {
	return s.EnableQualityMonitor
}

func (s *GetDingtalkMeetingListResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDingtalkMeetingListResponseBodyData) GetFreeType() *string {
	return s.FreeType
}

func (s *GetDingtalkMeetingListResponseBodyData) GetScene() *string {
	return s.Scene
}

func (s *GetDingtalkMeetingListResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDingtalkMeetingListResponseBodyData) GetTimeLength() *int64 {
	return s.TimeLength
}

func (s *GetDingtalkMeetingListResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetDingtalkMeetingListResponseBodyData) GetUserCount() *int32 {
	return s.UserCount
}

func (s *GetDingtalkMeetingListResponseBodyData) SetClusterName(v string) *GetDingtalkMeetingListResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetConferenceId(v string) *GetDingtalkMeetingListResponseBodyData {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetCreatorId(v string) *GetDingtalkMeetingListResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetCreatorNick(v string) *GetDingtalkMeetingListResponseBodyData {
	s.CreatorNick = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetCreatorWorkNo(v string) *GetDingtalkMeetingListResponseBodyData {
	s.CreatorWorkNo = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetDeptName(v string) *GetDingtalkMeetingListResponseBodyData {
	s.DeptName = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetEnableQualityMonitor(v bool) *GetDingtalkMeetingListResponseBodyData {
	s.EnableQualityMonitor = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetEndTime(v int64) *GetDingtalkMeetingListResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetFreeType(v string) *GetDingtalkMeetingListResponseBodyData {
	s.FreeType = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetScene(v string) *GetDingtalkMeetingListResponseBodyData {
	s.Scene = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetStartTime(v int64) *GetDingtalkMeetingListResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetTimeLength(v int64) *GetDingtalkMeetingListResponseBodyData {
	s.TimeLength = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetTitle(v string) *GetDingtalkMeetingListResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) SetUserCount(v int32) *GetDingtalkMeetingListResponseBodyData {
	s.UserCount = &v
	return s
}

func (s *GetDingtalkMeetingListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
