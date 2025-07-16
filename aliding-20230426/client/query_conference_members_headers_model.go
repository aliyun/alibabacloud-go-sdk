// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryConferenceMembersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryConferenceMembersHeadersAccountContext) *QueryConferenceMembersHeaders
	GetAccountContext() *QueryConferenceMembersHeadersAccountContext
}

type QueryConferenceMembersHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryConferenceMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryConferenceMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceMembersHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryConferenceMembersHeaders) GetAccountContext() *QueryConferenceMembersHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryConferenceMembersHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceMembersHeaders) SetAccountContext(v *QueryConferenceMembersHeadersAccountContext) *QueryConferenceMembersHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryConferenceMembersHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryConferenceMembersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryConferenceMembersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryConferenceMembersHeadersAccountContext) SetAccountId(v string) *QueryConferenceMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryConferenceMembersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
