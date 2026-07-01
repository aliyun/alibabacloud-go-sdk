// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBatchSmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SendBatchSmsResponseBody
	GetBizId() *string
	SetCode(v string) *SendBatchSmsResponseBody
	GetCode() *string
	SetMessage(v string) *SendBatchSmsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendBatchSmsResponseBody
	GetRequestId() *string
}

type SendBatchSmsResponseBody struct {
	// The delivery receipt ID.
	//
	// - You can use this ID to query the delivery status by calling the [QuerySendDetails](https://help.aliyun.com/document_detail/102352.html) operation.
	//
	// - Log on to the [Short Message Service console](https://dysms.console.aliyun.com/dysms.htm#/overview) and go to **Statistics*	- > **Delivery Logs*	- to view delivery details.
	//
	// example:
	//
	// 9006197469364984400
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The request status code.
	//
	// - A value of `OK` indicates a successful request.
	//
	// - For a list of other error codes, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the request status.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D230E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendBatchSmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendBatchSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendBatchSmsResponseBody) GetBizId() *string {
	return s.BizId
}

func (s *SendBatchSmsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendBatchSmsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendBatchSmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendBatchSmsResponseBody) SetBizId(v string) *SendBatchSmsResponseBody {
	s.BizId = &v
	return s
}

func (s *SendBatchSmsResponseBody) SetCode(v string) *SendBatchSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendBatchSmsResponseBody) SetMessage(v string) *SendBatchSmsResponseBody {
	s.Message = &v
	return s
}

func (s *SendBatchSmsResponseBody) SetRequestId(v string) *SendBatchSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendBatchSmsResponseBody) Validate() error {
	return dara.Validate(s)
}
