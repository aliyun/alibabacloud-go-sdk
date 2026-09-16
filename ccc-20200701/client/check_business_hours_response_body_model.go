// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBusinessHoursResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckBusinessHoursResponseBody
	GetCode() *string
	SetData(v bool) *CheckBusinessHoursResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CheckBusinessHoursResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckBusinessHoursResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckBusinessHoursResponseBody
	GetRequestId() *string
}

type CheckBusinessHoursResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 678F7002-CA01-4ABF-A112-585AFBDF3A3B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckBusinessHoursResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckBusinessHoursResponseBody) GoString() string {
	return s.String()
}

func (s *CheckBusinessHoursResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckBusinessHoursResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckBusinessHoursResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckBusinessHoursResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckBusinessHoursResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckBusinessHoursResponseBody) SetCode(v string) *CheckBusinessHoursResponseBody {
	s.Code = &v
	return s
}

func (s *CheckBusinessHoursResponseBody) SetData(v bool) *CheckBusinessHoursResponseBody {
	s.Data = &v
	return s
}

func (s *CheckBusinessHoursResponseBody) SetHttpStatusCode(v int32) *CheckBusinessHoursResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckBusinessHoursResponseBody) SetMessage(v string) *CheckBusinessHoursResponseBody {
	s.Message = &v
	return s
}

func (s *CheckBusinessHoursResponseBody) SetRequestId(v string) *CheckBusinessHoursResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckBusinessHoursResponseBody) Validate() error {
	return dara.Validate(s)
}
