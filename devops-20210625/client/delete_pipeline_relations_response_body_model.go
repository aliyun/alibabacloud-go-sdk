// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeletePipelineRelationsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeletePipelineRelationsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeletePipelineRelationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePipelineRelationsResponseBody
	GetSuccess() *bool
}

type DeletePipelineRelationsResponseBody struct {
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

func (s DeletePipelineRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineRelationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeletePipelineRelationsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeletePipelineRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePipelineRelationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePipelineRelationsResponseBody) SetErrorCode(v string) *DeletePipelineRelationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeletePipelineRelationsResponseBody) SetErrorMessage(v string) *DeletePipelineRelationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePipelineRelationsResponseBody) SetRequestId(v string) *DeletePipelineRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineRelationsResponseBody) SetSuccess(v bool) *DeletePipelineRelationsResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePipelineRelationsResponseBody) Validate() error {
	return dara.Validate(s)
}
