// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopIds(v []*string) *ModifyDesktopTimerResponseBody
	GetDesktopIds() []*string
	SetRequestId(v string) *ModifyDesktopTimerResponseBody
	GetRequestId() *string
}

type ModifyDesktopTimerResponseBody struct {
	// The IDs of the cloud computers for which you successfully configure the scheduled task.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4638719F-3CAB-5704-BD54-55617BFF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopTimerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopTimerResponseBody) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *ModifyDesktopTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopTimerResponseBody) SetDesktopIds(v []*string) *ModifyDesktopTimerResponseBody {
	s.DesktopIds = v
	return s
}

func (s *ModifyDesktopTimerResponseBody) SetRequestId(v string) *ModifyDesktopTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
