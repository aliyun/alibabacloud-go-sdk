// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchModifyEntitlementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntitlements(v *BatchModifyEntitlementResponseBodyEntitlements) *BatchModifyEntitlementResponseBody
	GetEntitlements() *BatchModifyEntitlementResponseBodyEntitlements
	SetRequestId(v string) *BatchModifyEntitlementResponseBody
	GetRequestId() *string
}

type BatchModifyEntitlementResponseBody struct {
	Entitlements *BatchModifyEntitlementResponseBodyEntitlements `json:"Entitlements,omitempty" xml:"Entitlements,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 51592A88-0F2C-55E6-AD2C-2AD9C10D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchModifyEntitlementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchModifyEntitlementResponseBody) GoString() string {
	return s.String()
}

func (s *BatchModifyEntitlementResponseBody) GetEntitlements() *BatchModifyEntitlementResponseBodyEntitlements {
	return s.Entitlements
}

func (s *BatchModifyEntitlementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchModifyEntitlementResponseBody) SetEntitlements(v *BatchModifyEntitlementResponseBodyEntitlements) *BatchModifyEntitlementResponseBody {
	s.Entitlements = v
	return s
}

func (s *BatchModifyEntitlementResponseBody) SetRequestId(v string) *BatchModifyEntitlementResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchModifyEntitlementResponseBody) Validate() error {
	if s.Entitlements != nil {
		if err := s.Entitlements.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchModifyEntitlementResponseBodyEntitlements struct {
	AssignModels []*BatchModifyEntitlementResponseBodyEntitlementsAssignModels `json:"AssignModels,omitempty" xml:"AssignModels,omitempty" type:"Repeated"`
	// The result.
	//
	// Valid values:
	//
	// 	- FAILED
	//
	// 	- NOT_STARTED
	//
	// 	- STARTED
	//
	// 	- PROCESSING
	//
	// 	- FINISHED
	//
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// B2F4F018-0EDF-159C-B285-117B5F1C****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchModifyEntitlementResponseBodyEntitlements) String() string {
	return dara.Prettify(s)
}

func (s BatchModifyEntitlementResponseBodyEntitlements) GoString() string {
	return s.String()
}

func (s *BatchModifyEntitlementResponseBodyEntitlements) GetAssignModels() []*BatchModifyEntitlementResponseBodyEntitlementsAssignModels {
	return s.AssignModels
}

func (s *BatchModifyEntitlementResponseBodyEntitlements) GetStatus() *string {
	return s.Status
}

func (s *BatchModifyEntitlementResponseBodyEntitlements) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchModifyEntitlementResponseBodyEntitlements) SetAssignModels(v []*BatchModifyEntitlementResponseBodyEntitlementsAssignModels) *BatchModifyEntitlementResponseBodyEntitlements {
	s.AssignModels = v
	return s
}

func (s *BatchModifyEntitlementResponseBodyEntitlements) SetStatus(v string) *BatchModifyEntitlementResponseBodyEntitlements {
	s.Status = &v
	return s
}

func (s *BatchModifyEntitlementResponseBodyEntitlements) SetTaskId(v string) *BatchModifyEntitlementResponseBodyEntitlements {
	s.TaskId = &v
	return s
}

func (s *BatchModifyEntitlementResponseBodyEntitlements) Validate() error {
	if s.AssignModels != nil {
		for _, item := range s.AssignModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchModifyEntitlementResponseBodyEntitlementsAssignModels struct {
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-e94kzikmpljjx99pl
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The authorized user IDs for the cloud computer.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The assign result for each cloud computer.
	//
	// Valid values:
	//
	// 	- FAILED
	//
	// 	- NOT_STARTED
	//
	// 	- STARTED
	//
	// 	- PROCESSING
	//
	// 	- FINISHED
	//
	// example:
	//
	// FINISHED
	InnerStatus *string `json:"InnerStatus,omitempty" xml:"InnerStatus,omitempty"`
}

func (s BatchModifyEntitlementResponseBodyEntitlementsAssignModels) String() string {
	return dara.Prettify(s)
}

func (s BatchModifyEntitlementResponseBodyEntitlementsAssignModels) GoString() string {
	return s.String()
}

func (s *BatchModifyEntitlementResponseBodyEntitlementsAssignModels) GetDesktopId() *string {
	return s.DesktopId
}

func (s *BatchModifyEntitlementResponseBodyEntitlementsAssignModels) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *BatchModifyEntitlementResponseBodyEntitlementsAssignModels) GetInnerStatus() *string {
	return s.InnerStatus
}

func (s *BatchModifyEntitlementResponseBodyEntitlementsAssignModels) SetDesktopId(v string) *BatchModifyEntitlementResponseBodyEntitlementsAssignModels {
	s.DesktopId = &v
	return s
}

func (s *BatchModifyEntitlementResponseBodyEntitlementsAssignModels) SetEndUserIds(v []*string) *BatchModifyEntitlementResponseBodyEntitlementsAssignModels {
	s.EndUserIds = v
	return s
}

func (s *BatchModifyEntitlementResponseBodyEntitlementsAssignModels) SetInnerStatus(v string) *BatchModifyEntitlementResponseBodyEntitlementsAssignModels {
	s.InnerStatus = &v
	return s
}

func (s *BatchModifyEntitlementResponseBodyEntitlementsAssignModels) Validate() error {
	return dara.Validate(s)
}
