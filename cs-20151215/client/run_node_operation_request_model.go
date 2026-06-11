// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNodeOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationAction(v string) *RunNodeOperationRequest
	GetOperationAction() *string
	SetOperationArgs(v []*string) *RunNodeOperationRequest
	GetOperationArgs() []*string
}

type RunNodeOperationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// restart.kubelet
	OperationAction *string   `json:"operationAction,omitempty" xml:"operationAction,omitempty"`
	OperationArgs   []*string `json:"operationArgs,omitempty" xml:"operationArgs,omitempty" type:"Repeated"`
}

func (s RunNodeOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunNodeOperationRequest) GoString() string {
	return s.String()
}

func (s *RunNodeOperationRequest) GetOperationAction() *string {
	return s.OperationAction
}

func (s *RunNodeOperationRequest) GetOperationArgs() []*string {
	return s.OperationArgs
}

func (s *RunNodeOperationRequest) SetOperationAction(v string) *RunNodeOperationRequest {
	s.OperationAction = &v
	return s
}

func (s *RunNodeOperationRequest) SetOperationArgs(v []*string) *RunNodeOperationRequest {
	s.OperationArgs = v
	return s
}

func (s *RunNodeOperationRequest) Validate() error {
	return dara.Validate(s)
}
