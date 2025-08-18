// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordList(v []*BatchCreateRecordsRequestRecordList) *BatchCreateRecordsRequest
	GetRecordList() []*BatchCreateRecordsRequestRecordList
	SetSiteId(v int64) *BatchCreateRecordsRequest
	GetSiteId() *int64
}

type BatchCreateRecordsRequest struct {
	// The list of DNS records to be created.
	//
	// This parameter is required.
	RecordList []*BatchCreateRecordsRequestRecordList `json:"RecordList,omitempty" xml:"RecordList,omitempty" type:"Repeated"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s BatchCreateRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsRequest) GetRecordList() []*BatchCreateRecordsRequestRecordList {
	return s.RecordList
}

func (s *BatchCreateRecordsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchCreateRecordsRequest) SetRecordList(v []*BatchCreateRecordsRequestRecordList) *BatchCreateRecordsRequest {
	s.RecordList = v
	return s
}

func (s *BatchCreateRecordsRequest) SetSiteId(v int64) *BatchCreateRecordsRequest {
	s.SiteId = &v
	return s
}

func (s *BatchCreateRecordsRequest) Validate() error {
	return dara.Validate(s)
}

type BatchCreateRecordsRequestRecordList struct {
	AuthConf *BatchCreateRecordsRequestRecordListAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The business scenario of the record for acceleration. Valid values:
	//
	// 	- **image_video**
	//
	// 	- **api**
	//
	// 	- **web**
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The DNS information of the record. Enter fields based on the record type.
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
	Data *BatchCreateRecordsRequestRecordListData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Specifies whether to proxy the record. Only CNAME and A/AAAA records can be proxied. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
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
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The DNS type of the record.
	//
	// This parameter is required.
	//
	// example:
	//
	// A/AAAA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BatchCreateRecordsRequestRecordList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsRequestRecordList) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsRequestRecordList) GetAuthConf() *BatchCreateRecordsRequestRecordListAuthConf {
	return s.AuthConf
}

func (s *BatchCreateRecordsRequestRecordList) GetBizName() *string {
	return s.BizName
}

func (s *BatchCreateRecordsRequestRecordList) GetData() *BatchCreateRecordsRequestRecordListData {
	return s.Data
}

func (s *BatchCreateRecordsRequestRecordList) GetProxied() *bool {
	return s.Proxied
}

func (s *BatchCreateRecordsRequestRecordList) GetRecordName() *string {
	return s.RecordName
}

func (s *BatchCreateRecordsRequestRecordList) GetSourceType() *string {
	return s.SourceType
}

func (s *BatchCreateRecordsRequestRecordList) GetTtl() *int32 {
	return s.Ttl
}

func (s *BatchCreateRecordsRequestRecordList) GetType() *string {
	return s.Type
}

func (s *BatchCreateRecordsRequestRecordList) SetAuthConf(v *BatchCreateRecordsRequestRecordListAuthConf) *BatchCreateRecordsRequestRecordList {
	s.AuthConf = v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetBizName(v string) *BatchCreateRecordsRequestRecordList {
	s.BizName = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetData(v *BatchCreateRecordsRequestRecordListData) *BatchCreateRecordsRequestRecordList {
	s.Data = v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetProxied(v bool) *BatchCreateRecordsRequestRecordList {
	s.Proxied = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetRecordName(v string) *BatchCreateRecordsRequestRecordList {
	s.RecordName = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetSourceType(v string) *BatchCreateRecordsRequestRecordList {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetTtl(v int32) *BatchCreateRecordsRequestRecordList {
	s.Ttl = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetType(v string) *BatchCreateRecordsRequestRecordList {
	s.Type = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) Validate() error {
	return dara.Validate(s)
}

type BatchCreateRecordsRequestRecordListAuthConf struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s BatchCreateRecordsRequestRecordListAuthConf) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsRequestRecordListAuthConf) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) GetRegion() *string {
	return s.Region
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) GetVersion() *string {
	return s.Version
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) SetAccessKey(v string) *BatchCreateRecordsRequestRecordListAuthConf {
	s.AccessKey = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) SetAuthType(v string) *BatchCreateRecordsRequestRecordListAuthConf {
	s.AuthType = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) SetRegion(v string) *BatchCreateRecordsRequestRecordListAuthConf {
	s.Region = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) SetSecretKey(v string) *BatchCreateRecordsRequestRecordListAuthConf {
	s.SecretKey = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) SetVersion(v string) *BatchCreateRecordsRequestRecordListAuthConf {
	s.Version = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListAuthConf) Validate() error {
	return dara.Validate(s)
}

type BatchCreateRecordsRequestRecordListData struct {
	// The encryption algorithm used for the record. Valid values: 0 to 255. Applicable to CERT and SSHFP records.
	//
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key of the certificate. Applicable to CERT, SMIMEA, and TLSA records.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint of the record. Applicable to SSHFP records.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The Flag for a CAA record indicates its priority and how it is processed. Valid values: 0 to 255.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identification for the record. Valid values: 0 to 65535. Applicable to CERT records.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used to match or validate the certificate. Valid values: 0 to 255. Applicable to SMIMEA, and TLSA records.
	//
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port of the record. Valid values: 0 to 65535. Exclusive to SRV records.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. Valid values: 0 to 65535. A smaller value indicates a higher priority. This parameter is required when you add MX, SRV, and URI records.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key. Valid values: 0 to 255. Applicable to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag of a CAA record, which indicates its specific type and purpose, such as issue, issuewild, and iodef.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type of the record (in CERT records), or the public key type (in SSHFP records).
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier of the record. Valid values: 0 to 255. Applicable to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value or part of the record content. A/AAAA: the IP address being pointed to. CNAME: the target domain name being pointed to. MX: valid target mail server domain name. TXT: valid text string. CAA: valid certificate authority domain name. SRV: valid target host domain name. URI: valid URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. Valid values: 0 to 65,535. Applicable to SRV and URI records.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s BatchCreateRecordsRequestRecordListData) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsRequestRecordListData) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsRequestRecordListData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *BatchCreateRecordsRequestRecordListData) GetCertificate() *string {
	return s.Certificate
}

func (s *BatchCreateRecordsRequestRecordListData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *BatchCreateRecordsRequestRecordListData) GetFlag() *int32 {
	return s.Flag
}

func (s *BatchCreateRecordsRequestRecordListData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *BatchCreateRecordsRequestRecordListData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *BatchCreateRecordsRequestRecordListData) GetPort() *int32 {
	return s.Port
}

func (s *BatchCreateRecordsRequestRecordListData) GetPriority() *int32 {
	return s.Priority
}

func (s *BatchCreateRecordsRequestRecordListData) GetSelector() *int32 {
	return s.Selector
}

func (s *BatchCreateRecordsRequestRecordListData) GetTag() *string {
	return s.Tag
}

func (s *BatchCreateRecordsRequestRecordListData) GetType() *int32 {
	return s.Type
}

func (s *BatchCreateRecordsRequestRecordListData) GetUsage() *int32 {
	return s.Usage
}

func (s *BatchCreateRecordsRequestRecordListData) GetValue() *string {
	return s.Value
}

func (s *BatchCreateRecordsRequestRecordListData) GetWeight() *int32 {
	return s.Weight
}

func (s *BatchCreateRecordsRequestRecordListData) SetAlgorithm(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Algorithm = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetCertificate(v string) *BatchCreateRecordsRequestRecordListData {
	s.Certificate = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetFingerprint(v string) *BatchCreateRecordsRequestRecordListData {
	s.Fingerprint = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetFlag(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Flag = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetKeyTag(v int32) *BatchCreateRecordsRequestRecordListData {
	s.KeyTag = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetMatchingType(v int32) *BatchCreateRecordsRequestRecordListData {
	s.MatchingType = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetPort(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Port = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetPriority(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Priority = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetSelector(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Selector = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetTag(v string) *BatchCreateRecordsRequestRecordListData {
	s.Tag = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetType(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Type = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetUsage(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Usage = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetValue(v string) *BatchCreateRecordsRequestRecordListData {
	s.Value = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetWeight(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Weight = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) Validate() error {
	return dara.Validate(s)
}
