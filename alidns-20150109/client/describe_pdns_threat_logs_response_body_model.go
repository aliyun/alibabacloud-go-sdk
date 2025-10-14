// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsThreatLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*DescribePdnsThreatLogsResponseBodyLogs) *DescribePdnsThreatLogsResponseBody
	GetLogs() []*DescribePdnsThreatLogsResponseBodyLogs
	SetPageNumber(v int64) *DescribePdnsThreatLogsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsThreatLogsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePdnsThreatLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePdnsThreatLogsResponseBody
	GetTotalCount() *int64
}

type DescribePdnsThreatLogsResponseBody struct {
	Logs       []*DescribePdnsThreatLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	PageNumber *int64                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePdnsThreatLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatLogsResponseBody) GetLogs() []*DescribePdnsThreatLogsResponseBodyLogs {
	return s.Logs
}

func (s *DescribePdnsThreatLogsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsThreatLogsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsThreatLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsThreatLogsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsThreatLogsResponseBody) SetLogs(v []*DescribePdnsThreatLogsResponseBodyLogs) *DescribePdnsThreatLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribePdnsThreatLogsResponseBody) SetPageNumber(v int64) *DescribePdnsThreatLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsThreatLogsResponseBody) SetPageSize(v int64) *DescribePdnsThreatLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsThreatLogsResponseBody) SetRequestId(v string) *DescribePdnsThreatLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsThreatLogsResponseBody) SetTotalCount(v int64) *DescribePdnsThreatLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsThreatLogsResponseBody) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePdnsThreatLogsResponseBodyLogs struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SubDomain   *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatTime  *string `json:"ThreatTime,omitempty" xml:"ThreatTime,omitempty"`
	ThreatType  *string `json:"ThreatType,omitempty" xml:"ThreatType,omitempty"`
}

func (s DescribePdnsThreatLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) GetThreatTime() *string {
	return s.ThreatTime
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) GetThreatType() *string {
	return s.ThreatType
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) SetSourceIp(v string) *DescribePdnsThreatLogsResponseBodyLogs {
	s.SourceIp = &v
	return s
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) SetSubDomain(v string) *DescribePdnsThreatLogsResponseBodyLogs {
	s.SubDomain = &v
	return s
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) SetThreatLevel(v string) *DescribePdnsThreatLogsResponseBodyLogs {
	s.ThreatLevel = &v
	return s
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) SetThreatTime(v string) *DescribePdnsThreatLogsResponseBodyLogs {
	s.ThreatTime = &v
	return s
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) SetThreatType(v string) *DescribePdnsThreatLogsResponseBodyLogs {
	s.ThreatType = &v
	return s
}

func (s *DescribePdnsThreatLogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
