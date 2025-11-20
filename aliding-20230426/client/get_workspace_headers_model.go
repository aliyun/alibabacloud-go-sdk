// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetWorkspaceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetWorkspaceHeadersAccountContext) *GetWorkspaceHeaders
	GetAccountContext() *GetWorkspaceHeadersAccountContext
}

type GetWorkspaceHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetWorkspaceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetWorkspaceHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *GetWorkspaceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetWorkspaceHeaders) GetAccountContext() *GetWorkspaceHeadersAccountContext {
	return s.AccountContext
}

func (s *GetWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *GetWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWorkspaceHeaders) SetAccountContext(v *GetWorkspaceHeadersAccountContext) *GetWorkspaceHeaders {
	s.AccountContext = v
	return s
}

func (s *GetWorkspaceHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetWorkspaceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetWorkspaceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetWorkspaceHeadersAccountContext) SetAccountId(v string) *GetWorkspaceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetWorkspaceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
