// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataSourceMetaResponseBodyData) *GetDataSourceMetaResponseBody
	GetData() *GetDataSourceMetaResponseBodyData
	SetRequestId(v string) *GetDataSourceMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataSourceMetaResponseBody
	GetSuccess() *bool
}

type GetDataSourceMetaResponseBody struct {
	// The returned result.
	Data *GetDataSourceMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0000-ABCD-EFG
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

func (s GetDataSourceMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceMetaResponseBody) GetData() *GetDataSourceMetaResponseBodyData {
	return s.Data
}

func (s *GetDataSourceMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataSourceMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataSourceMetaResponseBody) SetData(v *GetDataSourceMetaResponseBodyData) *GetDataSourceMetaResponseBody {
	s.Data = v
	return s
}

func (s *GetDataSourceMetaResponseBody) SetRequestId(v string) *GetDataSourceMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceMetaResponseBody) SetSuccess(v bool) *GetDataSourceMetaResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataSourceMetaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceMetaResponseBodyData struct {
	// The reason why the metadata of the data source fails to be obtained. If the metadata of the data source is obtained, no value is returned for this parameter.
	//
	// example:
	//
	// read datasource time out
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned metadata of the data source. The returned metadata is in the JSON format.
	//
	// `{"dbTables":[{"dbName":"testdb","schema":[{"tableInfos":[{"dbName":"testdb","enable":true,"table":"table1","tableName":"table1"}]},{"tableInfos":[{"dbName":"testdb","enable":true,"table":"table2","tableName":"table2"}]}]}]}`
	//
	// Parameter description:
	//
	// 	- dbName: the name of the database in which the data source resides.
	//
	// 	- schema: the schema of the database.
	//
	// 	- enable: indicates whether the database is available. The valid values are true and false. The value true indicates that the database is available. The value false indicates that the database is unavailable.
	//
	// 	- tableName: the name of the table in the database.
	//
	// 	- tableInfos: the information about the table in the database.
	//
	// example:
	//
	// {"dbTables":[{"dbName":"testdb","schema":[{"tableInfos":[{"dbName":"testdb","enable":true,"table":"table1","tableName":"table1"}]},{"tableInfos":[{"dbName":"testdb","enable":true,"table":"table2","tableName":"table2"}]}]}]}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// Indicates whether the metadata of the data source is obtained. Valid values:
	//
	// 	- success: The metadata of the data source is obtained.
	//
	// 	- fail: The metadata of the data source failed to be obtained. You can troubleshoot issues based on the Message parameter.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDataSourceMetaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataSourceMetaResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetDataSourceMetaResponseBodyData) GetMeta() *string {
	return s.Meta
}

func (s *GetDataSourceMetaResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDataSourceMetaResponseBodyData) SetMessage(v string) *GetDataSourceMetaResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetDataSourceMetaResponseBodyData) SetMeta(v string) *GetDataSourceMetaResponseBodyData {
	s.Meta = &v
	return s
}

func (s *GetDataSourceMetaResponseBodyData) SetStatus(v string) *GetDataSourceMetaResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDataSourceMetaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
