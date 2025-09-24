// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResellerUserStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *SetResellerUserStatusRequest
	GetBusinessType() *string
	SetOwnerId(v string) *SetResellerUserStatusRequest
	GetOwnerId() *string
	SetStatus(v string) *SetResellerUserStatusRequest
	GetStatus() *string
	SetStopMode(v string) *SetResellerUserStatusRequest
	GetStopMode() *string
}

type SetResellerUserStatusRequest struct {
	// The type of the business. Valid values: FREEZE: the frozen business of the account. TRUSTEESHIP: the hosted business of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// FREEZE
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// This parameter is required.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The account status that you want to set. Valid values: Freeze: The account is frozen. Thaw: The account is unfrozen. Trusteeship: The account is hosted. TrusteeshipCancel: The account is not hosted.
	//
	// This parameter is required.
	//
	// example:
	//
	// Freeze
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 停机模式
	//
	// 取值：
	//
	//     0：普通停机
	//
	//     1：立即停机
	//
	// example:
	//
	// 0
	StopMode *string `json:"StopMode,omitempty" xml:"StopMode,omitempty"`
}

func (s SetResellerUserStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetResellerUserStatusRequest) GoString() string {
	return s.String()
}

func (s *SetResellerUserStatusRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *SetResellerUserStatusRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetResellerUserStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetResellerUserStatusRequest) GetStopMode() *string {
	return s.StopMode
}

func (s *SetResellerUserStatusRequest) SetBusinessType(v string) *SetResellerUserStatusRequest {
	s.BusinessType = &v
	return s
}

func (s *SetResellerUserStatusRequest) SetOwnerId(v string) *SetResellerUserStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetResellerUserStatusRequest) SetStatus(v string) *SetResellerUserStatusRequest {
	s.Status = &v
	return s
}

func (s *SetResellerUserStatusRequest) SetStopMode(v string) *SetResellerUserStatusRequest {
	s.StopMode = &v
	return s
}

func (s *SetResellerUserStatusRequest) Validate() error {
	return dara.Validate(s)
}
