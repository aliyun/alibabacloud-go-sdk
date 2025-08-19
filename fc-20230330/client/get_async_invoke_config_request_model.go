// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncInvokeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *GetAsyncInvokeConfigRequest
	GetQualifier() *string
}

type GetAsyncInvokeConfigRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetAsyncInvokeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncInvokeConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncInvokeConfigRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *GetAsyncInvokeConfigRequest) SetQualifier(v string) *GetAsyncInvokeConfigRequest {
	s.Qualifier = &v
	return s
}

func (s *GetAsyncInvokeConfigRequest) Validate() error {
	return dara.Validate(s)
}
