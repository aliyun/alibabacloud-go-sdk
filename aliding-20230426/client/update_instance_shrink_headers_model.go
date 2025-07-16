// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateInstanceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateInstanceShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateInstanceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateInstanceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInstanceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateInstanceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateInstanceShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateInstanceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInstanceShrinkHeaders) SetAccountContextShrink(v string) *UpdateInstanceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateInstanceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
