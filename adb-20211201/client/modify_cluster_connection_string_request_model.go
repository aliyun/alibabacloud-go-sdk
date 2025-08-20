// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *ModifyClusterConnectionStringRequest
	GetConnectionStringPrefix() *string
	SetCurrentConnectionString(v string) *ModifyClusterConnectionStringRequest
	GetCurrentConnectionString() *string
	SetDBClusterId(v string) *ModifyClusterConnectionStringRequest
	GetDBClusterId() *string
	SetPort(v int32) *ModifyClusterConnectionStringRequest
	GetPort() *int32
}

type ModifyClusterConnectionStringRequest struct {
	// The prefix of the public endpoint.
	//
	// 	- The prefix can contain lowercase letters, digits, and hyphens (-). It must start with a lowercase letter.
	//
	// 	- The prefix can be up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-123
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The public endpoint of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****.ads.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The port number. Set the value to **3306**.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyClusterConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *ModifyClusterConnectionStringRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ModifyClusterConnectionStringRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyClusterConnectionStringRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifyClusterConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyClusterConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyClusterConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetDBClusterId(v string) *ModifyClusterConnectionStringRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetPort(v int32) *ModifyClusterConnectionStringRequest {
	s.Port = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
