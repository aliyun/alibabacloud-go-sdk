// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportPrometheusRulesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *ExportPrometheusRulesResponseBody
  GetData() *string 
  SetRequestId(v string) *ExportPrometheusRulesResponseBody
  GetRequestId() *string 
}

type ExportPrometheusRulesResponseBody struct {
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportPrometheusRulesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportPrometheusRulesResponseBody) GoString() string {
  return s.String()
}

func (s *ExportPrometheusRulesResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportPrometheusRulesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportPrometheusRulesResponseBody) SetData(v string) *ExportPrometheusRulesResponseBody {
  s.Data = &v
  return s
}

func (s *ExportPrometheusRulesResponseBody) SetRequestId(v string) *ExportPrometheusRulesResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportPrometheusRulesResponseBody) Validate() error {
  return dara.Validate(s)
}

