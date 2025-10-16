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
	// The list of stream collaboration models.
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
	// The ID of the stream collaboration.
	//
	// example:
	//
	// co-0sot77uale3****
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// The current status of the collaboration task.
	//
	// Valid values:
	//
	// 	- COORDINATING: The collaboration task is being executed.
	//
	// 	- TERMINATING: The collaboration task is being terminated.
	//
	// 	- TERMINATED: The collaboration task is terminated.
	//
	// 	- PENDING: The collaboration task is pending to be executed.
	//
	// example:
	//
	// PENDING
	CoordinateStatus *string `json:"CoordinateStatus,omitempty" xml:"CoordinateStatus,omitempty"`
	// The ticket that is used to establish the Adaptive Streaming Protocol (ASP)-based connection.
	//
	// example:
	//
	// 1VDQ0VTUw0KW0Rlc2t0b3BdDQpHV1Rva2VuPTAwTzgwL3liS25zUEVGdkF6eU1Pc1ExeHZWdmk4VEE3NFJvU1V1d0dPYm1BNkNJWklDMHVNQklWcjU2NS80S0ZQekQ4aGFTR0ZHelZqMTFGbkRpWWgvUFF1Zm1xSXNGdFRFNFRWMExJNit3TkU0L2RMb04wNXBBSE5Tc3M4dWFXY3lwWE****
	CoordinateTicket *string `json:"CoordinateTicket,omitempty" xml:"CoordinateTicket,omitempty"`
	// The type of the initiator.
	//
	// Valid values:
	//
	// 	- ADMIN_INITIATE_FORCE: The administrator forcibly initiates the collaboration request.
	//
	// 	- ADMIN_INITIATE: The administrator initiates the collaboration request.
	//
	// 	- COORDINATOR_INITIATE_FORCE: The coordinator forcibly initiates the collaboration request.
	//
	// example:
	//
	// COORDINATOR_INITIATE_FORCE
	InitiatorType *string `json:"InitiatorType,omitempty" xml:"InitiatorType,omitempty"`
	// The ID of the Alibaba Cloud account of the end user.
	//
	// example:
	//
	// alice
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The ID of the cloud desktop.
	//
	// example:
	//
	// ecd-96vi03f9emqnl****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the cloud desktop.
	//
	// example:
	//
	// TestDesktop
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
