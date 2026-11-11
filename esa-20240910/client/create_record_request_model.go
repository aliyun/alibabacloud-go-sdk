// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConf(v *CreateRecordRequestAuthConf) *CreateRecordRequest
	GetAuthConf() *CreateRecordRequestAuthConf
	SetBizName(v string) *CreateRecordRequest
	GetBizName() *string
	SetComment(v string) *CreateRecordRequest
	GetComment() *string
	SetData(v *CreateRecordRequestData) *CreateRecordRequest
	GetData() *CreateRecordRequestData
	SetHostPolicy(v string) *CreateRecordRequest
	GetHostPolicy() *string
	SetHttpPorts(v string) *CreateRecordRequest
	GetHttpPorts() *string
	SetHttpsPorts(v string) *CreateRecordRequest
	GetHttpsPorts() *string
	SetProxied(v bool) *CreateRecordRequest
	GetProxied() *bool
	SetRecordName(v string) *CreateRecordRequest
	GetRecordName() *string
	SetSiteId(v int64) *CreateRecordRequest
	GetSiteId() *int64
	SetSourceType(v string) *CreateRecordRequest
	GetSourceType() *string
	SetTtl(v int32) *CreateRecordRequest
	GetTtl() *int32
	SetType(v string) *CreateRecordRequest
	GetType() *string
}

type CreateRecordRequest struct {
	// The origin server authentication information for a CNAME record.
	AuthConf *CreateRecordRequestAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The business scenario used to tag the DNS record. This parameter is required when proxy acceleration is enabled for the DNS record. That is, the Proxied parameter is set to true. This parameter is not required when proxy acceleration is disabled. That is, the Proxied parameter is set to false. Valid values:
	//
	// - **image_video**: Images and videos.
	//
	// - **api**: API operations.
	//
	// - **web**: Web pages.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The description of the record. The description can be up to 100 characters in length.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS information of the record. The required content for this field varies based on the record type. For more information, see
	//
	// <props="china">[Documentation](https://help.aliyun.com/document_detail/2708761.html)<props="intl">[Documentation](https://www.alibabacloud.com/help/doc-detail/2708761.html)
	//
	// .
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "value":"2.2.2.2"
	//
	// }
	Data *CreateRecordRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The back-to-origin host policy. This parameter takes effect when the record type is CNAME. Settings for the host policy in back-to-origin requests. Two patterns are available:
	//
	// - **follow_hostname**: Follows the request host.
	//
	// - **follow_origin_domain**: Follows the Origin Domain Name.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// The HTTP port. By default, port 80 is supported.
	//
	// example:
	//
	// 80,2052,2082,2086,2095,8080,8880
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The HTTPS port. By default, port 443 is supported.
	//
	// example:
	//
	// 443,2053,2083,2087,2096,8443
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether to enable proxy acceleration for the record. Proxy acceleration can be enabled only for CNAME, A, or AAAA records. That is, the Type parameter is set to CNAME, A, or AAAA. Valid values:
	//
	// - **true**: Enables proxy acceleration.
	//
	// - **false**: Disables proxy acceleration.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The record name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The site ID. You can obtain the site ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The origin server type for a CNAME record. This parameter is required when you add a CNAME record. That is, the Type parameter is set to CNAME. Valid values:
	//
	// - **OSS**: OSS origin server.
	//
	// - **S3**: S3 origin server.
	//
	// - **LB**: Load balancing origin server.
	//
	// - **OP**: IPAM pool origin server.
	//
	// - **Domain**: Common domain name origin server.
	//
	// If this parameter is not specified or is left empty, the default value is Domain, which indicates a common domain name origin server.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time-to-live (TTL) of the record, in seconds. A value of 1 indicates that the TTL is automatically determined.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The DNS record type, such as **A/AAAA**, **CNAME**, or **TXT**.
	//
	// This parameter is required.
	//
	// example:
	//
	// A/AAAA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordRequest) GetAuthConf() *CreateRecordRequestAuthConf {
	return s.AuthConf
}

func (s *CreateRecordRequest) GetBizName() *string {
	return s.BizName
}

func (s *CreateRecordRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateRecordRequest) GetData() *CreateRecordRequestData {
	return s.Data
}

func (s *CreateRecordRequest) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *CreateRecordRequest) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *CreateRecordRequest) GetHttpsPorts() *string {
	return s.HttpsPorts
}

func (s *CreateRecordRequest) GetProxied() *bool {
	return s.Proxied
}

func (s *CreateRecordRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *CreateRecordRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateRecordRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateRecordRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateRecordRequest) GetType() *string {
	return s.Type
}

func (s *CreateRecordRequest) SetAuthConf(v *CreateRecordRequestAuthConf) *CreateRecordRequest {
	s.AuthConf = v
	return s
}

