// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNodeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetNodeHeadersAccountContext) *GetNodeHeaders
	GetAccountContext() *GetNodeHeadersAccountContext
}

type GetNodeHeaders struct {
	CommonHeaders  map[string]*string            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetNodeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetNodeHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNodeHeaders) GoString() string {
	return s.String()
}

func (s *GetNodeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNodeHeaders) GetAccountContext() *GetNodeHeadersAccountContext {
	return s.AccountContext
}

func (s *GetNodeHeaders) SetCommonHeaders(v map[string]*string) *GetNodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNodeHeaders) SetAccountContext(v *GetNodeHeadersAccountContext) *GetNodeHeaders {
	s.AccountContext = v
	return s
}

func (s *GetNodeHeaders) Validate() error {
	return dara.Validate(s)
}

type GetNodeHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetNodeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetNodeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetNodeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetNodeHeadersAccountContext) SetAccountId(v string) *GetNodeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetNodeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
