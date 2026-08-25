// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateBackupPlanResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateBackupPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBackupPlanResponseBody
	GetSuccess() *bool
}

type UpdateBackupPlanResponseBody struct {
	// The return code. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned. If the request is successful, \\`successful\\` is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s UpdateBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateBackupPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBackupPlanResponseBody) SetCode(v string) *UpdateBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBackupPlanResponseBody) SetMessage(v string) *UpdateBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBackupPlanResponseBody) SetRequestId(v string) *UpdateBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBackupPlanResponseBody) SetSuccess(v bool) *UpdateBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
