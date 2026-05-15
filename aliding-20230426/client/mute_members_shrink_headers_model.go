// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteMembersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MuteMembersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *MuteMembersShrinkHeaders
	GetAccountContextShrink() *string
}

type MuteMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s MuteMembersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s MuteMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *MuteMembersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MuteMembersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *MuteMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *MuteMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MuteMembersShrinkHeaders) SetAccountContextShrink(v string) *MuteMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *MuteMembersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
