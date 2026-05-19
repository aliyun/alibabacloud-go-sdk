// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelOperatorOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetModelOperatorOrderResponseBodyData) *GetModelOperatorOrderResponseBody
	GetData() *GetModelOperatorOrderResponseBodyData
	SetMessage(v string) *GetModelOperatorOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetModelOperatorOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetModelOperatorOrderResponseBody
	GetSuccess() *bool
}

type GetModelOperatorOrderResponseBody struct {
	// The query result.
	Data *GetModelOperatorOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetModelOperatorOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelOperatorOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelOperatorOrderResponseBody) GetData() *GetModelOperatorOrderResponseBodyData {
	return s.Data
}

func (s *GetModelOperatorOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelOperatorOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelOperatorOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetModelOperatorOrderResponseBody) SetData(v *GetModelOperatorOrderResponseBodyData) *GetModelOperatorOrderResponseBody {
	s.Data = v
	return s
}

func (s *GetModelOperatorOrderResponseBody) SetMessage(v string) *GetModelOperatorOrderResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelOperatorOrderResponseBody) SetRequestId(v string) *GetModelOperatorOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelOperatorOrderResponseBody) SetSuccess(v bool) *GetModelOperatorOrderResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelOperatorOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelOperatorOrderResponseBodyData struct {
	// Indicates whether a valid order exists.
	//
	// example:
	//
	// true
	HasValidOrder *bool `json:"HasValidOrder,omitempty" xml:"HasValidOrder,omitempty"`
	// The instance list.
	InstanceList []*GetModelOperatorOrderResponseBodyDataInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
}

func (s GetModelOperatorOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModelOperatorOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelOperatorOrderResponseBodyData) GetHasValidOrder() *bool {
	return s.HasValidOrder
}

func (s *GetModelOperatorOrderResponseBodyData) GetInstanceList() []*GetModelOperatorOrderResponseBodyDataInstanceList {
	return s.InstanceList
}

func (s *GetModelOperatorOrderResponseBodyData) SetHasValidOrder(v bool) *GetModelOperatorOrderResponseBodyData {
	s.HasValidOrder = &v
	return s
}

func (s *GetModelOperatorOrderResponseBodyData) SetInstanceList(v []*GetModelOperatorOrderResponseBodyDataInstanceList) *GetModelOperatorOrderResponseBodyData {
	s.InstanceList = v
	return s
}

func (s *GetModelOperatorOrderResponseBodyData) Validate() error {
	if s.InstanceList != nil {
		for _, item := range s.InstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelOperatorOrderResponseBodyDataInstanceList struct {
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The instance end time (format: Timestamp).
	//
	// example:
	//
	// 1775145600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// instance type
	//
	// example:
	//
	// xlarge
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance name.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance start time (format: Timestamp).
	//
	// example:
	//
	// 1772439028000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The instance status.
	//
	// example:
	//
	// active/creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetModelOperatorOrderResponseBodyDataInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetModelOperatorOrderResponseBodyDataInstanceList) GoString() string {
	return s.String()
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) GetStatus() *string {
	return s.Status
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) SetChargeType(v string) *GetModelOperatorOrderResponseBodyDataInstanceList {
	s.ChargeType = &v
	return s
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) SetEndTime(v int64) *GetModelOperatorOrderResponseBodyDataInstanceList {
	s.EndTime = &v
	return s
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) SetInstanceClass(v string) *GetModelOperatorOrderResponseBodyDataInstanceList {
	s.InstanceClass = &v
	return s
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) SetInstanceId(v string) *GetModelOperatorOrderResponseBodyDataInstanceList {
	s.InstanceId = &v
	return s
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) SetStartTime(v int64) *GetModelOperatorOrderResponseBodyDataInstanceList {
	s.StartTime = &v
	return s
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) SetStatus(v string) *GetModelOperatorOrderResponseBodyDataInstanceList {
	s.Status = &v
	return s
}

func (s *GetModelOperatorOrderResponseBodyDataInstanceList) Validate() error {
	return dara.Validate(s)
}
