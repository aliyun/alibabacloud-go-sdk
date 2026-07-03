// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResponseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteResponseRuleRequest
	GetLang() *string
	SetMaxResults(v int32) *DeleteResponseRuleRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DeleteResponseRuleRequest
	GetNextToken() *string
	SetRegionId(v string) *DeleteResponseRuleRequest
	GetRegionId() *string
	SetResponseRuleId(v string) *DeleteResponseRuleRequest
	GetResponseRuleId() *string
}

type DeleteResponseRuleRequest struct {
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. If you do not specify this parameter, the service returns the first page of results.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the data management center for threat analysis. This must be the region where your assets are located. Valid values:
	//
	// - `cn-hangzhou`: for assets in the Chinese mainland or Hong Kong (China).
	//
	// - `ap-southeast-1`: for assets in international regions.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the automatic response rule to delete.
	//
	// example:
	//
	// 403235
	ResponseRuleId *string `json:"ResponseRuleId,omitempty" xml:"ResponseRuleId,omitempty"`
}

func (s DeleteResponseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResponseRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteResponseRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteResponseRuleRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DeleteResponseRuleRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DeleteResponseRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteResponseRuleRequest) GetResponseRuleId() *string {
	return s.ResponseRuleId
}

func (s *DeleteResponseRuleRequest) SetLang(v string) *DeleteResponseRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteResponseRuleRequest) SetMaxResults(v int32) *DeleteResponseRuleRequest {
	s.MaxResults = &v
	return s
}

func (s *DeleteResponseRuleRequest) SetNextToken(v string) *DeleteResponseRuleRequest {
	s.NextToken = &v
	return s
}

func (s *DeleteResponseRuleRequest) SetRegionId(v string) *DeleteResponseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteResponseRuleRequest) SetResponseRuleId(v string) *DeleteResponseRuleRequest {
	s.ResponseRuleId = &v
	return s
}

func (s *DeleteResponseRuleRequest) Validate() error {
	return dara.Validate(s)
}
