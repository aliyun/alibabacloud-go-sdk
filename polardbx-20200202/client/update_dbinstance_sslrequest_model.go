// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertCommonName(v string) *UpdateDBInstanceSSLRequest
	GetCertCommonName() *string
	SetDBInstanceName(v string) *UpdateDBInstanceSSLRequest
	GetDBInstanceName() *string
	SetEnableSSL(v bool) *UpdateDBInstanceSSLRequest
	GetEnableSSL() *bool
	SetRegionId(v string) *UpdateDBInstanceSSLRequest
	GetRegionId() *string
}

type UpdateDBInstanceSSLRequest struct {
	// example:
	//
	// pxc-hzrqjarxdocd4t.polarx.rds.aliyuncs.com
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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

func (s UpdateDBInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLRequest) GetCertCommonName() *string {
	return s.CertCommonName
}

func (s *UpdateDBInstanceSSLRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpdateDBInstanceSSLRequest) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *UpdateDBInstanceSSLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDBInstanceSSLRequest) SetCertCommonName(v string) *UpdateDBInstanceSSLRequest {
	s.CertCommonName = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) SetDBInstanceName(v string) *UpdateDBInstanceSSLRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) SetEnableSSL(v bool) *UpdateDBInstanceSSLRequest {
	s.EnableSSL = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) SetRegionId(v string) *UpdateDBInstanceSSLRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDBInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
