// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditMetaKnowledgeAssetResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditMetaKnowledgeAssetResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditMetaKnowledgeAssetResponse
  GetStatusCode() *int32 
  SetBody(v *EditMetaKnowledgeAssetResponseBody) *EditMetaKnowledgeAssetResponse
  GetBody() *EditMetaKnowledgeAssetResponseBody 
}

type EditMetaKnowledgeAssetResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditMetaKnowledgeAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditMetaKnowledgeAssetResponse) String() string {
  return dara.Prettify(s)
}

func (s EditMetaKnowledgeAssetResponse) GoString() string {
  return s.String()
}

func (s *EditMetaKnowledgeAssetResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditMetaKnowledgeAssetResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditMetaKnowledgeAssetResponse) GetBody() *EditMetaKnowledgeAssetResponseBody  {
  return s.Body
}

func (s *EditMetaKnowledgeAssetResponse) SetHeaders(v map[string]*string) *EditMetaKnowledgeAssetResponse {
  s.Headers = v
  return s
}

func (s *EditMetaKnowledgeAssetResponse) SetStatusCode(v int32) *EditMetaKnowledgeAssetResponse {
  s.StatusCode = &v
  return s
}

func (s *EditMetaKnowledgeAssetResponse) SetBody(v *EditMetaKnowledgeAssetResponseBody) *EditMetaKnowledgeAssetResponse {
  s.Body = v
  return s
}

func (s *EditMetaKnowledgeAssetResponse) Validate() error {
  return dara.Validate(s)
}

