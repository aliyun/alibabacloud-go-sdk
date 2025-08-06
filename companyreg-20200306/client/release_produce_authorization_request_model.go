// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseProduceAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedUserId(v string) *ReleaseProduceAuthorizationRequest
	GetAuthorizedUserId() *string
	SetBizId(v string) *ReleaseProduceAuthorizationRequest
	GetBizId() *string
	SetBizType(v string) *ReleaseProduceAuthorizationRequest
	GetBizType() *string
}

type ReleaseProduceAuthorizationRequest struct {
	// example:
	//
	// 1219541161213000
	AuthorizedUserId *string `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// P20211117141528000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp.beian_assist
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s ReleaseProduceAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseProduceAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ReleaseProduceAuthorizationRequest) GetAuthorizedUserId() *string {
	return s.AuthorizedUserId
}

func (s *ReleaseProduceAuthorizationRequest) GetBizId() *string {
	return s.BizId
}

func (s *ReleaseProduceAuthorizationRequest) GetBizType() *string {
	return s.BizType
}

func (s *ReleaseProduceAuthorizationRequest) SetAuthorizedUserId(v string) *ReleaseProduceAuthorizationRequest {
	s.AuthorizedUserId = &v
	return s
}

func (s *ReleaseProduceAuthorizationRequest) SetBizId(v string) *ReleaseProduceAuthorizationRequest {
	s.BizId = &v
	return s
}

func (s *ReleaseProduceAuthorizationRequest) SetBizType(v string) *ReleaseProduceAuthorizationRequest {
	s.BizType = &v
	return s
}

func (s *ReleaseProduceAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
