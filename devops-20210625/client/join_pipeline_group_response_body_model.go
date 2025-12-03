// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinPipelineGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *JoinPipelineGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *JoinPipelineGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *JoinPipelineGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *JoinPipelineGroupResponseBody
	GetSuccess() *bool
}

type JoinPipelineGroupResponseBody struct {
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

func (s JoinPipelineGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinPipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *JoinPipelineGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *JoinPipelineGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *JoinPipelineGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinPipelineGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *JoinPipelineGroupResponseBody) SetErrorCode(v string) *JoinPipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *JoinPipelineGroupResponseBody) SetErrorMessage(v string) *JoinPipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *JoinPipelineGroupResponseBody) SetRequestId(v string) *JoinPipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinPipelineGroupResponseBody) SetSuccess(v bool) *JoinPipelineGroupResponseBody {
	s.Success = &v
	return s
}

func (s *JoinPipelineGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
