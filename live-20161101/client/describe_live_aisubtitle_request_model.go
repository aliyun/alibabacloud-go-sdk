// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAISubtitleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsDefault(v bool) *DescribeLiveAISubtitleRequest
	GetIsDefault() *bool
	SetOwnerId(v int64) *DescribeLiveAISubtitleRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *DescribeLiveAISubtitleRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeLiveAISubtitleRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeLiveAISubtitleRequest
	GetRegionId() *string
	SetSubtitleId(v string) *DescribeLiveAISubtitleRequest
	GetSubtitleId() *string
	SetSubtitleName(v string) *DescribeLiveAISubtitleRequest
	GetSubtitleName() *string
}

type DescribeLiveAISubtitleRequest struct {
	// Specifies whether to query the default caption templates. Valid values:
	//
	// - true: queries the default caption templates.
	//
	// - false: does not query the default caption templates.
	//
	// > Default templates are built-in parameter sets. You can use them by specifying the copyFrom parameter in the AddLiveAISubtitle operation.
	//
	// example:
	//
	// false
	IsDefault *bool  `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	OwnerId   *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of templates on each page. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the caption template.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	SubtitleId *string `json:"SubtitleId,omitempty" xml:"SubtitleId,omitempty"`
	// The name of the caption template. The name can contain only digits, uppercase letters, lowercase letters, and hyphens (-). The name cannot start with a hyphen.
	//
	// example:
	//
	// sub01
	SubtitleName *string `json:"SubtitleName,omitempty" xml:"SubtitleName,omitempty"`
}

func (s DescribeLiveAISubtitleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAISubtitleRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveAISubtitleRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeLiveAISubtitleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveAISubtitleRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeLiveAISubtitleRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeLiveAISubtitleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveAISubtitleRequest) GetSubtitleId() *string {
	return s.SubtitleId
}

func (s *DescribeLiveAISubtitleRequest) GetSubtitleName() *string {
	return s.SubtitleName
}

func (s *DescribeLiveAISubtitleRequest) SetIsDefault(v bool) *DescribeLiveAISubtitleRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeLiveAISubtitleRequest) SetOwnerId(v int64) *DescribeLiveAISubtitleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveAISubtitleRequest) SetPageNumber(v string) *DescribeLiveAISubtitleRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveAISubtitleRequest) SetPageSize(v string) *DescribeLiveAISubtitleRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveAISubtitleRequest) SetRegionId(v string) *DescribeLiveAISubtitleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveAISubtitleRequest) SetSubtitleId(v string) *DescribeLiveAISubtitleRequest {
	s.SubtitleId = &v
	return s
}

func (s *DescribeLiveAISubtitleRequest) SetSubtitleName(v string) *DescribeLiveAISubtitleRequest {
	s.SubtitleName = &v
	return s
}

func (s *DescribeLiveAISubtitleRequest) Validate() error {
	return dara.Validate(s)
}
