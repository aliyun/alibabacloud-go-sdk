// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListNodesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListNodesHeadersAccountContext) *ListNodesHeaders
	GetAccountContext() *ListNodesHeadersAccountContext
}

type ListNodesHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListNodesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListNodesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListNodesHeaders) GoString() string {
	return s.String()
}

func (s *ListNodesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListNodesHeaders) GetAccountContext() *ListNodesHeadersAccountContext {
	return s.AccountContext
}

func (s *ListNodesHeaders) SetCommonHeaders(v map[string]*string) *ListNodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListNodesHeaders) SetAccountContext(v *ListNodesHeadersAccountContext) *ListNodesHeaders {
	s.AccountContext = v
	return s
}

func (s *ListNodesHeaders) Validate() error {
	return dara.Validate(s)
}

type ListNodesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListNodesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListNodesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListNodesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListNodesHeadersAccountContext) SetAccountId(v string) *ListNodesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListNodesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
