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
	// -----BEGIN CERTIFICATE-----
	//
	// MIID****r4w=
	//
	// -----END CERTIFICATE-----
	CaCert *string `json:"CaCert,omitempty" xml:"CaCert,omitempty"`
	// The client certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDR****A==
	//
	// -----END CERTIFICATE-----
	ClientCert *string `json:"ClientCert,omitempty" xml:"ClientCert,omitempty"`
	// The client configuration.
	//
	// example:
	//
	// client
	//
	// dev tun
	//
	// proto tcp
	//
	// remote 120.XX.XX.33 1194
	//
	// resolv-retry infinite
	//
	// nobind
	//
	// persist-key
	//
	// persist-tun
	//
	// ca ca.crt
	//
	// cert vsc-bp15t7****.crt
	//
	// key vsc-bp15t7****.key
	//
	// cipher AES-128-CBC
	//
	// ;comp-lzo
	//
	// verb 4
	ClientConfig *string `json:"ClientConfig,omitempty" xml:"ClientConfig,omitempty"`
	// The client key.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----
	//
	// MIIEpAI****ZzNFhuM/za==
	//
	// -----END RSA PRIVATE KEY-----
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	// The timestamp when the SSL client certificate was created. Unit: milliseconds.
	//
	// The timestamp is in the Unix format, which represents the total number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC, to the time when the SSL client certificate was created.
	//
	// example:
	//
	// 1552550980000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp when the SSL client certificate expires. Unit: milliseconds.
	//
	// The timestamp is in the Unix format, which represents the total number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC, to the time when the SSL client certificate expires.
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
	// The region ID of the SSL client certificate.
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
	// The resource group ID of the SSL client certificate.
	//
	// The SSL client certificate belongs to the same resource group as its associated SSL server. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource group information.
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
	// The SSL server ID.
	//
	// example:
	//
	// vss-bp155e9yclsg1xgq4****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
	// The status of the SSL client certificate. Valid values:
	//
	// - **expiring-soon**: The certificate will expire within one week.
	//
	// - **normal**: Normal.
	//
	// - **expired**: Expired.
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
