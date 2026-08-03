// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutInsightSelectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInsightSelectors(v []*string) *PutInsightSelectorsResponseBody
	GetInsightSelectors() []*string
	SetRequestId(v string) *PutInsightSelectorsResponseBody
	GetRequestId() *string
	SetTrailArn(v string) *PutInsightSelectorsResponseBody
	GetTrailArn() *string
}

type PutInsightSelectorsResponseBody struct {
	// An array of Insights event types.
	InsightSelectors []*string `json:"InsightSelectors,omitempty" xml:"InsightSelectors,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7EC26DF0-35AC-5F37-82B3-F5545D0A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the trail.
	//
	// example:
	//
	// acs:actiontrail:cn-shanghai:159498693826****:trail/trail-name
	TrailArn *string `json:"TrailArn,omitempty" xml:"TrailArn,omitempty"`
}

func (s PutInsightSelectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutInsightSelectorsResponseBody) GoString() string {
	return s.String()
}

func (s *PutInsightSelectorsResponseBody) GetInsightSelectors() []*string {
	return s.InsightSelectors
}

func (s *PutInsightSelectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutInsightSelectorsResponseBody) GetTrailArn() *string {
	return s.TrailArn
}

func (s *PutInsightSelectorsResponseBody) SetInsightSelectors(v []*string) *PutInsightSelectorsResponseBody {
	s.InsightSelectors = v
	return s
}

func (s *PutInsightSelectorsResponseBody) SetRequestId(v string) *PutInsightSelectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutInsightSelectorsResponseBody) SetTrailArn(v string) *PutInsightSelectorsResponseBody {
	s.TrailArn = &v
	return s
}

func (s *PutInsightSelectorsResponseBody) Validate() error {
	return dara.Validate(s)
}
