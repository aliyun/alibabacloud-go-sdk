// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScenarioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyScenarioResponseBody
	GetRequestId() *string
	SetResult(v string) *ApplyScenarioResponseBody
	GetResult() *string
}

type ApplyScenarioResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EA24D522-AD35-47B8-8CB2-ADBC38******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code of the business monitoring job, which is the coloring sign.
	//
	// example:
	//
	// 2b97****
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ApplyScenarioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyScenarioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyScenarioResponseBody) GetResult() *string {
	return s.Result
}

func (s *ApplyScenarioResponseBody) SetRequestId(v string) *ApplyScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyScenarioResponseBody) SetResult(v string) *ApplyScenarioResponseBody {
	s.Result = &v
	return s
}

func (s *ApplyScenarioResponseBody) Validate() error {
	return dara.Validate(s)
}
