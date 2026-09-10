// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLhmDWResourceGroupStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetLhmDWResourceGroupStatusRequest
	GetRegionId() *string
}

type GetLhmDWResourceGroupStatusRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetLhmDWResourceGroupStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLhmDWResourceGroupStatusRequest) GoString() string {
	return s.String()
}

func (s *GetLhmDWResourceGroupStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLhmDWResourceGroupStatusRequest) SetRegionId(v string) *GetLhmDWResourceGroupStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetLhmDWResourceGroupStatusRequest) Validate() error {
	return dara.Validate(s)
}
