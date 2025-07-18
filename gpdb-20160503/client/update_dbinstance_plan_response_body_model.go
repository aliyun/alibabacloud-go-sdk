// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstancePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpdateDBInstancePlanResponseBody
	GetDBInstanceId() *string
	SetErrorMessage(v string) *UpdateDBInstancePlanResponseBody
	GetErrorMessage() *string
	SetPlanId(v string) *UpdateDBInstancePlanResponseBody
	GetPlanId() *string
	SetRequestId(v string) *UpdateDBInstancePlanResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateDBInstancePlanResponseBody
	GetStatus() *string
}

type UpdateDBInstancePlanResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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
	// 34b32a0a-08ef-4a87-b6be-cdd9f56fc3ad
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

func (s UpdateDBInstancePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstancePlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstancePlanResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpdateDBInstancePlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDBInstancePlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *UpdateDBInstancePlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDBInstancePlanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateDBInstancePlanResponseBody) SetDBInstanceId(v string) *UpdateDBInstancePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetErrorMessage(v string) *UpdateDBInstancePlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetPlanId(v string) *UpdateDBInstancePlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetRequestId(v string) *UpdateDBInstancePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetStatus(v string) *UpdateDBInstancePlanResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) Validate() error {
	return dara.Validate(s)
}
