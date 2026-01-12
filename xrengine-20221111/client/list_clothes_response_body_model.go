// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClothesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListClothesResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListClothesResponseBody
	GetCurrent() *int32
	SetData(v []*ListClothesResponseBodyData) *ListClothesResponseBody
	GetData() []*ListClothesResponseBodyData
	SetErrorName(v string) *ListClothesResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *ListClothesResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ListClothesResponseBody
	GetMessage() *string
	SetPages(v int32) *ListClothesResponseBody
	GetPages() *int32
	SetRequestId(v string) *ListClothesResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListClothesResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListClothesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListClothesResponseBody
	GetTotal() *int32
}

type ListClothesResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                         `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*ListClothesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                        `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                         `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                         `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                         `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClothesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClothesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClothesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListClothesResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListClothesResponseBody) GetData() []*ListClothesResponseBodyData {
	return s.Data
}

func (s *ListClothesResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *ListClothesResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListClothesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClothesResponseBody) GetPages() *int32 {
	return s.Pages
}

func (s *ListClothesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClothesResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListClothesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClothesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListClothesResponseBody) SetCode(v string) *ListClothesResponseBody {
	s.Code = &v
	return s
}

func (s *ListClothesResponseBody) SetCurrent(v int32) *ListClothesResponseBody {
	s.Current = &v
	return s
}

func (s *ListClothesResponseBody) SetData(v []*ListClothesResponseBodyData) *ListClothesResponseBody {
	s.Data = v
	return s
}

func (s *ListClothesResponseBody) SetErrorName(v string) *ListClothesResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ListClothesResponseBody) SetHttpCode(v int32) *ListClothesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListClothesResponseBody) SetMessage(v string) *ListClothesResponseBody {
	s.Message = &v
	return s
}

func (s *ListClothesResponseBody) SetPages(v int32) *ListClothesResponseBody {
	s.Pages = &v
	return s
}

func (s *ListClothesResponseBody) SetRequestId(v string) *ListClothesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClothesResponseBody) SetSize(v int32) *ListClothesResponseBody {
	s.Size = &v
	return s
}

func (s *ListClothesResponseBody) SetSuccess(v bool) *ListClothesResponseBody {
	s.Success = &v
	return s
}

func (s *ListClothesResponseBody) SetTotal(v int32) *ListClothesResponseBody {
	s.Total = &v
	return s
}

func (s *ListClothesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClothesResponseBodyData struct {
	CoverUrl     *string            `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id           *string            `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string            `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Part         *string            `json:"Part,omitempty" xml:"Part,omitempty"`
	Size         *string            `json:"Size,omitempty" xml:"Size,omitempty"`
	Status       map[string]*string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListClothesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClothesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClothesResponseBodyData) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListClothesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClothesResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListClothesResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListClothesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListClothesResponseBodyData) GetOssKey() *string {
	return s.OssKey
}

func (s *ListClothesResponseBodyData) GetPart() *string {
	return s.Part
}

func (s *ListClothesResponseBodyData) GetSize() *string {
	return s.Size
}

func (s *ListClothesResponseBodyData) GetStatus() map[string]*string {
	return s.Status
}

func (s *ListClothesResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListClothesResponseBodyData) SetCoverUrl(v string) *ListClothesResponseBodyData {
	s.CoverUrl = &v
	return s
}

func (s *ListClothesResponseBodyData) SetCreateTime(v string) *ListClothesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListClothesResponseBodyData) SetId(v string) *ListClothesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListClothesResponseBodyData) SetModifiedTime(v string) *ListClothesResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *ListClothesResponseBodyData) SetName(v string) *ListClothesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListClothesResponseBodyData) SetOssKey(v string) *ListClothesResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *ListClothesResponseBodyData) SetPart(v string) *ListClothesResponseBodyData {
	s.Part = &v
	return s
}

func (s *ListClothesResponseBodyData) SetSize(v string) *ListClothesResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListClothesResponseBodyData) SetStatus(v map[string]*string) *ListClothesResponseBodyData {
	s.Status = v
	return s
}

func (s *ListClothesResponseBodyData) SetType(v string) *ListClothesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListClothesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
