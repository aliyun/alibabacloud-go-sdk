// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstIp(v string) *GetDeploymentsByIpRequest
	GetDstIp() *string
	SetDstPort(v string) *GetDeploymentsByIpRequest
	GetDstPort() *string
	SetIgnoreJobSummary(v bool) *GetDeploymentsByIpRequest
	GetIgnoreJobSummary() *bool
	SetIgnoreResourceSetting(v bool) *GetDeploymentsByIpRequest
	GetIgnoreResourceSetting() *bool
	SetSrcIp(v string) *GetDeploymentsByIpRequest
	GetSrcIp() *string
	SetSrcPort(v string) *GetDeploymentsByIpRequest
	GetSrcPort() *string
}

type GetDeploymentsByIpRequest struct {
	// example:
	//
	// 10.100.2.200
	DstIp *string `json:"dstIp,omitempty" xml:"dstIp,omitempty"`
	// example:
	//
	// 9092
	DstPort *string `json:"dstPort,omitempty" xml:"dstPort,omitempty"`
	// example:
	//
	// false
	IgnoreJobSummary *bool `json:"ignoreJobSummary,omitempty" xml:"ignoreJobSummary,omitempty"`
	// example:
	//
	// false
	IgnoreResourceSetting *bool `json:"ignoreResourceSetting,omitempty" xml:"ignoreResourceSetting,omitempty"`
	// example:
	//
	// 192.168.1.100
	SrcIp *string `json:"srcIp,omitempty" xml:"srcIp,omitempty"`
	// example:
	//
	// 54321
	SrcPort *string `json:"srcPort,omitempty" xml:"srcPort,omitempty"`
}

func (s GetDeploymentsByIpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByIpRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByIpRequest) GetDstIp() *string {
	return s.DstIp
}

func (s *GetDeploymentsByIpRequest) GetDstPort() *string {
	return s.DstPort
}

func (s *GetDeploymentsByIpRequest) GetIgnoreJobSummary() *bool {
	return s.IgnoreJobSummary
}

func (s *GetDeploymentsByIpRequest) GetIgnoreResourceSetting() *bool {
	return s.IgnoreResourceSetting
}

func (s *GetDeploymentsByIpRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *GetDeploymentsByIpRequest) GetSrcPort() *string {
	return s.SrcPort
}

func (s *GetDeploymentsByIpRequest) SetDstIp(v string) *GetDeploymentsByIpRequest {
	s.DstIp = &v
	return s
}

func (s *GetDeploymentsByIpRequest) SetDstPort(v string) *GetDeploymentsByIpRequest {
	s.DstPort = &v
	return s
}

func (s *GetDeploymentsByIpRequest) SetIgnoreJobSummary(v bool) *GetDeploymentsByIpRequest {
	s.IgnoreJobSummary = &v
	return s
}

func (s *GetDeploymentsByIpRequest) SetIgnoreResourceSetting(v bool) *GetDeploymentsByIpRequest {
	s.IgnoreResourceSetting = &v
	return s
}

func (s *GetDeploymentsByIpRequest) SetSrcIp(v string) *GetDeploymentsByIpRequest {
	s.SrcIp = &v
	return s
}

func (s *GetDeploymentsByIpRequest) SetSrcPort(v string) *GetDeploymentsByIpRequest {
	s.SrcPort = &v
	return s
}

func (s *GetDeploymentsByIpRequest) Validate() error {
	return dara.Validate(s)
}
