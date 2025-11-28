// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainName(v string) *ListRouteRuleRequest
	GetBizChainName() *string
	SetChainUpMode(v string) *ListRouteRuleRequest
	GetChainUpMode() *string
	SetDeviceGroupName(v string) *ListRouteRuleRequest
	GetDeviceGroupName() *string
	SetNum(v int32) *ListRouteRuleRequest
	GetNum() *int32
	SetRegionId(v string) *ListRouteRuleRequest
	GetRegionId() *string
	SetSize(v int32) *ListRouteRuleRequest
	GetSize() *int32
}

type ListRouteRuleRequest struct {
	BizChainName    *string `json:"BizChainName,omitempty" xml:"BizChainName,omitempty"`
	ChainUpMode     *string `json:"ChainUpMode,omitempty" xml:"ChainUpMode,omitempty"`
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	// This parameter is required.
	Num      *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListRouteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *ListRouteRuleRequest) GetBizChainName() *string {
	return s.BizChainName
}

func (s *ListRouteRuleRequest) GetChainUpMode() *string {
	return s.ChainUpMode
}

func (s *ListRouteRuleRequest) GetDeviceGroupName() *string {
	return s.DeviceGroupName
}

func (s *ListRouteRuleRequest) GetNum() *int32 {
	return s.Num
}

func (s *ListRouteRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRouteRuleRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListRouteRuleRequest) SetBizChainName(v string) *ListRouteRuleRequest {
	s.BizChainName = &v
	return s
}

func (s *ListRouteRuleRequest) SetChainUpMode(v string) *ListRouteRuleRequest {
	s.ChainUpMode = &v
	return s
}

func (s *ListRouteRuleRequest) SetDeviceGroupName(v string) *ListRouteRuleRequest {
	s.DeviceGroupName = &v
	return s
}

func (s *ListRouteRuleRequest) SetNum(v int32) *ListRouteRuleRequest {
	s.Num = &v
	return s
}

func (s *ListRouteRuleRequest) SetRegionId(v string) *ListRouteRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListRouteRuleRequest) SetSize(v int32) *ListRouteRuleRequest {
	s.Size = &v
	return s
}

func (s *ListRouteRuleRequest) Validate() error {
	return dara.Validate(s)
}
