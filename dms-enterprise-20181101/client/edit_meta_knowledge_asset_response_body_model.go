// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditMetaKnowledgeAssetResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrorCode(v string) *EditMetaKnowledgeAssetResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EditMetaKnowledgeAssetResponseBody
  GetErrorMessage() *string 
  SetRequestId(v string) *EditMetaKnowledgeAssetResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditMetaKnowledgeAssetResponseBody
  GetSuccess() *bool 
}

type EditMetaKnowledgeAssetResponseBody struct {
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditMetaKnowledgeAssetResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditMetaKnowledgeAssetResponseBody) GoString() string {
  return s.String()
}

func (s *EditMetaKnowledgeAssetResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EditMetaKnowledgeAssetResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EditMetaKnowledgeAssetResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditMetaKnowledgeAssetResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditMetaKnowledgeAssetResponseBody) SetErrorCode(v string) *EditMetaKnowledgeAssetResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EditMetaKnowledgeAssetResponseBody) SetErrorMessage(v string) *EditMetaKnowledgeAssetResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EditMetaKnowledgeAssetResponseBody) SetRequestId(v string) *EditMetaKnowledgeAssetResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditMetaKnowledgeAssetResponseBody) SetSuccess(v bool) *EditMetaKnowledgeAssetResponseBody {
  s.Success = &v
  return s
}

func (s *EditMetaKnowledgeAssetResponseBody) Validate() error {
  return dara.Validate(s)
}

