// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveRunHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RetrieveRunHeaders
	GetCommonHeaders() map[string]*string
	SetAccountId(v string) *RetrieveRunHeaders
	GetAccountId() *string
}

type RetrieveRunHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s RetrieveRunHeaders) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRunHeaders) GoString() string {
	return s.String()
}

func (s *RetrieveRunHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RetrieveRunHeaders) GetAccountId() *string {
	return s.AccountId
}

func (s *RetrieveRunHeaders) SetCommonHeaders(v map[string]*string) *RetrieveRunHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RetrieveRunHeaders) SetAccountId(v string) *RetrieveRunHeaders {
	s.AccountId = &v
	return s
}

func (s *RetrieveRunHeaders) Validate() error {
	return dara.Validate(s)
}
