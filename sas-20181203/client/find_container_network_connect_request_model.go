// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindContainerNetworkConnectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaType(v string) *FindContainerNetworkConnectRequest
	GetCriteriaType() *string
	SetCurrentPage(v int64) *FindContainerNetworkConnectRequest
	GetCurrentPage() *int64
	SetDstNode(v *FindContainerNetworkConnectRequestDstNode) *FindContainerNetworkConnectRequest
	GetDstNode() *FindContainerNetworkConnectRequestDstNode
	SetEndTime(v int64) *FindContainerNetworkConnectRequest
	GetEndTime() *int64
	SetPageSize(v int64) *FindContainerNetworkConnectRequest
	GetPageSize() *int64
	SetSrcNode(v *FindContainerNetworkConnectRequestSrcNode) *FindContainerNetworkConnectRequest
	GetSrcNode() *FindContainerNetworkConnectRequestSrcNode
	SetStartTime(v int64) *FindContainerNetworkConnectRequest
	GetStartTime() *int64
}

type FindContainerNetworkConnectRequest struct {
	// The type of the information that you want to query. Valid values:
	//
	// 	- **EDGE**: connection information
	//
	// example:
	//
	// EDGE
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The information about the destination node.
	DstNode *FindContainerNetworkConnectRequestDstNode `json:"DstNode,omitempty" xml:"DstNode,omitempty" type:"Struct"`
	// The end time of the network connection.
	//
	// example:
	//
	// 1649260799999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the source node.
	SrcNode *FindContainerNetworkConnectRequestSrcNode `json:"SrcNode,omitempty" xml:"SrcNode,omitempty" type:"Struct"`
	// The start time of the network connection.
	//
	// example:
	//
	// 1666886400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s FindContainerNetworkConnectRequest) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectRequest) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *FindContainerNetworkConnectRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *FindContainerNetworkConnectRequest) GetDstNode() *FindContainerNetworkConnectRequestDstNode {
	return s.DstNode
}

func (s *FindContainerNetworkConnectRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *FindContainerNetworkConnectRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *FindContainerNetworkConnectRequest) GetSrcNode() *FindContainerNetworkConnectRequestSrcNode {
	return s.SrcNode
}

func (s *FindContainerNetworkConnectRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *FindContainerNetworkConnectRequest) SetCriteriaType(v string) *FindContainerNetworkConnectRequest {
	s.CriteriaType = &v
	return s
}

func (s *FindContainerNetworkConnectRequest) SetCurrentPage(v int64) *FindContainerNetworkConnectRequest {
	s.CurrentPage = &v
	return s
}

func (s *FindContainerNetworkConnectRequest) SetDstNode(v *FindContainerNetworkConnectRequestDstNode) *FindContainerNetworkConnectRequest {
	s.DstNode = v
	return s
}

func (s *FindContainerNetworkConnectRequest) SetEndTime(v int64) *FindContainerNetworkConnectRequest {
	s.EndTime = &v
	return s
}

func (s *FindContainerNetworkConnectRequest) SetPageSize(v int64) *FindContainerNetworkConnectRequest {
	s.PageSize = &v
	return s
}

func (s *FindContainerNetworkConnectRequest) SetSrcNode(v *FindContainerNetworkConnectRequestSrcNode) *FindContainerNetworkConnectRequest {
	s.SrcNode = v
	return s
}

func (s *FindContainerNetworkConnectRequest) SetStartTime(v int64) *FindContainerNetworkConnectRequest {
	s.StartTime = &v
	return s
}

func (s *FindContainerNetworkConnectRequest) Validate() error {
	return dara.Validate(s)
}

type FindContainerNetworkConnectRequestDstNode struct {
	// The name of the container application.
	//
	// example:
	//
	// nginx-ingress-controller
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// example:
	//
	// f5x833820xx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespace of the cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The node IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The type of the node. Valid values:
	//
	// 	- **app**: application, which indicates that the node type is application.
	//
	// example:
	//
	// app
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The name of the pod.
	//
	// example:
	//
	// abc-deployment-yacs-31144-39265-1384966-7f8c8cd578-h6mhb
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s FindContainerNetworkConnectRequestDstNode) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectRequestDstNode) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectRequestDstNode) GetAppName() *string {
	return s.AppName
}

