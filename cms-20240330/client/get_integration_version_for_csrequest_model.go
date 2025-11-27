// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationVersionForCSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetIntegrationVersionForCSRequest
	GetClusterId() *string
	SetClusterType(v string) *GetIntegrationVersionForCSRequest
	GetClusterType() *string
}

type GetIntegrationVersionForCSRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c502646fd0d1249baaf792b3a1b589e1b
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs.ack.cluster
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
}

func (s GetIntegrationVersionForCSRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationVersionForCSRequest) GoString() string {
	return s.String()
}

func (s *GetIntegrationVersionForCSRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetIntegrationVersionForCSRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *GetIntegrationVersionForCSRequest) SetClusterId(v string) *GetIntegrationVersionForCSRequest {
	s.ClusterId = &v
	return s
}

func (s *GetIntegrationVersionForCSRequest) SetClusterType(v string) *GetIntegrationVersionForCSRequest {
	s.ClusterType = &v
	return s
}

func (s *GetIntegrationVersionForCSRequest) Validate() error {
	return dara.Validate(s)
}
