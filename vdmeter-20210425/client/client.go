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

type DescribeHuYaRecordByDomainRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeHuYaRecordByDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaRecordByDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeHuYaRecordByDomainRequest) SetBusinessType(v string) *DescribeHuYaRecordByDomainRequest {
	s.BusinessType = &v
	return s
}

func (s *DescribeHuYaRecordByDomainRequest) SetDomain(v string) *DescribeHuYaRecordByDomainRequest {
	s.Domain = &v
	return s
}

func (s *DescribeHuYaRecordByDomainRequest) SetEndTime(v string) *DescribeHuYaRecordByDomainRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHuYaRecordByDomainRequest) SetStartTime(v string) *DescribeHuYaRecordByDomainRequest {
	s.StartTime = &v
	return s
}

type DescribeHuYaRecordByDomainResponseBody struct {
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     *string                                             `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultDesc []*DescribeHuYaRecordByDomainResponseBodyResultDesc `json:"ResultDesc,omitempty" xml:"ResultDesc,omitempty" type:"Repeated"`
	Status     *int32                                              `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHuYaRecordByDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaRecordByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHuYaRecordByDomainResponseBody) SetRequestId(v string) *DescribeHuYaRecordByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHuYaRecordByDomainResponseBody) SetResult(v string) *DescribeHuYaRecordByDomainResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeHuYaRecordByDomainResponseBody) SetResultDesc(v []*DescribeHuYaRecordByDomainResponseBodyResultDesc) *DescribeHuYaRecordByDomainResponseBody {
	s.ResultDesc = v
	return s
}

func (s *DescribeHuYaRecordByDomainResponseBody) SetStatus(v int32) *DescribeHuYaRecordByDomainResponseBody {
	s.Status = &v
	return s
}

type DescribeHuYaRecordByDomainResponseBodyResultDesc struct {
	BusinessType   *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RecordDuration *int64  `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	RecordNum      *int64  `json:"RecordNum,omitempty" xml:"RecordNum,omitempty"`
	RecordType     *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Time           *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeHuYaRecordByDomainResponseBodyResultDesc) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaRecordByDomainResponseBodyResultDesc) GoString() string {
	return s.String()
}

func (s *DescribeHuYaRecordByDomainResponseBodyResultDesc) SetBusinessType(v string) *DescribeHuYaRecordByDomainResponseBodyResultDesc {
	s.BusinessType = &v
	return s
}

func (s *DescribeHuYaRecordByDomainResponseBodyResultDesc) SetDomain(v string) *DescribeHuYaRecordByDomainResponseBodyResultDesc {
	s.Domain = &v
	return s
}

func (s *DescribeHuYaRecordByDomainResponseBodyResultDesc) SetRecordDuration(v int64) *DescribeHuYaRecordByDomainResponseBodyResultDesc {
	s.RecordDuration = &v
	return s
}

func (s *DescribeHuYaRecordByDomainResponseBodyResultDesc) SetRecordNum(v int64) *DescribeHuYaRecordByDomainResponseBodyResultDesc {
	s.RecordNum = &v
	return s
}

func (s *DescribeHuYaRecordByDomainResponseBodyResultDesc) SetRecordType(v string) *DescribeHuYaRecordByDomainResponseBodyResultDesc {
	s.RecordType = &v
	return s
}

func (s *DescribeHuYaRecordByDomainResponseBodyResultDesc) SetTime(v string) *DescribeHuYaRecordByDomainResponseBodyResultDesc {
	s.Time = &v
	return s
}

type DescribeHuYaRecordByDomainResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHuYaRecordByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHuYaRecordByDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaRecordByDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeHuYaRecordByDomainResponse) SetHeaders(v map[string]*string) *DescribeHuYaRecordByDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeHuYaRecordByDomainResponse) SetStatusCode(v int32) *DescribeHuYaRecordByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHuYaRecordByDomainResponse) SetBody(v *DescribeHuYaRecordByDomainResponseBody) *DescribeHuYaRecordByDomainResponse {
	s.Body = v
	return s
}

type DescribeHuYaScreenshotByDomainRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeHuYaScreenshotByDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaScreenshotByDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeHuYaScreenshotByDomainRequest) SetBusinessType(v string) *DescribeHuYaScreenshotByDomainRequest {
	s.BusinessType = &v
	return s
}

func (s *DescribeHuYaScreenshotByDomainRequest) SetDomain(v string) *DescribeHuYaScreenshotByDomainRequest {
	s.Domain = &v
	return s
}

func (s *DescribeHuYaScreenshotByDomainRequest) SetEndTime(v string) *DescribeHuYaScreenshotByDomainRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHuYaScreenshotByDomainRequest) SetStartTime(v string) *DescribeHuYaScreenshotByDomainRequest {
	s.StartTime = &v
	return s
}

type DescribeHuYaScreenshotByDomainResponseBody struct {
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     *string                                                 `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultDesc []*DescribeHuYaScreenshotByDomainResponseBodyResultDesc `json:"ResultDesc,omitempty" xml:"ResultDesc,omitempty" type:"Repeated"`
	Status     *int32                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHuYaScreenshotByDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaScreenshotByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHuYaScreenshotByDomainResponseBody) SetRequestId(v string) *DescribeHuYaScreenshotByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHuYaScreenshotByDomainResponseBody) SetResult(v string) *DescribeHuYaScreenshotByDomainResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeHuYaScreenshotByDomainResponseBody) SetResultDesc(v []*DescribeHuYaScreenshotByDomainResponseBodyResultDesc) *DescribeHuYaScreenshotByDomainResponseBody {
	s.ResultDesc = v
	return s
}

func (s *DescribeHuYaScreenshotByDomainResponseBody) SetStatus(v int32) *DescribeHuYaScreenshotByDomainResponseBody {
	s.Status = &v
	return s
}

type DescribeHuYaScreenshotByDomainResponseBodyResultDesc struct {
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ScreenshotNum *int64  `json:"ScreenshotNum,omitempty" xml:"ScreenshotNum,omitempty"`
	Time          *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeHuYaScreenshotByDomainResponseBodyResultDesc) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaScreenshotByDomainResponseBodyResultDesc) GoString() string {
	return s.String()
}

func (s *DescribeHuYaScreenshotByDomainResponseBodyResultDesc) SetBusinessType(v string) *DescribeHuYaScreenshotByDomainResponseBodyResultDesc {
	s.BusinessType = &v
	return s
}

func (s *DescribeHuYaScreenshotByDomainResponseBodyResultDesc) SetDomain(v string) *DescribeHuYaScreenshotByDomainResponseBodyResultDesc {
	s.Domain = &v
	return s
}

func (s *DescribeHuYaScreenshotByDomainResponseBodyResultDesc) SetScreenshotNum(v int64) *DescribeHuYaScreenshotByDomainResponseBodyResultDesc {
	s.ScreenshotNum = &v
	return s
}

func (s *DescribeHuYaScreenshotByDomainResponseBodyResultDesc) SetTime(v string) *DescribeHuYaScreenshotByDomainResponseBodyResultDesc {
	s.Time = &v
	return s
}

type DescribeHuYaScreenshotByDomainResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHuYaScreenshotByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHuYaScreenshotByDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaScreenshotByDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeHuYaScreenshotByDomainResponse) SetHeaders(v map[string]*string) *DescribeHuYaScreenshotByDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeHuYaScreenshotByDomainResponse) SetStatusCode(v int32) *DescribeHuYaScreenshotByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHuYaScreenshotByDomainResponse) SetBody(v *DescribeHuYaScreenshotByDomainResponseBody) *DescribeHuYaScreenshotByDomainResponse {
	s.Body = v
	return s
}

type DescribeHuYaTranscodeByDomainRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeHuYaTranscodeByDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaTranscodeByDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeHuYaTranscodeByDomainRequest) SetBusinessType(v string) *DescribeHuYaTranscodeByDomainRequest {
	s.BusinessType = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainRequest) SetDomain(v string) *DescribeHuYaTranscodeByDomainRequest {
	s.Domain = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainRequest) SetEndTime(v string) *DescribeHuYaTranscodeByDomainRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainRequest) SetStartTime(v string) *DescribeHuYaTranscodeByDomainRequest {
	s.StartTime = &v
	return s
}

type DescribeHuYaTranscodeByDomainResponseBody struct {
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     *string                                                `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultDesc []*DescribeHuYaTranscodeByDomainResponseBodyResultDesc `json:"ResultDesc,omitempty" xml:"ResultDesc,omitempty" type:"Repeated"`
	Status     *int32                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHuYaTranscodeByDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaTranscodeByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHuYaTranscodeByDomainResponseBody) SetRequestId(v string) *DescribeHuYaTranscodeByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponseBody) SetResult(v string) *DescribeHuYaTranscodeByDomainResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponseBody) SetResultDesc(v []*DescribeHuYaTranscodeByDomainResponseBodyResultDesc) *DescribeHuYaTranscodeByDomainResponseBody {
	s.ResultDesc = v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponseBody) SetStatus(v int32) *DescribeHuYaTranscodeByDomainResponseBody {
	s.Status = &v
	return s
}

type DescribeHuYaTranscodeByDomainResponseBodyResultDesc struct {
	BusinessType      *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Domain            *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Time              *string `json:"Time,omitempty" xml:"Time,omitempty"`
	TranscodeDuration *int64  `json:"TranscodeDuration,omitempty" xml:"TranscodeDuration,omitempty"`
	TranscodeNum      *int64  `json:"TranscodeNum,omitempty" xml:"TranscodeNum,omitempty"`
	TranscodeType     *string `json:"TranscodeType,omitempty" xml:"TranscodeType,omitempty"`
}

func (s DescribeHuYaTranscodeByDomainResponseBodyResultDesc) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaTranscodeByDomainResponseBodyResultDesc) GoString() string {
	return s.String()
}

func (s *DescribeHuYaTranscodeByDomainResponseBodyResultDesc) SetBusinessType(v string) *DescribeHuYaTranscodeByDomainResponseBodyResultDesc {
	s.BusinessType = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponseBodyResultDesc) SetDomain(v string) *DescribeHuYaTranscodeByDomainResponseBodyResultDesc {
	s.Domain = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponseBodyResultDesc) SetTime(v string) *DescribeHuYaTranscodeByDomainResponseBodyResultDesc {
	s.Time = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponseBodyResultDesc) SetTranscodeDuration(v int64) *DescribeHuYaTranscodeByDomainResponseBodyResultDesc {
	s.TranscodeDuration = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponseBodyResultDesc) SetTranscodeNum(v int64) *DescribeHuYaTranscodeByDomainResponseBodyResultDesc {
	s.TranscodeNum = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponseBodyResultDesc) SetTranscodeType(v string) *DescribeHuYaTranscodeByDomainResponseBodyResultDesc {
	s.TranscodeType = &v
	return s
}

type DescribeHuYaTranscodeByDomainResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHuYaTranscodeByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHuYaTranscodeByDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHuYaTranscodeByDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeHuYaTranscodeByDomainResponse) SetHeaders(v map[string]*string) *DescribeHuYaTranscodeByDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponse) SetStatusCode(v int32) *DescribeHuYaTranscodeByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHuYaTranscodeByDomainResponse) SetBody(v *DescribeHuYaTranscodeByDomainResponseBody) *DescribeHuYaTranscodeByDomainResponse {
	s.Body = v
	return s
}

type DescribeMeterBypassRtUsageByTaskProfileRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterBypassRtUsageByTaskProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterBypassRtUsageByTaskProfileRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterBypassRtUsageByTaskProfileRequest) SetAppId(v string) *DescribeMeterBypassRtUsageByTaskProfileRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileRequest) SetEndTs(v int64) *DescribeMeterBypassRtUsageByTaskProfileRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileRequest) SetInterval(v int64) *DescribeMeterBypassRtUsageByTaskProfileRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileRequest) SetServiceArea(v string) *DescribeMeterBypassRtUsageByTaskProfileRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileRequest) SetStartTs(v int64) *DescribeMeterBypassRtUsageByTaskProfileRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterBypassRtUsageByTaskProfileResponseBody struct {
	Data      []*DescribeMeterBypassRtUsageByTaskProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterBypassRtUsageByTaskProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterBypassRtUsageByTaskProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponseBody) SetData(v []*DescribeMeterBypassRtUsageByTaskProfileResponseBodyData) *DescribeMeterBypassRtUsageByTaskProfileResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponseBody) SetRequestId(v string) *DescribeMeterBypassRtUsageByTaskProfileResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterBypassRtUsageByTaskProfileResponseBodyData struct {
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Ratio       *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	Timestamp   *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeMeterBypassRtUsageByTaskProfileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterBypassRtUsageByTaskProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData) SetDuration(v int64) *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData {
	s.Duration = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData) SetRatio(v string) *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData) SetServiceArea(v string) *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData) SetTaskProfile(v string) *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData {
	s.TaskProfile = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData) SetTimestamp(v int64) *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData) SetType(v string) *DescribeMeterBypassRtUsageByTaskProfileResponseBodyData {
	s.Type = &v
	return s
}

type DescribeMeterBypassRtUsageByTaskProfileResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterBypassRtUsageByTaskProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterBypassRtUsageByTaskProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterBypassRtUsageByTaskProfileResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponse) SetHeaders(v map[string]*string) *DescribeMeterBypassRtUsageByTaskProfileResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponse) SetStatusCode(v int32) *DescribeMeterBypassRtUsageByTaskProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterBypassRtUsageByTaskProfileResponse) SetBody(v *DescribeMeterBypassRtUsageByTaskProfileResponseBody) *DescribeMeterBypassRtUsageByTaskProfileResponse {
	s.Body = v
	return s
}

type DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest) SetAppId(v string) *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest) SetEndTs(v int64) *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest) SetInterval(v int64) *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest) SetServiceArea(v string) *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest) SetStartTs(v int64) *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBody struct {
	Data      []*DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBody) SetData(v []*DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData) *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBody) SetRequestId(v string) *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData struct {
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	TimeStamp   *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData) SetDuration(v int64) *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData {
	s.Duration = &v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData) SetServiceArea(v string) *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData) SetTaskProfile(v string) *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData {
	s.TaskProfile = &v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData) SetTimeStamp(v int64) *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBodyData {
	s.TimeStamp = &v
	return s
}

type DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse) SetHeaders(v map[string]*string) *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse) SetStatusCode(v int32) *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse) SetBody(v *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponseBody) *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse {
	s.Body = v
	return s
}

type DescribeMeterRecordRtUsageByTaskProfileRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterRecordRtUsageByTaskProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRecordRtUsageByTaskProfileRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterRecordRtUsageByTaskProfileRequest) SetAppId(v string) *DescribeMeterRecordRtUsageByTaskProfileRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileRequest) SetEndTs(v int64) *DescribeMeterRecordRtUsageByTaskProfileRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileRequest) SetInterval(v int64) *DescribeMeterRecordRtUsageByTaskProfileRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileRequest) SetServiceArea(v string) *DescribeMeterRecordRtUsageByTaskProfileRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileRequest) SetStartTs(v int64) *DescribeMeterRecordRtUsageByTaskProfileRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterRecordRtUsageByTaskProfileResponseBody struct {
	Data      []*DescribeMeterRecordRtUsageByTaskProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterRecordRtUsageByTaskProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRecordRtUsageByTaskProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponseBody) SetData(v []*DescribeMeterRecordRtUsageByTaskProfileResponseBodyData) *DescribeMeterRecordRtUsageByTaskProfileResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponseBody) SetRequestId(v string) *DescribeMeterRecordRtUsageByTaskProfileResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterRecordRtUsageByTaskProfileResponseBodyData struct {
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Ratio       *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	Timestamp   *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeMeterRecordRtUsageByTaskProfileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRecordRtUsageByTaskProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData) SetDuration(v int64) *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData {
	s.Duration = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData) SetRatio(v string) *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData) SetServiceArea(v string) *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData) SetTaskProfile(v string) *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData {
	s.TaskProfile = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData) SetTimestamp(v int64) *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData) SetType(v string) *DescribeMeterRecordRtUsageByTaskProfileResponseBodyData {
	s.Type = &v
	return s
}

type DescribeMeterRecordRtUsageByTaskProfileResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterRecordRtUsageByTaskProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterRecordRtUsageByTaskProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRecordRtUsageByTaskProfileResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponse) SetHeaders(v map[string]*string) *DescribeMeterRecordRtUsageByTaskProfileResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponse) SetStatusCode(v int32) *DescribeMeterRecordRtUsageByTaskProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterRecordRtUsageByTaskProfileResponse) SetBody(v *DescribeMeterRecordRtUsageByTaskProfileResponseBody) *DescribeMeterRecordRtUsageByTaskProfileResponse {
	s.Body = v
	return s
}

type DescribeMeterRtcApproxPeakRateRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterRtcApproxPeakRateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcApproxPeakRateRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcApproxPeakRateRequest) SetAppId(v string) *DescribeMeterRtcApproxPeakRateRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateRequest) SetEndTs(v int64) *DescribeMeterRtcApproxPeakRateRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateRequest) SetInterval(v int64) *DescribeMeterRtcApproxPeakRateRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateRequest) SetServiceArea(v string) *DescribeMeterRtcApproxPeakRateRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateRequest) SetStartTs(v int64) *DescribeMeterRtcApproxPeakRateRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterRtcApproxPeakRateResponseBody struct {
	ApproxPeakRate *float32                                                    `json:"ApproxPeakRate,omitempty" xml:"ApproxPeakRate,omitempty"`
	ApproxPeakTs   *int64                                                      `json:"ApproxPeakTs,omitempty" xml:"ApproxPeakTs,omitempty"`
	PeakRateVoList []*DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList `json:"PeakRateVoList,omitempty" xml:"PeakRateVoList,omitempty" type:"Repeated"`
	RequestId      *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterRtcApproxPeakRateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcApproxPeakRateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcApproxPeakRateResponseBody) SetApproxPeakRate(v float32) *DescribeMeterRtcApproxPeakRateResponseBody {
	s.ApproxPeakRate = &v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateResponseBody) SetApproxPeakTs(v int64) *DescribeMeterRtcApproxPeakRateResponseBody {
	s.ApproxPeakTs = &v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateResponseBody) SetPeakRateVoList(v []*DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList) *DescribeMeterRtcApproxPeakRateResponseBody {
	s.PeakRateVoList = v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateResponseBody) SetRequestId(v string) *DescribeMeterRtcApproxPeakRateResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList struct {
	PeakRate  *float32 `json:"PeakRate,omitempty" xml:"PeakRate,omitempty"`
	PeakTs    *int64   `json:"PeakTs,omitempty" xml:"PeakTs,omitempty"`
	Timestamp *int64   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList) SetPeakRate(v float32) *DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList {
	s.PeakRate = &v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList) SetPeakTs(v int64) *DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList {
	s.PeakTs = &v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList) SetTimestamp(v int64) *DescribeMeterRtcApproxPeakRateResponseBodyPeakRateVoList {
	s.Timestamp = &v
	return s
}

type DescribeMeterRtcApproxPeakRateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterRtcApproxPeakRateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterRtcApproxPeakRateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcApproxPeakRateResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcApproxPeakRateResponse) SetHeaders(v map[string]*string) *DescribeMeterRtcApproxPeakRateResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateResponse) SetStatusCode(v int32) *DescribeMeterRtcApproxPeakRateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterRtcApproxPeakRateResponse) SetBody(v *DescribeMeterRtcApproxPeakRateResponseBody) *DescribeMeterRtcApproxPeakRateResponse {
	s.Body = v
	return s
}

type DescribeMeterRtcChannelCntDataRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterRtcChannelCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcChannelCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcChannelCntDataRequest) SetAppId(v string) *DescribeMeterRtcChannelCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterRtcChannelCntDataRequest) SetEndTs(v int64) *DescribeMeterRtcChannelCntDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterRtcChannelCntDataRequest) SetInterval(v int64) *DescribeMeterRtcChannelCntDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterRtcChannelCntDataRequest) SetServiceArea(v string) *DescribeMeterRtcChannelCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRtcChannelCntDataRequest) SetStartTs(v int64) *DescribeMeterRtcChannelCntDataRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterRtcChannelCntDataResponseBody struct {
	Data      []*DescribeMeterRtcChannelCntDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterRtcChannelCntDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcChannelCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcChannelCntDataResponseBody) SetData(v []*DescribeMeterRtcChannelCntDataResponseBodyData) *DescribeMeterRtcChannelCntDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterRtcChannelCntDataResponseBody) SetRequestId(v string) *DescribeMeterRtcChannelCntDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterRtcChannelCntDataResponseBodyData struct {
	ChannelCnt *int64  `json:"ChannelCnt,omitempty" xml:"ChannelCnt,omitempty"`
	Timestamp  *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMeterRtcChannelCntDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcChannelCntDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcChannelCntDataResponseBodyData) SetChannelCnt(v int64) *DescribeMeterRtcChannelCntDataResponseBodyData {
	s.ChannelCnt = &v
	return s
}

func (s *DescribeMeterRtcChannelCntDataResponseBodyData) SetTimestamp(v string) *DescribeMeterRtcChannelCntDataResponseBodyData {
	s.Timestamp = &v
	return s
}

type DescribeMeterRtcChannelCntDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterRtcChannelCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterRtcChannelCntDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcChannelCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcChannelCntDataResponse) SetHeaders(v map[string]*string) *DescribeMeterRtcChannelCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterRtcChannelCntDataResponse) SetStatusCode(v int32) *DescribeMeterRtcChannelCntDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterRtcChannelCntDataResponse) SetBody(v *DescribeMeterRtcChannelCntDataResponseBody) *DescribeMeterRtcChannelCntDataResponse {
	s.Body = v
	return s
}

type DescribeMeterRtcDurationRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterRtcDurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcDurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcDurationRequest) SetAppId(v string) *DescribeMeterRtcDurationRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterRtcDurationRequest) SetEndTs(v int64) *DescribeMeterRtcDurationRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterRtcDurationRequest) SetInterval(v int64) *DescribeMeterRtcDurationRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterRtcDurationRequest) SetServiceArea(v string) *DescribeMeterRtcDurationRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRtcDurationRequest) SetStartTs(v int64) *DescribeMeterRtcDurationRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterRtcDurationResponseBody struct {
	Data      []*DescribeMeterRtcDurationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ReadyTs   *int64                                      `json:"ReadyTs,omitempty" xml:"ReadyTs,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterRtcDurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcDurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcDurationResponseBody) SetData(v []*DescribeMeterRtcDurationResponseBodyData) *DescribeMeterRtcDurationResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterRtcDurationResponseBody) SetReadyTs(v int64) *DescribeMeterRtcDurationResponseBody {
	s.ReadyTs = &v
	return s
}

