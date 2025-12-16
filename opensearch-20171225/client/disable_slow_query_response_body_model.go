// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSlowQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableSlowQueryResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DisableSlowQueryResponseBody
	GetResult() map[string]interface{}
}

type DisableSlowQueryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 728E89C6-8673-D39B-39A1-5BA2B56D448F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DisableSlowQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSlowQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSlowQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSlowQueryResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DisableSlowQueryResponseBody) SetRequestId(v string) *DisableSlowQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSlowQueryResponseBody) SetResult(v map[string]interface{}) *DisableSlowQueryResponseBody {
	s.Result = v
	return s
}

func (s *DisableSlowQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
