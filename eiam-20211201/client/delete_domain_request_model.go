// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *DeleteDomainRequest
	GetDomainId() *string
	SetInstanceId(v string) *DeleteDomainRequest
	GetInstanceId() *string
}

type DeleteDomainRequest struct {
	// The ID of the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxx
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

func (s DeleteDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *DeleteDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDomainRequest) SetDomainId(v string) *DeleteDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteDomainRequest) SetInstanceId(v string) *DeleteDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDomainRequest) Validate() error {
	return dara.Validate(s)
}
