// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableBackupPlanResponseBody
	GetCode() *string
	SetMessage(v string) *DisableBackupPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableBackupPlanResponseBody
	GetSuccess() *bool
}

type DisableBackupPlanResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
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

func (s DisableBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DisableBackupPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableBackupPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableBackupPlanResponseBody) SetCode(v string) *DisableBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DisableBackupPlanResponseBody) SetMessage(v string) *DisableBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DisableBackupPlanResponseBody) SetRequestId(v string) *DisableBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableBackupPlanResponseBody) SetSuccess(v bool) *DisableBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *DisableBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
