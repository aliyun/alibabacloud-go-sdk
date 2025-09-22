// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindGuestCredentialsRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*FindGuestCredentialsRecordResponseBodyData) *FindGuestCredentialsRecordResponseBody
	GetData() []*FindGuestCredentialsRecordResponseBodyData
	SetErrCode(v string) *FindGuestCredentialsRecordResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *FindGuestCredentialsRecordResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *FindGuestCredentialsRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindGuestCredentialsRecordResponseBody
	GetSuccess() *bool
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
	return dara.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBody) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBody) GetData() []*FindGuestCredentialsRecordResponseBodyData {
	return s.Data
}

func (s *FindGuestCredentialsRecordResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *FindGuestCredentialsRecordResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *FindGuestCredentialsRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindGuestCredentialsRecordResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *FindGuestCredentialsRecordResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetAdmin() map[string]interface{} {
	return s.Admin
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetChannelId() *int64 {
	return s.ChannelId
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetChannelLevelInfo() *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo {
	return s.ChannelLevelInfo
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetCompanyName() *string {
	return s.CompanyName
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetCredentialsCode() *string {
	return s.CredentialsCode
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetCredentialsName() *string {
	return s.CredentialsName
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetIdNumber() *string {
	return s.IdNumber
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetIdType() *string {
	return s.IdType
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetName() *string {
	return s.Name
}

func (s *FindGuestCredentialsRecordResponseBodyData) GetStatus() *int32 {
	return s.Status
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

func (s *FindGuestCredentialsRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GetChannelId() *int64 {
	return s.ChannelId
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GetChannelName() *string {
	return s.ChannelName
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GetLevelOneChannelName() *string {
	return s.LevelOneChannelName
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GetLevelOneOwner() []*FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner {
	return s.LevelOneOwner
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GetLevelThreeChannelName() *string {
	return s.LevelThreeChannelName
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GetLevelThreeOwner() []*FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner {
	return s.LevelThreeOwner
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GetLevelTwoChannelName() *string {
	return s.LevelTwoChannelName
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) GetLevelTwoOwner() []*FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner {
	return s.LevelTwoOwner
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

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfo) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GetOwnerEmpIdOrTelephone() *string {
	return s.OwnerEmpIdOrTelephone
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GetOwnerName() *string {
	return s.OwnerName
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) GetOwnerNickName() *string {
	return s.OwnerNickName
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

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelOneOwner) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GetOwnerEmpIdOrTelephone() *string {
	return s.OwnerEmpIdOrTelephone
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GetOwnerName() *string {
	return s.OwnerName
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) GetOwnerNickName() *string {
	return s.OwnerNickName
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

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelThreeOwner) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GetOwnerEmpIdOrTelephone() *string {
	return s.OwnerEmpIdOrTelephone
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GetOwnerName() *string {
	return s.OwnerName
}

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) GetOwnerNickName() *string {
	return s.OwnerNickName
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

func (s *FindGuestCredentialsRecordResponseBodyDataChannelLevelInfoLevelTwoOwner) Validate() error {
	return dara.Validate(s)
}
