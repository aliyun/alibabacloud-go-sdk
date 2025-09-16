// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOnlineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOnlineConfigResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyOnlineConfigResponseBody
	GetResult() map[string]interface{}
}

type ModifyOnlineConfigResponseBody struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyOnlineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOnlineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOnlineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOnlineConfigResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyOnlineConfigResponseBody) SetRequestId(v string) *ModifyOnlineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOnlineConfigResponseBody) SetResult(v map[string]interface{}) *ModifyOnlineConfigResponseBody {
	s.Result = v
	return s
}

func (s *ModifyOnlineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
