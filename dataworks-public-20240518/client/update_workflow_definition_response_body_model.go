// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkflowDefinitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkflowDefinitionResponseBody
	GetSuccess() *bool
}

type UpdateWorkflowDefinitionResponseBody struct {
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 20BF7E80-668A-5620-8AD8-879B8FEAXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkflowDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowDefinitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkflowDefinitionResponseBody) SetRequestId(v string) *UpdateWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowDefinitionResponseBody) SetSuccess(v bool) *UpdateWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkflowDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
