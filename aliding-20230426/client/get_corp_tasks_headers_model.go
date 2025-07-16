// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpTasksHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCorpTasksHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetCorpTasksHeadersAccountContext) *GetCorpTasksHeaders
	GetAccountContext() *GetCorpTasksHeadersAccountContext
}

type GetCorpTasksHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetCorpTasksHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetCorpTasksHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCorpTasksHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpTasksHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCorpTasksHeaders) GetAccountContext() *GetCorpTasksHeadersAccountContext {
	return s.AccountContext
}

func (s *GetCorpTasksHeaders) SetCommonHeaders(v map[string]*string) *GetCorpTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpTasksHeaders) SetAccountContext(v *GetCorpTasksHeadersAccountContext) *GetCorpTasksHeaders {
	s.AccountContext = v
	return s
}

func (s *GetCorpTasksHeaders) Validate() error {
	return dara.Validate(s)
}

type GetCorpTasksHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetCorpTasksHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetCorpTasksHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetCorpTasksHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetCorpTasksHeadersAccountContext) SetAccountId(v string) *GetCorpTasksHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetCorpTasksHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
