// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinationForCoordinatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoordinateFlowModels(v []*ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) *ApplyCoordinationForCoordinatorResponseBody
	GetCoordinateFlowModels() []*ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels
	SetRequestId(v string) *ApplyCoordinationForCoordinatorResponseBody
	GetRequestId() *string
}

type ApplyCoordinationForCoordinatorResponseBody struct {
	CoordinateFlowModels []*ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels `json:"CoordinateFlowModels,omitempty" xml:"CoordinateFlowModels,omitempty" type:"Repeated"`
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyCoordinationForCoordinatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForCoordinatorResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForCoordinatorResponseBody) GetCoordinateFlowModels() []*ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels {
	return s.CoordinateFlowModels
}

func (s *ApplyCoordinationForCoordinatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyCoordinationForCoordinatorResponseBody) SetCoordinateFlowModels(v []*ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) *ApplyCoordinationForCoordinatorResponseBody {
	s.CoordinateFlowModels = v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponseBody) SetRequestId(v string) *ApplyCoordinationForCoordinatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponseBody) Validate() error {
	return dara.Validate(s)
}

type ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels struct {
	// example:
	//
	// co-9kt75fon9pj****
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// example:
	//
	// PENDING
	CoordinateStatus *string `json:"CoordinateStatus,omitempty" xml:"CoordinateStatus,omitempty"`
	// example:
	//
	// W0Rlc2t0b3BdDQpHV1Rva2VuPTAwT1A1bHp1RUlEdU00T0IzemdiZ****
	CoordinateTicket *string `json:"CoordinateTicket,omitempty" xml:"CoordinateTicket,omitempty"`
	// example:
	//
	// 10419178654***
	CoordinatorUserId *string `json:"CoordinatorUserId,omitempty" xml:"CoordinatorUserId,omitempty"`
	// example:
	//
	// alice
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// ai-ij4a6kd4bn2****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// test-resource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) GetCoId() *string {
	return s.CoId
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) GetCoordinateStatus() *string {
	return s.CoordinateStatus
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) GetCoordinateTicket() *string {
	return s.CoordinateTicket
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) GetCoordinatorUserId() *string {
	return s.CoordinatorUserId
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) GetResourceId() *string {
	return s.ResourceId
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) GetResourceName() *string {
	return s.ResourceName
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) SetCoId(v string) *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels {
	s.CoId = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) SetCoordinateStatus(v string) *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels {
	s.CoordinateStatus = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) SetCoordinateTicket(v string) *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels {
	s.CoordinateTicket = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) SetCoordinatorUserId(v string) *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels {
	s.CoordinatorUserId = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) SetOwnerUserId(v string) *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels {
	s.OwnerUserId = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) SetResourceId(v string) *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels {
	s.ResourceId = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) SetResourceName(v string) *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels {
	s.ResourceName = &v
	return s
}

func (s *ApplyCoordinationForCoordinatorResponseBodyCoordinateFlowModels) Validate() error {
	return dara.Validate(s)
}
