// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteCertificateRequest
	GetId() *int64
	SetProjectId(v int64) *DeleteCertificateRequest
	GetProjectId() *int64
}

type DeleteCertificateRequest struct {
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
	// example:
	//
	// 106560
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteCertificateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteCertificateRequest) SetId(v int64) *DeleteCertificateRequest {
	s.Id = &v
	return s
}

func (s *DeleteCertificateRequest) SetProjectId(v int64) *DeleteCertificateRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteCertificateRequest) Validate() error {
	return dara.Validate(s)
}
