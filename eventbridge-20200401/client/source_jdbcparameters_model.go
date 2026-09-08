// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceJDBCParameters interface {
	dara.Model
	String() string
	GoString() string
	SetBatchSize(v int32) *SourceJDBCParameters
	GetBatchSize() *int32
	SetCustomQuery(v string) *SourceJDBCParameters
	GetCustomQuery() *string
	SetIncrementingColumn(v string) *SourceJDBCParameters
	GetIncrementingColumn() *string
	SetJdbcUrl(v string) *SourceJDBCParameters
	GetJdbcUrl() *string
	SetNetwork(v string) *SourceJDBCParameters
	GetNetwork() *string
	SetPassword(v string) *SourceJDBCParameters
	GetPassword() *string
	SetPollingInterval(v int32) *SourceJDBCParameters
	GetPollingInterval() *int32
	SetQueryMode(v string) *SourceJDBCParameters
	GetQueryMode() *string
	SetQueryTimeout(v int32) *SourceJDBCParameters
	GetQueryTimeout() *int32
	SetSecurityGroupId(v string) *SourceJDBCParameters
	GetSecurityGroupId() *string
	SetTableName(v string) *SourceJDBCParameters
	GetTableName() *string
	SetTimestampColumn(v string) *SourceJDBCParameters
	GetTimestampColumn() *string
	SetUsername(v string) *SourceJDBCParameters
	GetUsername() *string
	SetVSwitchIds(v string) *SourceJDBCParameters
	GetVSwitchIds() *string
	SetVpcId(v string) *SourceJDBCParameters
	GetVpcId() *string
}

type SourceJDBCParameters struct {
	// The maximum number of rows returned per query. Default value: 1000. Maximum value: 10000.
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// The custom SQL query statement (advanced mode). This parameter is mutually exclusive with TableName.
	CustomQuery *string `json:"CustomQuery,omitempty" xml:"CustomQuery,omitempty"`
	// The incrementing column name. Required when QueryMode is set to `incrementing` or `timestamp_incrementing`.
	IncrementingColumn *string `json:"IncrementingColumn,omitempty" xml:"IncrementingColumn,omitempty"`
	// The JDBC connection URL. ClickHouse example: `jdbc:clickhouse://host:8123/database`; MySQL example: `jdbc:mysql://host:3306/database`. The system automatically identifies the database type.
	JdbcUrl *string `json:"JdbcUrl,omitempty" xml:"JdbcUrl,omitempty"`
	// The network type. Valid values: `PublicNetwork` (public network); `PrivateNetwork` (VPC private network, requires VpcId, VSwitchIds, and SecurityGroupId to be configured).
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The database password.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The polling interval, in seconds. Minimum value: 10. Default value: 60.
	PollingInterval *int32 `json:"PollingInterval,omitempty" xml:"PollingInterval,omitempty"`
	// The query mode. Valid values: `bulk` (full query); `incrementing` (incrementing column tracking); `timestamp` (timestamp tracking); `timestamp_incrementing` (timestamp and incrementing column dual tracking).
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// The SQL query timeout period, in seconds. Default value: 30. Maximum value: 300.
	QueryTimeout *int32 `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	// The security group ID. Required when Network is set to PrivateNetwork.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The target table name. This parameter is mutually exclusive with CustomQuery. Required when custom SQL is not used.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The timestamp column name. Required when QueryMode is set to `timestamp` or `timestamp_incrementing`.
	TimestampColumn *string `json:"TimestampColumn,omitempty" xml:"TimestampColumn,omitempty"`
	// The database username.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The vSwitch ID. Required when Network is set to PrivateNetwork.
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID. Required when Network is set to PrivateNetwork.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SourceJDBCParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceJDBCParameters) GoString() string {
	return s.String()
}

func (s *SourceJDBCParameters) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *SourceJDBCParameters) GetCustomQuery() *string {
	return s.CustomQuery
}

func (s *SourceJDBCParameters) GetIncrementingColumn() *string {
	return s.IncrementingColumn
}

func (s *SourceJDBCParameters) GetJdbcUrl() *string {
	return s.JdbcUrl
}

func (s *SourceJDBCParameters) GetNetwork() *string {
	return s.Network
}

func (s *SourceJDBCParameters) GetPassword() *string {
	return s.Password
}

func (s *SourceJDBCParameters) GetPollingInterval() *int32 {
	return s.PollingInterval
}

func (s *SourceJDBCParameters) GetQueryMode() *string {
	return s.QueryMode
}

func (s *SourceJDBCParameters) GetQueryTimeout() *int32 {
	return s.QueryTimeout
}

func (s *SourceJDBCParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SourceJDBCParameters) GetTableName() *string {
	return s.TableName
}

func (s *SourceJDBCParameters) GetTimestampColumn() *string {
	return s.TimestampColumn
}

func (s *SourceJDBCParameters) GetUsername() *string {
	return s.Username
}

func (s *SourceJDBCParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SourceJDBCParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SourceJDBCParameters) SetBatchSize(v int32) *SourceJDBCParameters {
	s.BatchSize = &v
	return s
}

func (s *SourceJDBCParameters) SetCustomQuery(v string) *SourceJDBCParameters {
	s.CustomQuery = &v
	return s
}

func (s *SourceJDBCParameters) SetIncrementingColumn(v string) *SourceJDBCParameters {
	s.IncrementingColumn = &v
	return s
}

func (s *SourceJDBCParameters) SetJdbcUrl(v string) *SourceJDBCParameters {
	s.JdbcUrl = &v
	return s
}

func (s *SourceJDBCParameters) SetNetwork(v string) *SourceJDBCParameters {
	s.Network = &v
	return s
}

func (s *SourceJDBCParameters) SetPassword(v string) *SourceJDBCParameters {
	s.Password = &v
	return s
}

func (s *SourceJDBCParameters) SetPollingInterval(v int32) *SourceJDBCParameters {
	s.PollingInterval = &v
	return s
}

func (s *SourceJDBCParameters) SetQueryMode(v string) *SourceJDBCParameters {
	s.QueryMode = &v
	return s
}

func (s *SourceJDBCParameters) SetQueryTimeout(v int32) *SourceJDBCParameters {
	s.QueryTimeout = &v
	return s
}

func (s *SourceJDBCParameters) SetSecurityGroupId(v string) *SourceJDBCParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SourceJDBCParameters) SetTableName(v string) *SourceJDBCParameters {
	s.TableName = &v
	return s
}

func (s *SourceJDBCParameters) SetTimestampColumn(v string) *SourceJDBCParameters {
	s.TimestampColumn = &v
	return s
}

func (s *SourceJDBCParameters) SetUsername(v string) *SourceJDBCParameters {
	s.Username = &v
	return s
}

func (s *SourceJDBCParameters) SetVSwitchIds(v string) *SourceJDBCParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SourceJDBCParameters) SetVpcId(v string) *SourceJDBCParameters {
	s.VpcId = &v
	return s
}

func (s *SourceJDBCParameters) Validate() error {
	return dara.Validate(s)
}
