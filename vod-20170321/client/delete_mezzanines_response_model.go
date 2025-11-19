// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMezzaninesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMezzaninesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMezzaninesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMezzaninesResponseBody) *DeleteMezzaninesResponse
	GetBody() *DeleteMezzaninesResponseBody
}

type DeleteMezzaninesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMezzaninesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMezzaninesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMezzaninesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMezzaninesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMezzaninesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMezzaninesResponse) GetBody() *DeleteMezzaninesResponseBody {
	return s.Body
}

func (s *DeleteMezzaninesResponse) SetHeaders(v map[string]*string) *DeleteMezzaninesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMezzaninesResponse) SetStatusCode(v int32) *DeleteMezzaninesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMezzaninesResponse) SetBody(v *DeleteMezzaninesResponseBody) *DeleteMezzaninesResponse {
	s.Body = v
	return s
}

func (s *DeleteMezzaninesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
