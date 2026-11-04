// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTimeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTimeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTimeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTimeTemplateResponseBody) *DeleteTimeTemplateResponse
	GetBody() *DeleteTimeTemplateResponseBody
}

type DeleteTimeTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTimeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTimeTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTimeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTimeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTimeTemplateResponse) GetBody() *DeleteTimeTemplateResponseBody {
	return s.Body
}

func (s *DeleteTimeTemplateResponse) SetHeaders(v map[string]*string) *DeleteTimeTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTimeTemplateResponse) SetStatusCode(v int32) *DeleteTimeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTimeTemplateResponse) SetBody(v *DeleteTimeTemplateResponseBody) *DeleteTimeTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteTimeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
