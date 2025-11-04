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
	SetPageNo(v int32) *ListAlertsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListAlertsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAlertsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAlertsResponseBody
	GetTotalCount() *int32
}

type ListAlertsResponseBody struct {
	// The alerts.
	Alerts []*ListAlertsResponseBodyAlerts `json:"Alerts,omitempty" xml:"Alerts,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
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

func (s *ListAlertsResponseBody) GetPageNo() *int32 {
	return s.PageNo
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

func (s *ListAlertsResponseBody) SetPageNo(v int32) *ListAlertsResponseBody {
	s.PageNo = &v
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
	// The alert type.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The error code.
	//
	// example:
	//
	// ScheduleError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the alert was received in UTC.
	//
	// example:
	//
	// 2024-07-16T10:03Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the alert was modified in UTC.
	//
	// example:
	//
	// 2024-07-16T10:03Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The error message.
	//
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ARN of the related resource.
	//
	// example:
	//
	// acs:ims:mediaweaver:<regionId>:<userId>:vodSource/mySourceLocation/MySource
	RelatedResourceArns *string `json:"RelatedResourceArns,omitempty" xml:"RelatedResourceArns,omitempty"`
	// The ARN of the resource.
	//
	// example:
	//
	// acs:ims:mediaweaver:<regionId>:<userId>:vodSource/mySourceLocation/MySource
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s ListAlertsResponseBodyAlerts) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsResponseBodyAlerts) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyAlerts) GetCategory() *string {
	return s.Category
}

func (s *ListAlertsResponseBodyAlerts) GetCode() *string {
	return s.Code
}

func (s *ListAlertsResponseBodyAlerts) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAlertsResponseBodyAlerts) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAlertsResponseBodyAlerts) GetMessage() *string {
	return s.Message
}

func (s *ListAlertsResponseBodyAlerts) GetRelatedResourceArns() *string {
	return s.RelatedResourceArns
}

func (s *ListAlertsResponseBodyAlerts) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListAlertsResponseBodyAlerts) SetCategory(v string) *ListAlertsResponseBodyAlerts {
	s.Category = &v
	return s
}

func (s *ListAlertsResponseBodyAlerts) SetCode(v string) *ListAlertsResponseBodyAlerts {
	s.Code = &v
	return s
}

func (s *ListAlertsResponseBodyAlerts) SetGmtCreate(v string) *ListAlertsResponseBodyAlerts {
	s.GmtCreate = &v
	return s
}

func (s *ListAlertsResponseBodyAlerts) SetGmtModified(v string) *ListAlertsResponseBodyAlerts {
	s.GmtModified = &v
	return s
}

func (s *ListAlertsResponseBodyAlerts) SetMessage(v string) *ListAlertsResponseBodyAlerts {
	s.Message = &v
	return s
}

func (s *ListAlertsResponseBodyAlerts) SetRelatedResourceArns(v string) *ListAlertsResponseBodyAlerts {
	s.RelatedResourceArns = &v
	return s
}

func (s *ListAlertsResponseBodyAlerts) SetResourceArn(v string) *ListAlertsResponseBodyAlerts {
	s.ResourceArn = &v
	return s
}

func (s *ListAlertsResponseBodyAlerts) Validate() error {
	return dara.Validate(s)
}
