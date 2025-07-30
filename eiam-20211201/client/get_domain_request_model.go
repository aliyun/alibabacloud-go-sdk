// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *GetDomainRequest
	GetDomainId() *string
	SetInstanceId(v string) *GetDomainRequest
	GetInstanceId() *string
}

type GetDomainRequest struct {
	// Domain ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *GetDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDomainRequest) SetDomainId(v string) *GetDomainRequest {
	s.DomainId = &v
	return s
}

func (s *GetDomainRequest) SetInstanceId(v string) *GetDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDomainRequest) Validate() error {
	return dara.Validate(s)
}
