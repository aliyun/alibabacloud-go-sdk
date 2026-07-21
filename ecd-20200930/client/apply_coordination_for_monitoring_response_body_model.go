// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinationForMonitoringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoordinateFlowModels(v []*ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) *ApplyCoordinationForMonitoringResponseBody
	GetCoordinateFlowModels() []*ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels
	SetRequestId(v string) *ApplyCoordinationForMonitoringResponseBody
	GetRequestId() *string
}

type ApplyCoordinationForMonitoringResponseBody struct {
	// The list of coordination flow data.
	CoordinateFlowModels []*ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels `json:"CoordinateFlowModels,omitempty" xml:"CoordinateFlowModels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyCoordinationForMonitoringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForMonitoringResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForMonitoringResponseBody) GetCoordinateFlowModels() []*ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels {
	return s.CoordinateFlowModels
}

func (s *ApplyCoordinationForMonitoringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyCoordinationForMonitoringResponseBody) SetCoordinateFlowModels(v []*ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) *ApplyCoordinationForMonitoringResponseBody {
	s.CoordinateFlowModels = v
	return s
}

func (s *ApplyCoordinationForMonitoringResponseBody) SetRequestId(v string) *ApplyCoordinationForMonitoringResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCoordinationForMonitoringResponseBody) Validate() error {
	if s.CoordinateFlowModels != nil {
		for _, item := range s.CoordinateFlowModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels struct {
	// The coordination flow ID.
	//
	// example:
	//
	// co-0sot77uale3****
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// The current coordination status.
	//
	// [_single.resp.200.props.CoordinateFlowModels.items.CoordinateStatus.enum.COORDINATING  ]coordinating
	//
	// [_single.resp.200.props.CoordinateFlowModels.items.CoordinateStatus.enum.TERMINATING  ] terminating
	//
	// [_single.resp.200.props.CoordinateFlowModels.items.CoordinateStatus.enum.TERMINATED ]terminated
	//
	// [_single.resp.200.props.CoordinateFlowModels.items.CoordinateStatus.enum.PENDING ]pending acceptance
	//
	// example:
	//
	// PENDING
	CoordinateStatus *string `json:"CoordinateStatus,omitempty" xml:"CoordinateStatus,omitempty"`
	// The ticket used by ASP to establish a connection.
	//
	// example:
	//
	// 1VDQ0VTUw0KW0Rlc2t0b3BdDQpHV1Rva2VuPTAwTzgwL3liS25zUEVGdkF6eU1Pc1ExeHZWdmk4VEE3NFJvU1V1d0dPYm1BNkNJWklDMHVNQklWcjU2NS80S0ZQekQ4aGFTR0ZHelZqMTFGbkRpWWgvUFF1Zm1xSXNGdFRFNFRWMExJNit3TkU0L2RMb04wNXBBSE5Tc3M4dWFXY3lwWE****
	CoordinateTicket *string `json:"CoordinateTicket,omitempty" xml:"CoordinateTicket,omitempty"`
	// The initiator type.
	//
	// example:
	//
	// COORDINATOR_INITIATE_FORCE
	InitiatorType *string `json:"InitiatorType,omitempty" xml:"InitiatorType,omitempty"`
	// The Alibaba Cloud account ID of the user on the user side.
	//
	// example:
	//
	// alice
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-96vi03f9emqnl****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The cloud computer name.
	//
	// example:
	//
	// DemoComputer
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) GetCoId() *string {
	return s.CoId
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) GetCoordinateStatus() *string {
	return s.CoordinateStatus
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) GetCoordinateTicket() *string {
	return s.CoordinateTicket
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) GetInitiatorType() *string {
	return s.InitiatorType
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) GetResourceId() *string {
	return s.ResourceId
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) GetResourceName() *string {
	return s.ResourceName
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) SetCoId(v string) *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels {
	s.CoId = &v
	return s
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) SetCoordinateStatus(v string) *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels {
	s.CoordinateStatus = &v
	return s
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) SetCoordinateTicket(v string) *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels {
	s.CoordinateTicket = &v
	return s
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) SetInitiatorType(v string) *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels {
	s.InitiatorType = &v
	return s
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) SetOwnerUserId(v string) *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels {
	s.OwnerUserId = &v
	return s
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) SetResourceId(v string) *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels {
	s.ResourceId = &v
	return s
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) SetResourceName(v string) *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels {
	s.ResourceName = &v
	return s
}

func (s *ApplyCoordinationForMonitoringResponseBodyCoordinateFlowModels) Validate() error {
	return dara.Validate(s)
}
