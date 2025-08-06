// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDRMCertInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *GetDRMCertInfoRequest
	GetCertId() *string
	SetVideoId(v string) *GetDRMCertInfoRequest
	GetVideoId() *string
}

type GetDRMCertInfoRequest struct {
	// This parameter is required.
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// This parameter is required.
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetDRMCertInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDRMCertInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDRMCertInfoRequest) GetCertId() *string {
	return s.CertId
}

func (s *GetDRMCertInfoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetDRMCertInfoRequest) SetCertId(v string) *GetDRMCertInfoRequest {
	s.CertId = &v
	return s
}

func (s *GetDRMCertInfoRequest) SetVideoId(v string) *GetDRMCertInfoRequest {
	s.VideoId = &v
	return s
}

func (s *GetDRMCertInfoRequest) Validate() error {
	return dara.Validate(s)
}
