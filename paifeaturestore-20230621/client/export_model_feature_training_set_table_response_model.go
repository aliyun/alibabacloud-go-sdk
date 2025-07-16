// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportModelFeatureTrainingSetTableResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportModelFeatureTrainingSetTableResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportModelFeatureTrainingSetTableResponse
  GetStatusCode() *int32 
  SetBody(v *ExportModelFeatureTrainingSetTableResponseBody) *ExportModelFeatureTrainingSetTableResponse
  GetBody() *ExportModelFeatureTrainingSetTableResponseBody 
}

type ExportModelFeatureTrainingSetTableResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportModelFeatureTrainingSetTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportModelFeatureTrainingSetTableResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableResponse) GoString() string {
  return s.String()
}

func (s *ExportModelFeatureTrainingSetTableResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportModelFeatureTrainingSetTableResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportModelFeatureTrainingSetTableResponse) GetBody() *ExportModelFeatureTrainingSetTableResponseBody  {
  return s.Body
}

func (s *ExportModelFeatureTrainingSetTableResponse) SetHeaders(v map[string]*string) *ExportModelFeatureTrainingSetTableResponse {
  s.Headers = v
  return s
}

func (s *ExportModelFeatureTrainingSetTableResponse) SetStatusCode(v int32) *ExportModelFeatureTrainingSetTableResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportModelFeatureTrainingSetTableResponse) SetBody(v *ExportModelFeatureTrainingSetTableResponseBody) *ExportModelFeatureTrainingSetTableResponse {
  s.Body = v
  return s
}

func (s *ExportModelFeatureTrainingSetTableResponse) Validate() error {
  return dara.Validate(s)
}

