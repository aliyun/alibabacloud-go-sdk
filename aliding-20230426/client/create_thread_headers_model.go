// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThreadHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateThreadHeaders
	GetCommonHeaders() map[string]*string
	SetAccountId(v string) *CreateThreadHeaders
	GetAccountId() *string
}

type CreateThreadHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateThreadHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadHeaders) GoString() string {
	return s.String()
}

func (s *CreateThreadHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateThreadHeaders) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateThreadHeaders) SetCommonHeaders(v map[string]*string) *CreateThreadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateThreadHeaders) SetAccountId(v string) *CreateThreadHeaders {
	s.AccountId = &v
	return s
}

func (s *CreateThreadHeaders) Validate() error {
	return dara.Validate(s)
}
