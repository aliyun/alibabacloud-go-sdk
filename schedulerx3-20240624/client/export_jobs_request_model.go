// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportJobsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *ExportJobsRequest
  GetAppName() *string 
  SetClusterId(v string) *ExportJobsRequest
  GetClusterId() *string 
  SetExportJobType(v int32) *ExportJobsRequest
  GetExportJobType() *int32 
  SetJobIds(v []*int64) *ExportJobsRequest
  GetJobIds() []*int64 
}

type ExportJobsRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // test-app
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // xxljob-b6ec1xxxx
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // example:
  // 
  // 1
  ExportJobType *int32 `json:"ExportJobType,omitempty" xml:"ExportJobType,omitempty"`
  // -
  JobIds []*int64 `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s ExportJobsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportJobsRequest) GoString() string {
  return s.String()
}

func (s *ExportJobsRequest) GetAppName() *string  {
  return s.AppName
}

func (s *ExportJobsRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *ExportJobsRequest) GetExportJobType() *int32  {
  return s.ExportJobType
}

func (s *ExportJobsRequest) GetJobIds() []*int64  {
  return s.JobIds
}

func (s *ExportJobsRequest) SetAppName(v string) *ExportJobsRequest {
  s.AppName = &v
  return s
}

func (s *ExportJobsRequest) SetClusterId(v string) *ExportJobsRequest {
  s.ClusterId = &v
  return s
}

func (s *ExportJobsRequest) SetExportJobType(v int32) *ExportJobsRequest {
  s.ExportJobType = &v
  return s
}

func (s *ExportJobsRequest) SetJobIds(v []*int64) *ExportJobsRequest {
  s.JobIds = v
  return s
}

func (s *ExportJobsRequest) Validate() error {
  return dara.Validate(s)
}

