// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCallCenterProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestination(v string) *UpdateCallCenterProviderRequest
	GetDestination() *string
	SetDisplayName(v string) *UpdateCallCenterProviderRequest
	GetDisplayName() *string
	SetExtras(v string) *UpdateCallCenterProviderRequest
	GetExtras() *string
	SetInstanceId(v string) *UpdateCallCenterProviderRequest
	GetInstanceId() *string
	SetName(v string) *UpdateCallCenterProviderRequest
	GetName() *string
	SetOriginator(v string) *UpdateCallCenterProviderRequest
	GetOriginator() *string
	SetProviderId(v string) *UpdateCallCenterProviderRequest
	GetProviderId() *string
	SetProviderType(v string) *UpdateCallCenterProviderRequest
	GetProviderType() *string
	SetReferTo(v string) *UpdateCallCenterProviderRequest
	GetReferTo() *string
	SetTrunkId(v string) *UpdateCallCenterProviderRequest
	GetTrunkId() *string
}

type UpdateCallCenterProviderRequest struct {
	// example:
	//
	// 153*********
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// key1=value1
	Extras *string `json:"Extras,omitempty" xml:"Extras,omitempty"`
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 010****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// example:
	//
	// xxxxxxxxx
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// CCC
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// example:
	//
	// -
	ReferTo *string `json:"ReferTo,omitempty" xml:"ReferTo,omitempty"`
	// example:
	//
	// trunk-xxx
	TrunkId *string `json:"TrunkId,omitempty" xml:"TrunkId,omitempty"`
}

func (s UpdateCallCenterProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCallCenterProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateCallCenterProviderRequest) GetDestination() *string {
	return s.Destination
}

func (s *UpdateCallCenterProviderRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateCallCenterProviderRequest) GetExtras() *string {
	return s.Extras
}

func (s *UpdateCallCenterProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCallCenterProviderRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCallCenterProviderRequest) GetOriginator() *string {
	return s.Originator
}

func (s *UpdateCallCenterProviderRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *UpdateCallCenterProviderRequest) GetProviderType() *string {
	return s.ProviderType
}

func (s *UpdateCallCenterProviderRequest) GetReferTo() *string {
	return s.ReferTo
}

func (s *UpdateCallCenterProviderRequest) GetTrunkId() *string {
	return s.TrunkId
}

func (s *UpdateCallCenterProviderRequest) SetDestination(v string) *UpdateCallCenterProviderRequest {
	s.Destination = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) SetDisplayName(v string) *UpdateCallCenterProviderRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) SetExtras(v string) *UpdateCallCenterProviderRequest {
	s.Extras = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) SetInstanceId(v string) *UpdateCallCenterProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) SetName(v string) *UpdateCallCenterProviderRequest {
	s.Name = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) SetOriginator(v string) *UpdateCallCenterProviderRequest {
	s.Originator = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) SetProviderId(v string) *UpdateCallCenterProviderRequest {
	s.ProviderId = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) SetProviderType(v string) *UpdateCallCenterProviderRequest {
	s.ProviderType = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) SetReferTo(v string) *UpdateCallCenterProviderRequest {
	s.ReferTo = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) SetTrunkId(v string) *UpdateCallCenterProviderRequest {
	s.TrunkId = &v
	return s
}

func (s *UpdateCallCenterProviderRequest) Validate() error {
	return dara.Validate(s)
}
