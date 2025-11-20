// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddWorkspaceMembersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddWorkspaceMembersHeadersAccountContext) *AddWorkspaceMembersHeaders
	GetAccountContext() *AddWorkspaceMembersHeadersAccountContext
}

type AddWorkspaceMembersHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddWorkspaceMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddWorkspaceMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceMembersHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddWorkspaceMembersHeaders) GetAccountContext() *AddWorkspaceMembersHeadersAccountContext {
	return s.AccountContext
}

func (s *AddWorkspaceMembersHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceMembersHeaders) SetAccountContext(v *AddWorkspaceMembersHeadersAccountContext) *AddWorkspaceMembersHeaders {
	s.AccountContext = v
	return s
}

func (s *AddWorkspaceMembersHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddWorkspaceMembersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddWorkspaceMembersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddWorkspaceMembersHeadersAccountContext) SetAccountId(v string) *AddWorkspaceMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddWorkspaceMembersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
