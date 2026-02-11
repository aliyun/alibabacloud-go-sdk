// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterFromGrafanaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPromClusterList(v []*ListClusterFromGrafanaResponseBodyPromClusterList) *ListClusterFromGrafanaResponseBody
	GetPromClusterList() []*ListClusterFromGrafanaResponseBodyPromClusterList
	SetRequestId(v string) *ListClusterFromGrafanaResponseBody
	GetRequestId() *string
}

type ListClusterFromGrafanaResponseBody struct {
	PromClusterList []*ListClusterFromGrafanaResponseBodyPromClusterList `json:"PromClusterList,omitempty" xml:"PromClusterList,omitempty" type:"Repeated"`
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterFromGrafanaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterFromGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponseBody) GetPromClusterList() []*ListClusterFromGrafanaResponseBodyPromClusterList {
	return s.PromClusterList
}

func (s *ListClusterFromGrafanaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterFromGrafanaResponseBody) SetPromClusterList(v []*ListClusterFromGrafanaResponseBodyPromClusterList) *ListClusterFromGrafanaResponseBody {
	s.PromClusterList = v
	return s
}

func (s *ListClusterFromGrafanaResponseBody) SetRequestId(v string) *ListClusterFromGrafanaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBody) Validate() error {
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

type ListClusterFromGrafanaResponseBodyPromClusterList struct {
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

func (s ListClusterFromGrafanaResponseBodyPromClusterList) String() string {
	return dara.Prettify(s)
}

func (s ListClusterFromGrafanaResponseBodyPromClusterList) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetControllerId() *string {
	return s.ControllerId
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetExtra() *string {
	return s.Extra
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetId() *int64 {
	return s.Id
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetInstallTime() *int64 {
	return s.InstallTime
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetIsControllerInstalled() *bool {
	return s.IsControllerInstalled
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetLastHeartBeatTime() *int64 {
	return s.LastHeartBeatTime
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetOptions() *string {
	return s.Options
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetPluginsJsonArray() *string {
	return s.PluginsJsonArray
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetStateJson() *string {
	return s.StateJson
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) GetUserId() *string {
	return s.UserId
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetAgentStatus(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.AgentStatus = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterName(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterName = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterType(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterType = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetControllerId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ControllerId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetCreateTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.CreateTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetExtra(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Extra = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetId(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Id = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetInstallTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.InstallTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetIsControllerInstalled(v bool) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.IsControllerInstalled = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetLastHeartBeatTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.LastHeartBeatTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetNodeNum(v int32) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.NodeNum = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetOptions(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Options = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetPluginsJsonArray(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.PluginsJsonArray = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetRegionId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.RegionId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetStateJson(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.StateJson = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetUpdateTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.UpdateTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetUserId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.UserId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) Validate() error {
	return dara.Validate(s)
}
