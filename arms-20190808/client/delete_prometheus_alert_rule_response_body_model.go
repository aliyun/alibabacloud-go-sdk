// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeletePrometheusAlertRuleResponseBody
	GetCode() *int64
	SetMessage(v string) *DeletePrometheusAlertRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePrometheusAlertRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePrometheusAlertRuleResponseBody
	GetSuccess() *bool
}

type DeletePrometheusAlertRuleResponseBody struct {
	// Status code. 200 means success, other status codes are exceptions.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// More Information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9FEA6D00-317F-45E3-9004-7FB8B0B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the alert rule was deleted. Valid values:
	//
	// 	- `true`: The alert rule was deleted.
	//
	// 	- `false`: The alert rule failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePrometheusAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeletePrometheusAlertRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePrometheusAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrometheusAlertRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePrometheusAlertRuleResponseBody) SetCode(v int64) *DeletePrometheusAlertRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponseBody) SetMessage(v string) *DeletePrometheusAlertRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponseBody) SetRequestId(v string) *DeletePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponseBody) SetSuccess(v bool) *DeletePrometheusAlertRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
