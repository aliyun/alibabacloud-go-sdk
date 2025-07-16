// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRunHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateRunHeaders
	GetCommonHeaders() map[string]*string
	SetAccountId(v string) *CreateRunHeaders
	GetAccountId() *string
}

type CreateRunHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateRunHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateRunHeaders) GoString() string {
	return s.String()
}

func (s *CreateRunHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateRunHeaders) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateRunHeaders) SetCommonHeaders(v map[string]*string) *CreateRunHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRunHeaders) SetAccountId(v string) *CreateRunHeaders {
	s.AccountId = &v
	return s
}

func (s *CreateRunHeaders) Validate() error {
	return dara.Validate(s)
}
