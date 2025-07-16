// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStatusShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateStatusShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateStatusShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateStatusShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateStatusShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateStatusShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateStatusShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateStatusShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateStatusShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateStatusShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateStatusShrinkHeaders) SetAccountContextShrink(v string) *UpdateStatusShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateStatusShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
