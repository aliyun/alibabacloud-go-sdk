// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateDataQualityRuleResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDataQualityRuleResponseBody
	GetRequestId() *string
}

type CreateDataQualityRuleResponseBody struct {
	// The ID of the rule.
	//
	// example:
	//
	// 19715
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDataQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataQualityRuleResponseBody) SetId(v int64) *CreateDataQualityRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataQualityRuleResponseBody) SetRequestId(v string) *CreateDataQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataQualityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
