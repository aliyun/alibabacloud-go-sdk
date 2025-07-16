// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgLiveListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOrgLiveListResponseBody
	GetRequestId() *string
	SetResult(v *GetOrgLiveListResponseBodyResult) *GetOrgLiveListResponseBody
	GetResult() *GetOrgLiveListResponseBodyResult
	SetVendorRequestId(v string) *GetOrgLiveListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetOrgLiveListResponseBody
	GetVendorType() *string
}

type GetOrgLiveListResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetOrgLiveListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetOrgLiveListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrgLiveListResponseBody) GetResult() *GetOrgLiveListResponseBodyResult {
	return s.Result
}

func (s *GetOrgLiveListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetOrgLiveListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetOrgLiveListResponseBody) SetRequestId(v string) *GetOrgLiveListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrgLiveListResponseBody) SetResult(v *GetOrgLiveListResponseBodyResult) *GetOrgLiveListResponseBody {
	s.Result = v
	return s
}

func (s *GetOrgLiveListResponseBody) SetVendorRequestId(v string) *GetOrgLiveListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetOrgLiveListResponseBody) SetVendorType(v string) *GetOrgLiveListResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetOrgLiveListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOrgLiveListResponseBodyResult struct {
	// example:
	//
	// 新建的直播列表
	NewLive *GetOrgLiveListResponseBodyResultNewLive `json:"NewLive,omitempty" xml:"NewLive,omitempty" type:"Struct"`
	// example:
	//
	// 修改的直播列表
	UpdateLive *GetOrgLiveListResponseBodyResultUpdateLive `json:"UpdateLive,omitempty" xml:"UpdateLive,omitempty" type:"Struct"`
}

func (s GetOrgLiveListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResult) GetNewLive() *GetOrgLiveListResponseBodyResultNewLive {
	return s.NewLive
}

func (s *GetOrgLiveListResponseBodyResult) GetUpdateLive() *GetOrgLiveListResponseBodyResultUpdateLive {
	return s.UpdateLive
}

func (s *GetOrgLiveListResponseBodyResult) SetNewLive(v *GetOrgLiveListResponseBodyResultNewLive) *GetOrgLiveListResponseBodyResult {
	s.NewLive = v
	return s
}

func (s *GetOrgLiveListResponseBodyResult) SetUpdateLive(v *GetOrgLiveListResponseBodyResultUpdateLive) *GetOrgLiveListResponseBodyResult {
	s.UpdateLive = v
	return s
}

