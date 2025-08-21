// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecordsResponseBody
	GetPageSize() *int32
	SetRecords(v []*ListRecordsResponseBodyRecords) *ListRecordsResponseBody
	GetRecords() []*ListRecordsResponseBodyRecords
	SetRequestId(v string) *ListRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRecordsResponseBody
	GetTotalCount() *int32
}

type ListRecordsResponseBody struct {
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DNS record information. The format of this field varies based on the record type. For more information, see Add DNS records.
	Records []*ListRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecordsResponseBody) GetRecords() []*ListRecordsResponseBodyRecords {
	return s.Records
}

func (s *ListRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRecordsResponseBody) SetPageNumber(v int32) *ListRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecordsResponseBody) SetPageSize(v int32) *ListRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecordsResponseBody) SetRecords(v []*ListRecordsResponseBodyRecords) *ListRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListRecordsResponseBody) SetRequestId(v string) *ListRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecordsResponseBody) SetTotalCount(v int32) *ListRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRecordsResponseBodyRecords struct {
	// The origin authentication information of the CNAME record.
	AuthConf *ListRecordsResponseBodyRecordsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The business scenario of the record for acceleration. Valid values:
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
	// The comments of the record.
	//
	// example:
	//
	// this is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the record was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The DNS record information. Different types of records contain different information.
	//
	// example:
	//
	// {"value":"1.1.1.1"}
	Data *ListRecordsResponseBodyRecordsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the record is proxied. Valid values:
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
	// The origin type for the CNAME record. This parameter is returned when you add a CNAME record. Valid values:
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
	// The DNS type of the record, such as **A/AAAA, CNAME, and TXT**.
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
	// 30
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The time when the record was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-07T10:02:59Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListRecordsResponseBodyRecords) GetAuthConf() *ListRecordsResponseBodyRecordsAuthConf {
	return s.AuthConf
}

func (s *ListRecordsResponseBodyRecords) GetBizName() *string {
	return s.BizName
}

func (s *ListRecordsResponseBodyRecords) GetComment() *string {
	return s.Comment
}

func (s *ListRecordsResponseBodyRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRecordsResponseBodyRecords) GetData() *ListRecordsResponseBodyRecordsData {
	return s.Data
}

func (s *ListRecordsResponseBodyRecords) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *ListRecordsResponseBodyRecords) GetProxied() *bool {
	return s.Proxied
}

func (s *ListRecordsResponseBodyRecords) GetRecordCname() *string {
	return s.RecordCname
}

func (s *ListRecordsResponseBodyRecords) GetRecordId() *int64 {
	return s.RecordId
}

func (s *ListRecordsResponseBodyRecords) GetRecordName() *string {
	return s.RecordName
}

func (s *ListRecordsResponseBodyRecords) GetRecordSourceType() *string {
	return s.RecordSourceType
}

func (s *ListRecordsResponseBodyRecords) GetRecordType() *string {
	return s.RecordType
}

func (s *ListRecordsResponseBodyRecords) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListRecordsResponseBodyRecords) GetSiteName() *string {
	return s.SiteName
}

func (s *ListRecordsResponseBodyRecords) GetTtl() *int64 {
	return s.Ttl
}

