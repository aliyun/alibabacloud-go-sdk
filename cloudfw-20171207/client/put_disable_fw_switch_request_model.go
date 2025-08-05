// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDisableFwSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpaddrList(v []*string) *PutDisableFwSwitchRequest
	GetIpaddrList() []*string
	SetLang(v string) *PutDisableFwSwitchRequest
	GetLang() *string
	SetRegionList(v []*string) *PutDisableFwSwitchRequest
	GetRegionList() []*string
	SetResourceTypeList(v []*string) *PutDisableFwSwitchRequest
	GetResourceTypeList() []*string
	SetSourceIp(v string) *PutDisableFwSwitchRequest
	GetSourceIp() *string
}

type PutDisableFwSwitchRequest struct {
	// The IP addresses.
	//
	// >  You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	//
	// example:
	//
	// ["192.0.XX.XX","192.0.XX.XX"]
	IpaddrList []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The regions.
	//
	// >  You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	//
	// example:
	//
	// ["cn-hangzhou","cn-shanghai"]
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The types of the assets.
	//
	// > You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	//
	// example:
	//
	// ["EcsPublicIp","NatEip"]
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutDisableFwSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s PutDisableFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchRequest) GetIpaddrList() []*string {
	return s.IpaddrList
}

func (s *PutDisableFwSwitchRequest) GetLang() *string {
	return s.Lang
}

func (s *PutDisableFwSwitchRequest) GetRegionList() []*string {
	return s.RegionList
}

func (s *PutDisableFwSwitchRequest) GetResourceTypeList() []*string {
	return s.ResourceTypeList
}

func (s *PutDisableFwSwitchRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *PutDisableFwSwitchRequest) SetIpaddrList(v []*string) *PutDisableFwSwitchRequest {
	s.IpaddrList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetLang(v string) *PutDisableFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutDisableFwSwitchRequest) SetRegionList(v []*string) *PutDisableFwSwitchRequest {
	s.RegionList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetResourceTypeList(v []*string) *PutDisableFwSwitchRequest {
	s.ResourceTypeList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetSourceIp(v string) *PutDisableFwSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *PutDisableFwSwitchRequest) Validate() error {
	return dara.Validate(s)
}
