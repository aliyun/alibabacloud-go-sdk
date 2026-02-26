// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMemberAccountResult interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v []*string) *MemberAccountResult
	GetAccountNo() []*string
	SetShopId(v string) *MemberAccountResult
	GetShopId() *string
}

type MemberAccountResult struct {
	// example:
	//
	// "yue***@newburn.cn"
	AccountNo []*string `json:"accountNo,omitempty" xml:"accountNo,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
}

func (s MemberAccountResult) String() string {
	return dara.Prettify(s)
}

func (s MemberAccountResult) GoString() string {
	return s.String()
}

func (s *MemberAccountResult) GetAccountNo() []*string {
	return s.AccountNo
}

func (s *MemberAccountResult) GetShopId() *string {
	return s.ShopId
}

func (s *MemberAccountResult) SetAccountNo(v []*string) *MemberAccountResult {
	s.AccountNo = v
	return s
}

func (s *MemberAccountResult) SetShopId(v string) *MemberAccountResult {
	s.ShopId = &v
	return s
}

func (s *MemberAccountResult) Validate() error {
	return dara.Validate(s)
}
