// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListDriveRequest
	GetLimit() *int32
	SetMarker(v string) *ListDriveRequest
	GetMarker() *string
	SetOwner(v string) *ListDriveRequest
	GetOwner() *string
	SetOwnerType(v string) *ListDriveRequest
	GetOwnerType() *string
}

type ListDriveRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The owner of the drive. If this parameter is not specified, all drives are returned.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The type of the owner. Valid values:
	//
	// user and group.
	//
	// By default, drives of all owner types are returned.
	//
	// example:
	//
	// user
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
}

func (s ListDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDriveRequest) GoString() string {
	return s.String()
}

func (s *ListDriveRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListDriveRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListDriveRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListDriveRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ListDriveRequest) SetLimit(v int32) *ListDriveRequest {
	s.Limit = &v
	return s
}

func (s *ListDriveRequest) SetMarker(v string) *ListDriveRequest {
	s.Marker = &v
	return s
}

func (s *ListDriveRequest) SetOwner(v string) *ListDriveRequest {
	s.Owner = &v
	return s
}

func (s *ListDriveRequest) SetOwnerType(v string) *ListDriveRequest {
	s.OwnerType = &v
	return s
}

func (s *ListDriveRequest) Validate() error {
	return dara.Validate(s)
}
