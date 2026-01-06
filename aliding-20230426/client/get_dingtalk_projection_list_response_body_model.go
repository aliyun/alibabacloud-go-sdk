// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetDingtalkProjectionListResponseBody
	GetCurrentPage() *int32
	SetData(v []*GetDingtalkProjectionListResponseBodyData) *GetDingtalkProjectionListResponseBody
	GetData() []*GetDingtalkProjectionListResponseBodyData
	SetRequestId(v string) *GetDingtalkProjectionListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetDingtalkProjectionListResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *GetDingtalkProjectionListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDingtalkProjectionListResponseBody
	GetVendorType() *string
}

type GetDingtalkProjectionListResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// []
	Data []*GetDingtalkProjectionListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetDingtalkProjectionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDingtalkProjectionListResponseBody) GetData() []*GetDingtalkProjectionListResponseBodyData {
	return s.Data
}

func (s *GetDingtalkProjectionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDingtalkProjectionListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetDingtalkProjectionListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDingtalkProjectionListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDingtalkProjectionListResponseBody) SetCurrentPage(v int32) *GetDingtalkProjectionListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBody) SetData(v []*GetDingtalkProjectionListResponseBodyData) *GetDingtalkProjectionListResponseBody {
	s.Data = v
	return s
}

func (s *GetDingtalkProjectionListResponseBody) SetRequestId(v string) *GetDingtalkProjectionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBody) SetTotalCount(v int64) *GetDingtalkProjectionListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBody) SetVendorRequestId(v string) *GetDingtalkProjectionListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBody) SetVendorType(v string) *GetDingtalkProjectionListResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBody) Validate() error {
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

type GetDingtalkProjectionListResponseBodyData struct {
	// example:
	//
	// IN
	CallStatus *string `json:"callStatus,omitempty" xml:"callStatus,omitempty"`
	// example:
	//
	// ABCD
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 3424242
	DevUid *int64 `json:"devUid,omitempty" xml:"devUid,omitempty"`
	// example:
	//
	// 517169
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1757729705000
	EndTs *int64 `json:"endTs,omitempty" xml:"endTs,omitempty"`
	// example:
	//
	// 342456
	NickCode *string `json:"nickCode,omitempty" xml:"nickCode,omitempty"`
	// example:
	//
	// 123456789
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// 4234242
	RecvClientId *string `json:"recvClientId,omitempty" xml:"recvClientId,omitempty"`
	// example:
	//
	// web
	RecvClientName *string `json:"recvClientName,omitempty" xml:"recvClientName,omitempty"`
	// example:
	//
	// room001
	SendClientId *string `json:"sendClientId,omitempty" xml:"sendClientId,omitempty"`
	// example:
	//
	// active
	SendClientName *string `json:"sendClientName,omitempty" xml:"sendClientName,omitempty"`
	// example:
	//
	// 24324
	SendClientWorkNo *string `json:"sendClientWorkNo,omitempty" xml:"sendClientWorkNo,omitempty"`
	// example:
	//
	// 1765502676356
	StartTs *int64 `json:"startTs,omitempty" xml:"startTs,omitempty"`
	// example:
	//
	// 2432424
	TimeStr *string `json:"timeStr,omitempty" xml:"timeStr,omitempty"`
}

func (s GetDingtalkProjectionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionListResponseBodyData) GetCallStatus() *string {
	return s.CallStatus
}

func (s *GetDingtalkProjectionListResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetDingtalkProjectionListResponseBodyData) GetDevUid() *int64 {
	return s.DevUid
}

func (s *GetDingtalkProjectionListResponseBodyData) GetDuration() *string {
	return s.Duration
}

func (s *GetDingtalkProjectionListResponseBodyData) GetEndTs() *int64 {
	return s.EndTs
}

func (s *GetDingtalkProjectionListResponseBodyData) GetNickCode() *string {
	return s.NickCode
}

func (s *GetDingtalkProjectionListResponseBodyData) GetOrgId() *int64 {
	return s.OrgId
}

func (s *GetDingtalkProjectionListResponseBodyData) GetRecvClientId() *string {
	return s.RecvClientId
}

func (s *GetDingtalkProjectionListResponseBodyData) GetRecvClientName() *string {
	return s.RecvClientName
}

func (s *GetDingtalkProjectionListResponseBodyData) GetSendClientId() *string {
	return s.SendClientId
}

func (s *GetDingtalkProjectionListResponseBodyData) GetSendClientName() *string {
	return s.SendClientName
}

func (s *GetDingtalkProjectionListResponseBodyData) GetSendClientWorkNo() *string {
	return s.SendClientWorkNo
}

func (s *GetDingtalkProjectionListResponseBodyData) GetStartTs() *int64 {
	return s.StartTs
}

func (s *GetDingtalkProjectionListResponseBodyData) GetTimeStr() *string {
	return s.TimeStr
}

func (s *GetDingtalkProjectionListResponseBodyData) SetCallStatus(v string) *GetDingtalkProjectionListResponseBodyData {
	s.CallStatus = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetCode(v string) *GetDingtalkProjectionListResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetDevUid(v int64) *GetDingtalkProjectionListResponseBodyData {
	s.DevUid = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetDuration(v string) *GetDingtalkProjectionListResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetEndTs(v int64) *GetDingtalkProjectionListResponseBodyData {
	s.EndTs = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetNickCode(v string) *GetDingtalkProjectionListResponseBodyData {
	s.NickCode = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetOrgId(v int64) *GetDingtalkProjectionListResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetRecvClientId(v string) *GetDingtalkProjectionListResponseBodyData {
	s.RecvClientId = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetRecvClientName(v string) *GetDingtalkProjectionListResponseBodyData {
	s.RecvClientName = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetSendClientId(v string) *GetDingtalkProjectionListResponseBodyData {
	s.SendClientId = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetSendClientName(v string) *GetDingtalkProjectionListResponseBodyData {
	s.SendClientName = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetSendClientWorkNo(v string) *GetDingtalkProjectionListResponseBodyData {
	s.SendClientWorkNo = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetStartTs(v int64) *GetDingtalkProjectionListResponseBodyData {
	s.StartTs = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) SetTimeStr(v string) *GetDingtalkProjectionListResponseBodyData {
	s.TimeStr = &v
	return s
}

func (s *GetDingtalkProjectionListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
