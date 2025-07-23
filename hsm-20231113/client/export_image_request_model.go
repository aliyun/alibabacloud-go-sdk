// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportImageRequest interface {
  dara.Model
  String() string
  GoString() string
  SetImageId(v string) *ExportImageRequest
  GetImageId() *string 
  SetInstanceId(v string) *ExportImageRequest
  GetInstanceId() *string 
}

type ExportImageRequest struct {
  // The ID of the image.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // image-8vbdd5uc6v10ecn5****
  ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
  // The ID of the HSM.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // hsm-cn-vj30bil8****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ExportImageRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportImageRequest) GoString() string {
  return s.String()
}

func (s *ExportImageRequest) GetImageId() *string  {
  return s.ImageId
}

func (s *ExportImageRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportImageRequest) SetImageId(v string) *ExportImageRequest {
  s.ImageId = &v
  return s
}

func (s *ExportImageRequest) SetInstanceId(v string) *ExportImageRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportImageRequest) Validate() error {
  return dara.Validate(s)
}

