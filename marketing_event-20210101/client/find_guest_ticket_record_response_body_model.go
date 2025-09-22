// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindGuestTicketRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*FindGuestTicketRecordResponseBodyData) *FindGuestTicketRecordResponseBody
	GetData() []*FindGuestTicketRecordResponseBodyData
	SetErrCode(v string) *FindGuestTicketRecordResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *FindGuestTicketRecordResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *FindGuestTicketRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindGuestTicketRecordResponseBody
	GetSuccess() *bool
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
	return dara.Prettify(s)
}

func (s FindGuestTicketRecordResponseBody) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBody) GetData() []*FindGuestTicketRecordResponseBodyData {
	return s.Data
}

func (s *FindGuestTicketRecordResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *FindGuestTicketRecordResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *FindGuestTicketRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindGuestTicketRecordResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *FindGuestTicketRecordResponseBody) Validate() error {
	return dara.Validate(s)
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
	TicketCode         *string `json:"TicketCode,omitempty" xml:"TicketCode,omitempty"`
	TicketName         *string `json:"TicketName,omitempty" xml:"TicketName,omitempty"`
	TicketReceiveDates *string `json:"TicketReceiveDates,omitempty" xml:"TicketReceiveDates,omitempty"`
	// example:
	//
	// 1
	TicketType *string `json:"TicketType,omitempty" xml:"TicketType,omitempty"`
}

func (s FindGuestTicketRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyData) GetChannelLevelInfo() *FindGuestTicketRecordResponseBodyDataChannelLevelInfo {
	return s.ChannelLevelInfo
}

func (s *FindGuestTicketRecordResponseBodyData) GetCompanyName() *string {
	return s.CompanyName
}

func (s *FindGuestTicketRecordResponseBodyData) GetEquityDates() *string {
	return s.EquityDates
}

func (s *FindGuestTicketRecordResponseBodyData) GetHealthCommitmentStatus() *int32 {
	return s.HealthCommitmentStatus
}

func (s *FindGuestTicketRecordResponseBodyData) GetIdNumber() *string {
	return s.IdNumber
}

func (s *FindGuestTicketRecordResponseBodyData) GetIdType() *string {
	return s.IdType
}

func (s *FindGuestTicketRecordResponseBodyData) GetName() *string {
	return s.Name
}

func (s *FindGuestTicketRecordResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *FindGuestTicketRecordResponseBodyData) GetTicketCode() *string {
	return s.TicketCode
}

func (s *FindGuestTicketRecordResponseBodyData) GetTicketName() *string {
	return s.TicketName
}

func (s *FindGuestTicketRecordResponseBodyData) GetTicketReceiveDates() *string {
	return s.TicketReceiveDates
}

func (s *FindGuestTicketRecordResponseBodyData) GetTicketType() *string {
	return s.TicketType
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

func (s *FindGuestTicketRecordResponseBodyData) SetTicketReceiveDates(v string) *FindGuestTicketRecordResponseBodyData {
	s.TicketReceiveDates = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) SetTicketType(v string) *FindGuestTicketRecordResponseBodyData {
	s.TicketType = &v
	return s
}

func (s *FindGuestTicketRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GetChannelId() *int64 {
	return s.ChannelId
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GetChannelName() *string {
	return s.ChannelName
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GetLevelOneChannelName() *string {
	return s.LevelOneChannelName
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GetLevelOneOwner() []*FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner {
	return s.LevelOneOwner
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GetLevelThreeChannelName() *string {
	return s.LevelThreeChannelName
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GetLevelThreeOwner() []*FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner {
	return s.LevelThreeOwner
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GetLevelTwoChannelName() *string {
	return s.LevelTwoChannelName
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) GetLevelTwoOwner() []*FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner {
	return s.LevelTwoOwner
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

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfo) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GetOwnerEmpIdOrTelephone() *string {
	return s.OwnerEmpIdOrTelephone
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GetOwnerName() *string {
	return s.OwnerName
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GetOwnerNickName() *string {
	return s.OwnerNickName
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

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelOneOwner) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GetOwnerEmpIdOrTelephone() *string {
	return s.OwnerEmpIdOrTelephone
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GetOwnerName() *string {
	return s.OwnerName
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GetOwnerNickName() *string {
	return s.OwnerNickName
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

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GetOwnerEmpIdOrTelephone() *string {
	return s.OwnerEmpIdOrTelephone
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GetOwnerName() *string {
	return s.OwnerName
}

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GetOwnerNickName() *string {
	return s.OwnerNickName
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

func (s *FindGuestTicketRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) Validate() error {
	return dara.Validate(s)
}
