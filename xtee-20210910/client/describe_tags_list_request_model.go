// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTagsListRequest
	GetLang() *string
	SetRegId(v string) *DescribeTagsListRequest
	GetRegId() *string
}

type DescribeTagsListRequest struct {
	// Sets the language type for requests and responses, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeTagsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTagsListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTagsListRequest) SetLang(v string) *DescribeTagsListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagsListRequest) SetRegId(v string) *DescribeTagsListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTagsListRequest) Validate() error {
	return dara.Validate(s)
}
