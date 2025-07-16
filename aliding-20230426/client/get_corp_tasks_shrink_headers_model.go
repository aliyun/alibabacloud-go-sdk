// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpTasksShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCorpTasksShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetCorpTasksShrinkHeaders
	GetAccountContextShrink() *string
}

type GetCorpTasksShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetCorpTasksShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCorpTasksShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpTasksShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCorpTasksShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetCorpTasksShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetCorpTasksShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpTasksShrinkHeaders) SetAccountContextShrink(v string) *GetCorpTasksShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetCorpTasksShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
