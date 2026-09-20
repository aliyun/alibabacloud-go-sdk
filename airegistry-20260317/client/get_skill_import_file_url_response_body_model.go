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
	// The returned result.
	Data *GetSkillImportFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D9E87E66-9EF0-5C10-A5E6-924020A0C9B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The file type.
	//
	// example:
	//
	// application/zip
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The maximum allowed file size for upload, in MB.
	//
	// example:
	//
	// 10
	MaxSize *string `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The file name in the authorized OSS bucket.
	//
	// > Among the four image input methods (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS), select one to use.
	//
	// example:
	//
	// 1190239587066411/skill/import/5e993afe-f629-4619-9ac2-51b125300cdd/2026/06/09/35059076-5992-4a71-a706-89230e57f2a2/ui-ux-pro-max.zip
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// The file upload URL. The client uses this URL to upload the file.
	//
	// example:
	//
	// https://mse-shared-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/skill/import/199xxxxxxxx0842/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx/2026/06/10/xxxx-xxxx-xxxx-xxxx-xxxxxxxxxx/1781082579097.zip?Expires=1781083479&OSSAccessKeyId=STS.NZXGXTD2yoDLd5PfsYxjFrvBJ&Signature=Loyyzzzzzzzz%3D&security-token=CAIStgxxxxxxx
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
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
