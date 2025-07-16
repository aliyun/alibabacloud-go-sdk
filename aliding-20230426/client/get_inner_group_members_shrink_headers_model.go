// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerGroupMembersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInnerGroupMembersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetInnerGroupMembersShrinkHeaders
	GetAccountContextShrink() *string
}

type GetInnerGroupMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetInnerGroupMembersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInnerGroupMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInnerGroupMembersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetInnerGroupMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetInnerGroupMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInnerGroupMembersShrinkHeaders) SetAccountContextShrink(v string) *GetInnerGroupMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetInnerGroupMembersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
