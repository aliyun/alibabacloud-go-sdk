// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetThirdpartId(v string) *CostCenterDeleteRequest
	GetThirdpartId() *string
}

type CostCenterDeleteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 340049
	ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s CostCenterDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s CostCenterDeleteRequest) GoString() string {
	return s.String()
}

func (s *CostCenterDeleteRequest) GetThirdpartId() *string {
	return s.ThirdpartId
}

func (s *CostCenterDeleteRequest) SetThirdpartId(v string) *CostCenterDeleteRequest {
	s.ThirdpartId = &v
	return s
}

func (s *CostCenterDeleteRequest) Validate() error {
	return dara.Validate(s)
}
