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
	// Details of the record.
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
	if s.RecordModel != nil {
		if err := s.RecordModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRecordResponseBodyRecordModel struct {
	// The origin authentication settings for the CNAME record.
	AuthConf *GetRecordResponseBodyRecordModelAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The business scenario for record acceleration. Valid values:
	//
	// - **image_video**: Images and videos.
	//
	// - **api**: APIs.
	//
	// - **web**: Web pages.
	//
	// example:
	//
	// image_video
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comment for the record.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the record was created. The time is in UTC and follows the ISO 8601 standard. Format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// example:
	//
	// 2023-03-10T13:30:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomPort *string `json:"CustomPort,omitempty" xml:"CustomPort,omitempty"`
	// DNS information for the record. The content varies based on the record type.
	//
	// example:
	//
	// {"value":"1.1.1.1"}
	Data *GetRecordResponseBodyRecordModelData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Specifies the policy for the `Host` header during origin-pull. This parameter applies only to CNAME records. Valid values:
	//
	// - **follow_hostname**: Follows the host record.
	//
	// - **follow_origin_domain**: Follows the origin domain.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	HttpPorts  *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Indicates whether the proxy is enabled for the record. This feature is supported only for A, AAAA, and CNAME records. Valid values:
	//
	// - **true**: The proxy is enabled.
	//
	// - **false**: The proxy is disabled.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The CNAME of the record.
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
	// The origin type for CNAME records. Valid values:
	//
	// - **OSS**: OSS origin.
	//
	// - **S3**: S3 origin.
	//
	// - **LB**: A load balancer.
	//
	// - **OP**: An origin pool.
	//
	// - **Domain**: A standard domain.
	//
	// If unspecified, the default is `Domain`.
	//
	// example:
	//
	// OSS
	RecordSourceType *string `json:"RecordSourceType,omitempty" xml:"RecordSourceType,omitempty"`
	// The DNS type of the record, such as **A/AAAA**, **CNAME**, or **TXT**.
	//
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The site ID.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The record\\"s Time to Live (TTL) in seconds. A value of 1 indicates an automatic TTL.
	//
	// example:
	//
	// 20
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The time when the record was last updated. The time is in UTC and follows the ISO 8601 standard. Format: `yyyy-MM-ddTHH:mm:ssZ`.
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

func (s *GetRecordResponseBodyRecordModel) GetCustomPort() *string {
	return s.CustomPort
}

func (s *GetRecordResponseBodyRecordModel) GetData() *GetRecordResponseBodyRecordModelData {
	return s.Data
}

func (s *GetRecordResponseBodyRecordModel) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *GetRecordResponseBodyRecordModel) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *GetRecordResponseBodyRecordModel) GetHttpsPorts() *string {
	return s.HttpsPorts
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

func (s *GetRecordResponseBodyRecordModel) SetCustomPort(v string) *GetRecordResponseBodyRecordModel {
	s.CustomPort = &v
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

func (s *GetRecordResponseBodyRecordModel) SetHttpPorts(v string) *GetRecordResponseBodyRecordModel {
	s.HttpPorts = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetHttpsPorts(v string) *GetRecordResponseBodyRecordModel {
	s.HttpsPorts = &v
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

type GetRecordResponseBodyRecordModelAuthConf struct {
	// The AccessKey ID of the account that owns the origin.
	//
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The origin authentication type. Valid values:
	//
	// - **public**: Public read. Select this value when the origin is an OSS or S3 origin with public read access.
	//
	// - **private**: Private read. Select this value when the origin is an S3 origin with private read access.
	//
	// - **private_same_account**: Private read within the same account. Select this value when the origin is an OSS origin with private read access under the same Alibaba Cloud account.
	//
	// - **private_cross_account**: Private read across accounts. Select this value when the origin is an OSS origin with private read access under a different Alibaba Cloud account.
	//
	// example:
	//
	// public
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The origin\\"s region. For S3 origins, you can find region codes on the official S3 website.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The AccessKey secret of the account that owns the origin.
	//
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The version of the signature algorithm. Supported versions:
	//
	// - **v2**
	//
	// - **v4**
	//
	// If unspecified, the default is `v4`.
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
	// The encryption algorithm of the record, from **0*	- to **255**.
	//
	// example:
	//
	// 1
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The record\\"s public key certificate.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint of the record.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The flag of the record. For CAA records, this flag indicates the priority and processing method. The value must be from **0*	- to **255**.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identifier of the record, from **0*	- to **65535**.
	//
	// example:
	//
	// 1
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy for matching or verifying the certificate, from **0*	- to **255**.
	//
	// example:
	//
	// 1
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port of the record, from **0*	- to **65535**.
	//
	// example:
	//
	// 8707
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record, from **0*	- to **65535**. A lower value indicates a higher priority.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The record\\"s certificate or public key type, from **0*	- to **255**.
	//
	// example:
	//
	// 1
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag of the record. For CAA records, the tag indicates the specific type and purpose.
	//
	// example:
	//
	// issue
	Tag  *string                `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The certificate type for a CERT record, or the public key type for an SSHFP record.
	//
	// example:
	//
	// RSA
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier of the record, from **0*	- to **255**.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value. The meaning of this parameter varies by the record type:
	//
	// - **A/AAAA**: The target IP address. Separate multiple IP addresses with a comma (`,`). Requires at least one IPv4 address.
	//
	// - **CNAME**: The target domain name.
	//
	// - **NS**: The name server for the specified domain.
	//
	// - **MX**: A valid target mail server domain name.
	//
	// - **TXT**: A valid text string.
	//
	// - **CAA**: A valid Certificate Authority domain name.
	//
	// - **SRV**: A valid target host domain name.
	//
	// - **URI**: A valid URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record, from **0*	- to **65535**.
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
