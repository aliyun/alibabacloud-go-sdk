// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterOnlineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterOnlineConfigResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyClusterOnlineConfigResponseBody
	GetResult() map[string]interface{}
}

type ModifyClusterOnlineConfigResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyClusterOnlineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterOnlineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterOnlineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterOnlineConfigResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyClusterOnlineConfigResponseBody) SetRequestId(v string) *ModifyClusterOnlineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterOnlineConfigResponseBody) SetResult(v map[string]interface{}) *ModifyClusterOnlineConfigResponseBody {
	s.Result = v
	return s
}

func (s *ModifyClusterOnlineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
