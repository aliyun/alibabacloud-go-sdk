// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *RenewInstanceRequest
	GetAutoRenew() *bool
	SetDuration(v int32) *RenewInstanceRequest
	GetDuration() *int32
}

type RenewInstanceRequest struct {
	// Specifies whether to enable monthly auto-renewal. The default value is false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  If you enable auto-renewal for an instance for which auto-renewal is enabled, an error is reported.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The renewal duration. Unit: month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *RenewInstanceRequest) SetAutoRenew(v bool) *RenewInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
