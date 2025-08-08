// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLinkageAttributesTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLinkageAttributesTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLinkageAttributesTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetLinkageAttributesTemplateResponseBody) *GetLinkageAttributesTemplateResponse
	GetBody() *GetLinkageAttributesTemplateResponseBody
}

type GetLinkageAttributesTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLinkageAttributesTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLinkageAttributesTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLinkageAttributesTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetLinkageAttributesTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLinkageAttributesTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLinkageAttributesTemplateResponse) GetBody() *GetLinkageAttributesTemplateResponseBody {
	return s.Body
}

func (s *GetLinkageAttributesTemplateResponse) SetHeaders(v map[string]*string) *GetLinkageAttributesTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetLinkageAttributesTemplateResponse) SetStatusCode(v int32) *GetLinkageAttributesTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponse) SetBody(v *GetLinkageAttributesTemplateResponseBody) *GetLinkageAttributesTemplateResponse {
	s.Body = v
	return s
}

func (s *GetLinkageAttributesTemplateResponse) Validate() error {
	return dara.Validate(s)
}
