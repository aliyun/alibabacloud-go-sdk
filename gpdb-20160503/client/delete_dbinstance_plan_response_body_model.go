// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstancePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBInstancePlanResponseBody
	GetDBInstanceId() *string
	SetErrorMessage(v string) *DeleteDBInstancePlanResponseBody
	GetErrorMessage() *string
	SetPlanId(v string) *DeleteDBInstancePlanResponseBody
	GetPlanId() *string
	SetRequestId(v string) *DeleteDBInstancePlanResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteDBInstancePlanResponseBody
	GetStatus() *string
}

type DeleteDBInstancePlanResponseBody struct {
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

func (s DeleteDBInstancePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstancePlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstancePlanResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstancePlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDBInstancePlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *DeleteDBInstancePlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBInstancePlanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteDBInstancePlanResponseBody) SetDBInstanceId(v string) *DeleteDBInstancePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetErrorMessage(v string) *DeleteDBInstancePlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetPlanId(v string) *DeleteDBInstancePlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetRequestId(v string) *DeleteDBInstancePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetStatus(v string) *DeleteDBInstancePlanResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) Validate() error {
	return dara.Validate(s)
}
