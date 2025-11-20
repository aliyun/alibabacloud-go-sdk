// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *BatchUpdateFormDataByInstanceIdHeadersAccountContext) *BatchUpdateFormDataByInstanceIdHeaders
	GetAccountContext() *BatchUpdateFormDataByInstanceIdHeadersAccountContext
}

type BatchUpdateFormDataByInstanceIdHeaders struct {
	CommonHeaders  map[string]*string                                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *BatchUpdateFormDataByInstanceIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s BatchUpdateFormDataByInstanceIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchUpdateFormDataByInstanceIdHeaders) GetAccountContext() *BatchUpdateFormDataByInstanceIdHeadersAccountContext {
	return s.AccountContext
}

func (s *BatchUpdateFormDataByInstanceIdHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdHeaders) SetAccountContext(v *BatchUpdateFormDataByInstanceIdHeadersAccountContext) *BatchUpdateFormDataByInstanceIdHeaders {
	s.AccountContext = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchUpdateFormDataByInstanceIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s BatchUpdateFormDataByInstanceIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchUpdateFormDataByInstanceIdHeadersAccountContext) SetAccountId(v string) *BatchUpdateFormDataByInstanceIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
