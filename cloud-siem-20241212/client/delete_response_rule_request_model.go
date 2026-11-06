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
	// The maximum number of data records to return in the current request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the position to start reading. If this parameter is left empty, the reading starts from the beginning.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the data management center for threat analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located outside China.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the automated response rule.
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
