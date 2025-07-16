// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenegroupMemberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteScenegroupMemberHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteScenegroupMemberHeadersAccountContext) *DeleteScenegroupMemberHeaders
	GetAccountContext() *DeleteScenegroupMemberHeadersAccountContext
}

type DeleteScenegroupMemberHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteScenegroupMemberHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteScenegroupMemberHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenegroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *DeleteScenegroupMemberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteScenegroupMemberHeaders) GetAccountContext() *DeleteScenegroupMemberHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteScenegroupMemberHeaders) SetCommonHeaders(v map[string]*string) *DeleteScenegroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteScenegroupMemberHeaders) SetAccountContext(v *DeleteScenegroupMemberHeadersAccountContext) *DeleteScenegroupMemberHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteScenegroupMemberHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteScenegroupMemberHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteScenegroupMemberHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenegroupMemberHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteScenegroupMemberHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteScenegroupMemberHeadersAccountContext) SetAccountId(v string) *DeleteScenegroupMemberHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteScenegroupMemberHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
