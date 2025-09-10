// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateQuotaItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTemplateQuotaItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTemplateQuotaItemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTemplateQuotaItemResponseBody) *DeleteTemplateQuotaItemResponse
	GetBody() *DeleteTemplateQuotaItemResponseBody
}

type DeleteTemplateQuotaItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateQuotaItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateQuotaItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateQuotaItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateQuotaItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTemplateQuotaItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTemplateQuotaItemResponse) GetBody() *DeleteTemplateQuotaItemResponseBody {
	return s.Body
}

func (s *DeleteTemplateQuotaItemResponse) SetHeaders(v map[string]*string) *DeleteTemplateQuotaItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateQuotaItemResponse) SetStatusCode(v int32) *DeleteTemplateQuotaItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateQuotaItemResponse) SetBody(v *DeleteTemplateQuotaItemResponseBody) *DeleteTemplateQuotaItemResponse {
	s.Body = v
	return s
}

func (s *DeleteTemplateQuotaItemResponse) Validate() error {
	return dara.Validate(s)
}
