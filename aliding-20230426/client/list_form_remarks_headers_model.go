// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFormRemarksHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListFormRemarksHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListFormRemarksHeadersAccountContext) *ListFormRemarksHeaders
	GetAccountContext() *ListFormRemarksHeadersAccountContext
}

type ListFormRemarksHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListFormRemarksHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListFormRemarksHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListFormRemarksHeaders) GoString() string {
	return s.String()
}

func (s *ListFormRemarksHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListFormRemarksHeaders) GetAccountContext() *ListFormRemarksHeadersAccountContext {
	return s.AccountContext
}

func (s *ListFormRemarksHeaders) SetCommonHeaders(v map[string]*string) *ListFormRemarksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFormRemarksHeaders) SetAccountContext(v *ListFormRemarksHeadersAccountContext) *ListFormRemarksHeaders {
	s.AccountContext = v
	return s
}

func (s *ListFormRemarksHeaders) Validate() error {
	return dara.Validate(s)
}

type ListFormRemarksHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListFormRemarksHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListFormRemarksHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListFormRemarksHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListFormRemarksHeadersAccountContext) SetAccountId(v string) *ListFormRemarksHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListFormRemarksHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
