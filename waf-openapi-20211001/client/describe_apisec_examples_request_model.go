// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecExamplesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormalTag(v string) *DescribeApisecExamplesRequest
	GetAbnormalTag() *string
	SetApiId(v string) *DescribeApisecExamplesRequest
	GetApiId() *string
	SetClusterId(v string) *DescribeApisecExamplesRequest
	GetClusterId() *string
	SetExampleType(v string) *DescribeApisecExamplesRequest
	GetExampleType() *string
	SetInstanceId(v string) *DescribeApisecExamplesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeApisecExamplesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeApisecExamplesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeApisecExamplesRequest
	GetRegionId() *string
	SetRequestSensitiveTypeList(v []*string) *DescribeApisecExamplesRequest
	GetRequestSensitiveTypeList() []*string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecExamplesRequest
	GetResourceManagerResourceGroupId() *string
	SetResponseSensitiveTypeList(v []*string) *DescribeApisecExamplesRequest
	GetResponseSensitiveTypeList() []*string
}

type DescribeApisecExamplesRequest struct {
	// example:
	//
	// LackOfSpeedLimit
	AbnormalTag *string `json:"AbnormalTag,omitempty" xml:"AbnormalTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 867ade***24ee6e205b8da82b8f84
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// 176
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// sensitive
	ExampleType *string `json:"ExampleType,omitempty" xml:"ExampleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqtm005
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestSensitiveTypeList []*string `json:"RequestSensitiveTypeList,omitempty" xml:"RequestSensitiveTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// example:
	//
	// LackOfSpeedLimit
	ResponseSensitiveTypeList []*string `json:"ResponseSensitiveTypeList,omitempty" xml:"ResponseSensitiveTypeList,omitempty" type:"Repeated"`
}

func (s DescribeApisecExamplesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecExamplesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecExamplesRequest) GetAbnormalTag() *string {
	return s.AbnormalTag
}

func (s *DescribeApisecExamplesRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisecExamplesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecExamplesRequest) GetExampleType() *string {
	return s.ExampleType
}

func (s *DescribeApisecExamplesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecExamplesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeApisecExamplesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeApisecExamplesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecExamplesRequest) GetRequestSensitiveTypeList() []*string {
	return s.RequestSensitiveTypeList
}

func (s *DescribeApisecExamplesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecExamplesRequest) GetResponseSensitiveTypeList() []*string {
	return s.ResponseSensitiveTypeList
}

func (s *DescribeApisecExamplesRequest) SetAbnormalTag(v string) *DescribeApisecExamplesRequest {
	s.AbnormalTag = &v
	return s
}

func (s *DescribeApisecExamplesRequest) SetApiId(v string) *DescribeApisecExamplesRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApisecExamplesRequest) SetClusterId(v string) *DescribeApisecExamplesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecExamplesRequest) SetExampleType(v string) *DescribeApisecExamplesRequest {
	s.ExampleType = &v
	return s
}

func (s *DescribeApisecExamplesRequest) SetInstanceId(v string) *DescribeApisecExamplesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecExamplesRequest) SetMaxResults(v int32) *DescribeApisecExamplesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeApisecExamplesRequest) SetNextToken(v string) *DescribeApisecExamplesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeApisecExamplesRequest) SetRegionId(v string) *DescribeApisecExamplesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecExamplesRequest) SetRequestSensitiveTypeList(v []*string) *DescribeApisecExamplesRequest {
	s.RequestSensitiveTypeList = v
	return s
}

func (s *DescribeApisecExamplesRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecExamplesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecExamplesRequest) SetResponseSensitiveTypeList(v []*string) *DescribeApisecExamplesRequest {
	s.ResponseSensitiveTypeList = v
	return s
}

func (s *DescribeApisecExamplesRequest) Validate() error {
	return dara.Validate(s)
}
