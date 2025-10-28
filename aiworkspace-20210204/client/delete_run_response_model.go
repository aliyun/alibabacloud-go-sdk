// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRunResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRunResponseBody) *DeleteRunResponse
	GetBody() *DeleteRunResponseBody
}

type DeleteRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRunResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRunResponse) GoString() string {
	return s.String()
}

func (s *DeleteRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRunResponse) GetBody() *DeleteRunResponseBody {
	return s.Body
}

func (s *DeleteRunResponse) SetHeaders(v map[string]*string) *DeleteRunResponse {
	s.Headers = v
	return s
}

func (s *DeleteRunResponse) SetStatusCode(v int32) *DeleteRunResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRunResponse) SetBody(v *DeleteRunResponseBody) *DeleteRunResponse {
	s.Body = v
	return s
}

func (s *DeleteRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
