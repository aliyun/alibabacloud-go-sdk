// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportJobsShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *ExportJobsShrinkRequest
  GetAppName() *string 
  SetClusterId(v string) *ExportJobsShrinkRequest
  GetClusterId() *string 
  SetExportJobType(v int32) *ExportJobsShrinkRequest
  GetExportJobType() *int32 
  SetJobIdsShrink(v string) *ExportJobsShrinkRequest
  GetJobIdsShrink() *string 
}

type ExportJobsShrinkRequest struct {
  // The application name.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test-app
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // The cluster ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // xxljob-b6ec1xxxx
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // The type of jobs to export. The default value is 1.
  // 
  // example:
  // 
  // 1
  ExportJobType *int32 `json:"ExportJobType,omitempty" xml:"ExportJobType,omitempty"`
  // A list of job IDs.
  JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s ExportJobsShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportJobsShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportJobsShrinkRequest) GetAppName() *string  {
  return s.AppName
}

func (s *ExportJobsShrinkRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *ExportJobsShrinkRequest) GetExportJobType() *int32  {
  return s.ExportJobType
}

func (s *ExportJobsShrinkRequest) GetJobIdsShrink() *string  {
  return s.JobIdsShrink
}

func (s *ExportJobsShrinkRequest) SetAppName(v string) *ExportJobsShrinkRequest {
  s.AppName = &v
  return s
}

func (s *ExportJobsShrinkRequest) SetClusterId(v string) *ExportJobsShrinkRequest {
  s.ClusterId = &v
  return s
}

func (s *ExportJobsShrinkRequest) SetExportJobType(v int32) *ExportJobsShrinkRequest {
  s.ExportJobType = &v
  return s
}

func (s *ExportJobsShrinkRequest) SetJobIdsShrink(v string) *ExportJobsShrinkRequest {
  s.JobIdsShrink = &v
  return s
}

func (s *ExportJobsShrinkRequest) Validate() error {
  return dara.Validate(s)
}

