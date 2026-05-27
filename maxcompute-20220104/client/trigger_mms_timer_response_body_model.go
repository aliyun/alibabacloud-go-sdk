// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerMmsTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *TriggerMmsTimerResponseBody
	GetData() *int64
	SetRequestId(v string) *TriggerMmsTimerResponseBody
	GetRequestId() *string
}

type TriggerMmsTimerResponseBody struct {
	// timer id
	//
	// example:
	//
	// 1
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 0be3e0b716671885050924814e3623
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TriggerMmsTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerMmsTimerResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerMmsTimerResponseBody) GetData() *int64 {
	return s.Data
}

func (s *TriggerMmsTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerMmsTimerResponseBody) SetData(v int64) *TriggerMmsTimerResponseBody {
	s.Data = &v
	return s
}

func (s *TriggerMmsTimerResponseBody) SetRequestId(v string) *TriggerMmsTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerMmsTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
