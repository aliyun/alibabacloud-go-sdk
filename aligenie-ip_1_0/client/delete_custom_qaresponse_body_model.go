// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomQAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteCustomQAResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomQAResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteCustomQAResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *DeleteCustomQAResponseBody
	GetStatusCode() *int32
}

type DeleteCustomQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteCustomQAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomQAResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomQAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomQAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomQAResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteCustomQAResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomQAResponseBody) SetMessage(v string) *DeleteCustomQAResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomQAResponseBody) SetRequestId(v string) *DeleteCustomQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomQAResponseBody) SetResult(v bool) *DeleteCustomQAResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteCustomQAResponseBody) SetStatusCode(v int32) *DeleteCustomQAResponseBody {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomQAResponseBody) Validate() error {
	return dara.Validate(s)
}
