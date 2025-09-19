// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWarningResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetFileName(v string) *ExportWarningResponseBody
  GetFileName() *string 
  SetId(v int64) *ExportWarningResponseBody
  GetId() *int64 
  SetRequestId(v string) *ExportWarningResponseBody
  GetRequestId() *string 
}

type ExportWarningResponseBody struct {
  // The name of the file that contains exported baseline check results.
  // 
  // example:
  // 
  // health_check_export_20220407
  FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
  // The ID of the task to export baseline check results.
  // 
  // > You can call use the value of this parameter to call the [DescribeHcExportInfo](~~DescribeHcExportInfo~~) operation to query the export progress.
  // 
  // example:
  // 
  // 439316
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  // The ID of the request, which is used to locate and troubleshoot issues.
  // 
  // example:
  // 
  // A7FC828B-C242-1005-9736-C7CC5DC09FF0
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportWarningResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportWarningResponseBody) GoString() string {
  return s.String()
}

func (s *ExportWarningResponseBody) GetFileName() *string  {
  return s.FileName
}

func (s *ExportWarningResponseBody) GetId() *int64  {
  return s.Id
}

func (s *ExportWarningResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportWarningResponseBody) SetFileName(v string) *ExportWarningResponseBody {
  s.FileName = &v
  return s
}

func (s *ExportWarningResponseBody) SetId(v int64) *ExportWarningResponseBody {
  s.Id = &v
  return s
}

func (s *ExportWarningResponseBody) SetRequestId(v string) *ExportWarningResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportWarningResponseBody) Validate() error {
  return dara.Validate(s)
}

