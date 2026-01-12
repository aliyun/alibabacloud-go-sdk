// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProjectDatasetResponseBody
	GetCode() *string
	SetData(v *GetProjectDatasetResponseBodyData) *GetProjectDatasetResponseBody
	GetData() *GetProjectDatasetResponseBodyData
	SetMessage(v string) *GetProjectDatasetResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetProjectDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProjectDatasetResponseBody
	GetSuccess() *bool
}

type GetProjectDatasetResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetProjectDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProjectDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProjectDatasetResponseBody) GetData() *GetProjectDatasetResponseBodyData {
	return s.Data
}

func (s *GetProjectDatasetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProjectDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectDatasetResponseBody) SetCode(v string) *GetProjectDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectDatasetResponseBody) SetData(v *GetProjectDatasetResponseBodyData) *GetProjectDatasetResponseBody {
	s.Data = v
	return s
}

func (s *GetProjectDatasetResponseBody) SetMessage(v string) *GetProjectDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *GetProjectDatasetResponseBody) SetRequestId(v string) *GetProjectDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectDatasetResponseBody) SetSuccess(v bool) *GetProjectDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectDatasetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectDatasetResponseBodyData struct {
	CreateMode   *string                                   `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime   *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset      *GetProjectDatasetResponseBodyDataDataset `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Id           *string                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string                                   `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ModifiedTime *string                                   `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status       *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string                                   `json:"Title,omitempty" xml:"Title,omitempty"`
	Type         *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	BizUsage     *string                                   `json:"bizUsage,omitempty" xml:"bizUsage,omitempty"`
}

func (s GetProjectDatasetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProjectDatasetResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *GetProjectDatasetResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProjectDatasetResponseBodyData) GetDataset() *GetProjectDatasetResponseBodyDataDataset {
	return s.Dataset
}

func (s *GetProjectDatasetResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetProjectDatasetResponseBodyData) GetIntro() *string {
	return s.Intro
}

func (s *GetProjectDatasetResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetProjectDatasetResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetProjectDatasetResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetProjectDatasetResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetProjectDatasetResponseBodyData) GetBizUsage() *string {
	return s.BizUsage
}

func (s *GetProjectDatasetResponseBodyData) SetCreateMode(v string) *GetProjectDatasetResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *GetProjectDatasetResponseBodyData) SetCreateTime(v string) *GetProjectDatasetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetProjectDatasetResponseBodyData) SetDataset(v *GetProjectDatasetResponseBodyDataDataset) *GetProjectDatasetResponseBodyData {
	s.Dataset = v
	return s
}

func (s *GetProjectDatasetResponseBodyData) SetId(v string) *GetProjectDatasetResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetProjectDatasetResponseBodyData) SetIntro(v string) *GetProjectDatasetResponseBodyData {
	s.Intro = &v
	return s
}

func (s *GetProjectDatasetResponseBodyData) SetModifiedTime(v string) *GetProjectDatasetResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetProjectDatasetResponseBodyData) SetStatus(v string) *GetProjectDatasetResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetProjectDatasetResponseBodyData) SetTitle(v string) *GetProjectDatasetResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetProjectDatasetResponseBodyData) SetType(v string) *GetProjectDatasetResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetProjectDatasetResponseBodyData) SetBizUsage(v string) *GetProjectDatasetResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *GetProjectDatasetResponseBodyData) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectDatasetResponseBodyDataDataset struct {
	CoverUrl        *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelUrl        *string `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PoseUrl         *string `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s GetProjectDatasetResponseBodyDataDataset) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDatasetResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *GetProjectDatasetResponseBodyDataDataset) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *GetProjectDatasetResponseBodyDataDataset) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProjectDatasetResponseBodyDataDataset) GetId() *string {
	return s.Id
}

func (s *GetProjectDatasetResponseBodyDataDataset) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *GetProjectDatasetResponseBodyDataDataset) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetProjectDatasetResponseBodyDataDataset) GetOriginResultUrl() *string {
	return s.OriginResultUrl
}

func (s *GetProjectDatasetResponseBodyDataDataset) GetOssKey() *string {
	return s.OssKey
}

func (s *GetProjectDatasetResponseBodyDataDataset) GetPoseUrl() *string {
	return s.PoseUrl
}

func (s *GetProjectDatasetResponseBodyDataDataset) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *GetProjectDatasetResponseBodyDataDataset) SetCoverUrl(v string) *GetProjectDatasetResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *GetProjectDatasetResponseBodyDataDataset) SetCreateTime(v string) *GetProjectDatasetResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *GetProjectDatasetResponseBodyDataDataset) SetId(v string) *GetProjectDatasetResponseBodyDataDataset {
	s.Id = &v
	return s
}

func (s *GetProjectDatasetResponseBodyDataDataset) SetModelUrl(v string) *GetProjectDatasetResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *GetProjectDatasetResponseBodyDataDataset) SetModifiedTime(v string) *GetProjectDatasetResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *GetProjectDatasetResponseBodyDataDataset) SetOriginResultUrl(v string) *GetProjectDatasetResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *GetProjectDatasetResponseBodyDataDataset) SetOssKey(v string) *GetProjectDatasetResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *GetProjectDatasetResponseBodyDataDataset) SetPoseUrl(v string) *GetProjectDatasetResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *GetProjectDatasetResponseBodyDataDataset) SetPreviewUrl(v string) *GetProjectDatasetResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

func (s *GetProjectDatasetResponseBodyDataDataset) Validate() error {
	return dara.Validate(s)
}
