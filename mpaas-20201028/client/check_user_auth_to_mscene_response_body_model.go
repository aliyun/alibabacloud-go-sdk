// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserAuthToMsceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMpaasUserAuthCheckResponse(v *CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse) *CheckUserAuthToMsceneResponseBody
	GetMpaasUserAuthCheckResponse() *CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse
	SetRequestId(v string) *CheckUserAuthToMsceneResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CheckUserAuthToMsceneResponseBody
	GetResultCode() *string
	SetResultMsg(v string) *CheckUserAuthToMsceneResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *CheckUserAuthToMsceneResponseBody
	GetSuccess() *bool
}

type CheckUserAuthToMsceneResponseBody struct {
	MpaasUserAuthCheckResponse *CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse `json:"MpaasUserAuthCheckResponse,omitempty" xml:"MpaasUserAuthCheckResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// SUCCESS
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckUserAuthToMsceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUserAuthToMsceneResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserAuthToMsceneResponseBody) GetMpaasUserAuthCheckResponse() *CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse {
	return s.MpaasUserAuthCheckResponse
}

func (s *CheckUserAuthToMsceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUserAuthToMsceneResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CheckUserAuthToMsceneResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CheckUserAuthToMsceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckUserAuthToMsceneResponseBody) SetMpaasUserAuthCheckResponse(v *CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse) *CheckUserAuthToMsceneResponseBody {
	s.MpaasUserAuthCheckResponse = v
	return s
}

func (s *CheckUserAuthToMsceneResponseBody) SetRequestId(v string) *CheckUserAuthToMsceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserAuthToMsceneResponseBody) SetResultCode(v string) *CheckUserAuthToMsceneResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CheckUserAuthToMsceneResponseBody) SetResultMsg(v string) *CheckUserAuthToMsceneResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *CheckUserAuthToMsceneResponseBody) SetSuccess(v bool) *CheckUserAuthToMsceneResponseBody {
	s.Success = &v
	return s
}

func (s *CheckUserAuthToMsceneResponseBody) Validate() error {
	if s.MpaasUserAuthCheckResponse != nil {
		if err := s.MpaasUserAuthCheckResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse struct {
	// example:
	//
	// true
	Matched *bool `json:"Matched,omitempty" xml:"Matched,omitempty"`
}

func (s CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse) GoString() string {
	return s.String()
}

func (s *CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse) GetMatched() *bool {
	return s.Matched
}

func (s *CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse) SetMatched(v bool) *CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse {
	s.Matched = &v
	return s
}

func (s *CheckUserAuthToMsceneResponseBodyMpaasUserAuthCheckResponse) Validate() error {
	return dara.Validate(s)
}
