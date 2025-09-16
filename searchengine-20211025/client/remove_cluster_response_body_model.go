// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveClusterResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *RemoveClusterResponseBody
	GetResult() map[string]interface{}
}

type RemoveClusterResponseBody struct {
	// id of request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveClusterResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *RemoveClusterResponseBody) SetRequestId(v string) *RemoveClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveClusterResponseBody) SetResult(v map[string]interface{}) *RemoveClusterResponseBody {
	s.Result = v
	return s
}

func (s *RemoveClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
