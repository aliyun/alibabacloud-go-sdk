// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhancePortraitVideoResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnhancePortraitVideoResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnhancePortraitVideoResponse
  GetStatusCode() *int32 
  SetBody(v *EnhancePortraitVideoResponseBody) *EnhancePortraitVideoResponse
  GetBody() *EnhancePortraitVideoResponseBody 
}

type EnhancePortraitVideoResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnhancePortraitVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnhancePortraitVideoResponse) String() string {
  return dara.Prettify(s)
}

func (s EnhancePortraitVideoResponse) GoString() string {
  return s.String()
}

func (s *EnhancePortraitVideoResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnhancePortraitVideoResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnhancePortraitVideoResponse) GetBody() *EnhancePortraitVideoResponseBody  {
  return s.Body
}

func (s *EnhancePortraitVideoResponse) SetHeaders(v map[string]*string) *EnhancePortraitVideoResponse {
  s.Headers = v
  return s
}

func (s *EnhancePortraitVideoResponse) SetStatusCode(v int32) *EnhancePortraitVideoResponse {
  s.StatusCode = &v
  return s
}

func (s *EnhancePortraitVideoResponse) SetBody(v *EnhancePortraitVideoResponseBody) *EnhancePortraitVideoResponse {
  s.Body = v
  return s
}

func (s *EnhancePortraitVideoResponse) Validate() error {
  return dara.Validate(s)
}

