// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABTestFixedFlowDividersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListABTestFixedFlowDividersResponseBody
	GetRequestId() *string
	SetResult(v []*string) *ListABTestFixedFlowDividersResponseBody
	GetResult() []*string
}

type ListABTestFixedFlowDividersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestFixedFlowDividersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListABTestFixedFlowDividersResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestFixedFlowDividersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListABTestFixedFlowDividersResponseBody) GetResult() []*string {
	return s.Result
}

func (s *ListABTestFixedFlowDividersResponseBody) SetRequestId(v string) *ListABTestFixedFlowDividersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestFixedFlowDividersResponseBody) SetResult(v []*string) *ListABTestFixedFlowDividersResponseBody {
	s.Result = v
	return s
}

func (s *ListABTestFixedFlowDividersResponseBody) Validate() error {
	return dara.Validate(s)
}
