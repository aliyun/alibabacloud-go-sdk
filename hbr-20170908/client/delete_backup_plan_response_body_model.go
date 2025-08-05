// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBackupPlanResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteBackupPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBackupPlanResponseBody
	GetSuccess() *bool
}

type DeleteBackupPlanResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBackupPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBackupPlanResponseBody) SetCode(v string) *DeleteBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBackupPlanResponseBody) SetMessage(v string) *DeleteBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBackupPlanResponseBody) SetRequestId(v string) *DeleteBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupPlanResponseBody) SetSuccess(v bool) *DeleteBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
