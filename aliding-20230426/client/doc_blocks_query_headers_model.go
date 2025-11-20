// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DocBlocksQueryHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DocBlocksQueryHeadersAccountContext) *DocBlocksQueryHeaders
	GetAccountContext() *DocBlocksQueryHeadersAccountContext
}

type DocBlocksQueryHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DocBlocksQueryHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DocBlocksQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksQueryHeaders) GoString() string {
	return s.String()
}

func (s *DocBlocksQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DocBlocksQueryHeaders) GetAccountContext() *DocBlocksQueryHeadersAccountContext {
	return s.AccountContext
}

func (s *DocBlocksQueryHeaders) SetCommonHeaders(v map[string]*string) *DocBlocksQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocBlocksQueryHeaders) SetAccountContext(v *DocBlocksQueryHeadersAccountContext) *DocBlocksQueryHeaders {
	s.AccountContext = v
	return s
}

func (s *DocBlocksQueryHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocBlocksQueryHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DocBlocksQueryHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksQueryHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DocBlocksQueryHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DocBlocksQueryHeadersAccountContext) SetAccountId(v string) *DocBlocksQueryHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DocBlocksQueryHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
