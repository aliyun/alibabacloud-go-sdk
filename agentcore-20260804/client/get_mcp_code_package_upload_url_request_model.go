// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpCodePackageUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *GetMcpCodePackageUploadUrlRequestBody) *GetMcpCodePackageUploadUrlRequest
	GetBody() *GetMcpCodePackageUploadUrlRequestBody
}

type GetMcpCodePackageUploadUrlRequest struct {
	// The request body.
	Body *GetMcpCodePackageUploadUrlRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s GetMcpCodePackageUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpCodePackageUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMcpCodePackageUploadUrlRequest) GetBody() *GetMcpCodePackageUploadUrlRequestBody {
	return s.Body
}

func (s *GetMcpCodePackageUploadUrlRequest) SetBody(v *GetMcpCodePackageUploadUrlRequestBody) *GetMcpCodePackageUploadUrlRequest {
	s.Body = v
	return s
}

func (s *GetMcpCodePackageUploadUrlRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpCodePackageUploadUrlRequestBody struct {
	// The file name of the ZIP code package.
	//
	// This parameter is required.
	//
	// example:
	//
	// server.zip
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetMcpCodePackageUploadUrlRequestBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpCodePackageUploadUrlRequestBody) GoString() string {
	return s.String()
}

func (s *GetMcpCodePackageUploadUrlRequestBody) GetFileName() *string {
	return s.FileName
}

func (s *GetMcpCodePackageUploadUrlRequestBody) SetFileName(v string) *GetMcpCodePackageUploadUrlRequestBody {
	s.FileName = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlRequestBody) Validate() error {
	return dara.Validate(s)
}
