// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAvatarHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateUserAvatarHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateUserAvatarHeadersAccountContext) *UpdateUserAvatarHeaders
	GetAccountContext() *UpdateUserAvatarHeadersAccountContext
}

type UpdateUserAvatarHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateUserAvatarHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateUserAvatarHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAvatarHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserAvatarHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateUserAvatarHeaders) GetAccountContext() *UpdateUserAvatarHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateUserAvatarHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserAvatarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserAvatarHeaders) SetAccountContext(v *UpdateUserAvatarHeadersAccountContext) *UpdateUserAvatarHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateUserAvatarHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateUserAvatarHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateUserAvatarHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAvatarHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateUserAvatarHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateUserAvatarHeadersAccountContext) SetAccountId(v string) *UpdateUserAvatarHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateUserAvatarHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
