// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordModel(v *GetRecordResponseBodyRecordModel) *GetRecordResponseBody
	GetRecordModel() *GetRecordResponseBodyRecordModel
	SetRequestId(v string) *GetRecordResponseBody
	GetRequestId() *string
}

type GetRecordResponseBody struct {
	// The information about the queried record.
	RecordModel *GetRecordResponseBodyRecordModel `json:"RecordModel,omitempty" xml:"RecordModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBody) GetRecordModel() *GetRecordResponseBodyRecordModel {
	return s.RecordModel
}

func (s *GetRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecordResponseBody) SetRecordModel(v *GetRecordResponseBodyRecordModel) *GetRecordResponseBody {
	s.RecordModel = v
	return s
}

func (s *GetRecordResponseBody) SetRequestId(v string) *GetRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRecordResponseBodyRecordModel struct {
	// The origin authentication information of the CNAME record.
	AuthConf *GetRecordResponseBodyRecordModelAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The business scenario of the record for acceleration. Leave this parameter empty if your record is not proxied. Valid values:
	//
	// 	- **image_video**
	//
	// 	- **api**
	//
	// 	- **web**
	//
	// example:
	//
	// image_video
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comments of the record.
	//
	// example:
	//
	// This is a comment.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the record was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-10T13:30:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The DNS record information. The content returned by this parameter varies based on the record type.
	//
	// example:
	//
	// {"value":"1.1.1.1"}
	Data *GetRecordResponseBodyRecordModelData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The origin host policy. This policy takes effect when the record type is CNAME. Valid values:
	//
	// 	- follow_hostname: matches the requested domain name.
	//
	// 	- follow_origin_domain: matches the origin\\"s domain name.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// Indicates whether the record is proxied. Only CNAME and A/AAAA records can be proxied. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The CNAME. If you use CNAME setup when you add your website to ESA, the value is the CNAME that you configured then.
	//
	// example:
	//
	// a.example.com.cnamezone.com
	RecordCname *string `json:"RecordCname,omitempty" xml:"RecordCname,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The record name.
	//
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
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
	// If you do not pass this parameter or if you leave its value empty, Domain is returned by default.
	//
	// example:
	//
	// OSS
	RecordSourceType *string `json:"RecordSourceType,omitempty" xml:"RecordSourceType,omitempty"`
	// The type of the DNS record, such as **A/AAAA, CNAME, and TXT**.
	//
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The TTL of the record. Unit: seconds. If the value is 1, the TTL of the record is determined by the system.
	//
	// example:
	//
	// 20
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The time when the record was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-27T02:26:22Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetRecordResponseBodyRecordModel) String() string {
	return dara.Prettify(s)
}

func (s GetRecordResponseBodyRecordModel) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBodyRecordModel) GetAuthConf() *GetRecordResponseBodyRecordModelAuthConf {
	return s.AuthConf
}

func (s *GetRecordResponseBodyRecordModel) GetBizName() *string {
	return s.BizName
}

func (s *GetRecordResponseBodyRecordModel) GetComment() *string {
	return s.Comment
}

func (s *GetRecordResponseBodyRecordModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRecordResponseBodyRecordModel) GetData() *GetRecordResponseBodyRecordModelData {
	return s.Data
}

func (s *GetRecordResponseBodyRecordModel) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *GetRecordResponseBodyRecordModel) GetProxied() *bool {
	return s.Proxied
}

func (s *GetRecordResponseBodyRecordModel) GetRecordCname() *string {
	return s.RecordCname
}

func (s *GetRecordResponseBodyRecordModel) GetRecordId() *int64 {
	return s.RecordId
}

func (s *GetRecordResponseBodyRecordModel) GetRecordName() *string {
	return s.RecordName
}

func (s *GetRecordResponseBodyRecordModel) GetRecordSourceType() *string {
	return s.RecordSourceType
}

func (s *GetRecordResponseBodyRecordModel) GetRecordType() *string {
	return s.RecordType
}

func (s *GetRecordResponseBodyRecordModel) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetRecordResponseBodyRecordModel) GetSiteName() *string {
	return s.SiteName
}

