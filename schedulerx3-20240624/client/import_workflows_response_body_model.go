// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImportWorkflowsResponseBody
	GetCode() *int32
	SetMessage(v string) *ImportWorkflowsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportWorkflowsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportWorkflowsResponseBody
	GetSuccess() *bool
}

type ImportWorkflowsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 27B1345D-5F71-5972-8E4C-AABA6C6232F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportWorkflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportWorkflowsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImportWorkflowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportWorkflowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportWorkflowsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportWorkflowsResponseBody) SetCode(v int32) *ImportWorkflowsResponseBody {
	s.Code = &v
	return s
}

func (s *ImportWorkflowsResponseBody) SetMessage(v string) *ImportWorkflowsResponseBody {
	s.Message = &v
	return s
}

func (s *ImportWorkflowsResponseBody) SetRequestId(v string) *ImportWorkflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportWorkflowsResponseBody) SetSuccess(v bool) *ImportWorkflowsResponseBody {
	s.Success = &v
	return s
}

func (s *ImportWorkflowsResponseBody) Validate() error {
	return dara.Validate(s)
}
