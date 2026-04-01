// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemovePluginResponseBody
	GetRequestId() *string
	SetResult(v bool) *RemovePluginResponseBody
	GetResult() *bool
}

type RemovePluginResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RemovePluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePluginResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePluginResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RemovePluginResponseBody) SetRequestId(v string) *RemovePluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePluginResponseBody) SetResult(v bool) *RemovePluginResponseBody {
	s.Result = &v
	return s
}

func (s *RemovePluginResponseBody) Validate() error {
	return dara.Validate(s)
}