func (s *GetOrgLiveListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetOrgLiveListResponseBodyResultNewLive struct {
	// example:
	//
	// true
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// []
	LiveList []*GetOrgLiveListResponseBodyResultNewLiveLiveList `json:"LiveList,omitempty" xml:"LiveList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetOrgLiveListResponseBodyResultNewLive) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResultNewLive) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResultNewLive) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetOrgLiveListResponseBodyResultNewLive) GetLiveList() []*GetOrgLiveListResponseBodyResultNewLiveLiveList {
	return s.LiveList
}

func (s *GetOrgLiveListResponseBodyResultNewLive) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetOrgLiveListResponseBodyResultNewLive) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetOrgLiveListResponseBodyResultNewLive) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetHasMore(v bool) *GetOrgLiveListResponseBodyResultNewLive {
	s.HasMore = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetLiveList(v []*GetOrgLiveListResponseBodyResultNewLiveLiveList) *GetOrgLiveListResponseBodyResultNewLive {
	s.LiveList = v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetPageNumber(v int64) *GetOrgLiveListResponseBodyResultNewLive {
	s.PageNumber = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetPageSize(v int64) *GetOrgLiveListResponseBodyResultNewLive {
	s.PageSize = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetTotalCount(v int64) *GetOrgLiveListResponseBodyResultNewLive {
	s.TotalCount = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLive) Validate() error {
	return dara.Validate(s)
}

type GetOrgLiveListResponseBodyResultNewLiveLiveList struct {
	// example:
	//
	// nickName
	AnchorNickname *string `json:"AnchorNickname,omitempty" xml:"AnchorNickname,omitempty"`
	AnchorUnionId  *string `json:"AnchorUnionId,omitempty" xml:"AnchorUnionId,omitempty"`
	// example:
	//
	// ersqqdddf
	AnchorUserId *string `json:"AnchorUserId,omitempty" xml:"AnchorUserId,omitempty"`
	// example:
	//
	// 1398324600000
	LiveEndTime *int64 `json:"LiveEndTime,omitempty" xml:"LiveEndTime,omitempty"`
	// example:
	//
	// 1398321600000
	LiveStartTime *int64 `json:"LiveStartTime,omitempty" xml:"LiveStartTime,omitempty"`
	// example:
	//
	// 4d38xxxxx
	LiveUuid *string `json:"LiveUuid,omitempty" xml:"LiveUuid,omitempty"`
	// example:
	//
	// 群OpenConversationId
	ShareOpenConversationIds []*string `json:"ShareOpenConversationIds,omitempty" xml:"ShareOpenConversationIds,omitempty" type:"Repeated"`
	// example:
	//
	// 直播标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetOrgLiveListResponseBodyResultNewLiveLiveList) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResultNewLiveLiveList) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) GetAnchorNickname() *string {
	return s.AnchorNickname
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) GetAnchorUnionId() *string {
	return s.AnchorUnionId
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) GetAnchorUserId() *string {
	return s.AnchorUserId
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) GetLiveEndTime() *int64 {
	return s.LiveEndTime
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) GetLiveStartTime() *int64 {
	return s.LiveStartTime
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) GetLiveUuid() *string {
	return s.LiveUuid
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) GetShareOpenConversationIds() []*string {
	return s.ShareOpenConversationIds
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) GetTitle() *string {
	return s.Title
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetAnchorNickname(v string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.AnchorNickname = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetAnchorUnionId(v string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.AnchorUnionId = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetAnchorUserId(v string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.AnchorUserId = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetLiveEndTime(v int64) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.LiveEndTime = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetLiveStartTime(v int64) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.LiveStartTime = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetLiveUuid(v string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.LiveUuid = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetShareOpenConversationIds(v []*string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.ShareOpenConversationIds = v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetTitle(v string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.Title = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) Validate() error {
	return dara.Validate(s)
}

type GetOrgLiveListResponseBodyResultUpdateLive struct {
	// example:
	//
	// true
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// []
	LiveList []*GetOrgLiveListResponseBodyResultUpdateLiveLiveList `json:"LiveList,omitempty" xml:"LiveList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetOrgLiveListResponseBodyResultUpdateLive) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResultUpdateLive) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) GetLiveList() []*GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	return s.LiveList
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetHasMore(v bool) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.HasMore = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetLiveList(v []*GetOrgLiveListResponseBodyResultUpdateLiveLiveList) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.LiveList = v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetPageNumber(v int64) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.PageNumber = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetPageSize(v int64) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.PageSize = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetTotalCount(v int64) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.TotalCount = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) Validate() error {
	return dara.Validate(s)
}

type GetOrgLiveListResponseBodyResultUpdateLiveLiveList struct {
	// example:
	//
	// nickName
	AnchorNickname *string `json:"AnchorNickname,omitempty" xml:"AnchorNickname,omitempty"`
	AnchorUnionId  *string `json:"AnchorUnionId,omitempty" xml:"AnchorUnionId,omitempty"`
	// example:
	//
	// 012345
	AnchorUserId *string `json:"AnchorUserId,omitempty" xml:"AnchorUserId,omitempty"`
	// example:
	//
	// 1398324600000
	LiveEndTime *int64 `json:"LiveEndTime,omitempty" xml:"LiveEndTime,omitempty"`
	// example:
	//
	// 1398321600000
	LiveStartTime *int64 `json:"LiveStartTime,omitempty" xml:"LiveStartTime,omitempty"`
	// example:
	//
	// 4d38xxxxx
	LiveUuid *string `json:"LiveUuid,omitempty" xml:"LiveUuid,omitempty"`
	// example:
	//
	// 直播标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetOrgLiveListResponseBodyResultUpdateLiveLiveList) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResultUpdateLiveLiveList) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) GetAnchorNickname() *string {
	return s.AnchorNickname
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) GetAnchorUnionId() *string {
	return s.AnchorUnionId
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) GetAnchorUserId() *string {
	return s.AnchorUserId
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) GetLiveEndTime() *int64 {
	return s.LiveEndTime
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) GetLiveStartTime() *int64 {
	return s.LiveStartTime
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) GetLiveUuid() *string {
	return s.LiveUuid
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) GetTitle() *string {
	return s.Title
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetAnchorNickname(v string) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.AnchorNickname = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetAnchorUnionId(v string) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.AnchorUnionId = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetAnchorUserId(v string) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.AnchorUserId = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetLiveEndTime(v int64) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.LiveEndTime = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetLiveStartTime(v int64) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.LiveStartTime = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetLiveUuid(v string) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.LiveUuid = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetTitle(v string) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.Title = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) Validate() error {
	return dara.Validate(s)
}
