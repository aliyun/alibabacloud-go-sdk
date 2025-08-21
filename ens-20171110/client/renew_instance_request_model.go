// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RenewInstanceRequest
	GetInstanceId() *string
	SetPeriod(v int64) *RenewInstanceRequest
	GetPeriod() *int64
}

type RenewInstanceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-56789acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The renewal duration of the subscription instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewInstanceRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v int64) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
