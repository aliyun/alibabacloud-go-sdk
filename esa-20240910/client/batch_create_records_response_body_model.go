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
	// The results of the batch record creation, with details for both successful and failed records.
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
	if s.RecordResultList != nil {
		if err := s.RecordResultList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateRecordsResponseBodyRecordResultList struct {
	// A list of records that failed to be created.
	Failed []*BatchCreateRecordsResponseBodyRecordResultListFailed `json:"Failed,omitempty" xml:"Failed,omitempty" type:"Repeated"`
	// A list of successfully created records.
	Success []*BatchCreateRecordsResponseBodyRecordResultListSuccess `json:"Success,omitempty" xml:"Success,omitempty" type:"Repeated"`
	// The total number of records in the creation operation.
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
	if s.Failed != nil {
		for _, item := range s.Failed {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Success != nil {
		for _, item := range s.Success {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateRecordsResponseBodyRecordResultListFailed struct {
	// The acceleration use case for the record. Valid values:
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
	// The DNS information for the record.
	//
	// example:
	//
	// {"value":"2.2.2.2"}
	Data *BatchCreateRecordsResponseBodyRecordResultListFailedData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The result description.
	//
	// example:
	//
	// 记录的名称非法
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HttpPorts   *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts  *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether proxy acceleration is enabled for the record. This option is available only for CNAME, A, and AAAA records. Valid values:
	//
	// - **true**: Proxy acceleration is enabled.
	//
	// - **false**: Proxy acceleration is disabled.
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
	// The DNS type of the record, such as **A/AAAA**, **CNAME**, or **TXT**.
	//
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The type of origin for a CNAME record. This parameter is empty for other record types. Valid values:
	//
	// - **OSS**: An OSS origin.
	//
	// - **S3**: An S3 origin.
	//
	// - **LB**: A load balancer origin.
	//
	// - **OP**: An origin pool.
	//
	// - **Domain**: A domain name origin.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The TTL of the record in seconds. A value of 1 sets the TTL to Automatic.
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

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) GetHttpsPorts() *string {
	return s.HttpsPorts
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

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetHttpPorts(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.HttpPorts = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetHttpsPorts(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.HttpsPorts = &v
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateRecordsResponseBodyRecordResultListFailedData struct {
	// The encryption algorithm used by the record. The value ranges from **0*	- to **255**. This parameter applies to CERT and SSHFP records.
	//
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key certificate for the record. This parameter applies to CERT, SMIMEA, and TLSA records.
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
	// The flag for the record, which indicates its priority and processing method. This parameter applies to CAA records.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identifier for the record. The value ranges from **0*	- to **65535**. This parameter applies to CERT records.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used by the record to match or validate certificates. The value ranges from **0*	- to **255**. This parameter applies to SMIMEA and TLSA records.
	//
	// example:
	//
	// RSA
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port for the record. The value ranges from 0 to 65535. This parameter applies only to SRV records.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. The value ranges from **0*	- to **65535**. A smaller value indicates a higher priority. This parameter applies to MX, SRV, and URI records.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key used by the record. The value ranges from **0*	- to **255**. This parameter applies to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag for a CAA record, which specifies its type and purpose, such as `issue`, `issuewild`, or `iodef`.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type for CERT records or the public key type for SSHFP records.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier for the record. The value ranges from **0*	- to **255**. This parameter applies to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value. This parameter applies to A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI records. The meaning of this parameter varies based on the record type:
	//
	// - **A/AAAA**: The IP address. To specify multiple addresses, separate them with a comma (,). At least one IPv4 address is required.
	//
	// - **CNAME**: The target domain name.
	//
	// - **NS**: The name server for the domain.
	//
	// - **MX**: The domain name of a valid target mail server.
	//
	// - **TXT**: A valid text string.
	//
	// - **CAA**: The domain name of a valid certificate authority.
	//
	// - **SRV**: The domain name of a valid target host.
	//
	// - **URI**: A valid URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. The value ranges from 0 to 65535. This parameter applies to SRV and URI records.
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
	// The acceleration use case for the record. Valid values:
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
	// The DNS information for the record.
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
	HttpPorts   *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts  *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether proxy acceleration is enabled for the record. This option is available only for CNAME, A, and AAAA records. Valid values:
	//
	// - **true**: Proxy acceleration is enabled.
	//
	// - **false**: Proxy acceleration is disabled.
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
	// The DNS type of the record, such as **A/AAAA**, **CNAME**, or **TXT**.
	//
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The type of origin for a CNAME record. This parameter is empty for other record types. Valid values:
	//
	// - **OSS**: An OSS origin.
	//
	// - **S3**: An S3 origin.
	//
	// - **LB**: A load balancer origin.
	//
	// - **OP**: An origin pool.
	//
	// - **Domain**: A domain name origin.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The TTL of the record in seconds. A value of 1 sets the TTL to Automatic.
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

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) GetHttpsPorts() *string {
	return s.HttpsPorts
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

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetHttpPorts(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.HttpPorts = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetHttpsPorts(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.HttpsPorts = &v
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateRecordsResponseBodyRecordResultListSuccessData struct {
	// The encryption algorithm used by the record. The value ranges from **0*	- to **255**. This parameter applies to CERT and SSHFP records.
	//
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The public key certificate for the record. This parameter applies to CERT, SMIMEA, and TLSA records.
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
	// The flag for the record, which indicates its priority and processing method. This parameter applies to CAA records.
	//
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The public key identifier for the record. The value ranges from **0*	- to **65535**. This parameter applies to CERT records.
	//
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The algorithm policy used by the record to match or validate certificates. The value ranges from **0*	- to **255**. This parameter applies to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// The port for the record. The value ranges from **0*	- to **65535**. This parameter applies only to SRV records.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the record. The value ranges from **0*	- to **65535**. A smaller value indicates a higher priority. This parameter applies to MX, SRV, and URI records.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of certificate or public key used by the record. The value ranges from **0*	- to **255**. This parameter applies to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The tag for a CAA record, which specifies its type and purpose, such as `issue`, `issuewild`, or `iodef`.
	//
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The certificate type for CERT records or the public key type for SSHFP records.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage identifier for the record. The value ranges from **0*	- to **255**. This parameter applies to SMIMEA and TLSA records.
	//
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The record value. This parameter applies to A/AAAA, CNAME, NS, MX, TXT, CAA, SRV, and URI records. The meaning of this parameter varies based on the record type:
	//
	// - **A/AAAA**: The IP address. To specify multiple addresses, separate them with a comma (,). At least one IPv4 address is required.
	//
	// - **CNAME**: The target domain name.
	//
	// - **NS**: The name server for the domain.
	//
	// - **MX**: The domain name of a valid target mail server.
	//
	// - **TXT**: A valid text string.
	//
	// - **CAA**: The domain name of a valid certificate authority.
	//
	// - **SRV**: The domain name of a valid target host.
	//
	// - **URI**: A valid URI string.
	//
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. The value ranges from **0*	- to **65535**. This parameter applies to SRV and URI records.
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
