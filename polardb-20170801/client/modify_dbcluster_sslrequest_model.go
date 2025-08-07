// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterSSLRequest
	GetDBClusterId() *string
	SetDBEndpointId(v string) *ModifyDBClusterSSLRequest
	GetDBEndpointId() *string
	SetNetType(v string) *ModifyDBClusterSSLRequest
	GetNetType() *string
	SetOwnerAccount(v string) *ModifyDBClusterSSLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterSSLRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterSSLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterSSLRequest
	GetResourceOwnerId() *int64
	SetSSLAutoRotate(v string) *ModifyDBClusterSSLRequest
	GetSSLAutoRotate() *string
	SetSSLEnabled(v string) *ModifyDBClusterSSLRequest
	GetSSLEnabled() *string
}

type ModifyDBClusterSSLRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the endpoint.
	//
	// >
	//
	// 	- This parameter is required for PolarDB for MySQL clusters.
	//
	// 	- This parameter is not required for PolarDB for PostgreSQL or PolarDB for PostgreSQL (Compatible with Oracle) clusters. By default, SSL encryption is enabled for all endpoints of the clusters.
	//
	// 	- You can call the [DescribeDBClusterSSL](https://help.aliyun.com/document_detail/2319159.html) operation to view the details of the endpoint.
	//
	// example:
	//
	// pe-******************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The network type supported by the endpoint that is specified by **DBEndpointId**. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// 	- **Inner**
	//
	// >
	//
	// 	- This parameter is required for a PolarDB for MySQL cluster.
	//
	// 	- This parameter is not required for a PolarDB for Oracle or PolarDB for PostgreSQL cluster. By default, SSL encryption is enabled for all endpoints.
	//
	// example:
	//
	// Public
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether automatic rotation of SSL certificates is enabled.
	//
	// 	- **Enable**: The feature is enabled.
	//
	// 	- **Disable**: The feature is disabled.
	//
	// example:
	//
	// Enable
	SSLAutoRotate *string `json:"SSLAutoRotate,omitempty" xml:"SSLAutoRotate,omitempty"`
	// The SSL encryption status. Valid values:
	//
	// 	- **Disable**: SSL encryption is disabled.
	//
	// 	- **Enable**: SSL encryption is enabled.
	//
	// 	- **Update**: The SSL certificate is updated.
	//
	// > After you enable SSL encryption or update the SSL certificate, you must download and configure the certificate. For more information, see [Configure SSL encryption](https://help.aliyun.com/document_detail/153182.html).
	//
	// example:
	//
	// Enable
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
}

func (s ModifyDBClusterSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterSSLRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *ModifyDBClusterSSLRequest) GetNetType() *string {
	return s.NetType
}

func (s *ModifyDBClusterSSLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterSSLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterSSLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterSSLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterSSLRequest) GetSSLAutoRotate() *string {
	return s.SSLAutoRotate
}

func (s *ModifyDBClusterSSLRequest) GetSSLEnabled() *string {
	return s.SSLEnabled
}

func (s *ModifyDBClusterSSLRequest) SetDBClusterId(v string) *ModifyDBClusterSSLRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetDBEndpointId(v string) *ModifyDBClusterSSLRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetNetType(v string) *ModifyDBClusterSSLRequest {
	s.NetType = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetOwnerAccount(v string) *ModifyDBClusterSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetOwnerId(v int64) *ModifyDBClusterSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetResourceOwnerId(v int64) *ModifyDBClusterSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetSSLAutoRotate(v string) *ModifyDBClusterSSLRequest {
	s.SSLAutoRotate = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetSSLEnabled(v string) *ModifyDBClusterSSLRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) Validate() error {
	return dara.Validate(s)
}
