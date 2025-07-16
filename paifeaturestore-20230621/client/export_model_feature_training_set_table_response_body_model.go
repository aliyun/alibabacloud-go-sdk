// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportModelFeatureTrainingSetTableResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExportModelFeatureTrainingSetTableResponseBody
  GetRequestId() *string 
  SetTaskId(v string) *ExportModelFeatureTrainingSetTableResponseBody
  GetTaskId() *string 
}

type ExportModelFeatureTrainingSetTableResponseBody struct {
  // example:
  // 
  // 0FBBE454-9BD1-5D8F-9129-D14DB7FAFE0B
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExportModelFeatureTrainingSetTableResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableResponseBody) GoString() string {
  return s.String()
}

func (s *ExportModelFeatureTrainingSetTableResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportModelFeatureTrainingSetTableResponseBody) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExportModelFeatureTrainingSetTableResponseBody) SetRequestId(v string) *ExportModelFeatureTrainingSetTableResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportModelFeatureTrainingSetTableResponseBody) SetTaskId(v string) *ExportModelFeatureTrainingSetTableResponseBody {
  s.TaskId = &v
  return s
}

func (s *ExportModelFeatureTrainingSetTableResponseBody) Validate() error {
  return dara.Validate(s)
}