func (s *ListRecordsResponseBodyRecords) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRecordsResponseBodyRecords) SetAuthConf(v *ListRecordsResponseBodyRecordsAuthConf) *ListRecordsResponseBodyRecords {
	s.AuthConf = v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetBizName(v string) *ListRecordsResponseBodyRecords {
	s.BizName = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetComment(v string) *ListRecordsResponseBodyRecords {
	s.Comment = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetCreateTime(v string) *ListRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetData(v *ListRecordsResponseBodyRecordsData) *ListRecordsResponseBodyRecords {
	s.Data = v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetHostPolicy(v string) *ListRecordsResponseBodyRecords {
	s.HostPolicy = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetProxied(v bool) *ListRecordsResponseBodyRecords {
	s.Proxied = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordCname(v string) *ListRecordsResponseBodyRecords {
	s.RecordCname = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordId(v int64) *ListRecordsResponseBodyRecords {
	s.RecordId = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordName(v string) *ListRecordsResponseBodyRecords {
	s.RecordName = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordSourceType(v string) *ListRecordsResponseBodyRecords {
	s.RecordSourceType = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordType(v string) *ListRecordsResponseBodyRecords {
	s.RecordType = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetSiteId(v int64) *ListRecordsResponseBodyRecords {
	s.SiteId = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetSiteName(v string) *ListRecordsResponseBodyRecords {
	s.SiteName = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetTtl(v int64) *ListRecordsResponseBodyRecords {
	s.Ttl = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetUpdateTime(v string) *ListRecordsResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}

type ListRecordsResponseBodyRecordsAuthConf struct {
	// The access key of the account to which the origin server belongs. This value is returned when the SourceType is OSS, and AuthType is private_cross_account, or when the SourceType is S3 and AuthType is private.
	//
	// example:
	//
	// u0Nkg5gBK***QF5wvKMM504JUHt
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication type of the origin server. Different origins support different authentication types. The type of origin refers to the SourceType parameter in this operation. If the type of origin is OSS or S3, the authentication type of the origin must be specified. Valid values:
	//
	// 	- **public**: public read. This value is returned when the origin is a public OSS or S3 bucket.
	//
	// 	- **private**: private read. This value is returned when the origin is a private S3 bucket.
	//
	// 	- **private_same_account**: private read under the same account. This value is returned when the origin is a private OSS bucket in your Alibaba Cloud account.
	//
	// 	- **private_cross_account**: private read across accounts. This value is returned when the origin is a private OSS bucket in a different Alibaba Cloud account.
	//
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region of the origin. This parameter is returned if the origin type is S3. You can get the region information from the official website of Amazon S3.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The secret access key of the account to which the origin server belongs. This value is returned when the SourceType is OSS, and AuthType is private_same_account, or when the SourceType is S3 and AuthType is private.
	//
	// example:
	//
	// VIxuvJSA2S03f***kp208dy5w7
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The version of the signature algorithm. This value is returned when the origin type is S3 and AuthType is private. Valid values:
	//
	// 	- **v2**
	//
	// 	- **v4**
	//
	// If this parameter is left empty, the default value v4 is used.
	//
	// example:
	//
	// v4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListRecordsResponseBodyRecordsAuthConf) String() string {
	return dara.Prettify(s)
}

func (s ListRecordsResponseBodyRecordsAuthConf) GoString() string {
	return s.String()
}

func (s *ListRecordsResponseBodyRecordsAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ListRecordsResponseBodyRecordsAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *ListRecordsResponseBodyRecordsAuthConf) GetRegion() *string {
	return s.Region
}

func (s *ListRecordsResponseBodyRecordsAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *ListRecordsResponseBodyRecordsAuthConf) GetVersion() *string {
	return s.Version
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetAccessKey(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.AccessKey = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetAuthType(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.AuthType = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetRegion(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.Region = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetSecretKey(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.SecretKey = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetVersion(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.Version = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsAuthConf) Validate() error {
	return dara.Validate(s)
}

type ListRecordsResponseBodyRecordsData struct {
	// The encryption algorithm used for the record. Valid values: 0 to 255. Exclusive to CERT and SSHFP records.
	//
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key of the certificate. Exclusive to CERT, SMIMEA, and TLSA records.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint of the record. Exclusive to SSHFP records.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The flag bit of the record. The Flag for a CAA record indicates its priority and how it is processed. Valid values: 0 to 255. Exclusive to CAA records.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identification for the record. Valid values: 0 to 65535. Exclusive to CERT records.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used to match or validate the certificate. Valid values: 0 to 255. Exclusive to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port of the record. Valid values: 0 to 65535. Exclusive to SRV records.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. Valid values: 0 to 65535. A smaller value indicates a higher priority. Exclusive to MX, SRV, and URI records.
	//
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key. Valid values: 0 to 255. Exclusive to SMIMEA, and TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag of the record. The Tag of a CAA record indicate its specific type and usage. Exclusive to CAA records.
	//
	// example:
	//
	// issue
	Tag  *string                `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The certificate type of the record (in CERT records), or the public key type (in SSHFP records).
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier of the record. Valid values: 0 to 255. Exclusive to SMIMEA, and TLSA records.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// Record value or part of the record content. This value is returned when the record is A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, or URI. It has different meanings based on types of records:
	//
	// 	- **A/AAAA**: the IP addresses. IP addresses are separated by commas (,). There is at least one IPv4 address.
	//
	// 	- **CNAME**: the pointed/mapped domain name.
	//
	// 	- **NS**: the nameservers for the domain name.
	//
	// 	- **MX**: a valid domain name of the mail server.
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
	// CNAME
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. Valid values: 0 to 65535. Exclusive to SRV and URI records.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListRecordsResponseBodyRecordsData) String() string {
	return dara.Prettify(s)
}

func (s ListRecordsResponseBodyRecordsData) GoString() string {
	return s.String()
}

func (s *ListRecordsResponseBodyRecordsData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *ListRecordsResponseBodyRecordsData) GetCertificate() *string {
	return s.Certificate
}

func (s *ListRecordsResponseBodyRecordsData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *ListRecordsResponseBodyRecordsData) GetFlag() *int32 {
	return s.Flag
}

func (s *ListRecordsResponseBodyRecordsData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *ListRecordsResponseBodyRecordsData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *ListRecordsResponseBodyRecordsData) GetPort() *int32 {
	return s.Port
}

func (s *ListRecordsResponseBodyRecordsData) GetPriority() *int32 {
	return s.Priority
}

func (s *ListRecordsResponseBodyRecordsData) GetSelector() *int32 {
	return s.Selector
}

func (s *ListRecordsResponseBodyRecordsData) GetTag() *string {
	return s.Tag
}

func (s *ListRecordsResponseBodyRecordsData) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListRecordsResponseBodyRecordsData) GetType() *int32 {
	return s.Type
}

func (s *ListRecordsResponseBodyRecordsData) GetUsage() *int32 {
	return s.Usage
}

func (s *ListRecordsResponseBodyRecordsData) GetValue() *string {
	return s.Value
}

func (s *ListRecordsResponseBodyRecordsData) GetWeight() *int32 {
	return s.Weight
}

func (s *ListRecordsResponseBodyRecordsData) SetAlgorithm(v int32) *ListRecordsResponseBodyRecordsData {
	s.Algorithm = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetCertificate(v string) *ListRecordsResponseBodyRecordsData {
	s.Certificate = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetFingerprint(v string) *ListRecordsResponseBodyRecordsData {
	s.Fingerprint = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetFlag(v int32) *ListRecordsResponseBodyRecordsData {
	s.Flag = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetKeyTag(v int32) *ListRecordsResponseBodyRecordsData {
	s.KeyTag = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetMatchingType(v int32) *ListRecordsResponseBodyRecordsData {
	s.MatchingType = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetPort(v int32) *ListRecordsResponseBodyRecordsData {
	s.Port = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetPriority(v int32) *ListRecordsResponseBodyRecordsData {
	s.Priority = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetSelector(v int32) *ListRecordsResponseBodyRecordsData {
	s.Selector = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetTag(v string) *ListRecordsResponseBodyRecordsData {
	s.Tag = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetTags(v map[string]interface{}) *ListRecordsResponseBodyRecordsData {
	s.Tags = v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetType(v int32) *ListRecordsResponseBodyRecordsData {
	s.Type = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetUsage(v int32) *ListRecordsResponseBodyRecordsData {
	s.Usage = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetValue(v string) *ListRecordsResponseBodyRecordsData {
	s.Value = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetWeight(v int32) *ListRecordsResponseBodyRecordsData {
	s.Weight = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) Validate() error {
	return dara.Validate(s)
}
