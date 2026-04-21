// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubYikeUserCreditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredit(v int32) *SubYikeUserCreditRequest
	GetCredit() *int32
	SetYikeUserId(v string) *SubYikeUserCreditRequest
	GetYikeUserId() *string
}

type SubYikeUserCreditRequest struct {
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

func (s SubYikeUserCreditRequest) String() string {
	return dara.Prettify(s)
}

func (s SubYikeUserCreditRequest) GoString() string {
	return s.String()
}

func (s *SubYikeUserCreditRequest) GetCredit() *int32 {
	return s.Credit
}

func (s *SubYikeUserCreditRequest) GetYikeUserId() *string {
	return s.YikeUserId
}

func (s *SubYikeUserCreditRequest) SetCredit(v int32) *SubYikeUserCreditRequest {
	s.Credit = &v
	return s
}

func (s *SubYikeUserCreditRequest) SetYikeUserId(v string) *SubYikeUserCreditRequest {
	s.YikeUserId = &v
	return s
}

func (s *SubYikeUserCreditRequest) Validate() error {
	return dara.Validate(s)
}
