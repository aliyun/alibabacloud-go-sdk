// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySearcherReplicaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySearcherReplicaResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifySearcherReplicaResponseBody
	GetResult() map[string]interface{}
}

type ModifySearcherReplicaResponseBody struct {
	// example:
	//
	// e1eef569-1ff7-4bf8-acf7-1cecca9894ce
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifySearcherReplicaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySearcherReplicaResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySearcherReplicaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySearcherReplicaResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifySearcherReplicaResponseBody) SetRequestId(v string) *ModifySearcherReplicaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySearcherReplicaResponseBody) SetResult(v map[string]interface{}) *ModifySearcherReplicaResponseBody {
	s.Result = v
	return s
}

func (s *ModifySearcherReplicaResponseBody) Validate() error {
	return dara.Validate(s)
}
