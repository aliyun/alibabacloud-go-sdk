// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcomVideoRecreationShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInputShrink(v string) *EcomVideoRecreationShrinkRequest
  GetInputShrink() *string 
  SetOutputShrink(v string) *EcomVideoRecreationShrinkRequest
  GetOutputShrink() *string 
}

type EcomVideoRecreationShrinkRequest struct {
  // The input parameters for video remix.
  // 
  // This parameter is required.
  InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
  // The output specifications for the final video.
  OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s EcomVideoRecreationShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EcomVideoRecreationShrinkRequest) GoString() string {
  return s.String()
}

func (s *EcomVideoRecreationShrinkRequest) GetInputShrink() *string  {
  return s.InputShrink
}

func (s *EcomVideoRecreationShrinkRequest) GetOutputShrink() *string  {
  return s.OutputShrink
}

func (s *EcomVideoRecreationShrinkRequest) SetInputShrink(v string) *EcomVideoRecreationShrinkRequest {
  s.InputShrink = &v
  return s
}

func (s *EcomVideoRecreationShrinkRequest) SetOutputShrink(v string) *EcomVideoRecreationShrinkRequest {
  s.OutputShrink = &v
  return s
}

func (s *EcomVideoRecreationShrinkRequest) Validate() error {
  return dara.Validate(s)
}

