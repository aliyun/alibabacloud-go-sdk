// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSupportDBTableRecoveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *SupportDBTableRecoveryRequest
	GetClusterName() *string
	SetInstanceId(v string) *SupportDBTableRecoveryRequest
	GetInstanceId() *string
	SetRegionCode(v string) *SupportDBTableRecoveryRequest
	GetRegionCode() *string
}

type SupportDBTableRecoveryRequest struct {
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode  *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s SupportDBTableRecoveryRequest) String() string {
	return dara.Prettify(s)
}

func (s SupportDBTableRecoveryRequest) GoString() string {
	return s.String()
}

func (s *SupportDBTableRecoveryRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *SupportDBTableRecoveryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SupportDBTableRecoveryRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *SupportDBTableRecoveryRequest) SetClusterName(v string) *SupportDBTableRecoveryRequest {
	s.ClusterName = &v
	return s
}

func (s *SupportDBTableRecoveryRequest) SetInstanceId(v string) *SupportDBTableRecoveryRequest {
	s.InstanceId = &v
	return s
}

func (s *SupportDBTableRecoveryRequest) SetRegionCode(v string) *SupportDBTableRecoveryRequest {
	s.RegionCode = &v
	return s
}

func (s *SupportDBTableRecoveryRequest) Validate() error {
	return dara.Validate(s)
}
