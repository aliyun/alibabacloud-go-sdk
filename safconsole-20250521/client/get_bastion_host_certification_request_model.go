// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBastionHostCertificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *GetBastionHostCertificationRequest
	GetProjectId() *int64
}

type GetBastionHostCertificationRequest struct {
	// Project ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetBastionHostCertificationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBastionHostCertificationRequest) GoString() string {
	return s.String()
}

func (s *GetBastionHostCertificationRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBastionHostCertificationRequest) SetProjectId(v int64) *GetBastionHostCertificationRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBastionHostCertificationRequest) Validate() error {
	return dara.Validate(s)
}
