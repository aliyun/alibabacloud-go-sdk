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
	// The origin authentication information of the CNAME record.
	AuthConf *CreateRecordRequestAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The business scenario of the record for acceleration. Leave the parameter empty if your record is not proxied. Valid values:
	//
	// 	- **image_video**: video and image.
	//
	// 	- **api**: API.
	//
	// 	- **web**: web page.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comment of the record. The maximum length is 100 characters.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS record information. The format of this field varies based on the record type. For more information, see [References](https://www.alibabacloud.com/help/doc-detail/2708761.html) .
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
	// The origin host policy. This policy takes effect when the record type is CNAME. You can set the policy in two modes:
	//
	// 	- follow_hostname: Follow the host record.
	//
	// 	- follow_origin_domain: match the origin\\"s domain name.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// Specifies whether to proxy the record. Only CNAME and A/AAAA records can be proxied. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The origin type for the CNAME record. This parameter is required when you add a CNAME record. Valid values:
	//
	// 	- **OSS**: OSS bucket.
	//
	// 	- **S3**: S3 bucket.
	//
	// 	- **LB**: load balancer.
	//
	// 	- **OP**: origin pool.
	//
	// 	- **Domain**: domain name.
	//
	// If you do not pass this parameter or if you leave its value empty, Domain is used by default.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The TTL of the record. Unit: seconds. If the value is 1, the TTL of the record is determined by the system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. For example, A/AAAA, TXT, MX, or CNAME.
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
	return dara.Validate(s)
}

type CreateRecordRequestAuthConf struct {
	// The access key of the account to which the origin server belongs. This parameter is required when the SourceType is OSS, and AuthType is private_cross_account, or when the SourceType is S3 and AuthType is private.
	//
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication type of the origin server. Different origins support different authentication types. The type of origin refers to the SourceType parameter in this operation. If the type of origin is OSS or S3, you must specify the authentication type of the origin. Valid values:
	//
	// 	- **public**: public read. Select this value when the origin type is OSS or S3 and the origin access is public read.
	//
	// 	- **private**: private read. Select this value when the origin type is S3 and the origin access is private read.
	//
	// 	- **private_same_account**: private read under the same account. Select this value when the origin type is OSS, the origins belong to the same Alibaba Cloud account, and the origins have private read access.
	//
	// 	- **private_cross_account**: private read cross accounts. Select this value when the origin type is OSS, the origins belong to different Alibaba Cloud accounts, and the origins have private read access.
	//
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region of the origin. If the origin type is S3, you must specify this value. You can get the region information from the official website of S3.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The secret access key of the account to which the origin server belongs. This parameter is required when the SourceType is OSS, and AuthType is private_same_account, or when the SourceType is S3 and AuthType is private.
	//
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The version of the signature algorithm. This parameter is required when the origin type is S3 and AuthType is private. The following two types are supported:
	//
	// 	- **v2**
	//
	// 	- **v4**
	//
	// If you leave this parameter empty, the default value v4 is used.
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
	// The encryption algorithm used for the record, specified within the range from 0 to 255. This parameter is required when you add CERT or SSHFP records.
	//
	// example:
	//
	// 1
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key of the certificate. This parameter is required when you add CERT, SMIMEA, or TLSA records.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint of the record. This parameter is required when you add a SSHFP record.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The flag bit of the record. The Flag for a CAA record indicates its priority and how it is processed, specified within the range of 0 to 255. This parameter is required when you add a CAA record.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identification for the record, specified within the range of 0 to 65,535. This parameter is required when you add a CAA record.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used to match or validate the certificate, specified within the range 0 to 255. This parameter is required when you add SMIMEA or TLSA records.
	//
	// example:
	//
	// 1
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port of the record, specified within the range of 0 to 65,535. This parameter is required when you add an SRV record.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record, specified within the range of 0 to 65,535. A smaller value indicates a higher priority. This parameter is required when you add MX, SRV, and URI records.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key, specified within the range of 0 to 255. This parameter is required when you add SMIMEA or TLSA records.
	//
	// example:
	//
	// 1
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The label of the record. The Tag of a CAA record indicate its specific type and usage. This parameter is required when you add a CAA record. Valid values:
	//
	// 	- **issue**: indicates that a CA is authorized to issue a certificate for the domain name. This is usually followed by the domain name of the CA.
	//
	// 	- **issuewild**: indicates that a CA is authorized to issue a wildcard certificate (such as \\*.example.com) for the domain name.
	//
	// 	- **iodef**: specifies a URI to receive reports about CAA record violations.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type of the record (in CERT records), or the public key type (in SSHFP records). This parameter is required when you add CERT or SSHFP records.
	//
	// example:
	//
	// RSA
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier of the record, specified within the range of 0 to 255. This parameter is required when you add SMIMEA or TLSA records.
	//
	// example:
	//
	// 1
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// Record value or part of the record content. This parameter is required when you add A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI records. It has different meanings based on types of records:
	//
	// 	- **A/AAAA**: the IP address(es). Separate IP addresses with commas (,). You must have at least one IPv4 address.
	//
	// 	- **CNAME**: the target domain name.
	//
	// 	- **NS**: the name servers for the domain name.
	//
	// 	- **MX**: a valid domain name of the target mail server.
	//
	// 	- **TXT**: a valid text string.
	//
	// 	- **CAA**: a valid domain name of the certificate authority.
	//
	// 	- **SRV**: a valid domain name of the target host.
	//
	// 	- **URI**: a valid URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record, specified within the range of 0 to 65,535. This parameter is required when you add SRV or URI records.
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
