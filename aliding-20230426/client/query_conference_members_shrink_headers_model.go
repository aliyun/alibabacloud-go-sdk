// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceMembersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryConferenceMembersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryConferenceMembersShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryConferenceMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryConferenceMembersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryConferenceMembersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryConferenceMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceMembersShrinkHeaders) SetAccountContextShrink(v string) *QueryConferenceMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryConferenceMembersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
