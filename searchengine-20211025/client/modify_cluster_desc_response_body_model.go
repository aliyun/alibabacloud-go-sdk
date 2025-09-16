// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterDescResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterDescResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyClusterDescResponseBody
	GetResult() map[string]interface{}
}

type ModifyClusterDescResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyClusterDescResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterDescResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterDescResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterDescResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyClusterDescResponseBody) SetRequestId(v string) *ModifyClusterDescResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterDescResponseBody) SetResult(v map[string]interface{}) *ModifyClusterDescResponseBody {
	s.Result = v
	return s
}

func (s *ModifyClusterDescResponseBody) Validate() error {
	return dara.Validate(s)
}
