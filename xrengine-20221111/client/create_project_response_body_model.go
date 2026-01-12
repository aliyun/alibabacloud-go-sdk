// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateProjectResponseBody
	GetCode() *string
	SetData(v *CreateProjectResponseBodyData) *CreateProjectResponseBody
	GetData() *CreateProjectResponseBodyData
	SetMessage(v string) *CreateProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProjectResponseBody
	GetSuccess() *bool
}

type CreateProjectResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateProjectResponseBody) GetData() *CreateProjectResponseBodyData {
	return s.Data
}

func (s *CreateProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProjectResponseBody) SetCode(v string) *CreateProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProjectResponseBody) SetData(v *CreateProjectResponseBodyData) *CreateProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateProjectResponseBody) SetMessage(v string) *CreateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectResponseBodyData struct {
	CreateMode   *string                               `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime   *string                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset      *CreateProjectResponseBodyDataDataset `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Id           *string                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string                               `json:"Intro,omitempty" xml:"Intro,omitempty"`
	Method       *string                               `json:"Method,omitempty" xml:"Method,omitempty"`
	ModifiedTime *string                               `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source       *CreateProjectResponseBodyDataSource  `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status       *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string                               `json:"Title,omitempty" xml:"Title,omitempty"`
	Type         *string                               `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *CreateProjectResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateProjectResponseBodyData) GetDataset() *CreateProjectResponseBodyDataDataset {
	return s.Dataset
}

func (s *CreateProjectResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateProjectResponseBodyData) GetIntro() *string {
	return s.Intro
}

func (s *CreateProjectResponseBodyData) GetMethod() *string {
	return s.Method
}

func (s *CreateProjectResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *CreateProjectResponseBodyData) GetSource() *CreateProjectResponseBodyDataSource {
	return s.Source
}

func (s *CreateProjectResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateProjectResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *CreateProjectResponseBodyData) GetType() *string {
	return s.Type
}

func (s *CreateProjectResponseBodyData) SetCreateMode(v string) *CreateProjectResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetCreateTime(v string) *CreateProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetDataset(v *CreateProjectResponseBodyDataDataset) *CreateProjectResponseBodyData {
	s.Dataset = v
	return s
}

func (s *CreateProjectResponseBodyData) SetId(v string) *CreateProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetIntro(v string) *CreateProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetMethod(v string) *CreateProjectResponseBodyData {
	s.Method = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetModifiedTime(v string) *CreateProjectResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetSource(v *CreateProjectResponseBodyDataSource) *CreateProjectResponseBodyData {
	s.Source = v
	return s
}

func (s *CreateProjectResponseBodyData) SetStatus(v string) *CreateProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetTitle(v string) *CreateProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetType(v string) *CreateProjectResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateProjectResponseBodyData) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectResponseBodyDataDataset struct {
	CreateTime   *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id           *string                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                     `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                     `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *CreateProjectResponseBodyDataDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
}

func (s CreateProjectResponseBodyDataDataset) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataDataset) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateProjectResponseBodyDataDataset) GetId() *string {
	return s.Id
}

func (s *CreateProjectResponseBodyDataDataset) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *CreateProjectResponseBodyDataDataset) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateProjectResponseBodyDataDataset) GetPolicy() *CreateProjectResponseBodyDataDatasetPolicy {
	return s.Policy
}

func (s *CreateProjectResponseBodyDataDataset) SetCreateTime(v string) *CreateProjectResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *CreateProjectResponseBodyDataDataset) SetId(v string) *CreateProjectResponseBodyDataDataset {
	s.Id = &v
	return s
}

func (s *CreateProjectResponseBodyDataDataset) SetModifiedTime(v string) *CreateProjectResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *CreateProjectResponseBodyDataDataset) SetOssKey(v string) *CreateProjectResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *CreateProjectResponseBodyDataDataset) SetPolicy(v *CreateProjectResponseBodyDataDatasetPolicy) *CreateProjectResponseBodyDataDataset {
	s.Policy = v
	return s
}

func (s *CreateProjectResponseBodyDataDataset) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectResponseBodyDataDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s CreateProjectResponseBodyDataDatasetPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyDataDatasetPolicy) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) GetAccessId() *string {
	return s.AccessId
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) GetDir() *string {
	return s.Dir
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) GetExpire() *string {
	return s.Expire
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) GetHost() *string {
	return s.Host
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) GetPolicy() *string {
	return s.Policy
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) GetSignature() *string {
	return s.Signature
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) SetAccessId(v string) *CreateProjectResponseBodyDataDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) SetDir(v string) *CreateProjectResponseBodyDataDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) SetExpire(v string) *CreateProjectResponseBodyDataDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) SetHost(v string) *CreateProjectResponseBodyDataDatasetPolicy {
	s.Host = &v
	return s
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) SetPolicy(v string) *CreateProjectResponseBodyDataDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) SetSignature(v string) *CreateProjectResponseBodyDataDatasetPolicy {
	s.Signature = &v
	return s
}

func (s *CreateProjectResponseBodyDataDatasetPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateProjectResponseBodyDataSource struct {
	OssKey *string                                    `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy *CreateProjectResponseBodyDataSourcePolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
}

