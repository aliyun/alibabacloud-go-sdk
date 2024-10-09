// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddSumRecordFlowPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 4546-100000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 会议名称
	ConferenceName *string `json:"ConferenceName,omitempty" xml:"ConferenceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Z10
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 入口名称
	EntryName *string `json:"EntryName,omitempty" xml:"EntryName,omitempty"`
	// example:
	//
	// 429005111100000
	Idcard *string `json:"Idcard,omitempty" xml:"Idcard,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-25 14:11
	SignTime *string `json:"SignTime,omitempty" xml:"SignTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddSumRecordFlowPopRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSumRecordFlowPopRequest) GoString() string {
	return s.String()
}

func (s *AddSumRecordFlowPopRequest) SetActivityId(v string) *AddSumRecordFlowPopRequest {
	s.ActivityId = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetCode(v string) *AddSumRecordFlowPopRequest {
	s.Code = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetConferenceName(v string) *AddSumRecordFlowPopRequest {
	s.ConferenceName = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetDeviceId(v string) *AddSumRecordFlowPopRequest {
	s.DeviceId = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetEntryName(v string) *AddSumRecordFlowPopRequest {
	s.EntryName = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetIdcard(v string) *AddSumRecordFlowPopRequest {
	s.Idcard = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetSignTime(v string) *AddSumRecordFlowPopRequest {
	s.SignTime = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetType(v int32) *AddSumRecordFlowPopRequest {
	s.Type = &v
	return s
}

type AddSumRecordFlowPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddSumRecordFlowPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSumRecordFlowPopResponseBody) GoString() string {
	return s.String()
}

func (s *AddSumRecordFlowPopResponseBody) SetAccessDeniedDetail(v string) *AddSumRecordFlowPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetData(v bool) *AddSumRecordFlowPopResponseBody {
	s.Data = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetErrCode(v string) *AddSumRecordFlowPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetErrMessage(v string) *AddSumRecordFlowPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetHttpStatusCode(v int32) *AddSumRecordFlowPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetRequestId(v string) *AddSumRecordFlowPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetSuccess(v bool) *AddSumRecordFlowPopResponseBody {
	s.Success = &v
	return s
}

type AddSumRecordFlowPopResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSumRecordFlowPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSumRecordFlowPopResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSumRecordFlowPopResponse) GoString() string {
	return s.String()
}

func (s *AddSumRecordFlowPopResponse) SetHeaders(v map[string]*string) *AddSumRecordFlowPopResponse {
	s.Headers = v
	return s
}

func (s *AddSumRecordFlowPopResponse) SetStatusCode(v int32) *AddSumRecordFlowPopResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSumRecordFlowPopResponse) SetBody(v *AddSumRecordFlowPopResponseBody) *AddSumRecordFlowPopResponse {
	s.Body = v
	return s
}

type BindExhibitorRfidPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Z10
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 451246
	GuestTicketRecordId *int64 `json:"GuestTicketRecordId,omitempty" xml:"GuestTicketRecordId,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asdojzopf
	Rfid *string `json:"Rfid,omitempty" xml:"Rfid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4546-100000
	TicketCode *string `json:"TicketCode,omitempty" xml:"TicketCode,omitempty"`
}

func (s BindExhibitorRfidPopRequest) String() string {
	return tea.Prettify(s)
}

func (s BindExhibitorRfidPopRequest) GoString() string {
	return s.String()
}

func (s *BindExhibitorRfidPopRequest) SetActivityId(v int64) *BindExhibitorRfidPopRequest {
	s.ActivityId = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetDeviceId(v string) *BindExhibitorRfidPopRequest {
	s.DeviceId = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetGmtCreate(v string) *BindExhibitorRfidPopRequest {
	s.GmtCreate = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetGmtModified(v string) *BindExhibitorRfidPopRequest {
	s.GmtModified = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetGuestTicketRecordId(v int64) *BindExhibitorRfidPopRequest {
	s.GuestTicketRecordId = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetId(v int64) *BindExhibitorRfidPopRequest {
	s.Id = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetRfid(v string) *BindExhibitorRfidPopRequest {
	s.Rfid = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetTicketCode(v string) *BindExhibitorRfidPopRequest {
	s.TicketCode = &v
	return s
}

type BindExhibitorRfidPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindExhibitorRfidPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindExhibitorRfidPopResponseBody) GoString() string {
	return s.String()
}

func (s *BindExhibitorRfidPopResponseBody) SetAccessDeniedDetail(v string) *BindExhibitorRfidPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetData(v bool) *BindExhibitorRfidPopResponseBody {
	s.Data = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetErrCode(v string) *BindExhibitorRfidPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetErrMessage(v string) *BindExhibitorRfidPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetHttpStatusCode(v int32) *BindExhibitorRfidPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetRequestId(v string) *BindExhibitorRfidPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetSuccess(v bool) *BindExhibitorRfidPopResponseBody {
	s.Success = &v
	return s
}

type BindExhibitorRfidPopResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindExhibitorRfidPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindExhibitorRfidPopResponse) String() string {
	return tea.Prettify(s)
}

func (s BindExhibitorRfidPopResponse) GoString() string {
	return s.String()
}

func (s *BindExhibitorRfidPopResponse) SetHeaders(v map[string]*string) *BindExhibitorRfidPopResponse {
	s.Headers = v
	return s
}

func (s *BindExhibitorRfidPopResponse) SetStatusCode(v int32) *BindExhibitorRfidPopResponse {
	s.StatusCode = &v
	return s
}

func (s *BindExhibitorRfidPopResponse) SetBody(v *BindExhibitorRfidPopResponseBody) *BindExhibitorRfidPopResponse {
	s.Body = v
	return s
}

type BindGuestRfidPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Z10
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 451246
	GuestTicketRecordId *int64 `json:"GuestTicketRecordId,omitempty" xml:"GuestTicketRecordId,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asdojzopf
	Rfid *string `json:"Rfid,omitempty" xml:"Rfid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4546-100000
	TicketCode *string `json:"TicketCode,omitempty" xml:"TicketCode,omitempty"`
}

func (s BindGuestRfidPopRequest) String() string {
	return tea.Prettify(s)
}

func (s BindGuestRfidPopRequest) GoString() string {
	return s.String()
}

func (s *BindGuestRfidPopRequest) SetActivityId(v int64) *BindGuestRfidPopRequest {
	s.ActivityId = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetDeviceId(v string) *BindGuestRfidPopRequest {
	s.DeviceId = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetGmtCreate(v string) *BindGuestRfidPopRequest {
	s.GmtCreate = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetGmtModified(v string) *BindGuestRfidPopRequest {
	s.GmtModified = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetGuestTicketRecordId(v int64) *BindGuestRfidPopRequest {
	s.GuestTicketRecordId = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetId(v int64) *BindGuestRfidPopRequest {
	s.Id = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetRfid(v string) *BindGuestRfidPopRequest {
	s.Rfid = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetTicketCode(v string) *BindGuestRfidPopRequest {
	s.TicketCode = &v
	return s
}

type BindGuestRfidPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindGuestRfidPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindGuestRfidPopResponseBody) GoString() string {
	return s.String()
}

func (s *BindGuestRfidPopResponseBody) SetAccessDeniedDetail(v string) *BindGuestRfidPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetData(v bool) *BindGuestRfidPopResponseBody {
	s.Data = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetErrCode(v string) *BindGuestRfidPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetErrMessage(v string) *BindGuestRfidPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetHttpStatusCode(v int32) *BindGuestRfidPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetRequestId(v string) *BindGuestRfidPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetSuccess(v bool) *BindGuestRfidPopResponseBody {
	s.Success = &v
	return s
}

type BindGuestRfidPopResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindGuestRfidPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindGuestRfidPopResponse) String() string {
	return tea.Prettify(s)
}

func (s BindGuestRfidPopResponse) GoString() string {
	return s.String()
}

func (s *BindGuestRfidPopResponse) SetHeaders(v map[string]*string) *BindGuestRfidPopResponse {
	s.Headers = v
	return s
}

func (s *BindGuestRfidPopResponse) SetStatusCode(v int32) *BindGuestRfidPopResponse {
	s.StatusCode = &v
	return s
}

func (s *BindGuestRfidPopResponse) SetBody(v *BindGuestRfidPopResponseBody) *BindGuestRfidPopResponse {
	s.Body = v
	return s
}

type CheckNFCBindPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asdojzopf
	NfcId *string `json:"NfcId,omitempty" xml:"NfcId,omitempty"`
}

func (s CheckNFCBindPopRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckNFCBindPopRequest) GoString() string {
	return s.String()
}

func (s *CheckNFCBindPopRequest) SetActivityId(v int64) *CheckNFCBindPopRequest {
	s.ActivityId = &v
	return s
}

func (s *CheckNFCBindPopRequest) SetNfcId(v string) *CheckNFCBindPopRequest {
	s.NfcId = &v
	return s
}

type CheckNFCBindPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *CheckNFCBindPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckNFCBindPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckNFCBindPopResponseBody) GoString() string {
	return s.String()
}

func (s *CheckNFCBindPopResponseBody) SetAccessDeniedDetail(v string) *CheckNFCBindPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetData(v *CheckNFCBindPopResponseBodyData) *CheckNFCBindPopResponseBody {
	s.Data = v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetErrCode(v string) *CheckNFCBindPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetErrMessage(v string) *CheckNFCBindPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetHttpStatusCode(v int32) *CheckNFCBindPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetRequestId(v string) *CheckNFCBindPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetSuccess(v bool) *CheckNFCBindPopResponseBody {
	s.Success = &v
	return s
}

type CheckNFCBindPopResponseBodyData struct {
	// example:
	//
	// 0
	IsGlobal *int32 `json:"IsGlobal,omitempty" xml:"IsGlobal,omitempty"`
	// example:
	//
	// true
	IsSign *bool `json:"IsSign,omitempty" xml:"IsSign,omitempty"`
}

func (s CheckNFCBindPopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckNFCBindPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckNFCBindPopResponseBodyData) SetIsGlobal(v int32) *CheckNFCBindPopResponseBodyData {
	s.IsGlobal = &v
	return s
}

func (s *CheckNFCBindPopResponseBodyData) SetIsSign(v bool) *CheckNFCBindPopResponseBodyData {
	s.IsSign = &v
	return s
}

type CheckNFCBindPopResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckNFCBindPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckNFCBindPopResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckNFCBindPopResponse) GoString() string {
	return s.String()
}

func (s *CheckNFCBindPopResponse) SetHeaders(v map[string]*string) *CheckNFCBindPopResponse {
	s.Headers = v
	return s
}

func (s *CheckNFCBindPopResponse) SetStatusCode(v int32) *CheckNFCBindPopResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckNFCBindPopResponse) SetBody(v *CheckNFCBindPopResponseBody) *CheckNFCBindPopResponse {
	s.Body = v
	return s
}

type FindGuestCredentialsRecordRequest struct {
	// example:
	//
	// 34429
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 2023-08-07 12:00:00
	DateTimeString *string `json:"DateTimeString,omitempty" xml:"DateTimeString,omitempty"`
	EndDateTime    *string `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	StartDateTime  *string `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
}

func (s FindGuestCredentialsRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s FindGuestCredentialsRecordRequest) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordRequest) SetActivityId(v string) *FindGuestCredentialsRecordRequest {
	s.ActivityId = &v
	return s
}

func (s *FindGuestCredentialsRecordRequest) SetDateTimeString(v string) *FindGuestCredentialsRecordRequest {
	s.DateTimeString = &v
	return s
}

func (s *FindGuestCredentialsRecordRequest) SetEndDateTime(v string) *FindGuestCredentialsRecordRequest {
	s.EndDateTime = &v
	return s
}

func (s *FindGuestCredentialsRecordRequest) SetStartDateTime(v string) *FindGuestCredentialsRecordRequest {
	s.StartDateTime = &v
	return s
}

type FindGuestCredentialsRecordResponseBody struct {
	Data []*FindGuestCredentialsRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 25294484-D133-5BDC-8952-243AD90CDF66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindGuestCredentialsRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBody) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBody) SetData(v []*FindGuestCredentialsRecordResponseBodyData) *FindGuestCredentialsRecordResponseBody {
	s.Data = v
	return s
}

func (s *FindGuestCredentialsRecordResponseBody) SetErrCode(v string) *FindGuestCredentialsRecordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBody) SetErrMessage(v string) *FindGuestCredentialsRecordResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBody) SetRequestId(v string) *FindGuestCredentialsRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBody) SetSuccess(v bool) *FindGuestCredentialsRecordResponseBody {
	s.Success = &v
	return s
}

type FindGuestCredentialsRecordResponseBodyData struct {
	// example:
	//
	// {}
	Admin map[string]interface{} `json:"Admin,omitempty" xml:"Admin,omitempty"`
	// example:
	//
	// 1401
	ChannelId        *int64                                                      `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelLevelInfo *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo `json:"ChannelLevelInfo,omitempty" xml:"ChannelLevelInfo,omitempty" type:"Struct"`
	CompanyName      *string                                                     `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// example:
	//
	// 3602-10010215
	CredentialsCode *string `json:"CredentialsCode,omitempty" xml:"CredentialsCode,omitempty"`
	CredentialsName *string `json:"CredentialsName,omitempty" xml:"CredentialsName,omitempty"`
	// example:
	//
	// IjEqZpp8Wn29+sqOr3hxXuOqn6CyKYNSQ5dmMA0txiM=
	IdNumber *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	IdType   *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FindGuestCredentialsRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetAdmin(v map[string]interface{}) *FindGuestCredentialsRecordResponseBodyData {
	s.Admin = v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetChannelId(v int64) *FindGuestCredentialsRecordResponseBodyData {
	s.ChannelId = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetChannelLevelInfo(v *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) *FindGuestCredentialsRecordResponseBodyData {
	s.ChannelLevelInfo = v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetCompanyName(v string) *FindGuestCredentialsRecordResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetCredentialsCode(v string) *FindGuestCredentialsRecordResponseBodyData {
	s.CredentialsCode = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetCredentialsName(v string) *FindGuestCredentialsRecordResponseBodyData {
	s.CredentialsName = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetIdNumber(v string) *FindGuestCredentialsRecordResponseBodyData {
	s.IdNumber = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetIdType(v string) *FindGuestCredentialsRecordResponseBodyData {
	s.IdType = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetName(v string) *FindGuestCredentialsRecordResponseBodyData {
	s.Name = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyData) SetStatus(v int32) *FindGuestCredentialsRecordResponseBodyData {
	s.Status = &v
	return s
}

type FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo struct {
	// example:
	//
	// 1401
	ChannelId             *int64                                                                       `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelName           *string                                                                      `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	LevelOneChannelName   *string                                                                      `json:"LevelOneChannelName,omitempty" xml:"LevelOneChannelName,omitempty"`
	LevelOneOwner         []*FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner   `json:"LevelOneOwner,omitempty" xml:"LevelOneOwner,omitempty" type:"Repeated"`
	LevelThreeChannelName *string                                                                      `json:"LevelThreeChannelName,omitempty" xml:"LevelThreeChannelName,omitempty"`
	LevelThreeOwner       []*FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner `json:"LevelThreeOwner,omitempty" xml:"LevelThreeOwner,omitempty" type:"Repeated"`
	LevelTwoChannelName   *string                                                                      `json:"LevelTwoChannelName,omitempty" xml:"LevelTwoChannelName,omitempty"`
	LevelTwoOwner         []*FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner   `json:"LevelTwoOwner,omitempty" xml:"LevelTwoOwner,omitempty" type:"Repeated"`
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) String() string {
	return tea.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) SetChannelId(v int64) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo {
	s.ChannelId = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) SetChannelName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo {
	s.ChannelName = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) SetLevelOneChannelName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo {
	s.LevelOneChannelName = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) SetLevelOneOwner(v []*FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo {
	s.LevelOneOwner = v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) SetLevelThreeChannelName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo {
	s.LevelThreeChannelName = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) SetLevelThreeOwner(v []*FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo {
	s.LevelThreeOwner = v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) SetLevelTwoChannelName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo {
	s.LevelTwoChannelName = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) SetLevelTwoOwner(v []*FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo {
	s.LevelTwoOwner = v
	return s
}

type FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner struct {
	// example:
	//
	// hBCycGELqJd6LEqSWKiLCQ==
	OwnerEmpIdOrTelephone *string `json:"OwnerEmpIdOrTelephone,omitempty" xml:"OwnerEmpIdOrTelephone,omitempty"`
	// example:
	//
	// buc_396545
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// hBCycGELqJd6LEqSWKiLCQ==
	OwnerNickName *string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty"`
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) String() string {
	return tea.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) SetOwnerEmpIdOrTelephone(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner {
	s.OwnerEmpIdOrTelephone = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) SetOwnerName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner {
	s.OwnerName = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) SetOwnerNickName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner {
	s.OwnerNickName = &v
	return s
}

type FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner struct {
	// example:
	//
	// hBCycGELqJd6LEqSWKiLCQ==
	OwnerEmpIdOrTelephone *string `json:"OwnerEmpIdOrTelephone,omitempty" xml:"OwnerEmpIdOrTelephone,omitempty"`
	// example:
	//
	// buc_160953
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// hBCycGELqJd6LEqSWKiLCQ==
	OwnerNickName *string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty"`
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) String() string {
	return tea.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) SetOwnerEmpIdOrTelephone(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner {
	s.OwnerEmpIdOrTelephone = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) SetOwnerName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner {
	s.OwnerName = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) SetOwnerNickName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner {
	s.OwnerNickName = &v
	return s
}

type FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner struct {
	// example:
	//
	// hBCycGELqJd6LEqSWKiLCQ==
	OwnerEmpIdOrTelephone *string `json:"OwnerEmpIdOrTelephone,omitempty" xml:"OwnerEmpIdOrTelephone,omitempty"`
	// example:
	//
	// buc_87239
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// hBCycGELqJd6LEqSWKiLCQ==
	OwnerNickName *string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty"`
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) String() string {
	return tea.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) SetOwnerEmpIdOrTelephone(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner {
	s.OwnerEmpIdOrTelephone = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) SetOwnerName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner {
	s.OwnerName = &v
	return s
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) SetOwnerNickName(v string) *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner {
	s.OwnerNickName = &v
	return s
}

type FindGuestCredentialsRecordResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindGuestCredentialsRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindGuestCredentialsRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s FindGuestCredentialsRecordResponse) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponse) SetHeaders(v map[string]*string) *FindGuestCredentialsRecordResponse {
	s.Headers = v
	return s
}

func (s *FindGuestCredentialsRecordResponse) SetStatusCode(v int32) *FindGuestCredentialsRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *FindGuestCredentialsRecordResponse) SetBody(v *FindGuestCredentialsRecordResponseBody) *FindGuestCredentialsRecordResponse {
	s.Body = v
	return s
}

