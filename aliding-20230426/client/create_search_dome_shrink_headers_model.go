// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchDomeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSearchDomeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateSearchDomeShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateSearchDomeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateSearchDomeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchDomeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateSearchDomeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSearchDomeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateSearchDomeShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateSearchDomeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSearchDomeShrinkHeaders) SetAccountContextShrink(v string) *CreateSearchDomeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateSearchDomeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
