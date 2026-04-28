// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyGroupDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveName(v string) *ListMyGroupDriveRequest
	GetDriveName() *string
	SetLimit(v int32) *ListMyGroupDriveRequest
	GetLimit() *int32
	SetMarker(v string) *ListMyGroupDriveRequest
	GetMarker() *string
}

type ListMyGroupDriveRequest struct {
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
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
}

func (s ListMyGroupDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyGroupDriveRequest) GoString() string {
	return s.String()
}

func (s *ListMyGroupDriveRequest) GetDriveName() *string {
	return s.DriveName
}

func (s *ListMyGroupDriveRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListMyGroupDriveRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListMyGroupDriveRequest) SetDriveName(v string) *ListMyGroupDriveRequest {
	s.DriveName = &v
	return s
}

func (s *ListMyGroupDriveRequest) SetLimit(v int32) *ListMyGroupDriveRequest {
	s.Limit = &v
	return s
}

func (s *ListMyGroupDriveRequest) SetMarker(v string) *ListMyGroupDriveRequest {
	s.Marker = &v
	return s
}

func (s *ListMyGroupDriveRequest) Validate() error {
	return dara.Validate(s)
}
