// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderRobotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelOrderRobotTaskResponseBody
	GetCode() *string
	SetData(v string) *CancelOrderRobotTaskResponseBody
	GetData() *string
	SetMessage(v string) *CancelOrderRobotTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelOrderRobotTaskResponseBody
	GetRequestId() *string
}

type CancelOrderRobotTaskResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOrderRobotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderRobotTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelOrderRobotTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CancelOrderRobotTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelOrderRobotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelOrderRobotTaskResponseBody) SetCode(v string) *CancelOrderRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) SetData(v string) *CancelOrderRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) SetMessage(v string) *CancelOrderRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) SetRequestId(v string) *CancelOrderRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
