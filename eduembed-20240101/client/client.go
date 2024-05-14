// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateLabReservationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 16600
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-29 18:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 875
	LabId *int64 `json:"LabId,omitempty" xml:"LabId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MemberCount *int64 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-29 16:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateLabReservationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLabReservationRequest) GoString() string {
	return s.String()
}

func (s *CreateLabReservationRequest) SetAccountId(v int64) *CreateLabReservationRequest {
	s.AccountId = &v
	return s
}

func (s *CreateLabReservationRequest) SetEndTime(v string) *CreateLabReservationRequest {
	s.EndTime = &v
	return s
}

func (s *CreateLabReservationRequest) SetLabId(v int64) *CreateLabReservationRequest {
	s.LabId = &v
	return s
}

func (s *CreateLabReservationRequest) SetMemberCount(v int64) *CreateLabReservationRequest {
	s.MemberCount = &v
	return s
}

func (s *CreateLabReservationRequest) SetStartTime(v string) *CreateLabReservationRequest {
	s.StartTime = &v
	return s
}

type CreateLabReservationResponseBody struct {
	// example:
	//
	// 00000
	Code           *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	LabReservation *CreateLabReservationResponseBodyLabReservation `json:"LabReservation,omitempty" xml:"LabReservation,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9ADC729B-...
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLabReservationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLabReservationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLabReservationResponseBody) SetCode(v string) *CreateLabReservationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLabReservationResponseBody) SetLabReservation(v *CreateLabReservationResponseBodyLabReservation) *CreateLabReservationResponseBody {
	s.LabReservation = v
	return s
}

func (s *CreateLabReservationResponseBody) SetMessage(v string) *CreateLabReservationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLabReservationResponseBody) SetRequestId(v string) *CreateLabReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLabReservationResponseBody) SetSuccess(v bool) *CreateLabReservationResponseBody {
	s.Success = &v
	return s
}

type CreateLabReservationResponseBodyLabReservation struct {
	// example:
	//
	// 10dbb36c-7047-41c7-92ee-bd24bcf45dca
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateLabReservationResponseBodyLabReservation) String() string {
	return tea.Prettify(s)
}

func (s CreateLabReservationResponseBodyLabReservation) GoString() string {
	return s.String()
}

func (s *CreateLabReservationResponseBodyLabReservation) SetId(v string) *CreateLabReservationResponseBodyLabReservation {
	s.Id = &v
	return s
}

type CreateLabReservationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLabReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLabReservationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLabReservationResponse) GoString() string {
	return s.String()
}

func (s *CreateLabReservationResponse) SetHeaders(v map[string]*string) *CreateLabReservationResponse {
	s.Headers = v
	return s
}

func (s *CreateLabReservationResponse) SetStatusCode(v int32) *CreateLabReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLabReservationResponse) SetBody(v *CreateLabReservationResponseBody) *CreateLabReservationResponse {
	s.Body = v
	return s
}

type CreateLabSessionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// -
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 875
	LabId *int64 `json:"LabId,omitempty" xml:"LabId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -
	RamAccountId *int64 `json:"RamAccountId,omitempty" xml:"RamAccountId,omitempty"`
}

func (s CreateLabSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLabSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateLabSessionRequest) SetAccountId(v int64) *CreateLabSessionRequest {
	s.AccountId = &v
	return s
}

func (s *CreateLabSessionRequest) SetLabId(v int64) *CreateLabSessionRequest {
	s.LabId = &v
	return s
}

func (s *CreateLabSessionRequest) SetRamAccountId(v int64) *CreateLabSessionRequest {
	s.RamAccountId = &v
	return s
}

type CreateLabSessionResponseBody struct {
	// example:
	//
	// 00000
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	LabSession *CreateLabSessionResponseBodyLabSession `json:"LabSession,omitempty" xml:"LabSession,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9ADC729B-...
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLabSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLabSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLabSessionResponseBody) SetCode(v string) *CreateLabSessionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLabSessionResponseBody) SetLabSession(v *CreateLabSessionResponseBodyLabSession) *CreateLabSessionResponseBody {
	s.LabSession = v
	return s
}

