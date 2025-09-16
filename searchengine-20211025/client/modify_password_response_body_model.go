// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPasswordResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyPasswordResponseBody
	GetResult() map[string]interface{}
}

type ModifyPasswordResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 407BFD91-DE7D-50BA-8F88-CDE52A3B5E46
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPasswordResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyPasswordResponseBody) SetRequestId(v string) *ModifyPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPasswordResponseBody) SetResult(v map[string]interface{}) *ModifyPasswordResponseBody {
	s.Result = v
	return s
}

func (s *ModifyPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
