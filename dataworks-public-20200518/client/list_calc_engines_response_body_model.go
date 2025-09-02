// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalcEnginesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCalcEnginesResponseBodyData) *ListCalcEnginesResponseBody
	GetData() *ListCalcEnginesResponseBodyData
	SetHttpStatusCode(v int32) *ListCalcEnginesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListCalcEnginesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCalcEnginesResponseBody
	GetSuccess() *bool
}

type ListCalcEnginesResponseBody struct {
	// The query results for compute engines that are returned on multiple pages.
	Data *ListCalcEnginesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCalcEnginesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCalcEnginesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalcEnginesResponseBody) GetData() *ListCalcEnginesResponseBodyData {
	return s.Data
}

func (s *ListCalcEnginesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCalcEnginesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCalcEnginesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCalcEnginesResponseBody) SetData(v *ListCalcEnginesResponseBodyData) *ListCalcEnginesResponseBody {
	s.Data = v
	return s
}

func (s *ListCalcEnginesResponseBody) SetHttpStatusCode(v int32) *ListCalcEnginesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCalcEnginesResponseBody) SetRequestId(v string) *ListCalcEnginesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCalcEnginesResponseBody) SetSuccess(v bool) *ListCalcEnginesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCalcEnginesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCalcEnginesResponseBodyData struct {
	// The compute engines.
	CalcEngines []*ListCalcEnginesResponseBodyDataCalcEngines `json:"CalcEngines,omitempty" xml:"CalcEngines,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of compute engine instances.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCalcEnginesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCalcEnginesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCalcEnginesResponseBodyData) GetCalcEngines() []*ListCalcEnginesResponseBodyDataCalcEngines {
	return s.CalcEngines
}

func (s *ListCalcEnginesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCalcEnginesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCalcEnginesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCalcEnginesResponseBodyData) SetCalcEngines(v []*ListCalcEnginesResponseBodyDataCalcEngines) *ListCalcEnginesResponseBodyData {
	s.CalcEngines = v
	return s
}

func (s *ListCalcEnginesResponseBodyData) SetPageNumber(v int32) *ListCalcEnginesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCalcEnginesResponseBodyData) SetPageSize(v int32) *ListCalcEnginesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCalcEnginesResponseBodyData) SetTotalCount(v int32) *ListCalcEnginesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCalcEnginesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListCalcEnginesResponseBodyDataCalcEngines struct {
	// The ID of the workspace with which the compute engine is associated.
	//
	// example:
	//
	// 27
	BindingProjectId *int32 `json:"BindingProjectId,omitempty" xml:"BindingProjectId,omitempty"`
	// The name of the workspace with which the compute engine is associated.
	//
	// example:
	//
	// abc
	BindingProjectName *string `json:"BindingProjectName,omitempty" xml:"BindingProjectName,omitempty"`
	// The type of the compute engine.
	//
	// example:
	//
	// ODPS
	CalcEngineType *string `json:"CalcEngineType,omitempty" xml:"CalcEngineType,omitempty"`
	// The region where the DataWorks workspace with which the compute engine is associated resides.
	//
	// example:
	//
	// cn-zhangjiakou
	DwRegion *string `json:"DwRegion,omitempty" xml:"DwRegion,omitempty"`
	// The compute engine ID.
	//
	// example:
	//
	// 35
	EngineId *int32 `json:"EngineId,omitempty" xml:"EngineId,omitempty"`
	// The details of the compute engine.
	//
	// 	- ODPS
	//
	//         {
	//
	//           "pubEndpoint": "service.cn.maxcompute.aliyun.com/api",
	//
	//           "endpoint": "service.cn.maxcompute.aliyun-inc.com/api",
	//
	//           "initProperties": "
	//
	//           {\\"odpsTypeMode\\":\\"STANDARD\\",\\"openPai\\":false,\\"openPaiGpu\\":false}",
	//
	//           "resourceGroupType": "ODPS",
	//
	//           "resourceGroupId": "361826516****",
	//
	//           "vpcEndpoint": "service.cn.maxcompute.aliyun-inc.com/api",
	//
	//           "projectName": "onefall_test_zjk",
	//
	//           "taskSameAsOwner": "true"
	//
	//         }
	//
	// 	- EMR
	//
	//         {
	//
	//           "emrClusterId": "C-xxx",
	//
	//           "specs": "{\\"emrClusterId\\":\\"C-xxx\\",\\"emrAccessMode\\":\\"simple\\",\\"emrResourceQueueName\\":\\"default\\",\\"emrProjectId\\":\\"FP-xxx\\"}",
	//
	//           "endpoint": "emr.aliyuncs.com",
	//
	//           "emrResourceQueueName": "default",
	//
	//           "emrAccessMode": "simple",
	//
	//           "resourceGroupType": "DW",
	//
	//           "projectName": "xx-xxxx",
	//
	//           "emrProjectId": "FP-xxxx",
	//
	//           "taskSameAsOwner": "false"
	//
	//         }
	//
	// 	- BLINK
	//
	//         {
	//
	//           "bayesProjectId": "xxxx",
	//
	//           "bayesProjectName": "xc_blxxixxxnk_1",
	//
	//           "cluster": "xxxssxsx",
	//
	//           "endpoint": "https://stream.console.aliyun.com",
	//
	//           "engineType": "BLINK",
	//
	//           "name": "xsxsxxxxx",
	//
	//           "projectName": "xc_blxxxsxink_1",
	//
	//           "queue": "root.xc_blxsxxxxxxink_1",
	//
	//           "resourceGroupType": "DW",
	//
	//           "specs": "{\\"cluster\\":\\"xxxxxx\\",\\"bayesProjectName\\":\\"xc_blxxixxxnk_1\\",\\"bayesProjectId\\":\\"ssxxxsa\\",\\"name\\":\\"sxsxsxxx\\",\\"queue\\":\\"root.sxxsxxsx\\"}",
	//
	//           "taskSameAsOwner": false
	//
	//         }
	//
	// 	- HOLO
	//
	//         {
	//
	//           "endpoint": "hgprecn-cn-xsxssxsx-cn-shanghai-internal.hologres.aliyuncs.com:80",
	//
	//           "engineType": "ODPS",
	//
	//           "odpsEndpoint": "hgprecn-cn-xsxssxxs-cn-shanghai-internal.hologres.aliyuncs.com:80",
	//
	//           "odpsProjectName": "xsxssxsxsx",
	//
	//           "projectName": "xsxssxsxsx",
	//
	//           "resourceGroupType": "DW",
	//
	//           "specs": "{\\"pubEndpoint\\":\\"hgprecn-cn-xsxssxsxs-cn-shanghai.hologres.aliyuncs.com:80\\",\\"commonBuyInstanceId\\":\\"hgprecn-cn-xsxsxsxs\\",\\"project\\":\\"holo_upxsxgrade1\\",\\"common_buy_instance_id\\":\\"hgprecn-cn-xsxsxs\\",\\"endpoint\\":\\"hgprecn-cn-xsxxsxs-cn-shanghai-internal.hologres.aliyuncs.com:80\\",\\"port\\":\\"80\\",\\"host\\":\\"hgprecn-cn-xsxsxsxs-cn-shanghai-internal.hologres.aliyuncs.com\\",\\"vpcEndpoint\\":\\"hgprecn-cn-xsxsxsxs-cn-shanghai-vpc.hologres.aliyuncs.com:80\\",\\"authType\\":2,\\"region\\":\\"cn-shanghai\\"}",
	//
	//           "taskSameAsOwner": false
	//
	//         }
	//
	// 	- MaxGraph
	//
	//         {
	//
	//           "endpoint": "http://pre-graphcompute.aliyuncs.com",
	//
	//           "engineType": "ODPS",
	//
	//           "odpsEndpoint": "http://pre-graphcompute.aliyuncs.com",
	//
	//           "odpsProjectName": "xsxsxsxs",
	//
	//           "projectName": "xsxsxsxs",
	//
	//           "resourceGroupType": "DW",
	//
	//           "taskSameAsOwner": false
	//
	//         }
	//
	// 	- HYBRIDDB_FOR_POSTGRESQL
	//
	//         {
	//
	//           "endpoint": "hybriddb_for_postgresql_mo12121ck_endpoint",
	//
	//           "engineType": "ODPS",
	//
	//           "odpsEndpoint": "hybriddb_for_postgresql_m121212ock_endpoint",
	//
	//           "odpsProjectName": "sxasaxsaxaxas",
	//
	//           "projectName": "sxasaxsaxaxas",
	//
	//           "resourceGroupType": "DW",
	//
	//           "specs": "{\\"connectionString\\":\\"gp-xsxsxsxxs.gpdb.rds.aliyuncs.com\\",\\"database\\":\\"xsxsxxsxs\\",\\"password\\":\\"xxxxxxx\\",\\"instanceId\\":\\"gp-cdcdacdacda\\",\\"port\\":\\"3432\\",\\"ownerId\\":\\"12121212\\",\\"username\\":\\"sdasaddsa\\"}",
	//
	//           "taskSameAsOwner": false
	//
	//         }
	//
	// 	- ADB_MYSQL
	//
	//         {
	//
	//           "endpoint": "adb_mysql_mock_endpoint",
	//
	//           "engineType": "ODPS",
	//
	//           "odpsEndpoint": "adb_mysql_mock_endpoint",
	//
	//           "odpsProjectName": "am-xsaxaxa",
	//
	//           "projectName": "am-xsxsaxa",
	//
	//           "resourceGroupType": "DW",
	//
	//           "specs": "{\\"connectionString\\":\\"am-xsaxsa.ads.aliyuncs.com:3306\\",\\"database\\":\\"xsaxsaxa\\",\\"password\\":\\"xsaxsaxassxsa\\",\\"instanceId\\":\\"am-xsaxsasx\\",\\"username\\":\\"xsaxsadsd\\"}",
	//
	//           "taskSameAsOwner": false
	//
	//         }
	//
	// 	- HADOOP_CDH
	//
	//         {
	//
	//           "bindingBaseId": "xsaxsaxs",
	//
	//           "endpoint": "xsaaaaa",
	//
	//           "engineType": "ODPS",
	//
	//           "odpsEndpoint": "axsxaxssxs",
	//
	//           "odpsProjectName": "ssxxax",
	//
	//           "projectName": "xsaxsaxsa",
	//
	//           "resourceGroupId": 45208xxxxxx,
	//
	//           "resourceGroupType": "GATEWAY",
	//
	//           "specs": "{\\"cluster\\":{\\"hive\\":{\\"hiveServer2Url\\":\\"jdbc:hive2://xxxxxxer-1-cn-shanghai-pre-kerberos-1:10000\\",\\"hiveMetastore\\":\\"thrift://xxxxxxxr-1-cn-shanghai-pre-kerberos-1:9083\\",\\"version\\":\\"2.1.1\\"},\\"configFiles\\":{\\"coreSite\\":\\"4534574xxxxxx\\",\\"hdfsSite\\":\\"453457919xxxxxxx\\",\\"mapredSite\\":\\"45345750xxxxxx\\",\\"yarnSite\\":\\"4534575xxxxx\\",\\"krb5Conf\\":\\"4534576xxxxx1\\",\\"hiveSite\\":\\"453457xxxxx20\\"},\\"spark\\":{\\"version\\":\\"2.4.0\\"},\\"cdh\\":{\\"version\\":\\"6.3.2\\"},\\"hdfs\\":{\\"version\\":\\"3.0.0\\"},\\"impala\\":{\\"impalaUrl\\":\\"jdbc:impala://cdh-xsxssxxsx-1-cn-shanghai-pre-kerberos-1:21050\\",\\"version\\":\\"3.2.0\\"},\\"yarn\\":{\\"YarnUrl\\":\\"http://cdh-xsxsxsxsxs-1-cn-shanghai-pre-kerberos-1:8032\\",\\"webUrl\\":\\"http://cdh-xsxsxssxxssx-1-cn-shanghai-pre-kerberos-1:8088\\",\\"version\\":\\"3.0.0\\"},\\"presto\\":{\\"prestoUrl\\":\\"jdbc:presto://cdh-xssxsxxsxsxs-1-cn-shanghai-pre-kerberos-1:8080/hive/default\\",\\"version\\":\\"0.244.1\\"}},\\"instanceId\\":161sdads733,\\"authDetail\\":{\\"principal\\":\\"hive@HADOOP.COM\\",\\"keytabFileId\\":\\"45345815xsxsxs3\\",\\"type\\":\\"kerberos\\",\\"username\\":\\"xsxsxsxsa@HADOOP.COM\\"},\\"resGroupStatus\\":\\"\\",\\"hadoopAuthType\\":\\"kerberos\\",\\"clusterIdentifier\\":\\"xssxsxsxsx\\",\\"clusterId\\":xsxsx,\\"resGroupId\\":4520870619xsxssxxs,\\"accessMode\\":\\"security\\",\\"authType\\":2}",
	//
	//           "taskSameAsOwner": false
	//
	//         }
	//
	// example:
	//
	// {"pubEndpoint":"http://service.cn.maxcompute.aliyun.com/api","endpoint":"http://service.cn.maxcompute.aliyun-inc.com/api","resourceGroupType":"ODPS","resourceGroupId":"361826516****","vpcEndpoint":"http://service.cn.maxcompute.aliyun-inc.com/api","projectName":"onefall_test_zjk","taskSameAsOwner":"true"}
	EngineInfo map[string]interface{} `json:"EngineInfo,omitempty" xml:"EngineInfo,omitempty"`
	// The environment in which the compute engine is used. Valid values:
	//
	// 	- **DEV**
	//
	// 	- **PRD**
	//
	// example:
	//
	// PRD
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The time when the compute engine was created.
	//
	// example:
	//
	// Oct 10, 2019 3:42:44 PM
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Indicates whether the compute engine is the default engine of the current type.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The display name of the compute engine.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region where the compute engine resides.
	//
	// example:
	//
	// cn-zhangjiakou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The identity that is used to access the compute engine. Valid values:
	//
	// 	- **USER**: the current user
	//
	// 	- **PROJECT**: the workspace executor
	//
	// 	- **SUBACCOUNT**: a RAM user
	//
	// 	- **STS_ROLE**: the Security Token Service (STS) role
	//
	// example:
	//
	// PROJECT
	TaskAuthType *string `json:"TaskAuthType,omitempty" xml:"TaskAuthType,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListCalcEnginesResponseBodyDataCalcEngines) String() string {
	return dara.Prettify(s)
}

