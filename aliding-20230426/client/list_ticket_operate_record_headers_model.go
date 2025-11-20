// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketOperateRecordHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListTicketOperateRecordHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListTicketOperateRecordHeadersAccountContext) *ListTicketOperateRecordHeaders
	GetAccountContext() *ListTicketOperateRecordHeadersAccountContext
}

type ListTicketOperateRecordHeaders struct {
	CommonHeaders  map[string]*string                            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListTicketOperateRecordHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListTicketOperateRecordHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordHeaders) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListTicketOperateRecordHeaders) GetAccountContext() *ListTicketOperateRecordHeadersAccountContext {
	return s.AccountContext
}

func (s *ListTicketOperateRecordHeaders) SetCommonHeaders(v map[string]*string) *ListTicketOperateRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTicketOperateRecordHeaders) SetAccountContext(v *ListTicketOperateRecordHeadersAccountContext) *ListTicketOperateRecordHeaders {
	s.AccountContext = v
	return s
}

func (s *ListTicketOperateRecordHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTicketOperateRecordHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListTicketOperateRecordHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListTicketOperateRecordHeadersAccountContext) SetAccountId(v string) *ListTicketOperateRecordHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListTicketOperateRecordHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
