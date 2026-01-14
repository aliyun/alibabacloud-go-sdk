// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *DeleteBandwidthPackageRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *DeleteBandwidthPackageRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteBandwidthPackageRequest
	GetRegionId() *string
}

type DeleteBandwidthPackageRequest struct {
	// The bandwidth plan ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DeleteBandwidthPackageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBandwidthPackageRequest) SetBandwidthPackageId(v string) *DeleteBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetClientToken(v string) *DeleteBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetRegionId(v string) *DeleteBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
