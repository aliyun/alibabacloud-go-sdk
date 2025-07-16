// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSaveFormDataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchSaveFormDataHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *BatchSaveFormDataHeadersAccountContext) *BatchSaveFormDataHeaders
	GetAccountContext() *BatchSaveFormDataHeadersAccountContext
}

type BatchSaveFormDataHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *BatchSaveFormDataHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s BatchSaveFormDataHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveFormDataHeaders) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchSaveFormDataHeaders) GetAccountContext() *BatchSaveFormDataHeadersAccountContext {
	return s.AccountContext
}

func (s *BatchSaveFormDataHeaders) SetCommonHeaders(v map[string]*string) *BatchSaveFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSaveFormDataHeaders) SetAccountContext(v *BatchSaveFormDataHeadersAccountContext) *BatchSaveFormDataHeaders {
	s.AccountContext = v
	return s
}

func (s *BatchSaveFormDataHeaders) Validate() error {
	return dara.Validate(s)
}

type BatchSaveFormDataHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s BatchSaveFormDataHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveFormDataHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchSaveFormDataHeadersAccountContext) SetAccountId(v string) *BatchSaveFormDataHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *BatchSaveFormDataHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
