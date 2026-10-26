// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAppPluginVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppPluginVersionsRequest
	GetPageSize() *int32
	SetPluginId(v string) *ListAppPluginVersionsRequest
	GetPluginId() *string
}

type ListAppPluginVersionsRequest struct {
	// The page number. Valid values: integers greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the gateway plug-in.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
}

func (s ListAppPluginVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAppPluginVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppPluginVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppPluginVersionsRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListAppPluginVersionsRequest) SetPageNumber(v int32) *ListAppPluginVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppPluginVersionsRequest) SetPageSize(v int32) *ListAppPluginVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppPluginVersionsRequest) SetPluginId(v string) *ListAppPluginVersionsRequest {
	s.PluginId = &v
	return s
}

func (s *ListAppPluginVersionsRequest) Validate() error {
	return dara.Validate(s)
}
