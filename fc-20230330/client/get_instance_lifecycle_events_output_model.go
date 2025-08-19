// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLifecycleEventsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*InstanceEventItem) *GetInstanceLifecycleEventsOutput
	GetEvents() []*InstanceEventItem
	SetRequestId(v string) *GetInstanceLifecycleEventsOutput
	GetRequestId() *string
}

type GetInstanceLifecycleEventsOutput struct {
	Events    []*InstanceEventItem `json:"events" xml:"events" type:"Repeated"`
	RequestId *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetInstanceLifecycleEventsOutput) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLifecycleEventsOutput) GoString() string {
	return s.String()
}

func (s *GetInstanceLifecycleEventsOutput) GetEvents() []*InstanceEventItem {
	return s.Events
}

func (s *GetInstanceLifecycleEventsOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceLifecycleEventsOutput) SetEvents(v []*InstanceEventItem) *GetInstanceLifecycleEventsOutput {
	s.Events = v
	return s
}

func (s *GetInstanceLifecycleEventsOutput) SetRequestId(v string) *GetInstanceLifecycleEventsOutput {
	s.RequestId = &v
	return s
}

func (s *GetInstanceLifecycleEventsOutput) Validate() error {
	return dara.Validate(s)
}
