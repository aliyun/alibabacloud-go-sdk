// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSourceFileResponseBody
	GetCode() *string
	SetData(v *ListSourceFileResponseBodyData) *ListSourceFileResponseBody
	GetData() *ListSourceFileResponseBodyData
	SetMessage(v string) *ListSourceFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSourceFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSourceFileResponseBody
	GetSuccess() *bool
}

type ListSourceFileResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListSourceFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSourceFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSourceFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListSourceFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSourceFileResponseBody) GetData() *ListSourceFileResponseBodyData {
	return s.Data
}

func (s *ListSourceFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSourceFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSourceFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSourceFileResponseBody) SetCode(v string) *ListSourceFileResponseBody {
	s.Code = &v
	return s
}

func (s *ListSourceFileResponseBody) SetData(v *ListSourceFileResponseBodyData) *ListSourceFileResponseBody {
	s.Data = v
	return s
}

func (s *ListSourceFileResponseBody) SetMessage(v string) *ListSourceFileResponseBody {
	s.Message = &v
	return s
}

func (s *ListSourceFileResponseBody) SetRequestId(v string) *ListSourceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSourceFileResponseBody) SetSuccess(v bool) *ListSourceFileResponseBody {
	s.Success = &v
	return s
}

func (s *ListSourceFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSourceFileResponseBodyData struct {
	PicList   []*ListSourceFileResponseBodyDataPicList   `json:"PicList,omitempty" xml:"PicList,omitempty" type:"Repeated"`
	VideoList []*ListSourceFileResponseBodyDataVideoList `json:"VideoList,omitempty" xml:"VideoList,omitempty" type:"Repeated"`
}

func (s ListSourceFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSourceFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSourceFileResponseBodyData) GetPicList() []*ListSourceFileResponseBodyDataPicList {
	return s.PicList
}

func (s *ListSourceFileResponseBodyData) GetVideoList() []*ListSourceFileResponseBodyDataVideoList {
	return s.VideoList
}

func (s *ListSourceFileResponseBodyData) SetPicList(v []*ListSourceFileResponseBodyDataPicList) *ListSourceFileResponseBodyData {
	s.PicList = v
	return s
}

func (s *ListSourceFileResponseBodyData) SetVideoList(v []*ListSourceFileResponseBodyDataVideoList) *ListSourceFileResponseBodyData {
	s.VideoList = v
	return s
}

func (s *ListSourceFileResponseBodyData) Validate() error {
	if s.PicList != nil {
		for _, item := range s.PicList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoList != nil {
		for _, item := range s.VideoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSourceFileResponseBodyDataPicList struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListSourceFileResponseBodyDataPicList) String() string {
	return dara.Prettify(s)
}

func (s ListSourceFileResponseBodyDataPicList) GoString() string {
	return s.String()
}

func (s *ListSourceFileResponseBodyDataPicList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSourceFileResponseBodyDataPicList) GetFileName() *string {
	return s.FileName
}

func (s *ListSourceFileResponseBodyDataPicList) GetId() *string {
	return s.Id
}

func (s *ListSourceFileResponseBodyDataPicList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListSourceFileResponseBodyDataPicList) GetOssKey() *string {
	return s.OssKey
}

func (s *ListSourceFileResponseBodyDataPicList) GetType() *string {
	return s.Type
}

func (s *ListSourceFileResponseBodyDataPicList) GetUrl() *string {
	return s.Url
}

func (s *ListSourceFileResponseBodyDataPicList) SetCreateTime(v string) *ListSourceFileResponseBodyDataPicList {
	s.CreateTime = &v
	return s
}

func (s *ListSourceFileResponseBodyDataPicList) SetFileName(v string) *ListSourceFileResponseBodyDataPicList {
	s.FileName = &v
	return s
}

func (s *ListSourceFileResponseBodyDataPicList) SetId(v string) *ListSourceFileResponseBodyDataPicList {
	s.Id = &v
	return s
}

func (s *ListSourceFileResponseBodyDataPicList) SetModifiedTime(v string) *ListSourceFileResponseBodyDataPicList {
	s.ModifiedTime = &v
	return s
}

func (s *ListSourceFileResponseBodyDataPicList) SetOssKey(v string) *ListSourceFileResponseBodyDataPicList {
	s.OssKey = &v
	return s
}

func (s *ListSourceFileResponseBodyDataPicList) SetType(v string) *ListSourceFileResponseBodyDataPicList {
	s.Type = &v
	return s
}

func (s *ListSourceFileResponseBodyDataPicList) SetUrl(v string) *ListSourceFileResponseBodyDataPicList {
	s.Url = &v
	return s
}

func (s *ListSourceFileResponseBodyDataPicList) Validate() error {
	return dara.Validate(s)
}

type ListSourceFileResponseBodyDataVideoList struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListSourceFileResponseBodyDataVideoList) String() string {
	return dara.Prettify(s)
}

func (s ListSourceFileResponseBodyDataVideoList) GoString() string {
	return s.String()
}

func (s *ListSourceFileResponseBodyDataVideoList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSourceFileResponseBodyDataVideoList) GetFileName() *string {
	return s.FileName
}

func (s *ListSourceFileResponseBodyDataVideoList) GetId() *string {
	return s.Id
}

func (s *ListSourceFileResponseBodyDataVideoList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListSourceFileResponseBodyDataVideoList) GetOssKey() *string {
	return s.OssKey
}

func (s *ListSourceFileResponseBodyDataVideoList) GetType() *string {
	return s.Type
}

func (s *ListSourceFileResponseBodyDataVideoList) GetUrl() *string {
	return s.Url
}

func (s *ListSourceFileResponseBodyDataVideoList) SetCreateTime(v string) *ListSourceFileResponseBodyDataVideoList {
	s.CreateTime = &v
	return s
}

func (s *ListSourceFileResponseBodyDataVideoList) SetFileName(v string) *ListSourceFileResponseBodyDataVideoList {
	s.FileName = &v
	return s
}

func (s *ListSourceFileResponseBodyDataVideoList) SetId(v string) *ListSourceFileResponseBodyDataVideoList {
	s.Id = &v
	return s
}

func (s *ListSourceFileResponseBodyDataVideoList) SetModifiedTime(v string) *ListSourceFileResponseBodyDataVideoList {
	s.ModifiedTime = &v
	return s
}

func (s *ListSourceFileResponseBodyDataVideoList) SetOssKey(v string) *ListSourceFileResponseBodyDataVideoList {
	s.OssKey = &v
	return s
}

func (s *ListSourceFileResponseBodyDataVideoList) SetType(v string) *ListSourceFileResponseBodyDataVideoList {
	s.Type = &v
	return s
}

func (s *ListSourceFileResponseBodyDataVideoList) SetUrl(v string) *ListSourceFileResponseBodyDataVideoList {
	s.Url = &v
	return s
}

func (s *ListSourceFileResponseBodyDataVideoList) Validate() error {
	return dara.Validate(s)
}
