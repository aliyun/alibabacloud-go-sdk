// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceControlResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceControlResponseBody) *DeleteResourceControlResponse
	GetBody() *DeleteResourceControlResponseBody
}

type DeleteResourceControlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceControlResponse) GetBody() *DeleteResourceControlResponseBody {
	return s.Body
}

func (s *DeleteResourceControlResponse) SetHeaders(v map[string]*string) *DeleteResourceControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceControlResponse) SetStatusCode(v int32) *DeleteResourceControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceControlResponse) SetBody(v *DeleteResourceControlResponseBody) *DeleteResourceControlResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
