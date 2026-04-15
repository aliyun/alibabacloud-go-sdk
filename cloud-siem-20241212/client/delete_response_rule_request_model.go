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
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
