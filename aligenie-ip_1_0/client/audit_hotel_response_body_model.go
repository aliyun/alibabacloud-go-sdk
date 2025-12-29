// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditHotelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AuditHotelResponseBody
	GetCode() *int32
	SetMessage(v string) *AuditHotelResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuditHotelResponseBody
	GetRequestId() *string
	SetResult(v bool) *AuditHotelResponseBody
	GetResult() *bool
}

type AuditHotelResponseBody struct {
	// example:
	//
	// 200
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AuditHotelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuditHotelResponseBody) GoString() string {
	return s.String()
}

func (s *AuditHotelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AuditHotelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuditHotelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuditHotelResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AuditHotelResponseBody) SetCode(v int32) *AuditHotelResponseBody {
	s.Code = &v
	return s
}

func (s *AuditHotelResponseBody) SetMessage(v string) *AuditHotelResponseBody {
	s.Message = &v
	return s
}

func (s *AuditHotelResponseBody) SetRequestId(v string) *AuditHotelResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuditHotelResponseBody) SetResult(v bool) *AuditHotelResponseBody {
	s.Result = &v
	return s
}

func (s *AuditHotelResponseBody) Validate() error {
	return dara.Validate(s)
}
