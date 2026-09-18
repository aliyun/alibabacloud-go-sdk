// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeBundleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiMetadata(v string) *CreateCodeBundleRequest
	GetCiMetadata() *string
	SetFilename(v string) *CreateCodeBundleRequest
	GetFilename() *string
}

type CreateCodeBundleRequest struct {
	// An optional CI/CD metadata JSON string.
	//
	// example:
	//
	// {"region":"cn-beijing"}
	CiMetadata *string `json:"ciMetadata,omitempty" xml:"ciMetadata,omitempty"`
	// The file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-cases.zip
	Filename *string `json:"filename,omitempty" xml:"filename,omitempty"`
}

func (s CreateCodeBundleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeBundleRequest) GoString() string {
	return s.String()
}

func (s *CreateCodeBundleRequest) GetCiMetadata() *string {
	return s.CiMetadata
}

func (s *CreateCodeBundleRequest) GetFilename() *string {
	return s.Filename
}

func (s *CreateCodeBundleRequest) SetCiMetadata(v string) *CreateCodeBundleRequest {
	s.CiMetadata = &v
	return s
}

func (s *CreateCodeBundleRequest) SetFilename(v string) *CreateCodeBundleRequest {
	s.Filename = &v
	return s
}

func (s *CreateCodeBundleRequest) Validate() error {
	return dara.Validate(s)
}
