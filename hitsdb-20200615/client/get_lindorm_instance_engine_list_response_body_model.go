// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormInstanceEngineListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLindormInstanceEngineListResponseBody
	GetAccessDeniedDetail() *string
	SetEngineList(v []*GetLindormInstanceEngineListResponseBodyEngineList) *GetLindormInstanceEngineListResponseBody
	GetEngineList() []*GetLindormInstanceEngineListResponseBodyEngineList
	SetInstanceId(v string) *GetLindormInstanceEngineListResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GetLindormInstanceEngineListResponseBody
	GetRequestId() *string
}

type GetLindormInstanceEngineListResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The list of engines that can run on the specified instance.
	EngineList []*GetLindormInstanceEngineListResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// Instance ID.
	//
	// example:
	//
	// ld-bp1nq34mv3smk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B496BA0E-520C-59FC-BA04-196D8F3B07EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLindormInstanceEngineListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceEngineListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLindormInstanceEngineListResponseBody) GetEngineList() []*GetLindormInstanceEngineListResponseBodyEngineList {
	return s.EngineList
}

func (s *GetLindormInstanceEngineListResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormInstanceEngineListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormInstanceEngineListResponseBody) SetAccessDeniedDetail(v string) *GetLindormInstanceEngineListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBody) SetEngineList(v []*GetLindormInstanceEngineListResponseBodyEngineList) *GetLindormInstanceEngineListResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormInstanceEngineListResponseBody) SetInstanceId(v string) *GetLindormInstanceEngineListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBody) SetRequestId(v string) *GetLindormInstanceEngineListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLindormInstanceEngineListResponseBodyEngineList struct {
	// The type of engine that can run on the instance. Valid values:
	//
	// 	- **lindorm**: LindormTable.
	//
	// 	- **tsdb**: LindormTSDB.
	//
	// 	- **solr**: LindormSearch.
	//
	// 	- **store**: LindormDFS.
	//
	// example:
	//
	// lindorm
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The list of connection information about the engine.
	NetInfoList []*GetLindormInstanceEngineListResponseBodyEngineListNetInfoList `json:"NetInfoList,omitempty" xml:"NetInfoList,omitempty" type:"Repeated"`
}

func (s GetLindormInstanceEngineListResponseBodyEngineList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceEngineListResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponseBodyEngineList) GetEngineType() *string {
	return s.EngineType
}

func (s *GetLindormInstanceEngineListResponseBodyEngineList) GetNetInfoList() []*GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	return s.NetInfoList
}

func (s *GetLindormInstanceEngineListResponseBodyEngineList) SetEngineType(v string) *GetLindormInstanceEngineListResponseBodyEngineList {
	s.EngineType = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineList) SetNetInfoList(v []*GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) *GetLindormInstanceEngineListResponseBodyEngineList {
	s.NetInfoList = v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineList) Validate() error {
	return dara.Validate(s)
}

type GetLindormInstanceEngineListResponseBodyEngineListNetInfoList struct {
	// The method by which the connection information can be used to access LindormTable. Valid values:
	//
	// 	- **0**: The default value. This value can be ignored.
	//
	// 	- **1**: The connection information can be used to access LindormTable by using ApsaraDB for HBase API for Java.
	//
	// 	- **2**: The connection information can be used to access LindormTable by using ApsaraDB for HBase API for a non-Java language.
	//
	// 	- **3**: The connection information can be used to access LindormTable by using the LindormTable endpoint for CQL.
	//
	// 	- **4**: The connection information can be used to access LindormTable by using the LindormTable endpoint for SQL.
	//
	// 	- **5**: The connection information can be used to access Lindorm by using the LindormTable endpoint for Amazon S3.
	//
	// 	- **6**: The connection information can be used to access Lindorm by using the LindormTable endpoint for MySQL.
	//
	// example:
	//
	// 1
	AccessType *int32 `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The endpoint that is used to connect to the engine.
	//
	// example:
	//
	// ld-bp1nq34mv3smk****-proxy-lindorm.lindorm.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// 	- **0**: Internet
	//
	// 	- **2**: virtual private cloud (VPC)
	//
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port number used to connect to the engine.
	//
	// example:
	//
	// 30020
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) GetAccessType() *int32 {
	return s.AccessType
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) GetNetType() *string {
	return s.NetType
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) GetPort() *int32 {
	return s.Port
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetAccessType(v int32) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.AccessType = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetConnectionString(v string) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.ConnectionString = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetNetType(v string) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.NetType = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetPort(v int32) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.Port = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) Validate() error {
	return dara.Validate(s)
}
