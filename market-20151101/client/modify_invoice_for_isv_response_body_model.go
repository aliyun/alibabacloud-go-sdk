// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInvoiceForIsvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyInvoiceForIsvResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyInvoiceForIsvResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyInvoiceForIsvResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyInvoiceForIsvResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyInvoiceForIsvResponseBody
	GetResult() *bool
}

type ModifyInvoiceForIsvResponseBody struct {
	// example:
	//
	// cmjj01**45
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5BD09171-MB74-18D8-890E-C70C067527BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyInvoiceForIsvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInvoiceForIsvResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInvoiceForIsvResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyInvoiceForIsvResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyInvoiceForIsvResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyInvoiceForIsvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInvoiceForIsvResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyInvoiceForIsvResponseBody) SetCode(v string) *ModifyInvoiceForIsvResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInvoiceForIsvResponseBody) SetHttpStatusCode(v int32) *ModifyInvoiceForIsvResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyInvoiceForIsvResponseBody) SetMessage(v string) *ModifyInvoiceForIsvResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyInvoiceForIsvResponseBody) SetRequestId(v string) *ModifyInvoiceForIsvResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInvoiceForIsvResponseBody) SetResult(v bool) *ModifyInvoiceForIsvResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyInvoiceForIsvResponseBody) Validate() error {
	return dara.Validate(s)
}
