// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCcoSmartCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendCcoSmartCallResponseBody
	GetCode() *string
	SetData(v string) *SendCcoSmartCallResponseBody
	GetData() *string
	SetMessage(v string) *SendCcoSmartCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendCcoSmartCallResponseBody
	GetRequestId() *string
}

type SendCcoSmartCallResponseBody struct {
	// example:
	//
	// 16012854210^10281427****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendCcoSmartCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendCcoSmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendCcoSmartCallResponseBody) GetData() *string {
	return s.Data
}

func (s *SendCcoSmartCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendCcoSmartCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendCcoSmartCallResponseBody) SetCode(v string) *SendCcoSmartCallResponseBody {
	s.Code = &v
	return s
}

func (s *SendCcoSmartCallResponseBody) SetData(v string) *SendCcoSmartCallResponseBody {
	s.Data = &v
	return s
}

func (s *SendCcoSmartCallResponseBody) SetMessage(v string) *SendCcoSmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *SendCcoSmartCallResponseBody) SetRequestId(v string) *SendCcoSmartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCcoSmartCallResponseBody) Validate() error {
	return dara.Validate(s)
}
