// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuyResourceIds(v string) *DedicatedIpPoolCreateRequest
	GetBuyResourceIds() *string
	SetName(v string) *DedicatedIpPoolCreateRequest
	GetName() *string
}

type DedicatedIpPoolCreateRequest struct {
	// The IDs of the purchased IP instances. Separate multiple IDs with commas (,). You can obtain the instance IDs from the response of the DedicatedIpNonePoolList operation.
	//
	// example:
	//
	// xxx,xxx
	BuyResourceIds *string `json:"BuyResourceIds,omitempty" xml:"BuyResourceIds,omitempty"`
	// The name of the IP pool. The name must be 1 to 50 characters in length. It can contain letters, digits, underscores (_), and hyphens (-). The name cannot be changed after the IP pool is created.
	//
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DedicatedIpPoolCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolCreateRequest) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolCreateRequest) GetBuyResourceIds() *string {
	return s.BuyResourceIds
}

func (s *DedicatedIpPoolCreateRequest) GetName() *string {
	return s.Name
}

func (s *DedicatedIpPoolCreateRequest) SetBuyResourceIds(v string) *DedicatedIpPoolCreateRequest {
	s.BuyResourceIds = &v
	return s
}

func (s *DedicatedIpPoolCreateRequest) SetName(v string) *DedicatedIpPoolCreateRequest {
	s.Name = &v
	return s
}

func (s *DedicatedIpPoolCreateRequest) Validate() error {
	return dara.Validate(s)
}
