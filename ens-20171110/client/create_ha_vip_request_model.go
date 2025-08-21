// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHaVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateHaVipRequest
	GetAmount() *int32
	SetDescription(v string) *CreateHaVipRequest
	GetDescription() *string
	SetIpAddress(v string) *CreateHaVipRequest
	GetIpAddress() *string
	SetName(v string) *CreateHaVipRequest
	GetName() *string
	SetVSwitchId(v string) *CreateHaVipRequest
	GetVSwitchId() *string
}

type CreateHaVipRequest struct {
	// The number of HAVIPs that you want to create. Valid values: 1 to 10. The value can be only 1 if you specify an IP address.
	//
	// Default value: 1.
	//
	// example:
	//
	// 6
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 120.24.243.91
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// example:
	//
	// yourName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// vsw-5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateHaVipRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHaVipRequest) GoString() string {
	return s.String()
}

func (s *CreateHaVipRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateHaVipRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHaVipRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *CreateHaVipRequest) GetName() *string {
	return s.Name
}

func (s *CreateHaVipRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateHaVipRequest) SetAmount(v int32) *CreateHaVipRequest {
	s.Amount = &v
	return s
}

func (s *CreateHaVipRequest) SetDescription(v string) *CreateHaVipRequest {
	s.Description = &v
	return s
}

func (s *CreateHaVipRequest) SetIpAddress(v string) *CreateHaVipRequest {
	s.IpAddress = &v
	return s
}

func (s *CreateHaVipRequest) SetName(v string) *CreateHaVipRequest {
	s.Name = &v
	return s
}

func (s *CreateHaVipRequest) SetVSwitchId(v string) *CreateHaVipRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateHaVipRequest) Validate() error {
	return dara.Validate(s)
}
