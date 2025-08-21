// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstallCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReinstallCollectorResponseBody
	GetRequestId() *string
	SetResult(v bool) *ReinstallCollectorResponseBody
	GetResult() *bool
}

type ReinstallCollectorResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ReinstallCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReinstallCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *ReinstallCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReinstallCollectorResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ReinstallCollectorResponseBody) SetRequestId(v string) *ReinstallCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReinstallCollectorResponseBody) SetResult(v bool) *ReinstallCollectorResponseBody {
	s.Result = &v
	return s
}

func (s *ReinstallCollectorResponseBody) Validate() error {
	return dara.Validate(s)
}
