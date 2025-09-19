// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckSaleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetCheckSaleRequest
	GetRegionId() *string
}

type GetCheckSaleRequest struct {
	// The region in which the Security Center instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou:*	- inside China
	//
	// 	- Global **ap-southeast-1:*	- outside China
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCheckSaleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSaleRequest) GoString() string {
	return s.String()
}

func (s *GetCheckSaleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCheckSaleRequest) SetRegionId(v string) *GetCheckSaleRequest {
	s.RegionId = &v
	return s
}

func (s *GetCheckSaleRequest) Validate() error {
	return dara.Validate(s)
}
