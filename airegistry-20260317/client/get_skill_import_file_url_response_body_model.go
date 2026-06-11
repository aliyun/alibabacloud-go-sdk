// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillImportFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSkillImportFileUrlResponseBodyData) *GetSkillImportFileUrlResponseBody
	GetData() *GetSkillImportFileUrlResponseBodyData
	SetRequestId(v string) *GetSkillImportFileUrlResponseBody
	GetRequestId() *string
}

type GetSkillImportFileUrlResponseBody struct {
	Data      *GetSkillImportFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSkillImportFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillImportFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillImportFileUrlResponseBody) GetData() *GetSkillImportFileUrlResponseBodyData {
	return s.Data
}

func (s *GetSkillImportFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillImportFileUrlResponseBody) SetData(v *GetSkillImportFileUrlResponseBodyData) *GetSkillImportFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillImportFileUrlResponseBody) SetRequestId(v string) *GetSkillImportFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillImportFileUrlResponseBodyData struct {
	ContentType   *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	MaxSize       *string `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	UploadUrl     *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s GetSkillImportFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillImportFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillImportFileUrlResponseBodyData) GetContentType() *string {
	return s.ContentType
}

func (s *GetSkillImportFileUrlResponseBodyData) GetMaxSize() *string {
	return s.MaxSize
}

func (s *GetSkillImportFileUrlResponseBodyData) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *GetSkillImportFileUrlResponseBodyData) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetSkillImportFileUrlResponseBodyData) SetContentType(v string) *GetSkillImportFileUrlResponseBodyData {
	s.ContentType = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBodyData) SetMaxSize(v string) *GetSkillImportFileUrlResponseBodyData {
	s.MaxSize = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBodyData) SetOssObjectName(v string) *GetSkillImportFileUrlResponseBodyData {
	s.OssObjectName = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBodyData) SetUploadUrl(v string) *GetSkillImportFileUrlResponseBodyData {
	s.UploadUrl = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
