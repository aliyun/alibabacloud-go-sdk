// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPreSignedUrlForPutObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetPreSignedUrlForPutObjectRequest
	GetFileName() *string
}

type GetPreSignedUrlForPutObjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetPreSignedUrlForPutObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPreSignedUrlForPutObjectRequest) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlForPutObjectRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetPreSignedUrlForPutObjectRequest) SetFileName(v string) *GetPreSignedUrlForPutObjectRequest {
	s.FileName = &v
	return s
}

func (s *GetPreSignedUrlForPutObjectRequest) Validate() error {
	return dara.Validate(s)
}
