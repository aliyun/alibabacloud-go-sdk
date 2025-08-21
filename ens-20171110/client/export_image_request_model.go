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
  SetOSSBucket(v string) *ExportImageRequest
  GetOSSBucket() *string 
  SetOSSPrefix(v string) *ExportImageRequest
  GetOSSPrefix() *string 
  SetOSSRegionId(v string) *ExportImageRequest
  GetOSSRegionId() *string 
  SetRoleName(v string) *ExportImageRequest
  GetRoleName() *string 
}

type ExportImageRequest struct {
  // The ID of the image.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // m-5ragaz3s74b7go8ks7jp9****
  ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
  // The OSS bucket to which you want to export the image.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // whxyl****
  OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
  // The prefix of the object as which you want to store the image in the OSS bucket. The prefix must be 1 to 30 characters in length and can contain digits and letters.
  OSSPrefix *string `json:"OSSPrefix,omitempty" xml:"OSSPrefix,omitempty"`
  // The region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-beijing
  OSSRegionId *string `json:"OSSRegionId,omitempty" xml:"OSSRegionId,omitempty"`
  // The name of the Resource Access Management (RAM) role.
  // 
  // example:
  // 
  // AliyunMNSLoggingRole
  RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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

func (s *ExportImageRequest) GetOSSBucket() *string  {
  return s.OSSBucket
}

func (s *ExportImageRequest) GetOSSPrefix() *string  {
  return s.OSSPrefix
}

func (s *ExportImageRequest) GetOSSRegionId() *string  {
  return s.OSSRegionId
}

func (s *ExportImageRequest) GetRoleName() *string  {
  return s.RoleName
}

func (s *ExportImageRequest) SetImageId(v string) *ExportImageRequest {
  s.ImageId = &v
  return s
}

func (s *ExportImageRequest) SetOSSBucket(v string) *ExportImageRequest {
  s.OSSBucket = &v
  return s
}

func (s *ExportImageRequest) SetOSSPrefix(v string) *ExportImageRequest {
  s.OSSPrefix = &v
  return s
}

func (s *ExportImageRequest) SetOSSRegionId(v string) *ExportImageRequest {
  s.OSSRegionId = &v
  return s
}

func (s *ExportImageRequest) SetRoleName(v string) *ExportImageRequest {
  s.RoleName = &v
  return s
}

func (s *ExportImageRequest) Validate() error {
  return dara.Validate(s)
}