func (s ListCalcEnginesResponseBodyDataCalcEngines) GoString() string {
	return s.String()
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetBindingProjectId() *int32 {
	return s.BindingProjectId
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetBindingProjectName() *string {
	return s.BindingProjectName
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetCalcEngineType() *string {
	return s.CalcEngineType
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetDwRegion() *string {
	return s.DwRegion
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetEngineId() *int32 {
	return s.EngineId
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetEngineInfo() map[string]interface{} {
	return s.EngineInfo
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetEnvType() *string {
	return s.EnvType
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetName() *string {
	return s.Name
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetRegion() *string {
	return s.Region
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetTaskAuthType() *string {
	return s.TaskAuthType
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetBindingProjectId(v int32) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.BindingProjectId = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetBindingProjectName(v string) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.BindingProjectName = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetCalcEngineType(v string) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.CalcEngineType = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetDwRegion(v string) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.DwRegion = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetEngineId(v int32) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.EngineId = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetEngineInfo(v map[string]interface{}) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.EngineInfo = v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetEnvType(v string) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.EnvType = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetGmtCreate(v string) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.GmtCreate = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetIsDefault(v bool) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.IsDefault = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetName(v string) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.Name = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetRegion(v string) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.Region = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetTaskAuthType(v string) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.TaskAuthType = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) SetTenantId(v int64) *ListCalcEnginesResponseBodyDataCalcEngines {
	s.TenantId = &v
	return s
}

func (s *ListCalcEnginesResponseBodyDataCalcEngines) Validate() error {
	return dara.Validate(s)
}
