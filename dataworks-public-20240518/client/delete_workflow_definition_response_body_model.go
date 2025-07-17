// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkflowDefinitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkflowDefinitionResponseBody
	GetSuccess() *bool
}

type DeleteWorkflowDefinitionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B17730C0-D959-548A-AE23-E754177CXXXX
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

func (s DeleteWorkflowDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkflowDefinitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkflowDefinitionResponseBody) SetRequestId(v string) *DeleteWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowDefinitionResponseBody) SetSuccess(v bool) *DeleteWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkflowDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
