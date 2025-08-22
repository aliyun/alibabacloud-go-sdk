// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserSecDropByMinuteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeDcdnUserSecDropByMinuteResponseBody
	GetDescription() *string
	SetLen(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody
	GetLen() *int32
	SetPageNumber(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDcdnUserSecDropByMinuteResponseBody
	GetRequestId() *string
	SetRows(v []*DescribeDcdnUserSecDropByMinuteResponseBodyRows) *DescribeDcdnUserSecDropByMinuteResponseBody
	GetRows() []*DescribeDcdnUserSecDropByMinuteResponseBodyRows
	SetTotalCount(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody
	GetTotalCount() *int32
}

type DescribeDcdnUserSecDropByMinuteResponseBody struct {
	// The description of HTTP responses.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Len *int32 `json:"Len,omitempty" xml:"Len,omitempty"`
	// The number of the returned page.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8143fA8A-B2B2-4915-538D-546B538D25FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array.
	Rows []*DescribeDcdnUserSecDropByMinuteResponseBodyRows `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5738
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnUserSecDropByMinuteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserSecDropByMinuteResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) GetLen() *int32 {
	return s.Len
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) GetRows() []*DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	return s.Rows
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetDescription(v string) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetLen(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.Len = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetPageNumber(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetPageSize(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetRequestId(v string) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetRows(v []*DescribeDcdnUserSecDropByMinuteResponseBodyRows) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.Rows = v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetTotalCount(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserSecDropByMinuteResponseBodyRows struct {
	// The domain name.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of packets blocked within 5 minutes.
	//
	// example:
	//
	// 264
	Drops *int32 `json:"Drops,omitempty" xml:"Drops,omitempty"`
	// The object that triggered rate limiting.
	//
	// example:
	//
	// Normal Mode
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The rule that was triggered.
	//
	// example:
	//
	// Normal Mode
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The security feature that blocked the packets.
	//
	// example:
	//
	// robot
	SecFunc *string `json:"SecFunc,omitempty" xml:"SecFunc,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-01-31T15:04:04Z
	TmStr *string `json:"TmStr,omitempty" xml:"TmStr,omitempty"`
}

func (s DescribeDcdnUserSecDropByMinuteResponseBodyRows) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserSecDropByMinuteResponseBodyRows) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) GetDrops() *int32 {
	return s.Drops
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) GetObject() *string {
	return s.Object
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) GetSecFunc() *string {
	return s.SecFunc
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) GetTmStr() *string {
	return s.TmStr
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetDomain(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetDrops(v int32) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.Drops = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetObject(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.Object = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetRuleName(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetSecFunc(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.SecFunc = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetTmStr(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.TmStr = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) Validate() error {
	return dara.Validate(s)
}
