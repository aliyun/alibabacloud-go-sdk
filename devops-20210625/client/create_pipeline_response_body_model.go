// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreatePipelineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreatePipelineResponseBody
	GetErrorMessage() *string
	SetPipelinId(v int64) *CreatePipelineResponseBody
	GetPipelinId() *int64
	SetRequestId(v string) *CreatePipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePipelineResponseBody
	GetSuccess() *bool
}

type CreatePipelineResponseBody struct {
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
	// 11XXX
	PipelinId *int64 `json:"pipelinId,omitempty" xml:"pipelinId,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreatePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreatePipelineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreatePipelineResponseBody) GetPipelinId() *int64 {
	return s.PipelinId
}

func (s *CreatePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePipelineResponseBody) SetErrorCode(v string) *CreatePipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreatePipelineResponseBody) SetErrorMessage(v string) *CreatePipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePipelineResponseBody) SetPipelinId(v int64) *CreatePipelineResponseBody {
	s.PipelinId = &v
	return s
}

func (s *CreatePipelineResponseBody) SetRequestId(v string) *CreatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineResponseBody) SetSuccess(v bool) *CreatePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