func (s *CreateLabSessionResponseBody) SetMessage(v string) *CreateLabSessionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLabSessionResponseBody) SetRequestId(v string) *CreateLabSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLabSessionResponseBody) SetSuccess(v bool) *CreateLabSessionResponseBody {
	s.Success = &v
	return s
}

type CreateLabSessionResponseBodyLabSession struct {
	// example:
	//
	// 875
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// -
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateLabSessionResponseBodyLabSession) String() string {
	return tea.Prettify(s)
}

func (s CreateLabSessionResponseBodyLabSession) GoString() string {
	return s.String()
}

func (s *CreateLabSessionResponseBodyLabSession) SetId(v string) *CreateLabSessionResponseBodyLabSession {
	s.Id = &v
	return s
}

func (s *CreateLabSessionResponseBodyLabSession) SetUrl(v string) *CreateLabSessionResponseBodyLabSession {
	s.Url = &v
	return s
}

type CreateLabSessionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLabSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLabSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLabSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateLabSessionResponse) SetHeaders(v map[string]*string) *CreateLabSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateLabSessionResponse) SetStatusCode(v int32) *CreateLabSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLabSessionResponse) SetBody(v *CreateLabSessionResponseBody) *CreateLabSessionResponse {
	s.Body = v
	return s
}

type DescribeLabRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 875
	LabId *int64 `json:"LabId,omitempty" xml:"LabId,omitempty"`
}

func (s DescribeLabRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabRequest) GoString() string {
	return s.String()
}

func (s *DescribeLabRequest) SetLabId(v int64) *DescribeLabRequest {
	s.LabId = &v
	return s
}

type DescribeLabResponseBody struct {
	// example:
	//
	// 00000
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Lab  *DescribeLabResponseBodyLab `json:"Lab,omitempty" xml:"Lab,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9ADC729B-...
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLabResponseBody) SetCode(v string) *DescribeLabResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLabResponseBody) SetLab(v *DescribeLabResponseBodyLab) *DescribeLabResponseBody {
	s.Lab = v
	return s
}

func (s *DescribeLabResponseBody) SetMessage(v string) *DescribeLabResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLabResponseBody) SetRequestId(v string) *DescribeLabResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLabResponseBody) SetSuccess(v bool) *DescribeLabResponseBody {
	s.Success = &v
	return s
}

type DescribeLabResponseBodyLab struct {
	// example:
	//
	// 10
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 875
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// -
	SubTitle *string `json:"SubTitle,omitempty" xml:"SubTitle,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeLabResponseBodyLab) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabResponseBodyLab) GoString() string {
	return s.String()
}

func (s *DescribeLabResponseBodyLab) SetDuration(v int64) *DescribeLabResponseBodyLab {
	s.Duration = &v
	return s
}

func (s *DescribeLabResponseBodyLab) SetId(v int64) *DescribeLabResponseBodyLab {
	s.Id = &v
	return s
}

func (s *DescribeLabResponseBodyLab) SetIntroduction(v string) *DescribeLabResponseBodyLab {
	s.Introduction = &v
	return s
}

func (s *DescribeLabResponseBodyLab) SetSubTitle(v string) *DescribeLabResponseBodyLab {
	s.SubTitle = &v
	return s
}

func (s *DescribeLabResponseBodyLab) SetTitle(v string) *DescribeLabResponseBodyLab {
	s.Title = &v
	return s
}

type DescribeLabResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLabResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabResponse) GoString() string {
	return s.String()
}

func (s *DescribeLabResponse) SetHeaders(v map[string]*string) *DescribeLabResponse {
	s.Headers = v
	return s
}

func (s *DescribeLabResponse) SetStatusCode(v int32) *DescribeLabResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLabResponse) SetBody(v *DescribeLabResponseBody) *DescribeLabResponse {
	s.Body = v
	return s
}

type DescribeLabReservationRequest struct {
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LabReservationId *string `json:"LabReservationId,omitempty" xml:"LabReservationId,omitempty"`
}

func (s DescribeLabReservationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabReservationRequest) GoString() string {
	return s.String()
}

func (s *DescribeLabReservationRequest) SetAccountId(v int64) *DescribeLabReservationRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeLabReservationRequest) SetLabReservationId(v string) *DescribeLabReservationRequest {
	s.LabReservationId = &v
	return s
}

