// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableImageRequest interface {
  dara.Model
  String() string
  GoString() string
  SetId(v string) *EnableImageRequest
  GetId() *string 
}

type EnableImageRequest struct {
  // The image ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Custom_image_xxxx_xxxx
  Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s EnableImageRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableImageRequest) GoString() string {
  return s.String()
}

func (s *EnableImageRequest) GetId() *string  {
  return s.Id
}

func (s *EnableImageRequest) SetId(v string) *EnableImageRequest {
  s.Id = &v
  return s
}

func (s *EnableImageRequest) Validate() error {
  return dara.Validate(s)
}

