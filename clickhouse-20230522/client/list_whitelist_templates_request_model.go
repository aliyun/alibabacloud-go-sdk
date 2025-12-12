// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWhitelistTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListWhitelistTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWhitelistTemplatesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListWhitelistTemplatesRequest
	GetRegionId() *string
	SetTemplateName(v string) *ListWhitelistTemplatesRequest
	GetTemplateName() *string
}

type ListWhitelistTemplatesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 25
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// default
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListWhitelistTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWhitelistTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListWhitelistTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWhitelistTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWhitelistTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWhitelistTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListWhitelistTemplatesRequest) SetPageNumber(v int32) *ListWhitelistTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWhitelistTemplatesRequest) SetPageSize(v int32) *ListWhitelistTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWhitelistTemplatesRequest) SetRegionId(v string) *ListWhitelistTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListWhitelistTemplatesRequest) SetTemplateName(v string) *ListWhitelistTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListWhitelistTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
