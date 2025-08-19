// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListAppsRequest
	GetAppName() *string
	SetCreateTime(v string) *ListAppsRequest
	GetCreateTime() *string
	SetDescription(v string) *ListAppsRequest
	GetDescription() *string
	SetOrderType(v string) *ListAppsRequest
	GetOrderType() *string
	SetPageNumber(v int32) *ListAppsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListAppsRequest
	GetStatus() *string
	SetTags(v string) *ListAppsRequest
	GetTags() *string
}

type ListAppsRequest struct {
	// example:
	//
	// es-severless-test-app
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// 2023-08-29T02:37:22Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// metrics-logs-online
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// desc
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	Tags   *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListAppsRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAppsRequest) GetDescription() *string {
	return s.Description
}

func (s *ListAppsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListAppsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAppsRequest) GetTags() *string {
	return s.Tags
}

func (s *ListAppsRequest) SetAppName(v string) *ListAppsRequest {
	s.AppName = &v
	return s
}

func (s *ListAppsRequest) SetCreateTime(v string) *ListAppsRequest {
	s.CreateTime = &v
	return s
}

func (s *ListAppsRequest) SetDescription(v string) *ListAppsRequest {
	s.Description = &v
	return s
}

func (s *ListAppsRequest) SetOrderType(v string) *ListAppsRequest {
	s.OrderType = &v
	return s
}

func (s *ListAppsRequest) SetPageNumber(v int32) *ListAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v int32) *ListAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppsRequest) SetStatus(v string) *ListAppsRequest {
	s.Status = &v
	return s
}

func (s *ListAppsRequest) SetTags(v string) *ListAppsRequest {
	s.Tags = &v
	return s
}

func (s *ListAppsRequest) Validate() error {
	return dara.Validate(s)
}
