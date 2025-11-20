// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksModifyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DocBlocksModifyHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DocBlocksModifyHeadersAccountContext) *DocBlocksModifyHeaders
	GetAccountContext() *DocBlocksModifyHeadersAccountContext
}

type DocBlocksModifyHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DocBlocksModifyHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DocBlocksModifyHeaders) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksModifyHeaders) GoString() string {
	return s.String()
}

func (s *DocBlocksModifyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DocBlocksModifyHeaders) GetAccountContext() *DocBlocksModifyHeadersAccountContext {
	return s.AccountContext
}

func (s *DocBlocksModifyHeaders) SetCommonHeaders(v map[string]*string) *DocBlocksModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocBlocksModifyHeaders) SetAccountContext(v *DocBlocksModifyHeadersAccountContext) *DocBlocksModifyHeaders {
	s.AccountContext = v
	return s
}

func (s *DocBlocksModifyHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocBlocksModifyHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DocBlocksModifyHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksModifyHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DocBlocksModifyHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DocBlocksModifyHeadersAccountContext) SetAccountId(v string) *DocBlocksModifyHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DocBlocksModifyHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
