// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorityTemplateItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthorityTemplateItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthorityTemplateItemResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthorityTemplateItemResponseBody) *GetAuthorityTemplateItemResponse
	GetBody() *GetAuthorityTemplateItemResponseBody
}

type GetAuthorityTemplateItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthorityTemplateItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthorityTemplateItemResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateItemResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthorityTemplateItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthorityTemplateItemResponse) GetBody() *GetAuthorityTemplateItemResponseBody {
	return s.Body
}

func (s *GetAuthorityTemplateItemResponse) SetHeaders(v map[string]*string) *GetAuthorityTemplateItemResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorityTemplateItemResponse) SetStatusCode(v int32) *GetAuthorityTemplateItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthorityTemplateItemResponse) SetBody(v *GetAuthorityTemplateItemResponseBody) *GetAuthorityTemplateItemResponse {
	s.Body = v
	return s
}

func (s *GetAuthorityTemplateItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
