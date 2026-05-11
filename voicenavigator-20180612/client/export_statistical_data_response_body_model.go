// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportStatisticalDataResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetExportTaskId(v string) *ExportStatisticalDataResponseBody
  GetExportTaskId() *string 
  SetRequestId(v string) *ExportStatisticalDataResponseBody
  GetRequestId() *string 
}

type ExportStatisticalDataResponseBody struct {
  // example:
  // 
  // 6be5a9f1-406e-424e-a17b-b6fb86ee3cc9
  ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
  // example:
  // 
  // c62e6789-28a8-41db-941e-171a01d3b3b9
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportStatisticalDataResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportStatisticalDataResponseBody) GoString() string {
  return s.String()
}

func (s *ExportStatisticalDataResponseBody) GetExportTaskId() *string  {
  return s.ExportTaskId
}

func (s *ExportStatisticalDataResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportStatisticalDataResponseBody) SetExportTaskId(v string) *ExportStatisticalDataResponseBody {
  s.ExportTaskId = &v
  return s
}

func (s *ExportStatisticalDataResponseBody) SetRequestId(v string) *ExportStatisticalDataResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportStatisticalDataResponseBody) Validate() error {
  return dara.Validate(s)
}