type FindGuestTicketRecordRequest struct {
	// example:
	//
	// 34434
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 2023-09-04 15:14:00
	DateTimeString *string `json:"DateTimeString,omitempty" xml:"DateTimeString,omitempty"`
	EndDateTime    *string `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	StartDateTime  *string `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
}

func (s FindGuestTicketRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s FindGuestTicketRecordRequest) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordRequest) SetActivityId(v string) *FindGuestTicketRecordRequest {
	s.ActivityId = &v
	return s
}

func (s *FindGuestTicketRecordRequest) SetDateTimeString(v string) *FindGuestTicketRecordRequest {
	s.DateTimeString = &v
	return s
}

func (s *FindGuestTicketRecordRequest) SetEndDateTime(v string) *FindGuestTicketRecordRequest {
	s.EndDateTime = &v
	return s
}

func (s *FindGuestTicketRecordRequest) SetStartDateTime(v string) *FindGuestTicketRecordRequest {
	s.StartDateTime = &v
	return s
}

type FindGuestTicketRecordResponseBody struct {
	Data []*FindGuestTicketRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 8D190CE8-7D76-5781-8055-0990BBD2249F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindGuestTicketRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindGuestTicketRecordResponseBody) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBody) SetData(v []*FindGuestTicketRecordResponseBodyData) *FindGuestTicketRecordResponseBody {
	s.Data = v
	return s
}

