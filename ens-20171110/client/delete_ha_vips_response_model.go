// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHaVipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHaVipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHaVipsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHaVipsResponseBody) *DeleteHaVipsResponse
	GetBody() *DeleteHaVipsResponseBody
}

type DeleteHaVipsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHaVipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHaVipsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHaVipsResponse) GoString() string {
	return s.String()
}

func (s *DeleteHaVipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHaVipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHaVipsResponse) GetBody() *DeleteHaVipsResponseBody {
	return s.Body
}

func (s *DeleteHaVipsResponse) SetHeaders(v map[string]*string) *DeleteHaVipsResponse {
	s.Headers = v
	return s
}

func (s *DeleteHaVipsResponse) SetStatusCode(v int32) *DeleteHaVipsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHaVipsResponse) SetBody(v *DeleteHaVipsResponseBody) *DeleteHaVipsResponse {
	s.Body = v
	return s
}

func (s *DeleteHaVipsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
