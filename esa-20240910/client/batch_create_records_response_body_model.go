// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordResultList(v *BatchCreateRecordsResponseBodyRecordResultList) *BatchCreateRecordsResponseBody
	GetRecordResultList() *BatchCreateRecordsResponseBodyRecordResultList
	SetRequestId(v string) *BatchCreateRecordsResponseBody
	GetRequestId() *string
}

type BatchCreateRecordsResponseBody struct {
	// The records that have been created and failed to be created.
	RecordResultList *BatchCreateRecordsResponseBodyRecordResultList `json:"RecordResultList,omitempty" xml:"RecordResultList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2430E05E-1340-5773-B5E1-B743929F46F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCreateRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBody) GetRecordResultList() *BatchCreateRecordsResponseBodyRecordResultList {
	return s.RecordResultList
}

func (s *BatchCreateRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateRecordsResponseBody) SetRecordResultList(v *BatchCreateRecordsResponseBodyRecordResultList) *BatchCreateRecordsResponseBody {
	s.RecordResultList = v
	return s
}

func (s *BatchCreateRecordsResponseBody) SetRequestId(v string) *BatchCreateRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchCreateRecordsResponseBodyRecordResultList struct {
	// The records that failed to be created.
	Failed []*BatchCreateRecordsResponseBodyRecordResultListFailed `json:"Failed,omitempty" xml:"Failed,omitempty" type:"Repeated"`
	// The records that have been created.
	Success []*BatchCreateRecordsResponseBodyRecordResultListSuccess `json:"Success,omitempty" xml:"Success,omitempty" type:"Repeated"`
	// The total number of returned records.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultList) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) GetFailed() []*BatchCreateRecordsResponseBodyRecordResultListFailed {
	return s.Failed
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) GetSuccess() []*BatchCreateRecordsResponseBodyRecordResultListSuccess {
	return s.Success
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) SetFailed(v []*BatchCreateRecordsResponseBodyRecordResultListFailed) *BatchCreateRecordsResponseBodyRecordResultList {
	s.Failed = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) SetSuccess(v []*BatchCreateRecordsResponseBodyRecordResultListSuccess) *BatchCreateRecordsResponseBodyRecordResultList {
	s.Success = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) SetTotalCount(v int32) *BatchCreateRecordsResponseBodyRecordResultList {
	s.TotalCount = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) Validate() error {
	return dara.Validate(s)
}

type BatchCreateRecordsResponseBodyRecordResultListFailed struct {
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
	// The DNS information about the record, which contains various types of record values and their related attributes.
	//
	// example:
	//
	// {"value":"2.2.2.2"}
	Data *BatchCreateRecordsResponseBodyRecordResultListFailedData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The result description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The DNS type of the record, such as **A/AAAA, CNAME, and TXT**.
	//
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The origin type of the CNAME record. This field is left empty for other types of records. The type of the origin server. Valid values:
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
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The TTL of the record. Unit: seconds. If the value is 1, the TTL of the record is determined by the system.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailed) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailed) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetBizName() *string {
	return s.BizName
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetData() *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	return s.Data
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetDescription() *string {
	return s.Description
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetProxied() *bool {
	return s.Proxied
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetRecordId() *int64 {
	return s.RecordId
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetRecordName() *string {
	return s.RecordName
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetRecordType() *string {
	return s.RecordType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetSourceType() *string {
	return s.SourceType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetTtl() *int32 {
	return s.Ttl
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetBizName(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.BizName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetData(v *BatchCreateRecordsResponseBodyRecordResultListFailedData) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Data = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetDescription(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Description = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetProxied(v bool) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Proxied = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetRecordId(v int64) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.RecordId = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetRecordName(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.RecordName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetRecordType(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.RecordType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetSourceType(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetTtl(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Ttl = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) Validate() error {
	return dara.Validate(s)
}

type BatchCreateRecordsResponseBodyRecordResultListFailedData struct {
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
	// The flag bit of the record. Indicates its priority and handling method, used in CAA records.
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
	// The algorithm policy used to match or validate the certificate. Valid values: 0 to 255. Applicable to SMIMEA and TLSA records.
	//
	// example:
	//
	// RSA
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port number of the record, associated with the SRV record. Exclusive to SRV records.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. Valid values: 0 to 65535. A smaller value indicates a higher priority. Applicable to MX, SRV, and URI records.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key. Valid values: 0 to 255. Applicable to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// Indicates its priority and handling method, used in CAA records.
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
	// The record value or part of the record content. This value is returned when the record is A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, or URI. It has different meanings based on types of records:
	//
	// 	- **A/AAAA**: the IP addresses. IP addresses are separated by commas (,). There is at least one IPv4 address.
	//
	// 	- **CNAME**: the mapped domain name.
	//
	// 	- **NS**: the nameservers for the domain name.
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
	// The weight of the record. Applicable to SRV and URI records.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailedData) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailedData) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetCertificate() *string {
	return s.Certificate
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetFlag() *int32 {
	return s.Flag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetPort() *int32 {
	return s.Port
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetPriority() *int32 {
	return s.Priority
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetSelector() *int32 {
	return s.Selector
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetTag() *string {
	return s.Tag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetType() *int32 {
	return s.Type
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetUsage() *int32 {
	return s.Usage
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetValue() *string {
	return s.Value
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) GetWeight() *int32 {
	return s.Weight
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetAlgorithm(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Algorithm = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetCertificate(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Certificate = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetFingerprint(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Fingerprint = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetFlag(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Flag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetKeyTag(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.KeyTag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetMatchingType(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.MatchingType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetPort(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Port = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetPriority(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Priority = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetSelector(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Selector = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetTag(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Tag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetType(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Type = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetUsage(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Usage = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetValue(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Value = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetWeight(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Weight = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) Validate() error {
	return dara.Validate(s)
}

type BatchCreateRecordsResponseBodyRecordResultListSuccess struct {
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
	// The DNS record information.
	//
	// example:
	//
	// {"value":"1.1.1.1"}
	Data *BatchCreateRecordsResponseBodyRecordResultListSuccessData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The result description.
	//
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The DNS type of the record, such as **A/AAAA, CNAME, and TXT**.
	//
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The origin type of the CNAME record. This field is left empty for other types of records. The type of the origin server. Valid values:
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
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The TTL of the record. Unit: seconds. If the value is 1, the TTL of the record is determined by the system.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccess) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccess) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetBizName() *string {
	return s.BizName
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetData() *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	return s.Data
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetDescription() *string {
	return s.Description
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetProxied() *bool {
	return s.Proxied
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetRecordId() *int64 {
	return s.RecordId
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetRecordName() *string {
	return s.RecordName
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetRecordType() *string {
	return s.RecordType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetSourceType() *string {
	return s.SourceType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetTtl() *int32 {
	return s.Ttl
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetBizName(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.BizName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetData(v *BatchCreateRecordsResponseBodyRecordResultListSuccessData) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Data = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetDescription(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Description = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetProxied(v bool) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Proxied = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetRecordId(v int64) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.RecordId = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetRecordName(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.RecordName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetRecordType(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.RecordType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetSourceType(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetTtl(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Ttl = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) Validate() error {
	return dara.Validate(s)
}

type BatchCreateRecordsResponseBodyRecordResultListSuccessData struct {
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
	// The flag bit of the record. Indicates its priority and handling method, used in CAA records.
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
	// The algorithm policy used to match or validate the certificate. Valid values: 0 to 255. Applicable to SMIMEA and TLSA records.
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
	// The priority of the record. Valid values: 0 to 65535. A smaller value indicates a higher priority. Applicable to MX, SRV, and URI records.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key. Valid values: 0 to 255. Applicable to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The label of a CAA record, which indicates its specific type and purpose, such as issue, issuewild, and iodef.
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
	// The record value or part of the record content. This value is returned when the record is A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, or URI. It has different meanings based on types of records:
	//
	// 	- **A/AAAA**: the IP addresses. Multiple IPs are separated by commas (,). There is at least one IPv4 address.
	//
	// 	- **CNAME**: the mapped domain name.
	//
	// 	- **NS**: the nameservers for the domain name.
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
	// The weight of the record. Valid values: 0 to 65535. Applicable to SRV and URI records.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccessData) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccessData) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetCertificate() *string {
	return s.Certificate
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetFlag() *int32 {
	return s.Flag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetPort() *int32 {
	return s.Port
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetPriority() *int32 {
	return s.Priority
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetSelector() *int32 {
	return s.Selector
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetTag() *string {
	return s.Tag
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetType() *int32 {
	return s.Type
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetUsage() *int32 {
	return s.Usage
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetValue() *string {
	return s.Value
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) GetWeight() *int32 {
	return s.Weight
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetAlgorithm(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Algorithm = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetCertificate(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Certificate = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetFingerprint(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Fingerprint = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetFlag(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Flag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetKeyTag(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.KeyTag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetMatchingType(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.MatchingType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetPort(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Port = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetPriority(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Priority = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetSelector(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Selector = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetTag(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Tag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetType(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Type = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetUsage(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Usage = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetValue(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Value = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetWeight(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Weight = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) Validate() error {
	return dara.Validate(s)
}
