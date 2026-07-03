// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlerts(v []*ListAlertsResponseBodyAlerts) *ListAlertsResponseBody
	GetAlerts() []*ListAlertsResponseBodyAlerts
	SetMaxResults(v int32) *ListAlertsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAlertsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListAlertsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAlertsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAlertsResponseBody
	GetTotalCount() *int32
}

type ListAlertsResponseBody struct {
	// The alert details.
	Alerts []*ListAlertsResponseBodyAlerts `json:"Alerts,omitempty" xml:"Alerts,omitempty" type:"Repeated"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBody) GetAlerts() []*ListAlertsResponseBodyAlerts {
	return s.Alerts
}

func (s *ListAlertsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlertsResponseBody) SetAlerts(v []*ListAlertsResponseBodyAlerts) *ListAlertsResponseBody {
	s.Alerts = v
	return s
}

func (s *ListAlertsResponseBody) SetMaxResults(v int32) *ListAlertsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlertsResponseBody) SetNextToken(v string) *ListAlertsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlertsResponseBody) SetPageNumber(v int32) *ListAlertsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlertsResponseBody) SetPageSize(v int32) *ListAlertsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlertsResponseBody) SetRequestId(v string) *ListAlertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertsResponseBody) SetTotalCount(v int32) *ListAlertsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlertsResponseBody) Validate() error {
	if s.Alerts != nil {
		for _, item := range s.Alerts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertsResponseBodyAlerts struct {
	// The alert details.
	//
	// example:
	//
	// {"alert_uuid":"a3f8c2e1-9b7d-4f6a-8c2e-1d5b9a7f****"}
	AlertRecord *string `json:"AlertRecord,omitempty" xml:"AlertRecord,omitempty"`
	// The alert UUID.
	//
	// example:
	//
	// 798341271677187
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
}

func (s ListAlertsResponseBodyAlerts) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsResponseBodyAlerts) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyAlerts) GetAlertRecord() *string {
	return s.AlertRecord
}

func (s *ListAlertsResponseBodyAlerts) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *ListAlertsResponseBodyAlerts) SetAlertRecord(v string) *ListAlertsResponseBodyAlerts {
	s.AlertRecord = &v
	return s
}

func (s *ListAlertsResponseBodyAlerts) SetAlertUuid(v string) *ListAlertsResponseBodyAlerts {
	s.AlertUuid = &v
	return s
}

func (s *ListAlertsResponseBodyAlerts) Validate() error {
	return dara.Validate(s)
}
