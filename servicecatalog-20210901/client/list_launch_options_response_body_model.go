// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLaunchOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchOptionSummaries(v []*ListLaunchOptionsResponseBodyLaunchOptionSummaries) *ListLaunchOptionsResponseBody
	GetLaunchOptionSummaries() []*ListLaunchOptionsResponseBodyLaunchOptionSummaries
	SetRequestId(v string) *ListLaunchOptionsResponseBody
	GetRequestId() *string
}

type ListLaunchOptionsResponseBody struct {
	// The launch options.
	LaunchOptionSummaries []*ListLaunchOptionsResponseBodyLaunchOptionSummaries `json:"LaunchOptionSummaries,omitempty" xml:"LaunchOptionSummaries,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLaunchOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLaunchOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsResponseBody) GetLaunchOptionSummaries() []*ListLaunchOptionsResponseBodyLaunchOptionSummaries {
	return s.LaunchOptionSummaries
}

func (s *ListLaunchOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLaunchOptionsResponseBody) SetLaunchOptionSummaries(v []*ListLaunchOptionsResponseBodyLaunchOptionSummaries) *ListLaunchOptionsResponseBody {
	s.LaunchOptionSummaries = v
	return s
}

func (s *ListLaunchOptionsResponseBody) SetRequestId(v string) *ListLaunchOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLaunchOptionsResponseBody) Validate() error {
	if s.LaunchOptionSummaries != nil {
		for _, item := range s.LaunchOptionSummaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLaunchOptionsResponseBodyLaunchOptionSummaries struct {
	// The constraints.
	ConstraintSummaries []*ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries `json:"ConstraintSummaries,omitempty" xml:"ConstraintSummaries,omitempty" type:"Repeated"`
	// The ID of the product portfolio.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The name of the product portfolio.
	//
	// example:
	//
	// DEMO-IT services
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
}

func (s ListLaunchOptionsResponseBodyLaunchOptionSummaries) String() string {
	return dara.Prettify(s)
}

func (s ListLaunchOptionsResponseBodyLaunchOptionSummaries) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) GetConstraintSummaries() []*ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries {
	return s.ConstraintSummaries
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) GetPortfolioName() *string {
	return s.PortfolioName
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) SetConstraintSummaries(v []*ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) *ListLaunchOptionsResponseBodyLaunchOptionSummaries {
	s.ConstraintSummaries = v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) SetPortfolioId(v string) *ListLaunchOptionsResponseBodyLaunchOptionSummaries {
	s.PortfolioId = &v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) SetPortfolioName(v string) *ListLaunchOptionsResponseBodyLaunchOptionSummaries {
	s.PortfolioName = &v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) Validate() error {
	if s.ConstraintSummaries != nil {
		for _, item := range s.ConstraintSummaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries struct {
	// The type of the constraint. Valid values:
	//
	// 1.  Launch
	//
	// 2.  Template
	//
	// 3.  Approval
	//
	// 4.  StackRegion
	//
	// example:
	//
	// Launch
	ConstraintType *string `json:"ConstraintType,omitempty" xml:"ConstraintType,omitempty"`
	// The description of the constraint.
	//
	// example:
	//
	// Launch as local role TestRole
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The types of operations that require approval. If the type of the constraint is Approval, this parameter is not empty.
	OperationTypes []*string `json:"OperationTypes,omitempty" xml:"OperationTypes,omitempty" type:"Repeated"`
	// The regions in which users can launch the service. If the type of the constraint is StackRegion, this parameter is not empty.
	StackRegions []*string `json:"StackRegions,omitempty" xml:"StackRegions,omitempty" type:"Repeated"`
}

func (s ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) String() string {
	return dara.Prettify(s)
}

func (s ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) GetConstraintType() *string {
	return s.ConstraintType
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) GetDescription() *string {
	return s.Description
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) GetOperationTypes() []*string {
	return s.OperationTypes
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) GetStackRegions() []*string {
	return s.StackRegions
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) SetConstraintType(v string) *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries {
	s.ConstraintType = &v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) SetDescription(v string) *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries {
	s.Description = &v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) SetOperationTypes(v []*string) *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries {
	s.OperationTypes = v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) SetStackRegions(v []*string) *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries {
	s.StackRegions = v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) Validate() error {
	return dara.Validate(s)
}
