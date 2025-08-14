// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocGetTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PocGetTokenResponseBody
	GetCode() *string
	SetData(v string) *PocGetTokenResponseBody
	GetData() *string
	SetHttpStatusCode(v string) *PocGetTokenResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *PocGetTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *PocGetTokenResponseBody
	GetRequestId() *string
	SetSuccess(v string) *PocGetTokenResponseBody
	GetSuccess() *string
}

type PocGetTokenResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message
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
	// Whether the call was successful. true: Call succeeded. false: Call failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PocGetTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PocGetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *PocGetTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *PocGetTokenResponseBody) GetData() *string {
	return s.Data
}

func (s *PocGetTokenResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PocGetTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PocGetTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PocGetTokenResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PocGetTokenResponseBody) SetCode(v string) *PocGetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *PocGetTokenResponseBody) SetData(v string) *PocGetTokenResponseBody {
	s.Data = &v
	return s
}

func (s *PocGetTokenResponseBody) SetHttpStatusCode(v string) *PocGetTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PocGetTokenResponseBody) SetMessage(v string) *PocGetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *PocGetTokenResponseBody) SetRequestId(v string) *PocGetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *PocGetTokenResponseBody) SetSuccess(v string) *PocGetTokenResponseBody {
	s.Success = &v
	return s
}

func (s *PocGetTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
