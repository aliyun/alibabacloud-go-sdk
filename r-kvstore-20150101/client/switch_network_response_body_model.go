// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchNetworkResponseBody
	GetRequestId() *string
	SetTaskId(v string) *SwitchNetworkResponseBody
	GetTaskId() *string
}

type SwitchNetworkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F0997EE8-F4C2-4503-9168-85177ED78C70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 578678678
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SwitchNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchNetworkResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SwitchNetworkResponseBody) SetRequestId(v string) *SwitchNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchNetworkResponseBody) SetTaskId(v string) *SwitchNetworkResponseBody {
	s.TaskId = &v
	return s
}

func (s *SwitchNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
