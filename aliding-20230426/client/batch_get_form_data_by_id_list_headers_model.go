// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFormDataByIdListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchGetFormDataByIdListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *BatchGetFormDataByIdListHeadersAccountContext) *BatchGetFormDataByIdListHeaders
	GetAccountContext() *BatchGetFormDataByIdListHeadersAccountContext
}

type BatchGetFormDataByIdListHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *BatchGetFormDataByIdListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s BatchGetFormDataByIdListHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchGetFormDataByIdListHeaders) GetAccountContext() *BatchGetFormDataByIdListHeadersAccountContext {
	return s.AccountContext
}

func (s *BatchGetFormDataByIdListHeaders) SetCommonHeaders(v map[string]*string) *BatchGetFormDataByIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetFormDataByIdListHeaders) SetAccountContext(v *BatchGetFormDataByIdListHeadersAccountContext) *BatchGetFormDataByIdListHeaders {
	s.AccountContext = v
	return s
}

func (s *BatchGetFormDataByIdListHeaders) Validate() error {
	return dara.Validate(s)
}

type BatchGetFormDataByIdListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s BatchGetFormDataByIdListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchGetFormDataByIdListHeadersAccountContext) SetAccountId(v string) *BatchGetFormDataByIdListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *BatchGetFormDataByIdListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
