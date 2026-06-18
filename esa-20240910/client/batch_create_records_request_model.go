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
	// The list of DNS records to create.
	//
	// This parameter is required.
	RecordList []*BatchCreateRecordsRequestRecordList `json:"RecordList,omitempty" xml:"RecordList,omitempty" type:"Repeated"`
	// The ID of the site. You can get this ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
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
	if s.RecordList != nil {
		for _, item := range s.RecordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateRecordsRequestRecordList struct {
	// The origin authentication information for the CNAME record.
	AuthConf *BatchCreateRecordsRequestRecordListAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// The use case for proxy acceleration. Valid values:
	//
	// - **image_video**: Images and videos.
	//
	// - **api**: APIs.
	//
	// - **web**: Web pages.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The content of the DNS record. The required fields depend on the record type.
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
	Data       *BatchCreateRecordsRequestRecordListData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpPorts  *string                                  `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts *string                                  `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether to enable proxy acceleration for the record. Only CNAME and A/AAAA records support proxy acceleration. Valid values:
	//
	// - **true**: Enables proxy acceleration.
	//
	// - **false**: Disables proxy acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The name of the record.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The origin type for the CNAME record. This parameter is required when you add a CNAME record. Valid values:
	//
	// - **OSS**: An OSS origin.
	//
	// - **S3**: An S3 origin.
	//
	// - **LB**: A load balancer origin.
	//
	// - **OP**: An origin pool origin.
	//
	// - **Domain**: A domain name origin.
	//
	// If omitted or left empty, this parameter defaults to `Domain`.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The Time to Live (TTL) for the record, in seconds. A value of `1` indicates an automatic TTL.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record.
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

func (s *BatchCreateRecordsRequestRecordList) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *BatchCreateRecordsRequestRecordList) GetHttpsPorts() *string {
	return s.HttpsPorts
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

func (s *BatchCreateRecordsRequestRecordList) SetHttpPorts(v string) *BatchCreateRecordsRequestRecordList {
	s.HttpPorts = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetHttpsPorts(v string) *BatchCreateRecordsRequestRecordList {
	s.HttpsPorts = &v
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

type BatchCreateRecordsRequestRecordListAuthConf struct {
	// The access key ID of the account that owns the origin. This parameter is required when the origin type is `OSS` and the authentication type is `private_cross_account`, or when the origin type is `S3` and the authentication type is `private`.
	//
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The type of origin authentication. Supported authentication types depend on the origin type, which is specified by the `SourceType` parameter. This parameter is required when the origin type is `OSS` or `S3`. Valid values:
	//
	// - **public**: For OSS or S3 origins with public read access.
	//
	// - **private**: For S3 origins with private read access.
	//
	// - **private_same_account**: For OSS origins with private read access within the same Alibaba Cloud account.
	//
	// - **private_cross_account**: For OSS origins with private read access from a different Alibaba Cloud account.
	//
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region where the S3 origin is located. This parameter is required when the origin type is `S3`. For a list of valid region IDs, refer to the official S3 documentation.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The secret key associated with the specified AccessKey. This parameter is required when the origin type is `OSS` and the authentication type is `private_cross_account`, or when the origin type is `S3` and the authentication type is `private`.
	//
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signature algorithm version. This parameter is applicable when the origin type is `S3` and the authentication type is `private`. Supported versions:
	//
	// - **v2**
	//
	// - **v4**
	//
	// If omitted, the default version is `v4`.
	//
	// example:
	//
	// v4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
	// The algorithm identifier for the record. Valid values range from **0-255**. This parameter applies to CERT and SSHFP records.
	//
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The certificate or public key data for the record. This parameter applies to CERT, SMIMEA, and TLSA records.
	//
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key fingerprint for the record. This parameter applies to SSHFP records.
	//
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The flag for the CAA record, which specifies how a Certificate Authority must handle the record. Valid values range from **0-255**.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identifier for the record. Valid values range from **0-65535**. This parameter applies to CERT records.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used to match or validate a certificate. Valid values range from **0-255**. This parameter applies to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port number for the record. Valid values range from **0-65535**. This parameter applies only to SRV records.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. Valid values range from **0-65535**. A lower value indicates a higher priority. This parameter is required for MX, SRV, or URI records.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key used by the record. Valid values range from **0-255**. This parameter applies to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag for the CAA record, which specifies its type and purpose, such as `issue`, `issuewild`, or `iodef`.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type for a CERT record or the public key type for an SSHFP record.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier for the record. Valid values range from **0-255**. This parameter applies to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value. The format depends on the record type.
	//
	// - **A/AAAA**: An IP address.
	//
	// - **CNAME**: The target domain name.
	//
	// - **MX**: The domain name of the target mail server.
	//
	// - **TXT**: A text string.
	//
	// - **CAA**: The domain name of a Certificate Authority.
	//
	// - **SRV**: The domain name of the target host.
	//
	// - **URI**: A URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. Valid values range from **0-65535**. This parameter applies to SRV and URI records.
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