type DescribeLabReservationResponseBody struct {
	// example:
	//
	// 00000
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	LabReservation *DescribeLabReservationResponseBodyLabReservation `json:"LabReservation,omitempty" xml:"LabReservation,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9ADC729B-...
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLabReservationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabReservationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLabReservationResponseBody) SetCode(v string) *DescribeLabReservationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLabReservationResponseBody) SetLabReservation(v *DescribeLabReservationResponseBodyLabReservation) *DescribeLabReservationResponseBody {
	s.LabReservation = v
	return s
}

func (s *DescribeLabReservationResponseBody) SetMessage(v string) *DescribeLabReservationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLabReservationResponseBody) SetRequestId(v string) *DescribeLabReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLabReservationResponseBody) SetSuccess(v bool) *DescribeLabReservationResponseBody {
	s.Success = &v
	return s
}

type DescribeLabReservationResponseBodyLabReservation struct {
	AccountId *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MemberCount *int64  `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLabReservationResponseBodyLabReservation) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabReservationResponseBodyLabReservation) GoString() string {
	return s.String()
}

func (s *DescribeLabReservationResponseBodyLabReservation) SetAccountId(v int64) *DescribeLabReservationResponseBodyLabReservation {
	s.AccountId = &v
	return s
}

func (s *DescribeLabReservationResponseBodyLabReservation) SetEndTime(v string) *DescribeLabReservationResponseBodyLabReservation {
	s.EndTime = &v
	return s
}

func (s *DescribeLabReservationResponseBodyLabReservation) SetId(v string) *DescribeLabReservationResponseBodyLabReservation {
	s.Id = &v
	return s
}

func (s *DescribeLabReservationResponseBodyLabReservation) SetMemberCount(v int64) *DescribeLabReservationResponseBodyLabReservation {
	s.MemberCount = &v
	return s
}

func (s *DescribeLabReservationResponseBodyLabReservation) SetStartTime(v string) *DescribeLabReservationResponseBodyLabReservation {
	s.StartTime = &v
	return s
}

type DescribeLabReservationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLabReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLabReservationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabReservationResponse) GoString() string {
	return s.String()
}

func (s *DescribeLabReservationResponse) SetHeaders(v map[string]*string) *DescribeLabReservationResponse {
	s.Headers = v
	return s
}

func (s *DescribeLabReservationResponse) SetStatusCode(v int32) *DescribeLabReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLabReservationResponse) SetBody(v *DescribeLabReservationResponseBody) *DescribeLabReservationResponse {
	s.Body = v
	return s
}

type DescribeLabSessionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// -
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LabSessionId *string `json:"LabSessionId,omitempty" xml:"LabSessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -
	RamAccountId *int64 `json:"RamAccountId,omitempty" xml:"RamAccountId,omitempty"`
}

func (s DescribeLabSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabSessionRequest) GoString() string {
	return s.String()
}

func (s *DescribeLabSessionRequest) SetAccountId(v int64) *DescribeLabSessionRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeLabSessionRequest) SetLabSessionId(v string) *DescribeLabSessionRequest {
	s.LabSessionId = &v
	return s
}

func (s *DescribeLabSessionRequest) SetRamAccountId(v int64) *DescribeLabSessionRequest {
	s.RamAccountId = &v
	return s
}

type DescribeLabSessionResponseBody struct {
	// example:
	//
	// 00000
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	LabSession *DescribeLabSessionResponseBodyLabSession `json:"LabSession,omitempty" xml:"LabSession,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9ADC729B-...
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLabSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLabSessionResponseBody) SetCode(v string) *DescribeLabSessionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLabSessionResponseBody) SetLabSession(v *DescribeLabSessionResponseBodyLabSession) *DescribeLabSessionResponseBody {
	s.LabSession = v
	return s
}

func (s *DescribeLabSessionResponseBody) SetMessage(v string) *DescribeLabSessionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLabSessionResponseBody) SetRequestId(v string) *DescribeLabSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLabSessionResponseBody) SetSuccess(v bool) *DescribeLabSessionResponseBody {
	s.Success = &v
	return s
}

