// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetUploadCredentialsRequest
	GetFileName() *string
	SetVisibility(v string) *GetUploadCredentialsRequest
	GetVisibility() *string
}

type GetUploadCredentialsRequest struct {
	// The file name to upload.
	//
	// This parameter is required.
	//
	// example:
	//
	// template.yaml
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Specifies whether the file is publicly accessible. Valid values: **public*	- or **private**. The default value is **private**.
	//
	// example:
	//
	// public
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s GetUploadCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetUploadCredentialsRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *GetUploadCredentialsRequest) SetFileName(v string) *GetUploadCredentialsRequest {
	s.FileName = &v
	return s
}

func (s *GetUploadCredentialsRequest) SetVisibility(v string) *GetUploadCredentialsRequest {
	s.Visibility = &v
	return s
}

func (s *GetUploadCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
