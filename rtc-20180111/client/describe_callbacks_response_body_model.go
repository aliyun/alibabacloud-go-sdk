// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCallbacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallbacks(v []*DescribeCallbacksResponseBodyCallbacks) *DescribeCallbacksResponseBody
	GetCallbacks() []*DescribeCallbacksResponseBodyCallbacks
	SetRequestId(v string) *DescribeCallbacksResponseBody
	GetRequestId() *string
}

type DescribeCallbacksResponseBody struct {
	// A list of subscribed events for the application.
	Callbacks []*DescribeCallbacksResponseBodyCallbacks `json:"Callbacks,omitempty" xml:"Callbacks,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCallbacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallbacksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCallbacksResponseBody) GetCallbacks() []*DescribeCallbacksResponseBodyCallbacks {
	return s.Callbacks
}

func (s *DescribeCallbacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCallbacksResponseBody) SetCallbacks(v []*DescribeCallbacksResponseBodyCallbacks) *DescribeCallbacksResponseBody {
	s.Callbacks = v
	return s
}

func (s *DescribeCallbacksResponseBody) SetRequestId(v string) *DescribeCallbacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCallbacksResponseBody) Validate() error {
	if s.Callbacks != nil {
		for _, item := range s.Callbacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCallbacksResponseBodyCallbacks struct {
	// The callback code.
	//
	// example:
	//
	// RecordEvent
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The callback validation status. 0 indicates validating. 1 indicates validation successful. 2 indicates validation failed.
	//
	// example:
	//
	// 1
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The callback validation code.
	//
	// example:
	//
	// RESPONSE_INVALID
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The webhook address.
	//
	// example:
	//
	// https://www.aliyun.com
	Conf *string `json:"Conf,omitempty" xml:"Conf,omitempty"`
	// The callback validation message.
	//
	// example:
	//
	// Success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The callback status. 0 indicates enabled. 1 indicates disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-events. For more information, see Callback overview.
	//
	// Recording events:
	//
	// - **2010**: The service status of the recording changes.
	//
	// - **2011**: The status of the recording audio stream changes.
	//
	// - **2012**: The status of the recording video stream changes.
	SubEvent []*int32 `json:"SubEvent,omitempty" xml:"SubEvent,omitempty" type:"Repeated"`
}

func (s DescribeCallbacksResponseBodyCallbacks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallbacksResponseBodyCallbacks) GoString() string {
	return s.String()
}

func (s *DescribeCallbacksResponseBodyCallbacks) GetCategory() *string {
	return s.Category
}

func (s *DescribeCallbacksResponseBodyCallbacks) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *DescribeCallbacksResponseBodyCallbacks) GetCode() *string {
	return s.Code
}

func (s *DescribeCallbacksResponseBodyCallbacks) GetConf() *string {
	return s.Conf
}

func (s *DescribeCallbacksResponseBodyCallbacks) GetMsg() *string {
	return s.Msg
}

func (s *DescribeCallbacksResponseBodyCallbacks) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCallbacksResponseBodyCallbacks) GetSubEvent() []*int32 {
	return s.SubEvent
}

func (s *DescribeCallbacksResponseBodyCallbacks) SetCategory(v string) *DescribeCallbacksResponseBodyCallbacks {
	s.Category = &v
	return s
}

func (s *DescribeCallbacksResponseBodyCallbacks) SetCheckStatus(v string) *DescribeCallbacksResponseBodyCallbacks {
	s.CheckStatus = &v
	return s
}

func (s *DescribeCallbacksResponseBodyCallbacks) SetCode(v string) *DescribeCallbacksResponseBodyCallbacks {
	s.Code = &v
	return s
}

func (s *DescribeCallbacksResponseBodyCallbacks) SetConf(v string) *DescribeCallbacksResponseBodyCallbacks {
	s.Conf = &v
	return s
}

func (s *DescribeCallbacksResponseBodyCallbacks) SetMsg(v string) *DescribeCallbacksResponseBodyCallbacks {
	s.Msg = &v
	return s
}

func (s *DescribeCallbacksResponseBodyCallbacks) SetStatus(v int32) *DescribeCallbacksResponseBodyCallbacks {
	s.Status = &v
	return s
}

func (s *DescribeCallbacksResponseBodyCallbacks) SetSubEvent(v []*int32) *DescribeCallbacksResponseBodyCallbacks {
	s.SubEvent = v
	return s
}

func (s *DescribeCallbacksResponseBodyCallbacks) Validate() error {
	return dara.Validate(s)
}