type DescribeLabSessionResponseBodyLabSession struct {
	// example:
	//
	// 2023-05-05 15:01:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// False
	Finished *bool `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// example:
	//
	// 838
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 875
	LabId *int64 `json:"LabId,omitempty" xml:"LabId,omitempty"`
	// example:
	//
	// -
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeLabSessionResponseBodyLabSession) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabSessionResponseBodyLabSession) GoString() string {
	return s.String()
}

func (s *DescribeLabSessionResponseBodyLabSession) SetCreateTime(v string) *DescribeLabSessionResponseBodyLabSession {
	s.CreateTime = &v
	return s
}

func (s *DescribeLabSessionResponseBodyLabSession) SetFinished(v bool) *DescribeLabSessionResponseBodyLabSession {
	s.Finished = &v
	return s
}

func (s *DescribeLabSessionResponseBodyLabSession) SetId(v string) *DescribeLabSessionResponseBodyLabSession {
	s.Id = &v
	return s
}

func (s *DescribeLabSessionResponseBodyLabSession) SetLabId(v int64) *DescribeLabSessionResponseBodyLabSession {
	s.LabId = &v
	return s
}

func (s *DescribeLabSessionResponseBodyLabSession) SetUrl(v string) *DescribeLabSessionResponseBodyLabSession {
	s.Url = &v
	return s
}

type DescribeLabSessionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLabSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLabSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLabSessionResponse) GoString() string {
	return s.String()
}

func (s *DescribeLabSessionResponse) SetHeaders(v map[string]*string) *DescribeLabSessionResponse {
	s.Headers = v
	return s
}

func (s *DescribeLabSessionResponse) SetStatusCode(v int32) *DescribeLabSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLabSessionResponse) SetBody(v *DescribeLabSessionResponseBody) *DescribeLabSessionResponse {
	s.Body = v
	return s
}

type PageListLabReservationsRequest struct {
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s PageListLabReservationsRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListLabReservationsRequest) GoString() string {
	return s.String()
}

func (s *PageListLabReservationsRequest) SetAccountId(v int64) *PageListLabReservationsRequest {
	s.AccountId = &v
	return s
}

func (s *PageListLabReservationsRequest) SetPage(v int64) *PageListLabReservationsRequest {
	s.Page = &v
	return s
}

func (s *PageListLabReservationsRequest) SetPageSize(v int64) *PageListLabReservationsRequest {
	s.PageSize = &v
	return s
}

type PageListLabReservationsResponseBody struct {
	// example:
	//
	// 00000
	Code            *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	LabReservations []*PageListLabReservationsResponseBodyLabReservations `json:"LabReservations,omitempty" xml:"LabReservations,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9ADC729B-...
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 32
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageListLabReservationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListLabReservationsResponseBody) GoString() string {
	return s.String()
}

func (s *PageListLabReservationsResponseBody) SetCode(v string) *PageListLabReservationsResponseBody {
	s.Code = &v
	return s
}

func (s *PageListLabReservationsResponseBody) SetLabReservations(v []*PageListLabReservationsResponseBodyLabReservations) *PageListLabReservationsResponseBody {
	s.LabReservations = v
	return s
}

func (s *PageListLabReservationsResponseBody) SetMessage(v string) *PageListLabReservationsResponseBody {
	s.Message = &v
	return s
}

func (s *PageListLabReservationsResponseBody) SetPage(v int64) *PageListLabReservationsResponseBody {
	s.Page = &v
	return s
}

func (s *PageListLabReservationsResponseBody) SetPageSize(v int64) *PageListLabReservationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *PageListLabReservationsResponseBody) SetRequestId(v string) *PageListLabReservationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageListLabReservationsResponseBody) SetSuccess(v bool) *PageListLabReservationsResponseBody {
	s.Success = &v
	return s
}

func (s *PageListLabReservationsResponseBody) SetTotalCount(v int64) *PageListLabReservationsResponseBody {
	s.TotalCount = &v
	return s
}

type PageListLabReservationsResponseBodyLabReservations struct {
	// example:
	//
	// -
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 2023-05-04 09:25:41
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 552
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	MemberCount *int64 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	// example:
	//
	// 2023-05-04 07:25:41
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PageListLabReservationsResponseBodyLabReservations) String() string {
	return tea.Prettify(s)
}

func (s PageListLabReservationsResponseBodyLabReservations) GoString() string {
	return s.String()
}

