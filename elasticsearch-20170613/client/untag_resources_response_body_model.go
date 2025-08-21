// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UntagResourcesResponseBody
	GetRequestId() *string
	SetResult(v bool) *UntagResourcesResponseBody
	GetResult() *bool
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return results:
	//
	// 	- true: deleted
	//
	// 	- false: Failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetResult(v bool) *UntagResourcesResponseBody {
	s.Result = &v
	return s
}

func (s *UntagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
