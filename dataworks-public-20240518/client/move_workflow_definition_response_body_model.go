// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveWorkflowDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveWorkflowDefinitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MoveWorkflowDefinitionResponseBody
	GetSuccess() *bool
}

type MoveWorkflowDefinitionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 05ADAF4F-7709-5FB1-B606-3513483FXXXX
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

func (s MoveWorkflowDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveWorkflowDefinitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveWorkflowDefinitionResponseBody) SetRequestId(v string) *MoveWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveWorkflowDefinitionResponseBody) SetSuccess(v bool) *MoveWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

func (s *MoveWorkflowDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
