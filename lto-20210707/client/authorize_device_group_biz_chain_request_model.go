// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeDeviceGroupBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainIdList(v string) *AuthorizeDeviceGroupBizChainRequest
	GetBizChainIdList() *string
	SetDeviceGroupId(v string) *AuthorizeDeviceGroupBizChainRequest
	GetDeviceGroupId() *string
	SetRegionId(v string) *AuthorizeDeviceGroupBizChainRequest
	GetRegionId() *string
}

type AuthorizeDeviceGroupBizChainRequest struct {
	BizChainIdList *string `json:"BizChainIdList,omitempty" xml:"BizChainIdList,omitempty"`
	// This parameter is required.
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AuthorizeDeviceGroupBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeDeviceGroupBizChainRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeDeviceGroupBizChainRequest) GetBizChainIdList() *string {
	return s.BizChainIdList
}

func (s *AuthorizeDeviceGroupBizChainRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *AuthorizeDeviceGroupBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeDeviceGroupBizChainRequest) SetBizChainIdList(v string) *AuthorizeDeviceGroupBizChainRequest {
	s.BizChainIdList = &v
	return s
}

func (s *AuthorizeDeviceGroupBizChainRequest) SetDeviceGroupId(v string) *AuthorizeDeviceGroupBizChainRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AuthorizeDeviceGroupBizChainRequest) SetRegionId(v string) *AuthorizeDeviceGroupBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeDeviceGroupBizChainRequest) Validate() error {
	return dara.Validate(s)
}
