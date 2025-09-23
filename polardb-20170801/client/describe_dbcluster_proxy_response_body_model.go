// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChildInstances(v []*DescribeDBClusterProxyResponseBodyChildInstances) *DescribeDBClusterProxyResponseBody
	GetChildInstances() []*DescribeDBClusterProxyResponseBodyChildInstances
	SetDBProxyClusterId(v string) *DescribeDBClusterProxyResponseBody
	GetDBProxyClusterId() *string
	SetDBProxyClusterNum(v int64) *DescribeDBClusterProxyResponseBody
	GetDBProxyClusterNum() *int64
	SetDBProxyClusterStatus(v string) *DescribeDBClusterProxyResponseBody
	GetDBProxyClusterStatus() *string
	SetRequestId(v string) *DescribeDBClusterProxyResponseBody
	GetRequestId() *string
}

type DescribeDBClusterProxyResponseBody struct {
	ChildInstances []*DescribeDBClusterProxyResponseBodyChildInstances `json:"ChildInstances,omitempty" xml:"ChildInstances,omitempty" type:"Repeated"`
	// example:
	//
	// pe-xxxxxxxxx
	DBProxyClusterId *string `json:"DBProxyClusterId,omitempty" xml:"DBProxyClusterId,omitempty"`
	// example:
	//
	// 2
	DBProxyClusterNum *int64 `json:"DBProxyClusterNum,omitempty" xml:"DBProxyClusterNum,omitempty"`
	// example:
	//
	// ClassChanging
	DBProxyClusterStatus *string `json:"DBProxyClusterStatus,omitempty" xml:"DBProxyClusterStatus,omitempty"`
	// example:
	//
	// 30E11ED2-C922-5B96-BCC6-11EE8C484AC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterProxyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterProxyResponseBody) GetChildInstances() []*DescribeDBClusterProxyResponseBodyChildInstances {
	return s.ChildInstances
}

func (s *DescribeDBClusterProxyResponseBody) GetDBProxyClusterId() *string {
	return s.DBProxyClusterId
}

func (s *DescribeDBClusterProxyResponseBody) GetDBProxyClusterNum() *int64 {
	return s.DBProxyClusterNum
}

func (s *DescribeDBClusterProxyResponseBody) GetDBProxyClusterStatus() *string {
	return s.DBProxyClusterStatus
}

func (s *DescribeDBClusterProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterProxyResponseBody) SetChildInstances(v []*DescribeDBClusterProxyResponseBodyChildInstances) *DescribeDBClusterProxyResponseBody {
	s.ChildInstances = v
	return s
}

func (s *DescribeDBClusterProxyResponseBody) SetDBProxyClusterId(v string) *DescribeDBClusterProxyResponseBody {
	s.DBProxyClusterId = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBody) SetDBProxyClusterNum(v int64) *DescribeDBClusterProxyResponseBody {
	s.DBProxyClusterNum = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBody) SetDBProxyClusterStatus(v string) *DescribeDBClusterProxyResponseBody {
	s.DBProxyClusterStatus = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBody) SetRequestId(v string) *DescribeDBClusterProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterProxyResponseBodyChildInstances struct {
	// example:
	//
	// polar.mysql.g4.medium
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// 10.*.*10
	DBNodeIP *string `json:"DBNodeIP,omitempty" xml:"DBNodeIP,omitempty"`
	// example:
	//
	// pi-wz97h479y364g9du7
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// example:
	//
	// 2450
	DBNodePort *string `json:"DBNodePort,omitempty" xml:"DBNodePort,omitempty"`
	// example:
	//
	// Running
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// example:
	//
	// sh-lsf01-144-37
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
}

func (s DescribeDBClusterProxyResponseBodyChildInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterProxyResponseBodyChildInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) GetDBNodeIP() *string {
	return s.DBNodeIP
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) GetDBNodePort() *string {
	return s.DBNodePort
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) GetDBNodeStatus() *string {
	return s.DBNodeStatus
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) GetHostName() *string {
	return s.HostName
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) SetDBNodeClass(v string) *DescribeDBClusterProxyResponseBodyChildInstances {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) SetDBNodeIP(v string) *DescribeDBClusterProxyResponseBodyChildInstances {
	s.DBNodeIP = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) SetDBNodeId(v string) *DescribeDBClusterProxyResponseBodyChildInstances {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) SetDBNodePort(v string) *DescribeDBClusterProxyResponseBodyChildInstances {
	s.DBNodePort = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) SetDBNodeStatus(v string) *DescribeDBClusterProxyResponseBodyChildInstances {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) SetHostName(v string) *DescribeDBClusterProxyResponseBodyChildInstances {
	s.HostName = &v
	return s
}

func (s *DescribeDBClusterProxyResponseBodyChildInstances) Validate() error {
	return dara.Validate(s)
}
