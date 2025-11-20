// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableRecordsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMultiDimTableRecordsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteMultiDimTableRecordsHeadersAccountContext) *DeleteMultiDimTableRecordsHeaders
	GetAccountContext() *DeleteMultiDimTableRecordsHeadersAccountContext
}

type DeleteMultiDimTableRecordsHeaders struct {
	CommonHeaders  map[string]*string                               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteMultiDimTableRecordsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteMultiDimTableRecordsHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableRecordsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableRecordsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMultiDimTableRecordsHeaders) GetAccountContext() *DeleteMultiDimTableRecordsHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteMultiDimTableRecordsHeaders) SetCommonHeaders(v map[string]*string) *DeleteMultiDimTableRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMultiDimTableRecordsHeaders) SetAccountContext(v *DeleteMultiDimTableRecordsHeadersAccountContext) *DeleteMultiDimTableRecordsHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteMultiDimTableRecordsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMultiDimTableRecordsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteMultiDimTableRecordsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableRecordsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableRecordsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteMultiDimTableRecordsHeadersAccountContext) SetAccountId(v string) *DeleteMultiDimTableRecordsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteMultiDimTableRecordsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
