// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHanaBackupPlanResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateHanaBackupPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHanaBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHanaBackupPlanResponseBody
	GetSuccess() *bool
}

type UpdateHanaBackupPlanResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
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
	// The request ID.
	//
	// example:
	//
	// F23BCC67-09B4-582C-AE70-C813C8548DCC
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

func (s UpdateHanaBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHanaBackupPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHanaBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHanaBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHanaBackupPlanResponseBody) SetCode(v string) *UpdateHanaBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHanaBackupPlanResponseBody) SetMessage(v string) *UpdateHanaBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHanaBackupPlanResponseBody) SetRequestId(v string) *UpdateHanaBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHanaBackupPlanResponseBody) SetSuccess(v bool) *UpdateHanaBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHanaBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
