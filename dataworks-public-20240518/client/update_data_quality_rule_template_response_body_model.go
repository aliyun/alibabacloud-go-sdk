// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityRuleTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataQualityRuleTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataQualityRuleTemplateResponseBody
	GetSuccess() *bool
}

type UpdateDataQualityRuleTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataQualityRuleTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataQualityRuleTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataQualityRuleTemplateResponseBody) SetRequestId(v string) *UpdateDataQualityRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateResponseBody) SetSuccess(v bool) *UpdateDataQualityRuleTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
