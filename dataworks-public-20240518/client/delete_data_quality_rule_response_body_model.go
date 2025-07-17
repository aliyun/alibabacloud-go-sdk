// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataQualityRuleResponseBody
	GetSuccess() *bool
}

type DeleteDataQualityRuleResponseBody struct {
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

func (s DeleteDataQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataQualityRuleResponseBody) SetRequestId(v string) *DeleteDataQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataQualityRuleResponseBody) SetSuccess(v bool) *DeleteDataQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataQualityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