func (s *FindGuestTicketRecordResponseBody) SetErrCode(v string) *FindGuestTicketRecordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *FindGuestTicketRecordResponseBody) SetErrMessage(v string) *FindGuestTicketRecordResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *FindGuestTicketRecordResponseBody) SetRequestId(v string) *FindGuestTicketRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindGuestTicketRecordResponseBody) SetSuccess(v bool) *FindGuestTicketRecordResponseBody {
	s.Success = &v
	return s
}

type FindGuestTicketRecordResponseBodyData struct {
	ChannelLevelInfo *FindGuestTicketRecordResponseBodyDataChannelLevelInfo `json:"ChannelLevelInfo,omitempty" xml:"ChannelLevelInfo,omitempty" type:"Struct"`
	CompanyName      *string                                                `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// example:
	//
	// 2023-10-31,2023-11-02
	EquityDates *string `json:"EquityDates,omitempty" xml:"EquityDates,omitempty"`
	// example:
	//
	// -1
	HealthCommitmentStatus *int32 `json:"HealthCommitmentStatus,omitempty" xml:"HealthCommitmentStatus,omitempty"`
	// example:
	//
	// IjEqZpp8Wn29+sqOr3hxXuOqn6CyKYNSQ5dmMA0txiM=
	IdNumber *string `json:"IdNumber,omitempty" xml:"IdNumber,omitempty"`
	IdType   *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 3702-10240842
	TicketCode *string `json:"TicketCode,omitempty" xml:"TicketCode,omitempty"`
	TicketName *string `json:"TicketName,omitempty" xml:"TicketName,omitempty"`
	// example:
	//
	// 1
	TicketType *string `json:"TicketType,omitempty" xml:"TicketType,omitempty"`
}

func (s FindGuestTicketRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyData) SetChannelLevelInfo(v *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) *FindGuestTicketRecordResponseBodyData {
	s.ChannelLevelInfo = v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetCompanyName(v string) *FindGuestTicketRecordResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetEquityDates(v string) *FindGuestTicketRecordResponseBodyData {
	s.EquityDates = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetHealthCommitmentStatus(v int32) *FindGuestTicketRecordResponseBodyData {
	s.HealthCommitmentStatus = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetIdNumber(v string) *FindGuestTicketRecordResponseBodyData {
	s.IdNumber = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetIdType(v string) *FindGuestTicketRecordResponseBodyData {
	s.IdType = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetName(v string) *FindGuestTicketRecordResponseBodyData {
	s.Name = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetStatus(v int32) *FindGuestTicketRecordResponseBodyData {
	s.Status = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetTicketCode(v string) *FindGuestTicketRecordResponseBodyData {
	s.TicketCode = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetTicketName(v string) *FindGuestTicketRecordResponseBodyData {
	s.TicketName = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetTicketType(v string) *FindGuestTicketRecordResponseBodyData {
	s.TicketType = &v
	return s
}

type FindGuestTicketRecordResponseBodyDataChannelLevelInfo struct {
	// example:
	//
	// 1401
	ChannelId *int64 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// VIP
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// VIP
	LevelOneChannelName *string                                                               `json:"LevelOneChannelName,omitempty" xml:"LevelOneChannelName,omitempty"`
	LevelOneOwner       []*FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner `json:"LevelOneOwner,omitempty" xml:"LevelOneOwner,omitempty" type:"Repeated"`
	// example:
	//
	// VIP
	LevelThreeChannelName *string                                                                 `json:"LevelThreeChannelName,omitempty" xml:"LevelThreeChannelName,omitempty"`
	LevelThreeOwner       []*FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner `json:"LevelThreeOwner,omitempty" xml:"LevelThreeOwner,omitempty" type:"Repeated"`
	// example:
	//
	// VIP
	LevelTwoChannelName *string                                                               `json:"LevelTwoChannelName,omitempty" xml:"LevelTwoChannelName,omitempty"`
	LevelTwoOwner       []*FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner `json:"LevelTwoOwner,omitempty" xml:"LevelTwoOwner,omitempty" type:"Repeated"`
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfo) String() string {
	return tea.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) SetChannelId(v int64) *FindGuestTicketRecordResponseBodyDataChannelLevelInfo {
	s.ChannelId = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) SetChannelName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfo {
	s.ChannelName = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) SetLevelOneChannelName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfo {
	s.LevelOneChannelName = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) SetLevelOneOwner(v []*FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) *FindGuestTicketRecordResponseBodyDataChannelLevelInfo {
	s.LevelOneOwner = v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) SetLevelThreeChannelName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfo {
	s.LevelThreeChannelName = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) SetLevelThreeOwner(v []*FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) *FindGuestTicketRecordResponseBodyDataChannelLevelInfo {
	s.LevelThreeOwner = v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) SetLevelTwoChannelName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfo {
	s.LevelTwoChannelName = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) SetLevelTwoOwner(v []*FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) *FindGuestTicketRecordResponseBodyDataChannelLevelInfo {
	s.LevelTwoOwner = v
	return s
}

type FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner struct {
	// example:
	//
	// IjEqZpp8Wn29+sqOr3hxXuOqn6CyKYNSQ5dmMA0txiM=
	OwnerEmpIdOrTelephone *string `json:"OwnerEmpIdOrTelephone,omitempty" xml:"OwnerEmpIdOrTelephone,omitempty"`
	OwnerName             *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// dUffKUpLXP5LFGeJa+Rs8Q==
	OwnerNickName *string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty"`
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) String() string {
	return tea.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) SetOwnerEmpIdOrTelephone(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner {
	s.OwnerEmpIdOrTelephone = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) SetOwnerName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner {
	s.OwnerName = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) SetOwnerNickName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner {
	s.OwnerNickName = &v
	return s
}

type FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner struct {
	// example:
	//
	// IjEqZpp8Wn29+sqOr3hxXuOqn6CyKYNSQ5dmMA0txiM=
	OwnerEmpIdOrTelephone *string `json:"OwnerEmpIdOrTelephone,omitempty" xml:"OwnerEmpIdOrTelephone,omitempty"`
	OwnerName             *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// dUffKUpLXP5LFGeJa+Rs8Q==
	OwnerNickName *string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty"`
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) String() string {
	return tea.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) SetOwnerEmpIdOrTelephone(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner {
	s.OwnerEmpIdOrTelephone = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) SetOwnerName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner {
	s.OwnerName = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) SetOwnerNickName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner {
	s.OwnerNickName = &v
	return s
}

type FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner struct {
	// example:
	//
	// IjEqZpp8Wn29+sqOr3hxXuOqn6CyKYNSQ5dmMA0txiM=
	OwnerEmpIdOrTelephone *string `json:"OwnerEmpIdOrTelephone,omitempty" xml:"OwnerEmpIdOrTelephone,omitempty"`
	OwnerName             *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// dUffKUpLXP5LFGeJa+Rs8Q==
	OwnerNickName *string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty"`
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) String() string {
	return tea.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) SetOwnerEmpIdOrTelephone(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner {
	s.OwnerEmpIdOrTelephone = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) SetOwnerName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner {
	s.OwnerName = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) SetOwnerNickName(v string) *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner {
	s.OwnerNickName = &v
	return s
}

type FindGuestTicketRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindGuestTicketRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindGuestTicketRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s FindGuestTicketRecordResponse) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponse) SetHeaders(v map[string]*string) *FindGuestTicketRecordResponse {
	s.Headers = v
	return s
}

func (s *FindGuestTicketRecordResponse) SetStatusCode(v int32) *FindGuestTicketRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *FindGuestTicketRecordResponse) SetBody(v *FindGuestTicketRecordResponseBody) *FindGuestTicketRecordResponse {
	s.Body = v
	return s
}

