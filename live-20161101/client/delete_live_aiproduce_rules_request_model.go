// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAIProduceRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DeleteLiveAIProduceRulesRequest
	GetApp() *string
	SetDomain(v string) *DeleteLiveAIProduceRulesRequest
	GetDomain() *string
	SetOwnerId(v int64) *DeleteLiveAIProduceRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveAIProduceRulesRequest
	GetRegionId() *string
	SetRulesId(v string) *DeleteLiveAIProduceRulesRequest
	GetRulesId() *string
	SetSuffixName(v string) *DeleteLiveAIProduceRulesRequest
	GetSuffixName() *string
}

type DeleteLiveAIProduceRulesRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppName
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the subtitle rule.
	//
	// example:
	//
	// 445409ec-7eaa-461d -8f29-4bec2eb9****
	RulesId *string `json:"RulesId,omitempty" xml:"RulesId,omitempty"`
	// The suffix of the subtitle rule.
	//
	// >  Set the value to the name of the subtitle template.
	//
	// example:
	//
	// et
	SuffixName *string `json:"SuffixName,omitempty" xml:"SuffixName,omitempty"`
}

func (s DeleteLiveAIProduceRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAIProduceRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAIProduceRulesRequest) GetApp() *string {
	return s.App
}

func (s *DeleteLiveAIProduceRulesRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteLiveAIProduceRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveAIProduceRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveAIProduceRulesRequest) GetRulesId() *string {
	return s.RulesId
}

func (s *DeleteLiveAIProduceRulesRequest) GetSuffixName() *string {
	return s.SuffixName
}

func (s *DeleteLiveAIProduceRulesRequest) SetApp(v string) *DeleteLiveAIProduceRulesRequest {
	s.App = &v
	return s
}

func (s *DeleteLiveAIProduceRulesRequest) SetDomain(v string) *DeleteLiveAIProduceRulesRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLiveAIProduceRulesRequest) SetOwnerId(v int64) *DeleteLiveAIProduceRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveAIProduceRulesRequest) SetRegionId(v string) *DeleteLiveAIProduceRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveAIProduceRulesRequest) SetRulesId(v string) *DeleteLiveAIProduceRulesRequest {
	s.RulesId = &v
	return s
}

func (s *DeleteLiveAIProduceRulesRequest) SetSuffixName(v string) *DeleteLiveAIProduceRulesRequest {
	s.SuffixName = &v
	return s
}

func (s *DeleteLiveAIProduceRulesRequest) Validate() error {
	return dara.Validate(s)
}
