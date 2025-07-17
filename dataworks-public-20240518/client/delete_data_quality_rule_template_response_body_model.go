// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityRuleTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataQualityRuleTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataQualityRuleTemplateResponseBody
	GetSuccess() *bool
}

type DeleteDataQualityRuleTemplateResponseBody struct {
	// The request ID.
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

func (s DeleteDataQualityRuleTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityRuleTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataQualityRuleTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataQualityRuleTemplateResponseBody) SetRequestId(v string) *DeleteDataQualityRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataQualityRuleTemplateResponseBody) SetSuccess(v bool) *DeleteDataQualityRuleTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataQualityRuleTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
