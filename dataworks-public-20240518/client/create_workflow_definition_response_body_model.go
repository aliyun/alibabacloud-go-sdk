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
	// The unique identifier of the DataStudio workflow.
	//
	// 	Notice: This field is of the Long type before SDK 8.0.0 and of the String type in SDK 8.0.0 and later. **This change does not affect the normal use of the SDK, and the parameter is still returned based on the type defined in the SDK**. The type change may cause project compilation failures only when you upgrade the SDK across version 8.0.0. In this case, you must manually correct the data type.</notice>
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 0EF298E5-0940-5AC7-9CB0-6502507****
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
