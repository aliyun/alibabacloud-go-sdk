// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iInvokeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v io.Reader) *InvokeFunctionRequest
	GetBody() io.Reader
	SetQualifier(v string) *InvokeFunctionRequest
	GetQualifier() *string
}

type InvokeFunctionRequest struct {
	// The request parameters of function invocation.
	//
	// example:
	//
	// event
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s InvokeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeFunctionRequest) GoString() string {
	return s.String()
}

func (s *InvokeFunctionRequest) GetBody() io.Reader {
	return s.Body
}

func (s *InvokeFunctionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *InvokeFunctionRequest) SetBody(v io.Reader) *InvokeFunctionRequest {
	s.Body = v
	return s
}

func (s *InvokeFunctionRequest) SetQualifier(v string) *InvokeFunctionRequest {
	s.Qualifier = &v
	return s
}

func (s *InvokeFunctionRequest) Validate() error {
	return dara.Validate(s)
}
