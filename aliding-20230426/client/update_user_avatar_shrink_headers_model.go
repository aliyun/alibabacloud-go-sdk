// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAvatarShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateUserAvatarShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateUserAvatarShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateUserAvatarShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateUserAvatarShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAvatarShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserAvatarShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateUserAvatarShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateUserAvatarShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserAvatarShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserAvatarShrinkHeaders) SetAccountContextShrink(v string) *UpdateUserAvatarShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateUserAvatarShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
