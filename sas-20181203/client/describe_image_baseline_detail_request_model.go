// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineItemKey(v string) *DescribeImageBaselineDetailRequest
	GetBaselineItemKey() *string
	SetImageUuid(v string) *DescribeImageBaselineDetailRequest
	GetImageUuid() *string
	SetLang(v string) *DescribeImageBaselineDetailRequest
	GetLang() *string
}

type DescribeImageBaselineDetailRequest struct {
	// The information about the baseline.
	//
	// example:
	//
	// Valid values include but are not limited to ak_leak, duplicate_uid, duplicate_pwd_hash, and non_pwd_user.
	BaselineItemKey *string `json:"BaselineItemKey,omitempty" xml:"BaselineItemKey,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// 06293273b67d19516cfcc712194f****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeImageBaselineDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineDetailRequest) GetBaselineItemKey() *string {
	return s.BaselineItemKey
}

func (s *DescribeImageBaselineDetailRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeImageBaselineDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageBaselineDetailRequest) SetBaselineItemKey(v string) *DescribeImageBaselineDetailRequest {
	s.BaselineItemKey = &v
	return s
}

func (s *DescribeImageBaselineDetailRequest) SetImageUuid(v string) *DescribeImageBaselineDetailRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeImageBaselineDetailRequest) SetLang(v string) *DescribeImageBaselineDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageBaselineDetailRequest) Validate() error {
	return dara.Validate(s)
}