func (s *GetRecordResponseBodyRecordModel) GetTtl() *int32 {
	return s.Ttl
}

func (s *GetRecordResponseBodyRecordModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetRecordResponseBodyRecordModel) SetAuthConf(v *GetRecordResponseBodyRecordModelAuthConf) *GetRecordResponseBodyRecordModel {
	s.AuthConf = v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetBizName(v string) *GetRecordResponseBodyRecordModel {
	s.BizName = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetComment(v string) *GetRecordResponseBodyRecordModel {
	s.Comment = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetCreateTime(v string) *GetRecordResponseBodyRecordModel {
	s.CreateTime = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetData(v *GetRecordResponseBodyRecordModelData) *GetRecordResponseBodyRecordModel {
	s.Data = v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetHostPolicy(v string) *GetRecordResponseBodyRecordModel {
	s.HostPolicy = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetProxied(v bool) *GetRecordResponseBodyRecordModel {
	s.Proxied = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordCname(v string) *GetRecordResponseBodyRecordModel {
	s.RecordCname = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordId(v int64) *GetRecordResponseBodyRecordModel {
	s.RecordId = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordName(v string) *GetRecordResponseBodyRecordModel {
	s.RecordName = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordSourceType(v string) *GetRecordResponseBodyRecordModel {
	s.RecordSourceType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordType(v string) *GetRecordResponseBodyRecordModel {
	s.RecordType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetSiteId(v int64) *GetRecordResponseBodyRecordModel {
	s.SiteId = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetSiteName(v string) *GetRecordResponseBodyRecordModel {
	s.SiteName = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetTtl(v int32) *GetRecordResponseBodyRecordModel {
	s.Ttl = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetUpdateTime(v string) *GetRecordResponseBodyRecordModel {
	s.UpdateTime = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) Validate() error {
	return dara.Validate(s)
}

type GetRecordResponseBodyRecordModelAuthConf struct {
	// The access key ID of the account to which the origin server belongs. This parameter is returned if the origin type is OSS and AuthType is set to private_cross_account, or the origin type is S3 and AuthType is set to private.
	//
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication type of the origin server. Different origins support different authentication types. The origin type refers to the SourceType parameter in this operation. This parameter is returned if the origin type is OSS or S3Valid values:
	//
	// 	- **public**: public read. This value is returned when the origin is a public OSS or S3 bucket.
	//
	// 	- **private**: private read. This value is returned when the origin is a private S3 bucket.
	//
	// 	- **private_same_account**: private read in the same account. This value is returned when the origin is a private OSS bucket in your account.
	//
	// 	- **private_cross_account**: private read across accounts. This value is returned when the origin is a private OSS bucket in a different Alibaba Cloud account.
	//
	// example:
	//
	// public
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region of the origin. If the origin type is S3, you must specify this value. You can obtain the region information from the official website of S3.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The secret access key of the account to which the origin server belongs. This parameter is returned if the origin type is OSS and AuthType is set to private_cross_account, or the origin type is S3 and AuthType is set to private.SecretKey
	//
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The version of the signature algorithm. This parameter is returned when the origin type is S3 and AuthType is private. The following two types are supported:
	//
	// 	- **v2**
	//
	// 	- **v4**
	//
	// If this parameter is left empty, the default value v4 is used.
	//
	// example:
	//
	// v2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetRecordResponseBodyRecordModelAuthConf) String() string {
	return dara.Prettify(s)
}

func (s GetRecordResponseBodyRecordModelAuthConf) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetRegion() *string {
	return s.Region
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetVersion() *string {
	return s.Version
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetAccessKey(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.AccessKey = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetAuthType(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.AuthType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetRegion(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.Region = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetSecretKey(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.SecretKey = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetVersion(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.Version = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) Validate() error {
	return dara.Validate(s)
}

type GetRecordResponseBodyRecordModelData struct {
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
	// The public key identification for the record. Valid values: 0 to 65535. This parameter is required when you add a CAA record.
	//
	// example:
	//
	// 1
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used to match or validate the certificate, specified within the range 0 to 255. This parameter is required when you add SMIMEA or TLSA records.
	//
	// example:
	//
	// 1
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port of the record. Valid values: 0 to 65535. This parameter is required when you add an SRV record.
	//
	// example:
	//
	// 8707
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. Valid values: 0 to 65535. A smaller value indicates a higher priority. This parameter is required when you add MX, SRV, and URI records.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the certificate or public key, specified within the range of 0 to 255. This parameter is required when you add SMIMEA or TLSA records.
	//
	// example:
	//
	// 1
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag of the record. The Tag of a CAA record indicate its specific type and usage.
	//
	// example:
	//
	// issue
	Tag  *string                `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
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
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value or part of the record content. This parameter is returned when you add A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI records. It has different meanings based on types of records.
	//
	// 	- **A/AAAA**: the IP address. Multiple IP addresses are separated with commas (,). There is at least one IPv4 address.
	//
	// 	- **CNAME**: the target domain name.
	//
	// 	- **NS**: the nameserver for the domain name.
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
	// The weight of the record, specified within the range of 0 to 65535. This parameter is required when you add SRV or URI records.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetRecordResponseBodyRecordModelData) String() string {
	return dara.Prettify(s)
}

func (s GetRecordResponseBodyRecordModelData) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBodyRecordModelData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *GetRecordResponseBodyRecordModelData) GetCertificate() *string {
	return s.Certificate
}

func (s *GetRecordResponseBodyRecordModelData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *GetRecordResponseBodyRecordModelData) GetFlag() *int32 {
	return s.Flag
}

func (s *GetRecordResponseBodyRecordModelData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *GetRecordResponseBodyRecordModelData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *GetRecordResponseBodyRecordModelData) GetPort() *int32 {
	return s.Port
}

func (s *GetRecordResponseBodyRecordModelData) GetPriority() *int32 {
	return s.Priority
}

func (s *GetRecordResponseBodyRecordModelData) GetSelector() *int32 {
	return s.Selector
}

func (s *GetRecordResponseBodyRecordModelData) GetTag() *string {
	return s.Tag
}

func (s *GetRecordResponseBodyRecordModelData) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetRecordResponseBodyRecordModelData) GetType() *int32 {
	return s.Type
}

func (s *GetRecordResponseBodyRecordModelData) GetUsage() *int32 {
	return s.Usage
}

func (s *GetRecordResponseBodyRecordModelData) GetValue() *string {
	return s.Value
}

func (s *GetRecordResponseBodyRecordModelData) GetWeight() *int32 {
	return s.Weight
}

func (s *GetRecordResponseBodyRecordModelData) SetAlgorithm(v int32) *GetRecordResponseBodyRecordModelData {
	s.Algorithm = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetCertificate(v string) *GetRecordResponseBodyRecordModelData {
	s.Certificate = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetFingerprint(v string) *GetRecordResponseBodyRecordModelData {
	s.Fingerprint = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetFlag(v int32) *GetRecordResponseBodyRecordModelData {
	s.Flag = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetKeyTag(v int32) *GetRecordResponseBodyRecordModelData {
	s.KeyTag = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetMatchingType(v int32) *GetRecordResponseBodyRecordModelData {
	s.MatchingType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetPort(v int32) *GetRecordResponseBodyRecordModelData {
	s.Port = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetPriority(v int32) *GetRecordResponseBodyRecordModelData {
	s.Priority = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetSelector(v int32) *GetRecordResponseBodyRecordModelData {
	s.Selector = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetTag(v string) *GetRecordResponseBodyRecordModelData {
	s.Tag = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetTags(v map[string]interface{}) *GetRecordResponseBodyRecordModelData {
	s.Tags = v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetType(v int32) *GetRecordResponseBodyRecordModelData {
	s.Type = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetUsage(v int32) *GetRecordResponseBodyRecordModelData {
	s.Usage = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetValue(v string) *GetRecordResponseBodyRecordModelData {
	s.Value = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetWeight(v int32) *GetRecordResponseBodyRecordModelData {
	s.Weight = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) Validate() error {
	return dara.Validate(s)
}
