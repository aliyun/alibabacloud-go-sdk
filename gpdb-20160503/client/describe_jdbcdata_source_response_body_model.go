// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJDBCDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeJDBCDataSourceResponseBody
	GetCreateTime() *string
	SetDataSourceDescription(v string) *DescribeJDBCDataSourceResponseBody
	GetDataSourceDescription() *string
	SetDataSourceId(v string) *DescribeJDBCDataSourceResponseBody
	GetDataSourceId() *string
	SetDataSourceName(v string) *DescribeJDBCDataSourceResponseBody
	GetDataSourceName() *string
	SetDataSourceStatus(v string) *DescribeJDBCDataSourceResponseBody
	GetDataSourceStatus() *string
	SetDataSourceType(v string) *DescribeJDBCDataSourceResponseBody
	GetDataSourceType() *string
	SetExternalDataServiceId(v string) *DescribeJDBCDataSourceResponseBody
	GetExternalDataServiceId() *string
	SetJDBCConnectionString(v string) *DescribeJDBCDataSourceResponseBody
	GetJDBCConnectionString() *string
	SetJDBCPassword(v string) *DescribeJDBCDataSourceResponseBody
	GetJDBCPassword() *string
	SetJDBCUserName(v string) *DescribeJDBCDataSourceResponseBody
	GetJDBCUserName() *string
	SetModifyTime(v string) *DescribeJDBCDataSourceResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *DescribeJDBCDataSourceResponseBody
	GetRequestId() *string
	SetStatusMessage(v string) *DescribeJDBCDataSourceResponseBody
	GetStatusMessage() *string
}

type DescribeJDBCDataSourceResponseBody struct {
	// The time when the service was created.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the service. The description can be up to 256 characters in length.
	//
	// example:
	//
	// mysql data source config
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of data soruce
	//
	// example:
	//
	// hdfs_pxf
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Init
	//
	// 	- Running
	//
	// 	- Exception
	//
	// example:
	//
	// Running
	DataSourceStatus *string `json:"DataSourceStatus,omitempty" xml:"DataSourceStatus,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// MySQL
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The id of the external data service
	//
	// example:
	//
	// 2989
	ExternalDataServiceId *string `json:"ExternalDataServiceId,omitempty" xml:"ExternalDataServiceId,omitempty"`
	// The JDBC connection string.
	//
	// example:
	//
	// xxxxxx
	JDBCConnectionString *string `json:"JDBCConnectionString,omitempty" xml:"JDBCConnectionString,omitempty"`
	// The password of the database account.
	//
	// example:
	//
	// xxxxxx
	JDBCPassword *string `json:"JDBCPassword,omitempty" xml:"JDBCPassword,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// xxxxxx
	JDBCUserName *string `json:"JDBCUserName,omitempty" xml:"JDBCUserName,omitempty"`
	// The time when the data source was last modified.
	//
	// example:
	//
	// 2024-08-27T02:01:10Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The message of the status
	//
	// example:
	//
	// ""
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DescribeJDBCDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJDBCDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJDBCDataSourceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeJDBCDataSourceResponseBody) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *DescribeJDBCDataSourceResponseBody) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeJDBCDataSourceResponseBody) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *DescribeJDBCDataSourceResponseBody) GetDataSourceStatus() *string {
	return s.DataSourceStatus
}

func (s *DescribeJDBCDataSourceResponseBody) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DescribeJDBCDataSourceResponseBody) GetExternalDataServiceId() *string {
	return s.ExternalDataServiceId
}

func (s *DescribeJDBCDataSourceResponseBody) GetJDBCConnectionString() *string {
	return s.JDBCConnectionString
}

func (s *DescribeJDBCDataSourceResponseBody) GetJDBCPassword() *string {
	return s.JDBCPassword
}

func (s *DescribeJDBCDataSourceResponseBody) GetJDBCUserName() *string {
	return s.JDBCUserName
}

func (s *DescribeJDBCDataSourceResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeJDBCDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJDBCDataSourceResponseBody) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *DescribeJDBCDataSourceResponseBody) SetCreateTime(v string) *DescribeJDBCDataSourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetDataSourceDescription(v string) *DescribeJDBCDataSourceResponseBody {
	s.DataSourceDescription = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetDataSourceId(v string) *DescribeJDBCDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetDataSourceName(v string) *DescribeJDBCDataSourceResponseBody {
	s.DataSourceName = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetDataSourceStatus(v string) *DescribeJDBCDataSourceResponseBody {
	s.DataSourceStatus = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetDataSourceType(v string) *DescribeJDBCDataSourceResponseBody {
	s.DataSourceType = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetExternalDataServiceId(v string) *DescribeJDBCDataSourceResponseBody {
	s.ExternalDataServiceId = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetJDBCConnectionString(v string) *DescribeJDBCDataSourceResponseBody {
	s.JDBCConnectionString = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetJDBCPassword(v string) *DescribeJDBCDataSourceResponseBody {
	s.JDBCPassword = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetJDBCUserName(v string) *DescribeJDBCDataSourceResponseBody {
	s.JDBCUserName = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetModifyTime(v string) *DescribeJDBCDataSourceResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetRequestId(v string) *DescribeJDBCDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) SetStatusMessage(v string) *DescribeJDBCDataSourceResponseBody {
	s.StatusMessage = &v
	return s
}

func (s *DescribeJDBCDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
