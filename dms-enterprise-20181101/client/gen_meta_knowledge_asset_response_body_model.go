// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenMetaKnowledgeAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GenMetaKnowledgeAssetResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GenMetaKnowledgeAssetResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GenMetaKnowledgeAssetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenMetaKnowledgeAssetResponseBody
	GetSuccess() *bool
}

type GenMetaKnowledgeAssetResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenMetaKnowledgeAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenMetaKnowledgeAssetResponseBody) GoString() string {
	return s.String()
}

func (s *GenMetaKnowledgeAssetResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GenMetaKnowledgeAssetResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GenMetaKnowledgeAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenMetaKnowledgeAssetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenMetaKnowledgeAssetResponseBody) SetErrorCode(v string) *GenMetaKnowledgeAssetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GenMetaKnowledgeAssetResponseBody) SetErrorMessage(v string) *GenMetaKnowledgeAssetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GenMetaKnowledgeAssetResponseBody) SetRequestId(v string) *GenMetaKnowledgeAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenMetaKnowledgeAssetResponseBody) SetSuccess(v bool) *GenMetaKnowledgeAssetResponseBody {
	s.Success = &v
	return s
}

func (s *GenMetaKnowledgeAssetResponseBody) Validate() error {
	return dara.Validate(s)
}
