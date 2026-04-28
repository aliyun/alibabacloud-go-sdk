// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFacegroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *ListFacegroupsRequest
	GetDriveId() *string
	SetLimit(v int32) *ListFacegroupsRequest
	GetLimit() *int32
	SetMarker(v string) *ListFacegroupsRequest
	GetMarker() *string
	SetRemarks(v string) *ListFacegroupsRequest
	GetRemarks() *string
	SetReturnTotalCount(v bool) *ListFacegroupsRequest
	GetReturnTotalCount() *bool
}

type ListFacegroupsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The filter condition that is used to query groups. The value can be up to 128 characters in length. An exact match is used.
	Remarks          *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
	ReturnTotalCount *bool   `json:"return_total_count,omitempty" xml:"return_total_count,omitempty"`
}

func (s ListFacegroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFacegroupsRequest) GoString() string {
	return s.String()
}

func (s *ListFacegroupsRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ListFacegroupsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListFacegroupsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListFacegroupsRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *ListFacegroupsRequest) GetReturnTotalCount() *bool {
	return s.ReturnTotalCount
}

func (s *ListFacegroupsRequest) SetDriveId(v string) *ListFacegroupsRequest {
	s.DriveId = &v
	return s
}

func (s *ListFacegroupsRequest) SetLimit(v int32) *ListFacegroupsRequest {
	s.Limit = &v
	return s
}

func (s *ListFacegroupsRequest) SetMarker(v string) *ListFacegroupsRequest {
	s.Marker = &v
	return s
}

func (s *ListFacegroupsRequest) SetRemarks(v string) *ListFacegroupsRequest {
	s.Remarks = &v
	return s
}

func (s *ListFacegroupsRequest) SetReturnTotalCount(v bool) *ListFacegroupsRequest {
	s.ReturnTotalCount = &v
	return s
}

func (s *ListFacegroupsRequest) Validate() error {
	return dara.Validate(s)
}
