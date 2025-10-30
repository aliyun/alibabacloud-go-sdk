// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAxgGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OperateAxgGroupResponseBody
	GetCode() *string
	SetMessage(v string) *OperateAxgGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateAxgGroupResponseBody
	GetRequestId() *string
}

type OperateAxgGroupResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 986BCB6D-C9BF-42F9-91CE-3A9901233D36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateAxgGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateAxgGroupResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAxgGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *OperateAxgGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateAxgGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateAxgGroupResponseBody) SetCode(v string) *OperateAxgGroupResponseBody {
	s.Code = &v
	return s
}

func (s *OperateAxgGroupResponseBody) SetMessage(v string) *OperateAxgGroupResponseBody {
	s.Message = &v
	return s
}

func (s *OperateAxgGroupResponseBody) SetRequestId(v string) *OperateAxgGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAxgGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
