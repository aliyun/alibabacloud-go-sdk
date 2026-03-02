// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetTokenRequest
	GetAccountId() *string
	SetPbcId(v int64) *GetTokenRequest
	GetPbcId() *int64
}

type GetTokenRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	PbcId     *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *GetTokenRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *GetTokenRequest) SetAccountId(v string) *GetTokenRequest {
	s.AccountId = &v
	return s
}

func (s *GetTokenRequest) SetPbcId(v int64) *GetTokenRequest {
	s.PbcId = &v
	return s
}

func (s *GetTokenRequest) Validate() error {
	return dara.Validate(s)
}
