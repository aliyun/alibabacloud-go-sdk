// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtendConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExtendConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateExtendConfigResponseBody
	GetResult() *bool
}

type UpdateExtendConfigResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateExtendConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtendConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExtendConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExtendConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateExtendConfigResponseBody) SetRequestId(v string) *UpdateExtendConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExtendConfigResponseBody) SetResult(v bool) *UpdateExtendConfigResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateExtendConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
