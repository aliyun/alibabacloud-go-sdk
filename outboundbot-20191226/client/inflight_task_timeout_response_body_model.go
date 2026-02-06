// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInflightTaskTimeoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InflightTaskTimeoutResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *InflightTaskTimeoutResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InflightTaskTimeoutResponseBody
	GetMessage() *string
	SetRequestId(v string) *InflightTaskTimeoutResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InflightTaskTimeoutResponseBody
	GetSuccess() *bool
}

type InflightTaskTimeoutResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InflightTaskTimeoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InflightTaskTimeoutResponseBody) GoString() string {
	return s.String()
}

func (s *InflightTaskTimeoutResponseBody) GetCode() *string {
	return s.Code
}

func (s *InflightTaskTimeoutResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InflightTaskTimeoutResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InflightTaskTimeoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InflightTaskTimeoutResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InflightTaskTimeoutResponseBody) SetCode(v string) *InflightTaskTimeoutResponseBody {
	s.Code = &v
	return s
}

func (s *InflightTaskTimeoutResponseBody) SetHttpStatusCode(v int32) *InflightTaskTimeoutResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InflightTaskTimeoutResponseBody) SetMessage(v string) *InflightTaskTimeoutResponseBody {
	s.Message = &v
	return s
}

func (s *InflightTaskTimeoutResponseBody) SetRequestId(v string) *InflightTaskTimeoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *InflightTaskTimeoutResponseBody) SetSuccess(v bool) *InflightTaskTimeoutResponseBody {
	s.Success = &v
	return s
}

func (s *InflightTaskTimeoutResponseBody) Validate() error {
	return dara.Validate(s)
}
