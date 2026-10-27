// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *GetAgentCodePackageUploadUrlRequestBody) *GetAgentCodePackageUploadUrlRequest
	GetBody() *GetAgentCodePackageUploadUrlRequestBody
}

type GetAgentCodePackageUploadUrlRequest struct {
	// The request body.
	//
	// This parameter is required.
	Body *GetAgentCodePackageUploadUrlRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s GetAgentCodePackageUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageUploadUrlRequest) GetBody() *GetAgentCodePackageUploadUrlRequestBody {
	return s.Body
}

func (s *GetAgentCodePackageUploadUrlRequest) SetBody(v *GetAgentCodePackageUploadUrlRequestBody) *GetAgentCodePackageUploadUrlRequest {
	s.Body = v
	return s
}

func (s *GetAgentCodePackageUploadUrlRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentCodePackageUploadUrlRequestBody struct {
	// The filename of the code package.
	//
	// This parameter is required.
	//
	// example:
	//
	// main.zip
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The declared compressed size of the code package.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The format of the code package.
	//
	// This parameter is required.
	//
	// example:
	//
	// zip
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
}

func (s GetAgentCodePackageUploadUrlRequestBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageUploadUrlRequestBody) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageUploadUrlRequestBody) GetFileName() *string {
	return s.FileName
}

func (s *GetAgentCodePackageUploadUrlRequestBody) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetAgentCodePackageUploadUrlRequestBody) GetFormat() *string {
	return s.Format
}

func (s *GetAgentCodePackageUploadUrlRequestBody) SetFileName(v string) *GetAgentCodePackageUploadUrlRequestBody {
	s.FileName = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlRequestBody) SetFileSize(v int64) *GetAgentCodePackageUploadUrlRequestBody {
	s.FileSize = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlRequestBody) SetFormat(v string) *GetAgentCodePackageUploadUrlRequestBody {
	s.Format = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlRequestBody) Validate() error {
	return dara.Validate(s)
}
