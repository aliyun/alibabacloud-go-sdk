// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MuteMembersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *MuteMembersHeadersAccountContext) *MuteMembersHeaders
	GetAccountContext() *MuteMembersHeadersAccountContext
}

type MuteMembersHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *MuteMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s MuteMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s MuteMembersHeaders) GoString() string {
	return s.String()
}

func (s *MuteMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MuteMembersHeaders) GetAccountContext() *MuteMembersHeadersAccountContext {
	return s.AccountContext
}

func (s *MuteMembersHeaders) SetCommonHeaders(v map[string]*string) *MuteMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MuteMembersHeaders) SetAccountContext(v *MuteMembersHeadersAccountContext) *MuteMembersHeaders {
	s.AccountContext = v
	return s
}

func (s *MuteMembersHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MuteMembersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s MuteMembersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s MuteMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *MuteMembersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *MuteMembersHeadersAccountContext) SetAccountId(v string) *MuteMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *MuteMembersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
