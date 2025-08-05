// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*DescribeContainerClusterResponseBodyClusters) *DescribeContainerClusterResponseBody
	GetClusters() []*DescribeContainerClusterResponseBodyClusters
	SetCode(v string) *DescribeContainerClusterResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeContainerClusterResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeContainerClusterResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeContainerClusterResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeContainerClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeContainerClusterResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeContainerClusterResponseBody
	GetTotalCount() *int64
}

type DescribeContainerClusterResponseBody struct {
	// The information of clusters.
	Clusters []*DescribeContainerClusterResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CC94B755-C3C2-5B9D-BD77-E0FE819A4DB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerClusterResponseBody) GetClusters() []*DescribeContainerClusterResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeContainerClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeContainerClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeContainerClusterResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeContainerClusterResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContainerClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeContainerClusterResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeContainerClusterResponseBody) SetClusters(v []*DescribeContainerClusterResponseBodyClusters) *DescribeContainerClusterResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetCode(v string) *DescribeContainerClusterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetMessage(v string) *DescribeContainerClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetPageNumber(v int32) *DescribeContainerClusterResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetPageSize(v int32) *DescribeContainerClusterResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetRequestId(v string) *DescribeContainerClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetSuccess(v bool) *DescribeContainerClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetTotalCount(v int64) *DescribeContainerClusterResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeContainerClusterResponseBodyClusters struct {
	// The status of the client. Valid values:
	//
	// 	- **MISS**: The client is disconnected.
	//
	// 	- **UNKNOWN**: The client is in an unknown state.
	//
	// 	- **READY**: The client is ready.
	//
	// example:
	//
	// READY
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// cl-0006gwppd0jtttpmb0ri
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the cluster. Valid value: ACK, which indicates ACK clusters.
	//
	// example:
	//
	// ACK
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The description.
	//
	// example:
	//
	// description ack pv backup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the cluster.
	//
	// example:
	//
	// c5bbd0931a30947f4ab85efd19380a72d
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// ack_pv_backup_location
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// 	- **CLASSIC**: the classic network
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The token that is used to register the Hybrid Backup Recovery (HBR) client in the cluster.
	//
	// example:
	//
	// eyJhY2Nvd******A/VnZpgXQC5A==
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeContainerClusterResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerClusterResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeContainerClusterResponseBodyClusters) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *DescribeContainerClusterResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerClusterResponseBodyClusters) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeContainerClusterResponseBodyClusters) GetDescription() *string {
	return s.Description
}

func (s *DescribeContainerClusterResponseBodyClusters) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeContainerClusterResponseBodyClusters) GetName() *string {
	return s.Name
}

func (s *DescribeContainerClusterResponseBodyClusters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeContainerClusterResponseBodyClusters) GetToken() *string {
	return s.Token
}

func (s *DescribeContainerClusterResponseBodyClusters) SetAgentStatus(v string) *DescribeContainerClusterResponseBodyClusters {
	s.AgentStatus = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetClusterId(v string) *DescribeContainerClusterResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetClusterType(v string) *DescribeContainerClusterResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetDescription(v string) *DescribeContainerClusterResponseBodyClusters {
	s.Description = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetIdentifier(v string) *DescribeContainerClusterResponseBodyClusters {
	s.Identifier = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetName(v string) *DescribeContainerClusterResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetNetworkType(v string) *DescribeContainerClusterResponseBodyClusters {
	s.NetworkType = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetToken(v string) *DescribeContainerClusterResponseBodyClusters {
	s.Token = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) Validate() error {
	return dara.Validate(s)
}
