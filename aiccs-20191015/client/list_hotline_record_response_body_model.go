// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotlineRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHotlineRecordResponseBody
	GetCode() *string
	SetData(v []*ListHotlineRecordResponseBodyData) *ListHotlineRecordResponseBody
	GetData() []*ListHotlineRecordResponseBodyData
	SetMessage(v string) *ListHotlineRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotlineRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHotlineRecordResponseBody
	GetSuccess() *bool
}

type ListHotlineRecordResponseBody struct {
	// Status code. A return value of Success indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Hotline session information.
	Data []*ListHotlineRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotlineRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHotlineRecordResponseBody) GetData() []*ListHotlineRecordResponseBodyData {
	return s.Data
}

func (s *ListHotlineRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotlineRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotlineRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHotlineRecordResponseBody) SetCode(v string) *ListHotlineRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetData(v []*ListHotlineRecordResponseBodyData) *ListHotlineRecordResponseBody {
	s.Data = v
	return s
}

func (s *ListHotlineRecordResponseBody) SetMessage(v string) *ListHotlineRecordResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetRequestId(v string) *ListHotlineRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetSuccess(v bool) *ListHotlineRecordResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotlineRecordResponseBody) Validate() error {
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

type ListHotlineRecordResponseBodyData struct {
	// Session ID. This corresponds to the acid in WebSocket after an inbound call.
	//
	// example:
	//
	// 100365558
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Connection ID.
	//
	// example:
	//
	// 100365548
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// Recording end UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 16128694810
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Recording start UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 16128694110
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Recording file URL.
	//
	// example:
	//
	// http://aliccrec-shvpc.oss-cn-shanghai.aliyuncs.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotlineRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponseBodyData) GetCallId() *string {
	return s.CallId
}

func (s *ListHotlineRecordResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *ListHotlineRecordResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListHotlineRecordResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHotlineRecordResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListHotlineRecordResponseBodyData) SetCallId(v string) *ListHotlineRecordResponseBodyData {
	s.CallId = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetConnectionId(v string) *ListHotlineRecordResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetEndTime(v int64) *ListHotlineRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetStartTime(v int64) *ListHotlineRecordResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetUrl(v string) *ListHotlineRecordResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
