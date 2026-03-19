// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduleExecuteTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyScheduleExecuteTimeResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyScheduleExecuteTimeResponseBody
	GetResult() *bool
}

type ModifyScheduleExecuteTimeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyScheduleExecuteTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduleExecuteTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduleExecuteTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScheduleExecuteTimeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyScheduleExecuteTimeResponseBody) SetRequestId(v string) *ModifyScheduleExecuteTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScheduleExecuteTimeResponseBody) SetResult(v bool) *ModifyScheduleExecuteTimeResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyScheduleExecuteTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
