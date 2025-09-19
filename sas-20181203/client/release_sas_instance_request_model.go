// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSasInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *ReleaseSasInstanceRequest
	GetChargeType() *string
	SetInstanceId(v string) *ReleaseSasInstanceRequest
	GetInstanceId() *string
}

type ReleaseSasInstanceRequest struct {
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseSasInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSasInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseSasInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ReleaseSasInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseSasInstanceRequest) SetChargeType(v string) *ReleaseSasInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *ReleaseSasInstanceRequest) SetInstanceId(v string) *ReleaseSasInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseSasInstanceRequest) Validate() error {
	return dara.Validate(s)
}
