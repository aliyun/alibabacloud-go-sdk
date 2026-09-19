// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTemplateCacheResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeTemplateCacheResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTemplateCacheResponseBody
	GetRequestId() *string
	SetTemplateCache(v *PublicTemplateCache) *DescribeTemplateCacheResponseBody
	GetTemplateCache() *PublicTemplateCache
}

type DescribeTemplateCacheResponseBody struct {
	Code          *string              `json:"code,omitempty" xml:"code,omitempty"`
	Message       *string              `json:"message,omitempty" xml:"message,omitempty"`
	RequestId     *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TemplateCache *PublicTemplateCache `json:"templateCache,omitempty" xml:"templateCache,omitempty"`
}

func (s DescribeTemplateCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateCacheResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTemplateCacheResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTemplateCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplateCacheResponseBody) GetTemplateCache() *PublicTemplateCache {
	return s.TemplateCache
}

func (s *DescribeTemplateCacheResponseBody) SetCode(v string) *DescribeTemplateCacheResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTemplateCacheResponseBody) SetMessage(v string) *DescribeTemplateCacheResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTemplateCacheResponseBody) SetRequestId(v string) *DescribeTemplateCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateCacheResponseBody) SetTemplateCache(v *PublicTemplateCache) *DescribeTemplateCacheResponseBody {
	s.TemplateCache = v
	return s
}

func (s *DescribeTemplateCacheResponseBody) Validate() error {
	if s.TemplateCache != nil {
		if err := s.TemplateCache.Validate(); err != nil {
			return err
		}
	}
	return nil
}