func (s *DescribeMeterRtcDurationResponseBody) SetRequestId(v string) *DescribeMeterRtcDurationResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterRtcDurationResponseBodyData struct {
	AudioDuration   *int64 `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	ContentDuration *int64 `json:"ContentDuration,omitempty" xml:"ContentDuration,omitempty"`
	Timestamp       *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TotalDuration   *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	V1080Duration   *int64 `json:"V1080Duration,omitempty" xml:"V1080Duration,omitempty"`
	V480Duration    *int64 `json:"V480Duration,omitempty" xml:"V480Duration,omitempty"`
	V720Duration    *int64 `json:"V720Duration,omitempty" xml:"V720Duration,omitempty"`
}

func (s DescribeMeterRtcDurationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcDurationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcDurationResponseBodyData) SetAudioDuration(v int64) *DescribeMeterRtcDurationResponseBodyData {
	s.AudioDuration = &v
	return s
}

func (s *DescribeMeterRtcDurationResponseBodyData) SetContentDuration(v int64) *DescribeMeterRtcDurationResponseBodyData {
	s.ContentDuration = &v
	return s
}

func (s *DescribeMeterRtcDurationResponseBodyData) SetTimestamp(v int64) *DescribeMeterRtcDurationResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeMeterRtcDurationResponseBodyData) SetTotalDuration(v int64) *DescribeMeterRtcDurationResponseBodyData {
	s.TotalDuration = &v
	return s
}

func (s *DescribeMeterRtcDurationResponseBodyData) SetV1080Duration(v int64) *DescribeMeterRtcDurationResponseBodyData {
	s.V1080Duration = &v
	return s
}

func (s *DescribeMeterRtcDurationResponseBodyData) SetV480Duration(v int64) *DescribeMeterRtcDurationResponseBodyData {
	s.V480Duration = &v
	return s
}

func (s *DescribeMeterRtcDurationResponseBodyData) SetV720Duration(v int64) *DescribeMeterRtcDurationResponseBodyData {
	s.V720Duration = &v
	return s
}

type DescribeMeterRtcDurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterRtcDurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterRtcDurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcDurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcDurationResponse) SetHeaders(v map[string]*string) *DescribeMeterRtcDurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterRtcDurationResponse) SetStatusCode(v int32) *DescribeMeterRtcDurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterRtcDurationResponse) SetBody(v *DescribeMeterRtcDurationResponseBody) *DescribeMeterRtcDurationResponse {
	s.Body = v
	return s
}

type DescribeMeterRtcPeakChannelCntDataRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterRtcPeakChannelCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcPeakChannelCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcPeakChannelCntDataRequest) SetAppId(v string) *DescribeMeterRtcPeakChannelCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterRtcPeakChannelCntDataRequest) SetEndTs(v int64) *DescribeMeterRtcPeakChannelCntDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterRtcPeakChannelCntDataRequest) SetInterval(v int64) *DescribeMeterRtcPeakChannelCntDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterRtcPeakChannelCntDataRequest) SetServiceArea(v string) *DescribeMeterRtcPeakChannelCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRtcPeakChannelCntDataRequest) SetStartTs(v int64) *DescribeMeterRtcPeakChannelCntDataRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterRtcPeakChannelCntDataResponseBody struct {
	Data      []*DescribeMeterRtcPeakChannelCntDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterRtcPeakChannelCntDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcPeakChannelCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcPeakChannelCntDataResponseBody) SetData(v []*DescribeMeterRtcPeakChannelCntDataResponseBodyData) *DescribeMeterRtcPeakChannelCntDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterRtcPeakChannelCntDataResponseBody) SetRequestId(v string) *DescribeMeterRtcPeakChannelCntDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterRtcPeakChannelCntDataResponseBodyData struct {
	ActiveChannelPeak     *int64 `json:"ActiveChannelPeak,omitempty" xml:"ActiveChannelPeak,omitempty"`
	ActiveChannelPeakTime *int64 `json:"ActiveChannelPeakTime,omitempty" xml:"ActiveChannelPeakTime,omitempty"`
	Timestamp             *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMeterRtcPeakChannelCntDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcPeakChannelCntDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcPeakChannelCntDataResponseBodyData) SetActiveChannelPeak(v int64) *DescribeMeterRtcPeakChannelCntDataResponseBodyData {
	s.ActiveChannelPeak = &v
	return s
}

func (s *DescribeMeterRtcPeakChannelCntDataResponseBodyData) SetActiveChannelPeakTime(v int64) *DescribeMeterRtcPeakChannelCntDataResponseBodyData {
	s.ActiveChannelPeakTime = &v
	return s
}

func (s *DescribeMeterRtcPeakChannelCntDataResponseBodyData) SetTimestamp(v int64) *DescribeMeterRtcPeakChannelCntDataResponseBodyData {
	s.Timestamp = &v
	return s
}

type DescribeMeterRtcPeakChannelCntDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterRtcPeakChannelCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterRtcPeakChannelCntDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcPeakChannelCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcPeakChannelCntDataResponse) SetHeaders(v map[string]*string) *DescribeMeterRtcPeakChannelCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterRtcPeakChannelCntDataResponse) SetStatusCode(v int32) *DescribeMeterRtcPeakChannelCntDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterRtcPeakChannelCntDataResponse) SetBody(v *DescribeMeterRtcPeakChannelCntDataResponseBody) *DescribeMeterRtcPeakChannelCntDataResponse {
	s.Body = v
	return s
}

type DescribeMeterRtcPeakUserCntDataRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterRtcPeakUserCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcPeakUserCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcPeakUserCntDataRequest) SetAppId(v string) *DescribeMeterRtcPeakUserCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterRtcPeakUserCntDataRequest) SetEndTs(v int64) *DescribeMeterRtcPeakUserCntDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterRtcPeakUserCntDataRequest) SetInterval(v int64) *DescribeMeterRtcPeakUserCntDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterRtcPeakUserCntDataRequest) SetServiceArea(v string) *DescribeMeterRtcPeakUserCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRtcPeakUserCntDataRequest) SetStartTs(v int64) *DescribeMeterRtcPeakUserCntDataRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterRtcPeakUserCntDataResponseBody struct {
	Data      []*DescribeMeterRtcPeakUserCntDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterRtcPeakUserCntDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcPeakUserCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcPeakUserCntDataResponseBody) SetData(v []*DescribeMeterRtcPeakUserCntDataResponseBodyData) *DescribeMeterRtcPeakUserCntDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterRtcPeakUserCntDataResponseBody) SetRequestId(v string) *DescribeMeterRtcPeakUserCntDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterRtcPeakUserCntDataResponseBodyData struct {
	ActiveUserPeak     *int64 `json:"ActiveUserPeak,omitempty" xml:"ActiveUserPeak,omitempty"`
	ActiveUserPeakTime *int64 `json:"ActiveUserPeakTime,omitempty" xml:"ActiveUserPeakTime,omitempty"`
	Timestamp          *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMeterRtcPeakUserCntDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcPeakUserCntDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcPeakUserCntDataResponseBodyData) SetActiveUserPeak(v int64) *DescribeMeterRtcPeakUserCntDataResponseBodyData {
	s.ActiveUserPeak = &v
	return s
}

func (s *DescribeMeterRtcPeakUserCntDataResponseBodyData) SetActiveUserPeakTime(v int64) *DescribeMeterRtcPeakUserCntDataResponseBodyData {
	s.ActiveUserPeakTime = &v
	return s
}

func (s *DescribeMeterRtcPeakUserCntDataResponseBodyData) SetTimestamp(v int64) *DescribeMeterRtcPeakUserCntDataResponseBodyData {
	s.Timestamp = &v
	return s
}

type DescribeMeterRtcPeakUserCntDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterRtcPeakUserCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterRtcPeakUserCntDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcPeakUserCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcPeakUserCntDataResponse) SetHeaders(v map[string]*string) *DescribeMeterRtcPeakUserCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterRtcPeakUserCntDataResponse) SetStatusCode(v int32) *DescribeMeterRtcPeakUserCntDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterRtcPeakUserCntDataResponse) SetBody(v *DescribeMeterRtcPeakUserCntDataResponseBody) *DescribeMeterRtcPeakUserCntDataResponse {
	s.Body = v
	return s
}

type DescribeMeterRtcRtBandWidthUsageRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterRtcRtBandWidthUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcRtBandWidthUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcRtBandWidthUsageRequest) SetAppId(v string) *DescribeMeterRtcRtBandWidthUsageRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageRequest) SetEndTs(v int64) *DescribeMeterRtcRtBandWidthUsageRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageRequest) SetInterval(v int64) *DescribeMeterRtcRtBandWidthUsageRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageRequest) SetServiceArea(v string) *DescribeMeterRtcRtBandWidthUsageRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageRequest) SetStartTs(v int64) *DescribeMeterRtcRtBandWidthUsageRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterRtcRtBandWidthUsageResponseBody struct {
	Data      []*DescribeMeterRtcRtBandWidthUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterRtcRtBandWidthUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcRtBandWidthUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcRtBandWidthUsageResponseBody) SetData(v []*DescribeMeterRtcRtBandWidthUsageResponseBodyData) *DescribeMeterRtcRtBandWidthUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageResponseBody) SetRequestId(v string) *DescribeMeterRtcRtBandWidthUsageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterRtcRtBandWidthUsageResponseBodyData struct {
	AnchorPeakRate   *float32 `json:"AnchorPeakRate,omitempty" xml:"AnchorPeakRate,omitempty"`
	AnchorPeakTs     *int64   `json:"AnchorPeakTs,omitempty" xml:"AnchorPeakTs,omitempty"`
	AudiencePeakRate *float32 `json:"AudiencePeakRate,omitempty" xml:"AudiencePeakRate,omitempty"`
	AudiencePeakTs   *int64   `json:"AudiencePeakTs,omitempty" xml:"AudiencePeakTs,omitempty"`
	PeakRate         *float32 `json:"PeakRate,omitempty" xml:"PeakRate,omitempty"`
	PeakTs           *int64   `json:"PeakTs,omitempty" xml:"PeakTs,omitempty"`
	Timestamp        *int64   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMeterRtcRtBandWidthUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcRtBandWidthUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcRtBandWidthUsageResponseBodyData) SetAnchorPeakRate(v float32) *DescribeMeterRtcRtBandWidthUsageResponseBodyData {
	s.AnchorPeakRate = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageResponseBodyData) SetAnchorPeakTs(v int64) *DescribeMeterRtcRtBandWidthUsageResponseBodyData {
	s.AnchorPeakTs = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageResponseBodyData) SetAudiencePeakRate(v float32) *DescribeMeterRtcRtBandWidthUsageResponseBodyData {
	s.AudiencePeakRate = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageResponseBodyData) SetAudiencePeakTs(v int64) *DescribeMeterRtcRtBandWidthUsageResponseBodyData {
	s.AudiencePeakTs = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageResponseBodyData) SetPeakRate(v float32) *DescribeMeterRtcRtBandWidthUsageResponseBodyData {
	s.PeakRate = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageResponseBodyData) SetPeakTs(v int64) *DescribeMeterRtcRtBandWidthUsageResponseBodyData {
	s.PeakTs = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageResponseBodyData) SetTimestamp(v int64) *DescribeMeterRtcRtBandWidthUsageResponseBodyData {
	s.Timestamp = &v
	return s
}

type DescribeMeterRtcRtBandWidthUsageResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterRtcRtBandWidthUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterRtcRtBandWidthUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcRtBandWidthUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcRtBandWidthUsageResponse) SetHeaders(v map[string]*string) *DescribeMeterRtcRtBandWidthUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageResponse) SetStatusCode(v int32) *DescribeMeterRtcRtBandWidthUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterRtcRtBandWidthUsageResponse) SetBody(v *DescribeMeterRtcRtBandWidthUsageResponseBody) *DescribeMeterRtcRtBandWidthUsageResponse {
	s.Body = v
	return s
}

type DescribeMeterRtcRtFlowUsageRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterRtcRtFlowUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcRtFlowUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcRtFlowUsageRequest) SetAppId(v string) *DescribeMeterRtcRtFlowUsageRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageRequest) SetEndTs(v int64) *DescribeMeterRtcRtFlowUsageRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageRequest) SetInterval(v int64) *DescribeMeterRtcRtFlowUsageRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageRequest) SetServiceArea(v string) *DescribeMeterRtcRtFlowUsageRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageRequest) SetStartTs(v int64) *DescribeMeterRtcRtFlowUsageRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterRtcRtFlowUsageResponseBody struct {
	Data      []*DescribeMeterRtcRtFlowUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterRtcRtFlowUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcRtFlowUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcRtFlowUsageResponseBody) SetData(v []*DescribeMeterRtcRtFlowUsageResponseBodyData) *DescribeMeterRtcRtFlowUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageResponseBody) SetRequestId(v string) *DescribeMeterRtcRtFlowUsageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterRtcRtFlowUsageResponseBodyData struct {
	AnchorFlowValue   *float32 `json:"AnchorFlowValue,omitempty" xml:"AnchorFlowValue,omitempty"`
	AudienceFlowValue *float32 `json:"AudienceFlowValue,omitempty" xml:"AudienceFlowValue,omitempty"`
	FlowValue         *float32 `json:"FlowValue,omitempty" xml:"FlowValue,omitempty"`
	Timestamp         *int64   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMeterRtcRtFlowUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcRtFlowUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcRtFlowUsageResponseBodyData) SetAnchorFlowValue(v float32) *DescribeMeterRtcRtFlowUsageResponseBodyData {
	s.AnchorFlowValue = &v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageResponseBodyData) SetAudienceFlowValue(v float32) *DescribeMeterRtcRtFlowUsageResponseBodyData {
	s.AudienceFlowValue = &v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageResponseBodyData) SetFlowValue(v float32) *DescribeMeterRtcRtFlowUsageResponseBodyData {
	s.FlowValue = &v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageResponseBodyData) SetTimestamp(v int64) *DescribeMeterRtcRtFlowUsageResponseBodyData {
	s.Timestamp = &v
	return s
}

type DescribeMeterRtcRtFlowUsageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterRtcRtFlowUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterRtcRtFlowUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcRtFlowUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcRtFlowUsageResponse) SetHeaders(v map[string]*string) *DescribeMeterRtcRtFlowUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageResponse) SetStatusCode(v int32) *DescribeMeterRtcRtFlowUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterRtcRtFlowUsageResponse) SetBody(v *DescribeMeterRtcRtFlowUsageResponseBody) *DescribeMeterRtcRtFlowUsageResponse {
	s.Body = v
	return s
}

type DescribeMeterRtcUserCntDataRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs       *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterRtcUserCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcUserCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcUserCntDataRequest) SetAppId(v string) *DescribeMeterRtcUserCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterRtcUserCntDataRequest) SetEndTs(v int64) *DescribeMeterRtcUserCntDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterRtcUserCntDataRequest) SetInterval(v int64) *DescribeMeterRtcUserCntDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterRtcUserCntDataRequest) SetServiceArea(v string) *DescribeMeterRtcUserCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeMeterRtcUserCntDataRequest) SetStartTs(v int64) *DescribeMeterRtcUserCntDataRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterRtcUserCntDataResponseBody struct {
	Data      []*DescribeMeterRtcUserCntDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterRtcUserCntDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcUserCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcUserCntDataResponseBody) SetData(v []*DescribeMeterRtcUserCntDataResponseBodyData) *DescribeMeterRtcUserCntDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterRtcUserCntDataResponseBody) SetRequestId(v string) *DescribeMeterRtcUserCntDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterRtcUserCntDataResponseBodyData struct {
	ActiveUserCnt *int64 `json:"ActiveUserCnt,omitempty" xml:"ActiveUserCnt,omitempty"`
	Timestamp     *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMeterRtcUserCntDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcUserCntDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcUserCntDataResponseBodyData) SetActiveUserCnt(v int64) *DescribeMeterRtcUserCntDataResponseBodyData {
	s.ActiveUserCnt = &v
	return s
}

func (s *DescribeMeterRtcUserCntDataResponseBodyData) SetTimestamp(v int64) *DescribeMeterRtcUserCntDataResponseBodyData {
	s.Timestamp = &v
	return s
}

type DescribeMeterRtcUserCntDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterRtcUserCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterRtcUserCntDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterRtcUserCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterRtcUserCntDataResponse) SetHeaders(v map[string]*string) *DescribeMeterRtcUserCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterRtcUserCntDataResponse) SetStatusCode(v int32) *DescribeMeterRtcUserCntDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterRtcUserCntDataResponse) SetBody(v *DescribeMeterRtcUserCntDataResponseBody) *DescribeMeterRtcUserCntDataResponse {
	s.Body = v
	return s
}

type DescribeNewPlayVideoPlaySessionEventDetailRequest struct {
	BizDate     *int64  `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	InputStatus *string `json:"InputStatus,omitempty" xml:"InputStatus,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VPS         *string `json:"VPS,omitempty" xml:"VPS,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessionEventDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessionEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailRequest) SetBizDate(v int64) *DescribeNewPlayVideoPlaySessionEventDetailRequest {
	s.BizDate = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailRequest) SetInputStatus(v string) *DescribeNewPlayVideoPlaySessionEventDetailRequest {
	s.InputStatus = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailRequest) SetPageNum(v int32) *DescribeNewPlayVideoPlaySessionEventDetailRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailRequest) SetPageSize(v int32) *DescribeNewPlayVideoPlaySessionEventDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailRequest) SetVPS(v string) *DescribeNewPlayVideoPlaySessionEventDetailRequest {
	s.VPS = &v
	return s
}

