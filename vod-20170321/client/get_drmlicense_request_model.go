// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDRMLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCDMData(v string) *GetDRMLicenseRequest
	GetCDMData() *string
	SetCertId(v string) *GetDRMLicenseRequest
	GetCertId() *string
	SetDRMType(v string) *GetDRMLicenseRequest
	GetDRMType() *string
	SetVideoId(v string) *GetDRMLicenseRequest
	GetVideoId() *string
}

type GetDRMLicenseRequest struct {
	// This parameter is required.
	CDMData *string `json:"CDMData,omitempty" xml:"CDMData,omitempty"`
	CertId  *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	DRMType *string `json:"DRMType,omitempty" xml:"DRMType,omitempty"`
	// This parameter is required.
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetDRMLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDRMLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseRequest) GetCDMData() *string {
	return s.CDMData
}

func (s *GetDRMLicenseRequest) GetCertId() *string {
	return s.CertId
}

func (s *GetDRMLicenseRequest) GetDRMType() *string {
	return s.DRMType
}

func (s *GetDRMLicenseRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetDRMLicenseRequest) SetCDMData(v string) *GetDRMLicenseRequest {
	s.CDMData = &v
	return s
}

func (s *GetDRMLicenseRequest) SetCertId(v string) *GetDRMLicenseRequest {
	s.CertId = &v
	return s
}

func (s *GetDRMLicenseRequest) SetDRMType(v string) *GetDRMLicenseRequest {
	s.DRMType = &v
	return s
}

func (s *GetDRMLicenseRequest) SetVideoId(v string) *GetDRMLicenseRequest {
	s.VideoId = &v
	return s
}

func (s *GetDRMLicenseRequest) Validate() error {
	return dara.Validate(s)
}
