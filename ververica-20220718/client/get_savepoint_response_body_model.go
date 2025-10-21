// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavepointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Savepoint) *GetSavepointResponseBody
	GetData() *Savepoint
	SetErrorCode(v string) *GetSavepointResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSavepointResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetSavepointResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetSavepointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSavepointResponseBody
	GetSuccess() *bool
}

type GetSavepointResponseBody struct {
	// 	- If the value of success was true, the savepoint information was returned.
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

func (s GetSavepointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSavepointResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavepointResponseBody) GetData() *Savepoint {
	return s.Data
}

func (s *GetSavepointResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSavepointResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSavepointResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetSavepointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSavepointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSavepointResponseBody) SetData(v *Savepoint) *GetSavepointResponseBody {
	s.Data = v
	return s
}

func (s *GetSavepointResponseBody) SetErrorCode(v string) *GetSavepointResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSavepointResponseBody) SetErrorMessage(v string) *GetSavepointResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSavepointResponseBody) SetHttpCode(v int32) *GetSavepointResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetSavepointResponseBody) SetRequestId(v string) *GetSavepointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSavepointResponseBody) SetSuccess(v bool) *GetSavepointResponseBody {
	s.Success = &v
	return s
}

func (s *GetSavepointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
