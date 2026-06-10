// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryUsageDurationRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *QueryHistoryUsageDurationRankResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryHistoryUsageDurationRankResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryHistoryUsageDurationRankResponseBody
	GetTotalCount() *int32
	SetUsageDurationList(v []*QueryHistoryUsageDurationRankResponseBodyUsageDurationList) *QueryHistoryUsageDurationRankResponseBody
	GetUsageDurationList() []*QueryHistoryUsageDurationRankResponseBodyUsageDurationList
}

type QueryHistoryUsageDurationRankResponseBody struct {
	// The token used to retrieve the next page of results. If this parameter is not returned, there are no more results.
	//
	// example:
	//
	// AAAAAWvmfbFWy0uSlxZ6pIAKAnuwt1ezsRqxI6hPibm27fMH
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of matching entries.
	//
	// example:
	//
	// 94
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of usage duration metrics.
	UsageDurationList []*QueryHistoryUsageDurationRankResponseBodyUsageDurationList `json:"UsageDurationList,omitempty" xml:"UsageDurationList,omitempty" type:"Repeated"`
}

func (s QueryHistoryUsageDurationRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryUsageDurationRankResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHistoryUsageDurationRankResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryHistoryUsageDurationRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHistoryUsageDurationRankResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryHistoryUsageDurationRankResponseBody) GetUsageDurationList() []*QueryHistoryUsageDurationRankResponseBodyUsageDurationList {
	return s.UsageDurationList
}

func (s *QueryHistoryUsageDurationRankResponseBody) SetNextToken(v string) *QueryHistoryUsageDurationRankResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBody) SetRequestId(v string) *QueryHistoryUsageDurationRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBody) SetTotalCount(v int32) *QueryHistoryUsageDurationRankResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBody) SetUsageDurationList(v []*QueryHistoryUsageDurationRankResponseBodyUsageDurationList) *QueryHistoryUsageDurationRankResponseBody {
	s.UsageDurationList = v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBody) Validate() error {
	if s.UsageDurationList != nil {
		for _, item := range s.UsageDurationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHistoryUsageDurationRankResponseBodyUsageDurationList struct {
	// The **charge type**. Valid values: `POST_PAID` (post-paid), `PRE_PAID` (pre-paid), `MONTH_PACKAGE` (monthly subscription), and `DURATION` (duration package).
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// **The ID of the cloud desktop.**
	//
	// example:
	//
	// ecd-8cndajrdrd424sb99
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// **The name of the cloud desktop.**
	//
	// example:
	//
	// TestName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// **The end user ID.**
	//
	// example:
	//
	// endUserId
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The **usage duration**, in seconds. Note: This is a `Long` value. Ensure your client can handle the precision.
	//
	// example:
	//
	// 20000
	UsageDuration *int64 `json:"UsageDuration,omitempty" xml:"UsageDuration,omitempty"`
}

func (s QueryHistoryUsageDurationRankResponseBodyUsageDurationList) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryUsageDurationRankResponseBodyUsageDurationList) GoString() string {
	return s.String()
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) GetChargeType() *string {
	return s.ChargeType
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) GetDesktopId() *string {
	return s.DesktopId
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) GetDesktopName() *string {
	return s.DesktopName
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) GetEndUserId() *string {
	return s.EndUserId
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) GetUsageDuration() *int64 {
	return s.UsageDuration
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) SetChargeType(v string) *QueryHistoryUsageDurationRankResponseBodyUsageDurationList {
	s.ChargeType = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) SetDesktopId(v string) *QueryHistoryUsageDurationRankResponseBodyUsageDurationList {
	s.DesktopId = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) SetDesktopName(v string) *QueryHistoryUsageDurationRankResponseBodyUsageDurationList {
	s.DesktopName = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) SetEndUserId(v string) *QueryHistoryUsageDurationRankResponseBodyUsageDurationList {
	s.EndUserId = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) SetRegionId(v string) *QueryHistoryUsageDurationRankResponseBodyUsageDurationList {
	s.RegionId = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) SetUsageDuration(v int64) *QueryHistoryUsageDurationRankResponseBodyUsageDurationList {
	s.UsageDuration = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponseBodyUsageDurationList) Validate() error {
	return dara.Validate(s)
}
