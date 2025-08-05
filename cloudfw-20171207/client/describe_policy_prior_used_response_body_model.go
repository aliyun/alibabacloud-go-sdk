// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyPriorUsedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v int32) *DescribePolicyPriorUsedResponseBody
	GetEnd() *int32
	SetRequestId(v string) *DescribePolicyPriorUsedResponseBody
	GetRequestId() *string
	SetStart(v int32) *DescribePolicyPriorUsedResponseBody
	GetStart() *int32
}

type DescribePolicyPriorUsedResponseBody struct {
	// The lowest priority of existing access control policies.
	//
	// >  The value -1 indicates the lowest priority.
	//
	// example:
	//
	// 150
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The highest priority of existing access control policies.
	//
	// >  The value 0 indicates the highest priority.
	//
	// example:
	//
	// -1
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribePolicyPriorUsedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedResponseBody) GetEnd() *int32 {
	return s.End
}

func (s *DescribePolicyPriorUsedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolicyPriorUsedResponseBody) GetStart() *int32 {
	return s.Start
}

func (s *DescribePolicyPriorUsedResponseBody) SetEnd(v int32) *DescribePolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

func (s *DescribePolicyPriorUsedResponseBody) SetRequestId(v string) *DescribePolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyPriorUsedResponseBody) SetStart(v int32) *DescribePolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

func (s *DescribePolicyPriorUsedResponseBody) Validate() error {
	return dara.Validate(s)
}
