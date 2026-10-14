// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTotalAttackMaxFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBps(v int64) *DescribeTotalAttackMaxFlowResponseBody
	GetBps() *int64
	SetPps(v int64) *DescribeTotalAttackMaxFlowResponseBody
	GetPps() *int64
	SetRequestId(v string) *DescribeTotalAttackMaxFlowResponseBody
	GetRequestId() *string
}

type DescribeTotalAttackMaxFlowResponseBody struct {
	Bps       *int64  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Pps       *int64  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTotalAttackMaxFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTotalAttackMaxFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTotalAttackMaxFlowResponseBody) GetBps() *int64 {
	return s.Bps
}

func (s *DescribeTotalAttackMaxFlowResponseBody) GetPps() *int64 {
	return s.Pps
}

func (s *DescribeTotalAttackMaxFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTotalAttackMaxFlowResponseBody) SetBps(v int64) *DescribeTotalAttackMaxFlowResponseBody {
	s.Bps = &v
	return s
}

func (s *DescribeTotalAttackMaxFlowResponseBody) SetPps(v int64) *DescribeTotalAttackMaxFlowResponseBody {
	s.Pps = &v
	return s
}

func (s *DescribeTotalAttackMaxFlowResponseBody) SetRequestId(v string) *DescribeTotalAttackMaxFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTotalAttackMaxFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
