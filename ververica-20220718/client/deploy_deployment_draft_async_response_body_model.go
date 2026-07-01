// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployDeploymentDraftAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeployDeploymentDraftAsyncResponseBodyData) *DeployDeploymentDraftAsyncResponseBody
	GetData() *DeployDeploymentDraftAsyncResponseBodyData
	SetErrorCode(v string) *DeployDeploymentDraftAsyncResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeployDeploymentDraftAsyncResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeployDeploymentDraftAsyncResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeployDeploymentDraftAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeployDeploymentDraftAsyncResponseBody
	GetSuccess() *bool
}

type DeployDeploymentDraftAsyncResponseBody struct {
	// The result of the asynchronous operation. This parameter is returned when the request is successful.
	Data *DeployDeploymentDraftAsyncResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// - If success is false, an error code is returned.
	//
	// - If success is true, an empty value is returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// - If success is false, an error message is returned.
	//
	// - If success is true, an empty value is returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code of the business operation. A value other than 200 indicates that the operation failed.
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

func (s DeployDeploymentDraftAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployDeploymentDraftAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncResponseBody) GetData() *DeployDeploymentDraftAsyncResponseBodyData {
	return s.Data
}

func (s *DeployDeploymentDraftAsyncResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeployDeploymentDraftAsyncResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeployDeploymentDraftAsyncResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeployDeploymentDraftAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployDeploymentDraftAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetData(v *DeployDeploymentDraftAsyncResponseBodyData) *DeployDeploymentDraftAsyncResponseBody {
	s.Data = v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetErrorCode(v string) *DeployDeploymentDraftAsyncResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetErrorMessage(v string) *DeployDeploymentDraftAsyncResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetHttpCode(v int32) *DeployDeploymentDraftAsyncResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetRequestId(v string) *DeployDeploymentDraftAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetSuccess(v bool) *DeployDeploymentDraftAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeployDeploymentDraftAsyncResponseBodyData struct {
	// The ID of the asynchronous ticket. Use this ID to query the result of the asynchronous operation.
	//
	// example:
	//
	// b3dcdb25-bf36-457d-92ba-a36077e8****
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
}

func (s DeployDeploymentDraftAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeployDeploymentDraftAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncResponseBodyData) GetTicketId() *string {
	return s.TicketId
}

func (s *DeployDeploymentDraftAsyncResponseBodyData) SetTicketId(v string) *DeployDeploymentDraftAsyncResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
