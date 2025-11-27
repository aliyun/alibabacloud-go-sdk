// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePoolResponse
	GetStatusCode() *int32
	SetBody(v *DeletePoolResponseBody) *DeletePoolResponse
	GetBody() *DeletePoolResponseBody
}

type DeletePoolResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePoolResponse) GoString() string {
	return s.String()
}

func (s *DeletePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePoolResponse) GetBody() *DeletePoolResponseBody {
	return s.Body
}

func (s *DeletePoolResponse) SetHeaders(v map[string]*string) *DeletePoolResponse {
	s.Headers = v
	return s
}

func (s *DeletePoolResponse) SetStatusCode(v int32) *DeletePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePoolResponse) SetBody(v *DeletePoolResponseBody) *DeletePoolResponse {
	s.Body = v
	return s
}

func (s *DeletePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
