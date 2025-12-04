// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteElasticNetworkInterfaceRequest
	GetClientToken() *string
	SetElasticNetworkInterfaceId(v string) *DeleteElasticNetworkInterfaceRequest
	GetElasticNetworkInterfaceId() *string
	SetRegionId(v string) *DeleteElasticNetworkInterfaceRequest
	GetRegionId() *string
}

type DeleteElasticNetworkInterfaceRequest struct {
	// By default, popApi is not ignored and idempotent
	//
	// example:
	//
	// 141cccd6-dfbd-11ec-b8e8-0242ac110003
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteElasticNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteElasticNetworkInterfaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteElasticNetworkInterfaceRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *DeleteElasticNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteElasticNetworkInterfaceRequest) SetClientToken(v string) *DeleteElasticNetworkInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *DeleteElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceRequest) SetRegionId(v string) *DeleteElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