type DescribeNewPlayVideoPlaySessionEventDetailResponseBody struct {
	Data      []*DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum   *int64                                                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int64                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReadyTs   *int64                                                        `json:"ReadyTs,omitempty" xml:"ReadyTs,omitempty"`
	RequestId *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum  *int64                                                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessionEventDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessionEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBody) SetData(v []*DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) *DescribeNewPlayVideoPlaySessionEventDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBody) SetPageNum(v int64) *DescribeNewPlayVideoPlaySessionEventDetailResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBody) SetPageSize(v int64) *DescribeNewPlayVideoPlaySessionEventDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBody) SetReadyTs(v int64) *DescribeNewPlayVideoPlaySessionEventDetailResponseBody {
	s.ReadyTs = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBody) SetRequestId(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBody) SetTotalNum(v int64) *DescribeNewPlayVideoPlaySessionEventDetailResponseBody {
	s.TotalNum = &v
	return s
}

type DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData struct {
	BizTime     *string `json:"BizTime,omitempty" xml:"BizTime,omitempty"`
	Cost        *string `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Details     *string `json:"Details,omitempty" xml:"Details,omitempty"`
	EventName   *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	IP          *string `json:"IP,omitempty" xml:"IP,omitempty"`
	ISP         *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IsNormal    *int32  `json:"IsNormal,omitempty" xml:"IsNormal,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Subject     *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetBizTime(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.BizTime = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetCost(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.Cost = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetDetails(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.Details = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetEventName(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.EventName = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetIP(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.IP = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetISP(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.ISP = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetIsNormal(v int32) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.IsNormal = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetNetworkType(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetRegion(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.Region = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData) SetSubject(v string) *DescribeNewPlayVideoPlaySessionEventDetailResponseBodyData {
	s.Subject = &v
	return s
}

