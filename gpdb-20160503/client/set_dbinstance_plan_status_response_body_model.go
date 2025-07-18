// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDBInstancePlanStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *SetDBInstancePlanStatusResponseBody
	GetErrorMessage() *string
	SetPlanId(v string) *SetDBInstancePlanStatusResponseBody
	GetPlanId() *string
	SetRequestId(v string) *SetDBInstancePlanStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetDBInstancePlanStatusResponseBody
	GetStatus() *string
}

type SetDBInstancePlanStatusResponseBody struct {
	// The error message returned.
	//
	// This parameter is returned only when the operation fails.
	//
	// example:
	//
	// ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the plan.
	//
	// example:
	//
	// 1234
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34b32a0a-08ef-4a87-b6be-cdd988888888
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation.
	//
	// If the operation is successful, **success*	- is returned. If the operation fails, this parameter is not returned.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDBInstancePlanStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDBInstancePlanStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDBInstancePlanStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SetDBInstancePlanStatusResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *SetDBInstancePlanStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDBInstancePlanStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetDBInstancePlanStatusResponseBody) SetErrorMessage(v string) *SetDBInstancePlanStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) SetPlanId(v string) *SetDBInstancePlanStatusResponseBody {
	s.PlanId = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) SetRequestId(v string) *SetDBInstancePlanStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) SetStatus(v string) *SetDBInstancePlanStatusResponseBody {
	s.Status = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
