// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateLiveShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateLiveShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateLiveShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateLiveShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateLiveShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateLiveShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateLiveShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateLiveShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateLiveShrinkHeaders) SetAccountContextShrink(v string) *UpdateLiveShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateLiveShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
