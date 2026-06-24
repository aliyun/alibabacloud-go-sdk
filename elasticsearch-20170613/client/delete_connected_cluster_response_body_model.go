// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectedClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConnectedClusterResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteConnectedClusterResponseBody
	GetResult() *bool
}

type DeleteConnectedClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation. Valid values:
	//
	// - true: The connected instance is removed.
	//
	// - false: The connected instance failed to be removed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteConnectedClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectedClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectedClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConnectedClusterResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteConnectedClusterResponseBody) SetRequestId(v string) *DeleteConnectedClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConnectedClusterResponseBody) SetResult(v bool) *DeleteConnectedClusterResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteConnectedClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
