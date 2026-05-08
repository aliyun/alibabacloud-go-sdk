// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddYikeProductionMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductionId(v string) *AddYikeProductionMembersRequest
	GetProductionId() *string
	SetYikeUserIds(v string) *AddYikeProductionMembersRequest
	GetYikeUserIds() *string
}

type AddYikeProductionMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pd_061716***
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21*****
	YikeUserIds *string `json:"YikeUserIds,omitempty" xml:"YikeUserIds,omitempty"`
}

func (s AddYikeProductionMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddYikeProductionMembersRequest) GoString() string {
	return s.String()
}

func (s *AddYikeProductionMembersRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *AddYikeProductionMembersRequest) GetYikeUserIds() *string {
	return s.YikeUserIds
}

func (s *AddYikeProductionMembersRequest) SetProductionId(v string) *AddYikeProductionMembersRequest {
	s.ProductionId = &v
	return s
}

func (s *AddYikeProductionMembersRequest) SetYikeUserIds(v string) *AddYikeProductionMembersRequest {
	s.YikeUserIds = &v
	return s
}

func (s *AddYikeProductionMembersRequest) Validate() error {
	return dara.Validate(s)
}
