// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AddCustomQAResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCustomQAResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddCustomQAResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *AddCustomQAResponseBody
	GetStatusCode() *int32
}

type AddCustomQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7***726E
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

func (s AddCustomQAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomQAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCustomQAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomQAResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddCustomQAResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomQAResponseBody) SetMessage(v string) *AddCustomQAResponseBody {
	s.Message = &v
	return s
}

func (s *AddCustomQAResponseBody) SetRequestId(v string) *AddCustomQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomQAResponseBody) SetResult(v bool) *AddCustomQAResponseBody {
	s.Result = &v
	return s
}

func (s *AddCustomQAResponseBody) SetStatusCode(v int32) *AddCustomQAResponseBody {
	s.StatusCode = &v
	return s
}

func (s *AddCustomQAResponseBody) Validate() error {
	return dara.Validate(s)
}
