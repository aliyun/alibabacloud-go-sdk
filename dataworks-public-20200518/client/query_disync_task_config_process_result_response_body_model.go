// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDISyncTaskConfigProcessResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryDISyncTaskConfigProcessResultResponseBodyData) *QueryDISyncTaskConfigProcessResultResponseBody
	GetData() *QueryDISyncTaskConfigProcessResultResponseBodyData
	SetRequestId(v string) *QueryDISyncTaskConfigProcessResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDISyncTaskConfigProcessResultResponseBody
	GetSuccess() *bool
}

type QueryDISyncTaskConfigProcessResultResponseBody struct {
	// The information returned for the parameters that are asynchronously generated and used to create or update a real-time synchronization task in Data Integration.
	Data *QueryDISyncTaskConfigProcessResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0bc1411515937635973****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDISyncTaskConfigProcessResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDISyncTaskConfigProcessResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDISyncTaskConfigProcessResultResponseBody) GetData() *QueryDISyncTaskConfigProcessResultResponseBodyData {
	return s.Data
}

func (s *QueryDISyncTaskConfigProcessResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDISyncTaskConfigProcessResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDISyncTaskConfigProcessResultResponseBody) SetData(v *QueryDISyncTaskConfigProcessResultResponseBodyData) *QueryDISyncTaskConfigProcessResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultResponseBody) SetRequestId(v string) *QueryDISyncTaskConfigProcessResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultResponseBody) SetSuccess(v bool) *QueryDISyncTaskConfigProcessResultResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDISyncTaskConfigProcessResultResponseBodyData struct {
	// The reason why the parameters fail to be obtained. If the parameters are obtained, the value null is returned.
	//
	// example:
	//
	// fileId:[100] is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the parameters are obtained. Valid values:
	//
	// 	- success: The parameters are obtained.
	//
	// 	- fail: The parameters fail to be obtained. You can view the reason for the failure and troubleshoot the issue based on the reason.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The parameters that are obtained. The parameters are used as the request parameters of the [CreateDISyncTask](https://help.aliyun.com/document_detail/278725.html) or [UpdateDISyncTask](https://help.aliyun.com/document_detail/289109.html) operation to create or update a real-time synchronization task in Data Integration.
	//
	// example:
	//
	// {"extend":{"mode":"migration_holo","resourceGroup":"280749","name":"h"},"type":"job","steps":[{"stepType":"mysql","parameter":{"connection":[{"datasourceType":"mysql","datasource":"mm","selectedTables":[{"schema":[{"tableInfos":[{"enable":true,"table":"m_v1","tableName":"m_v1"}]}],"dbName":"m"}]}]},"name":"reader","category":"reader"},{"stepType":"holo","parameter":{"datasource":"h","tableMappingRule":{"datasource":[{"tableRule":[{"srcTable":"m_v1","mergeIntoCycleType":"DEFAULT","hourDeltaEnable":false,"dstTable":"m.m_v1","dayDeltaEnable":false,"primaryKeyInfo":{"column":["id"],"type":"pk"},"dstCreateTableInfo":{"indexType":"m_v1","dataColumn":[{"columnSize":0,"name":"id","index":0,"comment":"","newDigit":0,"type":"int8","digit":0,"primaryKey":true}],"schemaName":"m","tableName":"m_v1"},"srcDbName":"m"],"srcDatasourceName":"mm"}],"totalTableMapping":1},"writeMode":"replay"},"name":"writer","category":"writer"}],"version":"2.0","order":{"hops":[{"from":"reader","to":"writer"}]}}
	TaskContent *string `json:"TaskContent,omitempty" xml:"TaskContent,omitempty"`
}

func (s QueryDISyncTaskConfigProcessResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDISyncTaskConfigProcessResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDISyncTaskConfigProcessResultResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *QueryDISyncTaskConfigProcessResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryDISyncTaskConfigProcessResultResponseBodyData) GetTaskContent() *string {
	return s.TaskContent
}

func (s *QueryDISyncTaskConfigProcessResultResponseBodyData) SetMessage(v string) *QueryDISyncTaskConfigProcessResultResponseBodyData {
	s.Message = &v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultResponseBodyData) SetStatus(v string) *QueryDISyncTaskConfigProcessResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultResponseBodyData) SetTaskContent(v string) *QueryDISyncTaskConfigProcessResultResponseBodyData {
	s.TaskContent = &v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
