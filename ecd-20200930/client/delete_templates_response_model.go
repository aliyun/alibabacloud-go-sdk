// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTemplatesResponseBody) *DeleteTemplatesResponse
	GetBody() *DeleteTemplatesResponseBody
}

type DeleteTemplatesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTemplatesResponse) GetBody() *DeleteTemplatesResponseBody {
	return s.Body
}

func (s *DeleteTemplatesResponse) SetHeaders(v map[string]*string) *DeleteTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplatesResponse) SetStatusCode(v int32) *DeleteTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplatesResponse) SetBody(v *DeleteTemplatesResponseBody) *DeleteTemplatesResponse {
	s.Body = v
	return s
}

func (s *DeleteTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
