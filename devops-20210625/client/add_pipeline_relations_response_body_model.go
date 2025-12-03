// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPipelineRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddPipelineRelationsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddPipelineRelationsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddPipelineRelationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddPipelineRelationsResponseBody
	GetSuccess() *bool
}

type AddPipelineRelationsResponseBody struct {
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
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddPipelineRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPipelineRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *AddPipelineRelationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddPipelineRelationsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddPipelineRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPipelineRelationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddPipelineRelationsResponseBody) SetErrorCode(v string) *AddPipelineRelationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddPipelineRelationsResponseBody) SetErrorMessage(v string) *AddPipelineRelationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddPipelineRelationsResponseBody) SetRequestId(v string) *AddPipelineRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPipelineRelationsResponseBody) SetSuccess(v bool) *AddPipelineRelationsResponseBody {
	s.Success = &v
	return s
}

func (s *AddPipelineRelationsResponseBody) Validate() error {
	return dara.Validate(s)
}
