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
	// The IDs of the purchased IP instances. Separate multiple IDs with commas. Obtain these IDs from the response of the DedicatedIpNonePoolList operation.
	//
	// example:
	//
	// xxx,xxx
	BuyResourceIds *string `json:"BuyResourceIds,omitempty" xml:"BuyResourceIds,omitempty"`
	// The ID of the IP pool.
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to change the associated IP addresses. Set this parameter to true.
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
