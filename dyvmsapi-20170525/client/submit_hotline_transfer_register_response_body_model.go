// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotlineTransferRegisterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitHotlineTransferRegisterResponseBody
	GetCode() *string
	SetData(v int64) *SubmitHotlineTransferRegisterResponseBody
	GetData() *int64
	SetMessage(v string) *SubmitHotlineTransferRegisterResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitHotlineTransferRegisterResponseBody
	GetRequestId() *string
}

type SubmitHotlineTransferRegisterResponseBody struct {
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
	// The registration ID.
	//
	// example:
	//
	// 2258****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitHotlineTransferRegisterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotlineTransferRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotlineTransferRegisterResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitHotlineTransferRegisterResponseBody) GetData() *int64 {
	return s.Data
}

func (s *SubmitHotlineTransferRegisterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitHotlineTransferRegisterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitHotlineTransferRegisterResponseBody) SetCode(v string) *SubmitHotlineTransferRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitHotlineTransferRegisterResponseBody) SetData(v int64) *SubmitHotlineTransferRegisterResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitHotlineTransferRegisterResponseBody) SetMessage(v string) *SubmitHotlineTransferRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitHotlineTransferRegisterResponseBody) SetRequestId(v string) *SubmitHotlineTransferRegisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotlineTransferRegisterResponseBody) Validate() error {
	return dara.Validate(s)
}
