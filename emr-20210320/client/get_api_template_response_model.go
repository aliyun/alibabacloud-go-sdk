// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetApiTemplateResponseBody) *GetApiTemplateResponse
	GetBody() *GetApiTemplateResponseBody
}

type GetApiTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetApiTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiTemplateResponse) GetBody() *GetApiTemplateResponseBody {
	return s.Body
}

func (s *GetApiTemplateResponse) SetHeaders(v map[string]*string) *GetApiTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetApiTemplateResponse) SetStatusCode(v int32) *GetApiTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiTemplateResponse) SetBody(v *GetApiTemplateResponseBody) *GetApiTemplateResponse {
	s.Body = v
	return s
}

func (s *GetApiTemplateResponse) Validate() error {
	return dara.Validate(s)
}
