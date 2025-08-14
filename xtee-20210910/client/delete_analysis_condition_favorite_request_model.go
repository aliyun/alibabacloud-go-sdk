// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAnalysisConditionFavoriteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteAnalysisConditionFavoriteRequest
	GetLang() *string
	SetId(v int64) *DeleteAnalysisConditionFavoriteRequest
	GetId() *int64
	SetRegId(v string) *DeleteAnalysisConditionFavoriteRequest
	GetRegId() *string
}

type DeleteAnalysisConditionFavoriteRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DeleteAnalysisConditionFavoriteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAnalysisConditionFavoriteRequest) GoString() string {
	return s.String()
}

func (s *DeleteAnalysisConditionFavoriteRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteAnalysisConditionFavoriteRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteAnalysisConditionFavoriteRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteAnalysisConditionFavoriteRequest) SetLang(v string) *DeleteAnalysisConditionFavoriteRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAnalysisConditionFavoriteRequest) SetId(v int64) *DeleteAnalysisConditionFavoriteRequest {
	s.Id = &v
	return s
}

func (s *DeleteAnalysisConditionFavoriteRequest) SetRegId(v string) *DeleteAnalysisConditionFavoriteRequest {
	s.RegId = &v
	return s
}

func (s *DeleteAnalysisConditionFavoriteRequest) Validate() error {
	return dara.Validate(s)
}
