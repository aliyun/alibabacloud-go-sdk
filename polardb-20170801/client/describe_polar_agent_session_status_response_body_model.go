// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarAgentSessionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePolarAgentSessionStatusResponseBody
	GetRequestId() *string
	SetStatus(v int64) *DescribePolarAgentSessionStatusResponseBody
	GetStatus() *int64
}

type DescribePolarAgentSessionStatusResponseBody struct {
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePolarAgentSessionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentSessionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentSessionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarAgentSessionStatusResponseBody) GetStatus() *int64 {
	return s.Status
}

func (s *DescribePolarAgentSessionStatusResponseBody) SetRequestId(v string) *DescribePolarAgentSessionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarAgentSessionStatusResponseBody) SetStatus(v int64) *DescribePolarAgentSessionStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePolarAgentSessionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
