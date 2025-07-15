// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVisitorLoginDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatDeviceId(v string) *GetVisitorLoginDetailsRequest
	GetChatDeviceId() *string
	SetInstanceId(v string) *GetVisitorLoginDetailsRequest
	GetInstanceId() *string
	SetToken(v string) *GetVisitorLoginDetailsRequest
	GetToken() *string
	SetVisitorId(v string) *GetVisitorLoginDetailsRequest
	GetVisitorId() *string
}

type GetVisitorLoginDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4c51c9116c36537cb850dc1081d745df
	ChatDeviceId *string `json:"ChatDeviceId,omitempty" xml:"ChatDeviceId,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 9XYGTGWtq2wFi_Bpg7aUnIoYi_vG_rO3b***YtsxbHRHrYHlz1LDBLJAyZcLxieRQR4h_6AnWvTjJeNU5jgxzO*****bHwej7WgWrmA
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fcd020fe-d8e4-40e5-8c77-1a272a174a7d
	VisitorId *string `json:"VisitorId,omitempty" xml:"VisitorId,omitempty"`
}

func (s GetVisitorLoginDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVisitorLoginDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetVisitorLoginDetailsRequest) GetChatDeviceId() *string {
	return s.ChatDeviceId
}

func (s *GetVisitorLoginDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVisitorLoginDetailsRequest) GetToken() *string {
	return s.Token
}

func (s *GetVisitorLoginDetailsRequest) GetVisitorId() *string {
	return s.VisitorId
}

func (s *GetVisitorLoginDetailsRequest) SetChatDeviceId(v string) *GetVisitorLoginDetailsRequest {
	s.ChatDeviceId = &v
	return s
}

func (s *GetVisitorLoginDetailsRequest) SetInstanceId(v string) *GetVisitorLoginDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVisitorLoginDetailsRequest) SetToken(v string) *GetVisitorLoginDetailsRequest {
	s.Token = &v
	return s
}

func (s *GetVisitorLoginDetailsRequest) SetVisitorId(v string) *GetVisitorLoginDetailsRequest {
	s.VisitorId = &v
	return s
}

func (s *GetVisitorLoginDetailsRequest) Validate() error {
	return dara.Validate(s)
}
