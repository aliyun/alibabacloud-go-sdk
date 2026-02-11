// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportPrometheusRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ImportPrometheusRulesResponseBody
	GetData() *string
	SetRequestId(v string) *ImportPrometheusRulesResponseBody
	GetRequestId() *string
}

type ImportPrometheusRulesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportPrometheusRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportPrometheusRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportPrometheusRulesResponseBody) GetData() *string {
	return s.Data
}

func (s *ImportPrometheusRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportPrometheusRulesResponseBody) SetData(v string) *ImportPrometheusRulesResponseBody {
	s.Data = &v
	return s
}

func (s *ImportPrometheusRulesResponseBody) SetRequestId(v string) *ImportPrometheusRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportPrometheusRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
