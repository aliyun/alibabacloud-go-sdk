// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddWorkspaceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddWorkspaceHeadersAccountContext) *AddWorkspaceHeaders
	GetAccountContext() *AddWorkspaceHeadersAccountContext
}

type AddWorkspaceHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddWorkspaceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddWorkspaceHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddWorkspaceHeaders) GetAccountContext() *AddWorkspaceHeadersAccountContext {
	return s.AccountContext
}

func (s *AddWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceHeaders) SetAccountContext(v *AddWorkspaceHeadersAccountContext) *AddWorkspaceHeaders {
	s.AccountContext = v
	return s
}

func (s *AddWorkspaceHeaders) Validate() error {
	return dara.Validate(s)
}

type AddWorkspaceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddWorkspaceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddWorkspaceHeadersAccountContext) SetAccountId(v string) *AddWorkspaceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddWorkspaceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
