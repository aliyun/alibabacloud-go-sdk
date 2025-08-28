// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRobotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRobotTaskResponseBody
	GetCode() *string
	SetData(v string) *DeleteRobotTaskResponseBody
	GetData() *string
	SetMessage(v string) *DeleteRobotTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRobotTaskResponseBody
	GetRequestId() *string
}

type DeleteRobotTaskResponseBody struct {
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

func (s DeleteRobotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRobotTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRobotTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteRobotTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRobotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRobotTaskResponseBody) SetCode(v string) *DeleteRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) SetData(v string) *DeleteRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) SetMessage(v string) *DeleteRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) SetRequestId(v string) *DeleteRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
