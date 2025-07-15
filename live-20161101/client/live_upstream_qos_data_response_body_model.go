// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveUpstreamQosDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*LiveUpstreamQosDataResponseBodyData) *LiveUpstreamQosDataResponseBody
	GetData() []*LiveUpstreamQosDataResponseBodyData
	SetEndTime(v string) *LiveUpstreamQosDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *LiveUpstreamQosDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *LiveUpstreamQosDataResponseBody
	GetStartTime() *string
}

type LiveUpstreamQosDataResponseBody struct {
	Data      []*LiveUpstreamQosDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	EndTime   *string                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime *string                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s LiveUpstreamQosDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LiveUpstreamQosDataResponseBody) GoString() string {
	return s.String()
}

func (s *LiveUpstreamQosDataResponseBody) GetData() []*LiveUpstreamQosDataResponseBodyData {
	return s.Data
}

func (s *LiveUpstreamQosDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *LiveUpstreamQosDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LiveUpstreamQosDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *LiveUpstreamQosDataResponseBody) SetData(v []*LiveUpstreamQosDataResponseBodyData) *LiveUpstreamQosDataResponseBody {
	s.Data = v
	return s
}

func (s *LiveUpstreamQosDataResponseBody) SetEndTime(v string) *LiveUpstreamQosDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBody) SetRequestId(v string) *LiveUpstreamQosDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBody) SetStartTime(v string) *LiveUpstreamQosDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type LiveUpstreamQosDataResponseBodyData struct {
	CdnDomain      *string                                          `json:"CdnDomain,omitempty" xml:"CdnDomain,omitempty"`
	CdnIsp         *string                                          `json:"CdnIsp,omitempty" xml:"CdnIsp,omitempty"`
	CdnOcid        *string                                          `json:"CdnOcid,omitempty" xml:"CdnOcid,omitempty"`
	CdnProvince    *string                                          `json:"CdnProvince,omitempty" xml:"CdnProvince,omitempty"`
	DetailData     []*LiveUpstreamQosDataResponseBodyDataDetailData `json:"DetailData,omitempty" xml:"DetailData,omitempty" type:"Repeated"`
	KwaiSidc       *string                                          `json:"KwaiSidc,omitempty" xml:"KwaiSidc,omitempty"`
	KwaiTsc        *int64                                           `json:"KwaiTsc,omitempty" xml:"KwaiTsc,omitempty"`
	UpstreamDomain *string                                          `json:"UpstreamDomain,omitempty" xml:"UpstreamDomain,omitempty"`
}

func (s LiveUpstreamQosDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LiveUpstreamQosDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *LiveUpstreamQosDataResponseBodyData) GetCdnDomain() *string {
	return s.CdnDomain
}

func (s *LiveUpstreamQosDataResponseBodyData) GetCdnIsp() *string {
	return s.CdnIsp
}

func (s *LiveUpstreamQosDataResponseBodyData) GetCdnOcid() *string {
	return s.CdnOcid
}

func (s *LiveUpstreamQosDataResponseBodyData) GetCdnProvince() *string {
	return s.CdnProvince
}

func (s *LiveUpstreamQosDataResponseBodyData) GetDetailData() []*LiveUpstreamQosDataResponseBodyDataDetailData {
	return s.DetailData
}

func (s *LiveUpstreamQosDataResponseBodyData) GetKwaiSidc() *string {
	return s.KwaiSidc
}

func (s *LiveUpstreamQosDataResponseBodyData) GetKwaiTsc() *int64 {
	return s.KwaiTsc
}

func (s *LiveUpstreamQosDataResponseBodyData) GetUpstreamDomain() *string {
	return s.UpstreamDomain
}

