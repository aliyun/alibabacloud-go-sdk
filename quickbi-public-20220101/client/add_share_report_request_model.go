// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShareReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthPoint(v int32) *AddShareReportRequest
	GetAuthPoint() *int32
	SetExpireDate(v int64) *AddShareReportRequest
	GetExpireDate() *int64
	SetShareToId(v string) *AddShareReportRequest
	GetShareToId() *string
	SetShareToType(v int32) *AddShareReportRequest
	GetShareToType() *int32
	SetWorksId(v string) *AddShareReportRequest
	GetWorksId() *string
}

type AddShareReportRequest struct {
	// The authorization type. Valid values:
	//
	// - 1: View only.
	//
	// - 3: View and export.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	AuthPoint *int32 `json:"AuthPoint,omitempty" xml:"AuthPoint,omitempty"`
	// The expiration time of the sharing, in millisecond-level UNIX timestamp format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1608202110838
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The ID of the share recipient. The value can be a Quick BI user ID or a user group ID.
	//
	// - ShareToType=0 (user): ShareTo is the user ID.
	//
	// - ShareToType=1 (user group): ShareTo is the user group ID.
	//
	// - ShareToType=2 (organization): ShareTo is the organization ID.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	ShareToId *string `json:"ShareToId,omitempty" xml:"ShareToId,omitempty"`
	// The sharing type. Valid values:
	//
	// - 0: user
	//
	// - 1: user group
	//
	// - 2: organization
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ShareToType *int32 `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
	// The ID of the work to share. Works include data portals, dashboards, workbooks, and ad hoc queries.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s AddShareReportRequest) String() string {
	return dara.Prettify(s)
}

func (s AddShareReportRequest) GoString() string {
	return s.String()
}

func (s *AddShareReportRequest) GetAuthPoint() *int32 {
	return s.AuthPoint
}

func (s *AddShareReportRequest) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *AddShareReportRequest) GetShareToId() *string {
	return s.ShareToId
}

func (s *AddShareReportRequest) GetShareToType() *int32 {
	return s.ShareToType
}

func (s *AddShareReportRequest) GetWorksId() *string {
	return s.WorksId
}

func (s *AddShareReportRequest) SetAuthPoint(v int32) *AddShareReportRequest {
	s.AuthPoint = &v
	return s
}

func (s *AddShareReportRequest) SetExpireDate(v int64) *AddShareReportRequest {
	s.ExpireDate = &v
	return s
}

func (s *AddShareReportRequest) SetShareToId(v string) *AddShareReportRequest {
	s.ShareToId = &v
	return s
}

func (s *AddShareReportRequest) SetShareToType(v int32) *AddShareReportRequest {
	s.ShareToType = &v
	return s
}

func (s *AddShareReportRequest) SetWorksId(v string) *AddShareReportRequest {
	s.WorksId = &v
	return s
}

func (s *AddShareReportRequest) Validate() error {
	return dara.Validate(s)
}
