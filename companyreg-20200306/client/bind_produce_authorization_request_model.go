// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindProduceAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedUserIds(v string) *BindProduceAuthorizationRequest
	GetAuthorizedUserIds() *string
	SetBizId(v string) *BindProduceAuthorizationRequest
	GetBizId() *string
	SetBizType(v string) *BindProduceAuthorizationRequest
	GetBizType() *string
}

type BindProduceAuthorizationRequest struct {
	// example:
	//
	// 1219541161213057,1219541161213059
	AuthorizedUserIds *string `json:"AuthorizedUserIds,omitempty" xml:"AuthorizedUserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// P20210815211849000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp.bookkeeping_ai
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s BindProduceAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s BindProduceAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationRequest) GetAuthorizedUserIds() *string {
	return s.AuthorizedUserIds
}

func (s *BindProduceAuthorizationRequest) GetBizId() *string {
	return s.BizId
}

func (s *BindProduceAuthorizationRequest) GetBizType() *string {
	return s.BizType
}

func (s *BindProduceAuthorizationRequest) SetAuthorizedUserIds(v string) *BindProduceAuthorizationRequest {
	s.AuthorizedUserIds = &v
	return s
}

func (s *BindProduceAuthorizationRequest) SetBizId(v string) *BindProduceAuthorizationRequest {
	s.BizId = &v
	return s
}

func (s *BindProduceAuthorizationRequest) SetBizType(v string) *BindProduceAuthorizationRequest {
	s.BizType = &v
	return s
}

func (s *BindProduceAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
