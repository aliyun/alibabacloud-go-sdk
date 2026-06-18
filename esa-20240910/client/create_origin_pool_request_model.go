// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *CreateOriginPoolRequest
	GetEnabled() *bool
	SetName(v string) *CreateOriginPoolRequest
	GetName() *string
	SetOrigins(v []*CreateOriginPoolRequestOrigins) *CreateOriginPoolRequest
	GetOrigins() []*CreateOriginPoolRequestOrigins
	SetSiteId(v int64) *CreateOriginPoolRequest
	GetSiteId() *int64
}

type CreateOriginPoolRequest struct {
	// Specifies whether the origin pool is enabled.
	//
	// - `true`: enabled
	//
	// - `false`: disabled
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the origin pool. The name must be unique within a site.
	//
	// This parameter is required.
	//
	// example:
	//
	// pool1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of origins to add to the origin pool. Use an array to specify multiple origins.
	Origins []*CreateOriginPoolRequestOrigins `json:"Origins,omitempty" xml:"Origins,omitempty" type:"Repeated"`
	// The site ID. To obtain this ID, call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21655860979****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateOriginPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolRequest) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateOriginPoolRequest) GetName() *string {
	return s.Name
}

func (s *CreateOriginPoolRequest) GetOrigins() []*CreateOriginPoolRequestOrigins {
	return s.Origins
}

func (s *CreateOriginPoolRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateOriginPoolRequest) SetEnabled(v bool) *CreateOriginPoolRequest {
	s.Enabled = &v
	return s
}

func (s *CreateOriginPoolRequest) SetName(v string) *CreateOriginPoolRequest {
	s.Name = &v
	return s
}

func (s *CreateOriginPoolRequest) SetOrigins(v []*CreateOriginPoolRequestOrigins) *CreateOriginPoolRequest {
	s.Origins = v
	return s
}

func (s *CreateOriginPoolRequest) SetSiteId(v int64) *CreateOriginPoolRequest {
	s.SiteId = &v
	return s
}

func (s *CreateOriginPoolRequest) Validate() error {
	if s.Origins != nil {
		for _, item := range s.Origins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOriginPoolRequestOrigins struct {
	// The address of the origin. For example, www\\.example.com.
	//
	// example:
	//
	// www.example.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The authentication configuration for the origin. Required if the origin is an OSS or AWS S3 bucket that requires authentication.
	AuthConf *CreateOriginPoolRequestOriginsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// Specifies whether the origin is enabled.
	//
	// - `true`: enabled
	//
	// - `false`: disabled
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request header to include in back-to-origin requests. Only the `Host` header is supported.
	//
	// example:
	//
	// {
	//
	//         "Host": [
	//
	//           "example.com"
	//
	//         ]
	//
	//       }
	Header interface{} `json:"Header,omitempty" xml:"Header,omitempty"`
	// The IP protocol version for back-to-origin requests. Valid values:
	//
	// - `round_robin`: Default. Randomly selects an IPv4 or IPv6 origin.
	//
	// - `ipv4_first`: Prioritizes IPv4 origins.
	//
	// - `ipv6_first`: Prioritizes IPv6 origins.
	//
	// - `follow`: Uses the same IP protocol version as the client\\"s request.
	//
	// example:
	//
	// round_robin
	IpVersionPolicy *string `json:"IpVersionPolicy,omitempty" xml:"IpVersionPolicy,omitempty"`
	// The name of the origin. The name must be unique within the origin pool.
	//
	// example:
	//
	// origin1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the origin. Valid values:
	//
	// - `ip_domain`: An IP address or a domain name.
	//
	// - `OSS`: An Alibaba Cloud OSS bucket.
	//
	// - `S3`: An AWS S3 bucket.
	//
	// example:
	//
	// ip_domain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin. The value must be an integer from 0 to 100.
	//
	// example:
	//
	// 50
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateOriginPoolRequestOrigins) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolRequestOrigins) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolRequestOrigins) GetAddress() *string {
	return s.Address
}

func (s *CreateOriginPoolRequestOrigins) GetAuthConf() *CreateOriginPoolRequestOriginsAuthConf {
	return s.AuthConf
}

func (s *CreateOriginPoolRequestOrigins) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateOriginPoolRequestOrigins) GetHeader() interface{} {
	return s.Header
}

func (s *CreateOriginPoolRequestOrigins) GetIpVersionPolicy() *string {
	return s.IpVersionPolicy
}

func (s *CreateOriginPoolRequestOrigins) GetName() *string {
	return s.Name
}

func (s *CreateOriginPoolRequestOrigins) GetType() *string {
	return s.Type
}

func (s *CreateOriginPoolRequestOrigins) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateOriginPoolRequestOrigins) SetAddress(v string) *CreateOriginPoolRequestOrigins {
	s.Address = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetAuthConf(v *CreateOriginPoolRequestOriginsAuthConf) *CreateOriginPoolRequestOrigins {
	s.AuthConf = v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetEnabled(v bool) *CreateOriginPoolRequestOrigins {
	s.Enabled = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetHeader(v interface{}) *CreateOriginPoolRequestOrigins {
	s.Header = v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetIpVersionPolicy(v string) *CreateOriginPoolRequestOrigins {
	s.IpVersionPolicy = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetName(v string) *CreateOriginPoolRequestOrigins {
	s.Name = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetType(v string) *CreateOriginPoolRequestOrigins {
	s.Type = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) SetWeight(v int32) *CreateOriginPoolRequestOrigins {
	s.Weight = &v
	return s
}

func (s *CreateOriginPoolRequestOrigins) Validate() error {
	if s.AuthConf != nil {
		if err := s.AuthConf.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOriginPoolRequestOriginsAuthConf struct {
	// The access key required for private authentication.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication type. Valid values:
	//
	// - `public`: Public read/write. Use this for public OSS or AWS S3 buckets.
	//
	// - `private_same_account`: Private authentication for an OSS bucket in the same Alibaba Cloud account.
	//
	// - `private_cross_account`: Private authentication for an OSS bucket in a different Alibaba Cloud account.
	//
	// - `private`: Private authentication for an AWS S3 bucket.
	//
	// example:
	//
	// public
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region of the origin. Required for AWS S3 origins.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The secret key required for private authentication.
	//
	// example:
	//
	// yourAccessKeySecret
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signature version. Required for AWS S3 origins.
	//
	// example:
	//
	// v2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateOriginPoolRequestOriginsAuthConf) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolRequestOriginsAuthConf) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetRegion() *string {
	return s.Region
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *CreateOriginPoolRequestOriginsAuthConf) GetVersion() *string {
	return s.Version
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetAccessKey(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.AccessKey = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetAuthType(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.AuthType = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetRegion(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.Region = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetSecretKey(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.SecretKey = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) SetVersion(v string) *CreateOriginPoolRequestOriginsAuthConf {
	s.Version = &v
	return s
}

func (s *CreateOriginPoolRequestOriginsAuthConf) Validate() error {
	return dara.Validate(s)
}
