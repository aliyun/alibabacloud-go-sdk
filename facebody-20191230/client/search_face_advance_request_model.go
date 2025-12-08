// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSearchFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *SearchFaceAdvanceRequest
	GetDbName() *string
	SetDbNames(v string) *SearchFaceAdvanceRequest
	GetDbNames() *string
	SetImageUrlObject(v io.Reader) *SearchFaceAdvanceRequest
	GetImageUrlObject() io.Reader
	SetLimit(v int32) *SearchFaceAdvanceRequest
	GetLimit() *int32
	SetMaxFaceNum(v int64) *SearchFaceAdvanceRequest
	GetMaxFaceNum() *int64
	SetQualityScoreThreshold(v float32) *SearchFaceAdvanceRequest
	GetQualityScoreThreshold() *float32
}

type SearchFaceAdvanceRequest struct {
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
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
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

func (s SearchFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SearchFaceAdvanceRequest) GetDbName() *string {
	return s.DbName
}

func (s *SearchFaceAdvanceRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *SearchFaceAdvanceRequest) GetImageUrlObject() io.Reader {
	return s.ImageUrlObject
}

func (s *SearchFaceAdvanceRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchFaceAdvanceRequest) GetMaxFaceNum() *int64 {
	return s.MaxFaceNum
}

func (s *SearchFaceAdvanceRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *SearchFaceAdvanceRequest) SetDbName(v string) *SearchFaceAdvanceRequest {
	s.DbName = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetDbNames(v string) *SearchFaceAdvanceRequest {
	s.DbNames = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetImageUrlObject(v io.Reader) *SearchFaceAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *SearchFaceAdvanceRequest) SetLimit(v int32) *SearchFaceAdvanceRequest {
	s.Limit = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetMaxFaceNum(v int64) *SearchFaceAdvanceRequest {
	s.MaxFaceNum = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetQualityScoreThreshold(v float32) *SearchFaceAdvanceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *SearchFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
