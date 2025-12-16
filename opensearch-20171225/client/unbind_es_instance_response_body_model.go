// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindEsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindEsInstanceResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *UnbindEsInstanceResponseBody
	GetResult() map[string]interface{}
}

type UnbindEsInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UnbindEsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindEsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindEsInstanceResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *UnbindEsInstanceResponseBody) SetRequestId(v string) *UnbindEsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindEsInstanceResponseBody) SetResult(v map[string]interface{}) *UnbindEsInstanceResponseBody {
	s.Result = v
	return s
}

func (s *UnbindEsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
