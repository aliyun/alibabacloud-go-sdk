// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateWorkflowDefinitionResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateWorkflowDefinitionResponseBody
	GetRequestId() *string
}

type CreateWorkflowDefinitionResponseBody struct {
	// The ID of the workflow.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0EF298E5-0940-5AC7-9CB0-65025070XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWorkflowDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateWorkflowDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkflowDefinitionResponseBody) SetId(v int64) *CreateWorkflowDefinitionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateWorkflowDefinitionResponseBody) SetRequestId(v string) *CreateWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkflowDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
