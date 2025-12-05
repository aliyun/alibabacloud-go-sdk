// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceVersionResponseBody) *DeleteServiceVersionResponse
	GetBody() *DeleteServiceVersionResponseBody
}

type DeleteServiceVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceVersionResponse) GetBody() *DeleteServiceVersionResponseBody {
	return s.Body
}

func (s *DeleteServiceVersionResponse) SetHeaders(v map[string]*string) *DeleteServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceVersionResponse) SetStatusCode(v int32) *DeleteServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceVersionResponse) SetBody(v *DeleteServiceVersionResponseBody) *DeleteServiceVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
