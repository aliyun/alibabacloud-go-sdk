// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAliPayUrlResult interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateAliPayUrlResult
	GetAccountId() *string
	SetMemberId(v string) *CreateAliPayUrlResult
	GetMemberId() *string
	SetZftWithholdSignUrl(v string) *CreateAliPayUrlResult
	GetZftWithholdSignUrl() *string
}

type CreateAliPayUrlResult struct {
	AccountId          *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	MemberId           *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	ZftWithholdSignUrl *string `json:"zftWithholdSignUrl,omitempty" xml:"zftWithholdSignUrl,omitempty"`
}

func (s CreateAliPayUrlResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAliPayUrlResult) GoString() string {
	return s.String()
}

func (s *CreateAliPayUrlResult) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateAliPayUrlResult) GetMemberId() *string {
	return s.MemberId
}

func (s *CreateAliPayUrlResult) GetZftWithholdSignUrl() *string {
	return s.ZftWithholdSignUrl
}

func (s *CreateAliPayUrlResult) SetAccountId(v string) *CreateAliPayUrlResult {
	s.AccountId = &v
	return s
}

func (s *CreateAliPayUrlResult) SetMemberId(v string) *CreateAliPayUrlResult {
	s.MemberId = &v
	return s
}

func (s *CreateAliPayUrlResult) SetZftWithholdSignUrl(v string) *CreateAliPayUrlResult {
	s.ZftWithholdSignUrl = &v
	return s
}

func (s *CreateAliPayUrlResult) Validate() error {
	return dara.Validate(s)
}
