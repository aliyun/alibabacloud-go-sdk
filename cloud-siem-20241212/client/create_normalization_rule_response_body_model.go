// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNormalizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNormalizationRuleId(v string) *CreateNormalizationRuleResponseBody
	GetNormalizationRuleId() *string
	SetRequestId(v string) *CreateNormalizationRuleResponseBody
	GetRequestId() *string
}

type CreateNormalizationRuleResponseBody struct {
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp。
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNormalizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNormalizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNormalizationRuleResponseBody) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *CreateNormalizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNormalizationRuleResponseBody) SetNormalizationRuleId(v string) *CreateNormalizationRuleResponseBody {
	s.NormalizationRuleId = &v
	return s
}

func (s *CreateNormalizationRuleResponseBody) SetRequestId(v string) *CreateNormalizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNormalizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
