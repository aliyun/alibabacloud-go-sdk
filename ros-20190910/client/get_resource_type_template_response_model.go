// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceTypeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceTypeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceTypeTemplateResponseBody) *GetResourceTypeTemplateResponse
	GetBody() *GetResourceTypeTemplateResponseBody
}

type GetResourceTypeTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceTypeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceTypeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTypeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceTypeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceTypeTemplateResponse) GetBody() *GetResourceTypeTemplateResponseBody {
	return s.Body
}

func (s *GetResourceTypeTemplateResponse) SetHeaders(v map[string]*string) *GetResourceTypeTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTypeTemplateResponse) SetStatusCode(v int32) *GetResourceTypeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceTypeTemplateResponse) SetBody(v *GetResourceTypeTemplateResponseBody) *GetResourceTypeTemplateResponse {
	s.Body = v
	return s
}

func (s *GetResourceTypeTemplateResponse) Validate() error {
	return dara.Validate(s)
}
