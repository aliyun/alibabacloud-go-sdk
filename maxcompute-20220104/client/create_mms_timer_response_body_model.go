// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateMmsTimerResponseBodyData) *CreateMmsTimerResponseBody
	GetData() *CreateMmsTimerResponseBodyData
	SetRequestId(v string) *CreateMmsTimerResponseBody
	GetRequestId() *string
}

type CreateMmsTimerResponseBody struct {
	Data *CreateMmsTimerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0be3e0b716671885050924814e3623
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMmsTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsTimerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMmsTimerResponseBody) GetData() *CreateMmsTimerResponseBodyData {
	return s.Data
}

func (s *CreateMmsTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMmsTimerResponseBody) SetData(v *CreateMmsTimerResponseBodyData) *CreateMmsTimerResponseBody {
	s.Data = v
	return s
}

func (s *CreateMmsTimerResponseBody) SetRequestId(v string) *CreateMmsTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMmsTimerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMmsTimerResponseBodyData struct {
	// timer id
	//
	// example:
	//
	// 1
	TimerId *int64 `json:"timerId,omitempty" xml:"timerId,omitempty"`
}

func (s CreateMmsTimerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsTimerResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMmsTimerResponseBodyData) GetTimerId() *int64 {
	return s.TimerId
}

func (s *CreateMmsTimerResponseBodyData) SetTimerId(v int64) *CreateMmsTimerResponseBodyData {
	s.TimerId = &v
	return s
}

func (s *CreateMmsTimerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
