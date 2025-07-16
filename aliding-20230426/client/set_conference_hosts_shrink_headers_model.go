// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetConferenceHostsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SetConferenceHostsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SetConferenceHostsShrinkHeaders
	GetAccountContextShrink() *string
}

type SetConferenceHostsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SetConferenceHostsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SetConferenceHostsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SetConferenceHostsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SetConferenceHostsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SetConferenceHostsShrinkHeaders) SetCommonHeaders(v map[string]*string) *SetConferenceHostsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetConferenceHostsShrinkHeaders) SetAccountContextShrink(v string) *SetConferenceHostsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SetConferenceHostsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
