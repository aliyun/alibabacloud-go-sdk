// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadIdentityRegistrationDocConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerId(v string) *GetUploadIdentityRegistrationDocConfigRequest
	GetCustomerId() *string
	SetFilePath(v string) *GetUploadIdentityRegistrationDocConfigRequest
	GetFilePath() *string
}

type GetUploadIdentityRegistrationDocConfigRequest struct {
	// This parameter is required.
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// This parameter is required.
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s GetUploadIdentityRegistrationDocConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadIdentityRegistrationDocConfigRequest) GoString() string {
	return s.String()
}

func (s *GetUploadIdentityRegistrationDocConfigRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetUploadIdentityRegistrationDocConfigRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetUploadIdentityRegistrationDocConfigRequest) SetCustomerId(v string) *GetUploadIdentityRegistrationDocConfigRequest {
	s.CustomerId = &v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigRequest) SetFilePath(v string) *GetUploadIdentityRegistrationDocConfigRequest {
	s.FilePath = &v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigRequest) Validate() error {
	return dara.Validate(s)
}
