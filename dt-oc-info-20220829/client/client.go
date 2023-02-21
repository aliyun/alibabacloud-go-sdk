// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetOcCompetitorsRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcCompetitorsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcCompetitorsRequest) GoString() string {
	return s.String()
}

func (s *GetOcCompetitorsRequest) SetPageNo(v int32) *GetOcCompetitorsRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcCompetitorsRequest) SetPageSize(v int32) *GetOcCompetitorsRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcCompetitorsRequest) SetSearchKey(v string) *GetOcCompetitorsRequest {
	s.SearchKey = &v
	return s
}

type GetOcCompetitorsResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcCompetitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                              `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                              `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                              `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcCompetitorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcCompetitorsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcCompetitorsResponseBody) SetCode(v string) *GetOcCompetitorsResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcCompetitorsResponseBody) SetData(v []*GetOcCompetitorsResponseBodyData) *GetOcCompetitorsResponseBody {
	s.Data = v
	return s
}

func (s *GetOcCompetitorsResponseBody) SetMessage(v string) *GetOcCompetitorsResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcCompetitorsResponseBody) SetPageIndex(v int32) *GetOcCompetitorsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcCompetitorsResponseBody) SetPageNum(v int32) *GetOcCompetitorsResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcCompetitorsResponseBody) SetRequestId(v string) *GetOcCompetitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcCompetitorsResponseBody) SetSuccess(v bool) *GetOcCompetitorsResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcCompetitorsResponseBody) SetTotalNum(v int32) *GetOcCompetitorsResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcCompetitorsResponseBodyData struct {
	CompetitionBrandIntroduction *string `json:"CompetitionBrandIntroduction,omitempty" xml:"CompetitionBrandIntroduction,omitempty"`
	CompetitionEntAddress        *string `json:"CompetitionEntAddress,omitempty" xml:"CompetitionEntAddress,omitempty"`
	CompetitionEntEsDate         *string `json:"CompetitionEntEsDate,omitempty" xml:"CompetitionEntEsDate,omitempty"`
	CompetitionEntFinTurn        *string `json:"CompetitionEntFinTurn,omitempty" xml:"CompetitionEntFinTurn,omitempty"`
	CompetitionEntName           *string `json:"CompetitionEntName,omitempty" xml:"CompetitionEntName,omitempty"`
	CompetitionIntroduction      *string `json:"CompetitionIntroduction,omitempty" xml:"CompetitionIntroduction,omitempty"`
	CompetitionLogoUrl           *string `json:"CompetitionLogoUrl,omitempty" xml:"CompetitionLogoUrl,omitempty"`
	CompetitionProductName       *string `json:"CompetitionProductName,omitempty" xml:"CompetitionProductName,omitempty"`
	CompetitionTag               *string `json:"CompetitionTag,omitempty" xml:"CompetitionTag,omitempty"`
	CompetitionWebsite           *string `json:"CompetitionWebsite,omitempty" xml:"CompetitionWebsite,omitempty"`
	EntName                      *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
}

func (s GetOcCompetitorsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcCompetitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionBrandIntroduction(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionBrandIntroduction = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionEntAddress(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionEntAddress = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionEntEsDate(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionEntEsDate = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionEntFinTurn(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionEntFinTurn = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionEntName(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionEntName = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionIntroduction(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionIntroduction = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionLogoUrl(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionLogoUrl = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionProductName(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionProductName = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionTag(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionTag = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetCompetitionWebsite(v string) *GetOcCompetitorsResponseBodyData {
	s.CompetitionWebsite = &v
	return s
}

func (s *GetOcCompetitorsResponseBodyData) SetEntName(v string) *GetOcCompetitorsResponseBodyData {
	s.EntName = &v
	return s
}

type GetOcCompetitorsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcCompetitorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcCompetitorsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcCompetitorsResponse) GoString() string {
	return s.String()
}

func (s *GetOcCompetitorsResponse) SetHeaders(v map[string]*string) *GetOcCompetitorsResponse {
	s.Headers = v
	return s
}

func (s *GetOcCompetitorsResponse) SetStatusCode(v int32) *GetOcCompetitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcCompetitorsResponse) SetBody(v *GetOcCompetitorsResponseBody) *GetOcCompetitorsResponse {
	s.Body = v
	return s
}

type GetOcCoreTeamsRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcCoreTeamsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcCoreTeamsRequest) GoString() string {
	return s.String()
}

func (s *GetOcCoreTeamsRequest) SetPageNo(v int32) *GetOcCoreTeamsRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcCoreTeamsRequest) SetPageSize(v int32) *GetOcCoreTeamsRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcCoreTeamsRequest) SetSearchKey(v string) *GetOcCoreTeamsRequest {
	s.SearchKey = &v
	return s
}

type GetOcCoreTeamsResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcCoreTeamsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                            `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcCoreTeamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcCoreTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcCoreTeamsResponseBody) SetCode(v string) *GetOcCoreTeamsResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcCoreTeamsResponseBody) SetData(v []*GetOcCoreTeamsResponseBodyData) *GetOcCoreTeamsResponseBody {
	s.Data = v
	return s
}

func (s *GetOcCoreTeamsResponseBody) SetMessage(v string) *GetOcCoreTeamsResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcCoreTeamsResponseBody) SetPageIndex(v int32) *GetOcCoreTeamsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcCoreTeamsResponseBody) SetPageNum(v int32) *GetOcCoreTeamsResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcCoreTeamsResponseBody) SetRequestId(v string) *GetOcCoreTeamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcCoreTeamsResponseBody) SetSuccess(v bool) *GetOcCoreTeamsResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcCoreTeamsResponseBody) SetTotalNum(v int32) *GetOcCoreTeamsResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcCoreTeamsResponseBodyData struct {
	EntName            *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	MemberIntroduction *string `json:"MemberIntroduction,omitempty" xml:"MemberIntroduction,omitempty"`
	MemberName         *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberPosition     *string `json:"MemberPosition,omitempty" xml:"MemberPosition,omitempty"`
}

func (s GetOcCoreTeamsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcCoreTeamsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcCoreTeamsResponseBodyData) SetEntName(v string) *GetOcCoreTeamsResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcCoreTeamsResponseBodyData) SetMemberIntroduction(v string) *GetOcCoreTeamsResponseBodyData {
	s.MemberIntroduction = &v
	return s
}

func (s *GetOcCoreTeamsResponseBodyData) SetMemberName(v string) *GetOcCoreTeamsResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *GetOcCoreTeamsResponseBodyData) SetMemberPosition(v string) *GetOcCoreTeamsResponseBodyData {
	s.MemberPosition = &v
	return s
}

type GetOcCoreTeamsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcCoreTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcCoreTeamsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcCoreTeamsResponse) GoString() string {
	return s.String()
}

func (s *GetOcCoreTeamsResponse) SetHeaders(v map[string]*string) *GetOcCoreTeamsResponse {
	s.Headers = v
	return s
}

func (s *GetOcCoreTeamsResponse) SetStatusCode(v int32) *GetOcCoreTeamsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcCoreTeamsResponse) SetBody(v *GetOcCoreTeamsResponseBody) *GetOcCoreTeamsResponse {
	s.Body = v
	return s
}

type GetOcFinancingRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcFinancingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcFinancingRequest) GoString() string {
	return s.String()
}

func (s *GetOcFinancingRequest) SetPageNo(v int32) *GetOcFinancingRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcFinancingRequest) SetPageSize(v int32) *GetOcFinancingRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcFinancingRequest) SetSearchKey(v string) *GetOcFinancingRequest {
	s.SearchKey = &v
	return s
}

type GetOcFinancingResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcFinancingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                            `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcFinancingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcFinancingResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcFinancingResponseBody) SetCode(v string) *GetOcFinancingResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcFinancingResponseBody) SetData(v []*GetOcFinancingResponseBodyData) *GetOcFinancingResponseBody {
	s.Data = v
	return s
}

func (s *GetOcFinancingResponseBody) SetMessage(v string) *GetOcFinancingResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcFinancingResponseBody) SetPageIndex(v int32) *GetOcFinancingResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcFinancingResponseBody) SetPageNum(v int32) *GetOcFinancingResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcFinancingResponseBody) SetRequestId(v string) *GetOcFinancingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcFinancingResponseBody) SetSuccess(v bool) *GetOcFinancingResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcFinancingResponseBody) SetTotalNum(v int32) *GetOcFinancingResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcFinancingResponseBodyData struct {
	EntName   *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	FinAmount *string `json:"FinAmount,omitempty" xml:"FinAmount,omitempty"`
	FinDate   *string `json:"FinDate,omitempty" xml:"FinDate,omitempty"`
	FinTurn   *string `json:"FinTurn,omitempty" xml:"FinTurn,omitempty"`
	Investors *string `json:"Investors,omitempty" xml:"Investors,omitempty"`
}

func (s GetOcFinancingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcFinancingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcFinancingResponseBodyData) SetEntName(v string) *GetOcFinancingResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcFinancingResponseBodyData) SetFinAmount(v string) *GetOcFinancingResponseBodyData {
	s.FinAmount = &v
	return s
}

func (s *GetOcFinancingResponseBodyData) SetFinDate(v string) *GetOcFinancingResponseBodyData {
	s.FinDate = &v
	return s
}

func (s *GetOcFinancingResponseBodyData) SetFinTurn(v string) *GetOcFinancingResponseBodyData {
	s.FinTurn = &v
	return s
}

func (s *GetOcFinancingResponseBodyData) SetInvestors(v string) *GetOcFinancingResponseBodyData {
	s.Investors = &v
	return s
}

type GetOcFinancingResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcFinancingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcFinancingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcFinancingResponse) GoString() string {
	return s.String()
}

func (s *GetOcFinancingResponse) SetHeaders(v map[string]*string) *GetOcFinancingResponse {
	s.Headers = v
	return s
}

func (s *GetOcFinancingResponse) SetStatusCode(v int32) *GetOcFinancingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcFinancingResponse) SetBody(v *GetOcFinancingResponseBody) *GetOcFinancingResponse {
	s.Body = v
	return s
}

type GetOcFuzzSearchRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcFuzzSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcFuzzSearchRequest) GoString() string {
	return s.String()
}

func (s *GetOcFuzzSearchRequest) SetPageNo(v int32) *GetOcFuzzSearchRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcFuzzSearchRequest) SetPageSize(v int32) *GetOcFuzzSearchRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcFuzzSearchRequest) SetSearchKey(v string) *GetOcFuzzSearchRequest {
	s.SearchKey = &v
	return s
}

type GetOcFuzzSearchResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcFuzzSearchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                             `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                             `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcFuzzSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcFuzzSearchResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcFuzzSearchResponseBody) SetCode(v string) *GetOcFuzzSearchResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcFuzzSearchResponseBody) SetData(v []*GetOcFuzzSearchResponseBodyData) *GetOcFuzzSearchResponseBody {
	s.Data = v
	return s
}

func (s *GetOcFuzzSearchResponseBody) SetMessage(v string) *GetOcFuzzSearchResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcFuzzSearchResponseBody) SetPageIndex(v int32) *GetOcFuzzSearchResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcFuzzSearchResponseBody) SetPageNum(v int32) *GetOcFuzzSearchResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcFuzzSearchResponseBody) SetRequestId(v string) *GetOcFuzzSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcFuzzSearchResponseBody) SetSuccess(v bool) *GetOcFuzzSearchResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcFuzzSearchResponseBody) SetTotalNum(v int32) *GetOcFuzzSearchResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcFuzzSearchResponseBodyData struct {
	EntName *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
}

func (s GetOcFuzzSearchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcFuzzSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcFuzzSearchResponseBodyData) SetEntName(v string) *GetOcFuzzSearchResponseBodyData {
	s.EntName = &v
	return s
}

type GetOcFuzzSearchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcFuzzSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcFuzzSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcFuzzSearchResponse) GoString() string {
	return s.String()
}

func (s *GetOcFuzzSearchResponse) SetHeaders(v map[string]*string) *GetOcFuzzSearchResponse {
	s.Headers = v
	return s
}

func (s *GetOcFuzzSearchResponse) SetStatusCode(v int32) *GetOcFuzzSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcFuzzSearchResponse) SetBody(v *GetOcFuzzSearchResponseBody) *GetOcFuzzSearchResponse {
	s.Body = v
	return s
}

type GetOcIcAbnormalOperationRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcAbnormalOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcAbnormalOperationRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcAbnormalOperationRequest) SetPageNo(v int32) *GetOcIcAbnormalOperationRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcAbnormalOperationRequest) SetPageSize(v int32) *GetOcIcAbnormalOperationRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcAbnormalOperationRequest) SetSearchKey(v string) *GetOcIcAbnormalOperationRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcAbnormalOperationResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcAbnormalOperationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                      `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                      `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                      `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcAbnormalOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcAbnormalOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcAbnormalOperationResponseBody) SetCode(v string) *GetOcIcAbnormalOperationResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBody) SetData(v []*GetOcIcAbnormalOperationResponseBodyData) *GetOcIcAbnormalOperationResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBody) SetMessage(v string) *GetOcIcAbnormalOperationResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBody) SetPageIndex(v int32) *GetOcIcAbnormalOperationResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBody) SetPageNum(v int32) *GetOcIcAbnormalOperationResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBody) SetRequestId(v string) *GetOcIcAbnormalOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBody) SetSuccess(v bool) *GetOcIcAbnormalOperationResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBody) SetTotalNum(v int32) *GetOcIcAbnormalOperationResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcAbnormalOperationResponseBodyData struct {
	InDate        *string `json:"InDate,omitempty" xml:"InDate,omitempty"`
	InDepartment  *string `json:"InDepartment,omitempty" xml:"InDepartment,omitempty"`
	InReason      *string `json:"InReason,omitempty" xml:"InReason,omitempty"`
	OutDate       *string `json:"OutDate,omitempty" xml:"OutDate,omitempty"`
	OutDepartment *string `json:"OutDepartment,omitempty" xml:"OutDepartment,omitempty"`
	OutReason     *string `json:"OutReason,omitempty" xml:"OutReason,omitempty"`
}

func (s GetOcIcAbnormalOperationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcAbnormalOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcAbnormalOperationResponseBodyData) SetInDate(v string) *GetOcIcAbnormalOperationResponseBodyData {
	s.InDate = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBodyData) SetInDepartment(v string) *GetOcIcAbnormalOperationResponseBodyData {
	s.InDepartment = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBodyData) SetInReason(v string) *GetOcIcAbnormalOperationResponseBodyData {
	s.InReason = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBodyData) SetOutDate(v string) *GetOcIcAbnormalOperationResponseBodyData {
	s.OutDate = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBodyData) SetOutDepartment(v string) *GetOcIcAbnormalOperationResponseBodyData {
	s.OutDepartment = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponseBodyData) SetOutReason(v string) *GetOcIcAbnormalOperationResponseBodyData {
	s.OutReason = &v
	return s
}

type GetOcIcAbnormalOperationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcAbnormalOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcAbnormalOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcAbnormalOperationResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcAbnormalOperationResponse) SetHeaders(v map[string]*string) *GetOcIcAbnormalOperationResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcAbnormalOperationResponse) SetStatusCode(v int32) *GetOcIcAbnormalOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcAbnormalOperationResponse) SetBody(v *GetOcIcAbnormalOperationResponseBody) *GetOcIcAbnormalOperationResponse {
	s.Body = v
	return s
}

type GetOcIcAdminLicenseRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcAdminLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcAdminLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcAdminLicenseRequest) SetPageNo(v int32) *GetOcIcAdminLicenseRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcAdminLicenseRequest) SetPageSize(v int32) *GetOcIcAdminLicenseRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcAdminLicenseRequest) SetSearchKey(v string) *GetOcIcAdminLicenseRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcAdminLicenseResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcAdminLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcAdminLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcAdminLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcAdminLicenseResponseBody) SetCode(v string) *GetOcIcAdminLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBody) SetData(v []*GetOcIcAdminLicenseResponseBodyData) *GetOcIcAdminLicenseResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcAdminLicenseResponseBody) SetMessage(v string) *GetOcIcAdminLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBody) SetPageIndex(v int32) *GetOcIcAdminLicenseResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBody) SetPageNum(v int32) *GetOcIcAdminLicenseResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBody) SetRequestId(v string) *GetOcIcAdminLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBody) SetSuccess(v bool) *GetOcIcAdminLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBody) SetTotalNum(v int32) *GetOcIcAdminLicenseResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcAdminLicenseResponseBodyData struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	LicenseName *string `json:"LicenseName,omitempty" xml:"LicenseName,omitempty"`
	LicenseNo   *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetOcIcAdminLicenseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcAdminLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcAdminLicenseResponseBodyData) SetContent(v string) *GetOcIcAdminLicenseResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBodyData) SetDepartment(v string) *GetOcIcAdminLicenseResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBodyData) SetEndDate(v string) *GetOcIcAdminLicenseResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBodyData) SetLicenseName(v string) *GetOcIcAdminLicenseResponseBodyData {
	s.LicenseName = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBodyData) SetLicenseNo(v string) *GetOcIcAdminLicenseResponseBodyData {
	s.LicenseNo = &v
	return s
}

func (s *GetOcIcAdminLicenseResponseBodyData) SetStartDate(v string) *GetOcIcAdminLicenseResponseBodyData {
	s.StartDate = &v
	return s
}

type GetOcIcAdminLicenseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcAdminLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcAdminLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcAdminLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcAdminLicenseResponse) SetHeaders(v map[string]*string) *GetOcIcAdminLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcAdminLicenseResponse) SetStatusCode(v int32) *GetOcIcAdminLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcAdminLicenseResponse) SetBody(v *GetOcIcAdminLicenseResponseBody) *GetOcIcAdminLicenseResponse {
	s.Body = v
	return s
}

type GetOcIcBasicRequest struct {
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcBasicRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcBasicRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcBasicRequest) SetSearchKey(v string) *GetOcIcBasicRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcBasicResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetOcIcBasicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                        `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcBasicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcBasicResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcBasicResponseBody) SetCode(v string) *GetOcIcBasicResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcBasicResponseBody) SetData(v *GetOcIcBasicResponseBodyData) *GetOcIcBasicResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcBasicResponseBody) SetMessage(v string) *GetOcIcBasicResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcBasicResponseBody) SetPageIndex(v int32) *GetOcIcBasicResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcBasicResponseBody) SetPageNum(v int32) *GetOcIcBasicResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcBasicResponseBody) SetRequestId(v string) *GetOcIcBasicResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcBasicResponseBody) SetSuccess(v bool) *GetOcIcBasicResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcBasicResponseBody) SetTotalNum(v int32) *GetOcIcBasicResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcBasicResponseBodyData struct {
	CheckDate        *string `json:"CheckDate,omitempty" xml:"CheckDate,omitempty"`
	EntAddress       *string `json:"EntAddress,omitempty" xml:"EntAddress,omitempty"`
	EntBrief         *string `json:"EntBrief,omitempty" xml:"EntBrief,omitempty"`
	EntName          *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	EntNameEng       *string `json:"EntNameEng,omitempty" xml:"EntNameEng,omitempty"`
	EntStatus        *string `json:"EntStatus,omitempty" xml:"EntStatus,omitempty"`
	EntType          *string `json:"EntType,omitempty" xml:"EntType,omitempty"`
	EsDate           *string `json:"EsDate,omitempty" xml:"EsDate,omitempty"`
	FormerNames      *string `json:"FormerNames,omitempty" xml:"FormerNames,omitempty"`
	IndustryNameLv1  *string `json:"IndustryNameLv1,omitempty" xml:"IndustryNameLv1,omitempty"`
	IndustryNameLv2  *string `json:"IndustryNameLv2,omitempty" xml:"IndustryNameLv2,omitempty"`
	InsuredNum       *string `json:"InsuredNum,omitempty" xml:"InsuredNum,omitempty"`
	LegalName        *string `json:"LegalName,omitempty" xml:"LegalName,omitempty"`
	LicenseNumber    *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	OpFrom           *string `json:"OpFrom,omitempty" xml:"OpFrom,omitempty"`
	OpScope          *string `json:"OpScope,omitempty" xml:"OpScope,omitempty"`
	OpTo             *string `json:"OpTo,omitempty" xml:"OpTo,omitempty"`
	OrgNo            *string `json:"OrgNo,omitempty" xml:"OrgNo,omitempty"`
	RecCap           *string `json:"RecCap,omitempty" xml:"RecCap,omitempty"`
	RegCap           *string `json:"RegCap,omitempty" xml:"RegCap,omitempty"`
	RegOrg           *string `json:"RegOrg,omitempty" xml:"RegOrg,omitempty"`
	RegOrgCity       *string `json:"RegOrgCity,omitempty" xml:"RegOrgCity,omitempty"`
	RegOrgDistrict   *string `json:"RegOrgDistrict,omitempty" xml:"RegOrgDistrict,omitempty"`
	RegOrgProvince   *string `json:"RegOrgProvince,omitempty" xml:"RegOrgProvince,omitempty"`
	SocialCreditCode *string `json:"SocialCreditCode,omitempty" xml:"SocialCreditCode,omitempty"`
	StaffNum         *string `json:"StaffNum,omitempty" xml:"StaffNum,omitempty"`
	TaxNum           *string `json:"TaxNum,omitempty" xml:"TaxNum,omitempty"`
}

func (s GetOcIcBasicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcBasicResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcBasicResponseBodyData) SetCheckDate(v string) *GetOcIcBasicResponseBodyData {
	s.CheckDate = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetEntAddress(v string) *GetOcIcBasicResponseBodyData {
	s.EntAddress = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetEntBrief(v string) *GetOcIcBasicResponseBodyData {
	s.EntBrief = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetEntName(v string) *GetOcIcBasicResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetEntNameEng(v string) *GetOcIcBasicResponseBodyData {
	s.EntNameEng = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetEntStatus(v string) *GetOcIcBasicResponseBodyData {
	s.EntStatus = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetEntType(v string) *GetOcIcBasicResponseBodyData {
	s.EntType = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetEsDate(v string) *GetOcIcBasicResponseBodyData {
	s.EsDate = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetFormerNames(v string) *GetOcIcBasicResponseBodyData {
	s.FormerNames = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetIndustryNameLv1(v string) *GetOcIcBasicResponseBodyData {
	s.IndustryNameLv1 = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetIndustryNameLv2(v string) *GetOcIcBasicResponseBodyData {
	s.IndustryNameLv2 = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetInsuredNum(v string) *GetOcIcBasicResponseBodyData {
	s.InsuredNum = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetLegalName(v string) *GetOcIcBasicResponseBodyData {
	s.LegalName = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetLicenseNumber(v string) *GetOcIcBasicResponseBodyData {
	s.LicenseNumber = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetOpFrom(v string) *GetOcIcBasicResponseBodyData {
	s.OpFrom = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetOpScope(v string) *GetOcIcBasicResponseBodyData {
	s.OpScope = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetOpTo(v string) *GetOcIcBasicResponseBodyData {
	s.OpTo = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetOrgNo(v string) *GetOcIcBasicResponseBodyData {
	s.OrgNo = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetRecCap(v string) *GetOcIcBasicResponseBodyData {
	s.RecCap = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetRegCap(v string) *GetOcIcBasicResponseBodyData {
	s.RegCap = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetRegOrg(v string) *GetOcIcBasicResponseBodyData {
	s.RegOrg = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetRegOrgCity(v string) *GetOcIcBasicResponseBodyData {
	s.RegOrgCity = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetRegOrgDistrict(v string) *GetOcIcBasicResponseBodyData {
	s.RegOrgDistrict = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetRegOrgProvince(v string) *GetOcIcBasicResponseBodyData {
	s.RegOrgProvince = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetSocialCreditCode(v string) *GetOcIcBasicResponseBodyData {
	s.SocialCreditCode = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetStaffNum(v string) *GetOcIcBasicResponseBodyData {
	s.StaffNum = &v
	return s
}

func (s *GetOcIcBasicResponseBodyData) SetTaxNum(v string) *GetOcIcBasicResponseBodyData {
	s.TaxNum = &v
	return s
}

type GetOcIcBasicResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcBasicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcBasicResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcBasicResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcBasicResponse) SetHeaders(v map[string]*string) *GetOcIcBasicResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcBasicResponse) SetStatusCode(v int32) *GetOcIcBasicResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcBasicResponse) SetBody(v *GetOcIcBasicResponseBody) *GetOcIcBasicResponse {
	s.Body = v
	return s
}

type GetOcIcBranchRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcBranchRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcBranchRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcBranchRequest) SetPageNo(v int32) *GetOcIcBranchRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcBranchRequest) SetPageSize(v int32) *GetOcIcBranchRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcBranchRequest) SetSearchKey(v string) *GetOcIcBranchRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcBranchResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcBranchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                           `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                           `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcBranchResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcBranchResponseBody) SetCode(v string) *GetOcIcBranchResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcBranchResponseBody) SetData(v []*GetOcIcBranchResponseBodyData) *GetOcIcBranchResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcBranchResponseBody) SetMessage(v string) *GetOcIcBranchResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcBranchResponseBody) SetPageIndex(v int32) *GetOcIcBranchResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcBranchResponseBody) SetPageNum(v int32) *GetOcIcBranchResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcBranchResponseBody) SetRequestId(v string) *GetOcIcBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcBranchResponseBody) SetSuccess(v bool) *GetOcIcBranchResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcBranchResponseBody) SetTotalNum(v int32) *GetOcIcBranchResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcBranchResponseBodyData struct {
	EntName   *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	EntStatus *string `json:"EntStatus,omitempty" xml:"EntStatus,omitempty"`
	EsDate    *string `json:"EsDate,omitempty" xml:"EsDate,omitempty"`
	OperName  *string `json:"OperName,omitempty" xml:"OperName,omitempty"`
}

func (s GetOcIcBranchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcBranchResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcBranchResponseBodyData) SetEntName(v string) *GetOcIcBranchResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIcBranchResponseBodyData) SetEntStatus(v string) *GetOcIcBranchResponseBodyData {
	s.EntStatus = &v
	return s
}

func (s *GetOcIcBranchResponseBodyData) SetEsDate(v string) *GetOcIcBranchResponseBodyData {
	s.EsDate = &v
	return s
}

func (s *GetOcIcBranchResponseBodyData) SetOperName(v string) *GetOcIcBranchResponseBodyData {
	s.OperName = &v
	return s
}

type GetOcIcBranchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcBranchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcBranchResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcBranchResponse) SetHeaders(v map[string]*string) *GetOcIcBranchResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcBranchResponse) SetStatusCode(v int32) *GetOcIcBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcBranchResponse) SetBody(v *GetOcIcBranchResponseBody) *GetOcIcBranchResponse {
	s.Body = v
	return s
}

type GetOcIcChangeRecordRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcChangeRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcChangeRecordRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcChangeRecordRequest) SetPageNo(v int32) *GetOcIcChangeRecordRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcChangeRecordRequest) SetPageSize(v int32) *GetOcIcChangeRecordRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcChangeRecordRequest) SetSearchKey(v string) *GetOcIcChangeRecordRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcChangeRecordResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcChangeRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcChangeRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcChangeRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcChangeRecordResponseBody) SetCode(v string) *GetOcIcChangeRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcChangeRecordResponseBody) SetData(v []*GetOcIcChangeRecordResponseBodyData) *GetOcIcChangeRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcChangeRecordResponseBody) SetMessage(v string) *GetOcIcChangeRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcChangeRecordResponseBody) SetPageIndex(v int32) *GetOcIcChangeRecordResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcChangeRecordResponseBody) SetPageNum(v int32) *GetOcIcChangeRecordResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcChangeRecordResponseBody) SetRequestId(v string) *GetOcIcChangeRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcChangeRecordResponseBody) SetSuccess(v bool) *GetOcIcChangeRecordResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcChangeRecordResponseBody) SetTotalNum(v int32) *GetOcIcChangeRecordResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcChangeRecordResponseBodyData struct {
	AfterContent  *string `json:"AfterContent,omitempty" xml:"AfterContent,omitempty"`
	BeforeContent *string `json:"BeforeContent,omitempty" xml:"BeforeContent,omitempty"`
	ChangeDate    *string `json:"ChangeDate,omitempty" xml:"ChangeDate,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetOcIcChangeRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcChangeRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcChangeRecordResponseBodyData) SetAfterContent(v string) *GetOcIcChangeRecordResponseBodyData {
	s.AfterContent = &v
	return s
}

