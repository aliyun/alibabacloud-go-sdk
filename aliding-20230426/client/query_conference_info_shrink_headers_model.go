// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryConferenceInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryConferenceInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryConferenceInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryConferenceInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryConferenceInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryConferenceInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceInfoShrinkHeaders) SetAccountContextShrink(v string) *QueryConferenceInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryConferenceInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
