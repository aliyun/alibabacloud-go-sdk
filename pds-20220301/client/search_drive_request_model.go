// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveName(v string) *SearchDriveRequest
	GetDriveName() *string
	SetLimit(v int32) *SearchDriveRequest
	GetLimit() *int32
	SetMarker(v string) *SearchDriveRequest
	GetMarker() *string
	SetOwner(v string) *SearchDriveRequest
	GetOwner() *string
	SetOwnerType(v string) *SearchDriveRequest
	GetOwnerType() *string
}

type SearchDriveRequest struct {
	// The drive name.
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// The maximum number of asynchronous tasks to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The owner of the drive.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The type of the owner. Valid values:
	//
	// user group
	//
	// example:
	//
	// user
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
}

func (s SearchDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDriveRequest) GoString() string {
	return s.String()
}

func (s *SearchDriveRequest) GetDriveName() *string {
	return s.DriveName
}

func (s *SearchDriveRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchDriveRequest) GetMarker() *string {
	return s.Marker
}

func (s *SearchDriveRequest) GetOwner() *string {
	return s.Owner
}

func (s *SearchDriveRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *SearchDriveRequest) SetDriveName(v string) *SearchDriveRequest {
	s.DriveName = &v
	return s
}

func (s *SearchDriveRequest) SetLimit(v int32) *SearchDriveRequest {
	s.Limit = &v
	return s
}

func (s *SearchDriveRequest) SetMarker(v string) *SearchDriveRequest {
	s.Marker = &v
	return s
}

func (s *SearchDriveRequest) SetOwner(v string) *SearchDriveRequest {
	s.Owner = &v
	return s
}

func (s *SearchDriveRequest) SetOwnerType(v string) *SearchDriveRequest {
	s.OwnerType = &v
	return s
}

func (s *SearchDriveRequest) Validate() error {
	return dara.Validate(s)
}
