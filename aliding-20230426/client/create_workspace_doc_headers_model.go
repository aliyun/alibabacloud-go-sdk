// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceDocHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateWorkspaceDocHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateWorkspaceDocHeadersAccountContext) *CreateWorkspaceDocHeaders
	GetAccountContext() *CreateWorkspaceDocHeadersAccountContext
}

type CreateWorkspaceDocHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateWorkspaceDocHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateWorkspaceDocHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceDocHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateWorkspaceDocHeaders) GetAccountContext() *CreateWorkspaceDocHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateWorkspaceDocHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceDocHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceDocHeaders) SetAccountContext(v *CreateWorkspaceDocHeadersAccountContext) *CreateWorkspaceDocHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateWorkspaceDocHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceDocHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateWorkspaceDocHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceDocHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateWorkspaceDocHeadersAccountContext) SetAccountId(v string) *CreateWorkspaceDocHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateWorkspaceDocHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
