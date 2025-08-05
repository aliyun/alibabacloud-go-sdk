// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectDefaultQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v string) *UpdateProjectDefaultQuotaRequest
	GetQuota() *string
}

type UpdateProjectDefaultQuotaRequest struct {
	// The default computing quota that is used to allocate computing resources, the jobs that are initiated by this project consume the computing resources in the default quota.
	//
	// example:
	//
	// os_PayAsYouGoQuota
	Quota *string `json:"quota,omitempty" xml:"quota,omitempty"`
}

func (s UpdateProjectDefaultQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectDefaultQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectDefaultQuotaRequest) GetQuota() *string {
	return s.Quota
}

func (s *UpdateProjectDefaultQuotaRequest) SetQuota(v string) *UpdateProjectDefaultQuotaRequest {
	s.Quota = &v
	return s
}

func (s *UpdateProjectDefaultQuotaRequest) Validate() error {
	return dara.Validate(s)
}
