// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisConditionFavoriteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAnalysisConditionFavoriteListRequest
	GetLang() *string
	SetRegId(v string) *DescribeAnalysisConditionFavoriteListRequest
	GetRegId() *string
}

type DescribeAnalysisConditionFavoriteListRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAnalysisConditionFavoriteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisConditionFavoriteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisConditionFavoriteListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAnalysisConditionFavoriteListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAnalysisConditionFavoriteListRequest) SetLang(v string) *DescribeAnalysisConditionFavoriteListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListRequest) SetRegId(v string) *DescribeAnalysisConditionFavoriteListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListRequest) Validate() error {
	return dara.Validate(s)
}
