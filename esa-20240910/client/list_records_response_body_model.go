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
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records    []*ListRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AuthConf         *ListRecordsResponseBodyRecordsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	BizName          *string                                 `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Comment          *string                                 `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime       *string                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Data             *ListRecordsResponseBodyRecordsData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HostPolicy       *string                                 `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	Proxied          *bool                                   `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	RecordCname      *string                                 `json:"RecordCname,omitempty" xml:"RecordCname,omitempty"`
	RecordId         *int64                                  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RecordName       *string                                 `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	RecordSourceType *string                                 `json:"RecordSourceType,omitempty" xml:"RecordSourceType,omitempty"`
	RecordType       *string                                 `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	SiteId           *int64                                  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName         *string                                 `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	Ttl              *int64                                  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	UpdateTime       *string                                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
	Algorithm    *int32                 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Certificate  *string                `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Fingerprint  *string                `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Flag         *int32                 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	KeyTag       *int32                 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	MatchingType *int32                 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	Port         *int32                 `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority     *int32                 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Selector     *int32                 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	Tag          *string                `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Tags         map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type         *int32                 `json:"Type,omitempty" xml:"Type,omitempty"`
	Usage        *int32                 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Value        *string                `json:"Value,omitempty" xml:"Value,omitempty"`
	Weight       *int32                 `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
