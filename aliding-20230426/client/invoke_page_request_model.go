// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *InvokePageRequest
	GetOperationId() *string
	SetParams(v string) *InvokePageRequest
	GetParams() *string
}

type InvokePageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// createPage
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	Params      *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s InvokePageRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokePageRequest) GoString() string {
	return s.String()
}

func (s *InvokePageRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *InvokePageRequest) GetParams() *string {
	return s.Params
}

func (s *InvokePageRequest) SetOperationId(v string) *InvokePageRequest {
	s.OperationId = &v
	return s
}

func (s *InvokePageRequest) SetParams(v string) *InvokePageRequest {
	s.Params = &v
	return s
}

func (s *InvokePageRequest) Validate() error {
	return dara.Validate(s)
}
