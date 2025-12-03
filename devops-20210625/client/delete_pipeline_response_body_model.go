// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeletePipelineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeletePipelineResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeletePipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePipelineResponseBody
	GetSuccess() *bool
}

type DeletePipelineResponseBody struct {
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

func (s DeletePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeletePipelineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeletePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePipelineResponseBody) SetErrorCode(v string) *DeletePipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeletePipelineResponseBody) SetErrorMessage(v string) *DeletePipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePipelineResponseBody) SetRequestId(v string) *DeletePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineResponseBody) SetSuccess(v bool) *DeletePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