func (s *CreateRecordRequest) SetBizName(v string) *CreateRecordRequest {
	s.BizName = &v
	return s
}

func (s *CreateRecordRequest) SetComment(v string) *CreateRecordRequest {
	s.Comment = &v
	return s
}

func (s *CreateRecordRequest) SetData(v *CreateRecordRequestData) *CreateRecordRequest {
	s.Data = v
	return s
}

func (s *CreateRecordRequest) SetHostPolicy(v string) *CreateRecordRequest {
	s.HostPolicy = &v
	return s
}

func (s *CreateRecordRequest) SetHttpPorts(v string) *CreateRecordRequest {
	s.HttpPorts = &v
	return s
}

func (s *CreateRecordRequest) SetHttpsPorts(v string) *CreateRecordRequest {
	s.HttpsPorts = &v
	return s
}

func (s *CreateRecordRequest) SetProxied(v bool) *CreateRecordRequest {
	s.Proxied = &v
	return s
}

func (s *CreateRecordRequest) SetRecordName(v string) *CreateRecordRequest {
	s.RecordName = &v
	return s
}

func (s *CreateRecordRequest) SetSiteId(v int64) *CreateRecordRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRecordRequest) SetSourceType(v string) *CreateRecordRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRecordRequest) SetTtl(v int32) *CreateRecordRequest {
	s.Ttl = &v
	return s
}

func (s *CreateRecordRequest) SetType(v string) *CreateRecordRequest {
	s.Type = &v
	return s
}

