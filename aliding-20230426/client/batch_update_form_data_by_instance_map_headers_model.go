// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceMapHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceMapHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *BatchUpdateFormDataByInstanceMapHeadersAccountContext) *BatchUpdateFormDataByInstanceMapHeaders
	GetAccountContext() *BatchUpdateFormDataByInstanceMapHeadersAccountContext
}

type BatchUpdateFormDataByInstanceMapHeaders struct {
	CommonHeaders  map[string]*string                                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *BatchUpdateFormDataByInstanceMapHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s BatchUpdateFormDataByInstanceMapHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchUpdateFormDataByInstanceMapHeaders) GetAccountContext() *BatchUpdateFormDataByInstanceMapHeadersAccountContext {
	return s.AccountContext
}

func (s *BatchUpdateFormDataByInstanceMapHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceMapHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapHeaders) SetAccountContext(v *BatchUpdateFormDataByInstanceMapHeadersAccountContext) *BatchUpdateFormDataByInstanceMapHeaders {
	s.AccountContext = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapHeaders) Validate() error {
	return dara.Validate(s)
}

type BatchUpdateFormDataByInstanceMapHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s BatchUpdateFormDataByInstanceMapHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchUpdateFormDataByInstanceMapHeadersAccountContext) SetAccountId(v string) *BatchUpdateFormDataByInstanceMapHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
