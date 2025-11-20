// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRangeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateRangeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateRangeHeadersAccountContext) *UpdateRangeHeaders
	GetAccountContext() *UpdateRangeHeadersAccountContext
}

type UpdateRangeHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateRangeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateRangeHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateRangeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRangeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateRangeHeaders) GetAccountContext() *UpdateRangeHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateRangeHeaders) SetCommonHeaders(v map[string]*string) *UpdateRangeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRangeHeaders) SetAccountContext(v *UpdateRangeHeadersAccountContext) *UpdateRangeHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateRangeHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRangeHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateRangeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateRangeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateRangeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateRangeHeadersAccountContext) SetAccountId(v string) *UpdateRangeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateRangeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
