// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserInvitationCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateUserInvitationCodeResponseBody
	GetCode() *string
	SetData(v []*CreateUserInvitationCodeResponseBodyData) *CreateUserInvitationCodeResponseBody
	GetData() []*CreateUserInvitationCodeResponseBodyData
	SetHttpStatusCode(v int32) *CreateUserInvitationCodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateUserInvitationCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateUserInvitationCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUserInvitationCodeResponseBody
	GetSuccess() *bool
}

type CreateUserInvitationCodeResponseBody struct {
	// The business error code. This field may be empty if the call is successful. If the call fails, a value such as PARAM_ERROR indicates that the parameter value is invalid.
	//
	// example:
	//
	// PARAM_ERROR
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The array of created management invitation code records. Each user corresponds to one record. An array is returned even for single-user creation.
	Data []*CreateUserInvitationCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code of the backend business response. A value of 200 indicates success. A value of 400 indicates parameter or business validation failure. A value of 500 indicates a system error.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business error message. This field may be empty if the call is successful. If the call fails, this field describes the cause of the error.
	//
	// example:
	//
	// PARAM_ERROR
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, used to locate call logs and troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the business processing is successful. A value of true indicates success. A value of false indicates failure.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUserInvitationCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserInvitationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserInvitationCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateUserInvitationCodeResponseBody) GetData() []*CreateUserInvitationCodeResponseBodyData {
	return s.Data
}

func (s *CreateUserInvitationCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateUserInvitationCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUserInvitationCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserInvitationCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUserInvitationCodeResponseBody) SetCode(v string) *CreateUserInvitationCodeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBody) SetData(v []*CreateUserInvitationCodeResponseBodyData) *CreateUserInvitationCodeResponseBody {
	s.Data = v
	return s
}

func (s *CreateUserInvitationCodeResponseBody) SetHttpStatusCode(v int32) *CreateUserInvitationCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBody) SetMessage(v string) *CreateUserInvitationCodeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBody) SetRequestId(v string) *CreateUserInvitationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBody) SetSuccess(v bool) *CreateUserInvitationCodeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBody) Validate() error {
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

type CreateUserInvitationCodeResponseBodyData struct {
	// The binding time when the terminal uses the code to complete management. This field is empty for a newly created code because it is not yet bound to a terminal.
	//
	// example:
	//
	// null
	BindTime *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	// The logon identifier (EndUserId) of the user corresponding to the management code.
	//
	// example:
	//
	// user-001
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The expiration time of the management code calculated by the server. This field is empty if the code is permanently valid. To determine whether the code has expired, use the Expired field.
	//
	// example:
	//
	// null
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the code has expired as calculated by the server. A value of true indicates that the validity period has expired. A value of false indicates that the code has not expired or is permanently valid. This field is independent of the usage status in the Status field.
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The creation time of the management code, generated by the server.
	//
	// example:
	//
	// 2026-10-15T08:00:00.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The generated eight-digit terminal management invitation code. It is processed as a string and leading zeros are preserved.
	//
	// example:
	//
	// 01234567
	InvitationCode *string `json:"InvitationCode,omitempty" xml:"InvitationCode,omitempty"`
	// The client type of the bound terminal. This field is empty for a newly created code because it is not yet bound to a terminal.
	//
	// example:
	//
	// null
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The usage status of the management code. Valid values:
	//
	// - ACTIVE: The code is unused.
	//
	// - USED: The code is used.
	//
	// The creation response returns ACTIVE. Expiration does not change the usage status. Use the Expired field to determine whether the code has expired.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the terminal group specified for the management code.
	//
	// example:
	//
	// tg-example
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	// The display name of the terminal group. This creation operation does not populate this display field, so it may be empty in the creation response.
	//
	// example:
	//
	// null
	TerminalGroupName *string `json:"TerminalGroupName,omitempty" xml:"TerminalGroupName,omitempty"`
	// The display name of the user. This creation operation does not populate this display field, so it may be empty in the creation response.
	//
	// example:
	//
	// null
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The UUID of the terminal that uses the code to complete management. This field is empty for a newly created code because it is not yet bound to a terminal.
	//
	// example:
	//
	// null
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The validity period of the management code, in minutes. The value is the same as the request value. A value of -1 indicates that the code is permanently valid.
	//
	// example:
	//
	// 1440
	ValidityMinutes *int32 `json:"ValidityMinutes,omitempty" xml:"ValidityMinutes,omitempty"`
}

func (s CreateUserInvitationCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateUserInvitationCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUserInvitationCodeResponseBodyData) GetBindTime() *string {
	return s.BindTime
}

func (s *CreateUserInvitationCodeResponseBodyData) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CreateUserInvitationCodeResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateUserInvitationCodeResponseBodyData) GetExpired() *bool {
	return s.Expired
}

func (s *CreateUserInvitationCodeResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateUserInvitationCodeResponseBodyData) GetInvitationCode() *string {
	return s.InvitationCode
}

func (s *CreateUserInvitationCodeResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *CreateUserInvitationCodeResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateUserInvitationCodeResponseBodyData) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *CreateUserInvitationCodeResponseBodyData) GetTerminalGroupName() *string {
	return s.TerminalGroupName
}

func (s *CreateUserInvitationCodeResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserInvitationCodeResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *CreateUserInvitationCodeResponseBodyData) GetValidityMinutes() *int32 {
	return s.ValidityMinutes
}

func (s *CreateUserInvitationCodeResponseBodyData) SetBindTime(v string) *CreateUserInvitationCodeResponseBodyData {
	s.BindTime = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetEndUserId(v string) *CreateUserInvitationCodeResponseBodyData {
	s.EndUserId = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetExpireTime(v string) *CreateUserInvitationCodeResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetExpired(v bool) *CreateUserInvitationCodeResponseBodyData {
	s.Expired = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetGmtCreate(v string) *CreateUserInvitationCodeResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetInvitationCode(v string) *CreateUserInvitationCodeResponseBodyData {
	s.InvitationCode = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetPlatform(v string) *CreateUserInvitationCodeResponseBodyData {
	s.Platform = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetStatus(v string) *CreateUserInvitationCodeResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetTerminalGroupId(v string) *CreateUserInvitationCodeResponseBodyData {
	s.TerminalGroupId = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetTerminalGroupName(v string) *CreateUserInvitationCodeResponseBodyData {
	s.TerminalGroupName = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetUserName(v string) *CreateUserInvitationCodeResponseBodyData {
	s.UserName = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetUuid(v string) *CreateUserInvitationCodeResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) SetValidityMinutes(v int32) *CreateUserInvitationCodeResponseBodyData {
	s.ValidityMinutes = &v
	return s
}

func (s *CreateUserInvitationCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
