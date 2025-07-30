// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainDnsChallengeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *GetDomainDnsChallengeRequest
	GetDomain() *string
	SetInstanceId(v string) *GetDomainDnsChallengeRequest
	GetInstanceId() *string
}

type GetDomainDnsChallengeRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxx
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDomainDnsChallengeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDomainDnsChallengeRequest) GoString() string {
	return s.String()
}

func (s *GetDomainDnsChallengeRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetDomainDnsChallengeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDomainDnsChallengeRequest) SetDomain(v string) *GetDomainDnsChallengeRequest {
	s.Domain = &v
	return s
}

func (s *GetDomainDnsChallengeRequest) SetInstanceId(v string) *GetDomainDnsChallengeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDomainDnsChallengeRequest) Validate() error {
	return dara.Validate(s)
}