type DescribeNewPlayVideoPlaySessionEventDetailResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNewPlayVideoPlaySessionEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessionEventDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessionEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponse) SetHeaders(v map[string]*string) *DescribeNewPlayVideoPlaySessionEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponse) SetStatusCode(v int32) *DescribeNewPlayVideoPlaySessionEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionEventDetailResponse) SetBody(v *DescribeNewPlayVideoPlaySessionEventDetailResponseBody) *DescribeNewPlayVideoPlaySessionEventDetailResponse {
	s.Body = v
	return s
}

type DescribeNewPlayVideoPlaySessionListRequest struct {
	EndTimeStamp   *string `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	InputStatus    *string `json:"InputStatus,omitempty" xml:"InputStatus,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNum        *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTimeStamp *string `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	UniqueId       *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessionListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessionListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessionListRequest) SetEndTimeStamp(v string) *DescribeNewPlayVideoPlaySessionListRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListRequest) SetInputStatus(v string) *DescribeNewPlayVideoPlaySessionListRequest {
	s.InputStatus = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListRequest) SetOrder(v string) *DescribeNewPlayVideoPlaySessionListRequest {
	s.Order = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListRequest) SetPageNum(v int32) *DescribeNewPlayVideoPlaySessionListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListRequest) SetPageSize(v int32) *DescribeNewPlayVideoPlaySessionListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListRequest) SetStartTimeStamp(v string) *DescribeNewPlayVideoPlaySessionListRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListRequest) SetUniqueId(v string) *DescribeNewPlayVideoPlaySessionListRequest {
	s.UniqueId = &v
	return s
}

