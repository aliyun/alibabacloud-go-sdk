// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskCopiesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetTaskCopiesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetTaskCopiesHeadersAccountContext) *GetTaskCopiesHeaders
	GetAccountContext() *GetTaskCopiesHeadersAccountContext
}

type GetTaskCopiesHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetTaskCopiesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetTaskCopiesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetTaskCopiesHeaders) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetTaskCopiesHeaders) GetAccountContext() *GetTaskCopiesHeadersAccountContext {
	return s.AccountContext
}

func (s *GetTaskCopiesHeaders) SetCommonHeaders(v map[string]*string) *GetTaskCopiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTaskCopiesHeaders) SetAccountContext(v *GetTaskCopiesHeadersAccountContext) *GetTaskCopiesHeaders {
	s.AccountContext = v
	return s
}

func (s *GetTaskCopiesHeaders) Validate() error {
	return dara.Validate(s)
}

type GetTaskCopiesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetTaskCopiesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetTaskCopiesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetTaskCopiesHeadersAccountContext) SetAccountId(v string) *GetTaskCopiesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetTaskCopiesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