func (s *GetOcIcChangeRecordResponseBodyData) SetBeforeContent(v string) *GetOcIcChangeRecordResponseBodyData {
	s.BeforeContent = &v
	return s
}

func (s *GetOcIcChangeRecordResponseBodyData) SetChangeDate(v string) *GetOcIcChangeRecordResponseBodyData {
	s.ChangeDate = &v
	return s
}

func (s *GetOcIcChangeRecordResponseBodyData) SetType(v string) *GetOcIcChangeRecordResponseBodyData {
	s.Type = &v
	return s
}

type GetOcIcChangeRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcChangeRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcChangeRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcChangeRecordResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcChangeRecordResponse) SetHeaders(v map[string]*string) *GetOcIcChangeRecordResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcChangeRecordResponse) SetStatusCode(v int32) *GetOcIcChangeRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcChangeRecordResponse) SetBody(v *GetOcIcChangeRecordResponseBody) *GetOcIcChangeRecordResponse {
	s.Body = v
	return s
}

type GetOcIcCheckupRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcCheckupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcCheckupRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcCheckupRequest) SetPageNo(v int32) *GetOcIcCheckupRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcCheckupRequest) SetPageSize(v int32) *GetOcIcCheckupRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcCheckupRequest) SetSearchKey(v string) *GetOcIcCheckupRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcCheckupResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcCheckupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                            `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcCheckupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcCheckupResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcCheckupResponseBody) SetCode(v string) *GetOcIcCheckupResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcCheckupResponseBody) SetData(v []*GetOcIcCheckupResponseBodyData) *GetOcIcCheckupResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcCheckupResponseBody) SetMessage(v string) *GetOcIcCheckupResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcCheckupResponseBody) SetPageIndex(v int32) *GetOcIcCheckupResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcCheckupResponseBody) SetPageNum(v int32) *GetOcIcCheckupResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcCheckupResponseBody) SetRequestId(v string) *GetOcIcCheckupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcCheckupResponseBody) SetSuccess(v bool) *GetOcIcCheckupResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcCheckupResponseBody) SetTotalNum(v int32) *GetOcIcCheckupResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcCheckupResponseBodyData struct {
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty"`
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	Result     *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetOcIcCheckupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcCheckupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcCheckupResponseBodyData) SetDate(v string) *GetOcIcCheckupResponseBodyData {
	s.Date = &v
	return s
}

func (s *GetOcIcCheckupResponseBodyData) SetDepartment(v string) *GetOcIcCheckupResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcIcCheckupResponseBodyData) SetResult(v string) *GetOcIcCheckupResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetOcIcCheckupResponseBodyData) SetType(v string) *GetOcIcCheckupResponseBodyData {
	s.Type = &v
	return s
}

type GetOcIcCheckupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcCheckupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcCheckupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcCheckupResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcCheckupResponse) SetHeaders(v map[string]*string) *GetOcIcCheckupResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcCheckupResponse) SetStatusCode(v int32) *GetOcIcCheckupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcCheckupResponse) SetBody(v *GetOcIcCheckupResponseBody) *GetOcIcCheckupResponse {
	s.Body = v
	return s
}

type GetOcIcClearAccountRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcClearAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcClearAccountRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcClearAccountRequest) SetPageNo(v int32) *GetOcIcClearAccountRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcClearAccountRequest) SetPageSize(v int32) *GetOcIcClearAccountRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcClearAccountRequest) SetSearchKey(v string) *GetOcIcClearAccountRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcClearAccountResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcClearAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcClearAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcClearAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcClearAccountResponseBody) SetCode(v string) *GetOcIcClearAccountResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcClearAccountResponseBody) SetData(v []*GetOcIcClearAccountResponseBodyData) *GetOcIcClearAccountResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcClearAccountResponseBody) SetMessage(v string) *GetOcIcClearAccountResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcClearAccountResponseBody) SetPageIndex(v int32) *GetOcIcClearAccountResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcClearAccountResponseBody) SetPageNum(v int32) *GetOcIcClearAccountResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcClearAccountResponseBody) SetRequestId(v string) *GetOcIcClearAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcClearAccountResponseBody) SetSuccess(v bool) *GetOcIcClearAccountResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcClearAccountResponseBody) SetTotalNum(v int32) *GetOcIcClearAccountResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcClearAccountResponseBodyData struct {
	Leader *string `json:"Leader,omitempty" xml:"Leader,omitempty"`
	Member *string `json:"Member,omitempty" xml:"Member,omitempty"`
}

func (s GetOcIcClearAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcClearAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcClearAccountResponseBodyData) SetLeader(v string) *GetOcIcClearAccountResponseBodyData {
	s.Leader = &v
	return s
}

func (s *GetOcIcClearAccountResponseBodyData) SetMember(v string) *GetOcIcClearAccountResponseBodyData {
	s.Member = &v
	return s
}

type GetOcIcClearAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcClearAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcClearAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcClearAccountResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcClearAccountResponse) SetHeaders(v map[string]*string) *GetOcIcClearAccountResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcClearAccountResponse) SetStatusCode(v int32) *GetOcIcClearAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcClearAccountResponse) SetBody(v *GetOcIcClearAccountResponseBody) *GetOcIcClearAccountResponse {
	s.Body = v
	return s
}

type GetOcIcDoubleCheckupRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcDoubleCheckupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcDoubleCheckupRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcDoubleCheckupRequest) SetPageNo(v int32) *GetOcIcDoubleCheckupRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcDoubleCheckupRequest) SetPageSize(v int32) *GetOcIcDoubleCheckupRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcDoubleCheckupRequest) SetSearchKey(v string) *GetOcIcDoubleCheckupRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcDoubleCheckupResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcDoubleCheckupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                  `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcDoubleCheckupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcDoubleCheckupResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcDoubleCheckupResponseBody) SetCode(v string) *GetOcIcDoubleCheckupResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBody) SetData(v []*GetOcIcDoubleCheckupResponseBodyData) *GetOcIcDoubleCheckupResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBody) SetMessage(v string) *GetOcIcDoubleCheckupResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBody) SetPageIndex(v int32) *GetOcIcDoubleCheckupResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBody) SetPageNum(v int32) *GetOcIcDoubleCheckupResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBody) SetRequestId(v string) *GetOcIcDoubleCheckupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBody) SetSuccess(v bool) *GetOcIcDoubleCheckupResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBody) SetTotalNum(v int32) *GetOcIcDoubleCheckupResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcDoubleCheckupResponseBodyData struct {
	InspectDate       *string `json:"InspectDate,omitempty" xml:"InspectDate,omitempty"`
	InspectDepartment *string `json:"InspectDepartment,omitempty" xml:"InspectDepartment,omitempty"`
	InspectItem       *string `json:"InspectItem,omitempty" xml:"InspectItem,omitempty"`
	InspectPlanId     *string `json:"InspectPlanId,omitempty" xml:"InspectPlanId,omitempty"`
	InspectPlanName   *string `json:"InspectPlanName,omitempty" xml:"InspectPlanName,omitempty"`
	InspectResult     *string `json:"InspectResult,omitempty" xml:"InspectResult,omitempty"`
	InspectTaskId     *string `json:"InspectTaskId,omitempty" xml:"InspectTaskId,omitempty"`
	InspectTaskName   *string `json:"InspectTaskName,omitempty" xml:"InspectTaskName,omitempty"`
	InspectTypeName   *string `json:"InspectTypeName,omitempty" xml:"InspectTypeName,omitempty"`
}

func (s GetOcIcDoubleCheckupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcDoubleCheckupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcDoubleCheckupResponseBodyData) SetInspectDate(v string) *GetOcIcDoubleCheckupResponseBodyData {
	s.InspectDate = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBodyData) SetInspectDepartment(v string) *GetOcIcDoubleCheckupResponseBodyData {
	s.InspectDepartment = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBodyData) SetInspectItem(v string) *GetOcIcDoubleCheckupResponseBodyData {
	s.InspectItem = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBodyData) SetInspectPlanId(v string) *GetOcIcDoubleCheckupResponseBodyData {
	s.InspectPlanId = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBodyData) SetInspectPlanName(v string) *GetOcIcDoubleCheckupResponseBodyData {
	s.InspectPlanName = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBodyData) SetInspectResult(v string) *GetOcIcDoubleCheckupResponseBodyData {
	s.InspectResult = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBodyData) SetInspectTaskId(v string) *GetOcIcDoubleCheckupResponseBodyData {
	s.InspectTaskId = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBodyData) SetInspectTaskName(v string) *GetOcIcDoubleCheckupResponseBodyData {
	s.InspectTaskName = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponseBodyData) SetInspectTypeName(v string) *GetOcIcDoubleCheckupResponseBodyData {
	s.InspectTypeName = &v
	return s
}

type GetOcIcDoubleCheckupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcDoubleCheckupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcDoubleCheckupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcDoubleCheckupResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcDoubleCheckupResponse) SetHeaders(v map[string]*string) *GetOcIcDoubleCheckupResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcDoubleCheckupResponse) SetStatusCode(v int32) *GetOcIcDoubleCheckupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcDoubleCheckupResponse) SetBody(v *GetOcIcDoubleCheckupResponseBody) *GetOcIcDoubleCheckupResponse {
	s.Body = v
	return s
}

type GetOcIcEmployeeRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcEmployeeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEmployeeRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcEmployeeRequest) SetPageNo(v int32) *GetOcIcEmployeeRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcEmployeeRequest) SetPageSize(v int32) *GetOcIcEmployeeRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcEmployeeRequest) SetRequestId(v string) *GetOcIcEmployeeRequest {
	s.RequestId = &v
	return s
}

func (s *GetOcIcEmployeeRequest) SetSearchKey(v string) *GetOcIcEmployeeRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcEmployeeResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcEmployeeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                             `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                             `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcEmployeeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEmployeeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcEmployeeResponseBody) SetCode(v string) *GetOcIcEmployeeResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcEmployeeResponseBody) SetData(v []*GetOcIcEmployeeResponseBodyData) *GetOcIcEmployeeResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcEmployeeResponseBody) SetMessage(v string) *GetOcIcEmployeeResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcEmployeeResponseBody) SetPageIndex(v int32) *GetOcIcEmployeeResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcEmployeeResponseBody) SetPageNum(v int32) *GetOcIcEmployeeResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcEmployeeResponseBody) SetRequestId(v string) *GetOcIcEmployeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcEmployeeResponseBody) SetSuccess(v bool) *GetOcIcEmployeeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcEmployeeResponseBody) SetTotalNum(v int32) *GetOcIcEmployeeResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcEmployeeResponseBodyData struct {
	JobTitle *string `json:"JobTitle,omitempty" xml:"JobTitle,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOcIcEmployeeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEmployeeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcEmployeeResponseBodyData) SetJobTitle(v string) *GetOcIcEmployeeResponseBodyData {
	s.JobTitle = &v
	return s
}

func (s *GetOcIcEmployeeResponseBodyData) SetName(v string) *GetOcIcEmployeeResponseBodyData {
	s.Name = &v
	return s
}

type GetOcIcEmployeeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcEmployeeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcEmployeeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEmployeeResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcEmployeeResponse) SetHeaders(v map[string]*string) *GetOcIcEmployeeResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcEmployeeResponse) SetStatusCode(v int32) *GetOcIcEmployeeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcEmployeeResponse) SetBody(v *GetOcIcEmployeeResponseBody) *GetOcIcEmployeeResponse {
	s.Body = v
	return s
}

type GetOcIcEquityFrozenRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcEquityFrozenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEquityFrozenRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcEquityFrozenRequest) SetPageNo(v int32) *GetOcIcEquityFrozenRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcEquityFrozenRequest) SetPageSize(v int32) *GetOcIcEquityFrozenRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcEquityFrozenRequest) SetSearchKey(v string) *GetOcIcEquityFrozenRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcEquityFrozenResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcEquityFrozenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcEquityFrozenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEquityFrozenResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcEquityFrozenResponseBody) SetCode(v string) *GetOcIcEquityFrozenResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBody) SetData(v []*GetOcIcEquityFrozenResponseBodyData) *GetOcIcEquityFrozenResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcEquityFrozenResponseBody) SetMessage(v string) *GetOcIcEquityFrozenResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBody) SetPageIndex(v int32) *GetOcIcEquityFrozenResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBody) SetPageNum(v int32) *GetOcIcEquityFrozenResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBody) SetRequestId(v string) *GetOcIcEquityFrozenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBody) SetSuccess(v bool) *GetOcIcEquityFrozenResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBody) SetTotalNum(v int32) *GetOcIcEquityFrozenResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcEquityFrozenResponseBodyData struct {
	FreezeAmount      *string `json:"FreezeAmount,omitempty" xml:"FreezeAmount,omitempty"`
	FreezeCardNum     *string `json:"FreezeCardNum,omitempty" xml:"FreezeCardNum,omitempty"`
	FreezeCardType    *string `json:"FreezeCardType,omitempty" xml:"FreezeCardType,omitempty"`
	FreezeCourt       *string `json:"FreezeCourt,omitempty" xml:"FreezeCourt,omitempty"`
	FreezeEndDate     *string `json:"FreezeEndDate,omitempty" xml:"FreezeEndDate,omitempty"`
	FreezeExecItem    *string `json:"FreezeExecItem,omitempty" xml:"FreezeExecItem,omitempty"`
	FreezeExecPerson  *string `json:"FreezeExecPerson,omitempty" xml:"FreezeExecPerson,omitempty"`
	FreezeNoticeNum   *string `json:"FreezeNoticeNum,omitempty" xml:"FreezeNoticeNum,omitempty"`
	FreezePublishDate *string `json:"FreezePublishDate,omitempty" xml:"FreezePublishDate,omitempty"`
	FreezeStartDate   *string `json:"FreezeStartDate,omitempty" xml:"FreezeStartDate,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnfreezeAdjustNum *string `json:"UnfreezeAdjustNum,omitempty" xml:"UnfreezeAdjustNum,omitempty"`
	UnfreezeCourt     *string `json:"UnfreezeCourt,omitempty" xml:"UnfreezeCourt,omitempty"`
	UnfreezeDate      *string `json:"UnfreezeDate,omitempty" xml:"UnfreezeDate,omitempty"`
	UnfreezeReason    *string `json:"UnfreezeReason,omitempty" xml:"UnfreezeReason,omitempty"`
}

func (s GetOcIcEquityFrozenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEquityFrozenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezeAmount(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezeAmount = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezeCardNum(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezeCardNum = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezeCardType(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezeCardType = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezeCourt(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezeCourt = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezeEndDate(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezeEndDate = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezeExecItem(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezeExecItem = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezeExecPerson(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezeExecPerson = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezeNoticeNum(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezeNoticeNum = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezePublishDate(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezePublishDate = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetFreezeStartDate(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.FreezeStartDate = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetStatus(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetUnfreezeAdjustNum(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.UnfreezeAdjustNum = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetUnfreezeCourt(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.UnfreezeCourt = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetUnfreezeDate(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.UnfreezeDate = &v
	return s
}

func (s *GetOcIcEquityFrozenResponseBodyData) SetUnfreezeReason(v string) *GetOcIcEquityFrozenResponseBodyData {
	s.UnfreezeReason = &v
	return s
}

type GetOcIcEquityFrozenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcEquityFrozenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcEquityFrozenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEquityFrozenResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcEquityFrozenResponse) SetHeaders(v map[string]*string) *GetOcIcEquityFrozenResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcEquityFrozenResponse) SetStatusCode(v int32) *GetOcIcEquityFrozenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcEquityFrozenResponse) SetBody(v *GetOcIcEquityFrozenResponseBody) *GetOcIcEquityFrozenResponse {
	s.Body = v
	return s
}

type GetOcIcEquityPledgeRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcEquityPledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEquityPledgeRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcEquityPledgeRequest) SetPageNo(v int32) *GetOcIcEquityPledgeRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcEquityPledgeRequest) SetPageSize(v int32) *GetOcIcEquityPledgeRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcEquityPledgeRequest) SetSearchKey(v string) *GetOcIcEquityPledgeRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcEquityPledgeResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcEquityPledgeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcEquityPledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEquityPledgeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcEquityPledgeResponseBody) SetCode(v string) *GetOcIcEquityPledgeResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBody) SetData(v []*GetOcIcEquityPledgeResponseBodyData) *GetOcIcEquityPledgeResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcEquityPledgeResponseBody) SetMessage(v string) *GetOcIcEquityPledgeResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBody) SetPageIndex(v int32) *GetOcIcEquityPledgeResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBody) SetPageNum(v int32) *GetOcIcEquityPledgeResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBody) SetRequestId(v string) *GetOcIcEquityPledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBody) SetSuccess(v bool) *GetOcIcEquityPledgeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBody) SetTotalNum(v int32) *GetOcIcEquityPledgeResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcEquityPledgeResponseBodyData struct {
	Number            *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Pawnee            *string `json:"Pawnee,omitempty" xml:"Pawnee,omitempty"`
	PawneeIdentifyNo  *string `json:"PawneeIdentifyNo,omitempty" xml:"PawneeIdentifyNo,omitempty"`
	Pledgor           *string `json:"Pledgor,omitempty" xml:"Pledgor,omitempty"`
	PledgorAmount     *string `json:"PledgorAmount,omitempty" xml:"PledgorAmount,omitempty"`
	PledgorIdentifyNo *string `json:"PledgorIdentifyNo,omitempty" xml:"PledgorIdentifyNo,omitempty"`
	PublicDate        *string `json:"PublicDate,omitempty" xml:"PublicDate,omitempty"`
	RegDate           *string `json:"RegDate,omitempty" xml:"RegDate,omitempty"`
	RelatedComp       *string `json:"RelatedComp,omitempty" xml:"RelatedComp,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOcIcEquityPledgeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEquityPledgeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetNumber(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.Number = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetPawnee(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.Pawnee = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetPawneeIdentifyNo(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.PawneeIdentifyNo = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetPledgor(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.Pledgor = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetPledgorAmount(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.PledgorAmount = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetPledgorIdentifyNo(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.PledgorIdentifyNo = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetPublicDate(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.PublicDate = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetRegDate(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.RegDate = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetRelatedComp(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.RelatedComp = &v
	return s
}

func (s *GetOcIcEquityPledgeResponseBodyData) SetStatus(v string) *GetOcIcEquityPledgeResponseBodyData {
	s.Status = &v
	return s
}

type GetOcIcEquityPledgeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcEquityPledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcEquityPledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcEquityPledgeResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcEquityPledgeResponse) SetHeaders(v map[string]*string) *GetOcIcEquityPledgeResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcEquityPledgeResponse) SetStatusCode(v int32) *GetOcIcEquityPledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcEquityPledgeResponse) SetBody(v *GetOcIcEquityPledgeResponseBody) *GetOcIcEquityPledgeResponse {
	s.Body = v
	return s
}

type GetOcIcInvestmentRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcInvestmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcInvestmentRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcInvestmentRequest) SetPageNo(v int32) *GetOcIcInvestmentRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcInvestmentRequest) SetPageSize(v int32) *GetOcIcInvestmentRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcInvestmentRequest) SetSearchKey(v string) *GetOcIcInvestmentRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcInvestmentResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcInvestmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                               `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                               `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                               `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcInvestmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcInvestmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcInvestmentResponseBody) SetCode(v string) *GetOcIcInvestmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcInvestmentResponseBody) SetData(v []*GetOcIcInvestmentResponseBodyData) *GetOcIcInvestmentResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcInvestmentResponseBody) SetMessage(v string) *GetOcIcInvestmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcInvestmentResponseBody) SetPageIndex(v int32) *GetOcIcInvestmentResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcInvestmentResponseBody) SetPageNum(v int32) *GetOcIcInvestmentResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcInvestmentResponseBody) SetRequestId(v string) *GetOcIcInvestmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcInvestmentResponseBody) SetSuccess(v bool) *GetOcIcInvestmentResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcInvestmentResponseBody) SetTotalNum(v int32) *GetOcIcInvestmentResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcInvestmentResponseBodyData struct {
	EntName          *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	InvestCreditCode *string `json:"InvestCreditCode,omitempty" xml:"InvestCreditCode,omitempty"`
	InvestEsDate     *string `json:"InvestEsDate,omitempty" xml:"InvestEsDate,omitempty"`
	InvestLegalName  *string `json:"InvestLegalName,omitempty" xml:"InvestLegalName,omitempty"`
	InvestLicenseNo  *string `json:"InvestLicenseNo,omitempty" xml:"InvestLicenseNo,omitempty"`
	InvestName       *string `json:"InvestName,omitempty" xml:"InvestName,omitempty"`
	InvestRegCap     *string `json:"InvestRegCap,omitempty" xml:"InvestRegCap,omitempty"`
	InvestStatus     *string `json:"InvestStatus,omitempty" xml:"InvestStatus,omitempty"`
	ShouldCap        *string `json:"ShouldCap,omitempty" xml:"ShouldCap,omitempty"`
	StockPercentage  *string `json:"StockPercentage,omitempty" xml:"StockPercentage,omitempty"`
}

func (s GetOcIcInvestmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcInvestmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcInvestmentResponseBodyData) SetEntName(v string) *GetOcIcInvestmentResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIcInvestmentResponseBodyData) SetInvestCreditCode(v string) *GetOcIcInvestmentResponseBodyData {
	s.InvestCreditCode = &v
	return s
}

func (s *GetOcIcInvestmentResponseBodyData) SetInvestEsDate(v string) *GetOcIcInvestmentResponseBodyData {
	s.InvestEsDate = &v
	return s
}

func (s *GetOcIcInvestmentResponseBodyData) SetInvestLegalName(v string) *GetOcIcInvestmentResponseBodyData {
	s.InvestLegalName = &v
	return s
}

func (s *GetOcIcInvestmentResponseBodyData) SetInvestLicenseNo(v string) *GetOcIcInvestmentResponseBodyData {
	s.InvestLicenseNo = &v
	return s
}

func (s *GetOcIcInvestmentResponseBodyData) SetInvestName(v string) *GetOcIcInvestmentResponseBodyData {
	s.InvestName = &v
	return s
}

func (s *GetOcIcInvestmentResponseBodyData) SetInvestRegCap(v string) *GetOcIcInvestmentResponseBodyData {
	s.InvestRegCap = &v
	return s
}

func (s *GetOcIcInvestmentResponseBodyData) SetInvestStatus(v string) *GetOcIcInvestmentResponseBodyData {
	s.InvestStatus = &v
	return s
}

func (s *GetOcIcInvestmentResponseBodyData) SetShouldCap(v string) *GetOcIcInvestmentResponseBodyData {
	s.ShouldCap = &v
	return s
}

func (s *GetOcIcInvestmentResponseBodyData) SetStockPercentage(v string) *GetOcIcInvestmentResponseBodyData {
	s.StockPercentage = &v
	return s
}

type GetOcIcInvestmentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcInvestmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcInvestmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcInvestmentResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcInvestmentResponse) SetHeaders(v map[string]*string) *GetOcIcInvestmentResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcInvestmentResponse) SetStatusCode(v int32) *GetOcIcInvestmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcInvestmentResponse) SetBody(v *GetOcIcInvestmentResponseBody) *GetOcIcInvestmentResponse {
	s.Body = v
	return s
}

type GetOcIcKnowledgePropertyPledgeRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcKnowledgePropertyPledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcKnowledgePropertyPledgeRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcKnowledgePropertyPledgeRequest) SetPageNo(v int32) *GetOcIcKnowledgePropertyPledgeRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeRequest) SetPageSize(v int32) *GetOcIcKnowledgePropertyPledgeRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeRequest) SetRequestId(v string) *GetOcIcKnowledgePropertyPledgeRequest {
	s.RequestId = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeRequest) SetSearchKey(v string) *GetOcIcKnowledgePropertyPledgeRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcKnowledgePropertyPledgeResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcKnowledgePropertyPledgeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                            `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcKnowledgePropertyPledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcKnowledgePropertyPledgeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBody) SetCode(v string) *GetOcIcKnowledgePropertyPledgeResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBody) SetData(v []*GetOcIcKnowledgePropertyPledgeResponseBodyData) *GetOcIcKnowledgePropertyPledgeResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBody) SetMessage(v string) *GetOcIcKnowledgePropertyPledgeResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBody) SetPageIndex(v int32) *GetOcIcKnowledgePropertyPledgeResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBody) SetPageNum(v int32) *GetOcIcKnowledgePropertyPledgeResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBody) SetRequestId(v string) *GetOcIcKnowledgePropertyPledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBody) SetSuccess(v bool) *GetOcIcKnowledgePropertyPledgeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBody) SetTotalNum(v int32) *GetOcIcKnowledgePropertyPledgeResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcKnowledgePropertyPledgeResponseBodyData struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Number     *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Pawnee     *string `json:"Pawnee,omitempty" xml:"Pawnee,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Pledgor    *string `json:"Pledgor,omitempty" xml:"Pledgor,omitempty"`
	PublicDate *string `json:"PublicDate,omitempty" xml:"PublicDate,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetOcIcKnowledgePropertyPledgeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcKnowledgePropertyPledgeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBodyData) SetName(v string) *GetOcIcKnowledgePropertyPledgeResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBodyData) SetNumber(v string) *GetOcIcKnowledgePropertyPledgeResponseBodyData {
	s.Number = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBodyData) SetPawnee(v string) *GetOcIcKnowledgePropertyPledgeResponseBodyData {
	s.Pawnee = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBodyData) SetPeriod(v string) *GetOcIcKnowledgePropertyPledgeResponseBodyData {
	s.Period = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBodyData) SetPledgor(v string) *GetOcIcKnowledgePropertyPledgeResponseBodyData {
	s.Pledgor = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBodyData) SetPublicDate(v string) *GetOcIcKnowledgePropertyPledgeResponseBodyData {
	s.PublicDate = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBodyData) SetStatus(v string) *GetOcIcKnowledgePropertyPledgeResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponseBodyData) SetType(v string) *GetOcIcKnowledgePropertyPledgeResponseBodyData {
	s.Type = &v
	return s
}

type GetOcIcKnowledgePropertyPledgeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcKnowledgePropertyPledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcKnowledgePropertyPledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcKnowledgePropertyPledgeResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcKnowledgePropertyPledgeResponse) SetHeaders(v map[string]*string) *GetOcIcKnowledgePropertyPledgeResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponse) SetStatusCode(v int32) *GetOcIcKnowledgePropertyPledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcKnowledgePropertyPledgeResponse) SetBody(v *GetOcIcKnowledgePropertyPledgeResponseBody) *GetOcIcKnowledgePropertyPledgeResponse {
	s.Body = v
	return s
}

type GetOcIcMortgageRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcMortgageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcMortgageRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcMortgageRequest) SetPageNo(v int32) *GetOcIcMortgageRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcMortgageRequest) SetPageSize(v int32) *GetOcIcMortgageRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcMortgageRequest) SetSearchKey(v string) *GetOcIcMortgageRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcMortgageResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcMortgageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                             `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                             `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcMortgageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcMortgageResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcMortgageResponseBody) SetCode(v string) *GetOcIcMortgageResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcMortgageResponseBody) SetData(v []*GetOcIcMortgageResponseBodyData) *GetOcIcMortgageResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcMortgageResponseBody) SetMessage(v string) *GetOcIcMortgageResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcMortgageResponseBody) SetPageIndex(v int32) *GetOcIcMortgageResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcMortgageResponseBody) SetPageNum(v int32) *GetOcIcMortgageResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcMortgageResponseBody) SetRequestId(v string) *GetOcIcMortgageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcMortgageResponseBody) SetSuccess(v bool) *GetOcIcMortgageResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcMortgageResponseBody) SetTotalNum(v int32) *GetOcIcMortgageResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcMortgageResponseBodyData struct {
	DebitAmount    *string `json:"DebitAmount,omitempty" xml:"DebitAmount,omitempty"`
	DebitPeriod    *string `json:"DebitPeriod,omitempty" xml:"DebitPeriod,omitempty"`
	DebitScope     *string `json:"DebitScope,omitempty" xml:"DebitScope,omitempty"`
	DebitType      *string `json:"DebitType,omitempty" xml:"DebitType,omitempty"`
	Department     *string `json:"Department,omitempty" xml:"Department,omitempty"`
	Guarantees     *string `json:"Guarantees,omitempty" xml:"Guarantees,omitempty"`
	IdentifyNo     *string `json:"IdentifyNo,omitempty" xml:"IdentifyNo,omitempty"`
	IdentifyType   *string `json:"IdentifyType,omitempty" xml:"IdentifyType,omitempty"`
	MortgageesName *string `json:"MortgageesName,omitempty" xml:"MortgageesName,omitempty"`
	Number         *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// ocid
	OnecompId  *string `json:"OnecompId,omitempty" xml:"OnecompId,omitempty"`
	PublicDate *string `json:"PublicDate,omitempty" xml:"PublicDate,omitempty"`
	RegDate    *string `json:"RegDate,omitempty" xml:"RegDate,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOcIcMortgageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcMortgageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcMortgageResponseBodyData) SetDebitAmount(v string) *GetOcIcMortgageResponseBodyData {
	s.DebitAmount = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetDebitPeriod(v string) *GetOcIcMortgageResponseBodyData {
	s.DebitPeriod = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetDebitScope(v string) *GetOcIcMortgageResponseBodyData {
	s.DebitScope = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetDebitType(v string) *GetOcIcMortgageResponseBodyData {
	s.DebitType = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetDepartment(v string) *GetOcIcMortgageResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetGuarantees(v string) *GetOcIcMortgageResponseBodyData {
	s.Guarantees = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetIdentifyNo(v string) *GetOcIcMortgageResponseBodyData {
	s.IdentifyNo = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetIdentifyType(v string) *GetOcIcMortgageResponseBodyData {
	s.IdentifyType = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetMortgageesName(v string) *GetOcIcMortgageResponseBodyData {
	s.MortgageesName = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetNumber(v string) *GetOcIcMortgageResponseBodyData {
	s.Number = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetOnecompId(v string) *GetOcIcMortgageResponseBodyData {
	s.OnecompId = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetPublicDate(v string) *GetOcIcMortgageResponseBodyData {
	s.PublicDate = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetRegDate(v string) *GetOcIcMortgageResponseBodyData {
	s.RegDate = &v
	return s
}

func (s *GetOcIcMortgageResponseBodyData) SetStatus(v string) *GetOcIcMortgageResponseBodyData {
	s.Status = &v
	return s
}

type GetOcIcMortgageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcMortgageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcMortgageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcMortgageResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcMortgageResponse) SetHeaders(v map[string]*string) *GetOcIcMortgageResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcMortgageResponse) SetStatusCode(v int32) *GetOcIcMortgageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcMortgageResponse) SetBody(v *GetOcIcMortgageResponseBody) *GetOcIcMortgageResponse {
	s.Body = v
	return s
}

type GetOcIcSeriousOffenseRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcSeriousOffenseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcSeriousOffenseRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcSeriousOffenseRequest) SetPageNo(v int32) *GetOcIcSeriousOffenseRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcSeriousOffenseRequest) SetPageSize(v int32) *GetOcIcSeriousOffenseRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcSeriousOffenseRequest) SetSearchKey(v string) *GetOcIcSeriousOffenseRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcSeriousOffenseResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcSeriousOffenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                   `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                   `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcSeriousOffenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcSeriousOffenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcSeriousOffenseResponseBody) SetCode(v string) *GetOcIcSeriousOffenseResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBody) SetData(v []*GetOcIcSeriousOffenseResponseBodyData) *GetOcIcSeriousOffenseResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBody) SetMessage(v string) *GetOcIcSeriousOffenseResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBody) SetPageIndex(v int32) *GetOcIcSeriousOffenseResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBody) SetPageNum(v int32) *GetOcIcSeriousOffenseResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBody) SetRequestId(v string) *GetOcIcSeriousOffenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBody) SetSuccess(v bool) *GetOcIcSeriousOffenseResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBody) SetTotalNum(v int32) *GetOcIcSeriousOffenseResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcSeriousOffenseResponseBodyData struct {
	EntName       *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	InDate        *string `json:"InDate,omitempty" xml:"InDate,omitempty"`
	InDepartment  *string `json:"InDepartment,omitempty" xml:"InDepartment,omitempty"`
	InReason      *string `json:"InReason,omitempty" xml:"InReason,omitempty"`
	OutDate       *string `json:"OutDate,omitempty" xml:"OutDate,omitempty"`
	OutDepartment *string `json:"OutDepartment,omitempty" xml:"OutDepartment,omitempty"`
	OutReason     *string `json:"OutReason,omitempty" xml:"OutReason,omitempty"`
}

func (s GetOcIcSeriousOffenseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcSeriousOffenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcSeriousOffenseResponseBodyData) SetEntName(v string) *GetOcIcSeriousOffenseResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBodyData) SetInDate(v string) *GetOcIcSeriousOffenseResponseBodyData {
	s.InDate = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBodyData) SetInDepartment(v string) *GetOcIcSeriousOffenseResponseBodyData {
	s.InDepartment = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBodyData) SetInReason(v string) *GetOcIcSeriousOffenseResponseBodyData {
	s.InReason = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBodyData) SetOutDate(v string) *GetOcIcSeriousOffenseResponseBodyData {
	s.OutDate = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBodyData) SetOutDepartment(v string) *GetOcIcSeriousOffenseResponseBodyData {
	s.OutDepartment = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponseBodyData) SetOutReason(v string) *GetOcIcSeriousOffenseResponseBodyData {
	s.OutReason = &v
	return s
}

type GetOcIcSeriousOffenseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcSeriousOffenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcSeriousOffenseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcSeriousOffenseResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcSeriousOffenseResponse) SetHeaders(v map[string]*string) *GetOcIcSeriousOffenseResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcSeriousOffenseResponse) SetStatusCode(v int32) *GetOcIcSeriousOffenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcSeriousOffenseResponse) SetBody(v *GetOcIcSeriousOffenseResponseBody) *GetOcIcSeriousOffenseResponse {
	s.Body = v
	return s
}

type GetOcIcShareholderRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcShareholderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcShareholderRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcShareholderRequest) SetPageNo(v int32) *GetOcIcShareholderRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcShareholderRequest) SetPageSize(v int32) *GetOcIcShareholderRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcShareholderRequest) SetSearchKey(v string) *GetOcIcShareholderRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcShareholderResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcShareholderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcShareholderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcShareholderResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcShareholderResponseBody) SetCode(v string) *GetOcIcShareholderResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcShareholderResponseBody) SetData(v []*GetOcIcShareholderResponseBodyData) *GetOcIcShareholderResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcShareholderResponseBody) SetMessage(v string) *GetOcIcShareholderResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcShareholderResponseBody) SetPageIndex(v int32) *GetOcIcShareholderResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcShareholderResponseBody) SetPageNum(v int32) *GetOcIcShareholderResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcShareholderResponseBody) SetRequestId(v string) *GetOcIcShareholderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcShareholderResponseBody) SetSuccess(v bool) *GetOcIcShareholderResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcShareholderResponseBody) SetTotalNum(v int32) *GetOcIcShareholderResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcShareholderResponseBodyData struct {
	ShouldCap     *string `json:"ShouldCap,omitempty" xml:"ShouldCap,omitempty"`
	ShouldCapTime *string `json:"ShouldCapTime,omitempty" xml:"ShouldCapTime,omitempty"`
	StockName     *string `json:"StockName,omitempty" xml:"StockName,omitempty"`
	StockPercent  *string `json:"StockPercent,omitempty" xml:"StockPercent,omitempty"`
	StockType     *string `json:"StockType,omitempty" xml:"StockType,omitempty"`
}

func (s GetOcIcShareholderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcShareholderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcShareholderResponseBodyData) SetShouldCap(v string) *GetOcIcShareholderResponseBodyData {
	s.ShouldCap = &v
	return s
}

func (s *GetOcIcShareholderResponseBodyData) SetShouldCapTime(v string) *GetOcIcShareholderResponseBodyData {
	s.ShouldCapTime = &v
	return s
}

func (s *GetOcIcShareholderResponseBodyData) SetStockName(v string) *GetOcIcShareholderResponseBodyData {
	s.StockName = &v
	return s
}

func (s *GetOcIcShareholderResponseBodyData) SetStockPercent(v string) *GetOcIcShareholderResponseBodyData {
	s.StockPercent = &v
	return s
}

func (s *GetOcIcShareholderResponseBodyData) SetStockType(v string) *GetOcIcShareholderResponseBodyData {
	s.StockType = &v
	return s
}

type GetOcIcShareholderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcShareholderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcShareholderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcShareholderResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcShareholderResponse) SetHeaders(v map[string]*string) *GetOcIcShareholderResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcShareholderResponse) SetStatusCode(v int32) *GetOcIcShareholderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcShareholderResponse) SetBody(v *GetOcIcShareholderResponseBody) *GetOcIcShareholderResponse {
	s.Body = v
	return s
}

type GetOcIcSimpleCancelRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIcSimpleCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcSimpleCancelRequest) GoString() string {
	return s.String()
}

func (s *GetOcIcSimpleCancelRequest) SetPageNo(v int32) *GetOcIcSimpleCancelRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIcSimpleCancelRequest) SetPageSize(v int32) *GetOcIcSimpleCancelRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIcSimpleCancelRequest) SetSearchKey(v string) *GetOcIcSimpleCancelRequest {
	s.SearchKey = &v
	return s
}

type GetOcIcSimpleCancelResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIcSimpleCancelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIcSimpleCancelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcSimpleCancelResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIcSimpleCancelResponseBody) SetCode(v string) *GetOcIcSimpleCancelResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBody) SetData(v []*GetOcIcSimpleCancelResponseBodyData) *GetOcIcSimpleCancelResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIcSimpleCancelResponseBody) SetMessage(v string) *GetOcIcSimpleCancelResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBody) SetPageIndex(v int32) *GetOcIcSimpleCancelResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBody) SetPageNum(v int32) *GetOcIcSimpleCancelResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBody) SetRequestId(v string) *GetOcIcSimpleCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBody) SetSuccess(v bool) *GetOcIcSimpleCancelResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBody) SetTotalNum(v int32) *GetOcIcSimpleCancelResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIcSimpleCancelResponseBodyData struct {
	Department       *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EntName          *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	NoticePeriod     *string `json:"NoticePeriod,omitempty" xml:"NoticePeriod,omitempty"`
	ScaContent       *string `json:"ScaContent,omitempty" xml:"ScaContent,omitempty"`
	ScaDate          *string `json:"ScaDate,omitempty" xml:"ScaDate,omitempty"`
	ScaProposer      *string `json:"ScaProposer,omitempty" xml:"ScaProposer,omitempty"`
	ScaResult        *string `json:"ScaResult,omitempty" xml:"ScaResult,omitempty"`
	ScaResultDate    *string `json:"ScaResultDate,omitempty" xml:"ScaResultDate,omitempty"`
	SocialCreditCode *string `json:"SocialCreditCode,omitempty" xml:"SocialCreditCode,omitempty"`
}

func (s GetOcIcSimpleCancelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcSimpleCancelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIcSimpleCancelResponseBodyData) SetDepartment(v string) *GetOcIcSimpleCancelResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBodyData) SetEntName(v string) *GetOcIcSimpleCancelResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBodyData) SetNoticePeriod(v string) *GetOcIcSimpleCancelResponseBodyData {
	s.NoticePeriod = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBodyData) SetScaContent(v string) *GetOcIcSimpleCancelResponseBodyData {
	s.ScaContent = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBodyData) SetScaDate(v string) *GetOcIcSimpleCancelResponseBodyData {
	s.ScaDate = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBodyData) SetScaProposer(v string) *GetOcIcSimpleCancelResponseBodyData {
	s.ScaProposer = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBodyData) SetScaResult(v string) *GetOcIcSimpleCancelResponseBodyData {
	s.ScaResult = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBodyData) SetScaResultDate(v string) *GetOcIcSimpleCancelResponseBodyData {
	s.ScaResultDate = &v
	return s
}

func (s *GetOcIcSimpleCancelResponseBodyData) SetSocialCreditCode(v string) *GetOcIcSimpleCancelResponseBodyData {
	s.SocialCreditCode = &v
	return s
}

type GetOcIcSimpleCancelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIcSimpleCancelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIcSimpleCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIcSimpleCancelResponse) GoString() string {
	return s.String()
}

func (s *GetOcIcSimpleCancelResponse) SetHeaders(v map[string]*string) *GetOcIcSimpleCancelResponse {
	s.Headers = v
	return s
}

func (s *GetOcIcSimpleCancelResponse) SetStatusCode(v int32) *GetOcIcSimpleCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIcSimpleCancelResponse) SetBody(v *GetOcIcSimpleCancelResponseBody) *GetOcIcSimpleCancelResponse {
	s.Body = v
	return s
}

type GetOcIpCertificateRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIpCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetOcIpCertificateRequest) SetPageNo(v int32) *GetOcIpCertificateRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIpCertificateRequest) SetPageSize(v int32) *GetOcIpCertificateRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIpCertificateRequest) SetSearchKey(v string) *GetOcIpCertificateRequest {
	s.SearchKey = &v
	return s
}

type GetOcIpCertificateResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIpCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIpCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIpCertificateResponseBody) SetCode(v string) *GetOcIpCertificateResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIpCertificateResponseBody) SetData(v []*GetOcIpCertificateResponseBodyData) *GetOcIpCertificateResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIpCertificateResponseBody) SetMessage(v string) *GetOcIpCertificateResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIpCertificateResponseBody) SetPageIndex(v int32) *GetOcIpCertificateResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIpCertificateResponseBody) SetPageNum(v int32) *GetOcIpCertificateResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIpCertificateResponseBody) SetRequestId(v string) *GetOcIpCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIpCertificateResponseBody) SetSuccess(v bool) *GetOcIpCertificateResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIpCertificateResponseBody) SetTotalNum(v int32) *GetOcIpCertificateResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIpCertificateResponseBodyData struct {
	AuthorizeDate       *string `json:"AuthorizeDate,omitempty" xml:"AuthorizeDate,omitempty"`
	AuthorizeDepartment *string `json:"AuthorizeDepartment,omitempty" xml:"AuthorizeDepartment,omitempty"`
	CertNum             *string `json:"CertNum,omitempty" xml:"CertNum,omitempty"`
	CertScope           *string `json:"CertScope,omitempty" xml:"CertScope,omitempty"`
	CertType            *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	EntName             *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
	PubDate             *string `json:"PubDate,omitempty" xml:"PubDate,omitempty"`
	ValidEndDate        *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	ValidStartDate      *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s GetOcIpCertificateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIpCertificateResponseBodyData) SetAuthorizeDate(v string) *GetOcIpCertificateResponseBodyData {
	s.AuthorizeDate = &v
	return s
}

func (s *GetOcIpCertificateResponseBodyData) SetAuthorizeDepartment(v string) *GetOcIpCertificateResponseBodyData {
	s.AuthorizeDepartment = &v
	return s
}

func (s *GetOcIpCertificateResponseBodyData) SetCertNum(v string) *GetOcIpCertificateResponseBodyData {
	s.CertNum = &v
	return s
}

func (s *GetOcIpCertificateResponseBodyData) SetCertScope(v string) *GetOcIpCertificateResponseBodyData {
	s.CertScope = &v
	return s
}

func (s *GetOcIpCertificateResponseBodyData) SetCertType(v string) *GetOcIpCertificateResponseBodyData {
	s.CertType = &v
	return s
}

func (s *GetOcIpCertificateResponseBodyData) SetEntName(v string) *GetOcIpCertificateResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIpCertificateResponseBodyData) SetProvince(v string) *GetOcIpCertificateResponseBodyData {
	s.Province = &v
	return s
}

func (s *GetOcIpCertificateResponseBodyData) SetPubDate(v string) *GetOcIpCertificateResponseBodyData {
	s.PubDate = &v
	return s
}

func (s *GetOcIpCertificateResponseBodyData) SetValidEndDate(v string) *GetOcIpCertificateResponseBodyData {
	s.ValidEndDate = &v
	return s
}

func (s *GetOcIpCertificateResponseBodyData) SetValidStartDate(v string) *GetOcIpCertificateResponseBodyData {
	s.ValidStartDate = &v
	return s
}

type GetOcIpCertificateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIpCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIpCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetOcIpCertificateResponse) SetHeaders(v map[string]*string) *GetOcIpCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetOcIpCertificateResponse) SetStatusCode(v int32) *GetOcIpCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIpCertificateResponse) SetBody(v *GetOcIpCertificateResponseBody) *GetOcIpCertificateResponse {
	s.Body = v
	return s
}

type GetOcIpDomainRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIpDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpDomainRequest) GoString() string {
	return s.String()
}

func (s *GetOcIpDomainRequest) SetPageNo(v int32) *GetOcIpDomainRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIpDomainRequest) SetPageSize(v int32) *GetOcIpDomainRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIpDomainRequest) SetSearchKey(v string) *GetOcIpDomainRequest {
	s.SearchKey = &v
	return s
}

type GetOcIpDomainResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIpDomainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                           `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                           `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIpDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIpDomainResponseBody) SetCode(v string) *GetOcIpDomainResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIpDomainResponseBody) SetData(v []*GetOcIpDomainResponseBodyData) *GetOcIpDomainResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIpDomainResponseBody) SetMessage(v string) *GetOcIpDomainResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIpDomainResponseBody) SetPageIndex(v int32) *GetOcIpDomainResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIpDomainResponseBody) SetPageNum(v int32) *GetOcIpDomainResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIpDomainResponseBody) SetRequestId(v string) *GetOcIpDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIpDomainResponseBody) SetSuccess(v bool) *GetOcIpDomainResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIpDomainResponseBody) SetTotalNum(v int32) *GetOcIpDomainResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIpDomainResponseBodyData struct {
	CheckDate *string `json:"CheckDate,omitempty" xml:"CheckDate,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EntName   *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	HomeUrl   *string `json:"HomeUrl,omitempty" xml:"HomeUrl,omitempty"`
	Number    *string `json:"Number,omitempty" xml:"Number,omitempty"`
	SiteName  *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s GetOcIpDomainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIpDomainResponseBodyData) SetCheckDate(v string) *GetOcIpDomainResponseBodyData {
	s.CheckDate = &v
	return s
}

func (s *GetOcIpDomainResponseBodyData) SetDomain(v string) *GetOcIpDomainResponseBodyData {
	s.Domain = &v
	return s
}

func (s *GetOcIpDomainResponseBodyData) SetEntName(v string) *GetOcIpDomainResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIpDomainResponseBodyData) SetHomeUrl(v string) *GetOcIpDomainResponseBodyData {
	s.HomeUrl = &v
	return s
}

func (s *GetOcIpDomainResponseBodyData) SetNumber(v string) *GetOcIpDomainResponseBodyData {
	s.Number = &v
	return s
}

func (s *GetOcIpDomainResponseBodyData) SetSiteName(v string) *GetOcIpDomainResponseBodyData {
	s.SiteName = &v
	return s
}

type GetOcIpDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIpDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIpDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpDomainResponse) GoString() string {
	return s.String()
}

func (s *GetOcIpDomainResponse) SetHeaders(v map[string]*string) *GetOcIpDomainResponse {
	s.Headers = v
	return s
}

func (s *GetOcIpDomainResponse) SetStatusCode(v int32) *GetOcIpDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIpDomainResponse) SetBody(v *GetOcIpDomainResponseBody) *GetOcIpDomainResponse {
	s.Body = v
	return s
}

type GetOcIpPatentRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIpPatentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpPatentRequest) GoString() string {
	return s.String()
}

func (s *GetOcIpPatentRequest) SetPageNo(v int32) *GetOcIpPatentRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIpPatentRequest) SetPageSize(v int32) *GetOcIpPatentRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIpPatentRequest) SetSearchKey(v string) *GetOcIpPatentRequest {
	s.SearchKey = &v
	return s
}

type GetOcIpPatentResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIpPatentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                           `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                           `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIpPatentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpPatentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIpPatentResponseBody) SetCode(v string) *GetOcIpPatentResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIpPatentResponseBody) SetData(v []*GetOcIpPatentResponseBodyData) *GetOcIpPatentResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIpPatentResponseBody) SetMessage(v string) *GetOcIpPatentResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIpPatentResponseBody) SetPageIndex(v int32) *GetOcIpPatentResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIpPatentResponseBody) SetPageNum(v int32) *GetOcIpPatentResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIpPatentResponseBody) SetRequestId(v string) *GetOcIpPatentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIpPatentResponseBody) SetSuccess(v bool) *GetOcIpPatentResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIpPatentResponseBody) SetTotalNum(v int32) *GetOcIpPatentResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIpPatentResponseBodyData struct {
	Agency       *string `json:"Agency,omitempty" xml:"Agency,omitempty"`
	Agent        *string `json:"Agent,omitempty" xml:"Agent,omitempty"`
	Brief        *string `json:"Brief,omitempty" xml:"Brief,omitempty"`
	CateNum      *string `json:"CateNum,omitempty" xml:"CateNum,omitempty"`
	EntName      *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	InventorList *string `json:"InventorList,omitempty" xml:"InventorList,omitempty"`
	MainClaim    *string `json:"MainClaim,omitempty" xml:"MainClaim,omitempty"`
	PatentName   *string `json:"PatentName,omitempty" xml:"PatentName,omitempty"`
	PatentStatus *string `json:"PatentStatus,omitempty" xml:"PatentStatus,omitempty"`
	PatentType   *string `json:"PatentType,omitempty" xml:"PatentType,omitempty"`
	PatenteeList *string `json:"PatenteeList,omitempty" xml:"PatenteeList,omitempty"`
	PrioDate     *string `json:"PrioDate,omitempty" xml:"PrioDate,omitempty"`
	PrioNum      *string `json:"PrioNum,omitempty" xml:"PrioNum,omitempty"`
	PublicDate   *string `json:"PublicDate,omitempty" xml:"PublicDate,omitempty"`
	PublicNum    *string `json:"PublicNum,omitempty" xml:"PublicNum,omitempty"`
	RequestDate  *string `json:"RequestDate,omitempty" xml:"RequestDate,omitempty"`
	RequestNum   *string `json:"RequestNum,omitempty" xml:"RequestNum,omitempty"`
}

