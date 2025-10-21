// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Member) *GetMemberResponseBody
	GetData() *Member
	SetErrorCode(v string) *GetMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMemberResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetMemberResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMemberResponseBody
	GetSuccess() *bool
}

type GetMemberResponseBody struct {
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
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemberResponseBody) GetData() *Member {
	return s.Data
}

func (s *GetMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMemberResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMemberResponseBody) SetData(v *Member) *GetMemberResponseBody {
	s.Data = v
	return s
}

func (s *GetMemberResponseBody) SetErrorCode(v string) *GetMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMemberResponseBody) SetErrorMessage(v string) *GetMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMemberResponseBody) SetHttpCode(v int32) *GetMemberResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetMemberResponseBody) SetRequestId(v string) *GetMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemberResponseBody) SetSuccess(v bool) *GetMemberResponseBody {
	s.Success = &v
	return s
}

func (s *GetMemberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
