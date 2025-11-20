// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserIsGroupMemberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CheckUserIsGroupMemberHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CheckUserIsGroupMemberHeadersAccountContext) *CheckUserIsGroupMemberHeaders
	GetAccountContext() *CheckUserIsGroupMemberHeadersAccountContext
}

type CheckUserIsGroupMemberHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CheckUserIsGroupMemberHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CheckUserIsGroupMemberHeaders) String() string {
	return dara.Prettify(s)
}

func (s CheckUserIsGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CheckUserIsGroupMemberHeaders) GetAccountContext() *CheckUserIsGroupMemberHeadersAccountContext {
	return s.AccountContext
}

func (s *CheckUserIsGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *CheckUserIsGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckUserIsGroupMemberHeaders) SetAccountContext(v *CheckUserIsGroupMemberHeadersAccountContext) *CheckUserIsGroupMemberHeaders {
	s.AccountContext = v
	return s
}

func (s *CheckUserIsGroupMemberHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckUserIsGroupMemberHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CheckUserIsGroupMemberHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CheckUserIsGroupMemberHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CheckUserIsGroupMemberHeadersAccountContext) SetAccountId(v string) *CheckUserIsGroupMemberHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CheckUserIsGroupMemberHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
