// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElasticNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticNetworkInterfaceId(v string) *GetElasticNetworkInterfaceRequest
	GetElasticNetworkInterfaceId() *string
	SetRegionId(v string) *GetElasticNetworkInterfaceRequest
	GetRegionId() *string
}

type GetElasticNetworkInterfaceRequest struct {
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

func (s GetElasticNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *GetElasticNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *GetElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetElasticNetworkInterfaceRequest) SetRegionId(v string) *GetElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *GetElasticNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
