// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResourcePlanWithFlinkConfAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) *GenerateResourcePlanWithFlinkConfAsyncResponseBody
	GetData() *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData
	SetErrorCode(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GenerateResourcePlanWithFlinkConfAsyncResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateResourcePlanWithFlinkConfAsyncResponseBody
	GetSuccess() *bool
}

type GenerateResourcePlanWithFlinkConfAsyncResponseBody struct {
	// 	- If the value of success was true, the asynchronous generation result was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GenerateResourcePlanWithFlinkConfAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) GetData() *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData {
	return s.Data
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetData(v *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.Data = v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetErrorCode(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetErrorMessage(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetHttpCode(v int32) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetRequestId(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetSuccess(v bool) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateResourcePlanWithFlinkConfAsyncResponseBodyData struct {
	// The ID of the ticket for you to query the asynchronous generation result.
	//
	// example:
	//
	// b3dcdb25-bf36-457d-92ba-a36077e8****
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) GetTicketId() *string {
	return s.TicketId
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) SetTicketId(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
