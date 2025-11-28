// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceGroupAuthorizedBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupId(v string) *ListDeviceGroupAuthorizedBizChainRequest
	GetDeviceGroupId() *string
	SetRegionId(v string) *ListDeviceGroupAuthorizedBizChainRequest
	GetRegionId() *string
}

type ListDeviceGroupAuthorizedBizChainRequest struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDeviceGroupAuthorizedBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupAuthorizedBizChainRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupAuthorizedBizChainRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *ListDeviceGroupAuthorizedBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDeviceGroupAuthorizedBizChainRequest) SetDeviceGroupId(v string) *ListDeviceGroupAuthorizedBizChainRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainRequest) SetRegionId(v string) *ListDeviceGroupAuthorizedBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainRequest) Validate() error {
	return dara.Validate(s)
}
