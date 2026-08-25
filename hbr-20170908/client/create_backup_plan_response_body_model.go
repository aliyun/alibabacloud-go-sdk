// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBackupPlanResponseBody
	GetCode() *string
	SetMessage(v string) *CreateBackupPlanResponseBody
	GetMessage() *string
	SetPlanId(v string) *CreateBackupPlanResponseBody
	GetPlanId() *string
	SetRequestId(v string) *CreateBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBackupPlanResponseBody
	GetSuccess() *bool
}

type CreateBackupPlanResponseBody struct {
	// The response code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. The value "successful" is returned for a success response. An error message is returned for a failure response.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The backup plan ID.
	//
	// example:
	//
	// plan-*********************
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBackupPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBackupPlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *CreateBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBackupPlanResponseBody) SetCode(v string) *CreateBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetMessage(v string) *CreateBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetPlanId(v string) *CreateBackupPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetRequestId(v string) *CreateBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetSuccess(v bool) *CreateBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
