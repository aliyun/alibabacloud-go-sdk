// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueRateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetQueueRateResponseBody
	GetCode() *int32
	SetData(v []*GetQueueRateResponseBodyData) *GetQueueRateResponseBody
	GetData() []*GetQueueRateResponseBodyData
	SetMessage(v string) *GetQueueRateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueueRateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQueueRateResponseBody
	GetSuccess() *bool
}

type GetQueueRateResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetQueueRateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueueRateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueueRateResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueRateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetQueueRateResponseBody) GetData() []*GetQueueRateResponseBodyData {
	return s.Data
}

func (s *GetQueueRateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueueRateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueueRateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQueueRateResponseBody) SetCode(v int32) *GetQueueRateResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueueRateResponseBody) SetData(v []*GetQueueRateResponseBodyData) *GetQueueRateResponseBody {
	s.Data = v
	return s
}

func (s *GetQueueRateResponseBody) SetMessage(v string) *GetQueueRateResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueueRateResponseBody) SetRequestId(v string) *GetQueueRateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueRateResponseBody) SetSuccess(v bool) *GetQueueRateResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueueRateResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQueueRateResponseBodyData struct {
	InQps      *int64  `json:"InQps,omitempty" xml:"InQps,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutQps     *int64  `json:"OutQps,omitempty" xml:"OutQps,omitempty"`
	QueueName  *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	VhostName  *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetQueueRateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueueRateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueueRateResponseBodyData) GetInQps() *int64 {
	return s.InQps
}

func (s *GetQueueRateResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueueRateResponseBodyData) GetOutQps() *int64 {
	return s.OutQps
}

func (s *GetQueueRateResponseBodyData) GetQueueName() *string {
	return s.QueueName
}

func (s *GetQueueRateResponseBodyData) GetVhostName() *string {
	return s.VhostName
}

func (s *GetQueueRateResponseBodyData) SetInQps(v int64) *GetQueueRateResponseBodyData {
	s.InQps = &v
	return s
}

func (s *GetQueueRateResponseBodyData) SetInstanceId(v string) *GetQueueRateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetQueueRateResponseBodyData) SetOutQps(v int64) *GetQueueRateResponseBodyData {
	s.OutQps = &v
	return s
}

func (s *GetQueueRateResponseBodyData) SetQueueName(v string) *GetQueueRateResponseBodyData {
	s.QueueName = &v
	return s
}

func (s *GetQueueRateResponseBodyData) SetVhostName(v string) *GetQueueRateResponseBodyData {
	s.VhostName = &v
	return s
}

func (s *GetQueueRateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
