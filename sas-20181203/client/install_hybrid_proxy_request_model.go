// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallHybridProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *InstallHybridProxyRequest
	GetClusterName() *string
	SetInstallCode(v string) *InstallHybridProxyRequest
	GetInstallCode() *string
	SetYundunUuids(v []*string) *InstallHybridProxyRequest
	GetYundunUuids() []*string
}

type InstallHybridProxyRequest struct {
	// The cluster name.
	//
	// example:
	//
	// proxy-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The installation code.
	//
	// example:
	//
	// Z9c8SA
	InstallCode *string `json:"InstallCode,omitempty" xml:"InstallCode,omitempty"`
	// The UUIDs of the proxy servers.
	YundunUuids []*string `json:"YundunUuids,omitempty" xml:"YundunUuids,omitempty" type:"Repeated"`
}

func (s InstallHybridProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallHybridProxyRequest) GoString() string {
	return s.String()
}

func (s *InstallHybridProxyRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *InstallHybridProxyRequest) GetInstallCode() *string {
	return s.InstallCode
}

func (s *InstallHybridProxyRequest) GetYundunUuids() []*string {
	return s.YundunUuids
}

func (s *InstallHybridProxyRequest) SetClusterName(v string) *InstallHybridProxyRequest {
	s.ClusterName = &v
	return s
}

func (s *InstallHybridProxyRequest) SetInstallCode(v string) *InstallHybridProxyRequest {
	s.InstallCode = &v
	return s
}

func (s *InstallHybridProxyRequest) SetYundunUuids(v []*string) *InstallHybridProxyRequest {
	s.YundunUuids = v
	return s
}

func (s *InstallHybridProxyRequest) Validate() error {
	return dara.Validate(s)
}
