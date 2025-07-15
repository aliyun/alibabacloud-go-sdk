// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomCallTaggingRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *ExportCustomCallTaggingRequest
  GetInstanceId() *string 
}

type ExportCustomCallTaggingRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // ccc-test
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ExportCustomCallTaggingRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomCallTaggingRequest) GoString() string {
  return s.String()
}

func (s *ExportCustomCallTaggingRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportCustomCallTaggingRequest) SetInstanceId(v string) *ExportCustomCallTaggingRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportCustomCallTaggingRequest) Validate() error {
  return dara.Validate(s)
}

