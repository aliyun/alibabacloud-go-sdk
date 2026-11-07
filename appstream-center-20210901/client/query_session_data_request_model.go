// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QuerySessionDataRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QuerySessionDataRequest
	GetBizRegionId() *string
	SetRegionId(v string) *QuerySessionDataRequest
	GetRegionId() *string
	SetStatDate(v string) *QuerySessionDataRequest
	GetStatDate() *string
}

type QuerySessionDataRequest struct {
	// example:
	//
	// aig-89ibriac2wudyph38
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 20261106
	StatDate *string `json:"StatDate,omitempty" xml:"StatDate,omitempty"`
}

func (s QuerySessionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionDataRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionDataRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QuerySessionDataRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QuerySessionDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QuerySessionDataRequest) GetStatDate() *string {
	return s.StatDate
}

func (s *QuerySessionDataRequest) SetAppInstanceGroupId(v string) *QuerySessionDataRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QuerySessionDataRequest) SetBizRegionId(v string) *QuerySessionDataRequest {
	s.BizRegionId = &v
	return s
}

func (s *QuerySessionDataRequest) SetRegionId(v string) *QuerySessionDataRequest {
	s.RegionId = &v
	return s
}

func (s *QuerySessionDataRequest) SetStatDate(v string) *QuerySessionDataRequest {
	s.StatDate = &v
	return s
}

func (s *QuerySessionDataRequest) Validate() error {
	return dara.Validate(s)
}
