// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *ListUserResourcesRequest
	GetCommodityCode() *string
}

type ListUserResourcesRequest struct {
	// example:
	//
	// BrainIndustrial
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s ListUserResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListUserResourcesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListUserResourcesRequest) SetCommodityCode(v string) *ListUserResourcesRequest {
	s.CommodityCode = &v
	return s
}

func (s *ListUserResourcesRequest) Validate() error {
	return dara.Validate(s)
}
