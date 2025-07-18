// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstancePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDBInstancePlanResponseBody
	GetDBInstanceId() *string
	SetErrorMessage(v string) *CreateDBInstancePlanResponseBody
	GetErrorMessage() *string
	SetPlanId(v string) *CreateDBInstancePlanResponseBody
	GetPlanId() *string
	SetRequestId(v string) *CreateDBInstancePlanResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateDBInstancePlanResponseBody
	GetStatus() *string
}

type CreateDBInstancePlanResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message.
	//
	// This parameter is returned only if the request fails.
	//
	// example:
	//
	// ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The plan ID.
	//
	// example:
	//
	// 1234
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34b32a0a-08ef-4a87-b6be-cdd9f56fc3ad
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// If the request was successful, **success*	- is returned. If the request failed, this parameter is not returned.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateDBInstancePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstancePlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstancePlanResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstancePlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDBInstancePlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *CreateDBInstancePlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstancePlanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateDBInstancePlanResponseBody) SetDBInstanceId(v string) *CreateDBInstancePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetErrorMessage(v string) *CreateDBInstancePlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetPlanId(v string) *CreateDBInstancePlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetRequestId(v string) *CreateDBInstancePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetStatus(v string) *CreateDBInstancePlanResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) Validate() error {
	return dara.Validate(s)
}
