// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNodesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetNodesHeadersAccountContext) *GetNodesHeaders
	GetAccountContext() *GetNodesHeadersAccountContext
}

type GetNodesHeaders struct {
	CommonHeaders  map[string]*string             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetNodesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetNodesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNodesHeaders) GoString() string {
	return s.String()
}

func (s *GetNodesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNodesHeaders) GetAccountContext() *GetNodesHeadersAccountContext {
	return s.AccountContext
}

func (s *GetNodesHeaders) SetCommonHeaders(v map[string]*string) *GetNodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNodesHeaders) SetAccountContext(v *GetNodesHeadersAccountContext) *GetNodesHeaders {
	s.AccountContext = v
	return s
}

func (s *GetNodesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNodesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetNodesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetNodesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetNodesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetNodesHeadersAccountContext) SetAccountId(v string) *GetNodesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetNodesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
