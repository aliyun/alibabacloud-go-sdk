// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConf(v *UpdateRecordRequestAuthConf) *UpdateRecordRequest
	GetAuthConf() *UpdateRecordRequestAuthConf
	SetBizName(v string) *UpdateRecordRequest
	GetBizName() *string
	SetComment(v string) *UpdateRecordRequest
	GetComment() *string
	SetData(v *UpdateRecordRequestData) *UpdateRecordRequest
	GetData() *UpdateRecordRequestData
	SetHostPolicy(v string) *UpdateRecordRequest
	GetHostPolicy() *string
	SetProxied(v bool) *UpdateRecordRequest
	GetProxied() *bool
	SetRecordId(v int64) *UpdateRecordRequest
	GetRecordId() *int64
	SetSourceType(v string) *UpdateRecordRequest
	GetSourceType() *string
	SetTtl(v int32) *UpdateRecordRequest
	GetTtl() *int32
	SetType(v string) *UpdateRecordRequest
	GetType() *string
}

type UpdateRecordRequest struct {
	// The origin authentication information of the CNAME record.
	AuthConf *UpdateRecordRequestAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The business scenario of the record for acceleration. Leave the parameter empty if your record is not proxied. Valid values:
	//
	// 	- **video_image**: video and image.
	//
	// 	- **api**: API.
	//
	// 	- **web**: web page.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comments of the record.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS record information. The format of this field varies based on the record type. For more information, see [Add DNS records](https://www.alibabacloud.com/help/doc-detail/2708761.html).
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
	Data *UpdateRecordRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The origin host policy. This policy takes effect when the record type is CNAME. You can set the policy in two modes:
	//
	// 	- **follow_hostname**: match the requested domain name.
	//
	// 	- **follow_origin_domain**: match the origin\\"s domain name.
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
	// The record ID, which can be obtained by calling [ListRecords](https://help.aliyun.com/document_detail/2850265.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The type of the origin for the CNAME record. This parameter is required when you add a CNAME record. Valid values:
	//
	// 	- **OSS*	- : OSS origin.
	//
	// 	- **S3*	- : S3 origin.
	//
	// 	- **LB**: Load Balancer origin.
	//
	// 	- **OP**: origin in an origin pool.
	//
	// 	- **Domain**: common domain name.
	//
	// If you leave the parameter empty or set its value as null, the default is Domain, which is common domain name.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The TTL of the record. Unit: seconds. The range is 30 to 86,400, or 1. If the value is 1, the TTL of the record is determined by the system.
	//
	// example:
	//
	// 30
	Ttl  *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordRequest) GetAuthConf() *UpdateRecordRequestAuthConf {
	return s.AuthConf
}

func (s *UpdateRecordRequest) GetBizName() *string {
	return s.BizName
}

func (s *UpdateRecordRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateRecordRequest) GetData() *UpdateRecordRequestData {
	return s.Data
}

func (s *UpdateRecordRequest) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *UpdateRecordRequest) GetProxied() *bool {
	return s.Proxied
}

func (s *UpdateRecordRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateRecordRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateRecordRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateRecordRequest) GetType() *string {
	return s.Type
}

func (s *UpdateRecordRequest) SetAuthConf(v *UpdateRecordRequestAuthConf) *UpdateRecordRequest {
	s.AuthConf = v
	return s
}

func (s *UpdateRecordRequest) SetBizName(v string) *UpdateRecordRequest {
	s.BizName = &v
	return s
}

func (s *UpdateRecordRequest) SetComment(v string) *UpdateRecordRequest {
	s.Comment = &v
	return s
}

func (s *UpdateRecordRequest) SetData(v *UpdateRecordRequestData) *UpdateRecordRequest {
	s.Data = v
	return s
}

func (s *UpdateRecordRequest) SetHostPolicy(v string) *UpdateRecordRequest {
	s.HostPolicy = &v
	return s
}

func (s *UpdateRecordRequest) SetProxied(v bool) *UpdateRecordRequest {
	s.Proxied = &v
	return s
}

func (s *UpdateRecordRequest) SetRecordId(v int64) *UpdateRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordRequest) SetSourceType(v string) *UpdateRecordRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateRecordRequest) SetTtl(v int32) *UpdateRecordRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateRecordRequest) SetType(v string) *UpdateRecordRequest {
	s.Type = &v
	return s
}

