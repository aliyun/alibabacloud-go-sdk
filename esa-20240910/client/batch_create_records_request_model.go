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
	// This parameter is required.
	RecordList []*BatchCreateRecordsRequestRecordList `json:"RecordList,omitempty" xml:"RecordList,omitempty" type:"Repeated"`
	// This parameter is required.
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
	AuthConf *BatchCreateRecordsRequestRecordListAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	BizName  *string                                      `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// This parameter is required.
	Data *BatchCreateRecordsRequestRecordListData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// This parameter is required.
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// This parameter is required.
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// This parameter is required.
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
	Algorithm    *int32  `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Certificate  *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Fingerprint  *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Flag         *int32  `json:"Flag,omitempty" xml:"Flag,omitempty"`
	KeyTag       *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	MatchingType *int32  `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	Port         *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Selector     *int32  `json:"Selector,omitempty" xml:"Selector,omitempty"`
	Tag          *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Usage        *int32  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Weight       *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
