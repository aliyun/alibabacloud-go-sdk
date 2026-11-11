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
	// The current page number. The value is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DNS information of the records. For more information, see <props="china">[Documentation](https://help.aliyun.com/document_detail/2708761.html)<props="intl">[Documentation](https://www.alibabacloud.com/help/doc-detail/2708761.html).
	Records []*ListRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
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
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecordsResponseBodyRecords struct {
	// The origin server authentication information of the CNAME record.
	AuthConf *ListRecordsResponseBodyRecordsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The business scenario for record acceleration. Valid values:
	//
	// - **image_video**: images and videos
	//
	// - **api**: APIs
	//
	// - **web**: web pages
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The description of the record.
	//
	// example:
	//
	// this is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time of the record. The time is displayed in UTC and follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Specifies whether a custom port is enabled for the record. Valid values:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// off
	CustomPort *string `json:"CustomPort,omitempty" xml:"CustomPort,omitempty"`
	// The specific DNS information of the record. Different types of records contain different information.
	//
	// example:
	//
	// {"value":"1.1.1.1"}
	Data *ListRecordsResponseBodyRecordsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The back-to-origin host policy. This policy takes effect when the record type is CNAME. Settings for the host policy in back-to-origin requests include two patterns:
	//
	// - **follow_hostname**: follow the request host
	//
	// - **follow_origin_domain**: follow the Origin Domain Name
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// The HTTP port. Port 80 is supported by default.
	//
	// example:
	//
	// 80,8080,8880,2052,2082,2086,2095
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The HTTPS port. Port 443 is supported by default.
	//
	// example:
	//
	// 443,2053,2083,2087,2096,8443
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether proxy acceleration is enabled for the record. Valid values:
	//
	// - **true**: Proxy acceleration is enabled.
	//
	// - **false**: Proxy acceleration is disabled.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The CNAME of the record. If the site is onboarded by using CNAME access, you must configure this CNAME value for the record.
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
	// The origin server type of the CNAME record. This parameter is required when you add a CNAME record. Valid values:
	//
	// - **OSS**: OSS origin server
	//
	// - **S3**: S3 origin server
	//
	// - **LB**: load balancing origin server
	//
	// - **OP**: IPAM pool origin server
	//
	// - **Domain**: regular domain name origin server
	//
	// If this parameter is not specified or is left empty, the default value is Domain, which indicates a regular domain name origin server.
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
	// The ID of the site to which the record belongs.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The name of the site to which the record belongs.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The time-to-live (TTL) of the record, in seconds. A value of 1 indicates that the TTL is automatic.
	//
	// example:
	//
	// 30
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The update time of the record. The time is displayed in UTC and follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
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

func (s *ListRecordsResponseBodyRecords) GetCustomPort() *string {
	return s.CustomPort
}

func (s *ListRecordsResponseBodyRecords) GetData() *ListRecordsResponseBodyRecordsData {
	return s.Data
}

func (s *ListRecordsResponseBodyRecords) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *ListRecordsResponseBodyRecords) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *ListRecordsResponseBodyRecords) GetHttpsPorts() *string {
	return s.HttpsPorts
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

func (s *ListRecordsResponseBodyRecords) SetCustomPort(v string) *ListRecordsResponseBodyRecords {
	s.CustomPort = &v
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

func (s *ListRecordsResponseBodyRecords) SetHttpPorts(v string) *ListRecordsResponseBodyRecords {
	s.HttpPorts = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetHttpsPorts(v string) *ListRecordsResponseBodyRecords {
	s.HttpsPorts = &v
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

type ListRecordsResponseBodyRecordsAuthConf struct {
	// The AccessKey ID of the account to which the origin server belongs. This parameter is returned when the origin server type is OSS and the origin server authentication type is private read across accounts, or when the origin server type is S3 and the origin server authentication type is private read.
	//
	// example:
	//
	// u0Nkg5gBK***QF5wvKMM504JUHt
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The origin server authentication type. Different origin server types support different authentication types. The origin server type refers to the SourceType parameter in this operation. When the origin server type is OSS or S3, you must specify the origin server authentication type. Valid values:
	//
	// - **public**: Public read. This value is used when the origin server type is OSS or S3 and the origin server allows public read access.
	//
	// - **private**: Private read. This value is used when the origin server type is S3 and the origin server allows private read access.
	//
	// - **private_same_account**: Private read within the same account. This value is used when the origin server type is OSS, the origin server and the current service are within the same Alibaba Cloud account, and the origin server allows private read access.
	//
	// - **private_cross_account**: Private read across accounts. This value is used when the origin server type is OSS, the origin server and the current service are within different Alibaba Cloud accounts, and the origin server allows private read access.
	//
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region to which the origin server belongs. This parameter is returned when the origin server type is S3. You can obtain the region information from the official website of S3.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The AccessKey secret of the account to which the origin server belongs. This parameter is returned when the origin server type is OSS and the origin server authentication type is private read across accounts, or when the origin server type is S3 and the origin server authentication type is private read.
	//
	// example:
	//
	// VIxuvJSA2S03f***kp208dy5w7
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signature algorithm version. This parameter is returned when the origin server type is S3 and the origin server authentication type is private read. Valid values:
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
	// The encryption algorithm used by the record. Valid values: **0*	- to **255**. This parameter is exclusive to CERT and SSHFP records.
	//
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key certificate information of the record. This parameter is exclusive to CERT, SMIMEA, and TLSA records.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint value of the record. This parameter is exclusive to SSHFP records.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The flag of the record. For CAA records, the flag indicates the priority and processing method. Valid values: **0*	- to **255**. This parameter is exclusive to CAA records.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key tag of the record. Valid values: **0*	- to **65535**. This parameter is exclusive to CERT records.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used by the record to match or verify the certificate. Valid values: **0*	- to **255**. This parameter is exclusive to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port of the record. Valid values: **0*	- to **65535**. This parameter is exclusive to SRV records.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. Valid values: **0*	- to **65535**. A smaller value indicates a higher priority. This parameter is exclusive to MX, SRV, and URI records.
	//
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the certificate or public key used by the record. Valid values: **0*	- to **255**. This parameter is exclusive to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag of the record. For CAA records, the tag indicates the specific type and usage. This parameter is exclusive to CAA records.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The tags of the record.
	//
	// example:
	//
	// {\\"Name\\":\\"PARAMETER\\"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The certificate type of the record in CERT records, or the public key type in SSHFP records.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier of the record. Valid values: **0*	- to **255**. This parameter is exclusive to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value or partial content. This parameter is included in A, AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI records. It indicates different meanings in different types of records:
	//
	// - **A/AAAA**: The destination IP address. If multiple IP addresses are specified, separate them with commas (,). At least one IPv4 address is required.
	//
	// - **CNAME**: The destination domain name.
	//
	// - **NS**: The name server of the specified domain name.
	//
	// - **MX**: The valid destination mail server domain name.
	//
	// - **TXT**: The valid text string.
	//
	// - **CAA**: The valid certification authority domain name.
	//
	// - **SRV**: The valid destination host domain name.
	//
	// - **URI**: The valid URI string.
	//
	// example:
	//
	// CNAME
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. Valid values: **0*	- to **65535**. This parameter is exclusive to SRV and URI records.
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
