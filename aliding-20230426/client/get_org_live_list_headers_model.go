// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgLiveListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOrgLiveListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetOrgLiveListHeadersAccountContext) *GetOrgLiveListHeaders
	GetAccountContext() *GetOrgLiveListHeadersAccountContext
}

type GetOrgLiveListHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetOrgLiveListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetOrgLiveListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListHeaders) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOrgLiveListHeaders) GetAccountContext() *GetOrgLiveListHeadersAccountContext {
	return s.AccountContext
}

func (s *GetOrgLiveListHeaders) SetCommonHeaders(v map[string]*string) *GetOrgLiveListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrgLiveListHeaders) SetAccountContext(v *GetOrgLiveListHeadersAccountContext) *GetOrgLiveListHeaders {
	s.AccountContext = v
	return s
}

func (s *GetOrgLiveListHeaders) Validate() error {
	return dara.Validate(s)
}

type GetOrgLiveListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetOrgLiveListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetOrgLiveListHeadersAccountContext) SetAccountId(v string) *GetOrgLiveListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetOrgLiveListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
