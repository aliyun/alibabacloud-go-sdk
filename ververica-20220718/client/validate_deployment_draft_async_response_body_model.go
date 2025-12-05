// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDeploymentDraftAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ValidateDeploymentDraftAsyncResponseBodyData) *ValidateDeploymentDraftAsyncResponseBody
	GetData() *ValidateDeploymentDraftAsyncResponseBodyData
	SetErrorCode(v string) *ValidateDeploymentDraftAsyncResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ValidateDeploymentDraftAsyncResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ValidateDeploymentDraftAsyncResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ValidateDeploymentDraftAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidateDeploymentDraftAsyncResponseBody
	GetSuccess() *bool
}

type ValidateDeploymentDraftAsyncResponseBody struct {
	Data *ValidateDeploymentDraftAsyncResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ValidateDeploymentDraftAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateDeploymentDraftAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateDeploymentDraftAsyncResponseBody) GetData() *ValidateDeploymentDraftAsyncResponseBodyData {
	return s.Data
}

func (s *ValidateDeploymentDraftAsyncResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ValidateDeploymentDraftAsyncResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ValidateDeploymentDraftAsyncResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ValidateDeploymentDraftAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateDeploymentDraftAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateDeploymentDraftAsyncResponseBody) SetData(v *ValidateDeploymentDraftAsyncResponseBodyData) *ValidateDeploymentDraftAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponseBody) SetErrorCode(v string) *ValidateDeploymentDraftAsyncResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponseBody) SetErrorMessage(v string) *ValidateDeploymentDraftAsyncResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponseBody) SetHttpCode(v int32) *ValidateDeploymentDraftAsyncResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponseBody) SetRequestId(v string) *ValidateDeploymentDraftAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponseBody) SetSuccess(v bool) *ValidateDeploymentDraftAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ValidateDeploymentDraftAsyncResponseBodyData struct {
	// example:
	//
	// b3dcdb25-bf36-457d-92ba-a36077e8****
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
}

func (s ValidateDeploymentDraftAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ValidateDeploymentDraftAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValidateDeploymentDraftAsyncResponseBodyData) GetTicketId() *string {
	return s.TicketId
}

func (s *ValidateDeploymentDraftAsyncResponseBodyData) SetTicketId(v string) *ValidateDeploymentDraftAsyncResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
