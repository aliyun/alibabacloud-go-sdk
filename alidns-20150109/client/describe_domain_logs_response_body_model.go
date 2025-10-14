// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogs(v *DescribeDomainLogsResponseBodyDomainLogs) *DescribeDomainLogsResponseBody
	GetDomainLogs() *DescribeDomainLogsResponseBodyDomainLogs
	SetPageNumber(v int64) *DescribeDomainLogsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainLogsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDomainLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDomainLogsResponseBody
	GetTotalCount() *int64
}

type DescribeDomainLogsResponseBody struct {
	// The operation logs.
	DomainLogs *DescribeDomainLogsResponseBodyDomainLogs `json:"DomainLogs,omitempty" xml:"DomainLogs,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 2
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainLogsResponseBody) GetDomainLogs() *DescribeDomainLogsResponseBodyDomainLogs {
	return s.DomainLogs
}

func (s *DescribeDomainLogsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainLogsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainLogsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainLogsResponseBody) SetDomainLogs(v *DescribeDomainLogsResponseBodyDomainLogs) *DescribeDomainLogsResponseBody {
	s.DomainLogs = v
	return s
}

func (s *DescribeDomainLogsResponseBody) SetPageNumber(v int64) *DescribeDomainLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainLogsResponseBody) SetPageSize(v int64) *DescribeDomainLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainLogsResponseBody) SetRequestId(v string) *DescribeDomainLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainLogsResponseBody) SetTotalCount(v int64) *DescribeDomainLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainLogsResponseBody) Validate() error {
	if s.DomainLogs != nil {
		if err := s.DomainLogs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainLogsResponseBodyDomainLogs struct {
	DomainLog []*DescribeDomainLogsResponseBodyDomainLogsDomainLog `json:"DomainLog,omitempty" xml:"DomainLog,omitempty" type:"Repeated"`
}

func (s DescribeDomainLogsResponseBodyDomainLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainLogsResponseBodyDomainLogs) GoString() string {
	return s.String()
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) GetDomainLog() []*DescribeDomainLogsResponseBodyDomainLogsDomainLog {
	return s.DomainLog
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) SetDomainLog(v []*DescribeDomainLogsResponseBodyDomainLogsDomainLog) *DescribeDomainLogsResponseBodyDomainLogs {
	s.DomainLog = v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogs) Validate() error {
	if s.DomainLog != nil {
		for _, item := range s.DomainLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainLogsResponseBodyDomainLogsDomainLog struct {
	// The operation.
	//
	// example:
	//
	// Add
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The time when the operation is performed. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-12T09:23Z
	ActionTime *string `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
	// The time when the operation was performed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 143562300000
	ActionTimestamp *int64 `json:"ActionTimestamp,omitempty" xml:"ActionTimestamp,omitempty"`
	// The IP address of the operator.
	//
	// example:
	//
	// 182.92.253.20
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The domain name.
	//
	// example:
	//
	// abc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The message for the operation.
	//
	// example:
	//
	// To the DNS record list
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the private zone.
	//
	// example:
	//
	// cxfd345sd234
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDomainLogsResponseBodyDomainLogsDomainLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainLogsResponseBodyDomainLogsDomainLog) GoString() string {
	return s.String()
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) GetAction() *string {
	return s.Action
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) GetActionTime() *string {
	return s.ActionTime
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) GetActionTimestamp() *int64 {
	return s.ActionTimestamp
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) GetMessage() *string {
	return s.Message
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) SetAction(v string) *DescribeDomainLogsResponseBodyDomainLogsDomainLog {
	s.Action = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) SetActionTime(v string) *DescribeDomainLogsResponseBodyDomainLogsDomainLog {
	s.ActionTime = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) SetActionTimestamp(v int64) *DescribeDomainLogsResponseBodyDomainLogsDomainLog {
	s.ActionTimestamp = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) SetClientIp(v string) *DescribeDomainLogsResponseBodyDomainLogsDomainLog {
	s.ClientIp = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) SetDomainName(v string) *DescribeDomainLogsResponseBodyDomainLogsDomainLog {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) SetMessage(v string) *DescribeDomainLogsResponseBodyDomainLogsDomainLog {
	s.Message = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) SetZoneId(v string) *DescribeDomainLogsResponseBodyDomainLogsDomainLog {
	s.ZoneId = &v
	return s
}

func (s *DescribeDomainLogsResponseBodyDomainLogsDomainLog) Validate() error {
	return dara.Validate(s)
}