type QueryAllActivityInfoRequest struct {
	// This parameter is required.
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
}

func (s QueryAllActivityInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllActivityInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAllActivityInfoRequest) SetActivityId(v string) *QueryAllActivityInfoRequest {
	s.ActivityId = &v
	return s
}

type QueryAllActivityInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*QueryAllActivityInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAllActivityInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllActivityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllActivityInfoResponseBody) SetCode(v string) *QueryAllActivityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetData(v []*QueryAllActivityInfoResponseBodyData) *QueryAllActivityInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetMessage(v string) *QueryAllActivityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetRequestId(v string) *QueryAllActivityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAllActivityInfoResponseBody) SetSuccess(v bool) *QueryAllActivityInfoResponseBody {
	s.Success = &v
	return s
}

type QueryAllActivityInfoResponseBodyData struct {
	// example:
	//
	// 1234
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// vip
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// example:
	//
	// test
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// example:
	//
	// xx@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsVipCustomer *string `json:"IsVipCustomer,omitempty" xml:"IsVipCustomer,omitempty"`
	// example:
	//
	// 12123455
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	QRCode *string `json:"QRCode,omitempty" xml:"QRCode,omitempty"`
	// example:
	//
	// {}
	ReportFields *string `json:"ReportFields,omitempty" xml:"ReportFields,omitempty"`
}

func (s QueryAllActivityInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAllActivityInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAllActivityInfoResponseBodyData) SetActivityId(v int64) *QueryAllActivityInfoResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetChannelName(v string) *QueryAllActivityInfoResponseBodyData {
	s.ChannelName = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetCompanyName(v string) *QueryAllActivityInfoResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetCustomerName(v string) *QueryAllActivityInfoResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetEmail(v string) *QueryAllActivityInfoResponseBodyData {
	s.Email = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetId(v int64) *QueryAllActivityInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetIsVipCustomer(v string) *QueryAllActivityInfoResponseBodyData {
	s.IsVipCustomer = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetMobile(v string) *QueryAllActivityInfoResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetQRCode(v string) *QueryAllActivityInfoResponseBodyData {
	s.QRCode = &v
	return s
}

func (s *QueryAllActivityInfoResponseBodyData) SetReportFields(v string) *QueryAllActivityInfoResponseBodyData {
	s.ReportFields = &v
	return s
}

type QueryAllActivityInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAllActivityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAllActivityInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllActivityInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAllActivityInfoResponse) SetHeaders(v map[string]*string) *QueryAllActivityInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAllActivityInfoResponse) SetStatusCode(v int32) *QueryAllActivityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllActivityInfoResponse) SetBody(v *QueryAllActivityInfoResponseBody) *QueryAllActivityInfoResponse {
	s.Body = v
	return s
}

type QueryOrderSessionListPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asdojzopf
	NfcId *string `json:"NfcId,omitempty" xml:"NfcId,omitempty"`
}

func (s QueryOrderSessionListPopRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderSessionListPopRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderSessionListPopRequest) SetActivityId(v int64) *QueryOrderSessionListPopRequest {
	s.ActivityId = &v
	return s
}

func (s *QueryOrderSessionListPopRequest) SetNfcId(v string) *QueryOrderSessionListPopRequest {
	s.NfcId = &v
	return s
}

type QueryOrderSessionListPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data []*QueryOrderSessionListPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryOrderSessionListPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderSessionListPopResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderSessionListPopResponseBody) SetAccessDeniedDetail(v string) *QueryOrderSessionListPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetData(v []*QueryOrderSessionListPopResponseBodyData) *QueryOrderSessionListPopResponseBody {
	s.Data = v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetErrCode(v string) *QueryOrderSessionListPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetErrMessage(v string) *QueryOrderSessionListPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetHttpStatusCode(v int32) *QueryOrderSessionListPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetRequestId(v string) *QueryOrderSessionListPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetSuccess(v bool) *QueryOrderSessionListPopResponseBody {
	s.Success = &v
	return s
}

type QueryOrderSessionListPopResponseBodyData struct {
	// example:
	//
	// 1
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	SignInDate *string `json:"SignInDate,omitempty" xml:"SignInDate,omitempty"`
}

func (s QueryOrderSessionListPopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderSessionListPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOrderSessionListPopResponseBodyData) SetSessionId(v int64) *QueryOrderSessionListPopResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBodyData) SetSignInDate(v string) *QueryOrderSessionListPopResponseBodyData {
	s.SignInDate = &v
	return s
}

type QueryOrderSessionListPopResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrderSessionListPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrderSessionListPopResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderSessionListPopResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderSessionListPopResponse) SetHeaders(v map[string]*string) *QueryOrderSessionListPopResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderSessionListPopResponse) SetStatusCode(v int32) *QueryOrderSessionListPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderSessionListPopResponse) SetBody(v *QueryOrderSessionListPopResponseBody) *QueryOrderSessionListPopResponse {
	s.Body = v
	return s
}

type QuerySessionByActivityIdPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
}

func (s QuerySessionByActivityIdPopRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionByActivityIdPopRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionByActivityIdPopRequest) SetActivityId(v int64) *QuerySessionByActivityIdPopRequest {
	s.ActivityId = &v
	return s
}

type QuerySessionByActivityIdPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data []*QuerySessionByActivityIdPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySessionByActivityIdPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionByActivityIdPopResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionByActivityIdPopResponseBody) SetAccessDeniedDetail(v string) *QuerySessionByActivityIdPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetData(v []*QuerySessionByActivityIdPopResponseBodyData) *QuerySessionByActivityIdPopResponseBody {
	s.Data = v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetErrCode(v string) *QuerySessionByActivityIdPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetErrMessage(v string) *QuerySessionByActivityIdPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetHttpStatusCode(v int32) *QuerySessionByActivityIdPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetRequestId(v string) *QuerySessionByActivityIdPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBody) SetSuccess(v bool) *QuerySessionByActivityIdPopResponseBody {
	s.Success = &v
	return s
}

type QuerySessionByActivityIdPopResponseBodyData struct {
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// description
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	EndDateTime *string `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 地点
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 1234
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// name
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	StartDateTime *string `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
}

func (s QuerySessionByActivityIdPopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionByActivityIdPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetDescription(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.Description = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetDescriptionEn(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.DescriptionEn = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetEndDateTime(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.EndDateTime = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetId(v int64) *QuerySessionByActivityIdPopResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetLocation(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.Location = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetName(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.Name = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetNameEn(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.NameEn = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponseBodyData) SetStartDateTime(v string) *QuerySessionByActivityIdPopResponseBodyData {
	s.StartDateTime = &v
	return s
}

type QuerySessionByActivityIdPopResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionByActivityIdPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySessionByActivityIdPopResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionByActivityIdPopResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionByActivityIdPopResponse) SetHeaders(v map[string]*string) *QuerySessionByActivityIdPopResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionByActivityIdPopResponse) SetStatusCode(v int32) *QuerySessionByActivityIdPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponse) SetBody(v *QuerySessionByActivityIdPopResponseBody) *QuerySessionByActivityIdPopResponse {
	s.Body = v
	return s
}

type QuerySessionListPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asdojzopf
	NfcId *string `json:"NfcId,omitempty" xml:"NfcId,omitempty"`
}

func (s QuerySessionListPopRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionListPopRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionListPopRequest) SetActivityId(v int64) *QuerySessionListPopRequest {
	s.ActivityId = &v
	return s
}

func (s *QuerySessionListPopRequest) SetNfcId(v string) *QuerySessionListPopRequest {
	s.NfcId = &v
	return s
}

type QuerySessionListPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data []*QuerySessionListPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySessionListPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionListPopResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionListPopResponseBody) SetAccessDeniedDetail(v string) *QuerySessionListPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetData(v []*QuerySessionListPopResponseBodyData) *QuerySessionListPopResponseBody {
	s.Data = v
	return s
}

func (s *QuerySessionListPopResponseBody) SetErrCode(v string) *QuerySessionListPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetErrMessage(v string) *QuerySessionListPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetHttpStatusCode(v int32) *QuerySessionListPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetRequestId(v string) *QuerySessionListPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySessionListPopResponseBody) SetSuccess(v bool) *QuerySessionListPopResponseBody {
	s.Success = &v
	return s
}

type QuerySessionListPopResponseBodyData struct {
	// code
	//
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// id
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// location
	//
	// example:
	//
	// location
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// name
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QuerySessionListPopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionListPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySessionListPopResponseBodyData) SetCode(v string) *QuerySessionListPopResponseBodyData {
	s.Code = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetEndTime(v string) *QuerySessionListPopResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetId(v int64) *QuerySessionListPopResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetLocation(v string) *QuerySessionListPopResponseBodyData {
	s.Location = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetName(v string) *QuerySessionListPopResponseBodyData {
	s.Name = &v
	return s
}

func (s *QuerySessionListPopResponseBodyData) SetStartTime(v string) *QuerySessionListPopResponseBodyData {
	s.StartTime = &v
	return s
}

type QuerySessionListPopResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionListPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySessionListPopResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionListPopResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionListPopResponse) SetHeaders(v map[string]*string) *QuerySessionListPopResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionListPopResponse) SetStatusCode(v int32) *QuerySessionListPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionListPopResponse) SetBody(v *QuerySessionListPopResponseBody) *QuerySessionListPopResponse {
	s.Body = v
	return s
}

type QuerySignInRecordPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QuerySignInRecordPopRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySignInRecordPopRequest) GoString() string {
	return s.String()
}

func (s *QuerySignInRecordPopRequest) SetActivityId(v int64) *QuerySignInRecordPopRequest {
	s.ActivityId = &v
	return s
}

func (s *QuerySignInRecordPopRequest) SetEndTime(v string) *QuerySignInRecordPopRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySignInRecordPopRequest) SetPageNum(v int32) *QuerySignInRecordPopRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySignInRecordPopRequest) SetPageSize(v int32) *QuerySignInRecordPopRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySignInRecordPopRequest) SetStartTime(v string) *QuerySignInRecordPopRequest {
	s.StartTime = &v
	return s
}

type QuerySignInRecordPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data []*QuerySignInRecordPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySignInRecordPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySignInRecordPopResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySignInRecordPopResponseBody) SetAccessDeniedDetail(v string) *QuerySignInRecordPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetData(v []*QuerySignInRecordPopResponseBodyData) *QuerySignInRecordPopResponseBody {
	s.Data = v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetErrCode(v string) *QuerySignInRecordPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetErrMessage(v string) *QuerySignInRecordPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetHttpStatusCode(v int32) *QuerySignInRecordPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetRequestId(v string) *QuerySignInRecordPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetSuccess(v bool) *QuerySignInRecordPopResponseBody {
	s.Success = &v
	return s
}

type QuerySignInRecordPopResponseBodyData struct {
	// example:
	//
	// boarding
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// nfcid
	//
	// example:
	//
	// cshdsaodhoashd
	Rfid *string `json:"Rfid,omitempty" xml:"Rfid,omitempty"`
	// sessionId
	//
	// example:
	//
	// 2001
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QuerySignInRecordPopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySignInRecordPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySignInRecordPopResponseBodyData) SetEvent(v string) *QuerySignInRecordPopResponseBodyData {
	s.Event = &v
	return s
}

func (s *QuerySignInRecordPopResponseBodyData) SetRfid(v string) *QuerySignInRecordPopResponseBodyData {
	s.Rfid = &v
	return s
}

func (s *QuerySignInRecordPopResponseBodyData) SetSessionId(v int64) *QuerySignInRecordPopResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *QuerySignInRecordPopResponseBodyData) SetTime(v string) *QuerySignInRecordPopResponseBodyData {
	s.Time = &v
	return s
}

type QuerySignInRecordPopResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySignInRecordPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySignInRecordPopResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySignInRecordPopResponse) GoString() string {
	return s.String()
}

func (s *QuerySignInRecordPopResponse) SetHeaders(v map[string]*string) *QuerySignInRecordPopResponse {
	s.Headers = v
	return s
}

func (s *QuerySignInRecordPopResponse) SetStatusCode(v int32) *QuerySignInRecordPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySignInRecordPopResponse) SetBody(v *QuerySignInRecordPopResponseBody) *QuerySignInRecordPopResponse {
	s.Body = v
	return s
}

type QuerySingleActivityInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ActivityId   *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	CompanyName  *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// example:
	//
	// 12233445
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	QRCode *string `json:"QRCode,omitempty" xml:"QRCode,omitempty"`
}

func (s QuerySingleActivityInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleActivityInfoRequest) GoString() string {
	return s.String()
}

func (s *QuerySingleActivityInfoRequest) SetActivityId(v string) *QuerySingleActivityInfoRequest {
	s.ActivityId = &v
	return s
}

func (s *QuerySingleActivityInfoRequest) SetCompanyName(v string) *QuerySingleActivityInfoRequest {
	s.CompanyName = &v
	return s
}

func (s *QuerySingleActivityInfoRequest) SetCustomerName(v string) *QuerySingleActivityInfoRequest {
	s.CustomerName = &v
	return s
}

func (s *QuerySingleActivityInfoRequest) SetMobile(v string) *QuerySingleActivityInfoRequest {
	s.Mobile = &v
	return s
}