func (s *CreateRecordRequest) Validate() error {
	if s.AuthConf != nil {
		if err := s.AuthConf.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRecordRequestAuthConf struct {
	// The AccessKey ID of the account to which the origin server belongs. This parameter is required when the origin server type is OSS and the authentication type is private read across accounts, or when the origin server type is S3 and the authentication type is private read.
	//
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The origin server authentication type. Different origin server types support different authentication types. The origin server type is specified by the SourceType parameter. This parameter is required when the origin server type is OSS or S3. Valid values:
	//
	// - **public**: Public read. Select this value when the origin server type is OSS or S3 and the origin server allows public read access.
	//
	// - **private**: Private read. Select this value when the origin server type is S3 and the origin server allows private read access.
	//
	// - **private_same_account**: Private read within the same account. Select this value when the origin server type is OSS, the origin server belongs to the same Alibaba Cloud account, and the origin server allows private read access.
	//
	// - **private_cross_account**: Private read across accounts. Select this value when the origin server type is OSS, the origin server belongs to a different Alibaba Cloud account, and the origin server allows private read access.
	//
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region to which the origin server belongs. This parameter is required when the origin server type is S3. You can obtain the region information from the official website of S3.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The AccessKey secret of the account to which the origin server belongs. This parameter is required when the origin server type is OSS and the authentication type is private read across accounts, or when the origin server type is S3 and the authentication type is private read.
	//
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signature algorithm version. This parameter is required when the origin server type is S3 and the authentication type is private read. Valid values:
	//
	// - **v2**
	//
	// - **v4**
	//
	// If this parameter is not specified, the default value is v4.
	//
	// example:
	//
	// v4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateRecordRequestAuthConf) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordRequestAuthConf) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateRecordRequestAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateRecordRequestAuthConf) GetRegion() *string {
	return s.Region
}

func (s *CreateRecordRequestAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *CreateRecordRequestAuthConf) GetVersion() *string {
	return s.Version
}

func (s *CreateRecordRequestAuthConf) SetAccessKey(v string) *CreateRecordRequestAuthConf {
	s.AccessKey = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetAuthType(v string) *CreateRecordRequestAuthConf {
	s.AuthType = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetRegion(v string) *CreateRecordRequestAuthConf {
	s.Region = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetSecretKey(v string) *CreateRecordRequestAuthConf {
	s.SecretKey = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetVersion(v string) *CreateRecordRequestAuthConf {
	s.Version = &v
	return s
}

func (s *CreateRecordRequestAuthConf) Validate() error {
	return dara.Validate(s)
}

type CreateRecordRequestData struct {
	// The encryption algorithm used by the record. Valid values: **0*	- to **255**. This parameter is required when you add a CERT or SSHFP record.
	//
	// example:
	//
	// 1
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key certificate information of the record. This parameter is required when you add a CERT, SMIMEA, or TLSA record.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint value of the record. This parameter is required when you add an SSHFP record.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The flag of the record. For a CAA record, the flag indicates its priority and processing method. Valid values: **0*	- to **255**. This parameter is required when you add a CAA record.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key tag of the record. Valid values: **0*	- to **65535**. This parameter is required when you add a CERT record.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used by the record to match or verify the certificate. Valid values: **0*	- to **255**. This parameter is required when you add an SMIMEA or TLSA record.
	//
	// example:
	//
	// 1
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port of the record. Valid values: **0*	- to **65535**. This parameter is required when you add an SRV record.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. Valid values: **0*	- to **65535**. A smaller value indicates a higher priority. This parameter is required when you add an MX, SRV, or URI record.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the certificate or public key used by the record. Valid values: **0*	- to **255**. This parameter is required when you add an SMIMEA or TLSA record.
	//
	// example:
	//
	// 1
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag of the record. For a CAA record, the tag indicates its specific type and purpose. This parameter is required when you add a CAA record. Valid values:
	//
	// - **issue**: Authorizes a specific CA to issue certificates for the domain name. This is typically followed by the domain name of the CA.
	//
	// - **issuewild**: Authorizes a specific CA to issue wildcard certificates for the domain name, such as *.example.com.
	//
	// - **iodef**: Specifies a URI to receive reports about CAA record violations.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type for a CERT record, or the public key type for an SSHFP record. This parameter is required when you add a CERT or SSHFP record.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier of the record. Valid values: **0*	- to **255**. This parameter is required when you add an SMIMEA or TLSA record.
	//
	// example:
	//
	// 1
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value or partial content. This parameter is required when the record type is A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, or URI. This parameter has different meanings for different record types:
	//
	// - **A/AAAA**: The destination IP address. Separate multiple IP addresses with commas (,). At least one IPv4 address is required.
	//
	// - **CNAME**: The destination domain name.
	//
	// - **NS**: The name server of the domain name.
	//
	// - **MX**: The domain name of the destination mail server.
	//
	// - **TXT**: A valid text string.
	//
	// - **CAA**: The domain name of the certification authority (CA).
	//
	// - **SRV**: The domain name of the destination host.
	//
	// - **URI**: A valid URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. Valid values: **0*	- to **65535**. This parameter is required when you add an SRV or URI record.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateRecordRequestData) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordRequestData) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *CreateRecordRequestData) GetCertificate() *string {
	return s.Certificate
}

func (s *CreateRecordRequestData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *CreateRecordRequestData) GetFlag() *int32 {
	return s.Flag
}

func (s *CreateRecordRequestData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *CreateRecordRequestData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *CreateRecordRequestData) GetPort() *int32 {
	return s.Port
}

func (s *CreateRecordRequestData) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateRecordRequestData) GetSelector() *int32 {
	return s.Selector
}

func (s *CreateRecordRequestData) GetTag() *string {
	return s.Tag
}

func (s *CreateRecordRequestData) GetType() *int32 {
	return s.Type
}

func (s *CreateRecordRequestData) GetUsage() *int32 {
	return s.Usage
}

func (s *CreateRecordRequestData) GetValue() *string {
	return s.Value
}

func (s *CreateRecordRequestData) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateRecordRequestData) SetAlgorithm(v int32) *CreateRecordRequestData {
	s.Algorithm = &v
	return s
}

func (s *CreateRecordRequestData) SetCertificate(v string) *CreateRecordRequestData {
	s.Certificate = &v
	return s
}

func (s *CreateRecordRequestData) SetFingerprint(v string) *CreateRecordRequestData {
	s.Fingerprint = &v
	return s
}

func (s *CreateRecordRequestData) SetFlag(v int32) *CreateRecordRequestData {
	s.Flag = &v
	return s
}

func (s *CreateRecordRequestData) SetKeyTag(v int32) *CreateRecordRequestData {
	s.KeyTag = &v
	return s
}

func (s *CreateRecordRequestData) SetMatchingType(v int32) *CreateRecordRequestData {
	s.MatchingType = &v
	return s
}

func (s *CreateRecordRequestData) SetPort(v int32) *CreateRecordRequestData {
	s.Port = &v
	return s
}

func (s *CreateRecordRequestData) SetPriority(v int32) *CreateRecordRequestData {
	s.Priority = &v
	return s
}

func (s *CreateRecordRequestData) SetSelector(v int32) *CreateRecordRequestData {
	s.Selector = &v
	return s
}

func (s *CreateRecordRequestData) SetTag(v string) *CreateRecordRequestData {
	s.Tag = &v
	return s
}

func (s *CreateRecordRequestData) SetType(v int32) *CreateRecordRequestData {
	s.Type = &v
	return s
}

func (s *CreateRecordRequestData) SetUsage(v int32) *CreateRecordRequestData {
	s.Usage = &v
	return s
}

func (s *CreateRecordRequestData) SetValue(v string) *CreateRecordRequestData {
	s.Value = &v
	return s
}

func (s *CreateRecordRequestData) SetWeight(v int32) *CreateRecordRequestData {
	s.Weight = &v
	return s
}

func (s *CreateRecordRequestData) Validate() error {
	return dara.Validate(s)
}
