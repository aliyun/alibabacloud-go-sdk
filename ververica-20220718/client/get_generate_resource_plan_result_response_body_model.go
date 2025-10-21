// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGenerateResourcePlanResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AsyncResourcePlanOperationResult) *GetGenerateResourcePlanResultResponseBody
	GetData() *AsyncResourcePlanOperationResult
	SetErrorCode(v string) *GetGenerateResourcePlanResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetGenerateResourcePlanResultResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetGenerateResourcePlanResultResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetGenerateResourcePlanResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGenerateResourcePlanResultResponseBody
	GetSuccess() *bool
}

type GetGenerateResourcePlanResultResponseBody struct {
	// 	- If the value of success was true, the asynchronous generation result was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *AsyncResourcePlanOperationResult `json:"data,omitempty" xml:"data,omitempty"`
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

func (s GetGenerateResourcePlanResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGenerateResourcePlanResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetGenerateResourcePlanResultResponseBody) GetData() *AsyncResourcePlanOperationResult {
	return s.Data
}

func (s *GetGenerateResourcePlanResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetGenerateResourcePlanResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetGenerateResourcePlanResultResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetGenerateResourcePlanResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGenerateResourcePlanResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGenerateResourcePlanResultResponseBody) SetData(v *AsyncResourcePlanOperationResult) *GetGenerateResourcePlanResultResponseBody {
	s.Data = v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetErrorCode(v string) *GetGenerateResourcePlanResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetErrorMessage(v string) *GetGenerateResourcePlanResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetHttpCode(v int32) *GetGenerateResourcePlanResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetRequestId(v string) *GetGenerateResourcePlanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetSuccess(v bool) *GetGenerateResourcePlanResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
