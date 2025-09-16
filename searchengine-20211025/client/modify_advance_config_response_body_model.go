// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAdvanceConfigResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyAdvanceConfigResponseBody
	GetResult() map[string]interface{}
}

type ModifyAdvanceConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyAdvanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAdvanceConfigResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyAdvanceConfigResponseBody) SetRequestId(v string) *ModifyAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAdvanceConfigResponseBody) SetResult(v map[string]interface{}) *ModifyAdvanceConfigResponseBody {
	s.Result = v
	return s
}

func (s *ModifyAdvanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
