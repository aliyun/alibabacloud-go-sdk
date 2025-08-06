// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicenseKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v int64) *GetLicenseKeyRequest
	GetUserId() *int64
}

type GetLicenseKeyRequest struct {
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetLicenseKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseKeyRequest) GoString() string {
	return s.String()
}

func (s *GetLicenseKeyRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *GetLicenseKeyRequest) SetUserId(v int64) *GetLicenseKeyRequest {
	s.UserId = &v
	return s
}

func (s *GetLicenseKeyRequest) Validate() error {
	return dara.Validate(s)
}