func (s *PageListLabReservationsResponseBodyLabReservations) SetAccountId(v int64) *PageListLabReservationsResponseBodyLabReservations {
	s.AccountId = &v
	return s
}

func (s *PageListLabReservationsResponseBodyLabReservations) SetEndTime(v string) *PageListLabReservationsResponseBodyLabReservations {
	s.EndTime = &v
	return s
}

func (s *PageListLabReservationsResponseBodyLabReservations) SetId(v string) *PageListLabReservationsResponseBodyLabReservations {
	s.Id = &v
	return s
}

func (s *PageListLabReservationsResponseBodyLabReservations) SetMemberCount(v int64) *PageListLabReservationsResponseBodyLabReservations {
	s.MemberCount = &v
	return s
}

func (s *PageListLabReservationsResponseBodyLabReservations) SetStartTime(v string) *PageListLabReservationsResponseBodyLabReservations {
	s.StartTime = &v
	return s
}

type PageListLabReservationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageListLabReservationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageListLabReservationsResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListLabReservationsResponse) GoString() string {
	return s.String()
}

func (s *PageListLabReservationsResponse) SetHeaders(v map[string]*string) *PageListLabReservationsResponse {
	s.Headers = v
	return s
}

func (s *PageListLabReservationsResponse) SetStatusCode(v int32) *PageListLabReservationsResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListLabReservationsResponse) SetBody(v *PageListLabReservationsResponseBody) *PageListLabReservationsResponse {
	s.Body = v
	return s
}

type PageListLabSessionsRequest struct {
	Finished *bool  `json:"Finished,omitempty" xml:"Finished,omitempty"`
	LabId    *int64 `json:"LabId,omitempty" xml:"LabId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize     *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RamAccountId *int64 `json:"RamAccountId,omitempty" xml:"RamAccountId,omitempty"`
}

func (s PageListLabSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListLabSessionsRequest) GoString() string {
	return s.String()
}

func (s *PageListLabSessionsRequest) SetFinished(v bool) *PageListLabSessionsRequest {
	s.Finished = &v
	return s
}

func (s *PageListLabSessionsRequest) SetLabId(v int64) *PageListLabSessionsRequest {
	s.LabId = &v
	return s
}

func (s *PageListLabSessionsRequest) SetPage(v int64) *PageListLabSessionsRequest {
	s.Page = &v
	return s
}

func (s *PageListLabSessionsRequest) SetPageSize(v int64) *PageListLabSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *PageListLabSessionsRequest) SetRamAccountId(v int64) *PageListLabSessionsRequest {
	s.RamAccountId = &v
	return s
}

type PageListLabSessionsResponseBody struct {
	// example:
	//
	// 00000
	Code        *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	LabSessions []*PageListLabSessionsResponseBodyLabSessions `json:"LabSessions,omitempty" xml:"LabSessions,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9ADC729B-...
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageListLabSessionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListLabSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *PageListLabSessionsResponseBody) SetCode(v string) *PageListLabSessionsResponseBody {
	s.Code = &v
	return s
}

func (s *PageListLabSessionsResponseBody) SetLabSessions(v []*PageListLabSessionsResponseBodyLabSessions) *PageListLabSessionsResponseBody {
	s.LabSessions = v
	return s
}

func (s *PageListLabSessionsResponseBody) SetMessage(v string) *PageListLabSessionsResponseBody {
	s.Message = &v
	return s
}

func (s *PageListLabSessionsResponseBody) SetPage(v int64) *PageListLabSessionsResponseBody {
	s.Page = &v
	return s
}

func (s *PageListLabSessionsResponseBody) SetPageSize(v int64) *PageListLabSessionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *PageListLabSessionsResponseBody) SetRequestId(v string) *PageListLabSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageListLabSessionsResponseBody) SetSuccess(v bool) *PageListLabSessionsResponseBody {
	s.Success = &v
	return s
}

func (s *PageListLabSessionsResponseBody) SetTotalCount(v int64) *PageListLabSessionsResponseBody {
	s.TotalCount = &v
	return s
}

