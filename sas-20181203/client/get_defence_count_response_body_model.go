// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefenceCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefenceCount15Days(v int32) *GetDefenceCountResponseBody
	GetDefenceCount15Days() *int32
	SetDefenceCountTotal(v int32) *GetDefenceCountResponseBody
	GetDefenceCountTotal() *int32
	SetRequestId(v string) *GetDefenceCountResponseBody
	GetRequestId() *string
	SetSuspiciousDealtCount(v int32) *GetDefenceCountResponseBody
	GetSuspiciousDealtCount() *int32
	SetTamperProof15Days(v int32) *GetDefenceCountResponseBody
	GetTamperProof15Days() *int32
	SetTamperProofTotal(v int32) *GetDefenceCountResponseBody
	GetTamperProofTotal() *int32
}

type GetDefenceCountResponseBody struct {
	// The number of handled alerts of the precise defense type in the last 15 days.
	//
	// example:
	//
	// 1
	DefenceCount15Days *int32 `json:"DefenceCount15Days,omitempty" xml:"DefenceCount15Days,omitempty"`
	// The number of handled alerts of the precision defense type.
	//
	// example:
	//
	// 1
	DefenceCountTotal *int32 `json:"DefenceCountTotal,omitempty" xml:"DefenceCountTotal,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 89AD16CC-97EE-50F3-9B12-9E28E5C8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of handled security alerts of Cloud Security Center.
	//
	// example:
	//
	// 10
	SuspiciousDealtCount *int32 `json:"SuspiciousDealtCount,omitempty" xml:"SuspiciousDealtCount,omitempty"`
	// The number of handled alerts of the web tamper proofing type in the last 15 days.
	//
	// example:
	//
	// 2
	TamperProof15Days *int32 `json:"TamperProof15Days,omitempty" xml:"TamperProof15Days,omitempty"`
	// The number of handled alerts of the web tamper proofing type.
	//
	// example:
	//
	// 1
	TamperProofTotal *int32 `json:"TamperProofTotal,omitempty" xml:"TamperProofTotal,omitempty"`
}

func (s GetDefenceCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDefenceCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefenceCountResponseBody) GetDefenceCount15Days() *int32 {
	return s.DefenceCount15Days
}

func (s *GetDefenceCountResponseBody) GetDefenceCountTotal() *int32 {
	return s.DefenceCountTotal
}

func (s *GetDefenceCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDefenceCountResponseBody) GetSuspiciousDealtCount() *int32 {
	return s.SuspiciousDealtCount
}

func (s *GetDefenceCountResponseBody) GetTamperProof15Days() *int32 {
	return s.TamperProof15Days
}

func (s *GetDefenceCountResponseBody) GetTamperProofTotal() *int32 {
	return s.TamperProofTotal
}

func (s *GetDefenceCountResponseBody) SetDefenceCount15Days(v int32) *GetDefenceCountResponseBody {
	s.DefenceCount15Days = &v
	return s
}

func (s *GetDefenceCountResponseBody) SetDefenceCountTotal(v int32) *GetDefenceCountResponseBody {
	s.DefenceCountTotal = &v
	return s
}

func (s *GetDefenceCountResponseBody) SetRequestId(v string) *GetDefenceCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefenceCountResponseBody) SetSuspiciousDealtCount(v int32) *GetDefenceCountResponseBody {
	s.SuspiciousDealtCount = &v
	return s
}

func (s *GetDefenceCountResponseBody) SetTamperProof15Days(v int32) *GetDefenceCountResponseBody {
	s.TamperProof15Days = &v
	return s
}

func (s *GetDefenceCountResponseBody) SetTamperProofTotal(v int32) *GetDefenceCountResponseBody {
	s.TamperProofTotal = &v
	return s
}

func (s *GetDefenceCountResponseBody) Validate() error {
	return dara.Validate(s)
}
