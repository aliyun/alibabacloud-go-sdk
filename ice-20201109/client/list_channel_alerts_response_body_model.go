// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChannelAlertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListChannelAlertsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListChannelAlertsResponseBody
	GetPageSize() *int32
	SetProgramAlerts(v []*ListChannelAlertsResponseBodyProgramAlerts) *ListChannelAlertsResponseBody
	GetProgramAlerts() []*ListChannelAlertsResponseBodyProgramAlerts
	SetRequestId(v string) *ListChannelAlertsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListChannelAlertsResponseBody
	GetTotalCount() *int32
}

type ListChannelAlertsResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The alerts.
	ProgramAlerts []*ListChannelAlertsResponseBodyProgramAlerts `json:"ProgramAlerts,omitempty" xml:"ProgramAlerts,omitempty" type:"Repeated"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of alerts returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChannelAlertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChannelAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChannelAlertsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChannelAlertsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChannelAlertsResponseBody) GetProgramAlerts() []*ListChannelAlertsResponseBodyProgramAlerts {
	return s.ProgramAlerts
}

func (s *ListChannelAlertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChannelAlertsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListChannelAlertsResponseBody) SetPageNo(v int32) *ListChannelAlertsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChannelAlertsResponseBody) SetPageSize(v int32) *ListChannelAlertsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChannelAlertsResponseBody) SetProgramAlerts(v []*ListChannelAlertsResponseBodyProgramAlerts) *ListChannelAlertsResponseBody {
	s.ProgramAlerts = v
	return s
}

func (s *ListChannelAlertsResponseBody) SetRequestId(v string) *ListChannelAlertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChannelAlertsResponseBody) SetTotalCount(v int32) *ListChannelAlertsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChannelAlertsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListChannelAlertsResponseBodyProgramAlerts struct {
	// The ARN of the program.
	//
	// example:
	//
	// acs:ims:mediaweaver:<regionId>:<userId>:program/myChannel/MyProgram
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The alert type.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The number of alerts.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the alert was last modified in UTC.
	//
	// example:
	//
	// 2024-07-16T10:03Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the program.
	//
	// example:
	//
	// program_name
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
}

func (s ListChannelAlertsResponseBodyProgramAlerts) String() string {
	return dara.Prettify(s)
}

func (s ListChannelAlertsResponseBodyProgramAlerts) GoString() string {
	return s.String()
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) GetArn() *string {
	return s.Arn
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) GetCategory() *string {
	return s.Category
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) GetCount() *int32 {
	return s.Count
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) GetProgramName() *string {
	return s.ProgramName
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) SetArn(v string) *ListChannelAlertsResponseBodyProgramAlerts {
	s.Arn = &v
	return s
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) SetCategory(v string) *ListChannelAlertsResponseBodyProgramAlerts {
	s.Category = &v
	return s
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) SetCount(v int32) *ListChannelAlertsResponseBodyProgramAlerts {
	s.Count = &v
	return s
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) SetGmtModified(v string) *ListChannelAlertsResponseBodyProgramAlerts {
	s.GmtModified = &v
	return s
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) SetProgramName(v string) *ListChannelAlertsResponseBodyProgramAlerts {
	s.ProgramName = &v
	return s
}

func (s *ListChannelAlertsResponseBodyProgramAlerts) Validate() error {
	return dara.Validate(s)
}
