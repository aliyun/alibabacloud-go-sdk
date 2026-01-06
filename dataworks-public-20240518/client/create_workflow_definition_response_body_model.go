// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateWorkflowDefinitionResponseBody
	GetId() *string
	SetRequestId(v string) *CreateWorkflowDefinitionResponseBody
	GetRequestId() *string
}

type CreateWorkflowDefinitionResponseBody struct {
	// The unique identifier of the workflow.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *CreateWorkflowDefinitionResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateWorkflowDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkflowDefinitionResponseBody) SetId(v string) *CreateWorkflowDefinitionResponseBody {
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
