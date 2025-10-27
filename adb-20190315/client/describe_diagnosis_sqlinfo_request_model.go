// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisSQLInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDiagnosisSQLInfoRequest
	GetDBClusterId() *string
	SetLang(v string) *DescribeDiagnosisSQLInfoRequest
	GetLang() *string
	SetProcessId(v string) *DescribeDiagnosisSQLInfoRequest
	GetProcessId() *string
	SetProcessRcHost(v string) *DescribeDiagnosisSQLInfoRequest
	GetProcessRcHost() *string
	SetProcessStartTime(v int64) *DescribeDiagnosisSQLInfoRequest
	GetProcessStartTime() *int64
	SetProcessState(v string) *DescribeDiagnosisSQLInfoRequest
	GetProcessState() *string
	SetRegionId(v string) *DescribeDiagnosisSQLInfoRequest
	GetRegionId() *string
}

type DescribeDiagnosisSQLInfoRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// 	- **zh**: simplified Chinese
	//
	// 	- **en**: English
	//
	// 	- **ja**: Japanese
	//
	// 	- **zh-tw**: traditional Chinese
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the query.
	//
	// >  You can call the [DescribeDiagnosisRecords](https://help.aliyun.com/document_detail/308207.html) operation to query the SQL summary information of a specified AnalyticDB for MySQL cluster, including the query ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021070216432217201616806503453******
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The IP address and port number of the AnalyticDB for MySQL frontend node on which the SQL statement is executed.
	//
	// >  You can call the [DescribeDiagnosisRecords](https://help.aliyun.com/document_detail/308207.html) operation to query the SQL summary information of a specified AnalyticDB for MySQL cluster, including the IP address and port number of the frontend node.
	//
	// example:
	//
	// 192.45.***.***:3145
	ProcessRcHost *string `json:"ProcessRcHost,omitempty" xml:"ProcessRcHost,omitempty"`
	// The execution start time of the SQL statement. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// >  You can call the [DescribeDiagnosisRecords](https://help.aliyun.com/document_detail/308207.html) operation to query the SQL summary information of a specified AnalyticDB for MySQL cluster, including the execution start time of the SQL statement.
	//
	// example:
	//
	// 1625215402000
	ProcessStartTime *int64 `json:"ProcessStartTime,omitempty" xml:"ProcessStartTime,omitempty"`
	// The state of the SQL statement. Valid values:
	//
	// 	- **running**
	//
	// 	- **finished**
	//
	// 	- **failed**
	//
	// > You can call the [DescribeDiagnosisRecords](https://help.aliyun.com/document_detail/308207.html) operation to query the SQL summary information of a specified AnalyticDB for MySQL cluster, including the state of the SQL statement.
	//
	// example:
	//
	// running
	ProcessState *string `json:"ProcessState,omitempty" xml:"ProcessState,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDiagnosisSQLInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDiagnosisSQLInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDiagnosisSQLInfoRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeDiagnosisSQLInfoRequest) GetProcessRcHost() *string {
	return s.ProcessRcHost
}

func (s *DescribeDiagnosisSQLInfoRequest) GetProcessStartTime() *int64 {
	return s.ProcessStartTime
}

func (s *DescribeDiagnosisSQLInfoRequest) GetProcessState() *string {
	return s.ProcessState
}

func (s *DescribeDiagnosisSQLInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosisSQLInfoRequest) SetDBClusterId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetLang(v string) *DescribeDiagnosisSQLInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessRcHost(v string) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessRcHost = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessStartTime(v int64) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessStartTime = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessState(v string) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessState = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetRegionId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) Validate() error {
	return dara.Validate(s)
}