type PageListLabSessionsResponseBodyLabSessions struct {
	// example:
	//
	// 2023-06-24 18:27:25
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// False
	Finished *bool `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// example:
	//
	// 10334
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 875
	LabId *int64 `json:"LabId,omitempty" xml:"LabId,omitempty"`
	// example:
	//
	// -
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PageListLabSessionsResponseBodyLabSessions) String() string {
	return tea.Prettify(s)
}

func (s PageListLabSessionsResponseBodyLabSessions) GoString() string {
	return s.String()
}

func (s *PageListLabSessionsResponseBodyLabSessions) SetCreateTime(v string) *PageListLabSessionsResponseBodyLabSessions {
	s.CreateTime = &v
	return s
}

func (s *PageListLabSessionsResponseBodyLabSessions) SetFinished(v bool) *PageListLabSessionsResponseBodyLabSessions {
	s.Finished = &v
	return s
}

func (s *PageListLabSessionsResponseBodyLabSessions) SetId(v string) *PageListLabSessionsResponseBodyLabSessions {
	s.Id = &v
	return s
}

func (s *PageListLabSessionsResponseBodyLabSessions) SetLabId(v int64) *PageListLabSessionsResponseBodyLabSessions {
	s.LabId = &v
	return s
}

func (s *PageListLabSessionsResponseBodyLabSessions) SetUrl(v string) *PageListLabSessionsResponseBodyLabSessions {
	s.Url = &v
	return s
}

type PageListLabSessionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageListLabSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageListLabSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListLabSessionsResponse) GoString() string {
	return s.String()
}

func (s *PageListLabSessionsResponse) SetHeaders(v map[string]*string) *PageListLabSessionsResponse {
	s.Headers = v
	return s
}

func (s *PageListLabSessionsResponse) SetStatusCode(v int32) *PageListLabSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListLabSessionsResponse) SetBody(v *PageListLabSessionsResponseBody) *PageListLabSessionsResponse {
	s.Body = v
	return s
}

type PageListLabsRequest struct {
	Id []*int64 `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s PageListLabsRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListLabsRequest) GoString() string {
	return s.String()
}

func (s *PageListLabsRequest) SetId(v []*int64) *PageListLabsRequest {
	s.Id = v
	return s
}

func (s *PageListLabsRequest) SetPage(v int64) *PageListLabsRequest {
	s.Page = &v
	return s
}

func (s *PageListLabsRequest) SetPageSize(v int64) *PageListLabsRequest {
	s.PageSize = &v
	return s
}

type PageListLabsResponseBody struct {
	// example:
	//
	// 00000
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Labs []*PageListLabsResponseBodyLabs `json:"Labs,omitempty" xml:"Labs,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9ADC729B-...
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageListLabsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListLabsResponseBody) GoString() string {
	return s.String()
}

func (s *PageListLabsResponseBody) SetCode(v string) *PageListLabsResponseBody {
	s.Code = &v
	return s
}

func (s *PageListLabsResponseBody) SetLabs(v []*PageListLabsResponseBodyLabs) *PageListLabsResponseBody {
	s.Labs = v
	return s
}

func (s *PageListLabsResponseBody) SetMessage(v string) *PageListLabsResponseBody {
	s.Message = &v
	return s
}

func (s *PageListLabsResponseBody) SetPage(v int64) *PageListLabsResponseBody {
	s.Page = &v
	return s
}

func (s *PageListLabsResponseBody) SetPageSize(v int64) *PageListLabsResponseBody {
	s.PageSize = &v
	return s
}

func (s *PageListLabsResponseBody) SetRequestId(v string) *PageListLabsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageListLabsResponseBody) SetSuccess(v bool) *PageListLabsResponseBody {
	s.Success = &v
	return s
}

func (s *PageListLabsResponseBody) SetTotalCount(v int64) *PageListLabsResponseBody {
	s.TotalCount = &v
	return s
}

type PageListLabsResponseBodyLabs struct {
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 981
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// -
	SubTitle *string `json:"SubTitle,omitempty" xml:"SubTitle,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PageListLabsResponseBodyLabs) String() string {
	return tea.Prettify(s)
}

func (s PageListLabsResponseBodyLabs) GoString() string {
	return s.String()
}

func (s *PageListLabsResponseBodyLabs) SetDuration(v int64) *PageListLabsResponseBodyLabs {
	s.Duration = &v
	return s
}

