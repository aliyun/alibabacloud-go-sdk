// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *RemoveOrgRequest
	GetBusinessChannel() *string
	SetOrgId(v string) *RemoveOrgRequest
	GetOrgId() *string
}

type RemoveOrgRequest struct {
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The organization ID.
	//
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

func (s *RemoveOrgRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *RemoveOrgRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *RemoveOrgRequest) SetBusinessChannel(v string) *RemoveOrgRequest {
	s.BusinessChannel = &v
	return s
}

func (s *RemoveOrgRequest) SetOrgId(v string) *RemoveOrgRequest {
	s.OrgId = &v
	return s
}

func (s *RemoveOrgRequest) Validate() error {
	return dara.Validate(s)
}
