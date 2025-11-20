// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableFieldHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMultiDimTableFieldHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteMultiDimTableFieldHeadersAccountContext) *DeleteMultiDimTableFieldHeaders
	GetAccountContext() *DeleteMultiDimTableFieldHeadersAccountContext
}

type DeleteMultiDimTableFieldHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteMultiDimTableFieldHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteMultiDimTableFieldHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableFieldHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableFieldHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMultiDimTableFieldHeaders) GetAccountContext() *DeleteMultiDimTableFieldHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteMultiDimTableFieldHeaders) SetCommonHeaders(v map[string]*string) *DeleteMultiDimTableFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMultiDimTableFieldHeaders) SetAccountContext(v *DeleteMultiDimTableFieldHeadersAccountContext) *DeleteMultiDimTableFieldHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteMultiDimTableFieldHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMultiDimTableFieldHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteMultiDimTableFieldHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableFieldHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableFieldHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteMultiDimTableFieldHeadersAccountContext) SetAccountId(v string) *DeleteMultiDimTableFieldHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteMultiDimTableFieldHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
