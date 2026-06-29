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
	SetHttpPorts(v string) *UpdateRecordRequest
	GetHttpPorts() *string
	SetHttpsPorts(v string) *UpdateRecordRequest
	GetHttpsPorts() *string
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
	// The business scenario for record acceleration. This parameter is not required for records without acceleration enabled. Valid values:
	//
	// - **video_image**: video and image.
	//
	// - **api**: API.
	//
	// - **web**: web page.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comment for the record.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS information of the record. The content varies depending on the record type. For more information, see <props="china">[documentation](https://help.aliyun.com/document_detail/2708761.html)<props="intl">[documentation](https://www.alibabacloud.com/help/doc-detail/2708761.html).
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
	// The back-to-origin HOST policy. This parameter takes effect when the record type is CNAME. Settings the HOST policy for back-to-origin requests. Valid values:
	//
	// - **follow_hostname**: follows the host record.
	//
	// - **follow_origin_domain**: follows the Origin Domain Name.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	HttpPorts  *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether to enable proxy acceleration for the record. Only CNAME records and A/AAAA records support proxy acceleration. Valid values:
	//
	// - **true**: Enable proxy acceleration.
	//
	// - **false**: Disable proxy acceleration.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The ID of the record. You can call [ListRecords](https://help.aliyun.com/document_detail/2850265.html) to obtain the record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The origin server type of the CNAME record. This parameter is required when you add a CNAME record. Valid values:
	//
	// - **OSS**: OSS origin server.
	//
	// - **S3**: S3 origin server.
	//
	// - **LB**: load balancing origin server.
	//
	// - **OP**: IPAM pool origin server.
	//
	// - **Domain**: standard domain name origin server.
	//
	// If this parameter is not specified or is left empty, the default value is Domain, which indicates a standard domain name origin server type.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time-to-live (TTL) of the record, in seconds. Valid values: **30 to 86400**, or 1. A value of 1 indicates that the TTL of the record is automatically determined.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The DNS type of the record, such as A/AAAA, CNAME, or TXT.
	//
	// example:
	//
	// A/AAAA
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

func (s *UpdateRecordRequest) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *UpdateRecordRequest) GetHttpsPorts() *string {
	return s.HttpsPorts
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

func (s *UpdateRecordRequest) SetHttpPorts(v string) *UpdateRecordRequest {
	s.HttpPorts = &v
	return s
}

func (s *UpdateRecordRequest) SetHttpsPorts(v string) *UpdateRecordRequest {
	s.HttpsPorts = &v
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
	// The AccessKey of the account to which the origin server belongs. This parameter is required when the origin server type is OSS and the origin authentication type is private cross-account read, or when the origin server type is S3 and the origin authentication type is private read.
	//
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The origin authentication type. Different origin server types support different authentication types. The origin server type refers to the SourceType parameter in this operation. When the origin server type is OSS or S3, you must specify the origin authentication type. Valid values:
	//
	// - **public**: public read. Select this value when the origin server type is OSS or S3 and the origin server allows public read access.
	//
	// - **private**: private read. Select this value when the origin server type is S3 and the origin server allows only private read access.
	//
	// - **private_same_account**: private same-account read. Select this value when the origin server type is OSS, the origin server is under the same Alibaba Cloud account, and the origin server allows only private read access.
	//
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The signature algorithm version. This parameter is required when the origin server type is S3 and the origin authentication type is private read. Valid values:
	//
	// - **v2**
	//
	// - **v4**
	//
	// Default value: v4.
	//
	// example:
	//
	// v2
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The SecretKey of the account to which the origin server belongs. This parameter is required when the origin server type is OSS and the origin authentication type is private cross-account read, or when the origin server type is S3 and the origin authentication type is private read.
	//
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The region to which the origin server belongs. This parameter is required when the origin server type is S3. Obtain the region from the official S3 website.
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
	// The encryption algorithm used by the record. Valid values: **0 to 255**. This parameter is required when you add CERT or SSHFP records.
	//
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key certificate information of the record. This parameter is required when you add CERT, SMIMEA, or TLSA records.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint value of the record. This parameter is required when you add SSHFP records.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The flag of the record. The Flag of a CAA record indicates its priority and processing method. Valid values: **0 to 255**. This parameter is required when you add CAA records.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identifier of the record. Valid values: **0 to 65535**. This parameter is required when you add CERT records.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used by the record for matching or verifying certificates. Valid values: **0 to 255**. This parameter is required when you add SMIMEA or TLSA records.
	//
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port of the record. Valid values: **0 to 65535**. This parameter is required when you add SRV records.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. Valid values: **0 to 65535**. A smaller value indicates a higher priority. This parameter is required when you add MX, SRV, or URI records.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key used by the record. Valid values: **0 to 255**. This parameter is required when you add SMIMEA or TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag of the record. The Tag of a CAA record indicates its specific type and purpose. This parameter is required when you add CAA records.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type of the record (for CERT records) or the public key type (for SSHFP records). This parameter is required when you add CERT or SSHFP records.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier of the record. Valid values: **0 to 255**. This parameter is required when you add SMIMEA or TLSA records.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value or partial content. This parameter is required when you add A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, or URI records. The meaning varies depending on the record type:
	//
	// - **A/AAAA**: The IP address to which the record points. Separate multiple IP addresses with commas (,). At least one IPv4 address is required.
	//
	// - **CNAME**: The target domain name to which the record points.
	//
	// - **NS**: The name server for the specified domain name.
	//
	// - **MX**: A valid target mail server domain name.
	//
	// - **TXT**: A valid text string.
	//
	// - **CAA**: A valid certification authority domain name.
	//
	// - **SRV**: A valid target host domain name.
	//
	// - **URI**: A valid URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. Valid values: **0 to 65535**. This parameter is required when you add SRV or URI records.
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
