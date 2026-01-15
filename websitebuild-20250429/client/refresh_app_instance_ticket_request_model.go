// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAppInstanceTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *RefreshAppInstanceTicketRequest
	GetBizId() *string
	SetToken(v string) *RefreshAppInstanceTicketRequest
	GetToken() *string
	SetUuid(v string) *RefreshAppInstanceTicketRequest
	GetUuid() *string
}

type RefreshAppInstanceTicketRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// c5c354d7-5e68-443a-b7fc-767e6ede9deb
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RefreshAppInstanceTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshAppInstanceTicketRequest) GoString() string {
	return s.String()
}

func (s *RefreshAppInstanceTicketRequest) GetBizId() *string {
	return s.BizId
}

func (s *RefreshAppInstanceTicketRequest) GetToken() *string {
	return s.Token
}

func (s *RefreshAppInstanceTicketRequest) GetUuid() *string {
	return s.Uuid
}

func (s *RefreshAppInstanceTicketRequest) SetBizId(v string) *RefreshAppInstanceTicketRequest {
	s.BizId = &v
	return s
}

func (s *RefreshAppInstanceTicketRequest) SetToken(v string) *RefreshAppInstanceTicketRequest {
	s.Token = &v
	return s
}

func (s *RefreshAppInstanceTicketRequest) SetUuid(v string) *RefreshAppInstanceTicketRequest {
	s.Uuid = &v
	return s
}

func (s *RefreshAppInstanceTicketRequest) Validate() error {
	return dara.Validate(s)
}
