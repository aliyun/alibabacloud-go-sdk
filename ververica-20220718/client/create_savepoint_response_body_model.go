// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavepointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Savepoint) *CreateSavepointResponseBody
	GetData() *Savepoint
	SetErrorCode(v string) *CreateSavepointResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateSavepointResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateSavepointResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateSavepointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSavepointResponseBody
	GetSuccess() *bool
}

type CreateSavepointResponseBody struct {
	// 	- If the value of success was true, the savepoint that was created was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Savepoint `json:"data,omitempty" xml:"data,omitempty"`
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
	// The value was fixed to 200.
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

func (s CreateSavepointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSavepointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSavepointResponseBody) GetData() *Savepoint {
	return s.Data
}

func (s *CreateSavepointResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateSavepointResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateSavepointResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateSavepointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSavepointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSavepointResponseBody) SetData(v *Savepoint) *CreateSavepointResponseBody {
	s.Data = v
	return s
}

func (s *CreateSavepointResponseBody) SetErrorCode(v string) *CreateSavepointResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSavepointResponseBody) SetErrorMessage(v string) *CreateSavepointResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSavepointResponseBody) SetHttpCode(v int32) *CreateSavepointResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateSavepointResponseBody) SetRequestId(v string) *CreateSavepointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSavepointResponseBody) SetSuccess(v bool) *CreateSavepointResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSavepointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
