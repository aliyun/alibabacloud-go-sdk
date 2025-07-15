// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportHotTopicPlanningProposalsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportHotTopicPlanningProposalsResponseBody
  GetCode() *string 
  SetData(v string) *ExportHotTopicPlanningProposalsResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportHotTopicPlanningProposalsResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportHotTopicPlanningProposalsResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportHotTopicPlanningProposalsResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportHotTopicPlanningProposalsResponseBody
  GetSuccess() *bool 
}

type ExportHotTopicPlanningProposalsResponseBody struct {
  // example:
  // 
  // NoData
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // 业务数据
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 1813ceee-7fe5-41b4-87e5-982a4d18cca5
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportHotTopicPlanningProposalsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportHotTopicPlanningProposalsResponseBody) GoString() string {
  return s.String()
}

func (s *ExportHotTopicPlanningProposalsResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportHotTopicPlanningProposalsResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportHotTopicPlanningProposalsResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportHotTopicPlanningProposalsResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportHotTopicPlanningProposalsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportHotTopicPlanningProposalsResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetCode(v string) *ExportHotTopicPlanningProposalsResponseBody {
  s.Code = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetData(v string) *ExportHotTopicPlanningProposalsResponseBody {
  s.Data = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetHttpStatusCode(v int32) *ExportHotTopicPlanningProposalsResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetMessage(v string) *ExportHotTopicPlanningProposalsResponseBody {
  s.Message = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetRequestId(v string) *ExportHotTopicPlanningProposalsResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetSuccess(v bool) *ExportHotTopicPlanningProposalsResponseBody {
  s.Success = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) Validate() error {
  return dara.Validate(s)
}

