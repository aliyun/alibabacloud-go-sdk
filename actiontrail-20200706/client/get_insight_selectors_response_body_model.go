// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInsightSelectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInsightSelectors(v []*string) *GetInsightSelectorsResponseBody
	GetInsightSelectors() []*string
	SetRequestId(v string) *GetInsightSelectorsResponseBody
	GetRequestId() *string
	SetTrailArn(v string) *GetInsightSelectorsResponseBody
	GetTrailArn() *string
}

type GetInsightSelectorsResponseBody struct {
	// An array of Insights event types.
	InsightSelectors []*string `json:"InsightSelectors,omitempty" xml:"InsightSelectors,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D0227506-AA8C-5998-8A62-74769106****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the trail.
	//
	// example:
	//
	// acs:actiontrail:cn-shanghai:159498693826****:trail/trail-name
	TrailArn *string `json:"TrailArn,omitempty" xml:"TrailArn,omitempty"`
}

func (s GetInsightSelectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInsightSelectorsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInsightSelectorsResponseBody) GetInsightSelectors() []*string {
	return s.InsightSelectors
}

func (s *GetInsightSelectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInsightSelectorsResponseBody) GetTrailArn() *string {
	return s.TrailArn
}

func (s *GetInsightSelectorsResponseBody) SetInsightSelectors(v []*string) *GetInsightSelectorsResponseBody {
	s.InsightSelectors = v
	return s
}

func (s *GetInsightSelectorsResponseBody) SetRequestId(v string) *GetInsightSelectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInsightSelectorsResponseBody) SetTrailArn(v string) *GetInsightSelectorsResponseBody {
	s.TrailArn = &v
	return s
}

func (s *GetInsightSelectorsResponseBody) Validate() error {
	return dara.Validate(s)
}
