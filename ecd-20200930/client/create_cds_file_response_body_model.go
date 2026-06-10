// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdsFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileModel(v *CreateCdsFileResponseBodyFileModel) *CreateCdsFileResponseBody
	GetFileModel() *CreateCdsFileResponseBodyFileModel
	SetRequestId(v string) *CreateCdsFileResponseBody
	GetRequestId() *string
}

type CreateCdsFileResponseBody struct {
	// File upload response object.
	FileModel *CreateCdsFileResponseBodyFileModel `json:"FileModel,omitempty" xml:"FileModel,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 2BAFE05D-FFB9-5938-96D0-08017DB9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCdsFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCdsFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdsFileResponseBody) GetFileModel() *CreateCdsFileResponseBodyFileModel {
	return s.FileModel
}

func (s *CreateCdsFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCdsFileResponseBody) SetFileModel(v *CreateCdsFileResponseBodyFileModel) *CreateCdsFileResponseBody {
	s.FileModel = v
	return s
}

func (s *CreateCdsFileResponseBody) SetRequestId(v string) *CreateCdsFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCdsFileResponseBody) Validate() error {
	if s.FileModel != nil {
		if err := s.FileModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCdsFileResponseBodyFileModel struct {
	// File ID.
	//
	// example:
	//
	// 646ec2f662b1e08f7a2b4d648982f906****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// File upload task ID.
	//
	// example:
	//
	// 20BE155AC20943BE9D405B0997A9****
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	// File upload URL. Valid for 15 minutes by default.
	//
	// example:
	//
	// https://data.aliyunpds.com/xxx/xxx?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx&partNumber=1&uploadId=0CC175B9C0F1B6A831C399E26977****
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s CreateCdsFileResponseBodyFileModel) String() string {
	return dara.Prettify(s)
}

func (s CreateCdsFileResponseBodyFileModel) GoString() string {
	return s.String()
}

func (s *CreateCdsFileResponseBodyFileModel) GetFileId() *string {
	return s.FileId
}

func (s *CreateCdsFileResponseBodyFileModel) GetUploadId() *string {
	return s.UploadId
}

func (s *CreateCdsFileResponseBodyFileModel) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *CreateCdsFileResponseBodyFileModel) SetFileId(v string) *CreateCdsFileResponseBodyFileModel {
	s.FileId = &v
	return s
}

func (s *CreateCdsFileResponseBodyFileModel) SetUploadId(v string) *CreateCdsFileResponseBodyFileModel {
	s.UploadId = &v
	return s
}

func (s *CreateCdsFileResponseBodyFileModel) SetUploadUrl(v string) *CreateCdsFileResponseBodyFileModel {
	s.UploadUrl = &v
	return s
}

func (s *CreateCdsFileResponseBodyFileModel) Validate() error {
	return dara.Validate(s)
}
