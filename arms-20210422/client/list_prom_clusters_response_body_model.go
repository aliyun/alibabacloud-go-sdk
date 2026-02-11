// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPromClusterList(v []*ListPromClustersResponseBodyPromClusterList) *ListPromClustersResponseBody
	GetPromClusterList() []*ListPromClustersResponseBodyPromClusterList
	SetRequestId(v string) *ListPromClustersResponseBody
	GetRequestId() *string
}

type ListPromClustersResponseBody struct {
	PromClusterList []*ListPromClustersResponseBodyPromClusterList `json:"PromClusterList,omitempty" xml:"PromClusterList,omitempty" type:"Repeated"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPromClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPromClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListPromClustersResponseBody) GetPromClusterList() []*ListPromClustersResponseBodyPromClusterList {
	return s.PromClusterList
}

func (s *ListPromClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPromClustersResponseBody) SetPromClusterList(v []*ListPromClustersResponseBodyPromClusterList) *ListPromClustersResponseBody {
	s.PromClusterList = v
	return s
}

func (s *ListPromClustersResponseBody) SetRequestId(v string) *ListPromClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPromClustersResponseBody) Validate() error {
	if s.PromClusterList != nil {
		for _, item := range s.PromClusterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPromClustersResponseBodyPromClusterList struct {
	AgentStatus           *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName           *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType           *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ControllerId          *string `json:"ControllerId,omitempty" xml:"ControllerId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Extra                 *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstallTime           *int64  `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	IsControllerInstalled *bool   `json:"IsControllerInstalled,omitempty" xml:"IsControllerInstalled,omitempty"`
	LastHeartBeatTime     *int64  `json:"LastHeartBeatTime,omitempty" xml:"LastHeartBeatTime,omitempty"`
	NodeNum               *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Options               *string `json:"Options,omitempty" xml:"Options,omitempty"`
	PluginsJsonArray      *string `json:"PluginsJsonArray,omitempty" xml:"PluginsJsonArray,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StateJson             *string `json:"StateJson,omitempty" xml:"StateJson,omitempty"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListPromClustersResponseBodyPromClusterList) String() string {
	return dara.Prettify(s)
}

func (s ListPromClustersResponseBodyPromClusterList) GoString() string {
	return s.String()
}

func (s *ListPromClustersResponseBodyPromClusterList) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *ListPromClustersResponseBodyPromClusterList) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPromClustersResponseBodyPromClusterList) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListPromClustersResponseBodyPromClusterList) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListPromClustersResponseBodyPromClusterList) GetControllerId() *string {
	return s.ControllerId
}

func (s *ListPromClustersResponseBodyPromClusterList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListPromClustersResponseBodyPromClusterList) GetExtra() *string {
	return s.Extra
}

func (s *ListPromClustersResponseBodyPromClusterList) GetId() *int64 {
	return s.Id
}

func (s *ListPromClustersResponseBodyPromClusterList) GetInstallTime() *int64 {
	return s.InstallTime
}

func (s *ListPromClustersResponseBodyPromClusterList) GetIsControllerInstalled() *bool {
	return s.IsControllerInstalled
}

func (s *ListPromClustersResponseBodyPromClusterList) GetLastHeartBeatTime() *int64 {
	return s.LastHeartBeatTime
}

func (s *ListPromClustersResponseBodyPromClusterList) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *ListPromClustersResponseBodyPromClusterList) GetOptions() *string {
	return s.Options
}

func (s *ListPromClustersResponseBodyPromClusterList) GetPluginsJsonArray() *string {
	return s.PluginsJsonArray
}

func (s *ListPromClustersResponseBodyPromClusterList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPromClustersResponseBodyPromClusterList) GetStateJson() *string {
	return s.StateJson
}

func (s *ListPromClustersResponseBodyPromClusterList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListPromClustersResponseBodyPromClusterList) GetUserId() *string {
	return s.UserId
}

func (s *ListPromClustersResponseBodyPromClusterList) SetAgentStatus(v string) *ListPromClustersResponseBodyPromClusterList {
	s.AgentStatus = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetClusterId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ClusterId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetClusterName(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ClusterName = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetClusterType(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ClusterType = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetControllerId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ControllerId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetCreateTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.CreateTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetExtra(v string) *ListPromClustersResponseBodyPromClusterList {
	s.Extra = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetId(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.Id = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetInstallTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.InstallTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetIsControllerInstalled(v bool) *ListPromClustersResponseBodyPromClusterList {
	s.IsControllerInstalled = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetLastHeartBeatTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.LastHeartBeatTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetNodeNum(v int32) *ListPromClustersResponseBodyPromClusterList {
	s.NodeNum = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetOptions(v string) *ListPromClustersResponseBodyPromClusterList {
	s.Options = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetPluginsJsonArray(v string) *ListPromClustersResponseBodyPromClusterList {
	s.PluginsJsonArray = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetRegionId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.RegionId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetStateJson(v string) *ListPromClustersResponseBodyPromClusterList {
	s.StateJson = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetUpdateTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.UpdateTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetUserId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.UserId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) Validate() error {
	return dara.Validate(s)
}
