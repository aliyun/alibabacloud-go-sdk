// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateYikeProductionMemberAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v string) *UpdateYikeProductionMemberAuthRequest
	GetAuth() *string
	SetProductionId(v string) *UpdateYikeProductionMemberAuthRequest
	GetProductionId() *string
	SetYikeUserId(v string) *UpdateYikeProductionMemberAuthRequest
	GetYikeUserId() *string
}

type UpdateYikeProductionMemberAuthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Manage
	Auth *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pd_061716****
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// This parameter is required.
	YikeUserId *string `json:"YikeUserId,omitempty" xml:"YikeUserId,omitempty"`
}

func (s UpdateYikeProductionMemberAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateYikeProductionMemberAuthRequest) GoString() string {
	return s.String()
}

func (s *UpdateYikeProductionMemberAuthRequest) GetAuth() *string {
	return s.Auth
}

func (s *UpdateYikeProductionMemberAuthRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *UpdateYikeProductionMemberAuthRequest) GetYikeUserId() *string {
	return s.YikeUserId
}

func (s *UpdateYikeProductionMemberAuthRequest) SetAuth(v string) *UpdateYikeProductionMemberAuthRequest {
	s.Auth = &v
	return s
}

func (s *UpdateYikeProductionMemberAuthRequest) SetProductionId(v string) *UpdateYikeProductionMemberAuthRequest {
	s.ProductionId = &v
	return s
}

func (s *UpdateYikeProductionMemberAuthRequest) SetYikeUserId(v string) *UpdateYikeProductionMemberAuthRequest {
	s.YikeUserId = &v
	return s
}

func (s *UpdateYikeProductionMemberAuthRequest) Validate() error {
	return dara.Validate(s)
}
