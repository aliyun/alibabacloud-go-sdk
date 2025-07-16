// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListMessageHeaders
	GetCommonHeaders() map[string]*string
	SetAccountId(v string) *ListMessageHeaders
	GetAccountId() *string
}

type ListMessageHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListMessageHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListMessageHeaders) GoString() string {
	return s.String()
}

func (s *ListMessageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListMessageHeaders) GetAccountId() *string {
	return s.AccountId
}

func (s *ListMessageHeaders) SetCommonHeaders(v map[string]*string) *ListMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMessageHeaders) SetAccountId(v string) *ListMessageHeaders {
	s.AccountId = &v
	return s
}

func (s *ListMessageHeaders) Validate() error {
	return dara.Validate(s)
}