type DescribeNewPlayVideoPlaySessionListResponseBody struct {
	Data      []*DescribeNewPlayVideoPlaySessionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum   *int64                                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int64                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReadyTs   *int64                                                 `json:"ReadyTs,omitempty" xml:"ReadyTs,omitempty"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum  *int64                                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessionListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBody) SetData(v []*DescribeNewPlayVideoPlaySessionListResponseBodyData) *DescribeNewPlayVideoPlaySessionListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBody) SetPageNum(v int64) *DescribeNewPlayVideoPlaySessionListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBody) SetPageSize(v int64) *DescribeNewPlayVideoPlaySessionListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBody) SetReadyTs(v int64) *DescribeNewPlayVideoPlaySessionListResponseBody {
	s.ReadyTs = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBody) SetRequestId(v string) *DescribeNewPlayVideoPlaySessionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBody) SetTotalNum(v int64) *DescribeNewPlayVideoPlaySessionListResponseBody {
	s.TotalNum = &v
	return s
}

type DescribeNewPlayVideoPlaySessionListResponseBodyData struct {
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TraceId         *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	UUID            *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	VPS             *string `json:"VPS,omitempty" xml:"VPS,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBodyData) SetGmtModifiedTime(v string) *DescribeNewPlayVideoPlaySessionListResponseBodyData {
	s.GmtModifiedTime = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBodyData) SetStatus(v string) *DescribeNewPlayVideoPlaySessionListResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBodyData) SetTraceId(v string) *DescribeNewPlayVideoPlaySessionListResponseBodyData {
	s.TraceId = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBodyData) SetUUID(v string) *DescribeNewPlayVideoPlaySessionListResponseBodyData {
	s.UUID = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponseBodyData) SetVPS(v string) *DescribeNewPlayVideoPlaySessionListResponseBodyData {
	s.VPS = &v
	return s
}

type DescribeNewPlayVideoPlaySessionListResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNewPlayVideoPlaySessionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessionListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessionListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessionListResponse) SetHeaders(v map[string]*string) *DescribeNewPlayVideoPlaySessionListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponse) SetStatusCode(v int32) *DescribeNewPlayVideoPlaySessionListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessionListResponse) SetBody(v *DescribeNewPlayVideoPlaySessionListResponseBody) *DescribeNewPlayVideoPlaySessionListResponse {
	s.Body = v
	return s
}

type DescribeNewPlayVideoPlaySessioninfoRequest struct {
	VPS *string `json:"VPS,omitempty" xml:"VPS,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessioninfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessioninfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessioninfoRequest) SetVPS(v string) *DescribeNewPlayVideoPlaySessioninfoRequest {
	s.VPS = &v
	return s
}

type DescribeNewPlayVideoPlaySessioninfoResponseBody struct {
	Data      []*DescribeNewPlayVideoPlaySessioninfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessioninfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessioninfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBody) SetData(v []*DescribeNewPlayVideoPlaySessioninfoResponseBodyData) *DescribeNewPlayVideoPlaySessioninfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBody) SetRequestId(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNewPlayVideoPlaySessioninfoResponseBodyData struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppVersion   *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	DeviceBrand  *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceModel  *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	OS           *string `json:"OS,omitempty" xml:"OS,omitempty"`
	OV           *string `json:"OV,omitempty" xml:"OV,omitempty"`
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
	VPS          *string `json:"VPS,omitempty" xml:"VPS,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessioninfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessioninfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBodyData) SetAppId(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBodyData) SetAppName(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBodyData) SetAppVersion(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBodyData) SetDeviceBrand(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBodyData {
	s.DeviceBrand = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBodyData) SetDeviceModel(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBodyData {
	s.DeviceModel = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBodyData) SetOS(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBodyData {
	s.OS = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBodyData) SetOV(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBodyData {
	s.OV = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBodyData) SetTerminalType(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBodyData {
	s.TerminalType = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponseBodyData) SetVPS(v string) *DescribeNewPlayVideoPlaySessioninfoResponseBodyData {
	s.VPS = &v
	return s
}

type DescribeNewPlayVideoPlaySessioninfoResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNewPlayVideoPlaySessioninfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNewPlayVideoPlaySessioninfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNewPlayVideoPlaySessioninfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponse) SetHeaders(v map[string]*string) *DescribeNewPlayVideoPlaySessioninfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponse) SetStatusCode(v int32) *DescribeNewPlayVideoPlaySessioninfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNewPlayVideoPlaySessioninfoResponse) SetBody(v *DescribeNewPlayVideoPlaySessioninfoResponseBody) *DescribeNewPlayVideoPlaySessioninfoResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("vdmeter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeHuYaRecordByDomainWithOptions(request *DescribeHuYaRecordByDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeHuYaRecordByDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHuYaRecordByDomain"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHuYaRecordByDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHuYaRecordByDomain(request *DescribeHuYaRecordByDomainRequest) (_result *DescribeHuYaRecordByDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHuYaRecordByDomainResponse{}
	_body, _err := client.DescribeHuYaRecordByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHuYaScreenshotByDomainWithOptions(request *DescribeHuYaScreenshotByDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeHuYaScreenshotByDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHuYaScreenshotByDomain"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHuYaScreenshotByDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHuYaScreenshotByDomain(request *DescribeHuYaScreenshotByDomainRequest) (_result *DescribeHuYaScreenshotByDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHuYaScreenshotByDomainResponse{}
	_body, _err := client.DescribeHuYaScreenshotByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHuYaTranscodeByDomainWithOptions(request *DescribeHuYaTranscodeByDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeHuYaTranscodeByDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHuYaTranscodeByDomain"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHuYaTranscodeByDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHuYaTranscodeByDomain(request *DescribeHuYaTranscodeByDomainRequest) (_result *DescribeHuYaTranscodeByDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHuYaTranscodeByDomainResponse{}
	_body, _err := client.DescribeHuYaTranscodeByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterBypassRtUsageByTaskProfileWithOptions(request *DescribeMeterBypassRtUsageByTaskProfileRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterBypassRtUsageByTaskProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterBypassRtUsageByTaskProfile"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterBypassRtUsageByTaskProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterBypassRtUsageByTaskProfile(request *DescribeMeterBypassRtUsageByTaskProfileRequest) (_result *DescribeMeterBypassRtUsageByTaskProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterBypassRtUsageByTaskProfileResponse{}
	_body, _err := client.DescribeMeterBypassRtUsageByTaskProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterMpuTranscodeRtUsageByTaskProfileWithOptions(request *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterMpuTranscodeRtUsageByTaskProfile"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterMpuTranscodeRtUsageByTaskProfile(request *DescribeMeterMpuTranscodeRtUsageByTaskProfileRequest) (_result *DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterMpuTranscodeRtUsageByTaskProfileResponse{}
	_body, _err := client.DescribeMeterMpuTranscodeRtUsageByTaskProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterRecordRtUsageByTaskProfileWithOptions(request *DescribeMeterRecordRtUsageByTaskProfileRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterRecordRtUsageByTaskProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterRecordRtUsageByTaskProfile"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterRecordRtUsageByTaskProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterRecordRtUsageByTaskProfile(request *DescribeMeterRecordRtUsageByTaskProfileRequest) (_result *DescribeMeterRecordRtUsageByTaskProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterRecordRtUsageByTaskProfileResponse{}
	_body, _err := client.DescribeMeterRecordRtUsageByTaskProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterRtcApproxPeakRateWithOptions(request *DescribeMeterRtcApproxPeakRateRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterRtcApproxPeakRateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterRtcApproxPeakRate"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterRtcApproxPeakRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterRtcApproxPeakRate(request *DescribeMeterRtcApproxPeakRateRequest) (_result *DescribeMeterRtcApproxPeakRateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterRtcApproxPeakRateResponse{}
	_body, _err := client.DescribeMeterRtcApproxPeakRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterRtcChannelCntDataWithOptions(request *DescribeMeterRtcChannelCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterRtcChannelCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterRtcChannelCntData"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterRtcChannelCntDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterRtcChannelCntData(request *DescribeMeterRtcChannelCntDataRequest) (_result *DescribeMeterRtcChannelCntDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterRtcChannelCntDataResponse{}
	_body, _err := client.DescribeMeterRtcChannelCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterRtcDurationWithOptions(request *DescribeMeterRtcDurationRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterRtcDurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterRtcDuration"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterRtcDurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterRtcDuration(request *DescribeMeterRtcDurationRequest) (_result *DescribeMeterRtcDurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterRtcDurationResponse{}
	_body, _err := client.DescribeMeterRtcDurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterRtcPeakChannelCntDataWithOptions(request *DescribeMeterRtcPeakChannelCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterRtcPeakChannelCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterRtcPeakChannelCntData"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterRtcPeakChannelCntDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterRtcPeakChannelCntData(request *DescribeMeterRtcPeakChannelCntDataRequest) (_result *DescribeMeterRtcPeakChannelCntDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterRtcPeakChannelCntDataResponse{}
	_body, _err := client.DescribeMeterRtcPeakChannelCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterRtcPeakUserCntDataWithOptions(request *DescribeMeterRtcPeakUserCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterRtcPeakUserCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterRtcPeakUserCntData"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterRtcPeakUserCntDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterRtcPeakUserCntData(request *DescribeMeterRtcPeakUserCntDataRequest) (_result *DescribeMeterRtcPeakUserCntDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterRtcPeakUserCntDataResponse{}
	_body, _err := client.DescribeMeterRtcPeakUserCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterRtcRtBandWidthUsageWithOptions(request *DescribeMeterRtcRtBandWidthUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterRtcRtBandWidthUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterRtcRtBandWidthUsage"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterRtcRtBandWidthUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterRtcRtBandWidthUsage(request *DescribeMeterRtcRtBandWidthUsageRequest) (_result *DescribeMeterRtcRtBandWidthUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterRtcRtBandWidthUsageResponse{}
	_body, _err := client.DescribeMeterRtcRtBandWidthUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterRtcRtFlowUsageWithOptions(request *DescribeMeterRtcRtFlowUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterRtcRtFlowUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterRtcRtFlowUsage"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterRtcRtFlowUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterRtcRtFlowUsage(request *DescribeMeterRtcRtFlowUsageRequest) (_result *DescribeMeterRtcRtFlowUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterRtcRtFlowUsageResponse{}
	_body, _err := client.DescribeMeterRtcRtFlowUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterRtcUserCntDataWithOptions(request *DescribeMeterRtcUserCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterRtcUserCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterRtcUserCntData"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterRtcUserCntDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterRtcUserCntData(request *DescribeMeterRtcUserCntDataRequest) (_result *DescribeMeterRtcUserCntDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterRtcUserCntDataResponse{}
	_body, _err := client.DescribeMeterRtcUserCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNewPlayVideoPlaySessionEventDetailWithOptions(request *DescribeNewPlayVideoPlaySessionEventDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeNewPlayVideoPlaySessionEventDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizDate)) {
		query["BizDate"] = request.BizDate
	}

	if !tea.BoolValue(util.IsUnset(request.InputStatus)) {
		query["InputStatus"] = request.InputStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.VPS)) {
		query["VPS"] = request.VPS
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNewPlayVideoPlaySessionEventDetail"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNewPlayVideoPlaySessionEventDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNewPlayVideoPlaySessionEventDetail(request *DescribeNewPlayVideoPlaySessionEventDetailRequest) (_result *DescribeNewPlayVideoPlaySessionEventDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNewPlayVideoPlaySessionEventDetailResponse{}
	_body, _err := client.DescribeNewPlayVideoPlaySessionEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNewPlayVideoPlaySessionListWithOptions(request *DescribeNewPlayVideoPlaySessionListRequest, runtime *util.RuntimeOptions) (_result *DescribeNewPlayVideoPlaySessionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimeStamp)) {
		query["EndTimeStamp"] = request.EndTimeStamp
	}

	if !tea.BoolValue(util.IsUnset(request.InputStatus)) {
		query["InputStatus"] = request.InputStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeStamp)) {
		query["StartTimeStamp"] = request.StartTimeStamp
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		query["UniqueId"] = request.UniqueId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNewPlayVideoPlaySessionList"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNewPlayVideoPlaySessionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNewPlayVideoPlaySessionList(request *DescribeNewPlayVideoPlaySessionListRequest) (_result *DescribeNewPlayVideoPlaySessionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNewPlayVideoPlaySessionListResponse{}
	_body, _err := client.DescribeNewPlayVideoPlaySessionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNewPlayVideoPlaySessioninfoWithOptions(request *DescribeNewPlayVideoPlaySessioninfoRequest, runtime *util.RuntimeOptions) (_result *DescribeNewPlayVideoPlaySessioninfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VPS)) {
		query["VPS"] = request.VPS
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNewPlayVideoPlaySessioninfo"),
		Version:     tea.String("2021-04-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNewPlayVideoPlaySessioninfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNewPlayVideoPlaySessioninfo(request *DescribeNewPlayVideoPlaySessioninfoRequest) (_result *DescribeNewPlayVideoPlaySessioninfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNewPlayVideoPlaySessioninfoResponse{}
	_body, _err := client.DescribeNewPlayVideoPlaySessioninfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
