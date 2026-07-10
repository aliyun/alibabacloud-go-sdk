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
	SetDBEndpointId(v string) *ModifyDBClusterSSLRequest
	GetDBEndpointId() *string
	SetNetType(v string) *ModifyDBClusterSSLRequest
	GetNetType() *string
	SetOwnerAccount(v string) *ModifyDBClusterSSLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterSSLRequest
	GetOwnerId() *int64
	SetPfsInstanceId(v string) *ModifyDBClusterSSLRequest
	GetPfsInstanceId() *string
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
	// example:
	//
	// xxx
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The endpoint ID.
	//
	// >	- If the cluster is a PolarDB for MySQL cluster, this parameter is required.
	//
	// >	- If the cluster is a PolarDB for PostgreSQL cluster or a PolarDB for PostgreSQL (Compatible with Oracle) cluster, you do not need to specify this parameter. SSL encryption is enabled for all endpoints by default.
	//
	// >	- You can call the [DescribeDBClusterSSL](https://help.aliyun.com/document_detail/2319159.html) operation to query endpoint details.
	//
	// example:
	//
	// pe-******************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The network type of the endpoint. The value must be the same as the network type of the endpoint specified by **DBEndpointId**. Valid values:
	//
	// 	- **Public**: public network
	//
	// 	- **Private**: private network
	//
	// 	- **Inner**: private network (classic network)
	//
	// >	- If the cluster is a PolarDB for MySQL cluster, this parameter is required.
	//
	// >	- If the cluster is a PolarDB for PostgreSQL cluster or a PolarDB for PostgreSQL (Compatible with Oracle) cluster, you do not need to specify this parameter. SSL encryption is enabled for all endpoints by default.
	//
	// example:
	//
	// Public
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// pfs-xxx
	PfsInstanceId        *string `json:"PfsInstanceId,omitempty" xml:"PfsInstanceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable automatic SSL certificate rotation. Valid values:
	//
	// - **Enable**: enables automatic SSL certificate rotation.
	//
	// - **Disable**: disables automatic SSL certificate rotation.
	//
	// example:
	//
	// Enable
	SSLAutoRotate *string `json:"SSLAutoRotate,omitempty" xml:"SSLAutoRotate,omitempty"`
	// The SSL status. Valid values:
	//
	// 	- **Disable**: shutdown SSL encryption.
	//
	// 	- **Enable**: enables SSL encryption.
	//
	// 	- **Update**: updates the CA certificate.
	//
	// > After you enable SSL encryption or update the CA certificate, you must download and configure the certificate. For details, see [Settings for SSL encryption](https://help.aliyun.com/document_detail/153182.html).
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

func (s *ModifyDBClusterSSLRequest) GetConnectionString() *string {
	return s.ConnectionString
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

func (s *ModifyDBClusterSSLRequest) GetPfsInstanceId() *string {
	return s.PfsInstanceId
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

func (s *ModifyDBClusterSSLRequest) SetConnectionString(v string) *ModifyDBClusterSSLRequest {
	s.ConnectionString = &v
	return s
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

func (s *ModifyDBClusterSSLRequest) SetPfsInstanceId(v string) *ModifyDBClusterSSLRequest {
	s.PfsInstanceId = &v
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
