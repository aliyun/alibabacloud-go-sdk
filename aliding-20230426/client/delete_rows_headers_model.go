// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteRowsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteRowsHeadersAccountContext) *DeleteRowsHeaders
	GetAccountContext() *DeleteRowsHeadersAccountContext
}

type DeleteRowsHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteRowsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteRowsHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRowsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteRowsHeaders) GetAccountContext() *DeleteRowsHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteRowsHeaders) SetCommonHeaders(v map[string]*string) *DeleteRowsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRowsHeaders) SetAccountContext(v *DeleteRowsHeadersAccountContext) *DeleteRowsHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteRowsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRowsHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteRowsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteRowsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteRowsHeadersAccountContext) SetAccountId(v string) *DeleteRowsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteRowsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
