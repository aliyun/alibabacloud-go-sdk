// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddYikeUserCreditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredit(v int32) *AddYikeUserCreditRequest
	GetCredit() *int32
	SetYikeUserId(v string) *AddYikeUserCreditRequest
	GetYikeUserId() *string
}

type AddYikeUserCreditRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 50
	Credit *int32 `json:"Credit,omitempty" xml:"Credit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// id
	YikeUserId *string `json:"YikeUserId,omitempty" xml:"YikeUserId,omitempty"`
}

func (s AddYikeUserCreditRequest) String() string {
	return dara.Prettify(s)
}

func (s AddYikeUserCreditRequest) GoString() string {
	return s.String()
}

func (s *AddYikeUserCreditRequest) GetCredit() *int32 {
	return s.Credit
}

func (s *AddYikeUserCreditRequest) GetYikeUserId() *string {
	return s.YikeUserId
}

func (s *AddYikeUserCreditRequest) SetCredit(v int32) *AddYikeUserCreditRequest {
	s.Credit = &v
	return s
}

func (s *AddYikeUserCreditRequest) SetYikeUserId(v string) *AddYikeUserCreditRequest {
	s.YikeUserId = &v
	return s
}

func (s *AddYikeUserCreditRequest) Validate() error {
	return dara.Validate(s)
}