func (s GetOcIpPatentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpPatentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIpPatentResponseBodyData) SetAgency(v string) *GetOcIpPatentResponseBodyData {
	s.Agency = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetAgent(v string) *GetOcIpPatentResponseBodyData {
	s.Agent = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetBrief(v string) *GetOcIpPatentResponseBodyData {
	s.Brief = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetCateNum(v string) *GetOcIpPatentResponseBodyData {
	s.CateNum = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetEntName(v string) *GetOcIpPatentResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetInventorList(v string) *GetOcIpPatentResponseBodyData {
	s.InventorList = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetMainClaim(v string) *GetOcIpPatentResponseBodyData {
	s.MainClaim = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetPatentName(v string) *GetOcIpPatentResponseBodyData {
	s.PatentName = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetPatentStatus(v string) *GetOcIpPatentResponseBodyData {
	s.PatentStatus = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetPatentType(v string) *GetOcIpPatentResponseBodyData {
	s.PatentType = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetPatenteeList(v string) *GetOcIpPatentResponseBodyData {
	s.PatenteeList = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetPrioDate(v string) *GetOcIpPatentResponseBodyData {
	s.PrioDate = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetPrioNum(v string) *GetOcIpPatentResponseBodyData {
	s.PrioNum = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetPublicDate(v string) *GetOcIpPatentResponseBodyData {
	s.PublicDate = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetPublicNum(v string) *GetOcIpPatentResponseBodyData {
	s.PublicNum = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetRequestDate(v string) *GetOcIpPatentResponseBodyData {
	s.RequestDate = &v
	return s
}

func (s *GetOcIpPatentResponseBodyData) SetRequestNum(v string) *GetOcIpPatentResponseBodyData {
	s.RequestNum = &v
	return s
}

type GetOcIpPatentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIpPatentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIpPatentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpPatentResponse) GoString() string {
	return s.String()
}

func (s *GetOcIpPatentResponse) SetHeaders(v map[string]*string) *GetOcIpPatentResponse {
	s.Headers = v
	return s
}

func (s *GetOcIpPatentResponse) SetStatusCode(v int32) *GetOcIpPatentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIpPatentResponse) SetBody(v *GetOcIpPatentResponseBody) *GetOcIpPatentResponse {
	s.Body = v
	return s
}

type GetOcIpSoftwareCopyrightRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIpSoftwareCopyrightRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpSoftwareCopyrightRequest) GoString() string {
	return s.String()
}

func (s *GetOcIpSoftwareCopyrightRequest) SetPageNo(v int32) *GetOcIpSoftwareCopyrightRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightRequest) SetPageSize(v int32) *GetOcIpSoftwareCopyrightRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightRequest) SetSearchKey(v string) *GetOcIpSoftwareCopyrightRequest {
	s.SearchKey = &v
	return s
}

type GetOcIpSoftwareCopyrightResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIpSoftwareCopyrightResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                      `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                      `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                      `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIpSoftwareCopyrightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpSoftwareCopyrightResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIpSoftwareCopyrightResponseBody) SetCode(v string) *GetOcIpSoftwareCopyrightResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBody) SetData(v []*GetOcIpSoftwareCopyrightResponseBodyData) *GetOcIpSoftwareCopyrightResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBody) SetMessage(v string) *GetOcIpSoftwareCopyrightResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBody) SetPageIndex(v int32) *GetOcIpSoftwareCopyrightResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBody) SetPageNum(v int32) *GetOcIpSoftwareCopyrightResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBody) SetRequestId(v string) *GetOcIpSoftwareCopyrightResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBody) SetSuccess(v bool) *GetOcIpSoftwareCopyrightResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBody) SetTotalNum(v int32) *GetOcIpSoftwareCopyrightResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIpSoftwareCopyrightResponseBodyData struct {
	ApprovalDate *string `json:"ApprovalDate,omitempty" xml:"ApprovalDate,omitempty"`
	CopyName     *string `json:"CopyName,omitempty" xml:"CopyName,omitempty"`
	CopyNum      *string `json:"CopyNum,omitempty" xml:"CopyNum,omitempty"`
	EntName      *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	FirstDate    *string `json:"FirstDate,omitempty" xml:"FirstDate,omitempty"`
	ShortName    *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
	SuccessDate  *string `json:"SuccessDate,omitempty" xml:"SuccessDate,omitempty"`
	TypeNum      *string `json:"TypeNum,omitempty" xml:"TypeNum,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetOcIpSoftwareCopyrightResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpSoftwareCopyrightResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIpSoftwareCopyrightResponseBodyData) SetApprovalDate(v string) *GetOcIpSoftwareCopyrightResponseBodyData {
	s.ApprovalDate = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBodyData) SetCopyName(v string) *GetOcIpSoftwareCopyrightResponseBodyData {
	s.CopyName = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBodyData) SetCopyNum(v string) *GetOcIpSoftwareCopyrightResponseBodyData {
	s.CopyNum = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBodyData) SetEntName(v string) *GetOcIpSoftwareCopyrightResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBodyData) SetFirstDate(v string) *GetOcIpSoftwareCopyrightResponseBodyData {
	s.FirstDate = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBodyData) SetShortName(v string) *GetOcIpSoftwareCopyrightResponseBodyData {
	s.ShortName = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBodyData) SetSuccessDate(v string) *GetOcIpSoftwareCopyrightResponseBodyData {
	s.SuccessDate = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBodyData) SetTypeNum(v string) *GetOcIpSoftwareCopyrightResponseBodyData {
	s.TypeNum = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponseBodyData) SetVersion(v string) *GetOcIpSoftwareCopyrightResponseBodyData {
	s.Version = &v
	return s
}

type GetOcIpSoftwareCopyrightResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIpSoftwareCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIpSoftwareCopyrightResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpSoftwareCopyrightResponse) GoString() string {
	return s.String()
}

func (s *GetOcIpSoftwareCopyrightResponse) SetHeaders(v map[string]*string) *GetOcIpSoftwareCopyrightResponse {
	s.Headers = v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponse) SetStatusCode(v int32) *GetOcIpSoftwareCopyrightResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIpSoftwareCopyrightResponse) SetBody(v *GetOcIpSoftwareCopyrightResponseBody) *GetOcIpSoftwareCopyrightResponse {
	s.Body = v
	return s
}

type GetOcIpTrademarkRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIpTrademarkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpTrademarkRequest) GoString() string {
	return s.String()
}

func (s *GetOcIpTrademarkRequest) SetPageNo(v int32) *GetOcIpTrademarkRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIpTrademarkRequest) SetPageSize(v int32) *GetOcIpTrademarkRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIpTrademarkRequest) SetSearchKey(v string) *GetOcIpTrademarkRequest {
	s.SearchKey = &v
	return s
}

type GetOcIpTrademarkResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIpTrademarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                              `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                              `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                              `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIpTrademarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpTrademarkResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIpTrademarkResponseBody) SetCode(v string) *GetOcIpTrademarkResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIpTrademarkResponseBody) SetData(v []*GetOcIpTrademarkResponseBodyData) *GetOcIpTrademarkResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIpTrademarkResponseBody) SetMessage(v string) *GetOcIpTrademarkResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIpTrademarkResponseBody) SetPageIndex(v int32) *GetOcIpTrademarkResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIpTrademarkResponseBody) SetPageNum(v int32) *GetOcIpTrademarkResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIpTrademarkResponseBody) SetRequestId(v string) *GetOcIpTrademarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIpTrademarkResponseBody) SetSuccess(v bool) *GetOcIpTrademarkResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIpTrademarkResponseBody) SetTotalNum(v int32) *GetOcIpTrademarkResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIpTrademarkResponseBodyData struct {
	Agent           *string `json:"Agent,omitempty" xml:"Agent,omitempty"`
	ApplyDate       *string `json:"ApplyDate,omitempty" xml:"ApplyDate,omitempty"`
	EntName         *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	FirstPubDate    *string `json:"FirstPubDate,omitempty" xml:"FirstPubDate,omitempty"`
	FirstPubNo      *string `json:"FirstPubNo,omitempty" xml:"FirstPubNo,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Period          *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegNum          *string `json:"RegNum,omitempty" xml:"RegNum,omitempty"`
	RegPubDate      *string `json:"RegPubDate,omitempty" xml:"RegPubDate,omitempty"`
	RegPubNo        *string `json:"RegPubNo,omitempty" xml:"RegPubNo,omitempty"`
	TrademarkForm   *string `json:"TrademarkForm,omitempty" xml:"TrademarkForm,omitempty"`
	TrademarkName   *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkStatus *string `json:"TrademarkStatus,omitempty" xml:"TrademarkStatus,omitempty"`
	TrademarkType   *string `json:"TrademarkType,omitempty" xml:"TrademarkType,omitempty"`
	TypeName        *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetOcIpTrademarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpTrademarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIpTrademarkResponseBodyData) SetAgent(v string) *GetOcIpTrademarkResponseBodyData {
	s.Agent = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetApplyDate(v string) *GetOcIpTrademarkResponseBodyData {
	s.ApplyDate = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetEntName(v string) *GetOcIpTrademarkResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetFirstPubDate(v string) *GetOcIpTrademarkResponseBodyData {
	s.FirstPubDate = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetFirstPubNo(v string) *GetOcIpTrademarkResponseBodyData {
	s.FirstPubNo = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetImageUrl(v string) *GetOcIpTrademarkResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetPeriod(v string) *GetOcIpTrademarkResponseBodyData {
	s.Period = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetRegNum(v string) *GetOcIpTrademarkResponseBodyData {
	s.RegNum = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetRegPubDate(v string) *GetOcIpTrademarkResponseBodyData {
	s.RegPubDate = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetRegPubNo(v string) *GetOcIpTrademarkResponseBodyData {
	s.RegPubNo = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetTrademarkForm(v string) *GetOcIpTrademarkResponseBodyData {
	s.TrademarkForm = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetTrademarkName(v string) *GetOcIpTrademarkResponseBodyData {
	s.TrademarkName = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetTrademarkStatus(v string) *GetOcIpTrademarkResponseBodyData {
	s.TrademarkStatus = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetTrademarkType(v string) *GetOcIpTrademarkResponseBodyData {
	s.TrademarkType = &v
	return s
}

func (s *GetOcIpTrademarkResponseBodyData) SetTypeName(v string) *GetOcIpTrademarkResponseBodyData {
	s.TypeName = &v
	return s
}

type GetOcIpTrademarkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIpTrademarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIpTrademarkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpTrademarkResponse) GoString() string {
	return s.String()
}

func (s *GetOcIpTrademarkResponse) SetHeaders(v map[string]*string) *GetOcIpTrademarkResponse {
	s.Headers = v
	return s
}

func (s *GetOcIpTrademarkResponse) SetStatusCode(v int32) *GetOcIpTrademarkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIpTrademarkResponse) SetBody(v *GetOcIpTrademarkResponseBody) *GetOcIpTrademarkResponse {
	s.Body = v
	return s
}

type GetOcIpWorksCopyrightRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcIpWorksCopyrightRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpWorksCopyrightRequest) GoString() string {
	return s.String()
}

func (s *GetOcIpWorksCopyrightRequest) SetPageNo(v int32) *GetOcIpWorksCopyrightRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcIpWorksCopyrightRequest) SetPageSize(v int32) *GetOcIpWorksCopyrightRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcIpWorksCopyrightRequest) SetSearchKey(v string) *GetOcIpWorksCopyrightRequest {
	s.SearchKey = &v
	return s
}

type GetOcIpWorksCopyrightResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcIpWorksCopyrightResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                   `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                   `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcIpWorksCopyrightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpWorksCopyrightResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcIpWorksCopyrightResponseBody) SetCode(v string) *GetOcIpWorksCopyrightResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBody) SetData(v []*GetOcIpWorksCopyrightResponseBodyData) *GetOcIpWorksCopyrightResponseBody {
	s.Data = v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBody) SetMessage(v string) *GetOcIpWorksCopyrightResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBody) SetPageIndex(v int32) *GetOcIpWorksCopyrightResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBody) SetPageNum(v int32) *GetOcIpWorksCopyrightResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBody) SetRequestId(v string) *GetOcIpWorksCopyrightResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBody) SetSuccess(v bool) *GetOcIpWorksCopyrightResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBody) SetTotalNum(v int32) *GetOcIpWorksCopyrightResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcIpWorksCopyrightResponseBodyData struct {
	ApprovalDate *string `json:"ApprovalDate,omitempty" xml:"ApprovalDate,omitempty"`
	CopyName     *string `json:"CopyName,omitempty" xml:"CopyName,omitempty"`
	CopyNum      *string `json:"CopyNum,omitempty" xml:"CopyNum,omitempty"`
	EntName      *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	FirstDate    *string `json:"FirstDate,omitempty" xml:"FirstDate,omitempty"`
	SuccessDate  *string `json:"SuccessDate,omitempty" xml:"SuccessDate,omitempty"`
	TypeName     *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetOcIpWorksCopyrightResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpWorksCopyrightResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcIpWorksCopyrightResponseBodyData) SetApprovalDate(v string) *GetOcIpWorksCopyrightResponseBodyData {
	s.ApprovalDate = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBodyData) SetCopyName(v string) *GetOcIpWorksCopyrightResponseBodyData {
	s.CopyName = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBodyData) SetCopyNum(v string) *GetOcIpWorksCopyrightResponseBodyData {
	s.CopyNum = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBodyData) SetEntName(v string) *GetOcIpWorksCopyrightResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBodyData) SetFirstDate(v string) *GetOcIpWorksCopyrightResponseBodyData {
	s.FirstDate = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBodyData) SetSuccessDate(v string) *GetOcIpWorksCopyrightResponseBodyData {
	s.SuccessDate = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponseBodyData) SetTypeName(v string) *GetOcIpWorksCopyrightResponseBodyData {
	s.TypeName = &v
	return s
}

type GetOcIpWorksCopyrightResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcIpWorksCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcIpWorksCopyrightResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcIpWorksCopyrightResponse) GoString() string {
	return s.String()
}

func (s *GetOcIpWorksCopyrightResponse) SetHeaders(v map[string]*string) *GetOcIpWorksCopyrightResponse {
	s.Headers = v
	return s
}

func (s *GetOcIpWorksCopyrightResponse) SetStatusCode(v int32) *GetOcIpWorksCopyrightResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcIpWorksCopyrightResponse) SetBody(v *GetOcIpWorksCopyrightResponseBody) *GetOcIpWorksCopyrightResponse {
	s.Body = v
	return s
}

type GetOcJusticeAuctionRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcJusticeAuctionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeAuctionRequest) GoString() string {
	return s.String()
}

func (s *GetOcJusticeAuctionRequest) SetPageNo(v int32) *GetOcJusticeAuctionRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcJusticeAuctionRequest) SetPageSize(v int32) *GetOcJusticeAuctionRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcJusticeAuctionRequest) SetSearchKey(v string) *GetOcJusticeAuctionRequest {
	s.SearchKey = &v
	return s
}

type GetOcJusticeAuctionResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcJusticeAuctionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcJusticeAuctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeAuctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcJusticeAuctionResponseBody) SetCode(v string) *GetOcJusticeAuctionResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBody) SetData(v []*GetOcJusticeAuctionResponseBodyData) *GetOcJusticeAuctionResponseBody {
	s.Data = v
	return s
}

func (s *GetOcJusticeAuctionResponseBody) SetMessage(v string) *GetOcJusticeAuctionResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBody) SetPageIndex(v int32) *GetOcJusticeAuctionResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBody) SetPageNum(v int32) *GetOcJusticeAuctionResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBody) SetRequestId(v string) *GetOcJusticeAuctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBody) SetSuccess(v bool) *GetOcJusticeAuctionResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBody) SetTotalNum(v int32) *GetOcJusticeAuctionResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcJusticeAuctionResponseBodyData struct {
	AuctionDate *string `json:"AuctionDate,omitempty" xml:"AuctionDate,omitempty"`
	AuctionName *string `json:"AuctionName,omitempty" xml:"AuctionName,omitempty"`
	Basis       *string `json:"Basis,omitempty" xml:"Basis,omitempty"`
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Court       *string `json:"Court,omitempty" xml:"Court,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Document    *string `json:"Document,omitempty" xml:"Document,omitempty"`
	EntName     *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	EstPrice    *string `json:"EstPrice,omitempty" xml:"EstPrice,omitempty"`
	Owner       *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Restrict    *string `json:"Restrict,omitempty" xml:"Restrict,omitempty"`
	StartPrice  *string `json:"StartPrice,omitempty" xml:"StartPrice,omitempty"`
}

func (s GetOcJusticeAuctionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeAuctionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcJusticeAuctionResponseBodyData) SetAuctionDate(v string) *GetOcJusticeAuctionResponseBodyData {
	s.AuctionDate = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetAuctionName(v string) *GetOcJusticeAuctionResponseBodyData {
	s.AuctionName = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetBasis(v string) *GetOcJusticeAuctionResponseBodyData {
	s.Basis = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetCertificate(v string) *GetOcJusticeAuctionResponseBodyData {
	s.Certificate = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetCourt(v string) *GetOcJusticeAuctionResponseBodyData {
	s.Court = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetDescription(v string) *GetOcJusticeAuctionResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetDocument(v string) *GetOcJusticeAuctionResponseBodyData {
	s.Document = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetEntName(v string) *GetOcJusticeAuctionResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetEstPrice(v string) *GetOcJusticeAuctionResponseBodyData {
	s.EstPrice = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetOwner(v string) *GetOcJusticeAuctionResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetRestrict(v string) *GetOcJusticeAuctionResponseBodyData {
	s.Restrict = &v
	return s
}

func (s *GetOcJusticeAuctionResponseBodyData) SetStartPrice(v string) *GetOcJusticeAuctionResponseBodyData {
	s.StartPrice = &v
	return s
}

type GetOcJusticeAuctionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcJusticeAuctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcJusticeAuctionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeAuctionResponse) GoString() string {
	return s.String()
}

func (s *GetOcJusticeAuctionResponse) SetHeaders(v map[string]*string) *GetOcJusticeAuctionResponse {
	s.Headers = v
	return s
}

func (s *GetOcJusticeAuctionResponse) SetStatusCode(v int32) *GetOcJusticeAuctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcJusticeAuctionResponse) SetBody(v *GetOcJusticeAuctionResponseBody) *GetOcJusticeAuctionResponse {
	s.Body = v
	return s
}

type GetOcJusticeCaseFilingRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcJusticeCaseFilingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCaseFilingRequest) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCaseFilingRequest) SetPageNo(v int32) *GetOcJusticeCaseFilingRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcJusticeCaseFilingRequest) SetPageSize(v int32) *GetOcJusticeCaseFilingRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcJusticeCaseFilingRequest) SetSearchKey(v string) *GetOcJusticeCaseFilingRequest {
	s.SearchKey = &v
	return s
}

type GetOcJusticeCaseFilingResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcJusticeCaseFilingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                    `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                    `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcJusticeCaseFilingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCaseFilingResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCaseFilingResponseBody) SetCode(v string) *GetOcJusticeCaseFilingResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBody) SetData(v []*GetOcJusticeCaseFilingResponseBodyData) *GetOcJusticeCaseFilingResponseBody {
	s.Data = v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBody) SetMessage(v string) *GetOcJusticeCaseFilingResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBody) SetPageIndex(v int32) *GetOcJusticeCaseFilingResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBody) SetPageNum(v int32) *GetOcJusticeCaseFilingResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBody) SetRequestId(v string) *GetOcJusticeCaseFilingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBody) SetSuccess(v bool) *GetOcJusticeCaseFilingResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBody) SetTotalNum(v int32) *GetOcJusticeCaseFilingResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcJusticeCaseFilingResponseBodyData struct {
	Assistant   *string `json:"Assistant,omitempty" xml:"Assistant,omitempty"`
	CaseNum     *string `json:"CaseNum,omitempty" xml:"CaseNum,omitempty"`
	CaseStatus  *string `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	CauseAction *string `json:"CauseAction,omitempty" xml:"CauseAction,omitempty"`
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	FilingDate  *string `json:"FilingDate,omitempty" xml:"FilingDate,omitempty"`
	HearingDate *string `json:"HearingDate,omitempty" xml:"HearingDate,omitempty"`
	Judge       *string `json:"Judge,omitempty" xml:"Judge,omitempty"`
	Party       *string `json:"Party,omitempty" xml:"Party,omitempty"`
	Role        *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s GetOcJusticeCaseFilingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCaseFilingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetAssistant(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.Assistant = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetCaseNum(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.CaseNum = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetCaseStatus(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.CaseStatus = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetCauseAction(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.CauseAction = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetEndDate(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetFilingDate(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.FilingDate = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetHearingDate(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.HearingDate = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetJudge(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.Judge = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetParty(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.Party = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponseBodyData) SetRole(v string) *GetOcJusticeCaseFilingResponseBodyData {
	s.Role = &v
	return s
}

type GetOcJusticeCaseFilingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcJusticeCaseFilingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcJusticeCaseFilingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCaseFilingResponse) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCaseFilingResponse) SetHeaders(v map[string]*string) *GetOcJusticeCaseFilingResponse {
	s.Headers = v
	return s
}

func (s *GetOcJusticeCaseFilingResponse) SetStatusCode(v int32) *GetOcJusticeCaseFilingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcJusticeCaseFilingResponse) SetBody(v *GetOcJusticeCaseFilingResponseBody) *GetOcJusticeCaseFilingResponse {
	s.Body = v
	return s
}

type GetOcJusticeCourtAnnouncementRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcJusticeCourtAnnouncementRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCourtAnnouncementRequest) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCourtAnnouncementRequest) SetPageNo(v int32) *GetOcJusticeCourtAnnouncementRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementRequest) SetPageSize(v int32) *GetOcJusticeCourtAnnouncementRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementRequest) SetSearchKey(v string) *GetOcJusticeCourtAnnouncementRequest {
	s.SearchKey = &v
	return s
}

type GetOcJusticeCourtAnnouncementResponseBody struct {
	Code      *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcJusticeCourtAnnouncementResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                           `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                           `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcJusticeCourtAnnouncementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCourtAnnouncementResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCourtAnnouncementResponseBody) SetCode(v string) *GetOcJusticeCourtAnnouncementResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBody) SetData(v []*GetOcJusticeCourtAnnouncementResponseBodyData) *GetOcJusticeCourtAnnouncementResponseBody {
	s.Data = v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBody) SetMessage(v string) *GetOcJusticeCourtAnnouncementResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBody) SetPageIndex(v int32) *GetOcJusticeCourtAnnouncementResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBody) SetPageNum(v int32) *GetOcJusticeCourtAnnouncementResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBody) SetRequestId(v string) *GetOcJusticeCourtAnnouncementResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBody) SetSuccess(v bool) *GetOcJusticeCourtAnnouncementResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBody) SetTotalNum(v int32) *GetOcJusticeCourtAnnouncementResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcJusticeCourtAnnouncementResponseBodyData struct {
	CaseNum     *string `json:"CaseNum,omitempty" xml:"CaseNum,omitempty"`
	CauseAction *string `json:"CauseAction,omitempty" xml:"CauseAction,omitempty"`
	Court       *string `json:"Court,omitempty" xml:"Court,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	HearingDate *string `json:"HearingDate,omitempty" xml:"HearingDate,omitempty"`
	Judge       *string `json:"Judge,omitempty" xml:"Judge,omitempty"`
	Party       *string `json:"Party,omitempty" xml:"Party,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Tribunal    *string `json:"Tribunal,omitempty" xml:"Tribunal,omitempty"`
}

func (s GetOcJusticeCourtAnnouncementResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCourtAnnouncementResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCourtAnnouncementResponseBodyData) SetCaseNum(v string) *GetOcJusticeCourtAnnouncementResponseBodyData {
	s.CaseNum = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBodyData) SetCauseAction(v string) *GetOcJusticeCourtAnnouncementResponseBodyData {
	s.CauseAction = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBodyData) SetCourt(v string) *GetOcJusticeCourtAnnouncementResponseBodyData {
	s.Court = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBodyData) SetDepartment(v string) *GetOcJusticeCourtAnnouncementResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBodyData) SetHearingDate(v string) *GetOcJusticeCourtAnnouncementResponseBodyData {
	s.HearingDate = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBodyData) SetJudge(v string) *GetOcJusticeCourtAnnouncementResponseBodyData {
	s.Judge = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBodyData) SetParty(v string) *GetOcJusticeCourtAnnouncementResponseBodyData {
	s.Party = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBodyData) SetTitle(v string) *GetOcJusticeCourtAnnouncementResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponseBodyData) SetTribunal(v string) *GetOcJusticeCourtAnnouncementResponseBodyData {
	s.Tribunal = &v
	return s
}

type GetOcJusticeCourtAnnouncementResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcJusticeCourtAnnouncementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcJusticeCourtAnnouncementResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCourtAnnouncementResponse) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCourtAnnouncementResponse) SetHeaders(v map[string]*string) *GetOcJusticeCourtAnnouncementResponse {
	s.Headers = v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponse) SetStatusCode(v int32) *GetOcJusticeCourtAnnouncementResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcJusticeCourtAnnouncementResponse) SetBody(v *GetOcJusticeCourtAnnouncementResponseBody) *GetOcJusticeCourtAnnouncementResponse {
	s.Body = v
	return s
}

type GetOcJusticeCourtNoticeRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcJusticeCourtNoticeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCourtNoticeRequest) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCourtNoticeRequest) SetPageNo(v int32) *GetOcJusticeCourtNoticeRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcJusticeCourtNoticeRequest) SetPageSize(v int32) *GetOcJusticeCourtNoticeRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcJusticeCourtNoticeRequest) SetRequestId(v string) *GetOcJusticeCourtNoticeRequest {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeCourtNoticeRequest) SetSearchKey(v string) *GetOcJusticeCourtNoticeRequest {
	s.SearchKey = &v
	return s
}

type GetOcJusticeCourtNoticeResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcJusticeCourtNoticeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                     `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                     `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                     `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcJusticeCourtNoticeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCourtNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCourtNoticeResponseBody) SetCode(v string) *GetOcJusticeCourtNoticeResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBody) SetData(v []*GetOcJusticeCourtNoticeResponseBodyData) *GetOcJusticeCourtNoticeResponseBody {
	s.Data = v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBody) SetMessage(v string) *GetOcJusticeCourtNoticeResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBody) SetPageIndex(v int32) *GetOcJusticeCourtNoticeResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBody) SetPageNum(v int32) *GetOcJusticeCourtNoticeResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBody) SetRequestId(v string) *GetOcJusticeCourtNoticeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBody) SetSuccess(v bool) *GetOcJusticeCourtNoticeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBody) SetTotalNum(v int32) *GetOcJusticeCourtNoticeResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcJusticeCourtNoticeResponseBodyData struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Court      *string `json:"Court,omitempty" xml:"Court,omitempty"`
	Party      *string `json:"Party,omitempty" xml:"Party,omitempty"`
	PublicDate *string `json:"PublicDate,omitempty" xml:"PublicDate,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetOcJusticeCourtNoticeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCourtNoticeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCourtNoticeResponseBodyData) SetContent(v string) *GetOcJusticeCourtNoticeResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBodyData) SetCourt(v string) *GetOcJusticeCourtNoticeResponseBodyData {
	s.Court = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBodyData) SetParty(v string) *GetOcJusticeCourtNoticeResponseBodyData {
	s.Party = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBodyData) SetPublicDate(v string) *GetOcJusticeCourtNoticeResponseBodyData {
	s.PublicDate = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponseBodyData) SetType(v string) *GetOcJusticeCourtNoticeResponseBodyData {
	s.Type = &v
	return s
}

type GetOcJusticeCourtNoticeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcJusticeCourtNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcJusticeCourtNoticeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeCourtNoticeResponse) GoString() string {
	return s.String()
}

func (s *GetOcJusticeCourtNoticeResponse) SetHeaders(v map[string]*string) *GetOcJusticeCourtNoticeResponse {
	s.Headers = v
	return s
}

func (s *GetOcJusticeCourtNoticeResponse) SetStatusCode(v int32) *GetOcJusticeCourtNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcJusticeCourtNoticeResponse) SetBody(v *GetOcJusticeCourtNoticeResponseBody) *GetOcJusticeCourtNoticeResponse {
	s.Body = v
	return s
}

type GetOcJusticeDishonestyRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcJusticeDishonestyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeDishonestyRequest) GoString() string {
	return s.String()
}

func (s *GetOcJusticeDishonestyRequest) SetPageNo(v int32) *GetOcJusticeDishonestyRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcJusticeDishonestyRequest) SetPageSize(v int32) *GetOcJusticeDishonestyRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcJusticeDishonestyRequest) SetSearchKey(v string) *GetOcJusticeDishonestyRequest {
	s.SearchKey = &v
	return s
}

type GetOcJusticeDishonestyResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcJusticeDishonestyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                    `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                    `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcJusticeDishonestyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeDishonestyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcJusticeDishonestyResponseBody) SetCode(v string) *GetOcJusticeDishonestyResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBody) SetData(v []*GetOcJusticeDishonestyResponseBodyData) *GetOcJusticeDishonestyResponseBody {
	s.Data = v
	return s
}

func (s *GetOcJusticeDishonestyResponseBody) SetMessage(v string) *GetOcJusticeDishonestyResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBody) SetPageIndex(v int32) *GetOcJusticeDishonestyResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBody) SetPageNum(v int32) *GetOcJusticeDishonestyResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBody) SetRequestId(v string) *GetOcJusticeDishonestyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBody) SetSuccess(v bool) *GetOcJusticeDishonestyResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBody) SetTotalNum(v int32) *GetOcJusticeDishonestyResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcJusticeDishonestyResponseBodyData struct {
	Amount            *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	CaseNum           *string `json:"CaseNum,omitempty" xml:"CaseNum,omitempty"`
	Court             *string `json:"Court,omitempty" xml:"Court,omitempty"`
	EntName           *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	ExecuteDepartment *string `json:"ExecuteDepartment,omitempty" xml:"ExecuteDepartment,omitempty"`
	ExecutionDesc     *string `json:"ExecutionDesc,omitempty" xml:"ExecutionDesc,omitempty"`
	ExecutionStatus   *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	FilingDate        *string `json:"FilingDate,omitempty" xml:"FilingDate,omitempty"`
	FinalDuty         *string `json:"FinalDuty,omitempty" xml:"FinalDuty,omitempty"`
	FromCaseNum       *string `json:"FromCaseNum,omitempty" xml:"FromCaseNum,omitempty"`
	LegalName         *string `json:"LegalName,omitempty" xml:"LegalName,omitempty"`
	Province          *string `json:"Province,omitempty" xml:"Province,omitempty"`
	PublishDate       *string `json:"PublishDate,omitempty" xml:"PublishDate,omitempty"`
	SocialCreditCode  *string `json:"SocialCreditCode,omitempty" xml:"SocialCreditCode,omitempty"`
}

func (s GetOcJusticeDishonestyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeDishonestyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetAmount(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.Amount = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetCaseNum(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.CaseNum = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetCourt(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.Court = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetEntName(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetExecuteDepartment(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.ExecuteDepartment = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetExecutionDesc(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.ExecutionDesc = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetExecutionStatus(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.ExecutionStatus = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetFilingDate(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.FilingDate = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetFinalDuty(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.FinalDuty = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetFromCaseNum(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.FromCaseNum = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetLegalName(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.LegalName = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetProvince(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.Province = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetPublishDate(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.PublishDate = &v
	return s
}

func (s *GetOcJusticeDishonestyResponseBodyData) SetSocialCreditCode(v string) *GetOcJusticeDishonestyResponseBodyData {
	s.SocialCreditCode = &v
	return s
}

type GetOcJusticeDishonestyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcJusticeDishonestyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcJusticeDishonestyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeDishonestyResponse) GoString() string {
	return s.String()
}

func (s *GetOcJusticeDishonestyResponse) SetHeaders(v map[string]*string) *GetOcJusticeDishonestyResponse {
	s.Headers = v
	return s
}

func (s *GetOcJusticeDishonestyResponse) SetStatusCode(v int32) *GetOcJusticeDishonestyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcJusticeDishonestyResponse) SetBody(v *GetOcJusticeDishonestyResponseBody) *GetOcJusticeDishonestyResponse {
	s.Body = v
	return s
}

type GetOcJusticeExecutedRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcJusticeExecutedRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeExecutedRequest) GoString() string {
	return s.String()
}

func (s *GetOcJusticeExecutedRequest) SetPageNo(v int32) *GetOcJusticeExecutedRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcJusticeExecutedRequest) SetPageSize(v int32) *GetOcJusticeExecutedRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcJusticeExecutedRequest) SetRequestId(v string) *GetOcJusticeExecutedRequest {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeExecutedRequest) SetSearchKey(v string) *GetOcJusticeExecutedRequest {
	s.SearchKey = &v
	return s
}

type GetOcJusticeExecutedResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcJusticeExecutedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                  `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcJusticeExecutedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeExecutedResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcJusticeExecutedResponseBody) SetCode(v string) *GetOcJusticeExecutedResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcJusticeExecutedResponseBody) SetData(v []*GetOcJusticeExecutedResponseBodyData) *GetOcJusticeExecutedResponseBody {
	s.Data = v
	return s
}

func (s *GetOcJusticeExecutedResponseBody) SetMessage(v string) *GetOcJusticeExecutedResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcJusticeExecutedResponseBody) SetPageIndex(v int32) *GetOcJusticeExecutedResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcJusticeExecutedResponseBody) SetPageNum(v int32) *GetOcJusticeExecutedResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcJusticeExecutedResponseBody) SetRequestId(v string) *GetOcJusticeExecutedResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeExecutedResponseBody) SetSuccess(v bool) *GetOcJusticeExecutedResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcJusticeExecutedResponseBody) SetTotalNum(v int32) *GetOcJusticeExecutedResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcJusticeExecutedResponseBodyData struct {
	Amount     *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	CaseNum    *string `json:"CaseNum,omitempty" xml:"CaseNum,omitempty"`
	Court      *string `json:"Court,omitempty" xml:"Court,omitempty"`
	FilingDate *string `json:"FilingDate,omitempty" xml:"FilingDate,omitempty"`
}

func (s GetOcJusticeExecutedResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeExecutedResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcJusticeExecutedResponseBodyData) SetAmount(v string) *GetOcJusticeExecutedResponseBodyData {
	s.Amount = &v
	return s
}

func (s *GetOcJusticeExecutedResponseBodyData) SetCaseNum(v string) *GetOcJusticeExecutedResponseBodyData {
	s.CaseNum = &v
	return s
}

func (s *GetOcJusticeExecutedResponseBodyData) SetCourt(v string) *GetOcJusticeExecutedResponseBodyData {
	s.Court = &v
	return s
}

func (s *GetOcJusticeExecutedResponseBodyData) SetFilingDate(v string) *GetOcJusticeExecutedResponseBodyData {
	s.FilingDate = &v
	return s
}

type GetOcJusticeExecutedResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcJusticeExecutedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcJusticeExecutedResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeExecutedResponse) GoString() string {
	return s.String()
}

func (s *GetOcJusticeExecutedResponse) SetHeaders(v map[string]*string) *GetOcJusticeExecutedResponse {
	s.Headers = v
	return s
}

func (s *GetOcJusticeExecutedResponse) SetStatusCode(v int32) *GetOcJusticeExecutedResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcJusticeExecutedResponse) SetBody(v *GetOcJusticeExecutedResponseBody) *GetOcJusticeExecutedResponse {
	s.Body = v
	return s
}

type GetOcJusticeJudgementDocRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcJusticeJudgementDocRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeJudgementDocRequest) GoString() string {
	return s.String()
}

func (s *GetOcJusticeJudgementDocRequest) SetPageNo(v int32) *GetOcJusticeJudgementDocRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcJusticeJudgementDocRequest) SetPageSize(v int32) *GetOcJusticeJudgementDocRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcJusticeJudgementDocRequest) SetSearchKey(v string) *GetOcJusticeJudgementDocRequest {
	s.SearchKey = &v
	return s
}

type GetOcJusticeJudgementDocResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcJusticeJudgementDocResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                      `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                      `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                      `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcJusticeJudgementDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeJudgementDocResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcJusticeJudgementDocResponseBody) SetCode(v string) *GetOcJusticeJudgementDocResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBody) SetData(v []*GetOcJusticeJudgementDocResponseBodyData) *GetOcJusticeJudgementDocResponseBody {
	s.Data = v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBody) SetMessage(v string) *GetOcJusticeJudgementDocResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBody) SetPageIndex(v int32) *GetOcJusticeJudgementDocResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBody) SetPageNum(v int32) *GetOcJusticeJudgementDocResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBody) SetRequestId(v string) *GetOcJusticeJudgementDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBody) SetSuccess(v bool) *GetOcJusticeJudgementDocResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBody) SetTotalNum(v int32) *GetOcJusticeJudgementDocResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcJusticeJudgementDocResponseBodyData struct {
	CaseFlow    *string `json:"CaseFlow,omitempty" xml:"CaseFlow,omitempty"`
	CaseNum     *string `json:"CaseNum,omitempty" xml:"CaseNum,omitempty"`
	CaseType    *string `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	CauseAction *string `json:"CauseAction,omitempty" xml:"CauseAction,omitempty"`
	Court       *string `json:"Court,omitempty" xml:"Court,omitempty"`
	Defendant   *string `json:"Defendant,omitempty" xml:"Defendant,omitempty"`
	JudgeDate   *string `json:"JudgeDate,omitempty" xml:"JudgeDate,omitempty"`
	JudgeResult *string `json:"JudgeResult,omitempty" xml:"JudgeResult,omitempty"`
	JudgeType   *string `json:"JudgeType,omitempty" xml:"JudgeType,omitempty"`
	Party       *string `json:"Party,omitempty" xml:"Party,omitempty"`
	Plaintiff   *string `json:"Plaintiff,omitempty" xml:"Plaintiff,omitempty"`
	PublicDate  *string `json:"PublicDate,omitempty" xml:"PublicDate,omitempty"`
	Role        *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SubAmount   *string `json:"SubAmount,omitempty" xml:"SubAmount,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetOcJusticeJudgementDocResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeJudgementDocResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetCaseFlow(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.CaseFlow = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetCaseNum(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.CaseNum = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetCaseType(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.CaseType = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetCauseAction(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.CauseAction = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetCourt(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.Court = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetDefendant(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.Defendant = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetJudgeDate(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.JudgeDate = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetJudgeResult(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.JudgeResult = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetJudgeType(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.JudgeType = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetParty(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.Party = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetPlaintiff(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.Plaintiff = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetPublicDate(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.PublicDate = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetRole(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.Role = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetSubAmount(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.SubAmount = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponseBodyData) SetTitle(v string) *GetOcJusticeJudgementDocResponseBodyData {
	s.Title = &v
	return s
}

type GetOcJusticeJudgementDocResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcJusticeJudgementDocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcJusticeJudgementDocResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeJudgementDocResponse) GoString() string {
	return s.String()
}

func (s *GetOcJusticeJudgementDocResponse) SetHeaders(v map[string]*string) *GetOcJusticeJudgementDocResponse {
	s.Headers = v
	return s
}

func (s *GetOcJusticeJudgementDocResponse) SetStatusCode(v int32) *GetOcJusticeJudgementDocResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcJusticeJudgementDocResponse) SetBody(v *GetOcJusticeJudgementDocResponseBody) *GetOcJusticeJudgementDocResponse {
	s.Body = v
	return s
}

type GetOcJusticeLimitHighConsumeRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcJusticeLimitHighConsumeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeLimitHighConsumeRequest) GoString() string {
	return s.String()
}

func (s *GetOcJusticeLimitHighConsumeRequest) SetPageNo(v int32) *GetOcJusticeLimitHighConsumeRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeRequest) SetPageSize(v int32) *GetOcJusticeLimitHighConsumeRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeRequest) SetRequestId(v string) *GetOcJusticeLimitHighConsumeRequest {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeRequest) SetSearchKey(v string) *GetOcJusticeLimitHighConsumeRequest {
	s.SearchKey = &v
	return s
}

type GetOcJusticeLimitHighConsumeResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcJusticeLimitHighConsumeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                          `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                          `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                          `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcJusticeLimitHighConsumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeLimitHighConsumeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcJusticeLimitHighConsumeResponseBody) SetCode(v string) *GetOcJusticeLimitHighConsumeResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBody) SetData(v []*GetOcJusticeLimitHighConsumeResponseBodyData) *GetOcJusticeLimitHighConsumeResponseBody {
	s.Data = v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBody) SetMessage(v string) *GetOcJusticeLimitHighConsumeResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBody) SetPageIndex(v int32) *GetOcJusticeLimitHighConsumeResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBody) SetPageNum(v int32) *GetOcJusticeLimitHighConsumeResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBody) SetRequestId(v string) *GetOcJusticeLimitHighConsumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBody) SetSuccess(v bool) *GetOcJusticeLimitHighConsumeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBody) SetTotalNum(v int32) *GetOcJusticeLimitHighConsumeResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcJusticeLimitHighConsumeResponseBodyData struct {
	CaseNum            *string `json:"CaseNum,omitempty" xml:"CaseNum,omitempty"`
	CauseAction        *string `json:"CauseAction,omitempty" xml:"CauseAction,omitempty"`
	CompanyName        *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	Court              *string `json:"Court,omitempty" xml:"Court,omitempty"`
	ExecutionApplicant *string `json:"ExecutionApplicant,omitempty" xml:"ExecutionApplicant,omitempty"`
	FilingDate         *string `json:"FilingDate,omitempty" xml:"FilingDate,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PublishDate        *string `json:"PublishDate,omitempty" xml:"PublishDate,omitempty"`
}

func (s GetOcJusticeLimitHighConsumeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeLimitHighConsumeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcJusticeLimitHighConsumeResponseBodyData) SetCaseNum(v string) *GetOcJusticeLimitHighConsumeResponseBodyData {
	s.CaseNum = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBodyData) SetCauseAction(v string) *GetOcJusticeLimitHighConsumeResponseBodyData {
	s.CauseAction = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBodyData) SetCompanyName(v string) *GetOcJusticeLimitHighConsumeResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBodyData) SetCourt(v string) *GetOcJusticeLimitHighConsumeResponseBodyData {
	s.Court = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBodyData) SetExecutionApplicant(v string) *GetOcJusticeLimitHighConsumeResponseBodyData {
	s.ExecutionApplicant = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBodyData) SetFilingDate(v string) *GetOcJusticeLimitHighConsumeResponseBodyData {
	s.FilingDate = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBodyData) SetName(v string) *GetOcJusticeLimitHighConsumeResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponseBodyData) SetPublishDate(v string) *GetOcJusticeLimitHighConsumeResponseBodyData {
	s.PublishDate = &v
	return s
}

type GetOcJusticeLimitHighConsumeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcJusticeLimitHighConsumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcJusticeLimitHighConsumeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeLimitHighConsumeResponse) GoString() string {
	return s.String()
}

func (s *GetOcJusticeLimitHighConsumeResponse) SetHeaders(v map[string]*string) *GetOcJusticeLimitHighConsumeResponse {
	s.Headers = v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponse) SetStatusCode(v int32) *GetOcJusticeLimitHighConsumeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcJusticeLimitHighConsumeResponse) SetBody(v *GetOcJusticeLimitHighConsumeResponseBody) *GetOcJusticeLimitHighConsumeResponse {
	s.Body = v
	return s
}

type GetOcJusticeTerminalCaseRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcJusticeTerminalCaseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeTerminalCaseRequest) GoString() string {
	return s.String()
}

func (s *GetOcJusticeTerminalCaseRequest) SetPageNo(v int32) *GetOcJusticeTerminalCaseRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcJusticeTerminalCaseRequest) SetPageSize(v int32) *GetOcJusticeTerminalCaseRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcJusticeTerminalCaseRequest) SetSearchKey(v string) *GetOcJusticeTerminalCaseRequest {
	s.SearchKey = &v
	return s
}

type GetOcJusticeTerminalCaseResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcJusticeTerminalCaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                      `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                      `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                      `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcJusticeTerminalCaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeTerminalCaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcJusticeTerminalCaseResponseBody) SetCode(v string) *GetOcJusticeTerminalCaseResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBody) SetData(v []*GetOcJusticeTerminalCaseResponseBodyData) *GetOcJusticeTerminalCaseResponseBody {
	s.Data = v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBody) SetMessage(v string) *GetOcJusticeTerminalCaseResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBody) SetPageIndex(v int32) *GetOcJusticeTerminalCaseResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBody) SetPageNum(v int32) *GetOcJusticeTerminalCaseResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBody) SetRequestId(v string) *GetOcJusticeTerminalCaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBody) SetSuccess(v bool) *GetOcJusticeTerminalCaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBody) SetTotalNum(v int32) *GetOcJusticeTerminalCaseResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcJusticeTerminalCaseResponseBodyData struct {
	CaseNum           *string `json:"CaseNum,omitempty" xml:"CaseNum,omitempty"`
	Court             *string `json:"Court,omitempty" xml:"Court,omitempty"`
	EntName           *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	ExecAmount        *string `json:"ExecAmount,omitempty" xml:"ExecAmount,omitempty"`
	FailPerformAmount *string `json:"FailPerformAmount,omitempty" xml:"FailPerformAmount,omitempty"`
	FilingDate        *string `json:"FilingDate,omitempty" xml:"FilingDate,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TerminalNum       *string `json:"TerminalNum,omitempty" xml:"TerminalNum,omitempty"`
	TerminateDate     *string `json:"TerminateDate,omitempty" xml:"TerminateDate,omitempty"`
}

func (s GetOcJusticeTerminalCaseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeTerminalCaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcJusticeTerminalCaseResponseBodyData) SetCaseNum(v string) *GetOcJusticeTerminalCaseResponseBodyData {
	s.CaseNum = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBodyData) SetCourt(v string) *GetOcJusticeTerminalCaseResponseBodyData {
	s.Court = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBodyData) SetEntName(v string) *GetOcJusticeTerminalCaseResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBodyData) SetExecAmount(v string) *GetOcJusticeTerminalCaseResponseBodyData {
	s.ExecAmount = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBodyData) SetFailPerformAmount(v string) *GetOcJusticeTerminalCaseResponseBodyData {
	s.FailPerformAmount = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBodyData) SetFilingDate(v string) *GetOcJusticeTerminalCaseResponseBodyData {
	s.FilingDate = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBodyData) SetName(v string) *GetOcJusticeTerminalCaseResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBodyData) SetTerminalNum(v string) *GetOcJusticeTerminalCaseResponseBodyData {
	s.TerminalNum = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponseBodyData) SetTerminateDate(v string) *GetOcJusticeTerminalCaseResponseBodyData {
	s.TerminateDate = &v
	return s
}

type GetOcJusticeTerminalCaseResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcJusticeTerminalCaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcJusticeTerminalCaseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcJusticeTerminalCaseResponse) GoString() string {
	return s.String()
}

func (s *GetOcJusticeTerminalCaseResponse) SetHeaders(v map[string]*string) *GetOcJusticeTerminalCaseResponse {
	s.Headers = v
	return s
}

func (s *GetOcJusticeTerminalCaseResponse) SetStatusCode(v int32) *GetOcJusticeTerminalCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcJusticeTerminalCaseResponse) SetBody(v *GetOcJusticeTerminalCaseResponseBody) *GetOcJusticeTerminalCaseResponse {
	s.Body = v
	return s
}

type GetOcListedCompanyRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcListedCompanyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcListedCompanyRequest) GoString() string {
	return s.String()
}

func (s *GetOcListedCompanyRequest) SetPageNo(v int32) *GetOcListedCompanyRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcListedCompanyRequest) SetPageSize(v int32) *GetOcListedCompanyRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcListedCompanyRequest) SetSearchKey(v string) *GetOcListedCompanyRequest {
	s.SearchKey = &v
	return s
}

type GetOcListedCompanyResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcListedCompanyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcListedCompanyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcListedCompanyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcListedCompanyResponseBody) SetCode(v string) *GetOcListedCompanyResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcListedCompanyResponseBody) SetData(v []*GetOcListedCompanyResponseBodyData) *GetOcListedCompanyResponseBody {
	s.Data = v
	return s
}

func (s *GetOcListedCompanyResponseBody) SetMessage(v string) *GetOcListedCompanyResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcListedCompanyResponseBody) SetPageIndex(v int32) *GetOcListedCompanyResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcListedCompanyResponseBody) SetPageNum(v int32) *GetOcListedCompanyResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcListedCompanyResponseBody) SetRequestId(v string) *GetOcListedCompanyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcListedCompanyResponseBody) SetSuccess(v bool) *GetOcListedCompanyResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcListedCompanyResponseBody) SetTotalNum(v int32) *GetOcListedCompanyResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcListedCompanyResponseBodyData struct {
	CirculationMarketValue *string `json:"CirculationMarketValue,omitempty" xml:"CirculationMarketValue,omitempty"`
	EntName                *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	EntNameEng             *string `json:"EntNameEng,omitempty" xml:"EntNameEng,omitempty"`
	ListDate               *string `json:"ListDate,omitempty" xml:"ListDate,omitempty"`
	SecuritiesCode         *string `json:"SecuritiesCode,omitempty" xml:"SecuritiesCode,omitempty"`
	SecuritiesMarket       *string `json:"SecuritiesMarket,omitempty" xml:"SecuritiesMarket,omitempty"`
	SecuritiesName         *string `json:"SecuritiesName,omitempty" xml:"SecuritiesName,omitempty"`
	TotalFlowShares        *string `json:"TotalFlowShares,omitempty" xml:"TotalFlowShares,omitempty"`
	TotalShares            *string `json:"TotalShares,omitempty" xml:"TotalShares,omitempty"`
}

func (s GetOcListedCompanyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcListedCompanyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcListedCompanyResponseBodyData) SetCirculationMarketValue(v string) *GetOcListedCompanyResponseBodyData {
	s.CirculationMarketValue = &v
	return s
}

func (s *GetOcListedCompanyResponseBodyData) SetEntName(v string) *GetOcListedCompanyResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcListedCompanyResponseBodyData) SetEntNameEng(v string) *GetOcListedCompanyResponseBodyData {
	s.EntNameEng = &v
	return s
}

func (s *GetOcListedCompanyResponseBodyData) SetListDate(v string) *GetOcListedCompanyResponseBodyData {
	s.ListDate = &v
	return s
}

func (s *GetOcListedCompanyResponseBodyData) SetSecuritiesCode(v string) *GetOcListedCompanyResponseBodyData {
	s.SecuritiesCode = &v
	return s
}

func (s *GetOcListedCompanyResponseBodyData) SetSecuritiesMarket(v string) *GetOcListedCompanyResponseBodyData {
	s.SecuritiesMarket = &v
	return s
}

func (s *GetOcListedCompanyResponseBodyData) SetSecuritiesName(v string) *GetOcListedCompanyResponseBodyData {
	s.SecuritiesName = &v
	return s
}

func (s *GetOcListedCompanyResponseBodyData) SetTotalFlowShares(v string) *GetOcListedCompanyResponseBodyData {
	s.TotalFlowShares = &v
	return s
}

func (s *GetOcListedCompanyResponseBodyData) SetTotalShares(v string) *GetOcListedCompanyResponseBodyData {
	s.TotalShares = &v
	return s
}

type GetOcListedCompanyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcListedCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcListedCompanyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcListedCompanyResponse) GoString() string {
	return s.String()
}

func (s *GetOcListedCompanyResponse) SetHeaders(v map[string]*string) *GetOcListedCompanyResponse {
	s.Headers = v
	return s
}

func (s *GetOcListedCompanyResponse) SetStatusCode(v int32) *GetOcListedCompanyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcListedCompanyResponse) SetBody(v *GetOcListedCompanyResponseBody) *GetOcListedCompanyResponse {
	s.Body = v
	return s
}

type GetOcNegativeAdminPunishmentRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcNegativeAdminPunishmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeAdminPunishmentRequest) GoString() string {
	return s.String()
}

func (s *GetOcNegativeAdminPunishmentRequest) SetPageNo(v int32) *GetOcNegativeAdminPunishmentRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentRequest) SetPageSize(v int32) *GetOcNegativeAdminPunishmentRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentRequest) SetSearchKey(v string) *GetOcNegativeAdminPunishmentRequest {
	s.SearchKey = &v
	return s
}

type GetOcNegativeAdminPunishmentResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcNegativeAdminPunishmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                          `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                          `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                          `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcNegativeAdminPunishmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeAdminPunishmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcNegativeAdminPunishmentResponseBody) SetCode(v string) *GetOcNegativeAdminPunishmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBody) SetData(v []*GetOcNegativeAdminPunishmentResponseBodyData) *GetOcNegativeAdminPunishmentResponseBody {
	s.Data = v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBody) SetMessage(v string) *GetOcNegativeAdminPunishmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBody) SetPageIndex(v int32) *GetOcNegativeAdminPunishmentResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBody) SetPageNum(v int32) *GetOcNegativeAdminPunishmentResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBody) SetRequestId(v string) *GetOcNegativeAdminPunishmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBody) SetSuccess(v bool) *GetOcNegativeAdminPunishmentResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBody) SetTotalNum(v int32) *GetOcNegativeAdminPunishmentResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcNegativeAdminPunishmentResponseBodyData struct {
	Department   *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EntName      *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	IllegalType  *string `json:"IllegalType,omitempty" xml:"IllegalType,omitempty"`
	LawBasis     *string `json:"LawBasis,omitempty" xml:"LawBasis,omitempty"`
	PublicDate   *string `json:"PublicDate,omitempty" xml:"PublicDate,omitempty"`
	PunishDate   *string `json:"PunishDate,omitempty" xml:"PunishDate,omitempty"`
	PunishNum    *string `json:"PunishNum,omitempty" xml:"PunishNum,omitempty"`
	PunishResult *string `json:"PunishResult,omitempty" xml:"PunishResult,omitempty"`
}

func (s GetOcNegativeAdminPunishmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeAdminPunishmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcNegativeAdminPunishmentResponseBodyData) SetDepartment(v string) *GetOcNegativeAdminPunishmentResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBodyData) SetEntName(v string) *GetOcNegativeAdminPunishmentResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBodyData) SetIllegalType(v string) *GetOcNegativeAdminPunishmentResponseBodyData {
	s.IllegalType = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBodyData) SetLawBasis(v string) *GetOcNegativeAdminPunishmentResponseBodyData {
	s.LawBasis = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBodyData) SetPublicDate(v string) *GetOcNegativeAdminPunishmentResponseBodyData {
	s.PublicDate = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBodyData) SetPunishDate(v string) *GetOcNegativeAdminPunishmentResponseBodyData {
	s.PunishDate = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBodyData) SetPunishNum(v string) *GetOcNegativeAdminPunishmentResponseBodyData {
	s.PunishNum = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponseBodyData) SetPunishResult(v string) *GetOcNegativeAdminPunishmentResponseBodyData {
	s.PunishResult = &v
	return s
}

type GetOcNegativeAdminPunishmentResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcNegativeAdminPunishmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcNegativeAdminPunishmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeAdminPunishmentResponse) GoString() string {
	return s.String()
}

func (s *GetOcNegativeAdminPunishmentResponse) SetHeaders(v map[string]*string) *GetOcNegativeAdminPunishmentResponse {
	s.Headers = v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponse) SetStatusCode(v int32) *GetOcNegativeAdminPunishmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcNegativeAdminPunishmentResponse) SetBody(v *GetOcNegativeAdminPunishmentResponseBody) *GetOcNegativeAdminPunishmentResponse {
	s.Body = v
	return s
}

type GetOcNegativeCustomsPunishmentRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcNegativeCustomsPunishmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeCustomsPunishmentRequest) GoString() string {
	return s.String()
}

func (s *GetOcNegativeCustomsPunishmentRequest) SetPageNo(v int32) *GetOcNegativeCustomsPunishmentRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentRequest) SetPageSize(v int32) *GetOcNegativeCustomsPunishmentRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentRequest) SetSearchKey(v string) *GetOcNegativeCustomsPunishmentRequest {
	s.SearchKey = &v
	return s
}

type GetOcNegativeCustomsPunishmentResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcNegativeCustomsPunishmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                            `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcNegativeCustomsPunishmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeCustomsPunishmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcNegativeCustomsPunishmentResponseBody) SetCode(v string) *GetOcNegativeCustomsPunishmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBody) SetData(v []*GetOcNegativeCustomsPunishmentResponseBodyData) *GetOcNegativeCustomsPunishmentResponseBody {
	s.Data = v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBody) SetMessage(v string) *GetOcNegativeCustomsPunishmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBody) SetPageIndex(v int32) *GetOcNegativeCustomsPunishmentResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBody) SetPageNum(v int32) *GetOcNegativeCustomsPunishmentResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBody) SetRequestId(v string) *GetOcNegativeCustomsPunishmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBody) SetSuccess(v bool) *GetOcNegativeCustomsPunishmentResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBody) SetTotalNum(v int32) *GetOcNegativeCustomsPunishmentResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcNegativeCustomsPunishmentResponseBodyData struct {
	Basis      *string `json:"Basis,omitempty" xml:"Basis,omitempty"`
	CaseNo     *string `json:"CaseNo,omitempty" xml:"CaseNo,omitempty"`
	Customs    *string `json:"Customs,omitempty" xml:"Customs,omitempty"`
	CustomsNo  *string `json:"CustomsNo,omitempty" xml:"CustomsNo,omitempty"`
	LegalName  *string `json:"LegalName,omitempty" xml:"LegalName,omitempty"`
	PunishDate *string `json:"PunishDate,omitempty" xml:"PunishDate,omitempty"`
	PunishType *string `json:"PunishType,omitempty" xml:"PunishType,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetOcNegativeCustomsPunishmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeCustomsPunishmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcNegativeCustomsPunishmentResponseBodyData) SetBasis(v string) *GetOcNegativeCustomsPunishmentResponseBodyData {
	s.Basis = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBodyData) SetCaseNo(v string) *GetOcNegativeCustomsPunishmentResponseBodyData {
	s.CaseNo = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBodyData) SetCustoms(v string) *GetOcNegativeCustomsPunishmentResponseBodyData {
	s.Customs = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBodyData) SetCustomsNo(v string) *GetOcNegativeCustomsPunishmentResponseBodyData {
	s.CustomsNo = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBodyData) SetLegalName(v string) *GetOcNegativeCustomsPunishmentResponseBodyData {
	s.LegalName = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBodyData) SetPunishDate(v string) *GetOcNegativeCustomsPunishmentResponseBodyData {
	s.PunishDate = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBodyData) SetPunishType(v string) *GetOcNegativeCustomsPunishmentResponseBodyData {
	s.PunishType = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponseBodyData) SetTitle(v string) *GetOcNegativeCustomsPunishmentResponseBodyData {
	s.Title = &v
	return s
}

type GetOcNegativeCustomsPunishmentResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcNegativeCustomsPunishmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcNegativeCustomsPunishmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeCustomsPunishmentResponse) GoString() string {
	return s.String()
}

func (s *GetOcNegativeCustomsPunishmentResponse) SetHeaders(v map[string]*string) *GetOcNegativeCustomsPunishmentResponse {
	s.Headers = v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponse) SetStatusCode(v int32) *GetOcNegativeCustomsPunishmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcNegativeCustomsPunishmentResponse) SetBody(v *GetOcNegativeCustomsPunishmentResponseBody) *GetOcNegativeCustomsPunishmentResponse {
	s.Body = v
	return s
}

type GetOcNegativeEnvironmentPunishmentRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcNegativeEnvironmentPunishmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeEnvironmentPunishmentRequest) GoString() string {
	return s.String()
}

func (s *GetOcNegativeEnvironmentPunishmentRequest) SetPageNo(v int32) *GetOcNegativeEnvironmentPunishmentRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentRequest) SetPageSize(v int32) *GetOcNegativeEnvironmentPunishmentRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentRequest) SetSearchKey(v string) *GetOcNegativeEnvironmentPunishmentRequest {
	s.SearchKey = &v
	return s
}

type GetOcNegativeEnvironmentPunishmentResponseBody struct {
	Code      *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcNegativeEnvironmentPunishmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                                `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcNegativeEnvironmentPunishmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeEnvironmentPunishmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBody) SetCode(v string) *GetOcNegativeEnvironmentPunishmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBody) SetData(v []*GetOcNegativeEnvironmentPunishmentResponseBodyData) *GetOcNegativeEnvironmentPunishmentResponseBody {
	s.Data = v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBody) SetMessage(v string) *GetOcNegativeEnvironmentPunishmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBody) SetPageIndex(v int32) *GetOcNegativeEnvironmentPunishmentResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBody) SetPageNum(v int32) *GetOcNegativeEnvironmentPunishmentResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBody) SetRequestId(v string) *GetOcNegativeEnvironmentPunishmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBody) SetSuccess(v bool) *GetOcNegativeEnvironmentPunishmentResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBody) SetTotalNum(v int32) *GetOcNegativeEnvironmentPunishmentResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcNegativeEnvironmentPunishmentResponseBodyData struct {
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EntName       *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	ExecStatus    *string `json:"ExecStatus,omitempty" xml:"ExecStatus,omitempty"`
	PunishBasis   *string `json:"PunishBasis,omitempty" xml:"PunishBasis,omitempty"`
	PunishContent *string `json:"PunishContent,omitempty" xml:"PunishContent,omitempty"`
	PunishDate    *string `json:"PunishDate,omitempty" xml:"PunishDate,omitempty"`
	PunishLaw     *string `json:"PunishLaw,omitempty" xml:"PunishLaw,omitempty"`
	PunishNum     *string `json:"PunishNum,omitempty" xml:"PunishNum,omitempty"`
	PunishRes     *string `json:"PunishRes,omitempty" xml:"PunishRes,omitempty"`
}

func (s GetOcNegativeEnvironmentPunishmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeEnvironmentPunishmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBodyData) SetDepartment(v string) *GetOcNegativeEnvironmentPunishmentResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBodyData) SetEntName(v string) *GetOcNegativeEnvironmentPunishmentResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBodyData) SetExecStatus(v string) *GetOcNegativeEnvironmentPunishmentResponseBodyData {
	s.ExecStatus = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBodyData) SetPunishBasis(v string) *GetOcNegativeEnvironmentPunishmentResponseBodyData {
	s.PunishBasis = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBodyData) SetPunishContent(v string) *GetOcNegativeEnvironmentPunishmentResponseBodyData {
	s.PunishContent = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBodyData) SetPunishDate(v string) *GetOcNegativeEnvironmentPunishmentResponseBodyData {
	s.PunishDate = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBodyData) SetPunishLaw(v string) *GetOcNegativeEnvironmentPunishmentResponseBodyData {
	s.PunishLaw = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBodyData) SetPunishNum(v string) *GetOcNegativeEnvironmentPunishmentResponseBodyData {
	s.PunishNum = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponseBodyData) SetPunishRes(v string) *GetOcNegativeEnvironmentPunishmentResponseBodyData {
	s.PunishRes = &v
	return s
}

type GetOcNegativeEnvironmentPunishmentResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcNegativeEnvironmentPunishmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcNegativeEnvironmentPunishmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeEnvironmentPunishmentResponse) GoString() string {
	return s.String()
}

func (s *GetOcNegativeEnvironmentPunishmentResponse) SetHeaders(v map[string]*string) *GetOcNegativeEnvironmentPunishmentResponse {
	s.Headers = v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponse) SetStatusCode(v int32) *GetOcNegativeEnvironmentPunishmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcNegativeEnvironmentPunishmentResponse) SetBody(v *GetOcNegativeEnvironmentPunishmentResponseBody) *GetOcNegativeEnvironmentPunishmentResponse {
	s.Body = v
	return s
}

type GetOcNegativeFoodDrugPunishmentRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcNegativeFoodDrugPunishmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeFoodDrugPunishmentRequest) GoString() string {
	return s.String()
}

func (s *GetOcNegativeFoodDrugPunishmentRequest) SetPageNo(v int32) *GetOcNegativeFoodDrugPunishmentRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentRequest) SetPageSize(v int32) *GetOcNegativeFoodDrugPunishmentRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentRequest) SetSearchKey(v string) *GetOcNegativeFoodDrugPunishmentRequest {
	s.SearchKey = &v
	return s
}

type GetOcNegativeFoodDrugPunishmentResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcNegativeFoodDrugPunishmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                             `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                             `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcNegativeFoodDrugPunishmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeFoodDrugPunishmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBody) SetCode(v string) *GetOcNegativeFoodDrugPunishmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBody) SetData(v []*GetOcNegativeFoodDrugPunishmentResponseBodyData) *GetOcNegativeFoodDrugPunishmentResponseBody {
	s.Data = v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBody) SetMessage(v string) *GetOcNegativeFoodDrugPunishmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBody) SetPageIndex(v int32) *GetOcNegativeFoodDrugPunishmentResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBody) SetPageNum(v int32) *GetOcNegativeFoodDrugPunishmentResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBody) SetRequestId(v string) *GetOcNegativeFoodDrugPunishmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBody) SetSuccess(v bool) *GetOcNegativeFoodDrugPunishmentResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBody) SetTotalNum(v int32) *GetOcNegativeFoodDrugPunishmentResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcNegativeFoodDrugPunishmentResponseBodyData struct {
	Department   *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EntName      *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	IllegalType  *string `json:"IllegalType,omitempty" xml:"IllegalType,omitempty"`
	LawBasis     *string `json:"LawBasis,omitempty" xml:"LawBasis,omitempty"`
	PublicDate   *string `json:"PublicDate,omitempty" xml:"PublicDate,omitempty"`
	PunishDate   *string `json:"PunishDate,omitempty" xml:"PunishDate,omitempty"`
	PunishNum    *string `json:"PunishNum,omitempty" xml:"PunishNum,omitempty"`
	PunishResult *string `json:"PunishResult,omitempty" xml:"PunishResult,omitempty"`
}

func (s GetOcNegativeFoodDrugPunishmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeFoodDrugPunishmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBodyData) SetDepartment(v string) *GetOcNegativeFoodDrugPunishmentResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBodyData) SetEntName(v string) *GetOcNegativeFoodDrugPunishmentResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBodyData) SetIllegalType(v string) *GetOcNegativeFoodDrugPunishmentResponseBodyData {
	s.IllegalType = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBodyData) SetLawBasis(v string) *GetOcNegativeFoodDrugPunishmentResponseBodyData {
	s.LawBasis = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBodyData) SetPublicDate(v string) *GetOcNegativeFoodDrugPunishmentResponseBodyData {
	s.PublicDate = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBodyData) SetPunishDate(v string) *GetOcNegativeFoodDrugPunishmentResponseBodyData {
	s.PunishDate = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBodyData) SetPunishNum(v string) *GetOcNegativeFoodDrugPunishmentResponseBodyData {
	s.PunishNum = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponseBodyData) SetPunishResult(v string) *GetOcNegativeFoodDrugPunishmentResponseBodyData {
	s.PunishResult = &v
	return s
}

type GetOcNegativeFoodDrugPunishmentResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcNegativeFoodDrugPunishmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcNegativeFoodDrugPunishmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeFoodDrugPunishmentResponse) GoString() string {
	return s.String()
}

func (s *GetOcNegativeFoodDrugPunishmentResponse) SetHeaders(v map[string]*string) *GetOcNegativeFoodDrugPunishmentResponse {
	s.Headers = v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponse) SetStatusCode(v int32) *GetOcNegativeFoodDrugPunishmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcNegativeFoodDrugPunishmentResponse) SetBody(v *GetOcNegativeFoodDrugPunishmentResponseBody) *GetOcNegativeFoodDrugPunishmentResponse {
	s.Body = v
	return s
}

type GetOcNegativeQualityPunishmentRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcNegativeQualityPunishmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeQualityPunishmentRequest) GoString() string {
	return s.String()
}

func (s *GetOcNegativeQualityPunishmentRequest) SetPageNo(v int32) *GetOcNegativeQualityPunishmentRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentRequest) SetPageSize(v int32) *GetOcNegativeQualityPunishmentRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentRequest) SetSearchKey(v string) *GetOcNegativeQualityPunishmentRequest {
	s.SearchKey = &v
	return s
}

type GetOcNegativeQualityPunishmentResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcNegativeQualityPunishmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                            `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcNegativeQualityPunishmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeQualityPunishmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcNegativeQualityPunishmentResponseBody) SetCode(v string) *GetOcNegativeQualityPunishmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBody) SetData(v []*GetOcNegativeQualityPunishmentResponseBodyData) *GetOcNegativeQualityPunishmentResponseBody {
	s.Data = v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBody) SetMessage(v string) *GetOcNegativeQualityPunishmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBody) SetPageIndex(v int32) *GetOcNegativeQualityPunishmentResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBody) SetPageNum(v int32) *GetOcNegativeQualityPunishmentResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBody) SetRequestId(v string) *GetOcNegativeQualityPunishmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBody) SetSuccess(v bool) *GetOcNegativeQualityPunishmentResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBody) SetTotalNum(v int32) *GetOcNegativeQualityPunishmentResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcNegativeQualityPunishmentResponseBodyData struct {
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EntName     *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	EventDate   *string `json:"EventDate,omitempty" xml:"EventDate,omitempty"`
	EventResult *string `json:"EventResult,omitempty" xml:"EventResult,omitempty"`
	PubDate     *string `json:"PubDate,omitempty" xml:"PubDate,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetOcNegativeQualityPunishmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeQualityPunishmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcNegativeQualityPunishmentResponseBodyData) SetDepartment(v string) *GetOcNegativeQualityPunishmentResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBodyData) SetEntName(v string) *GetOcNegativeQualityPunishmentResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBodyData) SetEventDate(v string) *GetOcNegativeQualityPunishmentResponseBodyData {
	s.EventDate = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBodyData) SetEventResult(v string) *GetOcNegativeQualityPunishmentResponseBodyData {
	s.EventResult = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBodyData) SetPubDate(v string) *GetOcNegativeQualityPunishmentResponseBodyData {
	s.PubDate = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponseBodyData) SetTitle(v string) *GetOcNegativeQualityPunishmentResponseBodyData {
	s.Title = &v
	return s
}

type GetOcNegativeQualityPunishmentResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcNegativeQualityPunishmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcNegativeQualityPunishmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcNegativeQualityPunishmentResponse) GoString() string {
	return s.String()
}

func (s *GetOcNegativeQualityPunishmentResponse) SetHeaders(v map[string]*string) *GetOcNegativeQualityPunishmentResponse {
	s.Headers = v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponse) SetStatusCode(v int32) *GetOcNegativeQualityPunishmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcNegativeQualityPunishmentResponse) SetBody(v *GetOcNegativeQualityPunishmentResponseBody) *GetOcNegativeQualityPunishmentResponse {
	s.Body = v
	return s
}

type GetOcOperationBiddingRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcOperationBiddingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationBiddingRequest) GoString() string {
	return s.String()
}

func (s *GetOcOperationBiddingRequest) SetPageNo(v int32) *GetOcOperationBiddingRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcOperationBiddingRequest) SetPageSize(v int32) *GetOcOperationBiddingRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcOperationBiddingRequest) SetSearchKey(v string) *GetOcOperationBiddingRequest {
	s.SearchKey = &v
	return s
}

type GetOcOperationBiddingResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcOperationBiddingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                   `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                   `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcOperationBiddingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationBiddingResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcOperationBiddingResponseBody) SetCode(v string) *GetOcOperationBiddingResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcOperationBiddingResponseBody) SetData(v []*GetOcOperationBiddingResponseBodyData) *GetOcOperationBiddingResponseBody {
	s.Data = v
	return s
}

func (s *GetOcOperationBiddingResponseBody) SetMessage(v string) *GetOcOperationBiddingResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcOperationBiddingResponseBody) SetPageIndex(v int32) *GetOcOperationBiddingResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcOperationBiddingResponseBody) SetPageNum(v int32) *GetOcOperationBiddingResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcOperationBiddingResponseBody) SetRequestId(v string) *GetOcOperationBiddingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcOperationBiddingResponseBody) SetSuccess(v bool) *GetOcOperationBiddingResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcOperationBiddingResponseBody) SetTotalNum(v int32) *GetOcOperationBiddingResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcOperationBiddingResponseBodyData struct {
	AgentEntName  *string `json:"AgentEntName,omitempty" xml:"AgentEntName,omitempty"`
	BidIndustry   *string `json:"BidIndustry,omitempty" xml:"BidIndustry,omitempty"`
	BidTitle      *string `json:"BidTitle,omitempty" xml:"BidTitle,omitempty"`
	BidType       *string `json:"BidType,omitempty" xml:"BidType,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EntName       *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	InfoType      *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	OpeningTime   *string `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	ProjectAmount *string `json:"ProjectAmount,omitempty" xml:"ProjectAmount,omitempty"`
	ProjectName   *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectNum    *string `json:"ProjectNum,omitempty" xml:"ProjectNum,omitempty"`
	PublicDate    *string `json:"PublicDate,omitempty" xml:"PublicDate,omitempty"`
	RegionName    *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	TenderEntName *string `json:"TenderEntName,omitempty" xml:"TenderEntName,omitempty"`
	WinnerEntName *string `json:"WinnerEntName,omitempty" xml:"WinnerEntName,omitempty"`
}

func (s GetOcOperationBiddingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationBiddingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcOperationBiddingResponseBodyData) SetAgentEntName(v string) *GetOcOperationBiddingResponseBodyData {
	s.AgentEntName = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetBidIndustry(v string) *GetOcOperationBiddingResponseBodyData {
	s.BidIndustry = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetBidTitle(v string) *GetOcOperationBiddingResponseBodyData {
	s.BidTitle = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetBidType(v string) *GetOcOperationBiddingResponseBodyData {
	s.BidType = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetContent(v string) *GetOcOperationBiddingResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetEntName(v string) *GetOcOperationBiddingResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetInfoType(v string) *GetOcOperationBiddingResponseBodyData {
	s.InfoType = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetOpeningTime(v string) *GetOcOperationBiddingResponseBodyData {
	s.OpeningTime = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetProjectAmount(v string) *GetOcOperationBiddingResponseBodyData {
	s.ProjectAmount = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetProjectName(v string) *GetOcOperationBiddingResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetProjectNum(v string) *GetOcOperationBiddingResponseBodyData {
	s.ProjectNum = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetPublicDate(v string) *GetOcOperationBiddingResponseBodyData {
	s.PublicDate = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetRegionName(v string) *GetOcOperationBiddingResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetSubType(v string) *GetOcOperationBiddingResponseBodyData {
	s.SubType = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetTenderEntName(v string) *GetOcOperationBiddingResponseBodyData {
	s.TenderEntName = &v
	return s
}

func (s *GetOcOperationBiddingResponseBodyData) SetWinnerEntName(v string) *GetOcOperationBiddingResponseBodyData {
	s.WinnerEntName = &v
	return s
}

type GetOcOperationBiddingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcOperationBiddingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcOperationBiddingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationBiddingResponse) GoString() string {
	return s.String()
}

func (s *GetOcOperationBiddingResponse) SetHeaders(v map[string]*string) *GetOcOperationBiddingResponse {
	s.Headers = v
	return s
}

func (s *GetOcOperationBiddingResponse) SetStatusCode(v int32) *GetOcOperationBiddingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcOperationBiddingResponse) SetBody(v *GetOcOperationBiddingResponseBody) *GetOcOperationBiddingResponse {
	s.Body = v
	return s
}

type GetOcOperationCustomsRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcOperationCustomsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationCustomsRequest) GoString() string {
	return s.String()
}

func (s *GetOcOperationCustomsRequest) SetPageNo(v int32) *GetOcOperationCustomsRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcOperationCustomsRequest) SetPageSize(v int32) *GetOcOperationCustomsRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcOperationCustomsRequest) SetSearchKey(v string) *GetOcOperationCustomsRequest {
	s.SearchKey = &v
	return s
}

type GetOcOperationCustomsResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcOperationCustomsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                   `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                   `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcOperationCustomsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationCustomsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcOperationCustomsResponseBody) SetCode(v string) *GetOcOperationCustomsResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcOperationCustomsResponseBody) SetData(v []*GetOcOperationCustomsResponseBodyData) *GetOcOperationCustomsResponseBody {
	s.Data = v
	return s
}

func (s *GetOcOperationCustomsResponseBody) SetMessage(v string) *GetOcOperationCustomsResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcOperationCustomsResponseBody) SetPageIndex(v int32) *GetOcOperationCustomsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcOperationCustomsResponseBody) SetPageNum(v int32) *GetOcOperationCustomsResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcOperationCustomsResponseBody) SetRequestId(v string) *GetOcOperationCustomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcOperationCustomsResponseBody) SetSuccess(v bool) *GetOcOperationCustomsResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcOperationCustomsResponseBody) SetTotalNum(v int32) *GetOcOperationCustomsResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcOperationCustomsResponseBodyData struct {
	AdminRegionName *string `json:"AdminRegionName,omitempty" xml:"AdminRegionName,omitempty"`
	AnnualReport    *string `json:"AnnualReport,omitempty" xml:"AnnualReport,omitempty"`
	BusinessCate    *string `json:"BusinessCate,omitempty" xml:"BusinessCate,omitempty"`
	CancelFlag      *string `json:"CancelFlag,omitempty" xml:"CancelFlag,omitempty"`
	CreditLevelsNew *string `json:"CreditLevelsNew,omitempty" xml:"CreditLevelsNew,omitempty"`
	CustomsNum      *string `json:"CustomsNum,omitempty" xml:"CustomsNum,omitempty"`
	CustomsReg      *string `json:"CustomsReg,omitempty" xml:"CustomsReg,omitempty"`
	EcoRegionName   *string `json:"EcoRegionName,omitempty" xml:"EcoRegionName,omitempty"`
	ElectType       *string `json:"ElectType,omitempty" xml:"ElectType,omitempty"`
	EntName         *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	IdentCode       *string `json:"IdentCode,omitempty" xml:"IdentCode,omitempty"`
	IdentDate       *string `json:"IdentDate,omitempty" xml:"IdentDate,omitempty"`
	IndustryType    *string `json:"IndustryType,omitempty" xml:"IndustryType,omitempty"`
	RegDate         *string `json:"RegDate,omitempty" xml:"RegDate,omitempty"`
	SpecialArea     *string `json:"SpecialArea,omitempty" xml:"SpecialArea,omitempty"`
	ValidDate       *string `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
}

func (s GetOcOperationCustomsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationCustomsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcOperationCustomsResponseBodyData) SetAdminRegionName(v string) *GetOcOperationCustomsResponseBodyData {
	s.AdminRegionName = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetAnnualReport(v string) *GetOcOperationCustomsResponseBodyData {
	s.AnnualReport = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetBusinessCate(v string) *GetOcOperationCustomsResponseBodyData {
	s.BusinessCate = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetCancelFlag(v string) *GetOcOperationCustomsResponseBodyData {
	s.CancelFlag = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetCreditLevelsNew(v string) *GetOcOperationCustomsResponseBodyData {
	s.CreditLevelsNew = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetCustomsNum(v string) *GetOcOperationCustomsResponseBodyData {
	s.CustomsNum = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetCustomsReg(v string) *GetOcOperationCustomsResponseBodyData {
	s.CustomsReg = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetEcoRegionName(v string) *GetOcOperationCustomsResponseBodyData {
	s.EcoRegionName = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetElectType(v string) *GetOcOperationCustomsResponseBodyData {
	s.ElectType = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetEntName(v string) *GetOcOperationCustomsResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetIdentCode(v string) *GetOcOperationCustomsResponseBodyData {
	s.IdentCode = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetIdentDate(v string) *GetOcOperationCustomsResponseBodyData {
	s.IdentDate = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetIndustryType(v string) *GetOcOperationCustomsResponseBodyData {
	s.IndustryType = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetRegDate(v string) *GetOcOperationCustomsResponseBodyData {
	s.RegDate = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetSpecialArea(v string) *GetOcOperationCustomsResponseBodyData {
	s.SpecialArea = &v
	return s
}

func (s *GetOcOperationCustomsResponseBodyData) SetValidDate(v string) *GetOcOperationCustomsResponseBodyData {
	s.ValidDate = &v
	return s
}

type GetOcOperationCustomsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcOperationCustomsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcOperationCustomsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationCustomsResponse) GoString() string {
	return s.String()
}

func (s *GetOcOperationCustomsResponse) SetHeaders(v map[string]*string) *GetOcOperationCustomsResponse {
	s.Headers = v
	return s
}

func (s *GetOcOperationCustomsResponse) SetStatusCode(v int32) *GetOcOperationCustomsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcOperationCustomsResponse) SetBody(v *GetOcOperationCustomsResponseBody) *GetOcOperationCustomsResponse {
	s.Body = v
	return s
}

type GetOcOperationPurchaseLandRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcOperationPurchaseLandRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationPurchaseLandRequest) GoString() string {
	return s.String()
}

func (s *GetOcOperationPurchaseLandRequest) SetPageNo(v int32) *GetOcOperationPurchaseLandRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcOperationPurchaseLandRequest) SetPageSize(v int32) *GetOcOperationPurchaseLandRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcOperationPurchaseLandRequest) SetSearchKey(v string) *GetOcOperationPurchaseLandRequest {
	s.SearchKey = &v
	return s
}

type GetOcOperationPurchaseLandResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcOperationPurchaseLandResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                        `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcOperationPurchaseLandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationPurchaseLandResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcOperationPurchaseLandResponseBody) SetCode(v string) *GetOcOperationPurchaseLandResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBody) SetData(v []*GetOcOperationPurchaseLandResponseBodyData) *GetOcOperationPurchaseLandResponseBody {
	s.Data = v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBody) SetMessage(v string) *GetOcOperationPurchaseLandResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBody) SetPageIndex(v int32) *GetOcOperationPurchaseLandResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBody) SetPageNum(v int32) *GetOcOperationPurchaseLandResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBody) SetRequestId(v string) *GetOcOperationPurchaseLandResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBody) SetSuccess(v bool) *GetOcOperationPurchaseLandResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBody) SetTotalNum(v int32) *GetOcOperationPurchaseLandResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcOperationPurchaseLandResponseBodyData struct {
	Area                     *string `json:"Area,omitempty" xml:"Area,omitempty"`
	Department               *string `json:"Department,omitempty" xml:"Department,omitempty"`
	ElectronicNo             *string `json:"ElectronicNo,omitempty" xml:"ElectronicNo,omitempty"`
	EntName                  *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	Industry                 *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	LandLevel                *string `json:"LandLevel,omitempty" xml:"LandLevel,omitempty"`
	LandSource               *string `json:"LandSource,omitempty" xml:"LandSource,omitempty"`
	LandUse                  *string `json:"LandUse,omitempty" xml:"LandUse,omitempty"`
	Location                 *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Price                    *string `json:"Price,omitempty" xml:"Price,omitempty"`
	ProjectName              *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	PromiseDeliveryDate      *string `json:"PromiseDeliveryDate,omitempty" xml:"PromiseDeliveryDate,omitempty"`
	PromiseEndDate           *string `json:"PromiseEndDate,omitempty" xml:"PromiseEndDate,omitempty"`
	PromiseStartDate         *string `json:"PromiseStartDate,omitempty" xml:"PromiseStartDate,omitempty"`
	RegionName               *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	ReleaseDate              *string `json:"ReleaseDate,omitempty" xml:"ReleaseDate,omitempty"`
	SigningMode              *string `json:"SigningMode,omitempty" xml:"SigningMode,omitempty"`
	UseYear                  *string `json:"UseYear,omitempty" xml:"UseYear,omitempty"`
	VolumeFractionLowerBound *string `json:"VolumeFractionLowerBound,omitempty" xml:"VolumeFractionLowerBound,omitempty"`
	VolumeFractionUpperBound *string `json:"VolumeFractionUpperBound,omitempty" xml:"VolumeFractionUpperBound,omitempty"`
}

func (s GetOcOperationPurchaseLandResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationPurchaseLandResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetArea(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.Area = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetDepartment(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetElectronicNo(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.ElectronicNo = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetEntName(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetIndustry(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.Industry = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetLandLevel(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.LandLevel = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetLandSource(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.LandSource = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetLandUse(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.LandUse = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetLocation(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetPrice(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.Price = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetProjectName(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetPromiseDeliveryDate(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.PromiseDeliveryDate = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetPromiseEndDate(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.PromiseEndDate = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetPromiseStartDate(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.PromiseStartDate = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetRegionName(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetReleaseDate(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.ReleaseDate = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetSigningMode(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.SigningMode = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetUseYear(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.UseYear = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetVolumeFractionLowerBound(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.VolumeFractionLowerBound = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponseBodyData) SetVolumeFractionUpperBound(v string) *GetOcOperationPurchaseLandResponseBodyData {
	s.VolumeFractionUpperBound = &v
	return s
}

type GetOcOperationPurchaseLandResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcOperationPurchaseLandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcOperationPurchaseLandResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationPurchaseLandResponse) GoString() string {
	return s.String()
}

func (s *GetOcOperationPurchaseLandResponse) SetHeaders(v map[string]*string) *GetOcOperationPurchaseLandResponse {
	s.Headers = v
	return s
}

func (s *GetOcOperationPurchaseLandResponse) SetStatusCode(v int32) *GetOcOperationPurchaseLandResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcOperationPurchaseLandResponse) SetBody(v *GetOcOperationPurchaseLandResponseBody) *GetOcOperationPurchaseLandResponse {
	s.Body = v
	return s
}

type GetOcOperationRecruitmentRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcOperationRecruitmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationRecruitmentRequest) GoString() string {
	return s.String()
}

func (s *GetOcOperationRecruitmentRequest) SetPageNo(v int32) *GetOcOperationRecruitmentRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcOperationRecruitmentRequest) SetPageSize(v int32) *GetOcOperationRecruitmentRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcOperationRecruitmentRequest) SetSearchKey(v string) *GetOcOperationRecruitmentRequest {
	s.SearchKey = &v
	return s
}

type GetOcOperationRecruitmentResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcOperationRecruitmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                       `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                       `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                       `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcOperationRecruitmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationRecruitmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcOperationRecruitmentResponseBody) SetCode(v string) *GetOcOperationRecruitmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBody) SetData(v []*GetOcOperationRecruitmentResponseBodyData) *GetOcOperationRecruitmentResponseBody {
	s.Data = v
	return s
}

func (s *GetOcOperationRecruitmentResponseBody) SetMessage(v string) *GetOcOperationRecruitmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBody) SetPageIndex(v int32) *GetOcOperationRecruitmentResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBody) SetPageNum(v int32) *GetOcOperationRecruitmentResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBody) SetRequestId(v string) *GetOcOperationRecruitmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBody) SetSuccess(v bool) *GetOcOperationRecruitmentResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBody) SetTotalNum(v int32) *GetOcOperationRecruitmentResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcOperationRecruitmentResponseBodyData struct {
	BenefitList       *string `json:"BenefitList,omitempty" xml:"BenefitList,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Education         *string `json:"Education,omitempty" xml:"Education,omitempty"`
	EndDate           *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	EntName           *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	Experience        *string `json:"Experience,omitempty" xml:"Experience,omitempty"`
	PageUrl           *string `json:"PageUrl,omitempty" xml:"PageUrl,omitempty"`
	PublishDate       *string `json:"PublishDate,omitempty" xml:"PublishDate,omitempty"`
	RecruitingAddress *string `json:"RecruitingAddress,omitempty" xml:"RecruitingAddress,omitempty"`
	RecruitingName    *string `json:"RecruitingName,omitempty" xml:"RecruitingName,omitempty"`
	Salary            *string `json:"Salary,omitempty" xml:"Salary,omitempty"`
	StartDate         *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetOcOperationRecruitmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationRecruitmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetBenefitList(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.BenefitList = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetDescription(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetEducation(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.Education = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetEndDate(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetEntName(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetExperience(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.Experience = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetPageUrl(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.PageUrl = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetPublishDate(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.PublishDate = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetRecruitingAddress(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.RecruitingAddress = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetRecruitingName(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.RecruitingName = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetSalary(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.Salary = &v
	return s
}

func (s *GetOcOperationRecruitmentResponseBodyData) SetStartDate(v string) *GetOcOperationRecruitmentResponseBodyData {
	s.StartDate = &v
	return s
}

type GetOcOperationRecruitmentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcOperationRecruitmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcOperationRecruitmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcOperationRecruitmentResponse) GoString() string {
	return s.String()
}

func (s *GetOcOperationRecruitmentResponse) SetHeaders(v map[string]*string) *GetOcOperationRecruitmentResponse {
	s.Headers = v
	return s
}

func (s *GetOcOperationRecruitmentResponse) SetStatusCode(v int32) *GetOcOperationRecruitmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcOperationRecruitmentResponse) SetBody(v *GetOcOperationRecruitmentResponseBody) *GetOcOperationRecruitmentResponse {
	s.Body = v
	return s
}

type GetOcProductBandRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcProductBandRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcProductBandRequest) GoString() string {
	return s.String()
}

func (s *GetOcProductBandRequest) SetPageNo(v int32) *GetOcProductBandRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcProductBandRequest) SetPageSize(v int32) *GetOcProductBandRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcProductBandRequest) SetSearchKey(v string) *GetOcProductBandRequest {
	s.SearchKey = &v
	return s
}

type GetOcProductBandResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcProductBandResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                              `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                              `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                              `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcProductBandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcProductBandResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcProductBandResponseBody) SetCode(v string) *GetOcProductBandResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcProductBandResponseBody) SetData(v []*GetOcProductBandResponseBodyData) *GetOcProductBandResponseBody {
	s.Data = v
	return s
}

func (s *GetOcProductBandResponseBody) SetMessage(v string) *GetOcProductBandResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcProductBandResponseBody) SetPageIndex(v int32) *GetOcProductBandResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcProductBandResponseBody) SetPageNum(v int32) *GetOcProductBandResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcProductBandResponseBody) SetRequestId(v string) *GetOcProductBandResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcProductBandResponseBody) SetSuccess(v bool) *GetOcProductBandResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcProductBandResponseBody) SetTotalNum(v int32) *GetOcProductBandResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcProductBandResponseBodyData struct {
	BrandIntroduction   *string `json:"BrandIntroduction,omitempty" xml:"BrandIntroduction,omitempty"`
	Device              *string `json:"Device,omitempty" xml:"Device,omitempty"`
	EntName             *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	ProductIntroduction *string `json:"ProductIntroduction,omitempty" xml:"ProductIntroduction,omitempty"`
	ProductLogo         *string `json:"ProductLogo,omitempty" xml:"ProductLogo,omitempty"`
	ProductName         *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductTag          *string `json:"ProductTag,omitempty" xml:"ProductTag,omitempty"`
	ProductWebsite      *string `json:"ProductWebsite,omitempty" xml:"ProductWebsite,omitempty"`
}

func (s GetOcProductBandResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcProductBandResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcProductBandResponseBodyData) SetBrandIntroduction(v string) *GetOcProductBandResponseBodyData {
	s.BrandIntroduction = &v
	return s
}

func (s *GetOcProductBandResponseBodyData) SetDevice(v string) *GetOcProductBandResponseBodyData {
	s.Device = &v
	return s
}

func (s *GetOcProductBandResponseBodyData) SetEntName(v string) *GetOcProductBandResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcProductBandResponseBodyData) SetProductIntroduction(v string) *GetOcProductBandResponseBodyData {
	s.ProductIntroduction = &v
	return s
}

func (s *GetOcProductBandResponseBodyData) SetProductLogo(v string) *GetOcProductBandResponseBodyData {
	s.ProductLogo = &v
	return s
}

func (s *GetOcProductBandResponseBodyData) SetProductName(v string) *GetOcProductBandResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetOcProductBandResponseBodyData) SetProductTag(v string) *GetOcProductBandResponseBodyData {
	s.ProductTag = &v
	return s
}

func (s *GetOcProductBandResponseBodyData) SetProductWebsite(v string) *GetOcProductBandResponseBodyData {
	s.ProductWebsite = &v
	return s
}

type GetOcProductBandResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcProductBandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcProductBandResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcProductBandResponse) GoString() string {
	return s.String()
}

func (s *GetOcProductBandResponse) SetHeaders(v map[string]*string) *GetOcProductBandResponse {
	s.Headers = v
	return s
}

func (s *GetOcProductBandResponse) SetStatusCode(v int32) *GetOcProductBandResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcProductBandResponse) SetBody(v *GetOcProductBandResponseBody) *GetOcProductBandResponse {
	s.Body = v
	return s
}

type GetOcTaxAbnormalRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcTaxAbnormalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxAbnormalRequest) GoString() string {
	return s.String()
}

func (s *GetOcTaxAbnormalRequest) SetPageNo(v int32) *GetOcTaxAbnormalRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcTaxAbnormalRequest) SetPageSize(v int32) *GetOcTaxAbnormalRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcTaxAbnormalRequest) SetSearchKey(v string) *GetOcTaxAbnormalRequest {
	s.SearchKey = &v
	return s
}

type GetOcTaxAbnormalResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcTaxAbnormalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                              `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                              `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                              `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcTaxAbnormalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxAbnormalResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcTaxAbnormalResponseBody) SetCode(v string) *GetOcTaxAbnormalResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBody) SetData(v []*GetOcTaxAbnormalResponseBodyData) *GetOcTaxAbnormalResponseBody {
	s.Data = v
	return s
}

func (s *GetOcTaxAbnormalResponseBody) SetMessage(v string) *GetOcTaxAbnormalResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBody) SetPageIndex(v int32) *GetOcTaxAbnormalResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBody) SetPageNum(v int32) *GetOcTaxAbnormalResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBody) SetRequestId(v string) *GetOcTaxAbnormalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBody) SetSuccess(v bool) *GetOcTaxAbnormalResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBody) SetTotalNum(v int32) *GetOcTaxAbnormalResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcTaxAbnormalResponseBodyData struct {
	CardNum         *string `json:"CardNum,omitempty" xml:"CardNum,omitempty"`
	CardType        *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	EntName         *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	JudgeDate       *string `json:"JudgeDate,omitempty" xml:"JudgeDate,omitempty"`
	JudgeDepartment *string `json:"JudgeDepartment,omitempty" xml:"JudgeDepartment,omitempty"`
	JudgeReason     *string `json:"JudgeReason,omitempty" xml:"JudgeReason,omitempty"`
	LegalName       *string `json:"LegalName,omitempty" xml:"LegalName,omitempty"`
	OverdueAmount   *string `json:"OverdueAmount,omitempty" xml:"OverdueAmount,omitempty"`
	OverdueType     *string `json:"OverdueType,omitempty" xml:"OverdueType,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaxpayerNum     *string `json:"TaxpayerNum,omitempty" xml:"TaxpayerNum,omitempty"`
}

func (s GetOcTaxAbnormalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxAbnormalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcTaxAbnormalResponseBodyData) SetCardNum(v string) *GetOcTaxAbnormalResponseBodyData {
	s.CardNum = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetCardType(v string) *GetOcTaxAbnormalResponseBodyData {
	s.CardType = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetEntName(v string) *GetOcTaxAbnormalResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetJudgeDate(v string) *GetOcTaxAbnormalResponseBodyData {
	s.JudgeDate = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetJudgeDepartment(v string) *GetOcTaxAbnormalResponseBodyData {
	s.JudgeDepartment = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetJudgeReason(v string) *GetOcTaxAbnormalResponseBodyData {
	s.JudgeReason = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetLegalName(v string) *GetOcTaxAbnormalResponseBodyData {
	s.LegalName = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetOverdueAmount(v string) *GetOcTaxAbnormalResponseBodyData {
	s.OverdueAmount = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetOverdueType(v string) *GetOcTaxAbnormalResponseBodyData {
	s.OverdueType = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetStatus(v string) *GetOcTaxAbnormalResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetOcTaxAbnormalResponseBodyData) SetTaxpayerNum(v string) *GetOcTaxAbnormalResponseBodyData {
	s.TaxpayerNum = &v
	return s
}

type GetOcTaxAbnormalResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcTaxAbnormalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcTaxAbnormalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxAbnormalResponse) GoString() string {
	return s.String()
}

func (s *GetOcTaxAbnormalResponse) SetHeaders(v map[string]*string) *GetOcTaxAbnormalResponse {
	s.Headers = v
	return s
}

func (s *GetOcTaxAbnormalResponse) SetStatusCode(v int32) *GetOcTaxAbnormalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcTaxAbnormalResponse) SetBody(v *GetOcTaxAbnormalResponseBody) *GetOcTaxAbnormalResponse {
	s.Body = v
	return s
}

type GetOcTaxClassARequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcTaxClassARequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxClassARequest) GoString() string {
	return s.String()
}

func (s *GetOcTaxClassARequest) SetPageNo(v int32) *GetOcTaxClassARequest {
	s.PageNo = &v
	return s
}

func (s *GetOcTaxClassARequest) SetPageSize(v int32) *GetOcTaxClassARequest {
	s.PageSize = &v
	return s
}

func (s *GetOcTaxClassARequest) SetSearchKey(v string) *GetOcTaxClassARequest {
	s.SearchKey = &v
	return s
}

type GetOcTaxClassAResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcTaxClassAResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                            `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcTaxClassAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxClassAResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcTaxClassAResponseBody) SetCode(v string) *GetOcTaxClassAResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcTaxClassAResponseBody) SetData(v []*GetOcTaxClassAResponseBodyData) *GetOcTaxClassAResponseBody {
	s.Data = v
	return s
}

func (s *GetOcTaxClassAResponseBody) SetMessage(v string) *GetOcTaxClassAResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcTaxClassAResponseBody) SetPageIndex(v int32) *GetOcTaxClassAResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcTaxClassAResponseBody) SetPageNum(v int32) *GetOcTaxClassAResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcTaxClassAResponseBody) SetRequestId(v string) *GetOcTaxClassAResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcTaxClassAResponseBody) SetSuccess(v bool) *GetOcTaxClassAResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcTaxClassAResponseBody) SetTotalNum(v int32) *GetOcTaxClassAResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcTaxClassAResponseBodyData struct {
	EntName     *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	TaxLevel    *string `json:"TaxLevel,omitempty" xml:"TaxLevel,omitempty"`
	TaxpayerNum *string `json:"TaxpayerNum,omitempty" xml:"TaxpayerNum,omitempty"`
	Year        *string `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s GetOcTaxClassAResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxClassAResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcTaxClassAResponseBodyData) SetEntName(v string) *GetOcTaxClassAResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcTaxClassAResponseBodyData) SetTaxLevel(v string) *GetOcTaxClassAResponseBodyData {
	s.TaxLevel = &v
	return s
}

func (s *GetOcTaxClassAResponseBodyData) SetTaxpayerNum(v string) *GetOcTaxClassAResponseBodyData {
	s.TaxpayerNum = &v
	return s
}

func (s *GetOcTaxClassAResponseBodyData) SetYear(v string) *GetOcTaxClassAResponseBodyData {
	s.Year = &v
	return s
}

type GetOcTaxClassAResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcTaxClassAResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcTaxClassAResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxClassAResponse) GoString() string {
	return s.String()
}

func (s *GetOcTaxClassAResponse) SetHeaders(v map[string]*string) *GetOcTaxClassAResponse {
	s.Headers = v
	return s
}

func (s *GetOcTaxClassAResponse) SetStatusCode(v int32) *GetOcTaxClassAResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcTaxClassAResponse) SetBody(v *GetOcTaxClassAResponseBody) *GetOcTaxClassAResponse {
	s.Body = v
	return s
}

type GetOcTaxGeneralTaxpayerRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcTaxGeneralTaxpayerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxGeneralTaxpayerRequest) GoString() string {
	return s.String()
}

func (s *GetOcTaxGeneralTaxpayerRequest) SetPageNo(v int32) *GetOcTaxGeneralTaxpayerRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerRequest) SetPageSize(v int32) *GetOcTaxGeneralTaxpayerRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerRequest) SetSearchKey(v string) *GetOcTaxGeneralTaxpayerRequest {
	s.SearchKey = &v
	return s
}

type GetOcTaxGeneralTaxpayerResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcTaxGeneralTaxpayerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                     `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                     `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                     `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcTaxGeneralTaxpayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxGeneralTaxpayerResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcTaxGeneralTaxpayerResponseBody) SetCode(v string) *GetOcTaxGeneralTaxpayerResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBody) SetData(v []*GetOcTaxGeneralTaxpayerResponseBodyData) *GetOcTaxGeneralTaxpayerResponseBody {
	s.Data = v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBody) SetMessage(v string) *GetOcTaxGeneralTaxpayerResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBody) SetPageIndex(v int32) *GetOcTaxGeneralTaxpayerResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBody) SetPageNum(v int32) *GetOcTaxGeneralTaxpayerResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBody) SetRequestId(v string) *GetOcTaxGeneralTaxpayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBody) SetSuccess(v bool) *GetOcTaxGeneralTaxpayerResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBody) SetTotalNum(v int32) *GetOcTaxGeneralTaxpayerResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcTaxGeneralTaxpayerResponseBodyData struct {
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	EntName       *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	JudgeDate     *string `json:"JudgeDate,omitempty" xml:"JudgeDate,omitempty"`
	Qualification *string `json:"Qualification,omitempty" xml:"Qualification,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TaxpayerNum   *string `json:"TaxpayerNum,omitempty" xml:"TaxpayerNum,omitempty"`
}

func (s GetOcTaxGeneralTaxpayerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxGeneralTaxpayerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcTaxGeneralTaxpayerResponseBodyData) SetDepartment(v string) *GetOcTaxGeneralTaxpayerResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBodyData) SetEndDate(v string) *GetOcTaxGeneralTaxpayerResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBodyData) SetEntName(v string) *GetOcTaxGeneralTaxpayerResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBodyData) SetJudgeDate(v string) *GetOcTaxGeneralTaxpayerResponseBodyData {
	s.JudgeDate = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBodyData) SetQualification(v string) *GetOcTaxGeneralTaxpayerResponseBodyData {
	s.Qualification = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBodyData) SetStartDate(v string) *GetOcTaxGeneralTaxpayerResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponseBodyData) SetTaxpayerNum(v string) *GetOcTaxGeneralTaxpayerResponseBodyData {
	s.TaxpayerNum = &v
	return s
}

type GetOcTaxGeneralTaxpayerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcTaxGeneralTaxpayerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcTaxGeneralTaxpayerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxGeneralTaxpayerResponse) GoString() string {
	return s.String()
}

func (s *GetOcTaxGeneralTaxpayerResponse) SetHeaders(v map[string]*string) *GetOcTaxGeneralTaxpayerResponse {
	s.Headers = v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponse) SetStatusCode(v int32) *GetOcTaxGeneralTaxpayerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcTaxGeneralTaxpayerResponse) SetBody(v *GetOcTaxGeneralTaxpayerResponseBody) *GetOcTaxGeneralTaxpayerResponse {
	s.Body = v
	return s
}

type GetOcTaxIllegalRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcTaxIllegalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxIllegalRequest) GoString() string {
	return s.String()
}

func (s *GetOcTaxIllegalRequest) SetPageNo(v int32) *GetOcTaxIllegalRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcTaxIllegalRequest) SetPageSize(v int32) *GetOcTaxIllegalRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcTaxIllegalRequest) SetSearchKey(v string) *GetOcTaxIllegalRequest {
	s.SearchKey = &v
	return s
}

type GetOcTaxIllegalResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcTaxIllegalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                             `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                             `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcTaxIllegalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxIllegalResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcTaxIllegalResponseBody) SetCode(v string) *GetOcTaxIllegalResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcTaxIllegalResponseBody) SetData(v []*GetOcTaxIllegalResponseBodyData) *GetOcTaxIllegalResponseBody {
	s.Data = v
	return s
}

func (s *GetOcTaxIllegalResponseBody) SetMessage(v string) *GetOcTaxIllegalResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcTaxIllegalResponseBody) SetPageIndex(v int32) *GetOcTaxIllegalResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcTaxIllegalResponseBody) SetPageNum(v int32) *GetOcTaxIllegalResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcTaxIllegalResponseBody) SetRequestId(v string) *GetOcTaxIllegalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcTaxIllegalResponseBody) SetSuccess(v bool) *GetOcTaxIllegalResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcTaxIllegalResponseBody) SetTotalNum(v int32) *GetOcTaxIllegalResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcTaxIllegalResponseBodyData struct {
	AgencyCardNum     *string `json:"AgencyCardNum,omitempty" xml:"AgencyCardNum,omitempty"`
	AgencyCardType    *string `json:"AgencyCardType,omitempty" xml:"AgencyCardType,omitempty"`
	AgencyEnt         *string `json:"AgencyEnt,omitempty" xml:"AgencyEnt,omitempty"`
	AgencyName        *string `json:"AgencyName,omitempty" xml:"AgencyName,omitempty"`
	AgencySex         *string `json:"AgencySex,omitempty" xml:"AgencySex,omitempty"`
	CaseType          *string `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	Department        *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EntAddress        *string `json:"EntAddress,omitempty" xml:"EntAddress,omitempty"`
	EntName           *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	FinancialCardNum  *string `json:"FinancialCardNum,omitempty" xml:"FinancialCardNum,omitempty"`
	FinancialCardType *string `json:"FinancialCardType,omitempty" xml:"FinancialCardType,omitempty"`
	FinancialName     *string `json:"FinancialName,omitempty" xml:"FinancialName,omitempty"`
	FinancialSex      *string `json:"FinancialSex,omitempty" xml:"FinancialSex,omitempty"`
	IllegalTruth      *string `json:"IllegalTruth,omitempty" xml:"IllegalTruth,omitempty"`
	LawBasis          *string `json:"LawBasis,omitempty" xml:"LawBasis,omitempty"`
	LegalCardNum      *string `json:"LegalCardNum,omitempty" xml:"LegalCardNum,omitempty"`
	LegalCardType     *string `json:"LegalCardType,omitempty" xml:"LegalCardType,omitempty"`
	LegalName         *string `json:"LegalName,omitempty" xml:"LegalName,omitempty"`
	LegalSex          *string `json:"LegalSex,omitempty" xml:"LegalSex,omitempty"`
	OrgCode           *string `json:"OrgCode,omitempty" xml:"OrgCode,omitempty"`
	PublishDate       *string `json:"PublishDate,omitempty" xml:"PublishDate,omitempty"`
	TaxpayerNum       *string `json:"TaxpayerNum,omitempty" xml:"TaxpayerNum,omitempty"`
}

func (s GetOcTaxIllegalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxIllegalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcTaxIllegalResponseBodyData) SetAgencyCardNum(v string) *GetOcTaxIllegalResponseBodyData {
	s.AgencyCardNum = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetAgencyCardType(v string) *GetOcTaxIllegalResponseBodyData {
	s.AgencyCardType = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetAgencyEnt(v string) *GetOcTaxIllegalResponseBodyData {
	s.AgencyEnt = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetAgencyName(v string) *GetOcTaxIllegalResponseBodyData {
	s.AgencyName = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetAgencySex(v string) *GetOcTaxIllegalResponseBodyData {
	s.AgencySex = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetCaseType(v string) *GetOcTaxIllegalResponseBodyData {
	s.CaseType = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetDepartment(v string) *GetOcTaxIllegalResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetEntAddress(v string) *GetOcTaxIllegalResponseBodyData {
	s.EntAddress = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetEntName(v string) *GetOcTaxIllegalResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetFinancialCardNum(v string) *GetOcTaxIllegalResponseBodyData {
	s.FinancialCardNum = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetFinancialCardType(v string) *GetOcTaxIllegalResponseBodyData {
	s.FinancialCardType = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetFinancialName(v string) *GetOcTaxIllegalResponseBodyData {
	s.FinancialName = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetFinancialSex(v string) *GetOcTaxIllegalResponseBodyData {
	s.FinancialSex = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetIllegalTruth(v string) *GetOcTaxIllegalResponseBodyData {
	s.IllegalTruth = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetLawBasis(v string) *GetOcTaxIllegalResponseBodyData {
	s.LawBasis = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetLegalCardNum(v string) *GetOcTaxIllegalResponseBodyData {
	s.LegalCardNum = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetLegalCardType(v string) *GetOcTaxIllegalResponseBodyData {
	s.LegalCardType = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetLegalName(v string) *GetOcTaxIllegalResponseBodyData {
	s.LegalName = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetLegalSex(v string) *GetOcTaxIllegalResponseBodyData {
	s.LegalSex = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetOrgCode(v string) *GetOcTaxIllegalResponseBodyData {
	s.OrgCode = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetPublishDate(v string) *GetOcTaxIllegalResponseBodyData {
	s.PublishDate = &v
	return s
}

func (s *GetOcTaxIllegalResponseBodyData) SetTaxpayerNum(v string) *GetOcTaxIllegalResponseBodyData {
	s.TaxpayerNum = &v
	return s
}

type GetOcTaxIllegalResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcTaxIllegalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcTaxIllegalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxIllegalResponse) GoString() string {
	return s.String()
}

func (s *GetOcTaxIllegalResponse) SetHeaders(v map[string]*string) *GetOcTaxIllegalResponse {
	s.Headers = v
	return s
}

func (s *GetOcTaxIllegalResponse) SetStatusCode(v int32) *GetOcTaxIllegalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcTaxIllegalResponse) SetBody(v *GetOcTaxIllegalResponseBody) *GetOcTaxIllegalResponse {
	s.Body = v
	return s
}

type GetOcTaxOverdueRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcTaxOverdueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxOverdueRequest) GoString() string {
	return s.String()
}

func (s *GetOcTaxOverdueRequest) SetPageNo(v int32) *GetOcTaxOverdueRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcTaxOverdueRequest) SetPageSize(v int32) *GetOcTaxOverdueRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcTaxOverdueRequest) SetSearchKey(v string) *GetOcTaxOverdueRequest {
	s.SearchKey = &v
	return s
}

type GetOcTaxOverdueResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcTaxOverdueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                             `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                             `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcTaxOverdueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxOverdueResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcTaxOverdueResponseBody) SetCode(v string) *GetOcTaxOverdueResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcTaxOverdueResponseBody) SetData(v []*GetOcTaxOverdueResponseBodyData) *GetOcTaxOverdueResponseBody {
	s.Data = v
	return s
}

func (s *GetOcTaxOverdueResponseBody) SetMessage(v string) *GetOcTaxOverdueResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcTaxOverdueResponseBody) SetPageIndex(v int32) *GetOcTaxOverdueResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcTaxOverdueResponseBody) SetPageNum(v int32) *GetOcTaxOverdueResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcTaxOverdueResponseBody) SetRequestId(v string) *GetOcTaxOverdueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcTaxOverdueResponseBody) SetSuccess(v bool) *GetOcTaxOverdueResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcTaxOverdueResponseBody) SetTotalNum(v int32) *GetOcTaxOverdueResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcTaxOverdueResponseBodyData struct {
	CurrOverdueAmount *string `json:"CurrOverdueAmount,omitempty" xml:"CurrOverdueAmount,omitempty"`
	Department        *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EntAddress        *string `json:"EntAddress,omitempty" xml:"EntAddress,omitempty"`
	EntName           *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	LegalName         *string `json:"LegalName,omitempty" xml:"LegalName,omitempty"`
	OverdueAmount     *string `json:"OverdueAmount,omitempty" xml:"OverdueAmount,omitempty"`
	OverdueType       *string `json:"OverdueType,omitempty" xml:"OverdueType,omitempty"`
	PublishDate       *string `json:"PublishDate,omitempty" xml:"PublishDate,omitempty"`
	TaxpayerNum       *string `json:"TaxpayerNum,omitempty" xml:"TaxpayerNum,omitempty"`
	TaxpayerType      *string `json:"TaxpayerType,omitempty" xml:"TaxpayerType,omitempty"`
}

func (s GetOcTaxOverdueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxOverdueResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcTaxOverdueResponseBodyData) SetCurrOverdueAmount(v string) *GetOcTaxOverdueResponseBodyData {
	s.CurrOverdueAmount = &v
	return s
}

func (s *GetOcTaxOverdueResponseBodyData) SetDepartment(v string) *GetOcTaxOverdueResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcTaxOverdueResponseBodyData) SetEntAddress(v string) *GetOcTaxOverdueResponseBodyData {
	s.EntAddress = &v
	return s
}

func (s *GetOcTaxOverdueResponseBodyData) SetEntName(v string) *GetOcTaxOverdueResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcTaxOverdueResponseBodyData) SetLegalName(v string) *GetOcTaxOverdueResponseBodyData {
	s.LegalName = &v
	return s
}

func (s *GetOcTaxOverdueResponseBodyData) SetOverdueAmount(v string) *GetOcTaxOverdueResponseBodyData {
	s.OverdueAmount = &v
	return s
}

func (s *GetOcTaxOverdueResponseBodyData) SetOverdueType(v string) *GetOcTaxOverdueResponseBodyData {
	s.OverdueType = &v
	return s
}

func (s *GetOcTaxOverdueResponseBodyData) SetPublishDate(v string) *GetOcTaxOverdueResponseBodyData {
	s.PublishDate = &v
	return s
}

func (s *GetOcTaxOverdueResponseBodyData) SetTaxpayerNum(v string) *GetOcTaxOverdueResponseBodyData {
	s.TaxpayerNum = &v
	return s
}

func (s *GetOcTaxOverdueResponseBodyData) SetTaxpayerType(v string) *GetOcTaxOverdueResponseBodyData {
	s.TaxpayerType = &v
	return s
}

type GetOcTaxOverdueResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcTaxOverdueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcTaxOverdueResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxOverdueResponse) GoString() string {
	return s.String()
}

func (s *GetOcTaxOverdueResponse) SetHeaders(v map[string]*string) *GetOcTaxOverdueResponse {
	s.Headers = v
	return s
}

func (s *GetOcTaxOverdueResponse) SetStatusCode(v int32) *GetOcTaxOverdueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcTaxOverdueResponse) SetBody(v *GetOcTaxOverdueResponseBody) *GetOcTaxOverdueResponse {
	s.Body = v
	return s
}

type GetOcTaxPunishmentRequest struct {
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s GetOcTaxPunishmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxPunishmentRequest) GoString() string {
	return s.String()
}

func (s *GetOcTaxPunishmentRequest) SetPageNo(v int32) *GetOcTaxPunishmentRequest {
	s.PageNo = &v
	return s
}

func (s *GetOcTaxPunishmentRequest) SetPageSize(v int32) *GetOcTaxPunishmentRequest {
	s.PageSize = &v
	return s
}

func (s *GetOcTaxPunishmentRequest) SetSearchKey(v string) *GetOcTaxPunishmentRequest {
	s.SearchKey = &v
	return s
}

type GetOcTaxPunishmentResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetOcTaxPunishmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                                `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOcTaxPunishmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxPunishmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOcTaxPunishmentResponseBody) SetCode(v string) *GetOcTaxPunishmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBody) SetData(v []*GetOcTaxPunishmentResponseBodyData) *GetOcTaxPunishmentResponseBody {
	s.Data = v
	return s
}

func (s *GetOcTaxPunishmentResponseBody) SetMessage(v string) *GetOcTaxPunishmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBody) SetPageIndex(v int32) *GetOcTaxPunishmentResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBody) SetPageNum(v int32) *GetOcTaxPunishmentResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBody) SetRequestId(v string) *GetOcTaxPunishmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBody) SetSuccess(v bool) *GetOcTaxPunishmentResponseBody {
	s.Success = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBody) SetTotalNum(v int32) *GetOcTaxPunishmentResponseBody {
	s.TotalNum = &v
	return s
}

type GetOcTaxPunishmentResponseBodyData struct {
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	EntName     *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	EventName   *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventType   *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	LegalName   *string `json:"LegalName,omitempty" xml:"LegalName,omitempty"`
	PunishDate  *string `json:"PunishDate,omitempty" xml:"PunishDate,omitempty"`
	TaxpayerNum *string `json:"TaxpayerNum,omitempty" xml:"TaxpayerNum,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetOcTaxPunishmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxPunishmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOcTaxPunishmentResponseBodyData) SetDepartment(v string) *GetOcTaxPunishmentResponseBodyData {
	s.Department = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBodyData) SetEntName(v string) *GetOcTaxPunishmentResponseBodyData {
	s.EntName = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBodyData) SetEventName(v string) *GetOcTaxPunishmentResponseBodyData {
	s.EventName = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBodyData) SetEventType(v string) *GetOcTaxPunishmentResponseBodyData {
	s.EventType = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBodyData) SetLegalName(v string) *GetOcTaxPunishmentResponseBodyData {
	s.LegalName = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBodyData) SetPunishDate(v string) *GetOcTaxPunishmentResponseBodyData {
	s.PunishDate = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBodyData) SetTaxpayerNum(v string) *GetOcTaxPunishmentResponseBodyData {
	s.TaxpayerNum = &v
	return s
}

func (s *GetOcTaxPunishmentResponseBodyData) SetTitle(v string) *GetOcTaxPunishmentResponseBodyData {
	s.Title = &v
	return s
}

type GetOcTaxPunishmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOcTaxPunishmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOcTaxPunishmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOcTaxPunishmentResponse) GoString() string {
	return s.String()
}

