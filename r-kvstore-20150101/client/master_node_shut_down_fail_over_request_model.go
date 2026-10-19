// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMasterNodeShutDownFailOverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *MasterNodeShutDownFailOverRequest
	GetCategory() *string
	SetDBFaultMode(v string) *MasterNodeShutDownFailOverRequest
	GetDBFaultMode() *string
	SetDBNodes(v string) *MasterNodeShutDownFailOverRequest
	GetDBNodes() *string
	SetFailMode(v string) *MasterNodeShutDownFailOverRequest
	GetFailMode() *string
	SetInstanceId(v string) *MasterNodeShutDownFailOverRequest
	GetInstanceId() *string
	SetProxyFaultMode(v string) *MasterNodeShutDownFailOverRequest
	GetProxyFaultMode() *string
	SetProxyInstanceIds(v string) *MasterNodeShutDownFailOverRequest
	GetProxyInstanceIds() *string
}

type MasterNodeShutDownFailOverRequest struct {
	// The resource type. Set the value to `instance`.
	//
	// example:
	//
	// instance
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// - **Specify**: Specifies the data nodes.
	//
	// - **Random**: The system randomly simulates a failure for one-fourth of the data nodes (rounded up to the nearest integer, with a minimum of 1). This parameter takes effect only when no data nodes are specified.
	//
	// example:
	//
	// Random
	DBFaultMode *string `json:"DBFaultMode,omitempty" xml:"DBFaultMode,omitempty"`
	// The specified data nodes.
	//
	// example:
	//
	// r-rdsdavinx01003-db-0,r-rdsdavinx01003-db-1
	DBNodes *string `json:"DBNodes,omitempty" xml:"DBNodes,omitempty"`
	// - **Hard**: Simulates an unrecoverable hardware failure. A switchover (HA) is directly triggered.
	//
	// - **Soft*	- (default): Simulates a recoverable hardware failure. The system preferentially attempts to recover the failed node. If the recovery fails, a switchover (HA) is triggered.
	//
	// example:
	//
	// Hard
	FailMode *string `json:"FailMode,omitempty" xml:"FailMode,omitempty"`
	// The instance ID. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// - **Specify**: Specifies the proxy nodes.
	//
	// - **Random**: The system randomly simulates a failure for one-fourth of the proxy nodes (rounded up to the nearest integer, with a minimum of 1). This parameter takes effect only when no proxy nodes are specified.
	//
	// example:
	//
	// Specify
	ProxyFaultMode *string `json:"ProxyFaultMode,omitempty" xml:"ProxyFaultMode,omitempty"`
	// The specified proxy nodes.
	//
	// example:
	//
	// 6981,6982
	ProxyInstanceIds *string `json:"ProxyInstanceIds,omitempty" xml:"ProxyInstanceIds,omitempty"`
}

func (s MasterNodeShutDownFailOverRequest) String() string {
	return dara.Prettify(s)
}

func (s MasterNodeShutDownFailOverRequest) GoString() string {
	return s.String()
}

func (s *MasterNodeShutDownFailOverRequest) GetCategory() *string {
	return s.Category
}

func (s *MasterNodeShutDownFailOverRequest) GetDBFaultMode() *string {
	return s.DBFaultMode
}

func (s *MasterNodeShutDownFailOverRequest) GetDBNodes() *string {
	return s.DBNodes
}

func (s *MasterNodeShutDownFailOverRequest) GetFailMode() *string {
	return s.FailMode
}

func (s *MasterNodeShutDownFailOverRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MasterNodeShutDownFailOverRequest) GetProxyFaultMode() *string {
	return s.ProxyFaultMode
}

func (s *MasterNodeShutDownFailOverRequest) GetProxyInstanceIds() *string {
	return s.ProxyInstanceIds
}

func (s *MasterNodeShutDownFailOverRequest) SetCategory(v string) *MasterNodeShutDownFailOverRequest {
	s.Category = &v
	return s
}

func (s *MasterNodeShutDownFailOverRequest) SetDBFaultMode(v string) *MasterNodeShutDownFailOverRequest {
	s.DBFaultMode = &v
	return s
}

func (s *MasterNodeShutDownFailOverRequest) SetDBNodes(v string) *MasterNodeShutDownFailOverRequest {
	s.DBNodes = &v
	return s
}

func (s *MasterNodeShutDownFailOverRequest) SetFailMode(v string) *MasterNodeShutDownFailOverRequest {
	s.FailMode = &v
	return s
}

func (s *MasterNodeShutDownFailOverRequest) SetInstanceId(v string) *MasterNodeShutDownFailOverRequest {
	s.InstanceId = &v
	return s
}

func (s *MasterNodeShutDownFailOverRequest) SetProxyFaultMode(v string) *MasterNodeShutDownFailOverRequest {
	s.ProxyFaultMode = &v
	return s
}

func (s *MasterNodeShutDownFailOverRequest) SetProxyInstanceIds(v string) *MasterNodeShutDownFailOverRequest {
	s.ProxyInstanceIds = &v
	return s
}

func (s *MasterNodeShutDownFailOverRequest) Validate() error {
	return dara.Validate(s)
}
