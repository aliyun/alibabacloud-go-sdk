// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpAccomplishmentTasksShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCorpAccomplishmentTasksShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetCorpAccomplishmentTasksShrinkHeaders
	GetAccountContextShrink() *string
}

type GetCorpAccomplishmentTasksShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetCorpAccomplishmentTasksShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCorpAccomplishmentTasksShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCorpAccomplishmentTasksShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetCorpAccomplishmentTasksShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetCorpAccomplishmentTasksShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpAccomplishmentTasksShrinkHeaders) SetAccountContextShrink(v string) *GetCorpAccomplishmentTasksShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetCorpAccomplishmentTasksShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
