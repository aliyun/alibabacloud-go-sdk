// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMineWorkspaceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMineWorkspaceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetMineWorkspaceHeadersAccountContext) *GetMineWorkspaceHeaders
	GetAccountContext() *GetMineWorkspaceHeadersAccountContext
}

type GetMineWorkspaceHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetMineWorkspaceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetMineWorkspaceHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMineWorkspaceHeaders) GetAccountContext() *GetMineWorkspaceHeadersAccountContext {
	return s.AccountContext
}

func (s *GetMineWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *GetMineWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMineWorkspaceHeaders) SetAccountContext(v *GetMineWorkspaceHeadersAccountContext) *GetMineWorkspaceHeaders {
	s.AccountContext = v
	return s
}

func (s *GetMineWorkspaceHeaders) Validate() error {
	return dara.Validate(s)
}

type GetMineWorkspaceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetMineWorkspaceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMineWorkspaceHeadersAccountContext) SetAccountId(v string) *GetMineWorkspaceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetMineWorkspaceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
