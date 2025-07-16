// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiverListReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ReceiverListReportHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ReceiverListReportHeadersAccountContext) *ReceiverListReportHeaders
	GetAccountContext() *ReceiverListReportHeadersAccountContext
}

type ReceiverListReportHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ReceiverListReportHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ReceiverListReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s ReceiverListReportHeaders) GoString() string {
	return s.String()
}

func (s *ReceiverListReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ReceiverListReportHeaders) GetAccountContext() *ReceiverListReportHeadersAccountContext {
	return s.AccountContext
}

func (s *ReceiverListReportHeaders) SetCommonHeaders(v map[string]*string) *ReceiverListReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReceiverListReportHeaders) SetAccountContext(v *ReceiverListReportHeadersAccountContext) *ReceiverListReportHeaders {
	s.AccountContext = v
	return s
}

func (s *ReceiverListReportHeaders) Validate() error {
	return dara.Validate(s)
}

type ReceiverListReportHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ReceiverListReportHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ReceiverListReportHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ReceiverListReportHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ReceiverListReportHeadersAccountContext) SetAccountId(v string) *ReceiverListReportHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ReceiverListReportHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
