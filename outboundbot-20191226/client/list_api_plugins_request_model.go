// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListApiPluginsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListApiPluginsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApiPluginsRequest
	GetPageSize() *int32
	SetUuidsJson(v string) *ListApiPluginsRequest
	GetUuidsJson() *string
}

type ListApiPluginsRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 024f8cf0-c842-4c01-b74b-c8667e4579c7
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries to display per page
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of UUIDs of API plugins
	//
	// example:
	//
	// ["d17d5bfa-4972-4389-9718-f9602edabe48"]
	UuidsJson *string `json:"UuidsJson,omitempty" xml:"UuidsJson,omitempty"`
}

func (s ListApiPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListApiPluginsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApiPluginsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApiPluginsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiPluginsRequest) GetUuidsJson() *string {
	return s.UuidsJson
}

func (s *ListApiPluginsRequest) SetInstanceId(v string) *ListApiPluginsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApiPluginsRequest) SetPageNumber(v int32) *ListApiPluginsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApiPluginsRequest) SetPageSize(v int32) *ListApiPluginsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApiPluginsRequest) SetUuidsJson(v string) *ListApiPluginsRequest {
	s.UuidsJson = &v
	return s
}

func (s *ListApiPluginsRequest) Validate() error {
	return dara.Validate(s)
}
