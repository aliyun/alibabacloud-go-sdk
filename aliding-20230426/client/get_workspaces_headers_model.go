// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspacesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetWorkspacesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetWorkspacesHeadersAccountContext) *GetWorkspacesHeaders
	GetAccountContext() *GetWorkspacesHeadersAccountContext
}

type GetWorkspacesHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetWorkspacesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetWorkspacesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesHeaders) GoString() string {
	return s.String()
}

func (s *GetWorkspacesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetWorkspacesHeaders) GetAccountContext() *GetWorkspacesHeadersAccountContext {
	return s.AccountContext
}

func (s *GetWorkspacesHeaders) SetCommonHeaders(v map[string]*string) *GetWorkspacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWorkspacesHeaders) SetAccountContext(v *GetWorkspacesHeadersAccountContext) *GetWorkspacesHeaders {
	s.AccountContext = v
	return s
}

func (s *GetWorkspacesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspacesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetWorkspacesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetWorkspacesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetWorkspacesHeadersAccountContext) SetAccountId(v string) *GetWorkspacesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetWorkspacesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
