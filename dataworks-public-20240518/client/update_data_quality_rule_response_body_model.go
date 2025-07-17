// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataQualityRuleResponseBody
	GetSuccess() *bool
}

type UpdateDataQualityRuleResponseBody struct {
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

func (s UpdateDataQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataQualityRuleResponseBody) SetRequestId(v string) *UpdateDataQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataQualityRuleResponseBody) SetSuccess(v bool) *UpdateDataQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataQualityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
