// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteColumnsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteColumnsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteColumnsHeadersAccountContext) *DeleteColumnsHeaders
	GetAccountContext() *DeleteColumnsHeadersAccountContext
}

type DeleteColumnsHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteColumnsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteColumnsHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteColumnsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteColumnsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteColumnsHeaders) GetAccountContext() *DeleteColumnsHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteColumnsHeaders) SetCommonHeaders(v map[string]*string) *DeleteColumnsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteColumnsHeaders) SetAccountContext(v *DeleteColumnsHeadersAccountContext) *DeleteColumnsHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteColumnsHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteColumnsHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteColumnsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteColumnsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteColumnsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteColumnsHeadersAccountContext) SetAccountId(v string) *DeleteColumnsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteColumnsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
