// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type FindGuestCredentialsRecordRequest struct {
	// example:
	//
	// 34429
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 2023-08-07 12:00:00
	DateTimeString *string `json:"DateTimeString,omitempty" xml:"DateTimeString,omitempty"`
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
