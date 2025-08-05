// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHanaBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHanaBackupPlanResponseBody
	GetCode() *string
	SetMessage(v string) *CreateHanaBackupPlanResponseBody
	GetMessage() *string
	SetPlanId(v string) *CreateHanaBackupPlanResponseBody
	GetPlanId() *string
	SetRequestId(v string) *CreateHanaBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHanaBackupPlanResponseBody
	GetSuccess() *bool
}

type CreateHanaBackupPlanResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the backup plan.
	//
	// example:
	//
	// pl-000756jdlk2zmqig2nea
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 33AA3AAE-89E1-5D3A-A51D-0C0A80850F68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHanaBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHanaBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHanaBackupPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHanaBackupPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHanaBackupPlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *CreateHanaBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHanaBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHanaBackupPlanResponseBody) SetCode(v string) *CreateHanaBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHanaBackupPlanResponseBody) SetMessage(v string) *CreateHanaBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHanaBackupPlanResponseBody) SetPlanId(v string) *CreateHanaBackupPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateHanaBackupPlanResponseBody) SetRequestId(v string) *CreateHanaBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHanaBackupPlanResponseBody) SetSuccess(v bool) *CreateHanaBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHanaBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
