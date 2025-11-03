// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventBusesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEventBusesResponseBody
	GetCode() *string
	SetData(v *ListEventBusesResponseBodyData) *ListEventBusesResponseBody
	GetData() *ListEventBusesResponseBodyData
	SetMessage(v string) *ListEventBusesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEventBusesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEventBusesResponseBody
	GetSuccess() *bool
}

type ListEventBusesResponseBody struct {
	// The response code. Valid values:
	//
	// 	- Success: The request was successful.
	//
	// 	- Other codes: The request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListEventBusesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// InvalidArgument
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D1DCF64A-3F2C-5323-ADCB-3F4DF30FAD2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. If the operation was successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEventBusesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventBusesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventBusesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEventBusesResponseBody) GetData() *ListEventBusesResponseBodyData {
	return s.Data
}

func (s *ListEventBusesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEventBusesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventBusesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEventBusesResponseBody) SetCode(v string) *ListEventBusesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventBusesResponseBody) SetData(v *ListEventBusesResponseBodyData) *ListEventBusesResponseBody {
	s.Data = v
	return s
}

func (s *ListEventBusesResponseBody) SetMessage(v string) *ListEventBusesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEventBusesResponseBody) SetRequestId(v string) *ListEventBusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventBusesResponseBody) SetSuccess(v bool) *ListEventBusesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEventBusesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEventBusesResponseBodyData struct {
	// The event buses.
	EventBuses []*ListEventBusesResponseBodyDataEventBuses `json:"EventBuses,omitempty" xml:"EventBuses,omitempty" type:"Repeated"`
	// If excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEventBusesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEventBusesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEventBusesResponseBodyData) GetEventBuses() []*ListEventBusesResponseBodyDataEventBuses {
	return s.EventBuses
}

func (s *ListEventBusesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventBusesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListEventBusesResponseBodyData) SetEventBuses(v []*ListEventBusesResponseBodyDataEventBuses) *ListEventBusesResponseBodyData {
	s.EventBuses = v
	return s
}

func (s *ListEventBusesResponseBodyData) SetNextToken(v string) *ListEventBusesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListEventBusesResponseBodyData) SetTotal(v int32) *ListEventBusesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEventBusesResponseBodyData) Validate() error {
	if s.EventBuses != nil {
		for _, item := range s.EventBuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEventBusesResponseBodyDataEventBuses struct {
	// The timestamp that indicates when the event bus was created.
	//
	// example:
	//
	// 1607071602000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description.
	//
	// example:
	//
	// bus_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the event bus.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789098***:eventbus/default
	EventBusARN *string `json:"EventBusARN,omitempty" xml:"EventBusARN,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// default
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s ListEventBusesResponseBodyDataEventBuses) String() string {
	return dara.Prettify(s)
}

func (s ListEventBusesResponseBodyDataEventBuses) GoString() string {
	return s.String()
}

func (s *ListEventBusesResponseBodyDataEventBuses) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListEventBusesResponseBodyDataEventBuses) GetDescription() *string {
	return s.Description
}

func (s *ListEventBusesResponseBodyDataEventBuses) GetEventBusARN() *string {
	return s.EventBusARN
}

func (s *ListEventBusesResponseBodyDataEventBuses) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListEventBusesResponseBodyDataEventBuses) SetCreateTimestamp(v int64) *ListEventBusesResponseBodyDataEventBuses {
	s.CreateTimestamp = &v
	return s
}

func (s *ListEventBusesResponseBodyDataEventBuses) SetDescription(v string) *ListEventBusesResponseBodyDataEventBuses {
	s.Description = &v
	return s
}

func (s *ListEventBusesResponseBodyDataEventBuses) SetEventBusARN(v string) *ListEventBusesResponseBodyDataEventBuses {
	s.EventBusARN = &v
	return s
}

func (s *ListEventBusesResponseBodyDataEventBuses) SetEventBusName(v string) *ListEventBusesResponseBodyDataEventBuses {
	s.EventBusName = &v
	return s
}

func (s *ListEventBusesResponseBodyDataEventBuses) Validate() error {
	return dara.Validate(s)
}
