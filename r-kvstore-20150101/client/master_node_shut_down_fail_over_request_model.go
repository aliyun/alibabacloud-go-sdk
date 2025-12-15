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
	// The resource category. Set the value to instance.
	//
	// example:
	//
	// instance
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 	- Specify: This mode allows you to specify a database node to use.
	//
	// 	- Random: In this mode, a random database node is selected when no database node is specified.
	//
	// example:
	//
	// Random
	DBFaultMode *string `json:"DBFaultMode,omitempty" xml:"DBFaultMode,omitempty"`
	// The IDs of the database nodes.
	//
	// example:
	//
	// r-rdsdavinx01003-db-0,r-rdsdavinx01003-db-1
	DBNodes *string `json:"DBNodes,omitempty" xml:"DBNodes,omitempty"`
	// 	- **Hard**: stimulates a hardware failure that cannot be recovered. In this case, a high-availability switchover is triggered.
	//
	// 	- **Soft*	- (default): stimulates a hardware failure that can be recovered. In this case, the system first attempts to recover the faulty node. If the attempt fails, a high-availability switchover is triggered.
	//
	// Valid values:
	//
	// 	- Safe
	//
	// 	- UnSafe
	//
	// 	- Hard
	//
	// 	- Soft
	//
	// example:
	//
	// Safe
	FailMode *string `json:"FailMode,omitempty" xml:"FailMode,omitempty"`
	// The instance ID. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 	- Specify: This mode allows you to specify a proxy node to use.
	//
	// 	- Random: In this mode, a random proxy node is selected when no proxy node is specified.
	//
	// example:
	//
	// Specify
	ProxyFaultMode *string `json:"ProxyFaultMode,omitempty" xml:"ProxyFaultMode,omitempty"`
	// The IDs of the proxy nodes.
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
