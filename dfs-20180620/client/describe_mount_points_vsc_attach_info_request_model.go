// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountPointsVscAttachInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputRegionId(v string) *DescribeMountPointsVscAttachInfoRequest
	GetInputRegionId() *string
	SetMaxResults(v int32) *DescribeMountPointsVscAttachInfoRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeMountPointsVscAttachInfoRequest
	GetNextToken() *string
	SetQueryInfos(v []*DescribeMountPointsVscAttachInfoRequestQueryInfos) *DescribeMountPointsVscAttachInfoRequest
	GetQueryInfos() []*DescribeMountPointsVscAttachInfoRequestQueryInfos
	SetUseAssumeRoleChkServerPerm(v bool) *DescribeMountPointsVscAttachInfoRequest
	GetUseAssumeRoleChkServerPerm() *bool
}

type DescribeMountPointsVscAttachInfoRequest struct {
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
	NextToken  *string                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QueryInfos []*DescribeMountPointsVscAttachInfoRequestQueryInfos `json:"QueryInfos,omitempty" xml:"QueryInfos,omitempty" type:"Repeated"`
	// example:
	//
	// false
	UseAssumeRoleChkServerPerm *bool `json:"UseAssumeRoleChkServerPerm,omitempty" xml:"UseAssumeRoleChkServerPerm,omitempty"`
}

func (s DescribeMountPointsVscAttachInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountPointsVscAttachInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountPointsVscAttachInfoRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DescribeMountPointsVscAttachInfoRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeMountPointsVscAttachInfoRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMountPointsVscAttachInfoRequest) GetQueryInfos() []*DescribeMountPointsVscAttachInfoRequestQueryInfos {
	return s.QueryInfos
}

func (s *DescribeMountPointsVscAttachInfoRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *DescribeMountPointsVscAttachInfoRequest) SetInputRegionId(v string) *DescribeMountPointsVscAttachInfoRequest {
	s.InputRegionId = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoRequest) SetMaxResults(v int32) *DescribeMountPointsVscAttachInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoRequest) SetNextToken(v string) *DescribeMountPointsVscAttachInfoRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoRequest) SetQueryInfos(v []*DescribeMountPointsVscAttachInfoRequestQueryInfos) *DescribeMountPointsVscAttachInfoRequest {
	s.QueryInfos = v
	return s
}

func (s *DescribeMountPointsVscAttachInfoRequest) SetUseAssumeRoleChkServerPerm(v bool) *DescribeMountPointsVscAttachInfoRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoRequest) Validate() error {
	if s.QueryInfos != nil {
		for _, item := range s.QueryInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMountPointsVscAttachInfoRequestQueryInfos struct {
	// example:
	//
	// i-2zehyz70ednszl6rrfj6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// f-9dd3c6bajmy110.cn-zhangjiakou.dfs.aliyuncs.com
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// example:
	//
	// vsc-bp19yqmujug2r18z0h9qal
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DescribeMountPointsVscAttachInfoRequestQueryInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountPointsVscAttachInfoRequestQueryInfos) GoString() string {
	return s.String()
}

func (s *DescribeMountPointsVscAttachInfoRequestQueryInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMountPointsVscAttachInfoRequestQueryInfos) GetMountPointId() *string {
	return s.MountPointId
}

func (s *DescribeMountPointsVscAttachInfoRequestQueryInfos) GetVscId() *string {
	return s.VscId
}

func (s *DescribeMountPointsVscAttachInfoRequestQueryInfos) SetInstanceId(v string) *DescribeMountPointsVscAttachInfoRequestQueryInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoRequestQueryInfos) SetMountPointId(v string) *DescribeMountPointsVscAttachInfoRequestQueryInfos {
	s.MountPointId = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoRequestQueryInfos) SetVscId(v string) *DescribeMountPointsVscAttachInfoRequestQueryInfos {
	s.VscId = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoRequestQueryInfos) Validate() error {
	return dara.Validate(s)
}
