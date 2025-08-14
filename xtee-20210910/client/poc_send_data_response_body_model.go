// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocSendDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PocSendDataResponseBody
	GetCode() *string
	SetData(v string) *PocSendDataResponseBody
	GetData() *string
	SetHttpStatusCode(v string) *PocSendDataResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *PocSendDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *PocSendDataResponseBody
	GetRequestId() *string
	SetSuccess(v string) *PocSendDataResponseBody
	GetSuccess() *string
}

type PocSendDataResponseBody struct {
	// API status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data result.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PocSendDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PocSendDataResponseBody) GoString() string {
	return s.String()
}

func (s *PocSendDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *PocSendDataResponseBody) GetData() *string {
	return s.Data
}

func (s *PocSendDataResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PocSendDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PocSendDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PocSendDataResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PocSendDataResponseBody) SetCode(v string) *PocSendDataResponseBody {
	s.Code = &v
	return s
}

func (s *PocSendDataResponseBody) SetData(v string) *PocSendDataResponseBody {
	s.Data = &v
	return s
}

func (s *PocSendDataResponseBody) SetHttpStatusCode(v string) *PocSendDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PocSendDataResponseBody) SetMessage(v string) *PocSendDataResponseBody {
	s.Message = &v
	return s
}

func (s *PocSendDataResponseBody) SetRequestId(v string) *PocSendDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PocSendDataResponseBody) SetSuccess(v string) *PocSendDataResponseBody {
	s.Success = &v
	return s
}

func (s *PocSendDataResponseBody) Validate() error {
	return dara.Validate(s)
}
