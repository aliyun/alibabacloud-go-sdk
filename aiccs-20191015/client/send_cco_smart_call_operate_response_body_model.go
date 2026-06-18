// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCcoSmartCallOperateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendCcoSmartCallOperateResponseBody
	GetCode() *string
	SetData(v string) *SendCcoSmartCallOperateResponseBody
	GetData() *string
	SetMessage(v string) *SendCcoSmartCallOperateResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendCcoSmartCallOperateResponseBody
	GetRequestId() *string
}

type SendCcoSmartCallOperateResponseBody struct {
	// Request status code. A return value of OK indicates that the request succeeded.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Result of the command execution.
	//
	// - **true**: The command executed successfully.
	//
	// - **false**: The command execution failed.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendCcoSmartCallOperateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendCcoSmartCallOperateResponseBody) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallOperateResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendCcoSmartCallOperateResponseBody) GetData() *string {
	return s.Data
}

func (s *SendCcoSmartCallOperateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendCcoSmartCallOperateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendCcoSmartCallOperateResponseBody) SetCode(v string) *SendCcoSmartCallOperateResponseBody {
	s.Code = &v
	return s
}

func (s *SendCcoSmartCallOperateResponseBody) SetData(v string) *SendCcoSmartCallOperateResponseBody {
	s.Data = &v
	return s
}

func (s *SendCcoSmartCallOperateResponseBody) SetMessage(v string) *SendCcoSmartCallOperateResponseBody {
	s.Message = &v
	return s
}

func (s *SendCcoSmartCallOperateResponseBody) SetRequestId(v string) *SendCcoSmartCallOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCcoSmartCallOperateResponseBody) Validate() error {
	return dara.Validate(s)
}
