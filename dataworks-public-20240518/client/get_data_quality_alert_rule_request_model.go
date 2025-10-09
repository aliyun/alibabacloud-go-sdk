// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataQualityAlertRuleRequest
	GetId() *int64
}

type GetDataQualityAlertRuleRequest struct {
	// The data quality monitoring alert rule ID.
	//
	// example:
	//
	// 113642
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataQualityAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityAlertRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityAlertRuleRequest) SetId(v int64) *GetDataQualityAlertRuleRequest {
	s.Id = &v
	return s
}

func (s *GetDataQualityAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
