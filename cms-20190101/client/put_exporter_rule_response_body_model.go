// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutExporterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutExporterRuleResponseBody
	GetCode() *string
	SetMessage(v string) *PutExporterRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutExporterRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutExporterRuleResponseBody
	GetSuccess() *bool
}

type PutExporterRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutExporterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutExporterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutExporterRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutExporterRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutExporterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutExporterRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutExporterRuleResponseBody) SetCode(v string) *PutExporterRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutExporterRuleResponseBody) SetMessage(v string) *PutExporterRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutExporterRuleResponseBody) SetRequestId(v string) *PutExporterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutExporterRuleResponseBody) SetSuccess(v bool) *PutExporterRuleResponseBody {
	s.Success = &v
	return s
}

func (s *PutExporterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
