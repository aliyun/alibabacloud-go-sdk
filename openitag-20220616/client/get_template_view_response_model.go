// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateViewResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateViewResponseBody) *GetTemplateViewResponse
	GetBody() *GetTemplateViewResponseBody
}

type GetTemplateViewResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateViewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateViewResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateViewResponse) GetBody() *GetTemplateViewResponseBody {
	return s.Body
}

func (s *GetTemplateViewResponse) SetHeaders(v map[string]*string) *GetTemplateViewResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateViewResponse) SetStatusCode(v int32) *GetTemplateViewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateViewResponse) SetBody(v *GetTemplateViewResponseBody) *GetTemplateViewResponse {
	s.Body = v
	return s
}

func (s *GetTemplateViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
