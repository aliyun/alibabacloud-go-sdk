// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInsightsEventsCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetInsightsEventsCountResponseBodyData) *GetInsightsEventsCountResponseBody
	GetData() []*GetInsightsEventsCountResponseBodyData
	SetNextToken(v string) *GetInsightsEventsCountResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetInsightsEventsCountResponseBody
	GetRequestId() *string
}

type GetInsightsEventsCountResponseBody struct {
	// The information about the Insights events.
	Data []*GetInsightsEventsCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// VjE6bHJlTGoxdm1M****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4ABAEA6E-C740-5CE2-A003-643E5519****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInsightsEventsCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInsightsEventsCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetInsightsEventsCountResponseBody) GetData() []*GetInsightsEventsCountResponseBodyData {
	return s.Data
}

func (s *GetInsightsEventsCountResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetInsightsEventsCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInsightsEventsCountResponseBody) SetData(v []*GetInsightsEventsCountResponseBodyData) *GetInsightsEventsCountResponseBody {
	s.Data = v
	return s
}

func (s *GetInsightsEventsCountResponseBody) SetNextToken(v string) *GetInsightsEventsCountResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetInsightsEventsCountResponseBody) SetRequestId(v string) *GetInsightsEventsCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInsightsEventsCountResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInsightsEventsCountResponseBodyData struct {
	// The number of Insights events.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The type of the Insights event. Valid values:
	//
	// - IpInsight: IP request events.
	//
	// - ApiCallRateInsight: High-risk API call events.
	//
	// - ApiErrorRateInsight: API error events.
	//
	// - AkInsight: AccessKey pair call events.
	//
	// - PolicyChangeInsight: Permission change events.
	//
	// - PasswordChangeInsight: Password change events.
	//
	// - TrailConcealmentInsight: Trail concealment events.
	//
	// example:
	//
	// IpInsight
	InsightType *string `json:"InsightType,omitempty" xml:"InsightType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInsightsEventsCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInsightsEventsCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInsightsEventsCountResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetInsightsEventsCountResponseBodyData) GetInsightType() *string {
	return s.InsightType
}

func (s *GetInsightsEventsCountResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInsightsEventsCountResponseBodyData) SetCount(v int32) *GetInsightsEventsCountResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetInsightsEventsCountResponseBodyData) SetInsightType(v string) *GetInsightsEventsCountResponseBodyData {
	s.InsightType = &v
	return s
}

func (s *GetInsightsEventsCountResponseBodyData) SetRegionId(v string) *GetInsightsEventsCountResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetInsightsEventsCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
