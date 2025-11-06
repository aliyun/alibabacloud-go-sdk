// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactTemplatesResponseBody) *DeleteContactTemplatesResponse
	GetBody() *DeleteContactTemplatesResponseBody
}

type DeleteContactTemplatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactTemplatesResponse) GetBody() *DeleteContactTemplatesResponseBody {
	return s.Body
}

func (s *DeleteContactTemplatesResponse) SetHeaders(v map[string]*string) *DeleteContactTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactTemplatesResponse) SetStatusCode(v int32) *DeleteContactTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactTemplatesResponse) SetBody(v *DeleteContactTemplatesResponseBody) *DeleteContactTemplatesResponse {
	s.Body = v
	return s
}

func (s *DeleteContactTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
