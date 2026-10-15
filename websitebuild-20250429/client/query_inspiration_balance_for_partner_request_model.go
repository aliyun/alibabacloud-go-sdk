// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationBalanceForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBelongId(v string) *QueryInspirationBalanceForPartnerRequest
	GetBelongId() *string
	SetBelongIdType(v string) *QueryInspirationBalanceForPartnerRequest
	GetBelongIdType() *string
}

type QueryInspirationBalanceForPartnerRequest struct {
	// The owner ID, which can be a website ID or an Alibaba Cloud account ID.
	//
	// example:
	//
	// 123456
	BelongId *string `json:"BelongId,omitempty" xml:"BelongId,omitempty"`
	// The type of the owner ID. Valid values: siteId and uid.
	//
	// example:
	//
	// USER
	BelongIdType *string `json:"BelongIdType,omitempty" xml:"BelongIdType,omitempty"`
}

func (s QueryInspirationBalanceForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationBalanceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *QueryInspirationBalanceForPartnerRequest) GetBelongId() *string {
	return s.BelongId
}

func (s *QueryInspirationBalanceForPartnerRequest) GetBelongIdType() *string {
	return s.BelongIdType
}

func (s *QueryInspirationBalanceForPartnerRequest) SetBelongId(v string) *QueryInspirationBalanceForPartnerRequest {
	s.BelongId = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerRequest) SetBelongIdType(v string) *QueryInspirationBalanceForPartnerRequest {
	s.BelongIdType = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
