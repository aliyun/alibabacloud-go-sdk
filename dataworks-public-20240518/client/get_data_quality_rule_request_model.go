// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataQualityRuleRequest
	GetId() *int64
}

type GetDataQualityRuleRequest struct {
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19715
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityRuleRequest) SetId(v int64) *GetDataQualityRuleRequest {
	s.Id = &v
	return s
}

func (s *GetDataQualityRuleRequest) Validate() error {
	return dara.Validate(s)
}
