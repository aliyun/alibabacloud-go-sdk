// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRCSTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRCSTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRCSTemplateResponse
	GetStatusCode() *int32
	SetBody(v *QueryRCSTemplateResponseBody) *QueryRCSTemplateResponse
	GetBody() *QueryRCSTemplateResponseBody
}

type QueryRCSTemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRCSTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRCSTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRCSTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryRCSTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRCSTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRCSTemplateResponse) GetBody() *QueryRCSTemplateResponseBody {
	return s.Body
}

func (s *QueryRCSTemplateResponse) SetHeaders(v map[string]*string) *QueryRCSTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryRCSTemplateResponse) SetStatusCode(v int32) *QueryRCSTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRCSTemplateResponse) SetBody(v *QueryRCSTemplateResponseBody) *QueryRCSTemplateResponse {
	s.Body = v
	return s
}

func (s *QueryRCSTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
