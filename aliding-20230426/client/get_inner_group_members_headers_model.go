// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerGroupMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInnerGroupMembersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetInnerGroupMembersHeadersAccountContext) *GetInnerGroupMembersHeaders
	GetAccountContext() *GetInnerGroupMembersHeadersAccountContext
}

type GetInnerGroupMembersHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetInnerGroupMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetInnerGroupMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInnerGroupMembersHeaders) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInnerGroupMembersHeaders) GetAccountContext() *GetInnerGroupMembersHeadersAccountContext {
	return s.AccountContext
}

func (s *GetInnerGroupMembersHeaders) SetCommonHeaders(v map[string]*string) *GetInnerGroupMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInnerGroupMembersHeaders) SetAccountContext(v *GetInnerGroupMembersHeadersAccountContext) *GetInnerGroupMembersHeaders {
	s.AccountContext = v
	return s
}

func (s *GetInnerGroupMembersHeaders) Validate() error {
	return dara.Validate(s)
}

type GetInnerGroupMembersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetInnerGroupMembersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetInnerGroupMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetInnerGroupMembersHeadersAccountContext) SetAccountId(v string) *GetInnerGroupMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetInnerGroupMembersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
