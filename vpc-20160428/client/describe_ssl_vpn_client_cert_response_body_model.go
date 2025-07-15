// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnClientCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaCert(v string) *DescribeSslVpnClientCertResponseBody
	GetCaCert() *string
	SetClientCert(v string) *DescribeSslVpnClientCertResponseBody
	GetClientCert() *string
	SetClientConfig(v string) *DescribeSslVpnClientCertResponseBody
	GetClientConfig() *string
	SetClientKey(v string) *DescribeSslVpnClientCertResponseBody
	GetClientKey() *string
	SetCreateTime(v int64) *DescribeSslVpnClientCertResponseBody
	GetCreateTime() *int64
	SetEndTime(v int64) *DescribeSslVpnClientCertResponseBody
	GetEndTime() *int64
	SetName(v string) *DescribeSslVpnClientCertResponseBody
	GetName() *string
	SetRegionId(v string) *DescribeSslVpnClientCertResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeSslVpnClientCertResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeSslVpnClientCertResponseBody
	GetResourceGroupId() *string
	SetSslVpnClientCertId(v string) *DescribeSslVpnClientCertResponseBody
	GetSslVpnClientCertId() *string
	SetSslVpnServerId(v string) *DescribeSslVpnClientCertResponseBody
	GetSslVpnServerId() *string
	SetStatus(v string) *DescribeSslVpnClientCertResponseBody
	GetStatus() *string
}

type DescribeSslVpnClientCertResponseBody struct {
	// The CA certificate.
	//
	// example:
	//
	// CA certificate
	CaCert *string `json:"CaCert,omitempty" xml:"CaCert,omitempty"`
	// The client certificate.
	//
	// example:
	//
	// Client certificate
	ClientCert *string `json:"ClientCert,omitempty" xml:"ClientCert,omitempty"`
	// The client configuration.
	//
	// example:
	//
	// Client configuration
	ClientConfig *string `json:"ClientConfig,omitempty" xml:"ClientConfig,omitempty"`
	// The client key.
	//
	// example:
	//
	// The key of the client
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	// The timestamp that indicates when the SSL client certificate was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552550980000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp that indicates when the SSL client certificate expires. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1647158980000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the SSL client certificate.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the SSL client certificate is created.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the SSL client certificate belongs.
	//
	// The SSL client certificate and the SSL server associated with the SSL client certificate belong to the same resource group. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the SSL client certificate.
	//
	// example:
	//
	// vsc-bp13k5mp4tg8v3z9b****
	SslVpnClientCertId *string `json:"SslVpnClientCertId,omitempty" xml:"SslVpnClientCertId,omitempty"`
	// The ID of the SSL server.
	//
	// example:
	//
	// vss-bp155e9yclsg1xgq4****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
	// The status of the SSL client certificate. Valid values:
	//
	// 	- **expiring-soon**
	//
	// 	- **normal**
	//
	// 	- **expired**
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSslVpnClientCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientCertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientCertResponseBody) GetCaCert() *string {
	return s.CaCert
}

func (s *DescribeSslVpnClientCertResponseBody) GetClientCert() *string {
	return s.ClientCert
}

func (s *DescribeSslVpnClientCertResponseBody) GetClientConfig() *string {
	return s.ClientConfig
}

func (s *DescribeSslVpnClientCertResponseBody) GetClientKey() *string {
	return s.ClientKey
}

func (s *DescribeSslVpnClientCertResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSslVpnClientCertResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSslVpnClientCertResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeSslVpnClientCertResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSslVpnClientCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSslVpnClientCertResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSslVpnClientCertResponseBody) GetSslVpnClientCertId() *string {
	return s.SslVpnClientCertId
}

func (s *DescribeSslVpnClientCertResponseBody) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *DescribeSslVpnClientCertResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeSslVpnClientCertResponseBody) SetCaCert(v string) *DescribeSslVpnClientCertResponseBody {
	s.CaCert = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetClientCert(v string) *DescribeSslVpnClientCertResponseBody {
	s.ClientCert = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetClientConfig(v string) *DescribeSslVpnClientCertResponseBody {
	s.ClientConfig = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetClientKey(v string) *DescribeSslVpnClientCertResponseBody {
	s.ClientKey = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetCreateTime(v int64) *DescribeSslVpnClientCertResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetEndTime(v int64) *DescribeSslVpnClientCertResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetName(v string) *DescribeSslVpnClientCertResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetRegionId(v string) *DescribeSslVpnClientCertResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetRequestId(v string) *DescribeSslVpnClientCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetResourceGroupId(v string) *DescribeSslVpnClientCertResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetSslVpnClientCertId(v string) *DescribeSslVpnClientCertResponseBody {
	s.SslVpnClientCertId = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetSslVpnServerId(v string) *DescribeSslVpnClientCertResponseBody {
	s.SslVpnServerId = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) SetStatus(v string) *DescribeSslVpnClientCertResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSslVpnClientCertResponseBody) Validate() error {
	return dara.Validate(s)
}
