// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartInstanceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *StartInstanceShrinkHeaders
	GetAccountContextShrink() *string
}

type StartInstanceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s StartInstanceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *StartInstanceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartInstanceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *StartInstanceShrinkHeaders) SetCommonHeaders(v map[string]*string) *StartInstanceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartInstanceShrinkHeaders) SetAccountContextShrink(v string) *StartInstanceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *StartInstanceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
