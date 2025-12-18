// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConfigRuleEvaluationByResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *StartConfigRuleEvaluationByResourceRequest
	GetResourceId() *string
}

type StartConfigRuleEvaluationByResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-bp151g9tpto890zr****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s StartConfigRuleEvaluationByResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartConfigRuleEvaluationByResourceRequest) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationByResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *StartConfigRuleEvaluationByResourceRequest) SetResourceId(v string) *StartConfigRuleEvaluationByResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *StartConfigRuleEvaluationByResourceRequest) Validate() error {
	return dara.Validate(s)
}
