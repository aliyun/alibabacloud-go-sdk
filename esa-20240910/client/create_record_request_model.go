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
	// The origin authentication information for the CNAME record.
	AuthConf *CreateRecordRequestAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// Used to tag the business scenario of the DNS record. This parameter is required when proxy acceleration is enabled for the DNS record (i.e., when the proxied parameter is set to true), and is not required when proxy acceleration is disabled (i.e., when the proxied parameter is set to false). Valid values:
	//
	// - **image_video**: Image and video.
	//
	// - **api**: API.
	//
	// - **web**: Web page.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comment for the record, with a maximum length of 100 characters.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS information of the record. Different types of records require different content for this field. For more information, see
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
	// The origin host policy. This takes effect when the record type is CNAME. It specifies the host policy for back-to-origin requests. Two modes are available:
	//
	// - **follow_hostname**: Follow the request host.
	//
	// - **follow_origin_domain**: Follow the origin domain.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	HttpPorts  *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether to enable proxy acceleration for the record. Only CNAME records or A/AAAA records (i.e., when the type parameter is set to A/AAAA or CNAME) can enable proxy acceleration. Valid values:
	//
	// - **true**: Enable proxy acceleration.
	//
	// - **false**: Disable proxy acceleration.
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
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The origin type of the CNAME record. This parameter is required when adding a CNAME record (i.e., when the type parameter is set to CNAME). Valid values:
	//
	// - **OSS**: OSS origin.
	//
	// - **S3**: S3 origin.
	//
	// - **LB**: Load balancer origin.
	//
	// - **OP**: Origin pool origin.
	//
	// - **Domain**: Standard domain origin.
	//
	// If this parameter is not specified or is left empty, it defaults to Domain, which is the standard domain origin type.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time-to-live (TTL) of the record, in seconds. When set to 1, the TTL is automatic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The DNS type of the record, such as **A/AAAA**, **CNAME**, **TXT**, etc.
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
	// The AccessKey of the account that owns the origin. This value is required when the origin type is OSS and the authentication type is private cross-account read, or when the origin type is S3 and the authentication type is private read.
	//
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The origin authentication type. Different origin types support different authentication types. The origin type refers to the SourceType parameter in this API. When the origin type is OSS or S3, you need to specify the origin authentication type. Valid values:
	//
	// - **public**: Public read. Select this value when the origin type is OSS or S3 and the origin has public read access.
	//
	// - **private**: Private read. Select this value when the origin type is S3 and the origin has private read access.
	//
	// - **private_same_account**: Private same-account read. Select this value when the origin type is OSS, the origin is under the same Alibaba Cloud account, and the origin has private read access.
	//
	// - **private_cross_account**: Private cross-account read. Select this value when the origin type is OSS, the origin is not under the same Alibaba Cloud account, and the origin has private read access.
	//
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region where the origin is located. This value is required when the origin type is S3. The region information can be obtained from the official S3 website.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The SecretKey of the account that owns the origin. This value is required when the origin type is OSS and the authentication type is private cross-account read, or when the origin type is S3 and the authentication type is private read.
	//
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signature algorithm version. This is required when the origin type is S3 and the authentication type is private read. The following two versions are supported:
	//
	// - **v2**
	//
	// - **v4**
	//
	// If not specified, the default value is v4.
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
	// The encryption algorithm used by the record, ranging from **0 to 255**. This field is required when adding CERT or SSHFP records.
	//
	// example:
	//
	// 1
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key certificate information of the record. This parameter is required when adding CERT, SMIMEA, or TLSA records.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint value of the record. This parameter is required when adding an SSHFP record.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The flag of the record. The Flag of a CAA record indicates its priority and processing method, with a value range of **0 to 255**. This parameter is required when adding a CAA record.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identifier of the record, ranging from **0 to 65535**. This parameter is required when adding a CERT record.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used by the record to match or verify certificates, ranging from **0 to 255**. This parameter is required when adding SMIMEA or TLSA records.
	//
	// example:
	//
	// 1
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port of the record, ranging from **0 to 65535**. This parameter is required when adding an SRV record.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record, ranging from **0 to 65535**. A smaller value indicates a higher priority. This parameter is required when adding MX, SRV, or URI records.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key used by the record, ranging from **0 to 255**. This parameter is required when adding SMIMEA or TLSA records.
	//
	// example:
	//
	// 1
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag of the record. The Tag of a CAA record indicates its specific type and purpose. This parameter is required when adding a CAA record. Valid values for Tag:
	//
	// - **issue**: Authorizes a specific CA to issue certificates for the domain. It is usually followed by the CA\\"s domain name.
	//
	// - **issuewild**: Authorizes a specific CA to issue wildcard certificates for the domain (e.g., *.example.com).
	//
	// - **iodef**: Specifies a URI for receiving reports about violations of CAA records.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type of the record (for CERT records) or the public key type (for SSHFP records). This parameter is required when adding CERT or SSHFP records.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier of the record, ranging from **0 to 255**. This parameter is required when adding SMIMEA or TLSA records.
	//
	// example:
	//
	// 1
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value or partial content. This parameter is required when the record type is A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, or URI. It represents different meanings for different record types:
	//
	// - **A/AAAA**: The IP address to point to. Multiple IPs are separated by commas (,). At least one IPv4 address is required.
	//
	// - **CNAME**: The target domain name to point to.
	//
	// - **NS**: The name server for the specified domain.
	//
	// - **MX**: A valid target mail server domain name.
	//
	// - **TXT**: A valid text string.
	//
	// - **CAA**: A valid certificate authority domain name.
	//
	// - **SRV**: A valid target host domain name.
	//
	// - **URI**: A valid URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record, ranging from **0 to 65535**. This parameter is required when adding SRV or URI records.
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