func (s *UpdateRecordRequest) Validate() error {
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

type UpdateRecordRequestAuthConf struct {
	// The access key of the account to which the origin server belongs. This parameter is required when the SourceType is OSS, and AuthType is private_same_account, or when the SourceType is S3 and AuthType is private.
	//
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication type of the origin server. Different origins support different authentication types. The type of origin refers to the SourceType parameter in this operation. If the type of origin is OSS or S3, you must specify the authentication type of the origin. Valid values:
	//
	// 	- **public**: public read. Select this value when the origin type is OSS or S3 and the origin access is public read.
	//
	// 	- **private**: private read. Select this value when the origin type is S3 and the origin access is private read.
	//
	// 	- **private_same_account**: private read under the same account. Select this value when the origin type is OSS, the origins belong to the same Alibaba Cloud account, and the origins have private read access.
	//
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
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
	// v2
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The secret access key of the account to which the origin server belongs. This parameter is required when the SourceType is OSS, and AuthType is private_same_account, or when the SourceType is S3 and AuthType is private.
	//
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The region of the origin. If the origin type is S3, you must specify this value. You can get the region information from the official website of S3.
	//
	// example:
	//
	// us-east-1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateRecordRequestAuthConf) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordRequestAuthConf) GoString() string {
	return s.String()
}

func (s *UpdateRecordRequestAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *UpdateRecordRequestAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateRecordRequestAuthConf) GetRegion() *string {
	return s.Region
}

func (s *UpdateRecordRequestAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *UpdateRecordRequestAuthConf) GetVersion() *string {
	return s.Version
}

func (s *UpdateRecordRequestAuthConf) SetAccessKey(v string) *UpdateRecordRequestAuthConf {
	s.AccessKey = &v
	return s
}

func (s *UpdateRecordRequestAuthConf) SetAuthType(v string) *UpdateRecordRequestAuthConf {
	s.AuthType = &v
	return s
}

func (s *UpdateRecordRequestAuthConf) SetRegion(v string) *UpdateRecordRequestAuthConf {
	s.Region = &v
	return s
}

func (s *UpdateRecordRequestAuthConf) SetSecretKey(v string) *UpdateRecordRequestAuthConf {
	s.SecretKey = &v
	return s
}

func (s *UpdateRecordRequestAuthConf) SetVersion(v string) *UpdateRecordRequestAuthConf {
	s.Version = &v
	return s
}

func (s *UpdateRecordRequestAuthConf) Validate() error {
	return dara.Validate(s)
}

type UpdateRecordRequestData struct {
	// The encryption algorithm used for the record, specified within the range from 0 to 255. This parameter is required when you add CERT or SSHFP records.
	//
	// example:
	//
	// 0
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
	// 0
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
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The label of the record. The Tag of a CAA record indicate its specific type and usage. This parameter is required when you add a CAA record.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type of the record (in CERT records), or the public key type (in SSHFP records). This parameter is required when you add CERT or SSHFP records.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier of the record, specified within the range of 0 to 255. This parameter is required when you add SMIMEA or TLSA records.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value or part of the record content. This parameter is required when you add A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI records. It has different meanings based on different types of records:
	//
	// 	- **A/AAAA**: the IP address(es). Separate multiple IPs with commas (,). You must have at least one IPv4 address.
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

func (s UpdateRecordRequestData) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordRequestData) GoString() string {
	return s.String()
}

func (s *UpdateRecordRequestData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *UpdateRecordRequestData) GetCertificate() *string {
	return s.Certificate
}

func (s *UpdateRecordRequestData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *UpdateRecordRequestData) GetFlag() *int32 {
	return s.Flag
}

func (s *UpdateRecordRequestData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *UpdateRecordRequestData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *UpdateRecordRequestData) GetPort() *int32 {
	return s.Port
}

func (s *UpdateRecordRequestData) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateRecordRequestData) GetSelector() *int32 {
	return s.Selector
}

func (s *UpdateRecordRequestData) GetTag() *string {
	return s.Tag
}

func (s *UpdateRecordRequestData) GetType() *int32 {
	return s.Type
}

func (s *UpdateRecordRequestData) GetUsage() *int32 {
	return s.Usage
}

func (s *UpdateRecordRequestData) GetValue() *string {
	return s.Value
}

func (s *UpdateRecordRequestData) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateRecordRequestData) SetAlgorithm(v int32) *UpdateRecordRequestData {
	s.Algorithm = &v
	return s
}

func (s *UpdateRecordRequestData) SetCertificate(v string) *UpdateRecordRequestData {
	s.Certificate = &v
	return s
}

func (s *UpdateRecordRequestData) SetFingerprint(v string) *UpdateRecordRequestData {
	s.Fingerprint = &v
	return s
}

func (s *UpdateRecordRequestData) SetFlag(v int32) *UpdateRecordRequestData {
	s.Flag = &v
	return s
}

func (s *UpdateRecordRequestData) SetKeyTag(v int32) *UpdateRecordRequestData {
	s.KeyTag = &v
	return s
}

func (s *UpdateRecordRequestData) SetMatchingType(v int32) *UpdateRecordRequestData {
	s.MatchingType = &v
	return s
}

func (s *UpdateRecordRequestData) SetPort(v int32) *UpdateRecordRequestData {
	s.Port = &v
	return s
}

func (s *UpdateRecordRequestData) SetPriority(v int32) *UpdateRecordRequestData {
	s.Priority = &v
	return s
}

func (s *UpdateRecordRequestData) SetSelector(v int32) *UpdateRecordRequestData {
	s.Selector = &v
	return s
}

func (s *UpdateRecordRequestData) SetTag(v string) *UpdateRecordRequestData {
	s.Tag = &v
	return s
}

func (s *UpdateRecordRequestData) SetType(v int32) *UpdateRecordRequestData {
	s.Type = &v
	return s
}

func (s *UpdateRecordRequestData) SetUsage(v int32) *UpdateRecordRequestData {
	s.Usage = &v
	return s
}

func (s *UpdateRecordRequestData) SetValue(v string) *UpdateRecordRequestData {
	s.Value = &v
	return s
}

func (s *UpdateRecordRequestData) SetWeight(v int32) *UpdateRecordRequestData {
	s.Weight = &v
	return s
}

func (s *UpdateRecordRequestData) Validate() error {
	return dara.Validate(s)
}
