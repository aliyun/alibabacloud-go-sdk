// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNodeConfigResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyNodeConfigResponseBody
	GetResult() map[string]interface{}
}

type ModifyNodeConfigResponseBody struct {
	// id of request
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyNodeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNodeConfigResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyNodeConfigResponseBody) SetRequestId(v string) *ModifyNodeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeConfigResponseBody) SetResult(v map[string]interface{}) *ModifyNodeConfigResponseBody {
	s.Result = v
	return s
}

func (s *ModifyNodeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
