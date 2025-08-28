// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartCallOperateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SmartCallOperateResponseBody
	GetCode() *string
	SetMessage(v string) *SmartCallOperateResponseBody
	GetMessage() *string
	SetRequestId(v string) *SmartCallOperateResponseBody
	GetRequestId() *string
	SetStatus(v bool) *SmartCallOperateResponseBody
	GetStatus() *bool
}

type SmartCallOperateResponseBody struct {
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
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The action result. Valid values:
	//
	// 	- **true**: The action was successful.
	//
	// 	- **false**: The action failed.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SmartCallOperateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SmartCallOperateResponseBody) GoString() string {
	return s.String()
}

func (s *SmartCallOperateResponseBody) GetCode() *string {
	return s.Code
}

func (s *SmartCallOperateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SmartCallOperateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SmartCallOperateResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *SmartCallOperateResponseBody) SetCode(v string) *SmartCallOperateResponseBody {
	s.Code = &v
	return s
}

func (s *SmartCallOperateResponseBody) SetMessage(v string) *SmartCallOperateResponseBody {
	s.Message = &v
	return s
}

func (s *SmartCallOperateResponseBody) SetRequestId(v string) *SmartCallOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmartCallOperateResponseBody) SetStatus(v bool) *SmartCallOperateResponseBody {
	s.Status = &v
	return s
}

func (s *SmartCallOperateResponseBody) Validate() error {
	return dara.Validate(s)
}
