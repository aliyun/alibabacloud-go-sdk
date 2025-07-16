// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRelatedWorkspacesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetRelatedWorkspacesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetRelatedWorkspacesHeadersAccountContext) *GetRelatedWorkspacesHeaders
	GetAccountContext() *GetRelatedWorkspacesHeadersAccountContext
}

type GetRelatedWorkspacesHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetRelatedWorkspacesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetRelatedWorkspacesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesHeaders) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetRelatedWorkspacesHeaders) GetAccountContext() *GetRelatedWorkspacesHeadersAccountContext {
	return s.AccountContext
}

func (s *GetRelatedWorkspacesHeaders) SetCommonHeaders(v map[string]*string) *GetRelatedWorkspacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRelatedWorkspacesHeaders) SetAccountContext(v *GetRelatedWorkspacesHeadersAccountContext) *GetRelatedWorkspacesHeaders {
	s.AccountContext = v
	return s
}

func (s *GetRelatedWorkspacesHeaders) Validate() error {
	return dara.Validate(s)
}

type GetRelatedWorkspacesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetRelatedWorkspacesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetRelatedWorkspacesHeadersAccountContext) SetAccountId(v string) *GetRelatedWorkspacesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetRelatedWorkspacesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
