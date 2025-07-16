// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTitleIntelligenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatLevelThreeId(v int64) *GetTitleIntelligenceRequest
	GetCatLevelThreeId() *int64
	SetCatLevelTwoId(v int64) *GetTitleIntelligenceRequest
	GetCatLevelTwoId() *int64
	SetExtra(v string) *GetTitleIntelligenceRequest
	GetExtra() *string
	SetKeywords(v string) *GetTitleIntelligenceRequest
	GetKeywords() *string
	SetPlatform(v string) *GetTitleIntelligenceRequest
	GetPlatform() *string
}

type GetTitleIntelligenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111
	CatLevelThreeId *int64 `json:"CatLevelThreeId,omitempty" xml:"CatLevelThreeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 222
	CatLevelTwoId *int64 `json:"CatLevelTwoId,omitempty" xml:"CatLevelTwoId,omitempty"`
	// example:
	//
	// {"product_id":"1212"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// hello,apple
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ae
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s GetTitleIntelligenceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTitleIntelligenceRequest) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceRequest) GetCatLevelThreeId() *int64 {
	return s.CatLevelThreeId
}

func (s *GetTitleIntelligenceRequest) GetCatLevelTwoId() *int64 {
	return s.CatLevelTwoId
}

func (s *GetTitleIntelligenceRequest) GetExtra() *string {
	return s.Extra
}

func (s *GetTitleIntelligenceRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *GetTitleIntelligenceRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GetTitleIntelligenceRequest) SetCatLevelThreeId(v int64) *GetTitleIntelligenceRequest {
	s.CatLevelThreeId = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetCatLevelTwoId(v int64) *GetTitleIntelligenceRequest {
	s.CatLevelTwoId = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetExtra(v string) *GetTitleIntelligenceRequest {
	s.Extra = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetKeywords(v string) *GetTitleIntelligenceRequest {
	s.Keywords = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetPlatform(v string) *GetTitleIntelligenceRequest {
	s.Platform = &v
	return s
}

func (s *GetTitleIntelligenceRequest) Validate() error {
	return dara.Validate(s)
}
