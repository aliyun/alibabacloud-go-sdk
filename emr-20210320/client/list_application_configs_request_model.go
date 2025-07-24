// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *ListApplicationConfigsRequest
	GetApplicationName() *string
	SetClusterId(v string) *ListApplicationConfigsRequest
	GetClusterId() *string
	SetConfigFileName(v string) *ListApplicationConfigsRequest
	GetConfigFileName() *string
	SetConfigItemKey(v string) *ListApplicationConfigsRequest
	GetConfigItemKey() *string
	SetConfigItemValue(v string) *ListApplicationConfigsRequest
	GetConfigItemValue() *string
	SetMaxResults(v int32) *ListApplicationConfigsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationConfigsRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListApplicationConfigsRequest
	GetNodeGroupId() *string
	SetNodeId(v string) *ListApplicationConfigsRequest
	GetNodeId() *string
	SetRegionId(v string) *ListApplicationConfigsRequest
	GetRegionId() *string
}

type ListApplicationConfigsRequest struct {
	// The name of the application.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-e6a9d46e9267****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the configuration file.
	//
	// example:
	//
	// hdfs-site.xml
	ConfigFileName *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	// The key of the configuration item.
	//
	// example:
	//
	// dfs.replication
	ConfigItemKey *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// 2
	ConfigItemValue *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number of the next page returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the node group.
	//
	// example:
	//
	// ng-d555335ced5c****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// i-bp10h9rezawz1i4o****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListApplicationConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationConfigsRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationConfigsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListApplicationConfigsRequest) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *ListApplicationConfigsRequest) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *ListApplicationConfigsRequest) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *ListApplicationConfigsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationConfigsRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListApplicationConfigsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListApplicationConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationConfigsRequest) SetApplicationName(v string) *ListApplicationConfigsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationConfigsRequest) SetClusterId(v string) *ListApplicationConfigsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListApplicationConfigsRequest) SetConfigFileName(v string) *ListApplicationConfigsRequest {
	s.ConfigFileName = &v
	return s
}

func (s *ListApplicationConfigsRequest) SetConfigItemKey(v string) *ListApplicationConfigsRequest {
	s.ConfigItemKey = &v
	return s
}

func (s *ListApplicationConfigsRequest) SetConfigItemValue(v string) *ListApplicationConfigsRequest {
	s.ConfigItemValue = &v
	return s
}

func (s *ListApplicationConfigsRequest) SetMaxResults(v int32) *ListApplicationConfigsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationConfigsRequest) SetNextToken(v string) *ListApplicationConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationConfigsRequest) SetNodeGroupId(v string) *ListApplicationConfigsRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListApplicationConfigsRequest) SetNodeId(v string) *ListApplicationConfigsRequest {
	s.NodeId = &v
	return s
}

func (s *ListApplicationConfigsRequest) SetRegionId(v string) *ListApplicationConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationConfigsRequest) Validate() error {
	return dara.Validate(s)
}
