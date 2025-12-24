// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerSophonPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TriggerSophonPlaybookResponseBodyData) *TriggerSophonPlaybookResponseBody
	GetData() *TriggerSophonPlaybookResponseBodyData
	SetRequestId(v string) *TriggerSophonPlaybookResponseBody
	GetRequestId() *string
}

type TriggerSophonPlaybookResponseBody struct {
	// The details that is returned after the command or playbook is triggered.
	Data *TriggerSophonPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0DFC9403-54EB-5672-B690-9AA93C9EBB54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerSophonPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerSophonPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookResponseBody) GetData() *TriggerSophonPlaybookResponseBodyData {
	return s.Data
}

func (s *TriggerSophonPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerSophonPlaybookResponseBody) SetData(v *TriggerSophonPlaybookResponseBodyData) *TriggerSophonPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *TriggerSophonPlaybookResponseBody) SetRequestId(v string) *TriggerSophonPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerSophonPlaybookResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TriggerSophonPlaybookResponseBodyData struct {
	// The custom ID. If you do not specify this parameter when the playbook is triggered, a random ID is generated for fault locating and troubleshooting.
	//
	// example:
	//
	// a7c6d055-a72f-4676-bc89-3cd9edc0284c
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
}

func (s TriggerSophonPlaybookResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TriggerSophonPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookResponseBodyData) GetSophonTaskId() *string {
	return s.SophonTaskId
}

func (s *TriggerSophonPlaybookResponseBodyData) SetSophonTaskId(v string) *TriggerSophonPlaybookResponseBodyData {
	s.SophonTaskId = &v
	return s
}

func (s *TriggerSophonPlaybookResponseBodyData) Validate() error {
	return dara.Validate(s)
}
