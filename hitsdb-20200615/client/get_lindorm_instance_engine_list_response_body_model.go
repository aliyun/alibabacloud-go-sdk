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
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The list of engine types.
	EngineList []*GetLindormInstanceEngineListResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// ld-bp1nq34mv3smk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
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
	if s.EngineList != nil {
		for _, item := range s.EngineList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormInstanceEngineListResponseBodyEngineList struct {
	// The engine type. Valid values:
	//
	// - **lindorm**: LindormTable.
	//
	// - **tsdb**: LindormTSDB.
	//
	// - **solr**: Search engine.
	//
	// - **store**: File engine.
	//
	// example:
	//
	// lindorm
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The list of database connection information for the engine.
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
	if s.NetInfoList != nil {
		for _, item := range s.NetInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormInstanceEngineListResponseBodyEngineListNetInfoList struct {
	// The connection method for LindormTable. Valid values:
	//
	// - **0**: This is the default value and can be ignored.
	//
	// - **1**: Use the HBase Java API to access LindormTable.
	//
	// - **2**: Use a non-Java HBase API to access LindormTable.
	//
	// - **3**: Use CQL to access LindormTable.
	//
	// - **4**: Use the LindormTable SQL endpoint.
	//
	// - **5**: Use the S3-compatible endpoint for LindormTable.
	//
	// - **6**: Use the MySQL-compatible endpoint for LindormTable.
	//
	// example:
	//
	// 1
	AccessType *int32 `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The database endpoint.
	//
	// example:
	//
	// ld-bp1nq34mv3smk****-proxy-lindorm.lindorm.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The network type of the database endpoint. Valid values:
	//
	// - **0**: Internet.
	//
	// - **2**: Virtual private cloud (VPC).
	//
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port number of the database endpoint.
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
