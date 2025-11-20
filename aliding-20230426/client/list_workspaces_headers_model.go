// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListWorkspacesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListWorkspacesHeadersAccountContext) *ListWorkspacesHeaders
	GetAccountContext() *ListWorkspacesHeadersAccountContext
}

type ListWorkspacesHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListWorkspacesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListWorkspacesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesHeaders) GoString() string {
	return s.String()
}

func (s *ListWorkspacesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListWorkspacesHeaders) GetAccountContext() *ListWorkspacesHeadersAccountContext {
	return s.AccountContext
}

func (s *ListWorkspacesHeaders) SetCommonHeaders(v map[string]*string) *ListWorkspacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListWorkspacesHeaders) SetAccountContext(v *ListWorkspacesHeadersAccountContext) *ListWorkspacesHeaders {
	s.AccountContext = v
	return s
}

func (s *ListWorkspacesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkspacesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListWorkspacesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListWorkspacesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListWorkspacesHeadersAccountContext) SetAccountId(v string) *ListWorkspacesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListWorkspacesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
