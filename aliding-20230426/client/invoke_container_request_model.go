// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeContainerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *InvokeContainerRequest
	GetOperationId() *string
	SetParams(v string) *InvokeContainerRequest
	GetParams() *string
}

type InvokeContainerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// createSandbox
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	Params      *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s InvokeContainerRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeContainerRequest) GoString() string {
	return s.String()
}

func (s *InvokeContainerRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *InvokeContainerRequest) GetParams() *string {
	return s.Params
}

func (s *InvokeContainerRequest) SetOperationId(v string) *InvokeContainerRequest {
	s.OperationId = &v
	return s
}

func (s *InvokeContainerRequest) SetParams(v string) *InvokeContainerRequest {
	s.Params = &v
	return s
}

func (s *InvokeContainerRequest) Validate() error {
	return dara.Validate(s)
}
