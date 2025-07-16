// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConvExtensionShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateConvExtensionShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateConvExtensionShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateConvExtensionShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateConvExtensionShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateConvExtensionShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateConvExtensionShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateConvExtensionShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateConvExtensionShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateConvExtensionShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateConvExtensionShrinkHeaders) SetAccountContextShrink(v string) *UpdateConvExtensionShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateConvExtensionShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