func (s *LiveUpstreamQosDataResponseBodyData) SetCdnDomain(v string) *LiveUpstreamQosDataResponseBodyData {
	s.CdnDomain = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyData) SetCdnIsp(v string) *LiveUpstreamQosDataResponseBodyData {
	s.CdnIsp = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyData) SetCdnOcid(v string) *LiveUpstreamQosDataResponseBodyData {
	s.CdnOcid = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyData) SetCdnProvince(v string) *LiveUpstreamQosDataResponseBodyData {
	s.CdnProvince = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyData) SetDetailData(v []*LiveUpstreamQosDataResponseBodyDataDetailData) *LiveUpstreamQosDataResponseBodyData {
	s.DetailData = v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyData) SetKwaiSidc(v string) *LiveUpstreamQosDataResponseBodyData {
	s.KwaiSidc = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyData) SetKwaiTsc(v int64) *LiveUpstreamQosDataResponseBodyData {
	s.KwaiTsc = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyData) SetUpstreamDomain(v string) *LiveUpstreamQosDataResponseBodyData {
	s.UpstreamDomain = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type LiveUpstreamQosDataResponseBodyDataDetailData struct {
	ConnectDur             *int64  `json:"ConnectDur,omitempty" xml:"ConnectDur,omitempty"`
	ConnectFailedCount     *int64  `json:"ConnectFailedCount,omitempty" xml:"ConnectFailedCount,omitempty"`
	Count                  *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	FirstDataDur           *int64  `json:"FirstDataDur,omitempty" xml:"FirstDataDur,omitempty"`
	FirstDataFailedCount   *int64  `json:"FirstDataFailedCount,omitempty" xml:"FirstDataFailedCount,omitempty"`
	FirstFrameDur          *int64  `json:"FirstFrameDur,omitempty" xml:"FirstFrameDur,omitempty"`
	FirstFrameSuccessCount *int64  `json:"FirstFrameSuccessCount,omitempty" xml:"FirstFrameSuccessCount,omitempty"`
	StatusCode2Xx          *int64  `json:"StatusCode2Xx,omitempty" xml:"StatusCode2Xx,omitempty"`
	StatusCode3Xx          *int64  `json:"StatusCode3Xx,omitempty" xml:"StatusCode3Xx,omitempty"`
	StatusCode4Xx          *int64  `json:"StatusCode4Xx,omitempty" xml:"StatusCode4Xx,omitempty"`
	StatusCode5Xx          *int64  `json:"StatusCode5Xx,omitempty" xml:"StatusCode5Xx,omitempty"`
	Timestamp              *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s LiveUpstreamQosDataResponseBodyDataDetailData) String() string {
	return dara.Prettify(s)
}

func (s LiveUpstreamQosDataResponseBodyDataDetailData) GoString() string {
	return s.String()
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetConnectDur() *int64 {
	return s.ConnectDur
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetConnectFailedCount() *int64 {
	return s.ConnectFailedCount
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetCount() *int64 {
	return s.Count
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetFirstDataDur() *int64 {
	return s.FirstDataDur
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetFirstDataFailedCount() *int64 {
	return s.FirstDataFailedCount
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetFirstFrameDur() *int64 {
	return s.FirstFrameDur
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetFirstFrameSuccessCount() *int64 {
	return s.FirstFrameSuccessCount
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetStatusCode2Xx() *int64 {
	return s.StatusCode2Xx
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetStatusCode3Xx() *int64 {
	return s.StatusCode3Xx
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetStatusCode4Xx() *int64 {
	return s.StatusCode4Xx
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetStatusCode5Xx() *int64 {
	return s.StatusCode5Xx
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) GetTimestamp() *string {
	return s.Timestamp
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetConnectDur(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.ConnectDur = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetConnectFailedCount(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.ConnectFailedCount = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetCount(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.Count = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetFirstDataDur(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.FirstDataDur = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetFirstDataFailedCount(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.FirstDataFailedCount = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetFirstFrameDur(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.FirstFrameDur = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetFirstFrameSuccessCount(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.FirstFrameSuccessCount = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetStatusCode2Xx(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.StatusCode2Xx = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetStatusCode3Xx(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.StatusCode3Xx = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetStatusCode4Xx(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.StatusCode4Xx = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetStatusCode5Xx(v int64) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.StatusCode5Xx = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) SetTimestamp(v string) *LiveUpstreamQosDataResponseBodyDataDetailData {
	s.Timestamp = &v
	return s
}

func (s *LiveUpstreamQosDataResponseBodyDataDetailData) Validate() error {
	return dara.Validate(s)
}
