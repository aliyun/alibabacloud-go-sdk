// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRobotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelRobotTaskResponseBody
	GetCode() *string
	SetData(v string) *CancelRobotTaskResponseBody
	GetData() *string
	SetMessage(v string) *CancelRobotTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelRobotTaskResponseBody
	GetRequestId() *string
}

type CancelRobotTaskResponseBody struct {
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

func (s CancelRobotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRobotTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelRobotTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CancelRobotTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelRobotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelRobotTaskResponseBody) SetCode(v string) *CancelRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRobotTaskResponseBody) SetData(v string) *CancelRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelRobotTaskResponseBody) SetMessage(v string) *CancelRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelRobotTaskResponseBody) SetRequestId(v string) *CancelRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRobotTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
