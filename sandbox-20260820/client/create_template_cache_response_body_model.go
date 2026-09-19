// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTemplateCacheResponseBody
	GetCode() *string
	SetMessage(v string) *CreateTemplateCacheResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTemplateCacheResponseBody
	GetRequestId() *string
	SetTemplateCache(v *PublicTemplateCache) *CreateTemplateCacheResponseBody
	GetTemplateCache() *PublicTemplateCache
}

type CreateTemplateCacheResponseBody struct {
	Code          *string              `json:"code,omitempty" xml:"code,omitempty"`
	Message       *string              `json:"message,omitempty" xml:"message,omitempty"`
	RequestId     *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TemplateCache *PublicTemplateCache `json:"templateCache,omitempty" xml:"templateCache,omitempty"`
}

func (s CreateTemplateCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateCacheResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTemplateCacheResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTemplateCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTemplateCacheResponseBody) GetTemplateCache() *PublicTemplateCache {
	return s.TemplateCache
}

func (s *CreateTemplateCacheResponseBody) SetCode(v string) *CreateTemplateCacheResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTemplateCacheResponseBody) SetMessage(v string) *CreateTemplateCacheResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTemplateCacheResponseBody) SetRequestId(v string) *CreateTemplateCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateCacheResponseBody) SetTemplateCache(v *PublicTemplateCache) *CreateTemplateCacheResponseBody {
	s.TemplateCache = v
	return s
}

func (s *CreateTemplateCacheResponseBody) Validate() error {
	if s.TemplateCache != nil {
		if err := s.TemplateCache.Validate(); err != nil {
			return err
		}
	}
	return nil
}
