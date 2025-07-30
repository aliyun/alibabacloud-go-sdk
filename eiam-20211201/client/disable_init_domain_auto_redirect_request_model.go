// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInitDomainAutoRedirectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableInitDomainAutoRedirectRequest
	GetInstanceId() *string
}

type DisableInitDomainAutoRedirectRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableInitDomainAutoRedirectRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableInitDomainAutoRedirectRequest) GoString() string {
	return s.String()
}

func (s *DisableInitDomainAutoRedirectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableInitDomainAutoRedirectRequest) SetInstanceId(v string) *DisableInitDomainAutoRedirectRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableInitDomainAutoRedirectRequest) Validate() error {
	return dara.Validate(s)
}
