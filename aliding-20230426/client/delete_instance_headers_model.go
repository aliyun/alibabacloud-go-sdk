// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteInstanceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteInstanceHeadersAccountContext) *DeleteInstanceHeaders
	GetAccountContext() *DeleteInstanceHeadersAccountContext
}

type DeleteInstanceHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteInstanceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteInstanceHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteInstanceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteInstanceHeaders) GetAccountContext() *DeleteInstanceHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteInstanceHeaders) SetCommonHeaders(v map[string]*string) *DeleteInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteInstanceHeaders) SetAccountContext(v *DeleteInstanceHeadersAccountContext) *DeleteInstanceHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteInstanceHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteInstanceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteInstanceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteInstanceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteInstanceHeadersAccountContext) SetAccountId(v string) *DeleteInstanceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteInstanceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
