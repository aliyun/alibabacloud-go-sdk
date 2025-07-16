// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupLiveListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetGroupLiveListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetGroupLiveListHeadersAccountContext) *GetGroupLiveListHeaders
	GetAccountContext() *GetGroupLiveListHeadersAccountContext
}

type GetGroupLiveListHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetGroupLiveListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetGroupLiveListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetGroupLiveListHeaders) GetAccountContext() *GetGroupLiveListHeadersAccountContext {
	return s.AccountContext
}

func (s *GetGroupLiveListHeaders) SetCommonHeaders(v map[string]*string) *GetGroupLiveListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupLiveListHeaders) SetAccountContext(v *GetGroupLiveListHeadersAccountContext) *GetGroupLiveListHeaders {
	s.AccountContext = v
	return s
}

func (s *GetGroupLiveListHeaders) Validate() error {
	return dara.Validate(s)
}

type GetGroupLiveListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetGroupLiveListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetGroupLiveListHeadersAccountContext) SetAccountId(v string) *GetGroupLiveListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetGroupLiveListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
