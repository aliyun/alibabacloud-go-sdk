// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchLawQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunSearchLawQueryResponseBody
	GetCode() *string
	SetData(v *RunSearchLawQueryResponseBodyData) *RunSearchLawQueryResponseBody
	GetData() *RunSearchLawQueryResponseBodyData
	SetHttpStatusCode(v int64) *RunSearchLawQueryResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *RunSearchLawQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunSearchLawQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunSearchLawQueryResponseBody
	GetSuccess() *bool
}

type RunSearchLawQueryResponseBody struct {
	// example:
	//
	// Ok
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *RunSearchLawQueryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 05062567-EB51-50F6-AF56-0BE44955848D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RunSearchLawQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryResponseBody) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunSearchLawQueryResponseBody) GetData() *RunSearchLawQueryResponseBodyData {
	return s.Data
}

func (s *RunSearchLawQueryResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *RunSearchLawQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunSearchLawQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSearchLawQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunSearchLawQueryResponseBody) SetCode(v string) *RunSearchLawQueryResponseBody {
	s.Code = &v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetData(v *RunSearchLawQueryResponseBodyData) *RunSearchLawQueryResponseBody {
	s.Data = v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetHttpStatusCode(v int64) *RunSearchLawQueryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetMessage(v string) *RunSearchLawQueryResponseBody {
	s.Message = &v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetRequestId(v string) *RunSearchLawQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetSuccess(v bool) *RunSearchLawQueryResponseBody {
	s.Success = &v
	return s
}

func (s *RunSearchLawQueryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchLawQueryResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int32                                        `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	LawResult   []*RunSearchLawQueryResponseBodyDataLawResult `json:"lawResult,omitempty" xml:"lawResult,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	PageSize            *int32                                                `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query               *string                                               `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywords       []*string                                             `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty" type:"Repeated"`
	SortKeyAndDirection *RunSearchLawQueryResponseBodyDataSortKeyAndDirection `json:"sortKeyAndDirection,omitempty" xml:"sortKeyAndDirection,omitempty" type:"Struct"`
	// example:
	//
	// 0
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s RunSearchLawQueryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *RunSearchLawQueryResponseBodyData) GetLawResult() []*RunSearchLawQueryResponseBodyDataLawResult {
	return s.LawResult
}

func (s *RunSearchLawQueryResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RunSearchLawQueryResponseBodyData) GetQuery() *string {
	return s.Query
}

func (s *RunSearchLawQueryResponseBodyData) GetQueryKeywords() []*string {
	return s.QueryKeywords
}

func (s *RunSearchLawQueryResponseBodyData) GetSortKeyAndDirection() *RunSearchLawQueryResponseBodyDataSortKeyAndDirection {
	return s.SortKeyAndDirection
}

func (s *RunSearchLawQueryResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *RunSearchLawQueryResponseBodyData) SetCurrentPage(v int32) *RunSearchLawQueryResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetLawResult(v []*RunSearchLawQueryResponseBodyDataLawResult) *RunSearchLawQueryResponseBodyData {
	s.LawResult = v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetPageSize(v int32) *RunSearchLawQueryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetQuery(v string) *RunSearchLawQueryResponseBodyData {
	s.Query = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetQueryKeywords(v []*string) *RunSearchLawQueryResponseBodyData {
	s.QueryKeywords = v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetSortKeyAndDirection(v *RunSearchLawQueryResponseBodyDataSortKeyAndDirection) *RunSearchLawQueryResponseBodyData {
	s.SortKeyAndDirection = v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetTotalCount(v int64) *RunSearchLawQueryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) Validate() error {
	if s.LawResult != nil {
		for _, item := range s.LawResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SortKeyAndDirection != nil {
		if err := s.SortKeyAndDirection.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchLawQueryResponseBodyDataLawResult struct {
	LawDomain *RunSearchLawQueryResponseBodyDataLawResultLawDomain `json:"lawDomain,omitempty" xml:"lawDomain,omitempty" type:"Struct"`
	// example:
	//
	// 0.0050
	Similarity *string `json:"similarity,omitempty" xml:"similarity,omitempty"`
}

func (s RunSearchLawQueryResponseBodyDataLawResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryResponseBodyDataLawResult) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBodyDataLawResult) GetLawDomain() *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	return s.LawDomain
}

func (s *RunSearchLawQueryResponseBodyDataLawResult) GetSimilarity() *string {
	return s.Similarity
}

func (s *RunSearchLawQueryResponseBodyDataLawResult) SetLawDomain(v *RunSearchLawQueryResponseBodyDataLawResultLawDomain) *RunSearchLawQueryResponseBodyDataLawResult {
	s.LawDomain = v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResult) SetSimilarity(v string) *RunSearchLawQueryResponseBodyDataLawResult {
	s.Similarity = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResult) Validate() error {
	if s.LawDomain != nil {
		if err := s.LawDomain.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchLawQueryResponseBodyDataLawResultLawDomain struct {
	AbolitionBasis         *string `json:"abolitionBasis,omitempty" xml:"abolitionBasis,omitempty"`
	ImplementYearMonthDate *string `json:"implementYearMonthDate,omitempty" xml:"implementYearMonthDate,omitempty"`
	// example:
	//
	// null
	InvalidBasis *string `json:"invalidBasis,omitempty" xml:"invalidBasis,omitempty"`
	// example:
	//
	// ""
	IssuingNo    *string `json:"issuingNo,omitempty" xml:"issuingNo,omitempty"`
	IssuingOrgan *string `json:"issuingOrgan,omitempty" xml:"issuingOrgan,omitempty"`
	// example:
	//
	// b2274825c8c3bc2343ca73680243ddc8
	LawId *string `json:"lawId,omitempty" xml:"lawId,omitempty"`
	// example:
	//
	// ccc209683be1509676174fd6890f24b8
	LawItemId        *string `json:"lawItemId,omitempty" xml:"lawItemId,omitempty"`
	LawName          *string `json:"lawName,omitempty" xml:"lawName,omitempty"`
	LawOrder         *string `json:"lawOrder,omitempty" xml:"lawOrder,omitempty"`
	LawSourceContent *string `json:"lawSourceContent,omitempty" xml:"lawSourceContent,omitempty"`
	LawTitle         *string `json:"lawTitle,omitempty" xml:"lawTitle,omitempty"`
	// example:
	//
	// "[]"
	ModifyBasis          *string `json:"modifyBasis,omitempty" xml:"modifyBasis,omitempty"`
	PotencyLevel         *string `json:"potencyLevel,omitempty" xml:"potencyLevel,omitempty"`
	ReleaseYearMonthDate *string `json:"releaseYearMonthDate,omitempty" xml:"releaseYearMonthDate,omitempty"`
	// example:
	//
	// null
	ThematicClassify *string `json:"thematicClassify,omitempty" xml:"thematicClassify,omitempty"`
	Timeliness       *string `json:"timeliness,omitempty" xml:"timeliness,omitempty"`
}

func (s RunSearchLawQueryResponseBodyDataLawResultLawDomain) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryResponseBodyDataLawResultLawDomain) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetAbolitionBasis() *string {
	return s.AbolitionBasis
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetImplementYearMonthDate() *string {
	return s.ImplementYearMonthDate
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetInvalidBasis() *string {
	return s.InvalidBasis
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetIssuingNo() *string {
	return s.IssuingNo
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetIssuingOrgan() *string {
	return s.IssuingOrgan
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetLawId() *string {
	return s.LawId
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetLawItemId() *string {
	return s.LawItemId
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetLawName() *string {
	return s.LawName
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetLawOrder() *string {
	return s.LawOrder
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetLawSourceContent() *string {
	return s.LawSourceContent
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetLawTitle() *string {
	return s.LawTitle
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetModifyBasis() *string {
	return s.ModifyBasis
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetPotencyLevel() *string {
	return s.PotencyLevel
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetReleaseYearMonthDate() *string {
	return s.ReleaseYearMonthDate
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetThematicClassify() *string {
	return s.ThematicClassify
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) GetTimeliness() *string {
	return s.Timeliness
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetAbolitionBasis(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.AbolitionBasis = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetImplementYearMonthDate(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.ImplementYearMonthDate = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetInvalidBasis(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.InvalidBasis = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetIssuingNo(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.IssuingNo = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetIssuingOrgan(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.IssuingOrgan = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawId(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawId = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawItemId(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawItemId = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawName(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawName = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawOrder(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawOrder = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawSourceContent(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawSourceContent = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawTitle(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawTitle = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetModifyBasis(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.ModifyBasis = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetPotencyLevel(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.PotencyLevel = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetReleaseYearMonthDate(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.ReleaseYearMonthDate = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetThematicClassify(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.ThematicClassify = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetTimeliness(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.Timeliness = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) Validate() error {
	return dara.Validate(s)
}

type RunSearchLawQueryResponseBodyDataSortKeyAndDirection struct {
	// example:
	//
	// desc
	ReleaseYearMonthDate *string `json:"releaseYearMonthDate,omitempty" xml:"releaseYearMonthDate,omitempty"`
	// example:
	//
	// desc
	Similarity *string `json:"similarity,omitempty" xml:"similarity,omitempty"`
}

func (s RunSearchLawQueryResponseBodyDataSortKeyAndDirection) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryResponseBodyDataSortKeyAndDirection) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBodyDataSortKeyAndDirection) GetReleaseYearMonthDate() *string {
	return s.ReleaseYearMonthDate
}

func (s *RunSearchLawQueryResponseBodyDataSortKeyAndDirection) GetSimilarity() *string {
	return s.Similarity
}

func (s *RunSearchLawQueryResponseBodyDataSortKeyAndDirection) SetReleaseYearMonthDate(v string) *RunSearchLawQueryResponseBodyDataSortKeyAndDirection {
	s.ReleaseYearMonthDate = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataSortKeyAndDirection) SetSimilarity(v string) *RunSearchLawQueryResponseBodyDataSortKeyAndDirection {
	s.Similarity = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataSortKeyAndDirection) Validate() error {
	return dara.Validate(s)
}
