// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTableResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyTableResponseBody
	GetResult() map[string]interface{}
}

type ModifyTableResponseBody struct {
	// id of request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTableResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyTableResponseBody) SetRequestId(v string) *ModifyTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTableResponseBody) SetResult(v map[string]interface{}) *ModifyTableResponseBody {
	s.Result = v
	return s
}

func (s *ModifyTableResponseBody) Validate() error {
	return dara.Validate(s)
}
