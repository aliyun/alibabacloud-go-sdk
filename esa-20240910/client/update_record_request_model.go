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
	AuthConf *UpdateRecordRequestAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	BizName  *string                      `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Comment  *string                      `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	Data       *UpdateRecordRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HostPolicy *string                  `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// 是否代理加速
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// This parameter is required.
	RecordId   *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Ttl        *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
