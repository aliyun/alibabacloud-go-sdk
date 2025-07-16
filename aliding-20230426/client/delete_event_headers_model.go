// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteEventHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteEventHeadersAccountContext) *DeleteEventHeaders
	GetAccountContext() *DeleteEventHeadersAccountContext
}

type DeleteEventHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteEventHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteEventHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventHeaders) GoString() string {
	return s.String()
}

func (s *DeleteEventHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteEventHeaders) GetAccountContext() *DeleteEventHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteEventHeaders) SetCommonHeaders(v map[string]*string) *DeleteEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteEventHeaders) SetAccountContext(v *DeleteEventHeadersAccountContext) *DeleteEventHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteEventHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteEventHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteEventHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteEventHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteEventHeadersAccountContext) SetAccountId(v string) *DeleteEventHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteEventHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
