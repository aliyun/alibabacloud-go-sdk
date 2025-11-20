// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksDeleteHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DocBlocksDeleteHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DocBlocksDeleteHeadersAccountContext) *DocBlocksDeleteHeaders
	GetAccountContext() *DocBlocksDeleteHeadersAccountContext
}

type DocBlocksDeleteHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DocBlocksDeleteHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DocBlocksDeleteHeaders) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksDeleteHeaders) GoString() string {
	return s.String()
}

func (s *DocBlocksDeleteHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DocBlocksDeleteHeaders) GetAccountContext() *DocBlocksDeleteHeadersAccountContext {
	return s.AccountContext
}

func (s *DocBlocksDeleteHeaders) SetCommonHeaders(v map[string]*string) *DocBlocksDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocBlocksDeleteHeaders) SetAccountContext(v *DocBlocksDeleteHeadersAccountContext) *DocBlocksDeleteHeaders {
	s.AccountContext = v
	return s
}

func (s *DocBlocksDeleteHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocBlocksDeleteHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DocBlocksDeleteHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksDeleteHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DocBlocksDeleteHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DocBlocksDeleteHeadersAccountContext) SetAccountId(v string) *DocBlocksDeleteHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DocBlocksDeleteHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
