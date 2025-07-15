// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportHotTopicPlanningProposalsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportHotTopicPlanningProposalsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportHotTopicPlanningProposalsResponse
  GetStatusCode() *int32 
  SetBody(v *ExportHotTopicPlanningProposalsResponseBody) *ExportHotTopicPlanningProposalsResponse
  GetBody() *ExportHotTopicPlanningProposalsResponseBody 
}

type ExportHotTopicPlanningProposalsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportHotTopicPlanningProposalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportHotTopicPlanningProposalsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportHotTopicPlanningProposalsResponse) GoString() string {
  return s.String()
}

func (s *ExportHotTopicPlanningProposalsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportHotTopicPlanningProposalsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportHotTopicPlanningProposalsResponse) GetBody() *ExportHotTopicPlanningProposalsResponseBody  {
  return s.Body
}

func (s *ExportHotTopicPlanningProposalsResponse) SetHeaders(v map[string]*string) *ExportHotTopicPlanningProposalsResponse {
  s.Headers = v
  return s
}

func (s *ExportHotTopicPlanningProposalsResponse) SetStatusCode(v int32) *ExportHotTopicPlanningProposalsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsResponse) SetBody(v *ExportHotTopicPlanningProposalsResponseBody) *ExportHotTopicPlanningProposalsResponse {
  s.Body = v
  return s
}

func (s *ExportHotTopicPlanningProposalsResponse) Validate() error {
  return dara.Validate(s)
}

