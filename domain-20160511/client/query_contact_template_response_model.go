// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContactTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryContactTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryContactTemplateResponse
	GetStatusCode() *int32
	SetBody(v *QueryContactTemplateResponseBody) *QueryContactTemplateResponse
	GetBody() *QueryContactTemplateResponseBody
}

type QueryContactTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContactTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContactTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryContactTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryContactTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryContactTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryContactTemplateResponse) GetBody() *QueryContactTemplateResponseBody {
	return s.Body
}

func (s *QueryContactTemplateResponse) SetHeaders(v map[string]*string) *QueryContactTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryContactTemplateResponse) SetStatusCode(v int32) *QueryContactTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContactTemplateResponse) SetBody(v *QueryContactTemplateResponseBody) *QueryContactTemplateResponse {
	s.Body = v
	return s
}

func (s *QueryContactTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
