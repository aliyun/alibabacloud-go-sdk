// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DisableJobResponseBody
	GetCode() *int32
	SetMessage(v string) *DisableJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableJobResponseBody
	GetSuccess() *bool
}

type DisableJobResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// jobid: 92583 not match groupId: testSchedulerx.defaultGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C8E5FB4A-6D8D-424D-9AAA-4FE06BB74FF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the job was disabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableJobResponseBody) GoString() string {
	return s.String()
}

func (s *DisableJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DisableJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableJobResponseBody) SetCode(v int32) *DisableJobResponseBody {
	s.Code = &v
	return s
}

func (s *DisableJobResponseBody) SetMessage(v string) *DisableJobResponseBody {
	s.Message = &v
	return s
}

func (s *DisableJobResponseBody) SetRequestId(v string) *DisableJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableJobResponseBody) SetSuccess(v bool) *DisableJobResponseBody {
	s.Success = &v
	return s
}

func (s *DisableJobResponseBody) Validate() error {
	return dara.Validate(s)
}
