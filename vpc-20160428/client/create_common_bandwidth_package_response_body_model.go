// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommonBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *CreateCommonBandwidthPackageResponseBody
	GetBandwidthPackageId() *string
	SetRequestId(v string) *CreateCommonBandwidthPackageResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateCommonBandwidthPackageResponseBody
	GetResourceGroupId() *string
}

type CreateCommonBandwidthPackageResponseBody struct {
	// The ID of the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// cbwp-bp1vevu8h3ieh****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FF39F653-033E-4CD9-9EDF-3CCA5A71FBC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateCommonBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCommonBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommonBandwidthPackageResponseBody) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateCommonBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCommonBandwidthPackageResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCommonBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *CreateCommonBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateCommonBandwidthPackageResponseBody) SetRequestId(v string) *CreateCommonBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCommonBandwidthPackageResponseBody) SetResourceGroupId(v string) *CreateCommonBandwidthPackageResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCommonBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
