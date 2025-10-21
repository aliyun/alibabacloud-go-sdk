// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Member) *CreateMemberResponseBody
	GetData() *Member
	SetErrorCode(v string) *CreateMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateMemberResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateMemberResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMemberResponseBody
	GetSuccess() *bool
}

type CreateMemberResponseBody struct {
	// 	- If the value of success was false, a null value was returned.
	//
	// 	- If the value of success was true, the authorization information was returned.
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
	// F989CA70-2925-5A94-92B7-20F5762B71C8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemberResponseBody) GetData() *Member {
	return s.Data
}

func (s *CreateMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateMemberResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMemberResponseBody) SetData(v *Member) *CreateMemberResponseBody {
	s.Data = v
	return s
}

func (s *CreateMemberResponseBody) SetErrorCode(v string) *CreateMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMemberResponseBody) SetErrorMessage(v string) *CreateMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMemberResponseBody) SetHttpCode(v int32) *CreateMemberResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateMemberResponseBody) SetRequestId(v string) *CreateMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemberResponseBody) SetSuccess(v bool) *CreateMemberResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMemberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
