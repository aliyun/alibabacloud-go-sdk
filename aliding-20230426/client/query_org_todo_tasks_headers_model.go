// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgTodoTasksHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryOrgTodoTasksHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryOrgTodoTasksHeadersAccountContext) *QueryOrgTodoTasksHeaders
	GetAccountContext() *QueryOrgTodoTasksHeadersAccountContext
}

type QueryOrgTodoTasksHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryOrgTodoTasksHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryOrgTodoTasksHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryOrgTodoTasksHeaders) GetAccountContext() *QueryOrgTodoTasksHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryOrgTodoTasksHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgTodoTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgTodoTasksHeaders) SetAccountContext(v *QueryOrgTodoTasksHeadersAccountContext) *QueryOrgTodoTasksHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryOrgTodoTasksHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryOrgTodoTasksHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryOrgTodoTasksHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryOrgTodoTasksHeadersAccountContext) SetAccountId(v string) *QueryOrgTodoTasksHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryOrgTodoTasksHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
