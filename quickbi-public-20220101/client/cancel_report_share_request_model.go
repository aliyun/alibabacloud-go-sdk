// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelReportShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *CancelReportShareRequest
	GetReportId() *string
	SetShareToIds(v string) *CancelReportShareRequest
	GetShareToIds() *string
	SetShareToType(v int32) *CancelReportShareRequest
	GetShareToType() *int32
}

type CancelReportShareRequest struct {
	// The ID of the work. The works here include BI portal, dashboards, spreadsheets, and self-service access.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the person to be shared, which may be the user ID of the Quick BI or the user group ID.
	//
	// 	- If ShareToType is 0 (user), ShareTo is the user ID.
	//
	// 	- When ShareToType is set to 1 (user group), ShareTo is the user group ID.
	//
	// 	- When ShareToType=2 (organization), ShareTo is the ID of the organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	ShareToIds *string `json:"ShareToIds,omitempty" xml:"ShareToIds,omitempty"`
	// The deletion method. Valid values:
	//
	// 	- 0: Delete by user
	//
	// 	- 1: Delete by user group
	//
	// 	- 2: Delete by organization
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ShareToType *int32 `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
}

func (s CancelReportShareRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelReportShareRequest) GoString() string {
	return s.String()
}

func (s *CancelReportShareRequest) GetReportId() *string {
	return s.ReportId
}

func (s *CancelReportShareRequest) GetShareToIds() *string {
	return s.ShareToIds
}

func (s *CancelReportShareRequest) GetShareToType() *int32 {
	return s.ShareToType
}

func (s *CancelReportShareRequest) SetReportId(v string) *CancelReportShareRequest {
	s.ReportId = &v
	return s
}

func (s *CancelReportShareRequest) SetShareToIds(v string) *CancelReportShareRequest {
	s.ShareToIds = &v
	return s
}

func (s *CancelReportShareRequest) SetShareToType(v int32) *CancelReportShareRequest {
	s.ShareToType = &v
	return s
}

func (s *CancelReportShareRequest) Validate() error {
	return dara.Validate(s)
}
