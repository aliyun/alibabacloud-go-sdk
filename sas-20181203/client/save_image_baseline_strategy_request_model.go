// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveImageBaselineStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineItemList(v string) *SaveImageBaselineStrategyRequest
	GetBaselineItemList() *string
	SetLang(v string) *SaveImageBaselineStrategyRequest
	GetLang() *string
	SetSource(v string) *SaveImageBaselineStrategyRequest
	GetSource() *string
	SetStrategyId(v int64) *SaveImageBaselineStrategyRequest
	GetStrategyId() *int64
	SetStrategyName(v string) *SaveImageBaselineStrategyRequest
	GetStrategyName() *string
}

type SaveImageBaselineStrategyRequest struct {
	// The baseline check items.
	//
	// > You can call the [DescribeImageBaselineStrategy](~~DescribeImageBaselineStrategy~~) operation to query baseline check items.
	//
	// This parameter is required.
	//
	// example:
	//
	// ak_leak
	BaselineItemList *string `json:"BaselineItemList,omitempty" xml:"BaselineItemList,omitempty"`
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
	// The data source. If this parameter is left empty, the baseline check policy for images is queried. Valid values:
	//
	// 	- **default**: the baseline check policy for images
	//
	// 	- **agentless**: agentless detection
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the baseline check policy.
	//
	// > You can call the [DescribeImageBaselineStrategy](~~DescribeImageBaselineStrategy~~) operation to query the IDs of baseline check policies.
	//
	// example:
	//
	// 8639
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the baseline check policy.
	//
	// example:
	//
	// default
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s SaveImageBaselineStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveImageBaselineStrategyRequest) GoString() string {
	return s.String()
}

func (s *SaveImageBaselineStrategyRequest) GetBaselineItemList() *string {
	return s.BaselineItemList
}

func (s *SaveImageBaselineStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveImageBaselineStrategyRequest) GetSource() *string {
	return s.Source
}

func (s *SaveImageBaselineStrategyRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *SaveImageBaselineStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *SaveImageBaselineStrategyRequest) SetBaselineItemList(v string) *SaveImageBaselineStrategyRequest {
	s.BaselineItemList = &v
	return s
}

func (s *SaveImageBaselineStrategyRequest) SetLang(v string) *SaveImageBaselineStrategyRequest {
	s.Lang = &v
	return s
}

func (s *SaveImageBaselineStrategyRequest) SetSource(v string) *SaveImageBaselineStrategyRequest {
	s.Source = &v
	return s
}

func (s *SaveImageBaselineStrategyRequest) SetStrategyId(v int64) *SaveImageBaselineStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *SaveImageBaselineStrategyRequest) SetStrategyName(v string) *SaveImageBaselineStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *SaveImageBaselineStrategyRequest) Validate() error {
	return dara.Validate(s)
}
