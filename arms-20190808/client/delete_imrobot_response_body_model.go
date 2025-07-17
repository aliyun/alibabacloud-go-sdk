// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIMRobotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteIMRobotResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteIMRobotResponseBody
	GetRequestId() *string
}

type DeleteIMRobotResponseBody struct {
	// Indicates whether the call was successful.
	//
	// 	- `true`: The call was successful.
	//
	// 	- `false`: The call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID. You can use the ID to find logs and troubleshoot issues.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIMRobotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIMRobotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIMRobotResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteIMRobotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIMRobotResponseBody) SetIsSuccess(v bool) *DeleteIMRobotResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteIMRobotResponseBody) SetRequestId(v string) *DeleteIMRobotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIMRobotResponseBody) Validate() error {
	return dara.Validate(s)
}
