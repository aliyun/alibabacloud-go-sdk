// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountPointsVscAttachInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputRegionId(v string) *DescribeMountPointsVscAttachInfoShrinkRequest
	GetInputRegionId() *string
	SetMaxResults(v int32) *DescribeMountPointsVscAttachInfoShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeMountPointsVscAttachInfoShrinkRequest
	GetNextToken() *string
	SetQueryInfosShrink(v string) *DescribeMountPointsVscAttachInfoShrinkRequest
	GetQueryInfosShrink() *string
	SetUseAssumeRoleChkServerPerm(v bool) *DescribeMountPointsVscAttachInfoShrinkRequest
	GetUseAssumeRoleChkServerPerm() *bool
}

type DescribeMountPointsVscAttachInfoShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MYR6sz6qkmauspAy8oxjHP-tOLtD4KSv3DzI7G6iywTx1ZCExO50GtSuiTB9z0JppvYQ2iUa8s4HbcFanMQfDIlds4da87_Ax4UJMva****
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QueryInfosShrink *string `json:"QueryInfos,omitempty" xml:"QueryInfos,omitempty"`
	// example:
	//
	// false
	UseAssumeRoleChkServerPerm *bool `json:"UseAssumeRoleChkServerPerm,omitempty" xml:"UseAssumeRoleChkServerPerm,omitempty"`
}

func (s DescribeMountPointsVscAttachInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountPointsVscAttachInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) GetQueryInfosShrink() *string {
	return s.QueryInfosShrink
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) SetInputRegionId(v string) *DescribeMountPointsVscAttachInfoShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) SetMaxResults(v int32) *DescribeMountPointsVscAttachInfoShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) SetNextToken(v string) *DescribeMountPointsVscAttachInfoShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) SetQueryInfosShrink(v string) *DescribeMountPointsVscAttachInfoShrinkRequest {
	s.QueryInfosShrink = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) SetUseAssumeRoleChkServerPerm(v bool) *DescribeMountPointsVscAttachInfoShrinkRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
