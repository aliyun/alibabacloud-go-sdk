// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExporterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteExporterRuleResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteExporterRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteExporterRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteExporterRuleResponseBody
	GetSuccess() *bool
}

type DeleteExporterRuleResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A5F022D-AC7C-460E-94AE-B9E75083D023
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExporterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExporterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExporterRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteExporterRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteExporterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExporterRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteExporterRuleResponseBody) SetCode(v string) *DeleteExporterRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExporterRuleResponseBody) SetMessage(v string) *DeleteExporterRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExporterRuleResponseBody) SetRequestId(v string) *DeleteExporterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExporterRuleResponseBody) SetSuccess(v bool) *DeleteExporterRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteExporterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
