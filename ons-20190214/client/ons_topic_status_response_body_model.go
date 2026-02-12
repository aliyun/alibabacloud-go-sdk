// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsTopicStatusResponseBodyData) *OnsTopicStatusResponseBody
	GetData() *OnsTopicStatusResponseBodyData
	SetRequestId(v string) *OnsTopicStatusResponseBody
	GetRequestId() *string
}

type OnsTopicStatusResponseBody struct {
	// The data structure of the queried topic.
	Data *OnsTopicStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 427EE49D-D762-41FB-8F3D-9BAC96C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicStatusResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusResponseBody) GetData() *OnsTopicStatusResponseBodyData {
	return s.Data
}

func (s *OnsTopicStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTopicStatusResponseBody) SetData(v *OnsTopicStatusResponseBodyData) *OnsTopicStatusResponseBody {
	s.Data = v
	return s
}

func (s *OnsTopicStatusResponseBody) SetRequestId(v string) *OnsTopicStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTopicStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTopicStatusResponseBodyData struct {
	// The point in time when the latest message is ready in the topic. If no message exists in the topic, the return value of this parameter is 0.
	//
	// The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// For information about the definition of ready messages and ready time, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// example:
	//
	// 1570864984364
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// Indicates the operations that you can perform on the topic. Valid values:
	//
	// 	- **2**: You can publish messages to the topic.
	//
	// 	- **4**: You can subscribe to the topic.
	//
	// 	- **6**: You can publish messages to and subscribe to the topic.
	//
	// example:
	//
	// 6
	Perm *int32 `json:"Perm,omitempty" xml:"Perm,omitempty"`
	// The total number of messages in all partitions of the topic.
	//
	// example:
	//
	// 2310
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s OnsTopicStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusResponseBodyData) GetLastTimeStamp() *int64 {
	return s.LastTimeStamp
}

func (s *OnsTopicStatusResponseBodyData) GetPerm() *int32 {
	return s.Perm
}

func (s *OnsTopicStatusResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *OnsTopicStatusResponseBodyData) SetLastTimeStamp(v int64) *OnsTopicStatusResponseBodyData {
	s.LastTimeStamp = &v
	return s
}

func (s *OnsTopicStatusResponseBodyData) SetPerm(v int32) *OnsTopicStatusResponseBodyData {
	s.Perm = &v
	return s
}

func (s *OnsTopicStatusResponseBodyData) SetTotalCount(v int64) *OnsTopicStatusResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *OnsTopicStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
