// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetQualityRuleDetailRequest
	GetInstanceId() *string
	SetQualityRuleId(v int64) *GetQualityRuleDetailRequest
	GetQualityRuleId() *int64
}

type GetQualityRuleDetailRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	QualityRuleId *int64 `json:"QualityRuleId,omitempty" xml:"QualityRuleId,omitempty"`
}

func (s GetQualityRuleDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQualityRuleDetailRequest) GetQualityRuleId() *int64 {
	return s.QualityRuleId
}

func (s *GetQualityRuleDetailRequest) SetInstanceId(v string) *GetQualityRuleDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityRuleDetailRequest) SetQualityRuleId(v int64) *GetQualityRuleDetailRequest {
	s.QualityRuleId = &v
	return s
}

func (s *GetQualityRuleDetailRequest) Validate() error {
	return dara.Validate(s)
}
