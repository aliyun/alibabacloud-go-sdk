// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFingerprint(v string) *ActivateLicenseRequest
	GetFingerprint() *string
	SetId(v string) *ActivateLicenseRequest
	GetId() *string
	SetInstanceId(v string) *ActivateLicenseRequest
	GetInstanceId() *string
	SetOrderId(v string) *ActivateLicenseRequest
	GetOrderId() *string
}

type ActivateLicenseRequest struct {
	// example:
	//
	// XXX
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 12bea6b4489fsde7b903fe05934a0adx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ActivateLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateLicenseRequest) GoString() string {
	return s.String()
}

func (s *ActivateLicenseRequest) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *ActivateLicenseRequest) GetId() *string {
	return s.Id
}

func (s *ActivateLicenseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ActivateLicenseRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *ActivateLicenseRequest) SetFingerprint(v string) *ActivateLicenseRequest {
	s.Fingerprint = &v
	return s
}

func (s *ActivateLicenseRequest) SetId(v string) *ActivateLicenseRequest {
	s.Id = &v
	return s
}

func (s *ActivateLicenseRequest) SetInstanceId(v string) *ActivateLicenseRequest {
	s.InstanceId = &v
	return s
}

func (s *ActivateLicenseRequest) SetOrderId(v string) *ActivateLicenseRequest {
	s.OrderId = &v
	return s
}

func (s *ActivateLicenseRequest) Validate() error {
	return dara.Validate(s)
}
