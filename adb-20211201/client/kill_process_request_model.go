// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *KillProcessRequest
	GetDBClusterId() *string
	SetProcessId(v string) *KillProcessRequest
	GetProcessId() *string
	SetRegionId(v string) *KillProcessRequest
	GetRegionId() *string
}

type KillProcessRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-xxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The query ID.
	//
	// >  You can call the [DescribeProcessList](https://help.aliyun.com/document_detail/612277.html) operation to query the IDs of queries that are being executed.
	//
	// example:
	//
	// 202011191048151921681492420315100****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s KillProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s KillProcessRequest) GoString() string {
	return s.String()
}

func (s *KillProcessRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *KillProcessRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *KillProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KillProcessRequest) SetDBClusterId(v string) *KillProcessRequest {
	s.DBClusterId = &v
	return s
}

func (s *KillProcessRequest) SetProcessId(v string) *KillProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *KillProcessRequest) SetRegionId(v string) *KillProcessRequest {
	s.RegionId = &v
	return s
}

func (s *KillProcessRequest) Validate() error {
	return dara.Validate(s)
}
