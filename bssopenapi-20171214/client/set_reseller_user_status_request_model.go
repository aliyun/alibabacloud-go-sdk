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
	// The type of business operation. Valid values:
	//
	// FREEZE: Freeze operations.
	//
	// TRUSTEESHIP: Hosting operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// FREEZE
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// This parameter is required.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The status to set for the user account. Valid values:
	//
	// Freeze: The user account is frozen.
	//
	// Thaw: The user account is restored.
	//
	// Trusteeship: The user account is hosted.
	//
	// TrusteeshipCancel: The user account is not hosted.
	//
	// This parameter is required.
	//
	// example:
	//
	// Freeze
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The stopping mode. Valid values:
	//
	// 0: Normal shutdown.
	//
	// 1: Immediate shutdown.
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
