// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetAppVersionsRequest
	GetAppName() *string
	SetImageCategory(v string) *GetAppVersionsRequest
	GetImageCategory() *string
	SetImageType(v string) *GetAppVersionsRequest
	GetImageType() *string
	SetPageNumber(v int64) *GetAppVersionsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetAppVersionsRequest
	GetPageSize() *int64
}

type GetAppVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// V-Ray
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// Public
	ImageCategory *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetAppVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppVersionsRequest) GoString() string {
	return s.String()
}

func (s *GetAppVersionsRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetAppVersionsRequest) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *GetAppVersionsRequest) GetImageType() *string {
	return s.ImageType
}

func (s *GetAppVersionsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetAppVersionsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAppVersionsRequest) SetAppName(v string) *GetAppVersionsRequest {
	s.AppName = &v
	return s
}

func (s *GetAppVersionsRequest) SetImageCategory(v string) *GetAppVersionsRequest {
	s.ImageCategory = &v
	return s
}

func (s *GetAppVersionsRequest) SetImageType(v string) *GetAppVersionsRequest {
	s.ImageType = &v
	return s
}

func (s *GetAppVersionsRequest) SetPageNumber(v int64) *GetAppVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAppVersionsRequest) SetPageSize(v int64) *GetAppVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *GetAppVersionsRequest) Validate() error {
	return dara.Validate(s)
}
