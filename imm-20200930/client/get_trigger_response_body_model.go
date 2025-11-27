// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTriggerResponseBody
	GetRequestId() *string
	SetTrigger(v *DataIngestion) *GetTriggerResponseBody
	GetTrigger() *DataIngestion
}

type GetTriggerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4A7A2D0E-D8B8-4DA0-8127-EB32C6******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The trigger information.
	Trigger *DataIngestion `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s GetTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *GetTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTriggerResponseBody) GetTrigger() *DataIngestion {
	return s.Trigger
}

func (s *GetTriggerResponseBody) SetRequestId(v string) *GetTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTriggerResponseBody) SetTrigger(v *DataIngestion) *GetTriggerResponseBody {
	s.Trigger = v
	return s
}

func (s *GetTriggerResponseBody) Validate() error {
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}
