// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexPartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIndexPartitionResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyIndexPartitionResponseBody
	GetResult() map[string]interface{}
}

type ModifyIndexPartitionResponseBody struct {
	// id of request
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexPartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexPartitionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIndexPartitionResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyIndexPartitionResponseBody) SetRequestId(v string) *ModifyIndexPartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexPartitionResponseBody) SetResult(v map[string]interface{}) *ModifyIndexPartitionResponseBody {
	s.Result = v
	return s
}

func (s *ModifyIndexPartitionResponseBody) Validate() error {
	return dara.Validate(s)
}
