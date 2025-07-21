// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuyResourceIds(v string) *DedicatedIpPoolUpdateRequest
	GetBuyResourceIds() *string
	SetId(v string) *DedicatedIpPoolUpdateRequest
	GetId() *string
	SetUpdateResource(v bool) *DedicatedIpPoolUpdateRequest
	GetUpdateResource() *bool
}

type DedicatedIpPoolUpdateRequest struct {
	// Purchased IP instance IDs, separated by commas; sourced from the DedicatedIpNonePoolList API\\"s returned IP purchase instance IDs
	//
	// example:
	//
	// xxx,xxx
	BuyResourceIds *string `json:"BuyResourceIds,omitempty" xml:"BuyResourceIds,omitempty"`
	// IP pool ID
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Whether to change the associated IP, enter true
	//
	// example:
	//
	// true
	UpdateResource *bool `json:"UpdateResource,omitempty" xml:"UpdateResource,omitempty"`
}

func (s DedicatedIpPoolUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolUpdateRequest) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolUpdateRequest) GetBuyResourceIds() *string {
	return s.BuyResourceIds
}

func (s *DedicatedIpPoolUpdateRequest) GetId() *string {
	return s.Id
}

func (s *DedicatedIpPoolUpdateRequest) GetUpdateResource() *bool {
	return s.UpdateResource
}

func (s *DedicatedIpPoolUpdateRequest) SetBuyResourceIds(v string) *DedicatedIpPoolUpdateRequest {
	s.BuyResourceIds = &v
	return s
}

func (s *DedicatedIpPoolUpdateRequest) SetId(v string) *DedicatedIpPoolUpdateRequest {
	s.Id = &v
	return s
}

func (s *DedicatedIpPoolUpdateRequest) SetUpdateResource(v bool) *DedicatedIpPoolUpdateRequest {
	s.UpdateResource = &v
	return s
}

func (s *DedicatedIpPoolUpdateRequest) Validate() error {
	return dara.Validate(s)
}
