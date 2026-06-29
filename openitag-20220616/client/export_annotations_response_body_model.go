// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnnotationsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExportAnnotationsResponseBody
  GetCode() *int32 
  SetDetails(v string) *ExportAnnotationsResponseBody
  GetDetails() *string 
  SetErrorCode(v string) *ExportAnnotationsResponseBody
  GetErrorCode() *string 
  SetFlowJob(v *FlowJobInfo) *ExportAnnotationsResponseBody
  GetFlowJob() *FlowJobInfo 
  SetMessage(v string) *ExportAnnotationsResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportAnnotationsResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportAnnotationsResponseBody
  GetSuccess() *bool 
}

type ExportAnnotationsResponseBody struct {
  // Return encoding. The default value is 0, indicating Normal execution.
  // 
  // example:
  // 
  // 0
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // Details.
  // 
  // example:
  // 
  // null
  Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
  // Error code.
  // 
  // - When Success is false, returns a business error code.
  // 
  // - When Success is true, returns an empty value.
  // 
  // example:
  // 
  // ""
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // Pipeline.
  FlowJob *FlowJobInfo `json:"FlowJob,omitempty" xml:"FlowJob,omitempty"`
  // The response message of the request.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Request ID.
  // 
  // example:
  // 
  // 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the operation succeeded. Valid values:
  // 
  // - true: Succeeded.
  // 
  // - false: Failed.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportAnnotationsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportAnnotationsResponseBody) GoString() string {
  return s.String()
}

func (s *ExportAnnotationsResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExportAnnotationsResponseBody) GetDetails() *string  {
  return s.Details
}

func (s *ExportAnnotationsResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExportAnnotationsResponseBody) GetFlowJob() *FlowJobInfo  {
  return s.FlowJob
}

func (s *ExportAnnotationsResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportAnnotationsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportAnnotationsResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportAnnotationsResponseBody) SetCode(v int32) *ExportAnnotationsResponseBody {
  s.Code = &v
  return s
}

func (s *ExportAnnotationsResponseBody) SetDetails(v string) *ExportAnnotationsResponseBody {
  s.Details = &v
  return s
}

func (s *ExportAnnotationsResponseBody) SetErrorCode(v string) *ExportAnnotationsResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExportAnnotationsResponseBody) SetFlowJob(v *FlowJobInfo) *ExportAnnotationsResponseBody {
  s.FlowJob = v
  return s
}

func (s *ExportAnnotationsResponseBody) SetMessage(v string) *ExportAnnotationsResponseBody {
  s.Message = &v
  return s
}

func (s *ExportAnnotationsResponseBody) SetRequestId(v string) *ExportAnnotationsResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportAnnotationsResponseBody) SetSuccess(v bool) *ExportAnnotationsResponseBody {
  s.Success = &v
  return s
}

func (s *ExportAnnotationsResponseBody) Validate() error {
  if s.FlowJob != nil {
    if err := s.FlowJob.Validate(); err != nil {
      return err
    }
  }
  return nil
}

