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
	SetTenantCode(v string) *DeleteKeywordRequest
	GetTenantCode() *string
}

type DeleteKeywordRequest struct {
	// The IDs of the keywords to delete.
	//
	// example:
	//
	// [6715465]
	KeywordIdList *string `json:"KeywordIdList,omitempty" xml:"KeywordIdList,omitempty"`
	// The IDs of the keywords to delete.
	//
	// example:
	//
	// [16754493]
	KeywordIds *string `json:"KeywordIds,omitempty" xml:"KeywordIds,omitempty"`
	// The ID of the keyword library.
	//
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The code of the keyword library.
	//
	// - desensitize: masking keyword library.
	//
	// example:
	//
	// desensitize
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
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

func (s *DeleteKeywordRequest) GetTenantCode() *string {
	return s.TenantCode
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

func (s *DeleteKeywordRequest) SetTenantCode(v string) *DeleteKeywordRequest {
	s.TenantCode = &v
	return s
}

func (s *DeleteKeywordRequest) Validate() error {
	return dara.Validate(s)
}
