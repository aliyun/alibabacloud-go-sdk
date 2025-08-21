// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateXpackMonitorConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateXpackMonitorConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateXpackMonitorConfigResponseBody
	GetResult() *bool
}

type UpdateXpackMonitorConfigResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateXpackMonitorConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateXpackMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateXpackMonitorConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateXpackMonitorConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateXpackMonitorConfigResponseBody) SetRequestId(v string) *UpdateXpackMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateXpackMonitorConfigResponseBody) SetResult(v bool) *UpdateXpackMonitorConfigResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateXpackMonitorConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
