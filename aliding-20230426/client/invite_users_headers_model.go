// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteUsersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InviteUsersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InviteUsersHeadersAccountContext) *InviteUsersHeaders
	GetAccountContext() *InviteUsersHeadersAccountContext
}

type InviteUsersHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InviteUsersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InviteUsersHeaders) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersHeaders) GoString() string {
	return s.String()
}

func (s *InviteUsersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InviteUsersHeaders) GetAccountContext() *InviteUsersHeadersAccountContext {
	return s.AccountContext
}

func (s *InviteUsersHeaders) SetCommonHeaders(v map[string]*string) *InviteUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InviteUsersHeaders) SetAccountContext(v *InviteUsersHeadersAccountContext) *InviteUsersHeaders {
	s.AccountContext = v
	return s
}

func (s *InviteUsersHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InviteUsersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s InviteUsersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InviteUsersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InviteUsersHeadersAccountContext) SetAccountId(v string) *InviteUsersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InviteUsersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