func (s *PageListLabsResponseBodyLabs) SetId(v int64) *PageListLabsResponseBodyLabs {
	s.Id = &v
	return s
}

func (s *PageListLabsResponseBodyLabs) SetIntroduction(v string) *PageListLabsResponseBodyLabs {
	s.Introduction = &v
	return s
}

func (s *PageListLabsResponseBodyLabs) SetSubTitle(v string) *PageListLabsResponseBodyLabs {
	s.SubTitle = &v
	return s
}

func (s *PageListLabsResponseBodyLabs) SetTitle(v string) *PageListLabsResponseBodyLabs {
	s.Title = &v
	return s
}

type PageListLabsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageListLabsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageListLabsResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListLabsResponse) GoString() string {
	return s.String()
}

func (s *PageListLabsResponse) SetHeaders(v map[string]*string) *PageListLabsResponse {
	s.Headers = v
	return s
}

func (s *PageListLabsResponse) SetStatusCode(v int32) *PageListLabsResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListLabsResponse) SetBody(v *PageListLabsResponseBody) *PageListLabsResponse {
	s.Body = v
	return s
}

type RemoveLabReservationRequest struct {
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LabReservationId *string `json:"LabReservationId,omitempty" xml:"LabReservationId,omitempty"`
}

func (s RemoveLabReservationRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveLabReservationRequest) GoString() string {
	return s.String()
}

func (s *RemoveLabReservationRequest) SetAccountId(v int64) *RemoveLabReservationRequest {
	s.AccountId = &v
	return s
}

func (s *RemoveLabReservationRequest) SetLabReservationId(v string) *RemoveLabReservationRequest {
	s.LabReservationId = &v
	return s
}

type RemoveLabReservationResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9ADC729B-...
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveLabReservationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveLabReservationResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveLabReservationResponseBody) SetCode(v string) *RemoveLabReservationResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveLabReservationResponseBody) SetMessage(v string) *RemoveLabReservationResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveLabReservationResponseBody) SetRequestId(v string) *RemoveLabReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveLabReservationResponseBody) SetResult(v bool) *RemoveLabReservationResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveLabReservationResponseBody) SetSuccess(v bool) *RemoveLabReservationResponseBody {
	s.Success = &v
	return s
}

type RemoveLabReservationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveLabReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveLabReservationResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveLabReservationResponse) GoString() string {
	return s.String()
}

func (s *RemoveLabReservationResponse) SetHeaders(v map[string]*string) *RemoveLabReservationResponse {
	s.Headers = v
	return s
}

