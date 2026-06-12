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
	// The name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// template.yaml
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The visibility of the bucket to which the file is uploaded. Valid values: public and private. A value of **public*	- means the file is uploaded to a public bucket. A value of **private*	- means the file is uploaded to a private bucket that requires authorization for access.
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
