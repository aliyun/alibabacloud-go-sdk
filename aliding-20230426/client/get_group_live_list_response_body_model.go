// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupLiveListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetGroupLiveListResponseBody
	GetRequestId() *string
	SetResult(v *GetGroupLiveListResponseBodyResult) *GetGroupLiveListResponseBody
	GetResult() *GetGroupLiveListResponseBodyResult
	SetVendorRequestId(v string) *GetGroupLiveListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetGroupLiveListResponseBody
	GetVendorType() *string
}

type GetGroupLiveListResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetGroupLiveListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetGroupLiveListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupLiveListResponseBody) GetResult() *GetGroupLiveListResponseBodyResult {
	return s.Result
}

func (s *GetGroupLiveListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetGroupLiveListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetGroupLiveListResponseBody) SetRequestId(v string) *GetGroupLiveListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupLiveListResponseBody) SetResult(v *GetGroupLiveListResponseBodyResult) *GetGroupLiveListResponseBody {
	s.Result = v
	return s
}

func (s *GetGroupLiveListResponseBody) SetVendorRequestId(v string) *GetGroupLiveListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetGroupLiveListResponseBody) SetVendorType(v string) *GetGroupLiveListResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetGroupLiveListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetGroupLiveListResponseBodyResult struct {
	// example:
	//
	// 直播列表
	GroupLiveList []*GetGroupLiveListResponseBodyResultGroupLiveList `json:"GroupLiveList,omitempty" xml:"GroupLiveList,omitempty" type:"Repeated"`
}

func (s GetGroupLiveListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListResponseBodyResult) GetGroupLiveList() []*GetGroupLiveListResponseBodyResultGroupLiveList {
	return s.GroupLiveList
}

func (s *GetGroupLiveListResponseBodyResult) SetGroupLiveList(v []*GetGroupLiveListResponseBodyResultGroupLiveList) *GetGroupLiveListResponseBodyResult {
	s.GroupLiveList = v
	return s
}

func (s *GetGroupLiveListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetGroupLiveListResponseBodyResultGroupLiveList struct {
	// example:
	//
	// nickName
	AnchorNickname *string `json:"AnchorNickname,omitempty" xml:"AnchorNickname,omitempty"`
	// example:
	//
	// Eijxx
	AnchorUnionId *string `json:"AnchorUnionId,omitempty" xml:"AnchorUnionId,omitempty"`
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

func (s GetGroupLiveListResponseBodyResultGroupLiveList) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListResponseBodyResultGroupLiveList) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) GetAnchorNickname() *string {
	return s.AnchorNickname
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) GetAnchorUnionId() *string {
	return s.AnchorUnionId
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) GetLiveEndTime() *int64 {
	return s.LiveEndTime
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) GetLiveStartTime() *int64 {
	return s.LiveStartTime
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) GetLiveUuid() *string {
	return s.LiveUuid
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) GetTitle() *string {
	return s.Title
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetAnchorNickname(v string) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.AnchorNickname = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetAnchorUnionId(v string) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.AnchorUnionId = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetLiveEndTime(v int64) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.LiveEndTime = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetLiveStartTime(v int64) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.LiveStartTime = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetLiveUuid(v string) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.LiveUuid = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetTitle(v string) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.Title = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) Validate() error {
	return dara.Validate(s)
}
