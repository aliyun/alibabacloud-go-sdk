// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateAppInstanceTicketRequest
	GetBizId() *string
	SetClientId(v string) *CreateAppInstanceTicketRequest
	GetClientId() *string
	SetRole(v string) *CreateAppInstanceTicketRequest
	GetRole() *string
}

type CreateAppInstanceTicketRequest struct {
	// The customer business ID.
	//
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The Client ID of the device for which you want to revoke the access credential.
	//
	// example:
	//
	// d566aaf2-7c88-40a4-982f-6abef0be13c9
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateAppInstanceTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceTicketRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateAppInstanceTicketRequest) GetClientId() *string {
	return s.ClientId
}

func (s *CreateAppInstanceTicketRequest) GetRole() *string {
	return s.Role
}

func (s *CreateAppInstanceTicketRequest) SetBizId(v string) *CreateAppInstanceTicketRequest {
	s.BizId = &v
	return s
}

func (s *CreateAppInstanceTicketRequest) SetClientId(v string) *CreateAppInstanceTicketRequest {
	s.ClientId = &v
	return s
}

func (s *CreateAppInstanceTicketRequest) SetRole(v string) *CreateAppInstanceTicketRequest {
	s.Role = &v
	return s
}

func (s *CreateAppInstanceTicketRequest) Validate() error {
	return dara.Validate(s)
}
