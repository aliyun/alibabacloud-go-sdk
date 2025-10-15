// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *PreviewDocumentResponseBody
	GetCost() *int64
	SetData(v *PreviewDocumentResponseBodyData) *PreviewDocumentResponseBody
	GetData() *PreviewDocumentResponseBodyData
	SetDataType(v string) *PreviewDocumentResponseBody
	GetDataType() *string
	SetErrCode(v string) *PreviewDocumentResponseBody
	GetErrCode() *string
	SetMessage(v string) *PreviewDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *PreviewDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PreviewDocumentResponseBody
	GetSuccess() *bool
	SetTime(v string) *PreviewDocumentResponseBody
	GetTime() *string
}

type PreviewDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *PreviewDocumentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// ff551395-1c8a-4f30-8ffd-ef7e87c70b4c
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s PreviewDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewDocumentResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *PreviewDocumentResponseBody) GetData() *PreviewDocumentResponseBodyData {
	return s.Data
}

func (s *PreviewDocumentResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *PreviewDocumentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *PreviewDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PreviewDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PreviewDocumentResponseBody) GetTime() *string {
	return s.Time
}

func (s *PreviewDocumentResponseBody) SetCost(v int64) *PreviewDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetData(v *PreviewDocumentResponseBodyData) *PreviewDocumentResponseBody {
	s.Data = v
	return s
}

func (s *PreviewDocumentResponseBody) SetDataType(v string) *PreviewDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetErrCode(v string) *PreviewDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetMessage(v string) *PreviewDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetRequestId(v string) *PreviewDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetSuccess(v bool) *PreviewDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetTime(v string) *PreviewDocumentResponseBody {
	s.Time = &v
	return s
}

func (s *PreviewDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewDocumentResponseBodyData struct {
	// example:
	//
	// pdf
	PreviewType *string `json:"previewType,omitempty" xml:"previewType,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	UploadTime *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty"`
	// example:
	//
	// https://agi.alicdn.com/user/d0o/d3c1f50d-a6c2-49b3-b0c8-3e613c3f20ee_16872_3236784461.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PreviewDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PreviewDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *PreviewDocumentResponseBodyData) GetPreviewType() *string {
	return s.PreviewType
}

func (s *PreviewDocumentResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *PreviewDocumentResponseBodyData) GetUploadTime() *string {
	return s.UploadTime
}

func (s *PreviewDocumentResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *PreviewDocumentResponseBodyData) SetPreviewType(v string) *PreviewDocumentResponseBodyData {
	s.PreviewType = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) SetTitle(v string) *PreviewDocumentResponseBodyData {
	s.Title = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) SetUploadTime(v string) *PreviewDocumentResponseBodyData {
	s.UploadTime = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) SetUrl(v string) *PreviewDocumentResponseBodyData {
	s.Url = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
