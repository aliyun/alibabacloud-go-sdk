// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAreaMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAreaMapResponseBody
	GetCode() *string
	SetData(v *QueryAreaMapResponseBodyData) *QueryAreaMapResponseBody
	GetData() *QueryAreaMapResponseBodyData
	SetErrorName(v string) *QueryAreaMapResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *QueryAreaMapResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *QueryAreaMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAreaMapResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAreaMapResponseBody
	GetSuccess() *bool
}

type QueryAreaMapResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryAreaMapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                       `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                        `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAreaMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAreaMapResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAreaMapResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAreaMapResponseBody) GetData() *QueryAreaMapResponseBodyData {
	return s.Data
}

func (s *QueryAreaMapResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *QueryAreaMapResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *QueryAreaMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAreaMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAreaMapResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAreaMapResponseBody) SetCode(v string) *QueryAreaMapResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAreaMapResponseBody) SetData(v *QueryAreaMapResponseBodyData) *QueryAreaMapResponseBody {
	s.Data = v
	return s
}

func (s *QueryAreaMapResponseBody) SetErrorName(v string) *QueryAreaMapResponseBody {
	s.ErrorName = &v
	return s
}

func (s *QueryAreaMapResponseBody) SetHttpCode(v int32) *QueryAreaMapResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryAreaMapResponseBody) SetMessage(v string) *QueryAreaMapResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAreaMapResponseBody) SetRequestId(v string) *QueryAreaMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAreaMapResponseBody) SetSuccess(v bool) *QueryAreaMapResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAreaMapResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAreaMapResponseBodyData struct {
	D3Oss       *string `json:"D3Oss,omitempty" xml:"D3Oss,omitempty"`
	ExtInfo     *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LocationOss *string `json:"LocationOss,omitempty" xml:"LocationOss,omitempty"`
	MapType     *string `json:"MapType,omitempty" xml:"MapType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note        *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryAreaMapResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAreaMapResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAreaMapResponseBodyData) GetD3Oss() *string {
	return s.D3Oss
}

func (s *QueryAreaMapResponseBodyData) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *QueryAreaMapResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QueryAreaMapResponseBodyData) GetLocationOss() *string {
	return s.LocationOss
}

func (s *QueryAreaMapResponseBodyData) GetMapType() *string {
	return s.MapType
}

func (s *QueryAreaMapResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryAreaMapResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *QueryAreaMapResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryAreaMapResponseBodyData) SetD3Oss(v string) *QueryAreaMapResponseBodyData {
	s.D3Oss = &v
	return s
}

func (s *QueryAreaMapResponseBodyData) SetExtInfo(v string) *QueryAreaMapResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *QueryAreaMapResponseBodyData) SetId(v int64) *QueryAreaMapResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryAreaMapResponseBodyData) SetLocationOss(v string) *QueryAreaMapResponseBodyData {
	s.LocationOss = &v
	return s
}

func (s *QueryAreaMapResponseBodyData) SetMapType(v string) *QueryAreaMapResponseBodyData {
	s.MapType = &v
	return s
}

func (s *QueryAreaMapResponseBodyData) SetName(v string) *QueryAreaMapResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryAreaMapResponseBodyData) SetNote(v string) *QueryAreaMapResponseBodyData {
	s.Note = &v
	return s
}

func (s *QueryAreaMapResponseBodyData) SetStatus(v string) *QueryAreaMapResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryAreaMapResponseBodyData) Validate() error {
	return dara.Validate(s)
}