func (s *RemoveLabReservationResponse) SetStatusCode(v int32) *RemoveLabReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveLabReservationResponse) SetBody(v *RemoveLabReservationResponseBody) *RemoveLabReservationResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("eduembed"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 创建实验预约
//
// @param request - CreateLabReservationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLabReservationResponse
func (client *Client) CreateLabReservationWithOptions(request *CreateLabReservationRequest, runtime *util.RuntimeOptions) (_result *CreateLabReservationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.LabId)) {
		body["LabId"] = request.LabId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberCount)) {
		body["MemberCount"] = request.MemberCount
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLabReservation"),
		Version:     tea.String("2024-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLabReservationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验预约
//
// @param request - CreateLabReservationRequest
//
// @return CreateLabReservationResponse
func (client *Client) CreateLabReservation(request *CreateLabReservationRequest) (_result *CreateLabReservationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLabReservationResponse{}
	_body, _err := client.CreateLabReservationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实验会话
//
// @param request - CreateLabSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLabSessionResponse
func (client *Client) CreateLabSessionWithOptions(request *CreateLabSessionRequest, runtime *util.RuntimeOptions) (_result *CreateLabSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.LabId)) {
		body["LabId"] = request.LabId
	}

	if !tea.BoolValue(util.IsUnset(request.RamAccountId)) {
		body["RamAccountId"] = request.RamAccountId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLabSession"),
		Version:     tea.String("2024-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLabSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验会话
//
// @param request - CreateLabSessionRequest
//
// @return CreateLabSessionResponse
func (client *Client) CreateLabSession(request *CreateLabSessionRequest) (_result *CreateLabSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLabSessionResponse{}
	_body, _err := client.CreateLabSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看实验详情
//
// @param request - DescribeLabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLabResponse
func (client *Client) DescribeLabWithOptions(request *DescribeLabRequest, runtime *util.RuntimeOptions) (_result *DescribeLabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLab"),
		Version:     tea.String("2024-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看实验详情
//
// @param request - DescribeLabRequest
//
// @return DescribeLabResponse
func (client *Client) DescribeLab(request *DescribeLabRequest) (_result *DescribeLabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLabResponse{}
	_body, _err := client.DescribeLabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实验预约
//
// @param request - DescribeLabReservationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLabReservationResponse
func (client *Client) DescribeLabReservationWithOptions(request *DescribeLabReservationRequest, runtime *util.RuntimeOptions) (_result *DescribeLabReservationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLabReservation"),
		Version:     tea.String("2024-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLabReservationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实验预约
//
// @param request - DescribeLabReservationRequest
//
// @return DescribeLabReservationResponse
func (client *Client) DescribeLabReservation(request *DescribeLabReservationRequest) (_result *DescribeLabReservationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLabReservationResponse{}
	_body, _err := client.DescribeLabReservationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看实验会话信息
//
// @param request - DescribeLabSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLabSessionResponse
func (client *Client) DescribeLabSessionWithOptions(request *DescribeLabSessionRequest, runtime *util.RuntimeOptions) (_result *DescribeLabSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLabSession"),
		Version:     tea.String("2024-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLabSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看实验会话信息
//
// @param request - DescribeLabSessionRequest
//
// @return DescribeLabSessionResponse
func (client *Client) DescribeLabSession(request *DescribeLabSessionRequest) (_result *DescribeLabSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLabSessionResponse{}
	_body, _err := client.DescribeLabSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询实验预约
//
// @param request - PageListLabReservationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageListLabReservationsResponse
func (client *Client) PageListLabReservationsWithOptions(request *PageListLabReservationsRequest, runtime *util.RuntimeOptions) (_result *PageListLabReservationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PageListLabReservations"),
		Version:     tea.String("2024-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageListLabReservationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询实验预约
//
// @param request - PageListLabReservationsRequest
//
// @return PageListLabReservationsResponse
func (client *Client) PageListLabReservations(request *PageListLabReservationsRequest) (_result *PageListLabReservationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageListLabReservationsResponse{}
	_body, _err := client.PageListLabReservationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询实验会话
//
// @param request - PageListLabSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageListLabSessionsResponse
func (client *Client) PageListLabSessionsWithOptions(request *PageListLabSessionsRequest, runtime *util.RuntimeOptions) (_result *PageListLabSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PageListLabSessions"),
		Version:     tea.String("2024-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageListLabSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询实验会话
//
// @param request - PageListLabSessionsRequest
//
// @return PageListLabSessionsResponse
func (client *Client) PageListLabSessions(request *PageListLabSessionsRequest) (_result *PageListLabSessionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageListLabSessionsResponse{}
	_body, _err := client.PageListLabSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询实验
//
// @param request - PageListLabsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageListLabsResponse
func (client *Client) PageListLabsWithOptions(request *PageListLabsRequest, runtime *util.RuntimeOptions) (_result *PageListLabsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PageListLabs"),
		Version:     tea.String("2024-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageListLabsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询实验
//
// @param request - PageListLabsRequest
//
// @return PageListLabsResponse
func (client *Client) PageListLabs(request *PageListLabsRequest) (_result *PageListLabsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageListLabsResponse{}
	_body, _err := client.PageListLabsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除实验预约
//
// @param request - RemoveLabReservationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveLabReservationResponse
func (client *Client) RemoveLabReservationWithOptions(request *RemoveLabReservationRequest, runtime *util.RuntimeOptions) (_result *RemoveLabReservationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.LabReservationId)) {
		body["LabReservationId"] = request.LabReservationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveLabReservation"),
		Version:     tea.String("2024-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveLabReservationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除实验预约
//
// @param request - RemoveLabReservationRequest
//
// @return RemoveLabReservationResponse
func (client *Client) RemoveLabReservation(request *RemoveLabReservationRequest) (_result *RemoveLabReservationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveLabReservationResponse{}
	_body, _err := client.RemoveLabReservationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
