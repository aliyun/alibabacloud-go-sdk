// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Member) *UpdateMemberResponseBody
	GetData() *Member
	SetErrorCode(v string) *UpdateMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateMemberResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateMemberResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMemberResponseBody
	GetSuccess() *bool
}

type UpdateMemberResponseBody struct {
	// If the value of success was true, the member that was created was returned. If the value of success was false, a null value was returned.
	Data *Member `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemberResponseBody) GetData() *Member {
	return s.Data
}

func (s *UpdateMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateMemberResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMemberResponseBody) SetData(v *Member) *UpdateMemberResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMemberResponseBody) SetErrorCode(v string) *UpdateMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMemberResponseBody) SetErrorMessage(v string) *UpdateMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMemberResponseBody) SetHttpCode(v int32) *UpdateMemberResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateMemberResponseBody) SetRequestId(v string) *UpdateMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemberResponseBody) SetSuccess(v bool) *UpdateMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMemberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
