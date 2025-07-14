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
	// The scope of authorization. Valid values:
	//
	// 	- 1: view only
	//
	// 	- 3: View and export
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	AuthPoint *int32 `json:"AuthPoint,omitempty" xml:"AuthPoint,omitempty"`
	// The validity period of the share. The value is a timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1608202110838
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The ID of the person to be shared, which may be the user ID of the Quick BI or the user group ID.
	//
	// 	- If ShareToType is 0 (user), ShareTo is the user ID.
	//
	// 	- When ShareToType is set to 1 (user group), ShareTo is the user group ID.
	//
	// 	- When ShareToType=2 (organization), ShareTo is the ID of the organization.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	ShareToId *string `json:"ShareToId,omitempty" xml:"ShareToId,omitempty"`
	// The share type of the template. Valid values:
	//
	// 	- 0: user
	//
	// 	- 1: user group
	//
	// 	- 2: organization
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ShareToType *int32 `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
	// The ID of the shared work. The works here include BI portal, dashboards, spreadsheets, and self-service access.
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
