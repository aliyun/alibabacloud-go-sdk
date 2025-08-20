// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *ModifyDBClusterSSLRequest
	GetConnectionString() *string
	SetDBClusterId(v string) *ModifyDBClusterSSLRequest
	GetDBClusterId() *string
	SetEnableSSL(v bool) *ModifyDBClusterSSLRequest
	GetEnableSSL() *bool
	SetRegionId(v string) *ModifyDBClusterSSLRequest
	GetRegionId() *string
}

type ModifyDBClusterSSLRequest struct {
	// example:
	//
	// amv-***********.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// amv-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBClusterSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBClusterSSLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterSSLRequest) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *ModifyDBClusterSSLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterSSLRequest) SetConnectionString(v string) *ModifyDBClusterSSLRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetDBClusterId(v string) *ModifyDBClusterSSLRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetEnableSSL(v bool) *ModifyDBClusterSSLRequest {
	s.EnableSSL = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetRegionId(v string) *ModifyDBClusterSSLRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) Validate() error {
	return dara.Validate(s)
}
