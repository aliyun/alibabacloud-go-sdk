// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConsumerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConsumerResponseBody) *DeleteConsumerResponse
	GetBody() *DeleteConsumerResponseBody
}

type DeleteConsumerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConsumerResponse) GetBody() *DeleteConsumerResponseBody {
	return s.Body
}

func (s *DeleteConsumerResponse) SetHeaders(v map[string]*string) *DeleteConsumerResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerResponse) SetStatusCode(v int32) *DeleteConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConsumerResponse) SetBody(v *DeleteConsumerResponseBody) *DeleteConsumerResponse {
	s.Body = v
	return s
}

func (s *DeleteConsumerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
