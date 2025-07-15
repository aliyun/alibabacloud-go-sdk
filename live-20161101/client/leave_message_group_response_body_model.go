// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LeaveMessageGroupResponseBody
	GetRequestId() *string
	SetResult(v *LeaveMessageGroupResponseBodyResult) *LeaveMessageGroupResponseBody
	GetResult() *LeaveMessageGroupResponseBodyResult
}

type LeaveMessageGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *LeaveMessageGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s LeaveMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LeaveMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LeaveMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LeaveMessageGroupResponseBody) GetResult() *LeaveMessageGroupResponseBodyResult {
	return s.Result
}

func (s *LeaveMessageGroupResponseBody) SetRequestId(v string) *LeaveMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *LeaveMessageGroupResponseBody) SetResult(v *LeaveMessageGroupResponseBodyResult) *LeaveMessageGroupResponseBody {
	s.Result = v
	return s
}

func (s *LeaveMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type LeaveMessageGroupResponseBodyResult struct {
	// Indicates whether the user left the message group. Valid values:
	//
	// 	- true: The user left the message group.
	//
	// 	- false: The user failed to leave the message group.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LeaveMessageGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s LeaveMessageGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *LeaveMessageGroupResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *LeaveMessageGroupResponseBodyResult) SetSuccess(v bool) *LeaveMessageGroupResponseBodyResult {
	s.Success = &v
	return s
}

func (s *LeaveMessageGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
