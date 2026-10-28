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
	SetSSLEnabled(v string) *ModifyDBClusterSSLRequest
	GetSSLEnabled() *string
}

type ModifyDBClusterSSLRequest struct {
	// The internal or public endpoint for which to create or update a server certificate.
	//
	// example:
	//
	// amv-***********.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable SSL connection encryption. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The SSL operation. This parameter and EnableSSL are mutually exclusive. Valid values:
	//
	// - Enable: Enables SSL.
	//
	// - Disable: Disables SSL.
	//
	// - Update: Updates the validity period of an enabled SSL certificate and retains the current endpoint bound to the certificate. When you set this parameter to Update, do not specify ConnectionString, or specify an endpoint that is the same as the current endpoint bound to the certificate.
	//
	// example:
	//
	// Update
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
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

func (s *ModifyDBClusterSSLRequest) GetSSLEnabled() *string {
	return s.SSLEnabled
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

func (s *ModifyDBClusterSSLRequest) SetSSLEnabled(v string) *ModifyDBClusterSSLRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) Validate() error {
	return dara.Validate(s)
}