func (s *QuerySingleActivityInfoRequest) SetQRCode(v string) *QuerySingleActivityInfoRequest {
	s.QRCode = &v
	return s
}

type QuerySingleActivityInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*QuerySingleActivityInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySingleActivityInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleActivityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySingleActivityInfoResponseBody) SetCode(v string) *QuerySingleActivityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetData(v []*QuerySingleActivityInfoResponseBodyData) *QuerySingleActivityInfoResponseBody {
	s.Data = v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetMessage(v string) *QuerySingleActivityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetRequestId(v string) *QuerySingleActivityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBody) SetSuccess(v bool) *QuerySingleActivityInfoResponseBody {
	s.Success = &v
	return s
}

type QuerySingleActivityInfoResponseBodyData struct {
	// example:
	//
	// 123
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// vip
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// aliyun
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// example:
	//
	// xxx
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// example:
	//
	// xx@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsVipCustomer *string `json:"IsVipCustomer,omitempty" xml:"IsVipCustomer,omitempty"`
	// example:
	//
	// 234355**
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// http://xxx.com?a=xx
	QRCode *string `json:"QRCode,omitempty" xml:"QRCode,omitempty"`
	// example:
	//
	// {}
	ReportFields *string `json:"ReportFields,omitempty" xml:"ReportFields,omitempty"`
}

func (s QuerySingleActivityInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleActivityInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySingleActivityInfoResponseBodyData) SetActivityId(v int64) *QuerySingleActivityInfoResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetChannelName(v string) *QuerySingleActivityInfoResponseBodyData {
	s.ChannelName = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetCompanyName(v string) *QuerySingleActivityInfoResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetCustomerName(v string) *QuerySingleActivityInfoResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetEmail(v string) *QuerySingleActivityInfoResponseBodyData {
	s.Email = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetId(v int64) *QuerySingleActivityInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetIsVipCustomer(v string) *QuerySingleActivityInfoResponseBodyData {
	s.IsVipCustomer = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetMobile(v string) *QuerySingleActivityInfoResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetQRCode(v string) *QuerySingleActivityInfoResponseBodyData {
	s.QRCode = &v
	return s
}

func (s *QuerySingleActivityInfoResponseBodyData) SetReportFields(v string) *QuerySingleActivityInfoResponseBodyData {
	s.ReportFields = &v
	return s
}

type QuerySingleActivityInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySingleActivityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySingleActivityInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleActivityInfoResponse) GoString() string {
	return s.String()
}

func (s *QuerySingleActivityInfoResponse) SetHeaders(v map[string]*string) *QuerySingleActivityInfoResponse {
	s.Headers = v
	return s
}

func (s *QuerySingleActivityInfoResponse) SetStatusCode(v int32) *QuerySingleActivityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySingleActivityInfoResponse) SetBody(v *QuerySingleActivityInfoResponseBody) *QuerySingleActivityInfoResponse {
	s.Body = v
	return s
}

type SyncSignInInfoRequest struct {
	// This parameter is required.
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	QRCode *string `json:"QRCode,omitempty" xml:"QRCode,omitempty"`
}

func (s SyncSignInInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncSignInInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncSignInInfoRequest) SetActivityId(v string) *SyncSignInInfoRequest {
	s.ActivityId = &v
	return s
}

func (s *SyncSignInInfoRequest) SetQRCode(v string) *SyncSignInInfoRequest {
	s.QRCode = &v
	return s
}

type SyncSignInInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncSignInInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncSignInInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncSignInInfoResponseBody) SetCode(v string) *SyncSignInInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetData(v int32) *SyncSignInInfoResponseBody {
	s.Data = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetMessage(v string) *SyncSignInInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetRequestId(v string) *SyncSignInInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetSuccess(v bool) *SyncSignInInfoResponseBody {
	s.Success = &v
	return s
}

type SyncSignInInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncSignInInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncSignInInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncSignInInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncSignInInfoResponse) SetHeaders(v map[string]*string) *SyncSignInInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncSignInInfoResponse) SetStatusCode(v int32) *SyncSignInInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncSignInInfoResponse) SetBody(v *SyncSignInInfoResponseBody) *SyncSignInInfoResponse {
	s.Body = v
	return s
}

type TicketOrCredentialsSignInPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 4546-100000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 会议名称
	ConferenceName *string `json:"ConferenceName,omitempty" xml:"ConferenceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Z10
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 入口名称
	EntryName *string `json:"EntryName,omitempty" xml:"EntryName,omitempty"`
	// example:
	//
	// 429005111100000
	Idcard *string `json:"Idcard,omitempty" xml:"Idcard,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-25 14:11
	SignTime *string `json:"SignTime,omitempty" xml:"SignTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TicketOrCredentialsSignInPopRequest) String() string {
	return tea.Prettify(s)
}

func (s TicketOrCredentialsSignInPopRequest) GoString() string {
	return s.String()
}

func (s *TicketOrCredentialsSignInPopRequest) SetActivityId(v string) *TicketOrCredentialsSignInPopRequest {
	s.ActivityId = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetCode(v string) *TicketOrCredentialsSignInPopRequest {
	s.Code = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetConferenceName(v string) *TicketOrCredentialsSignInPopRequest {
	s.ConferenceName = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetDeviceId(v string) *TicketOrCredentialsSignInPopRequest {
	s.DeviceId = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetEntryName(v string) *TicketOrCredentialsSignInPopRequest {
	s.EntryName = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetIdcard(v string) *TicketOrCredentialsSignInPopRequest {
	s.Idcard = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetSignTime(v string) *TicketOrCredentialsSignInPopRequest {
	s.SignTime = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetType(v int32) *TicketOrCredentialsSignInPopRequest {
	s.Type = &v
	return s
}

type TicketOrCredentialsSignInPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TicketOrCredentialsSignInPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TicketOrCredentialsSignInPopResponseBody) GoString() string {
	return s.String()
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetAccessDeniedDetail(v string) *TicketOrCredentialsSignInPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetData(v interface{}) *TicketOrCredentialsSignInPopResponseBody {
	s.Data = v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetErrCode(v string) *TicketOrCredentialsSignInPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetErrMessage(v string) *TicketOrCredentialsSignInPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetHttpStatusCode(v int32) *TicketOrCredentialsSignInPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetRequestId(v string) *TicketOrCredentialsSignInPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetSuccess(v bool) *TicketOrCredentialsSignInPopResponseBody {
	s.Success = &v
	return s
}

type TicketOrCredentialsSignInPopResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketOrCredentialsSignInPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketOrCredentialsSignInPopResponse) String() string {
	return tea.Prettify(s)
}

func (s TicketOrCredentialsSignInPopResponse) GoString() string {
	return s.String()
}

func (s *TicketOrCredentialsSignInPopResponse) SetHeaders(v map[string]*string) *TicketOrCredentialsSignInPopResponse {
	s.Headers = v
	return s
}

func (s *TicketOrCredentialsSignInPopResponse) SetStatusCode(v int32) *TicketOrCredentialsSignInPopResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponse) SetBody(v *TicketOrCredentialsSignInPopResponseBody) *TicketOrCredentialsSignInPopResponse {
	s.Body = v
	return s
}

type UpdateCredentialsStatusPopRequest struct {
	// example:
	//
	// 4546-100000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 张三
	ProxyRecipientName *string `json:"ProxyRecipientName,omitempty" xml:"ProxyRecipientName,omitempty"`
	// example:
	//
	// 14787627188
	ProxyRecipientPhoneNumber *string `json:"ProxyRecipientPhoneNumber,omitempty" xml:"ProxyRecipientPhoneNumber,omitempty"`
	// example:
	//
	// Z30
	ReceiptLocation *string `json:"ReceiptLocation,omitempty" xml:"ReceiptLocation,omitempty"`
	// example:
	//
	// 429005111100000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s UpdateCredentialsStatusPopRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCredentialsStatusPopRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsStatusPopRequest) SetCode(v string) *UpdateCredentialsStatusPopRequest {
	s.Code = &v
	return s
}

func (s *UpdateCredentialsStatusPopRequest) SetProxyRecipientName(v string) *UpdateCredentialsStatusPopRequest {
	s.ProxyRecipientName = &v
	return s
}

func (s *UpdateCredentialsStatusPopRequest) SetProxyRecipientPhoneNumber(v string) *UpdateCredentialsStatusPopRequest {
	s.ProxyRecipientPhoneNumber = &v
	return s
}

func (s *UpdateCredentialsStatusPopRequest) SetReceiptLocation(v string) *UpdateCredentialsStatusPopRequest {
	s.ReceiptLocation = &v
	return s
}

func (s *UpdateCredentialsStatusPopRequest) SetTime(v string) *UpdateCredentialsStatusPopRequest {
	s.Time = &v
	return s
}

type UpdateCredentialsStatusPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCredentialsStatusPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCredentialsStatusPopResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsStatusPopResponseBody) SetAccessDeniedDetail(v string) *UpdateCredentialsStatusPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetData(v bool) *UpdateCredentialsStatusPopResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetErrCode(v string) *UpdateCredentialsStatusPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetErrMessage(v string) *UpdateCredentialsStatusPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetHttpStatusCode(v int32) *UpdateCredentialsStatusPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetRequestId(v string) *UpdateCredentialsStatusPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetSuccess(v bool) *UpdateCredentialsStatusPopResponseBody {
	s.Success = &v
	return s
}

type UpdateCredentialsStatusPopResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCredentialsStatusPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCredentialsStatusPopResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCredentialsStatusPopResponse) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsStatusPopResponse) SetHeaders(v map[string]*string) *UpdateCredentialsStatusPopResponse {
	s.Headers = v
	return s
}

func (s *UpdateCredentialsStatusPopResponse) SetStatusCode(v int32) *UpdateCredentialsStatusPopResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponse) SetBody(v *UpdateCredentialsStatusPopResponseBody) *UpdateCredentialsStatusPopResponse {
	s.Body = v
	return s
}

type UpdateTicketRecordByticketCodePopRequest struct {
	// example:
	//
	// 3402
	AgendaId *string `json:"AgendaId,omitempty" xml:"AgendaId,omitempty"`
	// example:
	//
	// 4546-100000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// boarding
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 12345
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s UpdateTicketRecordByticketCodePopRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRecordByticketCodePopRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetAgendaId(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.AgendaId = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetCode(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.Code = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetEvent(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.Event = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetSceneId(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetTime(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.Time = &v
	return s
}

type UpdateTicketRecordByticketCodePopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTicketRecordByticketCodePopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRecordByticketCodePopResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetAccessDeniedDetail(v string) *UpdateTicketRecordByticketCodePopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetData(v bool) *UpdateTicketRecordByticketCodePopResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetErrCode(v string) *UpdateTicketRecordByticketCodePopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetErrMessage(v string) *UpdateTicketRecordByticketCodePopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetHttpStatusCode(v int32) *UpdateTicketRecordByticketCodePopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetRequestId(v string) *UpdateTicketRecordByticketCodePopResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetSuccess(v bool) *UpdateTicketRecordByticketCodePopResponseBody {
	s.Success = &v
	return s
}

type UpdateTicketRecordByticketCodePopResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTicketRecordByticketCodePopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTicketRecordByticketCodePopResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRecordByticketCodePopResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketRecordByticketCodePopResponse) SetHeaders(v map[string]*string) *UpdateTicketRecordByticketCodePopResponse {
	s.Headers = v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponse) SetStatusCode(v int32) *UpdateTicketRecordByticketCodePopResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponse) SetBody(v *UpdateTicketRecordByticketCodePopResponseBody) *UpdateTicketRecordByticketCodePopResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("marketing_event"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddSumRecordFlowPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSumRecordFlowPopResponse
func (client *Client) AddSumRecordFlowPopWithOptions(request *AddSumRecordFlowPopRequest, runtime *util.RuntimeOptions) (_result *AddSumRecordFlowPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ConferenceName)) {
		query["ConferenceName"] = request.ConferenceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.EntryName)) {
		query["EntryName"] = request.EntryName
	}

	if !tea.BoolValue(util.IsUnset(request.Idcard)) {
		query["Idcard"] = request.Idcard
	}

	if !tea.BoolValue(util.IsUnset(request.SignTime)) {
		query["SignTime"] = request.SignTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSumRecordFlowPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSumRecordFlowPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddSumRecordFlowPopRequest
//
// @return AddSumRecordFlowPopResponse
func (client *Client) AddSumRecordFlowPop(request *AddSumRecordFlowPopRequest) (_result *AddSumRecordFlowPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSumRecordFlowPopResponse{}
	_body, _err := client.AddSumRecordFlowPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BindExhibitorRfidPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindExhibitorRfidPopResponse
func (client *Client) BindExhibitorRfidPopWithOptions(request *BindExhibitorRfidPopRequest, runtime *util.RuntimeOptions) (_result *BindExhibitorRfidPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.GmtCreate)) {
		query["GmtCreate"] = request.GmtCreate
	}

	if !tea.BoolValue(util.IsUnset(request.GmtModified)) {
		query["GmtModified"] = request.GmtModified
	}

	if !tea.BoolValue(util.IsUnset(request.GuestTicketRecordId)) {
		query["GuestTicketRecordId"] = request.GuestTicketRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Rfid)) {
		query["Rfid"] = request.Rfid
	}

	if !tea.BoolValue(util.IsUnset(request.TicketCode)) {
		query["TicketCode"] = request.TicketCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindExhibitorRfidPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindExhibitorRfidPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindExhibitorRfidPopRequest
//
// @return BindExhibitorRfidPopResponse
func (client *Client) BindExhibitorRfidPop(request *BindExhibitorRfidPopRequest) (_result *BindExhibitorRfidPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindExhibitorRfidPopResponse{}
	_body, _err := client.BindExhibitorRfidPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BindGuestRfidPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindGuestRfidPopResponse
func (client *Client) BindGuestRfidPopWithOptions(request *BindGuestRfidPopRequest, runtime *util.RuntimeOptions) (_result *BindGuestRfidPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.GmtCreate)) {
		query["GmtCreate"] = request.GmtCreate
	}

	if !tea.BoolValue(util.IsUnset(request.GmtModified)) {
		query["GmtModified"] = request.GmtModified
	}

	if !tea.BoolValue(util.IsUnset(request.GuestTicketRecordId)) {
		query["GuestTicketRecordId"] = request.GuestTicketRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Rfid)) {
		query["Rfid"] = request.Rfid
	}

	if !tea.BoolValue(util.IsUnset(request.TicketCode)) {
		query["TicketCode"] = request.TicketCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindGuestRfidPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindGuestRfidPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindGuestRfidPopRequest
//
// @return BindGuestRfidPopResponse
func (client *Client) BindGuestRfidPop(request *BindGuestRfidPopRequest) (_result *BindGuestRfidPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindGuestRfidPopResponse{}
	_body, _err := client.BindGuestRfidPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CheckNFCBindPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckNFCBindPopResponse
func (client *Client) CheckNFCBindPopWithOptions(request *CheckNFCBindPopRequest, runtime *util.RuntimeOptions) (_result *CheckNFCBindPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.NfcId)) {
		query["NfcId"] = request.NfcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckNFCBindPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckNFCBindPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckNFCBindPopRequest
//
// @return CheckNFCBindPopResponse
func (client *Client) CheckNFCBindPop(request *CheckNFCBindPopRequest) (_result *CheckNFCBindPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckNFCBindPopResponse{}
	_body, _err := client.CheckNFCBindPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拉取领证人员记录
//
// @param request - FindGuestCredentialsRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindGuestCredentialsRecordResponse
func (client *Client) FindGuestCredentialsRecordWithOptions(request *FindGuestCredentialsRecordRequest, runtime *util.RuntimeOptions) (_result *FindGuestCredentialsRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.DateTimeString)) {
		query["DateTimeString"] = request.DateTimeString
	}

	if !tea.BoolValue(util.IsUnset(request.EndDateTime)) {
		query["EndDateTime"] = request.EndDateTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateTime)) {
		query["StartDateTime"] = request.StartDateTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindGuestCredentialsRecord"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindGuestCredentialsRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拉取领证人员记录
//
// @param request - FindGuestCredentialsRecordRequest
//
// @return FindGuestCredentialsRecordResponse
func (client *Client) FindGuestCredentialsRecord(request *FindGuestCredentialsRecordRequest) (_result *FindGuestCredentialsRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindGuestCredentialsRecordResponse{}
	_body, _err := client.FindGuestCredentialsRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 云栖大会获取取票人信息list接口
//
// @param request - FindGuestTicketRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindGuestTicketRecordResponse
func (client *Client) FindGuestTicketRecordWithOptions(request *FindGuestTicketRecordRequest, runtime *util.RuntimeOptions) (_result *FindGuestTicketRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.DateTimeString)) {
		query["DateTimeString"] = request.DateTimeString
	}

	if !tea.BoolValue(util.IsUnset(request.EndDateTime)) {
		query["EndDateTime"] = request.EndDateTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateTime)) {
		query["StartDateTime"] = request.StartDateTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FindGuestTicketRecord"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindGuestTicketRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 云栖大会获取取票人信息list接口
//
// @param request - FindGuestTicketRecordRequest
//
// @return FindGuestTicketRecordResponse
func (client *Client) FindGuestTicketRecord(request *FindGuestTicketRecordRequest) (_result *FindGuestTicketRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindGuestTicketRecordResponse{}
	_body, _err := client.FindGuestTicketRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAllActivityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllActivityInfoResponse
func (client *Client) QueryAllActivityInfoWithOptions(request *QueryAllActivityInfoRequest, runtime *util.RuntimeOptions) (_result *QueryAllActivityInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAllActivityInfo"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllActivityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAllActivityInfoRequest
//
// @return QueryAllActivityInfoResponse
func (client *Client) QueryAllActivityInfo(request *QueryAllActivityInfoRequest) (_result *QueryAllActivityInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAllActivityInfoResponse{}
	_body, _err := client.QueryAllActivityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryOrderSessionListPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrderSessionListPopResponse
func (client *Client) QueryOrderSessionListPopWithOptions(request *QueryOrderSessionListPopRequest, runtime *util.RuntimeOptions) (_result *QueryOrderSessionListPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.NfcId)) {
		query["NfcId"] = request.NfcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrderSessionListPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrderSessionListPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryOrderSessionListPopRequest
//
// @return QueryOrderSessionListPopResponse
func (client *Client) QueryOrderSessionListPop(request *QueryOrderSessionListPopRequest) (_result *QueryOrderSessionListPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderSessionListPopResponse{}
	_body, _err := client.QueryOrderSessionListPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySessionByActivityIdPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionByActivityIdPopResponse
func (client *Client) QuerySessionByActivityIdPopWithOptions(request *QuerySessionByActivityIdPopRequest, runtime *util.RuntimeOptions) (_result *QuerySessionByActivityIdPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySessionByActivityIdPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySessionByActivityIdPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySessionByActivityIdPopRequest
//
// @return QuerySessionByActivityIdPopResponse
func (client *Client) QuerySessionByActivityIdPop(request *QuerySessionByActivityIdPopRequest) (_result *QuerySessionByActivityIdPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySessionByActivityIdPopResponse{}
	_body, _err := client.QuerySessionByActivityIdPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySessionListPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionListPopResponse
func (client *Client) QuerySessionListPopWithOptions(request *QuerySessionListPopRequest, runtime *util.RuntimeOptions) (_result *QuerySessionListPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.NfcId)) {
		query["NfcId"] = request.NfcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySessionListPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySessionListPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySessionListPopRequest
//
// @return QuerySessionListPopResponse
func (client *Client) QuerySessionListPop(request *QuerySessionListPopRequest) (_result *QuerySessionListPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySessionListPopResponse{}
	_body, _err := client.QuerySessionListPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySignInRecordPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySignInRecordPopResponse
func (client *Client) QuerySignInRecordPopWithOptions(request *QuerySignInRecordPopRequest, runtime *util.RuntimeOptions) (_result *QuerySignInRecordPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySignInRecordPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySignInRecordPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySignInRecordPopRequest
//
// @return QuerySignInRecordPopResponse
func (client *Client) QuerySignInRecordPop(request *QuerySignInRecordPopRequest) (_result *QuerySignInRecordPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySignInRecordPopResponse{}
	_body, _err := client.QuerySignInRecordPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySingleActivityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySingleActivityInfoResponse
func (client *Client) QuerySingleActivityInfoWithOptions(request *QuerySingleActivityInfoRequest, runtime *util.RuntimeOptions) (_result *QuerySingleActivityInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySingleActivityInfo"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySingleActivityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySingleActivityInfoRequest
//
// @return QuerySingleActivityInfoResponse
func (client *Client) QuerySingleActivityInfo(request *QuerySingleActivityInfoRequest) (_result *QuerySingleActivityInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySingleActivityInfoResponse{}
	_body, _err := client.QuerySingleActivityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SyncSignInInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncSignInInfoResponse
func (client *Client) SyncSignInInfoWithOptions(request *SyncSignInInfoRequest, runtime *util.RuntimeOptions) (_result *SyncSignInInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncSignInInfo"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncSignInInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SyncSignInInfoRequest
//
// @return SyncSignInInfoResponse
func (client *Client) SyncSignInInfo(request *SyncSignInInfoRequest) (_result *SyncSignInInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncSignInInfoResponse{}
	_body, _err := client.SyncSignInInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - TicketOrCredentialsSignInPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketOrCredentialsSignInPopResponse
func (client *Client) TicketOrCredentialsSignInPopWithOptions(request *TicketOrCredentialsSignInPopRequest, runtime *util.RuntimeOptions) (_result *TicketOrCredentialsSignInPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ConferenceName)) {
		query["ConferenceName"] = request.ConferenceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.EntryName)) {
		query["EntryName"] = request.EntryName
	}

	if !tea.BoolValue(util.IsUnset(request.Idcard)) {
		query["Idcard"] = request.Idcard
	}

	if !tea.BoolValue(util.IsUnset(request.SignTime)) {
		query["SignTime"] = request.SignTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TicketOrCredentialsSignInPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TicketOrCredentialsSignInPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - TicketOrCredentialsSignInPopRequest
//
// @return TicketOrCredentialsSignInPopResponse
func (client *Client) TicketOrCredentialsSignInPop(request *TicketOrCredentialsSignInPopRequest) (_result *TicketOrCredentialsSignInPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TicketOrCredentialsSignInPopResponse{}
	_body, _err := client.TicketOrCredentialsSignInPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateCredentialsStatusPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialsStatusPopResponse
func (client *Client) UpdateCredentialsStatusPopWithOptions(request *UpdateCredentialsStatusPopRequest, runtime *util.RuntimeOptions) (_result *UpdateCredentialsStatusPopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyRecipientName)) {
		query["ProxyRecipientName"] = request.ProxyRecipientName
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyRecipientPhoneNumber)) {
		query["ProxyRecipientPhoneNumber"] = request.ProxyRecipientPhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptLocation)) {
		query["ReceiptLocation"] = request.ReceiptLocation
	}

	if !tea.BoolValue(util.IsUnset(request.Time)) {
		query["Time"] = request.Time
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCredentialsStatusPop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCredentialsStatusPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCredentialsStatusPopRequest
//
// @return UpdateCredentialsStatusPopResponse
func (client *Client) UpdateCredentialsStatusPop(request *UpdateCredentialsStatusPopRequest) (_result *UpdateCredentialsStatusPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCredentialsStatusPopResponse{}
	_body, _err := client.UpdateCredentialsStatusPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateTicketRecordByticketCodePopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketRecordByticketCodePopResponse
func (client *Client) UpdateTicketRecordByticketCodePopWithOptions(request *UpdateTicketRecordByticketCodePopRequest, runtime *util.RuntimeOptions) (_result *UpdateTicketRecordByticketCodePopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgendaId)) {
		query["AgendaId"] = request.AgendaId
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Event)) {
		query["Event"] = request.Event
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.Time)) {
		query["Time"] = request.Time
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTicketRecordByticketCodePop"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTicketRecordByticketCodePopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateTicketRecordByticketCodePopRequest
//
// @return UpdateTicketRecordByticketCodePopResponse
func (client *Client) UpdateTicketRecordByticketCodePop(request *UpdateTicketRecordByticketCodePopRequest) (_result *UpdateTicketRecordByticketCodePopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTicketRecordByticketCodePopResponse{}
	_body, _err := client.UpdateTicketRecordByticketCodePopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
