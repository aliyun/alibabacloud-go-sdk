// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigNetStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ConfigNetStatusResponseBody
	GetCode() *int64
	SetMessage(v string) *ConfigNetStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfigNetStatusResponseBody
	GetRequestId() *string
}

type ConfigNetStatusResponseBody struct {
	// The status code.
	//
	// - **200**: The request was successful.
	//
	// - Other codes: The request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B0949F09-B9C1-1D5E-8F27-0A5BF3CD5D95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigNetStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigNetStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigNetStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ConfigNetStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfigNetStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigNetStatusResponseBody) SetCode(v int64) *ConfigNetStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigNetStatusResponseBody) SetMessage(v string) *ConfigNetStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigNetStatusResponseBody) SetRequestId(v string) *ConfigNetStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigNetStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
