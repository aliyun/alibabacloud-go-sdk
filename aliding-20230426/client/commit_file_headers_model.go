// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitFileHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CommitFileHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CommitFileHeadersAccountContext) *CommitFileHeaders
	GetAccountContext() *CommitFileHeadersAccountContext
}

type CommitFileHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CommitFileHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CommitFileHeaders) String() string {
	return dara.Prettify(s)
}

func (s CommitFileHeaders) GoString() string {
	return s.String()
}

func (s *CommitFileHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CommitFileHeaders) GetAccountContext() *CommitFileHeadersAccountContext {
	return s.AccountContext
}

func (s *CommitFileHeaders) SetCommonHeaders(v map[string]*string) *CommitFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommitFileHeaders) SetAccountContext(v *CommitFileHeadersAccountContext) *CommitFileHeaders {
	s.AccountContext = v
	return s
}

func (s *CommitFileHeaders) Validate() error {
	return dara.Validate(s)
}

type CommitFileHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CommitFileHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CommitFileHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CommitFileHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CommitFileHeadersAccountContext) SetAccountId(v string) *CommitFileHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CommitFileHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
