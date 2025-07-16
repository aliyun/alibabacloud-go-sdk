// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateLiveShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateLiveShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateLiveShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateLiveShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateLiveShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateLiveShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateLiveShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateLiveShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateLiveShrinkHeaders) SetAccountContextShrink(v string) *CreateLiveShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateLiveShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
