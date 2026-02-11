// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenarioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteScenarioResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteScenarioResponseBody
	GetResult() *bool
}

type DeleteScenarioResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteScenarioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScenarioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScenarioResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteScenarioResponseBody) SetRequestId(v string) *DeleteScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScenarioResponseBody) SetResult(v bool) *DeleteScenarioResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteScenarioResponseBody) Validate() error {
	return dara.Validate(s)
}
