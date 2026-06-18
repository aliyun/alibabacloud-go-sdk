// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateOriginPoolRequest
	GetEnabled() *bool
	SetId(v int64) *UpdateOriginPoolRequest
	GetId() *int64
	SetOrigins(v []*UpdateOriginPoolRequestOrigins) *UpdateOriginPoolRequest
	GetOrigins() []*UpdateOriginPoolRequestOrigins
	SetSiteId(v int64) *UpdateOriginPoolRequest
	GetSiteId() *int64
}

type UpdateOriginPoolRequest struct {
	// Specifies whether to enable the origin pool:
	//
	// - true: Enables the origin pool.
	//
	// - false: Disables the origin pool.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The origin pool ID. Get this ID by calling the [ListOriginPools](~~ListOriginPools~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1038520525196928
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// An array of origin configurations.
	Origins []*UpdateOriginPoolRequestOrigins `json:"Origins,omitempty" xml:"Origins,omitempty" type:"Repeated"`
	// The site ID. Get this ID by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 216558609793952
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateOriginPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateOriginPoolRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateOriginPoolRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateOriginPoolRequest) GetOrigins() []*UpdateOriginPoolRequestOrigins {
	return s.Origins
}

func (s *UpdateOriginPoolRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateOriginPoolRequest) SetEnabled(v bool) *UpdateOriginPoolRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateOriginPoolRequest) SetId(v int64) *UpdateOriginPoolRequest {
	s.Id = &v
	return s
}

func (s *UpdateOriginPoolRequest) SetOrigins(v []*UpdateOriginPoolRequestOrigins) *UpdateOriginPoolRequest {
	s.Origins = v
	return s
}

func (s *UpdateOriginPoolRequest) SetSiteId(v int64) *UpdateOriginPoolRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateOriginPoolRequest) Validate() error {
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

type UpdateOriginPoolRequestOrigins struct {
	// The origin\\"s domain name or IP address.
	//
	// example:
	//
	// www.example.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The authentication configuration. Required for private OSS or S3 origins.
	AuthConf *UpdateOriginPoolRequestOriginsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// Specifies whether to enable the origin:
	//
	// - true: Enables the origin.
	//
	// - false: Disables the origin.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request header to add to back-to-origin requests. Only the Host header is supported.
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
	// The IP version policy for back-to-origin requests. Valid values:
	//
	// - round_robin: (Default) Randomly selects an IPv4 or IPv6 origin.
	//
	// - ipv4_first: Prioritizes IPv4 origins.
	//
	// - ipv6_first: Prioritizes IPv6 origins.
	//
	// - follow: Uses the same IP version as the client request.
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
	// The origin type. Valid values:
	//
	// - ip_domain: An IP address or a domain name.
	//
	// - OSS: An OSS origin.
	//
	// - S3: An AWS S3 origin.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin. The value must be an integer from 0 to 100.
	//
	// example:
	//
	// 50
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateOriginPoolRequestOrigins) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginPoolRequestOrigins) GoString() string {
	return s.String()
}

func (s *UpdateOriginPoolRequestOrigins) GetAddress() *string {
	return s.Address
}

func (s *UpdateOriginPoolRequestOrigins) GetAuthConf() *UpdateOriginPoolRequestOriginsAuthConf {
	return s.AuthConf
}

func (s *UpdateOriginPoolRequestOrigins) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateOriginPoolRequestOrigins) GetHeader() interface{} {
	return s.Header
}

func (s *UpdateOriginPoolRequestOrigins) GetIpVersionPolicy() *string {
	return s.IpVersionPolicy
}

func (s *UpdateOriginPoolRequestOrigins) GetName() *string {
	return s.Name
}

func (s *UpdateOriginPoolRequestOrigins) GetType() *string {
	return s.Type
}

func (s *UpdateOriginPoolRequestOrigins) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateOriginPoolRequestOrigins) SetAddress(v string) *UpdateOriginPoolRequestOrigins {
	s.Address = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetAuthConf(v *UpdateOriginPoolRequestOriginsAuthConf) *UpdateOriginPoolRequestOrigins {
	s.AuthConf = v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetEnabled(v bool) *UpdateOriginPoolRequestOrigins {
	s.Enabled = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetHeader(v interface{}) *UpdateOriginPoolRequestOrigins {
	s.Header = v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetIpVersionPolicy(v string) *UpdateOriginPoolRequestOrigins {
	s.IpVersionPolicy = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetName(v string) *UpdateOriginPoolRequestOrigins {
	s.Name = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetType(v string) *UpdateOriginPoolRequestOrigins {
	s.Type = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) SetWeight(v int32) *UpdateOriginPoolRequestOrigins {
	s.Weight = &v
	return s
}

func (s *UpdateOriginPoolRequestOrigins) Validate() error {
	if s.AuthConf != nil {
		if err := s.AuthConf.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateOriginPoolRequestOriginsAuthConf struct {
	// The access key for private authentication. Required for private origins.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication type. Valid values:
	//
	// - public: For public OSS or S3 origins.
	//
	// - private_same_account: For private OSS origins in the same account.
	//
	// - private_cross_account: For private OSS origins that use cross-account authentication.
	//
	// - private: For private S3 origins.
	//
	// example:
	//
	// public
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region of the origin. This parameter is required if the origin type is S3.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The secret key for private authentication. Required for private origins.
	//
	// example:
	//
	// yourAccessKeySecret
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signature version. This parameter is required if the origin type is S3.
	//
	// example:
	//
	// v2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateOriginPoolRequestOriginsAuthConf) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginPoolRequestOriginsAuthConf) GoString() string {
	return s.String()
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetRegion() *string {
	return s.Region
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) GetVersion() *string {
	return s.Version
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetAccessKey(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.AccessKey = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetAuthType(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.AuthType = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetRegion(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.Region = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetSecretKey(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.SecretKey = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) SetVersion(v string) *UpdateOriginPoolRequestOriginsAuthConf {
	s.Version = &v
	return s
}

func (s *UpdateOriginPoolRequestOriginsAuthConf) Validate() error {
	return dara.Validate(s)
}
