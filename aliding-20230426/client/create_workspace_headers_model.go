// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateWorkspaceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateWorkspaceHeadersAccountContext) *CreateWorkspaceHeaders
	GetAccountContext() *CreateWorkspaceHeadersAccountContext
}

type CreateWorkspaceHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateWorkspaceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateWorkspaceHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateWorkspaceHeaders) GetAccountContext() *CreateWorkspaceHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceHeaders) SetAccountContext(v *CreateWorkspaceHeadersAccountContext) *CreateWorkspaceHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateWorkspaceHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateWorkspaceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateWorkspaceHeadersAccountContext) SetAccountId(v string) *CreateWorkspaceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateWorkspaceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
