// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserIsGroupMemberShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CheckUserIsGroupMemberShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CheckUserIsGroupMemberShrinkHeaders
	GetAccountContextShrink() *string
}

type CheckUserIsGroupMemberShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CheckUserIsGroupMemberShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CheckUserIsGroupMemberShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CheckUserIsGroupMemberShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CheckUserIsGroupMemberShrinkHeaders) SetCommonHeaders(v map[string]*string) *CheckUserIsGroupMemberShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckUserIsGroupMemberShrinkHeaders) SetAccountContextShrink(v string) *CheckUserIsGroupMemberShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CheckUserIsGroupMemberShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
