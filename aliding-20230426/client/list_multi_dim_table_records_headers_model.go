// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiDimTableRecordsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListMultiDimTableRecordsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListMultiDimTableRecordsHeadersAccountContext) *ListMultiDimTableRecordsHeaders
	GetAccountContext() *ListMultiDimTableRecordsHeadersAccountContext
}

type ListMultiDimTableRecordsHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListMultiDimTableRecordsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListMultiDimTableRecordsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsHeaders) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListMultiDimTableRecordsHeaders) GetAccountContext() *ListMultiDimTableRecordsHeadersAccountContext {
	return s.AccountContext
}

func (s *ListMultiDimTableRecordsHeaders) SetCommonHeaders(v map[string]*string) *ListMultiDimTableRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMultiDimTableRecordsHeaders) SetAccountContext(v *ListMultiDimTableRecordsHeadersAccountContext) *ListMultiDimTableRecordsHeaders {
	s.AccountContext = v
	return s
}

func (s *ListMultiDimTableRecordsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMultiDimTableRecordsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListMultiDimTableRecordsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListMultiDimTableRecordsHeadersAccountContext) SetAccountId(v string) *ListMultiDimTableRecordsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListMultiDimTableRecordsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
