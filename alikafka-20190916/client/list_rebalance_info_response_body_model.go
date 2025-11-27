// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRebalanceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListRebalanceInfoResponseBody
	GetCode() *int32
	SetData(v *ListRebalanceInfoResponseBodyData) *ListRebalanceInfoResponseBody
	GetData() *ListRebalanceInfoResponseBodyData
	SetMessage(v string) *ListRebalanceInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRebalanceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRebalanceInfoResponseBody
	GetSuccess() *bool
}

type ListRebalanceInfoResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListRebalanceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRebalanceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRebalanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListRebalanceInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListRebalanceInfoResponseBody) GetData() *ListRebalanceInfoResponseBodyData {
	return s.Data
}

func (s *ListRebalanceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRebalanceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRebalanceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRebalanceInfoResponseBody) SetCode(v int32) *ListRebalanceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListRebalanceInfoResponseBody) SetData(v *ListRebalanceInfoResponseBodyData) *ListRebalanceInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListRebalanceInfoResponseBody) SetMessage(v string) *ListRebalanceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListRebalanceInfoResponseBody) SetRequestId(v string) *ListRebalanceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRebalanceInfoResponseBody) SetSuccess(v bool) *ListRebalanceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ListRebalanceInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRebalanceInfoResponseBodyData struct {
	RebalanceInfoList []*ListRebalanceInfoResponseBodyDataRebalanceInfoList `json:"RebalanceInfoList,omitempty" xml:"RebalanceInfoList,omitempty" type:"Repeated"`
}

func (s ListRebalanceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRebalanceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRebalanceInfoResponseBodyData) GetRebalanceInfoList() []*ListRebalanceInfoResponseBodyDataRebalanceInfoList {
	return s.RebalanceInfoList
}

func (s *ListRebalanceInfoResponseBodyData) SetRebalanceInfoList(v []*ListRebalanceInfoResponseBodyDataRebalanceInfoList) *ListRebalanceInfoResponseBodyData {
	s.RebalanceInfoList = v
	return s
}

func (s *ListRebalanceInfoResponseBodyData) Validate() error {
	if s.RebalanceInfoList != nil {
		for _, item := range s.RebalanceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRebalanceInfoResponseBodyDataRebalanceInfoList struct {
	// example:
	//
	// 100
	Generation *int64 `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// example:
	//
	// kafka-test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 1709199270
	LastRebalanceTimestamp *int64 `json:"LastRebalanceTimestamp,omitempty" xml:"LastRebalanceTimestamp,omitempty"`
	// example:
	//
	// removing member consumer-1-cd14eb9c-379b-4b8e-9bbd-76f147f8536f on LeaveGroup
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// true
	RebalanceSuccess *bool `json:"RebalanceSuccess,omitempty" xml:"RebalanceSuccess,omitempty"`
	// example:
	//
	// 12
	RebalanceTimeConsuming *int64 `json:"RebalanceTimeConsuming,omitempty" xml:"RebalanceTimeConsuming,omitempty"`
}

func (s ListRebalanceInfoResponseBodyDataRebalanceInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListRebalanceInfoResponseBodyDataRebalanceInfoList) GoString() string {
	return s.String()
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) GetGeneration() *int64 {
	return s.Generation
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) GetLastRebalanceTimestamp() *int64 {
	return s.LastRebalanceTimestamp
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) GetReason() *string {
	return s.Reason
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) GetRebalanceSuccess() *bool {
	return s.RebalanceSuccess
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) GetRebalanceTimeConsuming() *int64 {
	return s.RebalanceTimeConsuming
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) SetGeneration(v int64) *ListRebalanceInfoResponseBodyDataRebalanceInfoList {
	s.Generation = &v
	return s
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) SetGroupId(v string) *ListRebalanceInfoResponseBodyDataRebalanceInfoList {
	s.GroupId = &v
	return s
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) SetLastRebalanceTimestamp(v int64) *ListRebalanceInfoResponseBodyDataRebalanceInfoList {
	s.LastRebalanceTimestamp = &v
	return s
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) SetReason(v string) *ListRebalanceInfoResponseBodyDataRebalanceInfoList {
	s.Reason = &v
	return s
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) SetRebalanceSuccess(v bool) *ListRebalanceInfoResponseBodyDataRebalanceInfoList {
	s.RebalanceSuccess = &v
	return s
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) SetRebalanceTimeConsuming(v int64) *ListRebalanceInfoResponseBodyDataRebalanceInfoList {
	s.RebalanceTimeConsuming = &v
	return s
}

func (s *ListRebalanceInfoResponseBodyDataRebalanceInfoList) Validate() error {
	return dara.Validate(s)
}
