// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallCenterProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestination(v string) *CreateCallCenterProviderRequest
	GetDestination() *string
	SetDisplayName(v string) *CreateCallCenterProviderRequest
	GetDisplayName() *string
	SetExtras(v string) *CreateCallCenterProviderRequest
	GetExtras() *string
	SetInstanceId(v string) *CreateCallCenterProviderRequest
	GetInstanceId() *string
	SetName(v string) *CreateCallCenterProviderRequest
	GetName() *string
	SetOriginator(v string) *CreateCallCenterProviderRequest
	GetOriginator() *string
	SetProviderType(v string) *CreateCallCenterProviderRequest
	GetProviderType() *string
	SetReferTo(v string) *CreateCallCenterProviderRequest
	GetReferTo() *string
	SetTrunkId(v string) *CreateCallCenterProviderRequest
	GetTrunkId() *string
}

type CreateCallCenterProviderRequest struct {
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

func (s CreateCallCenterProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCallCenterProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateCallCenterProviderRequest) GetDestination() *string {
	return s.Destination
}

func (s *CreateCallCenterProviderRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateCallCenterProviderRequest) GetExtras() *string {
	return s.Extras
}

func (s *CreateCallCenterProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCallCenterProviderRequest) GetName() *string {
	return s.Name
}

func (s *CreateCallCenterProviderRequest) GetOriginator() *string {
	return s.Originator
}

func (s *CreateCallCenterProviderRequest) GetProviderType() *string {
	return s.ProviderType
}

func (s *CreateCallCenterProviderRequest) GetReferTo() *string {
	return s.ReferTo
}

func (s *CreateCallCenterProviderRequest) GetTrunkId() *string {
	return s.TrunkId
}

func (s *CreateCallCenterProviderRequest) SetDestination(v string) *CreateCallCenterProviderRequest {
	s.Destination = &v
	return s
}

func (s *CreateCallCenterProviderRequest) SetDisplayName(v string) *CreateCallCenterProviderRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateCallCenterProviderRequest) SetExtras(v string) *CreateCallCenterProviderRequest {
	s.Extras = &v
	return s
}

func (s *CreateCallCenterProviderRequest) SetInstanceId(v string) *CreateCallCenterProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCallCenterProviderRequest) SetName(v string) *CreateCallCenterProviderRequest {
	s.Name = &v
	return s
}

func (s *CreateCallCenterProviderRequest) SetOriginator(v string) *CreateCallCenterProviderRequest {
	s.Originator = &v
	return s
}

func (s *CreateCallCenterProviderRequest) SetProviderType(v string) *CreateCallCenterProviderRequest {
	s.ProviderType = &v
	return s
}

func (s *CreateCallCenterProviderRequest) SetReferTo(v string) *CreateCallCenterProviderRequest {
	s.ReferTo = &v
	return s
}

func (s *CreateCallCenterProviderRequest) SetTrunkId(v string) *CreateCallCenterProviderRequest {
	s.TrunkId = &v
	return s
}

func (s *CreateCallCenterProviderRequest) Validate() error {
	return dara.Validate(s)
}
