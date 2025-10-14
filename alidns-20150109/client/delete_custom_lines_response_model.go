// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomLinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomLinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomLinesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomLinesResponseBody) *DeleteCustomLinesResponse
	GetBody() *DeleteCustomLinesResponseBody
}

type DeleteCustomLinesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomLinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomLinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomLinesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomLinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomLinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomLinesResponse) GetBody() *DeleteCustomLinesResponseBody {
	return s.Body
}

func (s *DeleteCustomLinesResponse) SetHeaders(v map[string]*string) *DeleteCustomLinesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomLinesResponse) SetStatusCode(v int32) *DeleteCustomLinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomLinesResponse) SetBody(v *DeleteCustomLinesResponseBody) *DeleteCustomLinesResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomLinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
