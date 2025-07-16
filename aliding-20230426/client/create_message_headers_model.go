// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateMessageHeaders
	GetCommonHeaders() map[string]*string
	SetAccountId(v string) *CreateMessageHeaders
	GetAccountId() *string
}

type CreateMessageHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateMessageHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageHeaders) GoString() string {
	return s.String()
}

func (s *CreateMessageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateMessageHeaders) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateMessageHeaders) SetCommonHeaders(v map[string]*string) *CreateMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMessageHeaders) SetAccountId(v string) *CreateMessageHeaders {
	s.AccountId = &v
	return s
}

func (s *CreateMessageHeaders) Validate() error {
	return dara.Validate(s)
}
