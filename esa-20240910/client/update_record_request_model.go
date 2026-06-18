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
	// The origin authentication settings for the CNAME record.
	AuthConf *UpdateRecordRequestAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The use case for proxy acceleration. Omit this parameter if proxy acceleration is disabled. Valid values:
	//
	// - **video_image**: Video and images.
	//
	// - **api**: APIs.
	//
	// - **web**: Web pages.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// A comment for the record.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS data for the record. The required content varies based on the record type. For more information, see <props="china">[Documentation](https://help.aliyun.com/document_detail/2708761.html)<props="intl">[Documentation](https://www.alibabacloud.com/help/doc-detail/2708761.html).
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
	// The origin HOST policy. This policy, which applies only to CNAME records, determines the value of the `HOST` header in requests sent to the origin. Valid values:
	//
	// - **follow_hostname**: Follows the host record.
	//
	// - **follow_origin_domain**: Follows the origin domain name.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	HttpPorts  *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Indicates whether to enable proxy acceleration for the record. Only CNAME and A/AAAA records support proxy acceleration. Valid values:
	//
	// - **true**: Enables proxy acceleration.
	//
	// - **false**: Disables proxy acceleration.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The record ID. Call the [ListRecords](https://help.aliyun.com/document_detail/2850265.html) operation to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The origin type for the CNAME record. This parameter is required for CNAME records. Valid values:
	//
	// - **OSS**: An OSS origin.
	//
	// - **S3**: An S3 origin.
	//
	// - **LB**: A load balancer origin.
	//
	// - **OP**: An origin address pool origin.
	//
	// - **Domain**: A standard domain name origin.
	//
	// If this parameter is omitted or left empty, the default value is `Domain`.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The record\\"s time to live (TTL) in seconds. The value must be an integer from **30 to 86400*	- or 1. A value of 1 sets the TTL to automatic.
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
	// The access key for the account that owns the origin. This is required for private, cross-account access to OSS origins, and for S3 origins where the authentication type is **private**.
	//
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The origin authentication type. This parameter is required when the **SourceType*	- is **OSS*	- or **S3**. Supported authentication types vary depending on the origin type. Valid values:
	//
	// - **public**: Public read. Use for publicly readable OSS or S3 origins.
	//
	// - **private**: Private read. Use for private S3 origins.
	//
	// - **private_same_account**: Private read within the same account. Use for private OSS origins accessed from the same Alibaba Cloud account.
	//
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region where the origin is located. This parameter is required when the origin type is S3. You can find the region ID on the official S3 website.
	//
	// - **v2**
	//
	// - **v4**
	//
	// If you do not specify a value, it defaults to v4.
	//
	// example:
	//
	// v2
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The secret key for the account that owns the origin. This is required for private, cross-account access to OSS origins, and for S3 origins where the authentication type is **private**.
	//
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signing algorithm version. This parameter is required when the origin type is S3 and the authentication type is **private**. Supported versions: v2 and v4. If this parameter is not specified, the default value is v4.
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
	// The encryption algorithm specified in the record. The value must be an integer from **0 to 255**. This parameter is required for CERT and SSHFP records.
	//
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key certificate data for the record. This parameter is required for CERT, SMIMEA, and TLSA records.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint for the record. This parameter is required for SSHFP records.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The flag for the record. For a CAA record, this flag indicates its priority and handling behavior. The value must be an integer from **0 to 255**. This parameter is required for CAA records.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identifier for the record. The value must be an integer from **0 to 65535**. This parameter is required for CERT records.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used to match or validate the certificate. The value must be an integer from **0 to 255**. This parameter is required for SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port number for the record. The value must be an integer from **0 to 65535**. This parameter is required for SRV records.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The record\\"s priority. The value must be an integer from **0 to 65535**, where a lower value indicates higher priority. This parameter is required for MX, SRV, and URI records.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key specified in the record. The value must be an integer from **0 to 255**. This parameter is required for SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag for the record. For a CAA record, the tag specifies the record\\"s type and purpose. This parameter is required for CAA records.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type for a CERT record, or the public key type for an SSHFP record. This parameter is required for CERT and SSHFP records.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier for the record. The value must be an integer from **0 to 255**. This parameter is required for SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The value of the record or part of its content. This parameter is required for A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI records. The meaning of this parameter varies by record type:
	//
	// - **A/AAAA**: The target IP address. To specify multiple IP addresses, separate them with a comma (,). At least one IPv4 address is required.
	//
	// - **CNAME**: The target domain name.
	//
	// - **NS**: The name server for the domain.
	//
	// - **MX**: A valid domain name for the target mail server.
	//
	// - **TXT**: A valid text string.
	//
	// - **CAA**: A valid domain name for the certificate authority.
	//
	// - **SRV**: A valid domain name for the target host.
	//
	// - **URI**: A valid URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. The value must be an integer from **0 to 65535**. This parameter is required for SRV and URI records.
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
