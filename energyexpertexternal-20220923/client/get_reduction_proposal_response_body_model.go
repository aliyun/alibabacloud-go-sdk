// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReductionProposalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetReductionProposalResponseBodyData) *GetReductionProposalResponseBody
	GetData() *GetReductionProposalResponseBodyData
	SetRequestId(v string) *GetReductionProposalResponseBody
	GetRequestId() *string
}

type GetReductionProposalResponseBody struct {
	// The returned data.
	Data *GetReductionProposalResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetReductionProposalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReductionProposalResponseBody) GoString() string {
	return s.String()
}

func (s *GetReductionProposalResponseBody) GetData() *GetReductionProposalResponseBodyData {
	return s.Data
}

func (s *GetReductionProposalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetReductionProposalResponseBody) SetData(v *GetReductionProposalResponseBodyData) *GetReductionProposalResponseBody {
	s.Data = v
	return s
}

func (s *GetReductionProposalResponseBody) SetRequestId(v string) *GetReductionProposalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReductionProposalResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReductionProposalResponseBodyData struct {
	// Proactive carbon reduction recommendations and advice.
	//
	// example:
	//
	// Reduce one-drop usage
	Reduction *string `json:"reduction,omitempty" xml:"reduction,omitempty"`
	// Active carbon reduction assessment.
	//
	// example:
	//
	// Trying Energy Expert for a more detailed assessment.
	ReductionEvaluation *string `json:"reductionEvaluation,omitempty" xml:"reductionEvaluation,omitempty"`
}

func (s GetReductionProposalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetReductionProposalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetReductionProposalResponseBodyData) GetReduction() *string {
	return s.Reduction
}

func (s *GetReductionProposalResponseBodyData) GetReductionEvaluation() *string {
	return s.ReductionEvaluation
}

func (s *GetReductionProposalResponseBodyData) SetReduction(v string) *GetReductionProposalResponseBodyData {
	s.Reduction = &v
	return s
}

func (s *GetReductionProposalResponseBodyData) SetReductionEvaluation(v string) *GetReductionProposalResponseBodyData {
	s.ReductionEvaluation = &v
	return s
}

func (s *GetReductionProposalResponseBodyData) Validate() error {
	return dara.Validate(s)
}