func (s CreateProjectResponseBodyDataSource) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataSource) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateProjectResponseBodyDataSource) GetPolicy() *CreateProjectResponseBodyDataSourcePolicy {
	return s.Policy
}

func (s *CreateProjectResponseBodyDataSource) SetOssKey(v string) *CreateProjectResponseBodyDataSource {
	s.OssKey = &v
	return s
}

func (s *CreateProjectResponseBodyDataSource) SetPolicy(v *CreateProjectResponseBodyDataSourcePolicy) *CreateProjectResponseBodyDataSource {
	s.Policy = v
	return s
}

func (s *CreateProjectResponseBodyDataSource) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectResponseBodyDataSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s CreateProjectResponseBodyDataSourcePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyDataSourcePolicy) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataSourcePolicy) GetAccessId() *string {
	return s.AccessId
}

func (s *CreateProjectResponseBodyDataSourcePolicy) GetDir() *string {
	return s.Dir
}

func (s *CreateProjectResponseBodyDataSourcePolicy) GetExpire() *string {
	return s.Expire
}

func (s *CreateProjectResponseBodyDataSourcePolicy) GetHost() *string {
	return s.Host
}

func (s *CreateProjectResponseBodyDataSourcePolicy) GetPolicy() *string {
	return s.Policy
}

func (s *CreateProjectResponseBodyDataSourcePolicy) GetSignature() *string {
	return s.Signature
}

func (s *CreateProjectResponseBodyDataSourcePolicy) SetAccessId(v string) *CreateProjectResponseBodyDataSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *CreateProjectResponseBodyDataSourcePolicy) SetDir(v string) *CreateProjectResponseBodyDataSourcePolicy {
	s.Dir = &v
	return s
}

func (s *CreateProjectResponseBodyDataSourcePolicy) SetExpire(v string) *CreateProjectResponseBodyDataSourcePolicy {
	s.Expire = &v
	return s
}

func (s *CreateProjectResponseBodyDataSourcePolicy) SetHost(v string) *CreateProjectResponseBodyDataSourcePolicy {
	s.Host = &v
	return s
}

func (s *CreateProjectResponseBodyDataSourcePolicy) SetPolicy(v string) *CreateProjectResponseBodyDataSourcePolicy {
	s.Policy = &v
	return s
}

func (s *CreateProjectResponseBodyDataSourcePolicy) SetSignature(v string) *CreateProjectResponseBodyDataSourcePolicy {
	s.Signature = &v
	return s
}

func (s *CreateProjectResponseBodyDataSourcePolicy) Validate() error {
	return dara.Validate(s)
}
