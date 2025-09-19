// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBuildRiskByKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeImageBuildRiskByKeyRequest
	GetCurrentPage() *int32
	SetImageUuid(v string) *DescribeImageBuildRiskByKeyRequest
	GetImageUuid() *string
	SetLang(v string) *DescribeImageBuildRiskByKeyRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageBuildRiskByKeyRequest
	GetPageSize() *int32
	SetRiskKey(v string) *DescribeImageBuildRiskByKeyRequest
	GetRiskKey() *string
	SetStatus(v int32) *DescribeImageBuildRiskByKeyRequest
	GetStatus() *int32
}

type DescribeImageBuildRiskByKeyRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// a910053dd4710173ecc9e9d8931f****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The key of the risk rule.
	//
	// > You can call the [DescribeImageBuildRiskList](~~DescribeImageBuildRiskList~~) operation to obtain the value of **RiskKey**.
	//
	// example:
	//
	// no_user
	RiskKey *string `json:"RiskKey,omitempty" xml:"RiskKey,omitempty"`
	// The status of the alert event. Valid values:
	//
	// 	- **0**: unhandled.
	//
	// 	- **1**: ignored.
	//
	// 	- **2**: false positive.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImageBuildRiskByKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskByKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskByKeyRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBuildRiskByKeyRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeImageBuildRiskByKeyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageBuildRiskByKeyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBuildRiskByKeyRequest) GetRiskKey() *string {
	return s.RiskKey
}

func (s *DescribeImageBuildRiskByKeyRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageBuildRiskByKeyRequest) SetCurrentPage(v int32) *DescribeImageBuildRiskByKeyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyRequest) SetImageUuid(v string) *DescribeImageBuildRiskByKeyRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyRequest) SetLang(v string) *DescribeImageBuildRiskByKeyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyRequest) SetPageSize(v int32) *DescribeImageBuildRiskByKeyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyRequest) SetRiskKey(v string) *DescribeImageBuildRiskByKeyRequest {
	s.RiskKey = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyRequest) SetStatus(v int32) *DescribeImageBuildRiskByKeyRequest {
	s.Status = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyRequest) Validate() error {
	return dara.Validate(s)
}
