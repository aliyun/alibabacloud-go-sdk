// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestFixedFlowDividersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateABTestFixedFlowDividersResponseBody
	GetRequestId() *string
	SetResult(v []*string) *UpdateABTestFixedFlowDividersResponseBody
	GetResult() []*string
}

type UpdateABTestFixedFlowDividersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s UpdateABTestFixedFlowDividersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestFixedFlowDividersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestFixedFlowDividersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateABTestFixedFlowDividersResponseBody) GetResult() []*string {
	return s.Result
}

func (s *UpdateABTestFixedFlowDividersResponseBody) SetRequestId(v string) *UpdateABTestFixedFlowDividersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponseBody) SetResult(v []*string) *UpdateABTestFixedFlowDividersResponseBody {
	s.Result = v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponseBody) Validate() error {
	return dara.Validate(s)
}
