// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByVpcAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiVpcAccessInfos(v *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos) *DescribeApisByVpcAccessResponseBody
	GetApiVpcAccessInfos() *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos
	SetPageNumber(v int32) *DescribeApisByVpcAccessResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByVpcAccessResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApisByVpcAccessResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApisByVpcAccessResponseBody
	GetTotalCount() *int32
}

type DescribeApisByVpcAccessResponseBody struct {
	// The returned API information. It is an array consisting of ApiInfo data.
	ApiVpcAccessInfos *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos `json:"ApiVpcAccessInfos,omitempty" xml:"ApiVpcAccessInfos,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4E707B25-5119-5ACF-9D26-7D2A2762F05C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisByVpcAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByVpcAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisByVpcAccessResponseBody) GetApiVpcAccessInfos() *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos {
	return s.ApiVpcAccessInfos
}

func (s *DescribeApisByVpcAccessResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByVpcAccessResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByVpcAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisByVpcAccessResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApisByVpcAccessResponseBody) SetApiVpcAccessInfos(v *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos) *DescribeApisByVpcAccessResponseBody {
	s.ApiVpcAccessInfos = v
	return s
}

func (s *DescribeApisByVpcAccessResponseBody) SetPageNumber(v int32) *DescribeApisByVpcAccessResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBody) SetPageSize(v int32) *DescribeApisByVpcAccessResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBody) SetRequestId(v string) *DescribeApisByVpcAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBody) SetTotalCount(v int32) *DescribeApisByVpcAccessResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBody) Validate() error {
	if s.ApiVpcAccessInfos != nil {
		if err := s.ApiVpcAccessInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos struct {
	ApiVpcAccessInfo []*DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo `json:"ApiVpcAccessInfo,omitempty" xml:"ApiVpcAccessInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos) GetApiVpcAccessInfo() []*DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	return s.ApiVpcAccessInfo
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos) SetApiVpcAccessInfo(v []*DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos {
	s.ApiVpcAccessInfo = v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfos) Validate() error {
	if s.ApiVpcAccessInfo != nil {
		for _, item := range s.ApiVpcAccessInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo struct {
	// The API ID.
	//
	// example:
	//
	// 09839002de484e76b5a213b040a6a3ca
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// iwc
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The description, which can be up to 200 characters in length.
	//
	// example:
	//
	// 123
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group to which the API belongs.
	//
	// example:
	//
	// 41c33748cbfb41f6b00870156203b72a
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group to which the API belongs.
	//
	// example:
	//
	// RT_PLU_IP_CTRL_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The instance ID or IP address in the VPC access authorization.
	//
	// example:
	//
	// Ib-04e41XXXXXd95e9c1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The HTTP request method of the API.
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /api/v1/friends/rc/status
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port number.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// 57e17906a69b4424914fb1e05f67c78a
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The environment to which the API is published. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the staging environment
	//
	// 	- **TEST**: the test environment
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// vpc id
	//
	// example:
	//
	// vpc-2ze7bj64wstznvftrskbk
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC access authorization.
	//
	// example:
	//
	// aliYun_service_prod
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetMethod() *string {
	return s.Method
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetPath() *string {
	return s.Path
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetPort() *int32 {
	return s.Port
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetStageId() *string {
	return s.StageId
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetApiId(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetApiName(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetDescription(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetGroupId(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetGroupName(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetInstanceId(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetMethod(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.Method = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetPath(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.Path = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetPort(v int32) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.Port = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetRegionId(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetStageId(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.StageId = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetStageName(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetVpcId(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) SetVpcName(v string) *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo {
	s.VpcName = &v
	return s
}

func (s *DescribeApisByVpcAccessResponseBodyApiVpcAccessInfosApiVpcAccessInfo) Validate() error {
	return dara.Validate(s)
}
