// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenMetaKnowledgeAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenMetaKnowledgeAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenMetaKnowledgeAssetResponse
	GetStatusCode() *int32
	SetBody(v *GenMetaKnowledgeAssetResponseBody) *GenMetaKnowledgeAssetResponse
	GetBody() *GenMetaKnowledgeAssetResponseBody
}

type GenMetaKnowledgeAssetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenMetaKnowledgeAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenMetaKnowledgeAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s GenMetaKnowledgeAssetResponse) GoString() string {
	return s.String()
}

func (s *GenMetaKnowledgeAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenMetaKnowledgeAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenMetaKnowledgeAssetResponse) GetBody() *GenMetaKnowledgeAssetResponseBody {
	return s.Body
}

func (s *GenMetaKnowledgeAssetResponse) SetHeaders(v map[string]*string) *GenMetaKnowledgeAssetResponse {
	s.Headers = v
	return s
}

func (s *GenMetaKnowledgeAssetResponse) SetStatusCode(v int32) *GenMetaKnowledgeAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *GenMetaKnowledgeAssetResponse) SetBody(v *GenMetaKnowledgeAssetResponseBody) *GenMetaKnowledgeAssetResponse {
	s.Body = v
	return s
}

func (s *GenMetaKnowledgeAssetResponse) Validate() error {
	return dara.Validate(s)
}
