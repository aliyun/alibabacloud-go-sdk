// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgId(v string) *RemoveOrgRequest
	GetOrgId() *string
}

type RemoveOrgRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// org-5yy5icj981xe5****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s RemoveOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveOrgRequest) GoString() string {
	return s.String()
}

func (s *RemoveOrgRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *RemoveOrgRequest) SetOrgId(v string) *RemoveOrgRequest {
	s.OrgId = &v
	return s
}

func (s *RemoveOrgRequest) Validate() error {
	return dara.Validate(s)
}
