// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAsyncInvokeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PutAsyncInvokeConfigInput) *PutAsyncInvokeConfigRequest
	GetBody() *PutAsyncInvokeConfigInput
	SetQualifier(v string) *PutAsyncInvokeConfigRequest
	GetQualifier() *string
}

type PutAsyncInvokeConfigRequest struct {
	// The configurations of asynchronous function invocations.
	//
	// This parameter is required.
	Body *PutAsyncInvokeConfigInput `json:"body,omitempty" xml:"body,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s PutAsyncInvokeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PutAsyncInvokeConfigRequest) GoString() string {
	return s.String()
}

func (s *PutAsyncInvokeConfigRequest) GetBody() *PutAsyncInvokeConfigInput {
	return s.Body
}

func (s *PutAsyncInvokeConfigRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *PutAsyncInvokeConfigRequest) SetBody(v *PutAsyncInvokeConfigInput) *PutAsyncInvokeConfigRequest {
	s.Body = v
	return s
}

func (s *PutAsyncInvokeConfigRequest) SetQualifier(v string) *PutAsyncInvokeConfigRequest {
	s.Qualifier = &v
	return s
}

func (s *PutAsyncInvokeConfigRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
