// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByUrlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNodeByUrlHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetNodeByUrlHeadersAccountContext) *GetNodeByUrlHeaders
	GetAccountContext() *GetNodeByUrlHeadersAccountContext
}

type GetNodeByUrlHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetNodeByUrlHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetNodeByUrlHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNodeByUrlHeaders) GetAccountContext() *GetNodeByUrlHeadersAccountContext {
	return s.AccountContext
}

func (s *GetNodeByUrlHeaders) SetCommonHeaders(v map[string]*string) *GetNodeByUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNodeByUrlHeaders) SetAccountContext(v *GetNodeByUrlHeadersAccountContext) *GetNodeByUrlHeaders {
	s.AccountContext = v
	return s
}

func (s *GetNodeByUrlHeaders) Validate() error {
	return dara.Validate(s)
}

type GetNodeByUrlHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetNodeByUrlHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetNodeByUrlHeadersAccountContext) SetAccountId(v string) *GetNodeByUrlHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetNodeByUrlHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
