// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableHanaBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableHanaBackupPlanResponseBody
	GetCode() *string
	SetMessage(v string) *DisableHanaBackupPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableHanaBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableHanaBackupPlanResponseBody
	GetSuccess() *bool
}

type DisableHanaBackupPlanResponseBody struct {
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
	// FFC87EC8-8126-5967-9C4D-82715F8DFC97
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

func (s DisableHanaBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableHanaBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DisableHanaBackupPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableHanaBackupPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableHanaBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableHanaBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableHanaBackupPlanResponseBody) SetCode(v string) *DisableHanaBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DisableHanaBackupPlanResponseBody) SetMessage(v string) *DisableHanaBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DisableHanaBackupPlanResponseBody) SetRequestId(v string) *DisableHanaBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableHanaBackupPlanResponseBody) SetSuccess(v bool) *DisableHanaBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *DisableHanaBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
