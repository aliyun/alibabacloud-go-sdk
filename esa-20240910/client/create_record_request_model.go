// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConf(v *CreateRecordRequestAuthConf) *CreateRecordRequest
	GetAuthConf() *CreateRecordRequestAuthConf
	SetBizName(v string) *CreateRecordRequest
	GetBizName() *string
	SetComment(v string) *CreateRecordRequest
	GetComment() *string
	SetData(v *CreateRecordRequestData) *CreateRecordRequest
	GetData() *CreateRecordRequestData
	SetHostPolicy(v string) *CreateRecordRequest
	GetHostPolicy() *string
	SetProxied(v bool) *CreateRecordRequest
	GetProxied() *bool
	SetRecordName(v string) *CreateRecordRequest
	GetRecordName() *string
	SetSiteId(v int64) *CreateRecordRequest
	GetSiteId() *int64
	SetSourceType(v string) *CreateRecordRequest
	GetSourceType() *string
	SetTtl(v int32) *CreateRecordRequest
	GetTtl() *int32
	SetType(v string) *CreateRecordRequest
	GetType() *string
}

type CreateRecordRequest struct {
	AuthConf *CreateRecordRequestAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// 业务场景
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	Data       *CreateRecordRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HostPolicy *string                  `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// 是否代理加速
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// 记录名称
	//
	// This parameter is required.
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// 记录类型
	//
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordRequest) GetAuthConf() *CreateRecordRequestAuthConf {
	return s.AuthConf
}

func (s *CreateRecordRequest) GetBizName() *string {
	return s.BizName
}

func (s *CreateRecordRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateRecordRequest) GetData() *CreateRecordRequestData {
	return s.Data
}

func (s *CreateRecordRequest) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *CreateRecordRequest) GetProxied() *bool {
	return s.Proxied
}

func (s *CreateRecordRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *CreateRecordRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateRecordRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateRecordRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateRecordRequest) GetType() *string {
	return s.Type
}

func (s *CreateRecordRequest) SetAuthConf(v *CreateRecordRequestAuthConf) *CreateRecordRequest {
	s.AuthConf = v
	return s
}

func (s *CreateRecordRequest) SetBizName(v string) *CreateRecordRequest {
	s.BizName = &v
	return s
}

func (s *CreateRecordRequest) SetComment(v string) *CreateRecordRequest {
	s.Comment = &v
	return s
}

func (s *CreateRecordRequest) SetData(v *CreateRecordRequestData) *CreateRecordRequest {
	s.Data = v
	return s
}

func (s *CreateRecordRequest) SetHostPolicy(v string) *CreateRecordRequest {
	s.HostPolicy = &v
	return s
}

func (s *CreateRecordRequest) SetProxied(v bool) *CreateRecordRequest {
	s.Proxied = &v
	return s
}

func (s *CreateRecordRequest) SetRecordName(v string) *CreateRecordRequest {
	s.RecordName = &v
	return s
}

func (s *CreateRecordRequest) SetSiteId(v int64) *CreateRecordRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRecordRequest) SetSourceType(v string) *CreateRecordRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRecordRequest) SetTtl(v int32) *CreateRecordRequest {
	s.Ttl = &v
	return s
}

func (s *CreateRecordRequest) SetType(v string) *CreateRecordRequest {
	s.Type = &v
	return s
}

func (s *CreateRecordRequest) Validate() error {
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

type CreateRecordRequestAuthConf struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateRecordRequestAuthConf) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordRequestAuthConf) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateRecordRequestAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateRecordRequestAuthConf) GetRegion() *string {
	return s.Region
}

func (s *CreateRecordRequestAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *CreateRecordRequestAuthConf) GetVersion() *string {
	return s.Version
}

func (s *CreateRecordRequestAuthConf) SetAccessKey(v string) *CreateRecordRequestAuthConf {
	s.AccessKey = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetAuthType(v string) *CreateRecordRequestAuthConf {
	s.AuthType = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetRegion(v string) *CreateRecordRequestAuthConf {
	s.Region = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetSecretKey(v string) *CreateRecordRequestAuthConf {
	s.SecretKey = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetVersion(v string) *CreateRecordRequestAuthConf {
	s.Version = &v
	return s
}

func (s *CreateRecordRequestAuthConf) Validate() error {
	return dara.Validate(s)
}

type CreateRecordRequestData struct {
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

func (s CreateRecordRequestData) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordRequestData) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *CreateRecordRequestData) GetCertificate() *string {
	return s.Certificate
}

func (s *CreateRecordRequestData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *CreateRecordRequestData) GetFlag() *int32 {
	return s.Flag
}

func (s *CreateRecordRequestData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *CreateRecordRequestData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *CreateRecordRequestData) GetPort() *int32 {
	return s.Port
}

func (s *CreateRecordRequestData) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateRecordRequestData) GetSelector() *int32 {
	return s.Selector
}

func (s *CreateRecordRequestData) GetTag() *string {
	return s.Tag
}

func (s *CreateRecordRequestData) GetType() *int32 {
	return s.Type
}

func (s *CreateRecordRequestData) GetUsage() *int32 {
	return s.Usage
}

func (s *CreateRecordRequestData) GetValue() *string {
	return s.Value
}

func (s *CreateRecordRequestData) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateRecordRequestData) SetAlgorithm(v int32) *CreateRecordRequestData {
	s.Algorithm = &v
	return s
}

func (s *CreateRecordRequestData) SetCertificate(v string) *CreateRecordRequestData {
	s.Certificate = &v
	return s
}

func (s *CreateRecordRequestData) SetFingerprint(v string) *CreateRecordRequestData {
	s.Fingerprint = &v
	return s
}

func (s *CreateRecordRequestData) SetFlag(v int32) *CreateRecordRequestData {
	s.Flag = &v
	return s
}

func (s *CreateRecordRequestData) SetKeyTag(v int32) *CreateRecordRequestData {
	s.KeyTag = &v
	return s
}

func (s *CreateRecordRequestData) SetMatchingType(v int32) *CreateRecordRequestData {
	s.MatchingType = &v
	return s
}

func (s *CreateRecordRequestData) SetPort(v int32) *CreateRecordRequestData {
	s.Port = &v
	return s
}

func (s *CreateRecordRequestData) SetPriority(v int32) *CreateRecordRequestData {
	s.Priority = &v
	return s
}

func (s *CreateRecordRequestData) SetSelector(v int32) *CreateRecordRequestData {
	s.Selector = &v
	return s
}

func (s *CreateRecordRequestData) SetTag(v string) *CreateRecordRequestData {
	s.Tag = &v
	return s
}

func (s *CreateRecordRequestData) SetType(v int32) *CreateRecordRequestData {
	s.Type = &v
	return s
}

func (s *CreateRecordRequestData) SetUsage(v int32) *CreateRecordRequestData {
	s.Usage = &v
	return s
}

func (s *CreateRecordRequestData) SetValue(v string) *CreateRecordRequestData {
	s.Value = &v
	return s
}

func (s *CreateRecordRequestData) SetWeight(v int32) *CreateRecordRequestData {
	s.Weight = &v
	return s
}

func (s *CreateRecordRequestData) Validate() error {
	return dara.Validate(s)
}