func (s *GetOcTaxPunishmentResponse) SetHeaders(v map[string]*string) *GetOcTaxPunishmentResponse {
	s.Headers = v
	return s
}

func (s *GetOcTaxPunishmentResponse) SetStatusCode(v int32) *GetOcTaxPunishmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOcTaxPunishmentResponse) SetBody(v *GetOcTaxPunishmentResponseBody) *GetOcTaxPunishmentResponse {
	s.Body = v
	return s
}

type GetQccCertificationDetailByIdRequest struct {
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
}

func (s GetQccCertificationDetailByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQccCertificationDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetQccCertificationDetailByIdRequest) SetCertId(v string) *GetQccCertificationDetailByIdRequest {
	s.CertId = &v
	return s
}

type GetQccCertificationDetailByIdResponseBody struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetQccCertificationDetailByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQccCertificationDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetQccCertificationDetailByIdResponseBody) SetCode(v string) *GetQccCertificationDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetQccCertificationDetailByIdResponseBody) SetData(v map[string]interface{}) *GetQccCertificationDetailByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetQccCertificationDetailByIdResponseBody) SetMessage(v string) *GetQccCertificationDetailByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetQccCertificationDetailByIdResponseBody) SetPageIndex(v int32) *GetQccCertificationDetailByIdResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetQccCertificationDetailByIdResponseBody) SetPageNum(v int32) *GetQccCertificationDetailByIdResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetQccCertificationDetailByIdResponseBody) SetRequestId(v string) *GetQccCertificationDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQccCertificationDetailByIdResponseBody) SetSuccess(v bool) *GetQccCertificationDetailByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetQccCertificationDetailByIdResponseBody) SetTotalNum(v int32) *GetQccCertificationDetailByIdResponseBody {
	s.TotalNum = &v
	return s
}

type GetQccCertificationDetailByIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQccCertificationDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQccCertificationDetailByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQccCertificationDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetQccCertificationDetailByIdResponse) SetHeaders(v map[string]*string) *GetQccCertificationDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetQccCertificationDetailByIdResponse) SetStatusCode(v int32) *GetQccCertificationDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQccCertificationDetailByIdResponse) SetBody(v *GetQccCertificationDetailByIdResponseBody) *GetQccCertificationDetailByIdResponse {
	s.Body = v
	return s
}

type GetQccSearchCertificationRequest struct {
	CertCategory *string `json:"CertCategory,omitempty" xml:"CertCategory,omitempty"`
	EntName      *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	PageNo       *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetQccSearchCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQccSearchCertificationRequest) GoString() string {
	return s.String()
}

func (s *GetQccSearchCertificationRequest) SetCertCategory(v string) *GetQccSearchCertificationRequest {
	s.CertCategory = &v
	return s
}

func (s *GetQccSearchCertificationRequest) SetEntName(v string) *GetQccSearchCertificationRequest {
	s.EntName = &v
	return s
}

func (s *GetQccSearchCertificationRequest) SetPageNo(v int32) *GetQccSearchCertificationRequest {
	s.PageNo = &v
	return s
}

func (s *GetQccSearchCertificationRequest) SetPageSize(v int32) *GetQccSearchCertificationRequest {
	s.PageSize = &v
	return s
}

type GetQccSearchCertificationResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageIndex *int32                   `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageNum   *int32                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalNum  *int32                   `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetQccSearchCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQccSearchCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *GetQccSearchCertificationResponseBody) SetCode(v string) *GetQccSearchCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *GetQccSearchCertificationResponseBody) SetData(v []map[string]interface{}) *GetQccSearchCertificationResponseBody {
	s.Data = v
	return s
}

func (s *GetQccSearchCertificationResponseBody) SetMessage(v string) *GetQccSearchCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *GetQccSearchCertificationResponseBody) SetPageIndex(v int32) *GetQccSearchCertificationResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetQccSearchCertificationResponseBody) SetPageNum(v int32) *GetQccSearchCertificationResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetQccSearchCertificationResponseBody) SetRequestId(v string) *GetQccSearchCertificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQccSearchCertificationResponseBody) SetSuccess(v bool) *GetQccSearchCertificationResponseBody {
	s.Success = &v
	return s
}

func (s *GetQccSearchCertificationResponseBody) SetTotalNum(v int32) *GetQccSearchCertificationResponseBody {
	s.TotalNum = &v
	return s
}

type GetQccSearchCertificationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQccSearchCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQccSearchCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQccSearchCertificationResponse) GoString() string {
	return s.String()
}

func (s *GetQccSearchCertificationResponse) SetHeaders(v map[string]*string) *GetQccSearchCertificationResponse {
	s.Headers = v
	return s
}

func (s *GetQccSearchCertificationResponse) SetStatusCode(v int32) *GetQccSearchCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQccSearchCertificationResponse) SetBody(v *GetQccSearchCertificationResponseBody) *GetQccSearchCertificationResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dt-oc-info"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcCompetitorsWithOptions(request *GetOcCompetitorsRequest, runtime *util.RuntimeOptions) (_result *GetOcCompetitorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcCompetitors"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcCompetitorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcCompetitors(request *GetOcCompetitorsRequest) (_result *GetOcCompetitorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcCompetitorsResponse{}
	_body, _err := client.GetOcCompetitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcCoreTeamsWithOptions(request *GetOcCoreTeamsRequest, runtime *util.RuntimeOptions) (_result *GetOcCoreTeamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcCoreTeams"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcCoreTeamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcCoreTeams(request *GetOcCoreTeamsRequest) (_result *GetOcCoreTeamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcCoreTeamsResponse{}
	_body, _err := client.GetOcCoreTeamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcFinancingWithOptions(request *GetOcFinancingRequest, runtime *util.RuntimeOptions) (_result *GetOcFinancingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcFinancing"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcFinancingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcFinancing(request *GetOcFinancingRequest) (_result *GetOcFinancingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcFinancingResponse{}
	_body, _err := client.GetOcFinancingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcFuzzSearchWithOptions(request *GetOcFuzzSearchRequest, runtime *util.RuntimeOptions) (_result *GetOcFuzzSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcFuzzSearch"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcFuzzSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcFuzzSearch(request *GetOcFuzzSearchRequest) (_result *GetOcFuzzSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcFuzzSearchResponse{}
	_body, _err := client.GetOcFuzzSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcAbnormalOperationWithOptions(request *GetOcIcAbnormalOperationRequest, runtime *util.RuntimeOptions) (_result *GetOcIcAbnormalOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcAbnormalOperation"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcAbnormalOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcAbnormalOperation(request *GetOcIcAbnormalOperationRequest) (_result *GetOcIcAbnormalOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcAbnormalOperationResponse{}
	_body, _err := client.GetOcIcAbnormalOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcAdminLicenseWithOptions(request *GetOcIcAdminLicenseRequest, runtime *util.RuntimeOptions) (_result *GetOcIcAdminLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcAdminLicense"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcAdminLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcAdminLicense(request *GetOcIcAdminLicenseRequest) (_result *GetOcIcAdminLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcAdminLicenseResponse{}
	_body, _err := client.GetOcIcAdminLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcBasicWithOptions(request *GetOcIcBasicRequest, runtime *util.RuntimeOptions) (_result *GetOcIcBasicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcBasic"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcBasicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcBasic(request *GetOcIcBasicRequest) (_result *GetOcIcBasicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcBasicResponse{}
	_body, _err := client.GetOcIcBasicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcBranchWithOptions(request *GetOcIcBranchRequest, runtime *util.RuntimeOptions) (_result *GetOcIcBranchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcBranch"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcBranch(request *GetOcIcBranchRequest) (_result *GetOcIcBranchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcBranchResponse{}
	_body, _err := client.GetOcIcBranchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcChangeRecordWithOptions(request *GetOcIcChangeRecordRequest, runtime *util.RuntimeOptions) (_result *GetOcIcChangeRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcChangeRecord"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcChangeRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcChangeRecord(request *GetOcIcChangeRecordRequest) (_result *GetOcIcChangeRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcChangeRecordResponse{}
	_body, _err := client.GetOcIcChangeRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcCheckupWithOptions(request *GetOcIcCheckupRequest, runtime *util.RuntimeOptions) (_result *GetOcIcCheckupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcCheckup"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcCheckupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcCheckup(request *GetOcIcCheckupRequest) (_result *GetOcIcCheckupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcCheckupResponse{}
	_body, _err := client.GetOcIcCheckupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcClearAccountWithOptions(request *GetOcIcClearAccountRequest, runtime *util.RuntimeOptions) (_result *GetOcIcClearAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcClearAccount"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcClearAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcClearAccount(request *GetOcIcClearAccountRequest) (_result *GetOcIcClearAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcClearAccountResponse{}
	_body, _err := client.GetOcIcClearAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcDoubleCheckupWithOptions(request *GetOcIcDoubleCheckupRequest, runtime *util.RuntimeOptions) (_result *GetOcIcDoubleCheckupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcDoubleCheckup"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcDoubleCheckupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcDoubleCheckup(request *GetOcIcDoubleCheckupRequest) (_result *GetOcIcDoubleCheckupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcDoubleCheckupResponse{}
	_body, _err := client.GetOcIcDoubleCheckupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcEmployeeWithOptions(request *GetOcIcEmployeeRequest, runtime *util.RuntimeOptions) (_result *GetOcIcEmployeeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcEmployee"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcEmployeeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcEmployee(request *GetOcIcEmployeeRequest) (_result *GetOcIcEmployeeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcEmployeeResponse{}
	_body, _err := client.GetOcIcEmployeeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcEquityFrozenWithOptions(request *GetOcIcEquityFrozenRequest, runtime *util.RuntimeOptions) (_result *GetOcIcEquityFrozenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcEquityFrozen"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcEquityFrozenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcEquityFrozen(request *GetOcIcEquityFrozenRequest) (_result *GetOcIcEquityFrozenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcEquityFrozenResponse{}
	_body, _err := client.GetOcIcEquityFrozenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcEquityPledgeWithOptions(request *GetOcIcEquityPledgeRequest, runtime *util.RuntimeOptions) (_result *GetOcIcEquityPledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcEquityPledge"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcEquityPledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcEquityPledge(request *GetOcIcEquityPledgeRequest) (_result *GetOcIcEquityPledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcEquityPledgeResponse{}
	_body, _err := client.GetOcIcEquityPledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcInvestmentWithOptions(request *GetOcIcInvestmentRequest, runtime *util.RuntimeOptions) (_result *GetOcIcInvestmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcInvestment"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcInvestmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcInvestment(request *GetOcIcInvestmentRequest) (_result *GetOcIcInvestmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcInvestmentResponse{}
	_body, _err := client.GetOcIcInvestmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcKnowledgePropertyPledgeWithOptions(request *GetOcIcKnowledgePropertyPledgeRequest, runtime *util.RuntimeOptions) (_result *GetOcIcKnowledgePropertyPledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcKnowledgePropertyPledge"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcKnowledgePropertyPledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcKnowledgePropertyPledge(request *GetOcIcKnowledgePropertyPledgeRequest) (_result *GetOcIcKnowledgePropertyPledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcKnowledgePropertyPledgeResponse{}
	_body, _err := client.GetOcIcKnowledgePropertyPledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcMortgageWithOptions(request *GetOcIcMortgageRequest, runtime *util.RuntimeOptions) (_result *GetOcIcMortgageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcMortgage"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcMortgageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcMortgage(request *GetOcIcMortgageRequest) (_result *GetOcIcMortgageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcMortgageResponse{}
	_body, _err := client.GetOcIcMortgageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcSeriousOffenseWithOptions(request *GetOcIcSeriousOffenseRequest, runtime *util.RuntimeOptions) (_result *GetOcIcSeriousOffenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcSeriousOffense"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcSeriousOffenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcSeriousOffense(request *GetOcIcSeriousOffenseRequest) (_result *GetOcIcSeriousOffenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcSeriousOffenseResponse{}
	_body, _err := client.GetOcIcSeriousOffenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcShareholderWithOptions(request *GetOcIcShareholderRequest, runtime *util.RuntimeOptions) (_result *GetOcIcShareholderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcShareholder"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcShareholderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcShareholder(request *GetOcIcShareholderRequest) (_result *GetOcIcShareholderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcShareholderResponse{}
	_body, _err := client.GetOcIcShareholderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIcSimpleCancelWithOptions(request *GetOcIcSimpleCancelRequest, runtime *util.RuntimeOptions) (_result *GetOcIcSimpleCancelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIcSimpleCancel"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIcSimpleCancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIcSimpleCancel(request *GetOcIcSimpleCancelRequest) (_result *GetOcIcSimpleCancelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIcSimpleCancelResponse{}
	_body, _err := client.GetOcIcSimpleCancelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIpCertificateWithOptions(request *GetOcIpCertificateRequest, runtime *util.RuntimeOptions) (_result *GetOcIpCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIpCertificate"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIpCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIpCertificate(request *GetOcIpCertificateRequest) (_result *GetOcIpCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIpCertificateResponse{}
	_body, _err := client.GetOcIpCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIpDomainWithOptions(request *GetOcIpDomainRequest, runtime *util.RuntimeOptions) (_result *GetOcIpDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIpDomain"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIpDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIpDomain(request *GetOcIpDomainRequest) (_result *GetOcIpDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIpDomainResponse{}
	_body, _err := client.GetOcIpDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIpPatentWithOptions(request *GetOcIpPatentRequest, runtime *util.RuntimeOptions) (_result *GetOcIpPatentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIpPatent"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIpPatentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIpPatent(request *GetOcIpPatentRequest) (_result *GetOcIpPatentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIpPatentResponse{}
	_body, _err := client.GetOcIpPatentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIpSoftwareCopyrightWithOptions(request *GetOcIpSoftwareCopyrightRequest, runtime *util.RuntimeOptions) (_result *GetOcIpSoftwareCopyrightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIpSoftwareCopyright"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIpSoftwareCopyrightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIpSoftwareCopyright(request *GetOcIpSoftwareCopyrightRequest) (_result *GetOcIpSoftwareCopyrightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIpSoftwareCopyrightResponse{}
	_body, _err := client.GetOcIpSoftwareCopyrightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIpTrademarkWithOptions(request *GetOcIpTrademarkRequest, runtime *util.RuntimeOptions) (_result *GetOcIpTrademarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIpTrademark"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIpTrademarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIpTrademark(request *GetOcIpTrademarkRequest) (_result *GetOcIpTrademarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIpTrademarkResponse{}
	_body, _err := client.GetOcIpTrademarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcIpWorksCopyrightWithOptions(request *GetOcIpWorksCopyrightRequest, runtime *util.RuntimeOptions) (_result *GetOcIpWorksCopyrightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcIpWorksCopyright"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcIpWorksCopyrightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcIpWorksCopyright(request *GetOcIpWorksCopyrightRequest) (_result *GetOcIpWorksCopyrightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcIpWorksCopyrightResponse{}
	_body, _err := client.GetOcIpWorksCopyrightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcJusticeAuctionWithOptions(request *GetOcJusticeAuctionRequest, runtime *util.RuntimeOptions) (_result *GetOcJusticeAuctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcJusticeAuction"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcJusticeAuctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcJusticeAuction(request *GetOcJusticeAuctionRequest) (_result *GetOcJusticeAuctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcJusticeAuctionResponse{}
	_body, _err := client.GetOcJusticeAuctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcJusticeCaseFilingWithOptions(request *GetOcJusticeCaseFilingRequest, runtime *util.RuntimeOptions) (_result *GetOcJusticeCaseFilingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcJusticeCaseFiling"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcJusticeCaseFilingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcJusticeCaseFiling(request *GetOcJusticeCaseFilingRequest) (_result *GetOcJusticeCaseFilingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcJusticeCaseFilingResponse{}
	_body, _err := client.GetOcJusticeCaseFilingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcJusticeCourtAnnouncementWithOptions(request *GetOcJusticeCourtAnnouncementRequest, runtime *util.RuntimeOptions) (_result *GetOcJusticeCourtAnnouncementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcJusticeCourtAnnouncement"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcJusticeCourtAnnouncementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcJusticeCourtAnnouncement(request *GetOcJusticeCourtAnnouncementRequest) (_result *GetOcJusticeCourtAnnouncementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcJusticeCourtAnnouncementResponse{}
	_body, _err := client.GetOcJusticeCourtAnnouncementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcJusticeCourtNoticeWithOptions(request *GetOcJusticeCourtNoticeRequest, runtime *util.RuntimeOptions) (_result *GetOcJusticeCourtNoticeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcJusticeCourtNotice"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcJusticeCourtNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcJusticeCourtNotice(request *GetOcJusticeCourtNoticeRequest) (_result *GetOcJusticeCourtNoticeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcJusticeCourtNoticeResponse{}
	_body, _err := client.GetOcJusticeCourtNoticeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcJusticeDishonestyWithOptions(request *GetOcJusticeDishonestyRequest, runtime *util.RuntimeOptions) (_result *GetOcJusticeDishonestyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcJusticeDishonesty"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcJusticeDishonestyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcJusticeDishonesty(request *GetOcJusticeDishonestyRequest) (_result *GetOcJusticeDishonestyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcJusticeDishonestyResponse{}
	_body, _err := client.GetOcJusticeDishonestyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcJusticeExecutedWithOptions(request *GetOcJusticeExecutedRequest, runtime *util.RuntimeOptions) (_result *GetOcJusticeExecutedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcJusticeExecuted"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcJusticeExecutedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcJusticeExecuted(request *GetOcJusticeExecutedRequest) (_result *GetOcJusticeExecutedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcJusticeExecutedResponse{}
	_body, _err := client.GetOcJusticeExecutedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcJusticeJudgementDocWithOptions(request *GetOcJusticeJudgementDocRequest, runtime *util.RuntimeOptions) (_result *GetOcJusticeJudgementDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcJusticeJudgementDoc"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcJusticeJudgementDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcJusticeJudgementDoc(request *GetOcJusticeJudgementDocRequest) (_result *GetOcJusticeJudgementDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcJusticeJudgementDocResponse{}
	_body, _err := client.GetOcJusticeJudgementDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcJusticeLimitHighConsumeWithOptions(request *GetOcJusticeLimitHighConsumeRequest, runtime *util.RuntimeOptions) (_result *GetOcJusticeLimitHighConsumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcJusticeLimitHighConsume"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcJusticeLimitHighConsumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcJusticeLimitHighConsume(request *GetOcJusticeLimitHighConsumeRequest) (_result *GetOcJusticeLimitHighConsumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcJusticeLimitHighConsumeResponse{}
	_body, _err := client.GetOcJusticeLimitHighConsumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcJusticeTerminalCaseWithOptions(request *GetOcJusticeTerminalCaseRequest, runtime *util.RuntimeOptions) (_result *GetOcJusticeTerminalCaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcJusticeTerminalCase"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcJusticeTerminalCaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcJusticeTerminalCase(request *GetOcJusticeTerminalCaseRequest) (_result *GetOcJusticeTerminalCaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcJusticeTerminalCaseResponse{}
	_body, _err := client.GetOcJusticeTerminalCaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcListedCompanyWithOptions(request *GetOcListedCompanyRequest, runtime *util.RuntimeOptions) (_result *GetOcListedCompanyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcListedCompany"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcListedCompanyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcListedCompany(request *GetOcListedCompanyRequest) (_result *GetOcListedCompanyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcListedCompanyResponse{}
	_body, _err := client.GetOcListedCompanyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcNegativeAdminPunishmentWithOptions(request *GetOcNegativeAdminPunishmentRequest, runtime *util.RuntimeOptions) (_result *GetOcNegativeAdminPunishmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcNegativeAdminPunishment"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcNegativeAdminPunishmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcNegativeAdminPunishment(request *GetOcNegativeAdminPunishmentRequest) (_result *GetOcNegativeAdminPunishmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcNegativeAdminPunishmentResponse{}
	_body, _err := client.GetOcNegativeAdminPunishmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcNegativeCustomsPunishmentWithOptions(request *GetOcNegativeCustomsPunishmentRequest, runtime *util.RuntimeOptions) (_result *GetOcNegativeCustomsPunishmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcNegativeCustomsPunishment"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcNegativeCustomsPunishmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcNegativeCustomsPunishment(request *GetOcNegativeCustomsPunishmentRequest) (_result *GetOcNegativeCustomsPunishmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcNegativeCustomsPunishmentResponse{}
	_body, _err := client.GetOcNegativeCustomsPunishmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcNegativeEnvironmentPunishmentWithOptions(request *GetOcNegativeEnvironmentPunishmentRequest, runtime *util.RuntimeOptions) (_result *GetOcNegativeEnvironmentPunishmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcNegativeEnvironmentPunishment"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcNegativeEnvironmentPunishmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcNegativeEnvironmentPunishment(request *GetOcNegativeEnvironmentPunishmentRequest) (_result *GetOcNegativeEnvironmentPunishmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcNegativeEnvironmentPunishmentResponse{}
	_body, _err := client.GetOcNegativeEnvironmentPunishmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcNegativeFoodDrugPunishmentWithOptions(request *GetOcNegativeFoodDrugPunishmentRequest, runtime *util.RuntimeOptions) (_result *GetOcNegativeFoodDrugPunishmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcNegativeFoodDrugPunishment"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcNegativeFoodDrugPunishmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcNegativeFoodDrugPunishment(request *GetOcNegativeFoodDrugPunishmentRequest) (_result *GetOcNegativeFoodDrugPunishmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcNegativeFoodDrugPunishmentResponse{}
	_body, _err := client.GetOcNegativeFoodDrugPunishmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcNegativeQualityPunishmentWithOptions(request *GetOcNegativeQualityPunishmentRequest, runtime *util.RuntimeOptions) (_result *GetOcNegativeQualityPunishmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcNegativeQualityPunishment"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcNegativeQualityPunishmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcNegativeQualityPunishment(request *GetOcNegativeQualityPunishmentRequest) (_result *GetOcNegativeQualityPunishmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcNegativeQualityPunishmentResponse{}
	_body, _err := client.GetOcNegativeQualityPunishmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcOperationBiddingWithOptions(request *GetOcOperationBiddingRequest, runtime *util.RuntimeOptions) (_result *GetOcOperationBiddingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcOperationBidding"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcOperationBiddingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcOperationBidding(request *GetOcOperationBiddingRequest) (_result *GetOcOperationBiddingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcOperationBiddingResponse{}
	_body, _err := client.GetOcOperationBiddingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcOperationCustomsWithOptions(request *GetOcOperationCustomsRequest, runtime *util.RuntimeOptions) (_result *GetOcOperationCustomsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcOperationCustoms"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcOperationCustomsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcOperationCustoms(request *GetOcOperationCustomsRequest) (_result *GetOcOperationCustomsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcOperationCustomsResponse{}
	_body, _err := client.GetOcOperationCustomsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcOperationPurchaseLandWithOptions(request *GetOcOperationPurchaseLandRequest, runtime *util.RuntimeOptions) (_result *GetOcOperationPurchaseLandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcOperationPurchaseLand"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcOperationPurchaseLandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcOperationPurchaseLand(request *GetOcOperationPurchaseLandRequest) (_result *GetOcOperationPurchaseLandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcOperationPurchaseLandResponse{}
	_body, _err := client.GetOcOperationPurchaseLandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcOperationRecruitmentWithOptions(request *GetOcOperationRecruitmentRequest, runtime *util.RuntimeOptions) (_result *GetOcOperationRecruitmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcOperationRecruitment"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcOperationRecruitmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcOperationRecruitment(request *GetOcOperationRecruitmentRequest) (_result *GetOcOperationRecruitmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcOperationRecruitmentResponse{}
	_body, _err := client.GetOcOperationRecruitmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcProductBandWithOptions(request *GetOcProductBandRequest, runtime *util.RuntimeOptions) (_result *GetOcProductBandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcProductBand"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcProductBandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcProductBand(request *GetOcProductBandRequest) (_result *GetOcProductBandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcProductBandResponse{}
	_body, _err := client.GetOcProductBandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcTaxAbnormalWithOptions(request *GetOcTaxAbnormalRequest, runtime *util.RuntimeOptions) (_result *GetOcTaxAbnormalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcTaxAbnormal"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcTaxAbnormalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcTaxAbnormal(request *GetOcTaxAbnormalRequest) (_result *GetOcTaxAbnormalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcTaxAbnormalResponse{}
	_body, _err := client.GetOcTaxAbnormalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcTaxClassAWithOptions(request *GetOcTaxClassARequest, runtime *util.RuntimeOptions) (_result *GetOcTaxClassAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcTaxClassA"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcTaxClassAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcTaxClassA(request *GetOcTaxClassARequest) (_result *GetOcTaxClassAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcTaxClassAResponse{}
	_body, _err := client.GetOcTaxClassAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcTaxGeneralTaxpayerWithOptions(request *GetOcTaxGeneralTaxpayerRequest, runtime *util.RuntimeOptions) (_result *GetOcTaxGeneralTaxpayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcTaxGeneralTaxpayer"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcTaxGeneralTaxpayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcTaxGeneralTaxpayer(request *GetOcTaxGeneralTaxpayerRequest) (_result *GetOcTaxGeneralTaxpayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcTaxGeneralTaxpayerResponse{}
	_body, _err := client.GetOcTaxGeneralTaxpayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcTaxIllegalWithOptions(request *GetOcTaxIllegalRequest, runtime *util.RuntimeOptions) (_result *GetOcTaxIllegalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcTaxIllegal"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcTaxIllegalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcTaxIllegal(request *GetOcTaxIllegalRequest) (_result *GetOcTaxIllegalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcTaxIllegalResponse{}
	_body, _err := client.GetOcTaxIllegalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcTaxOverdueWithOptions(request *GetOcTaxOverdueRequest, runtime *util.RuntimeOptions) (_result *GetOcTaxOverdueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcTaxOverdue"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcTaxOverdueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcTaxOverdue(request *GetOcTaxOverdueRequest) (_result *GetOcTaxOverdueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcTaxOverdueResponse{}
	_body, _err := client.GetOcTaxOverdueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOcTaxPunishmentWithOptions(request *GetOcTaxPunishmentRequest, runtime *util.RuntimeOptions) (_result *GetOcTaxPunishmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOcTaxPunishment"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOcTaxPunishmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOcTaxPunishment(request *GetOcTaxPunishmentRequest) (_result *GetOcTaxPunishmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOcTaxPunishmentResponse{}
	_body, _err := client.GetOcTaxPunishmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQccCertificationDetailByIdWithOptions(request *GetQccCertificationDetailByIdRequest, runtime *util.RuntimeOptions) (_result *GetQccCertificationDetailByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertId)) {
		body["CertId"] = request.CertId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQccCertificationDetailById"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQccCertificationDetailByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQccCertificationDetailById(request *GetQccCertificationDetailByIdRequest) (_result *GetQccCertificationDetailByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQccCertificationDetailByIdResponse{}
	_body, _err := client.GetQccCertificationDetailByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQccSearchCertificationWithOptions(request *GetQccSearchCertificationRequest, runtime *util.RuntimeOptions) (_result *GetQccSearchCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertCategory)) {
		body["CertCategory"] = request.CertCategory
	}

	if !tea.BoolValue(util.IsUnset(request.EntName)) {
		body["EntName"] = request.EntName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQccSearchCertification"),
		Version:     tea.String("2022-08-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQccSearchCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQccSearchCertification(request *GetQccSearchCertificationRequest) (_result *GetQccSearchCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQccSearchCertificationResponse{}
	_body, _err := client.GetQccSearchCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
