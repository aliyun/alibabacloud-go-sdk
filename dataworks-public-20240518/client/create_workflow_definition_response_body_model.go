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
	// 	Notice: This field was of the Long type in SDK versions earlier than 8.0.0 and is of the String type in SDK 8.0.0 and later. **This change does not affect normal SDK usage, and the parameter is still returned in the type defined in the SDK**. Only when you upgrade across SDK version 8.0.0, the type change may cause project compilation failures, and you need to manually correct the data type.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. Used for locating logs and troubleshooting issues.
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