func (s *FindContainerNetworkConnectRequestDstNode) GetClusterId() *string {
	return s.ClusterId
}

func (s *FindContainerNetworkConnectRequestDstNode) GetNamespace() *string {
	return s.Namespace
}

func (s *FindContainerNetworkConnectRequestDstNode) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *FindContainerNetworkConnectRequestDstNode) GetNodeType() *string {
	return s.NodeType
}

func (s *FindContainerNetworkConnectRequestDstNode) GetPodName() *string {
	return s.PodName
}

func (s *FindContainerNetworkConnectRequestDstNode) SetAppName(v string) *FindContainerNetworkConnectRequestDstNode {
	s.AppName = &v
	return s
}

func (s *FindContainerNetworkConnectRequestDstNode) SetClusterId(v string) *FindContainerNetworkConnectRequestDstNode {
	s.ClusterId = &v
	return s
}

func (s *FindContainerNetworkConnectRequestDstNode) SetNamespace(v string) *FindContainerNetworkConnectRequestDstNode {
	s.Namespace = &v
	return s
}

func (s *FindContainerNetworkConnectRequestDstNode) SetNodeIds(v []*string) *FindContainerNetworkConnectRequestDstNode {
	s.NodeIds = v
	return s
}

func (s *FindContainerNetworkConnectRequestDstNode) SetNodeType(v string) *FindContainerNetworkConnectRequestDstNode {
	s.NodeType = &v
	return s
}

func (s *FindContainerNetworkConnectRequestDstNode) SetPodName(v string) *FindContainerNetworkConnectRequestDstNode {
	s.PodName = &v
	return s
}

func (s *FindContainerNetworkConnectRequestDstNode) Validate() error {
	return dara.Validate(s)
}

type FindContainerNetworkConnectRequestSrcNode struct {
	// The name of the container application.
	//
	// example:
	//
	// arms-prometheus-ack-arms-prometheus
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// example:
	//
	// c56xxx1775dea0
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespace of the cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The node IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The type of the node. Valid values:
	//
	// 	- **app**: application, which indicates that the node type is application.
	//
	// example:
	//
	// app
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The name of the pod.
	//
	// example:
	//
	// abc-deployment-yacs-31144-39265-1384966-7f8c8cd578-h6mhb
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s FindContainerNetworkConnectRequestSrcNode) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectRequestSrcNode) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectRequestSrcNode) GetAppName() *string {
	return s.AppName
}

func (s *FindContainerNetworkConnectRequestSrcNode) GetClusterId() *string {
	return s.ClusterId
}

func (s *FindContainerNetworkConnectRequestSrcNode) GetNamespace() *string {
	return s.Namespace
}

func (s *FindContainerNetworkConnectRequestSrcNode) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *FindContainerNetworkConnectRequestSrcNode) GetNodeType() *string {
	return s.NodeType
}

func (s *FindContainerNetworkConnectRequestSrcNode) GetPodName() *string {
	return s.PodName
}

func (s *FindContainerNetworkConnectRequestSrcNode) SetAppName(v string) *FindContainerNetworkConnectRequestSrcNode {
	s.AppName = &v
	return s
}

func (s *FindContainerNetworkConnectRequestSrcNode) SetClusterId(v string) *FindContainerNetworkConnectRequestSrcNode {
	s.ClusterId = &v
	return s
}

func (s *FindContainerNetworkConnectRequestSrcNode) SetNamespace(v string) *FindContainerNetworkConnectRequestSrcNode {
	s.Namespace = &v
	return s
}

func (s *FindContainerNetworkConnectRequestSrcNode) SetNodeIds(v []*string) *FindContainerNetworkConnectRequestSrcNode {
	s.NodeIds = v
	return s
}

func (s *FindContainerNetworkConnectRequestSrcNode) SetNodeType(v string) *FindContainerNetworkConnectRequestSrcNode {
	s.NodeType = &v
	return s
}

func (s *FindContainerNetworkConnectRequestSrcNode) SetPodName(v string) *FindContainerNetworkConnectRequestSrcNode {
	s.PodName = &v
	return s
}

func (s *FindContainerNetworkConnectRequestSrcNode) Validate() error {
	return dara.Validate(s)
}
