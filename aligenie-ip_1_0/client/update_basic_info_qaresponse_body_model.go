// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicInfoQAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateBasicInfoQAResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBasicInfoQAResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateBasicInfoQAResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *UpdateBasicInfoQAResponseBody
	GetStatusCode() *int32
}

type UpdateBasicInfoQAResponseBody struct {
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

func (s UpdateBasicInfoQAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicInfoQAResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicInfoQAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBasicInfoQAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBasicInfoQAResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateBasicInfoQAResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBasicInfoQAResponseBody) SetMessage(v string) *UpdateBasicInfoQAResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBasicInfoQAResponseBody) SetRequestId(v string) *UpdateBasicInfoQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBasicInfoQAResponseBody) SetResult(v bool) *UpdateBasicInfoQAResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateBasicInfoQAResponseBody) SetStatusCode(v int32) *UpdateBasicInfoQAResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicInfoQAResponseBody) Validate() error {
	return dara.Validate(s)
}
