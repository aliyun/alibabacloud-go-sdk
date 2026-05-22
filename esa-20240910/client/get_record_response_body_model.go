// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordModel(v *GetRecordResponseBodyRecordModel) *GetRecordResponseBody
	GetRecordModel() *GetRecordResponseBodyRecordModel
	SetRequestId(v string) *GetRecordResponseBody
	GetRequestId() *string
}

type GetRecordResponseBody struct {
	RecordModel *GetRecordResponseBodyRecordModel `json:"RecordModel,omitempty" xml:"RecordModel,omitempty" type:"Struct"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBody) GetRecordModel() *GetRecordResponseBodyRecordModel {
	return s.RecordModel
}

func (s *GetRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecordResponseBody) SetRecordModel(v *GetRecordResponseBodyRecordModel) *GetRecordResponseBody {
	s.RecordModel = v
	return s
}

func (s *GetRecordResponseBody) SetRequestId(v string) *GetRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordResponseBody) Validate() error {
	if s.RecordModel != nil {
		if err := s.RecordModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRecordResponseBodyRecordModel struct {
	AuthConf         *GetRecordResponseBodyRecordModelAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	BizName          *string                                   `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Comment          *string                                   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime       *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Data             *GetRecordResponseBodyRecordModelData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HostPolicy       *string                                   `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	Proxied          *bool                                     `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	RecordCname      *string                                   `json:"RecordCname,omitempty" xml:"RecordCname,omitempty"`
	RecordId         *int64                                    `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RecordName       *string                                   `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	RecordSourceType *string                                   `json:"RecordSourceType,omitempty" xml:"RecordSourceType,omitempty"`
	RecordType       *string                                   `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	SiteId           *int64                                    `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName         *string                                   `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	Ttl              *int32                                    `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	UpdateTime       *string                                   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetRecordResponseBodyRecordModel) String() string {
	return dara.Prettify(s)
}

func (s GetRecordResponseBodyRecordModel) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBodyRecordModel) GetAuthConf() *GetRecordResponseBodyRecordModelAuthConf {
	return s.AuthConf
}

func (s *GetRecordResponseBodyRecordModel) GetBizName() *string {
	return s.BizName
}

func (s *GetRecordResponseBodyRecordModel) GetComment() *string {
	return s.Comment
}

func (s *GetRecordResponseBodyRecordModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRecordResponseBodyRecordModel) GetData() *GetRecordResponseBodyRecordModelData {
	return s.Data
}

func (s *GetRecordResponseBodyRecordModel) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *GetRecordResponseBodyRecordModel) GetProxied() *bool {
	return s.Proxied
}

func (s *GetRecordResponseBodyRecordModel) GetRecordCname() *string {
	return s.RecordCname
}

func (s *GetRecordResponseBodyRecordModel) GetRecordId() *int64 {
	return s.RecordId
}

func (s *GetRecordResponseBodyRecordModel) GetRecordName() *string {
	return s.RecordName
}

func (s *GetRecordResponseBodyRecordModel) GetRecordSourceType() *string {
	return s.RecordSourceType
}

func (s *GetRecordResponseBodyRecordModel) GetRecordType() *string {
	return s.RecordType
}

func (s *GetRecordResponseBodyRecordModel) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetRecordResponseBodyRecordModel) GetSiteName() *string {
	return s.SiteName
}

func (s *GetRecordResponseBodyRecordModel) GetTtl() *int32 {
	return s.Ttl
}

func (s *GetRecordResponseBodyRecordModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetRecordResponseBodyRecordModel) SetAuthConf(v *GetRecordResponseBodyRecordModelAuthConf) *GetRecordResponseBodyRecordModel {
	s.AuthConf = v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetBizName(v string) *GetRecordResponseBodyRecordModel {
	s.BizName = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetComment(v string) *GetRecordResponseBodyRecordModel {
	s.Comment = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetCreateTime(v string) *GetRecordResponseBodyRecordModel {
	s.CreateTime = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetData(v *GetRecordResponseBodyRecordModelData) *GetRecordResponseBodyRecordModel {
	s.Data = v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetHostPolicy(v string) *GetRecordResponseBodyRecordModel {
	s.HostPolicy = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetProxied(v bool) *GetRecordResponseBodyRecordModel {
	s.Proxied = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordCname(v string) *GetRecordResponseBodyRecordModel {
	s.RecordCname = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordId(v int64) *GetRecordResponseBodyRecordModel {
	s.RecordId = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordName(v string) *GetRecordResponseBodyRecordModel {
	s.RecordName = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordSourceType(v string) *GetRecordResponseBodyRecordModel {
	s.RecordSourceType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordType(v string) *GetRecordResponseBodyRecordModel {
	s.RecordType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetSiteId(v int64) *GetRecordResponseBodyRecordModel {
	s.SiteId = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetSiteName(v string) *GetRecordResponseBodyRecordModel {
	s.SiteName = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetTtl(v int32) *GetRecordResponseBodyRecordModel {
	s.Ttl = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetUpdateTime(v string) *GetRecordResponseBodyRecordModel {
	s.UpdateTime = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) Validate() error {
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

type GetRecordResponseBodyRecordModelAuthConf struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetRecordResponseBodyRecordModelAuthConf) String() string {
	return dara.Prettify(s)
}

func (s GetRecordResponseBodyRecordModelAuthConf) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetRegion() *string {
	return s.Region
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *GetRecordResponseBodyRecordModelAuthConf) GetVersion() *string {
	return s.Version
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetAccessKey(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.AccessKey = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetAuthType(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.AuthType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetRegion(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.Region = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetSecretKey(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.SecretKey = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetVersion(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.Version = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) Validate() error {
	return dara.Validate(s)
}

type GetRecordResponseBodyRecordModelData struct {
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

func (s GetRecordResponseBodyRecordModelData) String() string {
	return dara.Prettify(s)
}

func (s GetRecordResponseBodyRecordModelData) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBodyRecordModelData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *GetRecordResponseBodyRecordModelData) GetCertificate() *string {
	return s.Certificate
}

func (s *GetRecordResponseBodyRecordModelData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *GetRecordResponseBodyRecordModelData) GetFlag() *int32 {
	return s.Flag
}

func (s *GetRecordResponseBodyRecordModelData) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *GetRecordResponseBodyRecordModelData) GetMatchingType() *int32 {
	return s.MatchingType
}

func (s *GetRecordResponseBodyRecordModelData) GetPort() *int32 {
	return s.Port
}

func (s *GetRecordResponseBodyRecordModelData) GetPriority() *int32 {
	return s.Priority
}

func (s *GetRecordResponseBodyRecordModelData) GetSelector() *int32 {
	return s.Selector
}

func (s *GetRecordResponseBodyRecordModelData) GetTag() *string {
	return s.Tag
}

func (s *GetRecordResponseBodyRecordModelData) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetRecordResponseBodyRecordModelData) GetType() *int32 {
	return s.Type
}

func (s *GetRecordResponseBodyRecordModelData) GetUsage() *int32 {
	return s.Usage
}

func (s *GetRecordResponseBodyRecordModelData) GetValue() *string {
	return s.Value
}

func (s *GetRecordResponseBodyRecordModelData) GetWeight() *int32 {
	return s.Weight
}

func (s *GetRecordResponseBodyRecordModelData) SetAlgorithm(v int32) *GetRecordResponseBodyRecordModelData {
	s.Algorithm = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetCertificate(v string) *GetRecordResponseBodyRecordModelData {
	s.Certificate = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetFingerprint(v string) *GetRecordResponseBodyRecordModelData {
	s.Fingerprint = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetFlag(v int32) *GetRecordResponseBodyRecordModelData {
	s.Flag = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetKeyTag(v int32) *GetRecordResponseBodyRecordModelData {
	s.KeyTag = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetMatchingType(v int32) *GetRecordResponseBodyRecordModelData {
	s.MatchingType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetPort(v int32) *GetRecordResponseBodyRecordModelData {
	s.Port = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetPriority(v int32) *GetRecordResponseBodyRecordModelData {
	s.Priority = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetSelector(v int32) *GetRecordResponseBodyRecordModelData {
	s.Selector = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetTag(v string) *GetRecordResponseBodyRecordModelData {
	s.Tag = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetTags(v map[string]interface{}) *GetRecordResponseBodyRecordModelData {
	s.Tags = v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetType(v int32) *GetRecordResponseBodyRecordModelData {
	s.Type = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetUsage(v int32) *GetRecordResponseBodyRecordModelData {
	s.Usage = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetValue(v string) *GetRecordResponseBodyRecordModelData {
	s.Value = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetWeight(v int32) *GetRecordResponseBodyRecordModelData {
	s.Weight = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) Validate() error {
	return dara.Validate(s)
}
