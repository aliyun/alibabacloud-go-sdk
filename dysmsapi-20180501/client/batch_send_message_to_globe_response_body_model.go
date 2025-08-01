// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMessageToGlobeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedList(v string) *BatchSendMessageToGlobeResponseBody
	GetFailedList() *string
	SetFrom(v string) *BatchSendMessageToGlobeResponseBody
	GetFrom() *string
	SetMessageIdList(v string) *BatchSendMessageToGlobeResponseBody
	GetMessageIdList() *string
	SetRequestId(v string) *BatchSendMessageToGlobeResponseBody
	GetRequestId() *string
	SetResponseCode(v string) *BatchSendMessageToGlobeResponseBody
	GetResponseCode() *string
	SetResponseDescription(v string) *BatchSendMessageToGlobeResponseBody
	GetResponseDescription() *string
	SetSuccessCount(v string) *BatchSendMessageToGlobeResponseBody
	GetSuccessCount() *string
}

type BatchSendMessageToGlobeResponseBody struct {
	// The list of the mobile phone numbers that failed to receive the messages.
	//
	// example:
	//
	// ["931520581****","931530581****"]
	FailedList *string `json:"FailedList,omitempty" xml:"FailedList,omitempty"`
	// The sender ID that was returned. The API operation returns the sender ID that you have specified in the request parameters.
	//
	// example:
	//
	// Alicloud321
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the sent message.
	//
	// example:
	//
	// ["123****","124****"]
	MessageIdList *string `json:"MessageIdList,omitempty" xml:"MessageIdList,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D28D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code. If OK is returned, the request is successful. For more information, see [Error codes](https://help.aliyun.com/document_detail/180674.html).
	//
	// example:
	//
	// OK
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// The SMS Send Request was accepted
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	// The number of sent messages.
	//
	// example:
	//
	// 2
	SuccessCount *string `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s BatchSendMessageToGlobeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMessageToGlobeResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendMessageToGlobeResponseBody) GetFailedList() *string {
	return s.FailedList
}

func (s *BatchSendMessageToGlobeResponseBody) GetFrom() *string {
	return s.From
}

func (s *BatchSendMessageToGlobeResponseBody) GetMessageIdList() *string {
	return s.MessageIdList
}

func (s *BatchSendMessageToGlobeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSendMessageToGlobeResponseBody) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *BatchSendMessageToGlobeResponseBody) GetResponseDescription() *string {
	return s.ResponseDescription
}

func (s *BatchSendMessageToGlobeResponseBody) GetSuccessCount() *string {
	return s.SuccessCount
}

func (s *BatchSendMessageToGlobeResponseBody) SetFailedList(v string) *BatchSendMessageToGlobeResponseBody {
	s.FailedList = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetFrom(v string) *BatchSendMessageToGlobeResponseBody {
	s.From = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetMessageIdList(v string) *BatchSendMessageToGlobeResponseBody {
	s.MessageIdList = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetRequestId(v string) *BatchSendMessageToGlobeResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetResponseCode(v string) *BatchSendMessageToGlobeResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetResponseDescription(v string) *BatchSendMessageToGlobeResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetSuccessCount(v string) *BatchSendMessageToGlobeResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) Validate() error {
	return dara.Validate(s)
}
