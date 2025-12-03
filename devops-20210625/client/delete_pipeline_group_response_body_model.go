// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeletePipelineGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeletePipelineGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeletePipelineGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePipelineGroupResponseBody
	GetSuccess() *bool
}

type DeletePipelineGroupResponseBody struct {
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

func (s DeletePipelineGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeletePipelineGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeletePipelineGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePipelineGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePipelineGroupResponseBody) SetErrorCode(v string) *DeletePipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeletePipelineGroupResponseBody) SetErrorMessage(v string) *DeletePipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePipelineGroupResponseBody) SetRequestId(v string) *DeletePipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineGroupResponseBody) SetSuccess(v bool) *DeletePipelineGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePipelineGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
