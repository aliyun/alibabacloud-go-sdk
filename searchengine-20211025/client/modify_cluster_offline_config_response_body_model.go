// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterOfflineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterOfflineConfigResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyClusterOfflineConfigResponseBody
	GetResult() map[string]interface{}
}

type ModifyClusterOfflineConfigResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyClusterOfflineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterOfflineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterOfflineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterOfflineConfigResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyClusterOfflineConfigResponseBody) SetRequestId(v string) *ModifyClusterOfflineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterOfflineConfigResponseBody) SetResult(v map[string]interface{}) *ModifyClusterOfflineConfigResponseBody {
	s.Result = v
	return s
}

func (s *ModifyClusterOfflineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
