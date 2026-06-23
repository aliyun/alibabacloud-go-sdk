// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserInfoToMsenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMpaasUserInfoShareResponse(v *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) *QueryUserInfoToMsenceResponseBody
	GetMpaasUserInfoShareResponse() *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse
	SetRequestId(v string) *QueryUserInfoToMsenceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryUserInfoToMsenceResponseBody
	GetResultCode() *string
	SetResultMsg(v string) *QueryUserInfoToMsenceResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *QueryUserInfoToMsenceResponseBody
	GetSuccess() *bool
}

type QueryUserInfoToMsenceResponseBody struct {
	MpaasUserInfoShareResponse *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse `json:"MpaasUserInfoShareResponse,omitempty" xml:"MpaasUserInfoShareResponse,omitempty" type:"Struct"`
	RequestId                  *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                 *string                                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMsg                  *string                                                      `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success                    *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserInfoToMsenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoToMsenceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserInfoToMsenceResponseBody) GetMpaasUserInfoShareResponse() *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse {
	return s.MpaasUserInfoShareResponse
}

func (s *QueryUserInfoToMsenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserInfoToMsenceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryUserInfoToMsenceResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryUserInfoToMsenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserInfoToMsenceResponseBody) SetMpaasUserInfoShareResponse(v *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) *QueryUserInfoToMsenceResponseBody {
	s.MpaasUserInfoShareResponse = v
	return s
}

func (s *QueryUserInfoToMsenceResponseBody) SetRequestId(v string) *QueryUserInfoToMsenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserInfoToMsenceResponseBody) SetResultCode(v string) *QueryUserInfoToMsenceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryUserInfoToMsenceResponseBody) SetResultMsg(v string) *QueryUserInfoToMsenceResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *QueryUserInfoToMsenceResponseBody) SetSuccess(v bool) *QueryUserInfoToMsenceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserInfoToMsenceResponseBody) Validate() error {
	if s.MpaasUserInfoShareResponse != nil {
		if err := s.MpaasUserInfoShareResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse struct {
	Avatar   *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Gender   *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) GetAvatar() *string {
	return s.Avatar
}

func (s *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) GetGender() *string {
	return s.Gender
}

func (s *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) GetNickName() *string {
	return s.NickName
}

func (s *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) SetAvatar(v string) *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse {
	s.Avatar = &v
	return s
}

func (s *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) SetGender(v string) *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse {
	s.Gender = &v
	return s
}

func (s *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) SetNickName(v string) *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse {
	s.NickName = &v
	return s
}

func (s *QueryUserInfoToMsenceResponseBodyMpaasUserInfoShareResponse) Validate() error {
	return dara.Validate(s)
}
