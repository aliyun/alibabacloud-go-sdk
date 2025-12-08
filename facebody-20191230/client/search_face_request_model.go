// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *SearchFaceRequest
	GetDbName() *string
	SetDbNames(v string) *SearchFaceRequest
	GetDbNames() *string
	SetImageUrl(v string) *SearchFaceRequest
	GetImageUrl() *string
	SetLimit(v int32) *SearchFaceRequest
	GetLimit() *int32
	SetMaxFaceNum(v int64) *SearchFaceRequest
	GetMaxFaceNum() *int64
	SetQualityScoreThreshold(v float32) *SearchFaceRequest
	GetQualityScoreThreshold() *float32
}

type SearchFaceRequest struct {
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// test1,test2,test3
	DbNames *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/SearchFace/SearchFace1.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// 5
	MaxFaceNum *int64 `json:"MaxFaceNum,omitempty" xml:"MaxFaceNum,omitempty"`
	// example:
	//
	// 50.0
	QualityScoreThreshold *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
}

func (s SearchFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchFaceRequest) GoString() string {
	return s.String()
}

func (s *SearchFaceRequest) GetDbName() *string {
	return s.DbName
}

func (s *SearchFaceRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *SearchFaceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SearchFaceRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchFaceRequest) GetMaxFaceNum() *int64 {
	return s.MaxFaceNum
}

func (s *SearchFaceRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *SearchFaceRequest) SetDbName(v string) *SearchFaceRequest {
	s.DbName = &v
	return s
}

func (s *SearchFaceRequest) SetDbNames(v string) *SearchFaceRequest {
	s.DbNames = &v
	return s
}

func (s *SearchFaceRequest) SetImageUrl(v string) *SearchFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *SearchFaceRequest) SetLimit(v int32) *SearchFaceRequest {
	s.Limit = &v
	return s
}

func (s *SearchFaceRequest) SetMaxFaceNum(v int64) *SearchFaceRequest {
	s.MaxFaceNum = &v
	return s
}

func (s *SearchFaceRequest) SetQualityScoreThreshold(v float32) *SearchFaceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *SearchFaceRequest) Validate() error {
	return dara.Validate(s)
}
