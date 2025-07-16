// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocUpdateContentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DocUpdateContentHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DocUpdateContentHeadersAccountContext) *DocUpdateContentHeaders
	GetAccountContext() *DocUpdateContentHeadersAccountContext
}

type DocUpdateContentHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DocUpdateContentHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DocUpdateContentHeaders) String() string {
	return dara.Prettify(s)
}

func (s DocUpdateContentHeaders) GoString() string {
	return s.String()
}

func (s *DocUpdateContentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DocUpdateContentHeaders) GetAccountContext() *DocUpdateContentHeadersAccountContext {
	return s.AccountContext
}

func (s *DocUpdateContentHeaders) SetCommonHeaders(v map[string]*string) *DocUpdateContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocUpdateContentHeaders) SetAccountContext(v *DocUpdateContentHeadersAccountContext) *DocUpdateContentHeaders {
	s.AccountContext = v
	return s
}

func (s *DocUpdateContentHeaders) Validate() error {
	return dara.Validate(s)
}

type DocUpdateContentHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DocUpdateContentHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DocUpdateContentHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DocUpdateContentHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DocUpdateContentHeadersAccountContext) SetAccountId(v string) *DocUpdateContentHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DocUpdateContentHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
