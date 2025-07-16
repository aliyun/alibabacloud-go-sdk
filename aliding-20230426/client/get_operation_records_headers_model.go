// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOperationRecordsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetOperationRecordsHeadersAccountContext) *GetOperationRecordsHeaders
	GetAccountContext() *GetOperationRecordsHeadersAccountContext
}

type GetOperationRecordsHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetOperationRecordsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetOperationRecordsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordsHeaders) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOperationRecordsHeaders) GetAccountContext() *GetOperationRecordsHeadersAccountContext {
	return s.AccountContext
}

func (s *GetOperationRecordsHeaders) SetCommonHeaders(v map[string]*string) *GetOperationRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOperationRecordsHeaders) SetAccountContext(v *GetOperationRecordsHeadersAccountContext) *GetOperationRecordsHeaders {
	s.AccountContext = v
	return s
}

func (s *GetOperationRecordsHeaders) Validate() error {
	return dara.Validate(s)
}

type GetOperationRecordsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetOperationRecordsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetOperationRecordsHeadersAccountContext) SetAccountId(v string) *GetOperationRecordsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetOperationRecordsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
