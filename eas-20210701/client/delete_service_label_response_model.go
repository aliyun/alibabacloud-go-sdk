// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceLabelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceLabelResponseBody) *DeleteServiceLabelResponse
	GetBody() *DeleteServiceLabelResponseBody
}

type DeleteServiceLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceLabelResponse) GetBody() *DeleteServiceLabelResponseBody {
	return s.Body
}

func (s *DeleteServiceLabelResponse) SetHeaders(v map[string]*string) *DeleteServiceLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceLabelResponse) SetStatusCode(v int32) *DeleteServiceLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceLabelResponse) SetBody(v *DeleteServiceLabelResponseBody) *DeleteServiceLabelResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
