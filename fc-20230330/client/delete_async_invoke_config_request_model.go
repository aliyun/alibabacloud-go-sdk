// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAsyncInvokeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *DeleteAsyncInvokeConfigRequest
	GetQualifier() *string
}

type DeleteAsyncInvokeConfigRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s DeleteAsyncInvokeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAsyncInvokeConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAsyncInvokeConfigRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeleteAsyncInvokeConfigRequest) SetQualifier(v string) *DeleteAsyncInvokeConfigRequest {
	s.Qualifier = &v
	return s
}

func (s *DeleteAsyncInvokeConfigRequest) Validate() error {
	return dara.Validate(s)
}
