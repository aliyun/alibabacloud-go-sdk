// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRobotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartRobotTaskResponseBody
	GetCode() *string
	SetData(v string) *StartRobotTaskResponseBody
	GetData() *string
	SetMessage(v string) *StartRobotTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartRobotTaskResponseBody
	GetRequestId() *string
}

type StartRobotTaskResponseBody struct {
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

func (s StartRobotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartRobotTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartRobotTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *StartRobotTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartRobotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRobotTaskResponseBody) SetCode(v string) *StartRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StartRobotTaskResponseBody) SetData(v string) *StartRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StartRobotTaskResponseBody) SetMessage(v string) *StartRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartRobotTaskResponseBody) SetRequestId(v string) *StartRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRobotTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
