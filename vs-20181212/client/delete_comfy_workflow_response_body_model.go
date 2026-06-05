// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComfyWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteComfyWorkflowResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteComfyWorkflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteComfyWorkflowResponseBody
	GetRequestId() *string
}

type DeleteComfyWorkflowResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteComfyWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComfyWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComfyWorkflowResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteComfyWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteComfyWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComfyWorkflowResponseBody) SetCode(v string) *DeleteComfyWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteComfyWorkflowResponseBody) SetMessage(v string) *DeleteComfyWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteComfyWorkflowResponseBody) SetRequestId(v string) *DeleteComfyWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComfyWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
