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
	// The delivery group ID, used to filter session data for a specified delivery group. If not provided, the summary of session data for all delivery groups under the current account is returned.
	//
	// example:
	//
	// aig-89ibriac2wudyph38
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The business region ID, which specifies the region for which to query session data. The server uses this value as the region filter condition for session data. If not provided, no region filtering is applied, and the summary of session data for all regions is returned.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The region ID, which is a common request parameter for the gateway. The region actually used to filter session data in this operation is specified by BizRegionId.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The statistical date in the yyyyMMdd format (without hyphens), such as 20261106. If not provided, the data for yesterday is queried by default. Only dates within the last 90 days are supported. If the date exceeds 90 days or the format is invalid, a parameter error is returned.
	//
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
