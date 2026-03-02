// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeveloper interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *Developer
	GetAccountId() *string
	SetCodeupAccountId(v string) *Developer
	GetCodeupAccountId() *string
	SetEmail(v string) *Developer
	GetEmail() *string
	SetEnterpriseId(v int64) *Developer
	GetEnterpriseId() *int64
	SetName(v string) *Developer
	GetName() *string
	SetPhone(v string) *Developer
	GetPhone() *string
	SetRequestId(v string) *Developer
	GetRequestId() *string
}

type Developer struct {
	// This parameter is required.
	//
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 121321412341
	CodeupAccountId *string `json:"codeupAccountId,omitempty" xml:"codeupAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx@alibaba.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 尚仁
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15113456789
	Phone     *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s Developer) String() string {
	return dara.Prettify(s)
}

func (s Developer) GoString() string {
	return s.String()
}

func (s *Developer) GetAccountId() *string {
	return s.AccountId
}

func (s *Developer) GetCodeupAccountId() *string {
	return s.CodeupAccountId
}

func (s *Developer) GetEmail() *string {
	return s.Email
}

func (s *Developer) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *Developer) GetName() *string {
	return s.Name
}

func (s *Developer) GetPhone() *string {
	return s.Phone
}

func (s *Developer) GetRequestId() *string {
	return s.RequestId
}

func (s *Developer) SetAccountId(v string) *Developer {
	s.AccountId = &v
	return s
}

func (s *Developer) SetCodeupAccountId(v string) *Developer {
	s.CodeupAccountId = &v
	return s
}

func (s *Developer) SetEmail(v string) *Developer {
	s.Email = &v
	return s
}

func (s *Developer) SetEnterpriseId(v int64) *Developer {
	s.EnterpriseId = &v
	return s
}

func (s *Developer) SetName(v string) *Developer {
	s.Name = &v
	return s
}

func (s *Developer) SetPhone(v string) *Developer {
	s.Phone = &v
	return s
}

func (s *Developer) SetRequestId(v string) *Developer {
	s.RequestId = &v
	return s
}

func (s *Developer) Validate() error {
	return dara.Validate(s)
}
