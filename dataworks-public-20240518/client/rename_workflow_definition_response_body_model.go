// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameWorkflowDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenameWorkflowDefinitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenameWorkflowDefinitionResponseBody
	GetSuccess() *bool
}

type RenameWorkflowDefinitionResponseBody struct {
	// The request ID. You can troubleshoot issues based on the ID.
	//
	// example:
	//
	// 975BD43D-C421-595C-A29C-565A8AD5XXXX
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

func (s RenameWorkflowDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameWorkflowDefinitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenameWorkflowDefinitionResponseBody) SetRequestId(v string) *RenameWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameWorkflowDefinitionResponseBody) SetSuccess(v bool) *RenameWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

func (s *RenameWorkflowDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
