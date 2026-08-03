// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutInsightSelectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInsightSelectors(v string) *PutInsightSelectorsRequest
	GetInsightSelectors() *string
	SetTrailName(v string) *PutInsightSelectorsRequest
	GetTrailName() *string
}

type PutInsightSelectorsRequest struct {
	// The types of Insights events that the trail should deliver.
	//
	// example:
	//
	// [{"insightType":"AkInsight"},{"insightType":"IpInsight"}]
	InsightSelectors *string `json:"InsightSelectors,omitempty" xml:"InsightSelectors,omitempty"`
	// The name of the trail.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
}

func (s PutInsightSelectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutInsightSelectorsRequest) GoString() string {
	return s.String()
}

func (s *PutInsightSelectorsRequest) GetInsightSelectors() *string {
	return s.InsightSelectors
}

func (s *PutInsightSelectorsRequest) GetTrailName() *string {
	return s.TrailName
}

func (s *PutInsightSelectorsRequest) SetInsightSelectors(v string) *PutInsightSelectorsRequest {
	s.InsightSelectors = &v
	return s
}

func (s *PutInsightSelectorsRequest) SetTrailName(v string) *PutInsightSelectorsRequest {
	s.TrailName = &v
	return s
}

func (s *PutInsightSelectorsRequest) Validate() error {
	return dara.Validate(s)
}
