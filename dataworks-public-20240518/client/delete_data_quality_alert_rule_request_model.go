// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteDataQualityAlertRuleRequest
	GetId() *int64
}

type DeleteDataQualityAlertRuleRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDataQualityAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityAlertRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDataQualityAlertRuleRequest) SetId(v int64) *DeleteDataQualityAlertRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataQualityAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
