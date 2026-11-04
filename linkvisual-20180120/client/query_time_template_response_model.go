// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTimeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTimeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTimeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *QueryTimeTemplateResponseBody) *QueryTimeTemplateResponse
	GetBody() *QueryTimeTemplateResponseBody
}

type QueryTimeTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTimeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTimeTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTimeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTimeTemplateResponse) GetBody() *QueryTimeTemplateResponseBody {
	return s.Body
}

func (s *QueryTimeTemplateResponse) SetHeaders(v map[string]*string) *QueryTimeTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryTimeTemplateResponse) SetStatusCode(v int32) *QueryTimeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTimeTemplateResponse) SetBody(v *QueryTimeTemplateResponseBody) *QueryTimeTemplateResponse {
	s.Body = v
	return s
}

func (s *QueryTimeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
