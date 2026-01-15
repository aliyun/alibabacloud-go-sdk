// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetCertificateRequest
	GetId() *int64
	SetProjectId(v int64) *GetCertificateRequest
	GetProjectId() *int64
}

type GetCertificateRequest struct {
	// The ID of the certificate file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 676303114031776
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the workspace to which the certificate file belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1065601
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetCertificateRequest) GetId() *int64 {
	return s.Id
}

func (s *GetCertificateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCertificateRequest) SetId(v int64) *GetCertificateRequest {
	s.Id = &v
	return s
}

func (s *GetCertificateRequest) SetProjectId(v int64) *GetCertificateRequest {
	s.ProjectId = &v
	return s
}

func (s *GetCertificateRequest) Validate() error {
	return dara.Validate(s)
}
