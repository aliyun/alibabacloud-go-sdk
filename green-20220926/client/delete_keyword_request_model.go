// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeywordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeywordIdList(v string) *DeleteKeywordRequest
	GetKeywordIdList() *string
	SetKeywordIds(v string) *DeleteKeywordRequest
	GetKeywordIds() *string
	SetLibId(v string) *DeleteKeywordRequest
	GetLibId() *string
	SetRegionId(v string) *DeleteKeywordRequest
	GetRegionId() *string
}

type DeleteKeywordRequest struct {
	KeywordIdList *string `json:"KeywordIdList,omitempty" xml:"KeywordIdList,omitempty"`
	// example:
	//
	// [16754493]
	KeywordIds *string `json:"KeywordIds,omitempty" xml:"KeywordIds,omitempty"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteKeywordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeywordRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeywordRequest) GetKeywordIdList() *string {
	return s.KeywordIdList
}

func (s *DeleteKeywordRequest) GetKeywordIds() *string {
	return s.KeywordIds
}

func (s *DeleteKeywordRequest) GetLibId() *string {
	return s.LibId
}

func (s *DeleteKeywordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKeywordRequest) SetKeywordIdList(v string) *DeleteKeywordRequest {
	s.KeywordIdList = &v
	return s
}

func (s *DeleteKeywordRequest) SetKeywordIds(v string) *DeleteKeywordRequest {
	s.KeywordIds = &v
	return s
}

func (s *DeleteKeywordRequest) SetLibId(v string) *DeleteKeywordRequest {
	s.LibId = &v
	return s
}

func (s *DeleteKeywordRequest) SetRegionId(v string) *DeleteKeywordRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKeywordRequest) Validate() error {
	return dara.Validate(s)
}
