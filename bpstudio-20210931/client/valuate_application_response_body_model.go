// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValuateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ValuateApplicationResponseBody
	GetCode() *string
	SetData(v int64) *ValuateApplicationResponseBody
	GetData() *int64
	SetMessage(v string) *ValuateApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ValuateApplicationResponseBody
	GetRequestId() *string
}

type ValuateApplicationResponseBody struct {
	// The code of the query task.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 123
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Idempotent notation
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned message.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValuateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValuateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ValuateApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ValuateApplicationResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ValuateApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ValuateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValuateApplicationResponseBody) SetCode(v string) *ValuateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ValuateApplicationResponseBody) SetData(v int64) *ValuateApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *ValuateApplicationResponseBody) SetMessage(v string) *ValuateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ValuateApplicationResponseBody) SetRequestId(v string) *ValuateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValuateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
