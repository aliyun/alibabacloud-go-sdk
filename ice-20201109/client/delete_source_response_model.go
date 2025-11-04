// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSourceResponseBody) *DeleteSourceResponse
	GetBody() *DeleteSourceResponseBody
}

type DeleteSourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSourceResponse) GetBody() *DeleteSourceResponseBody {
	return s.Body
}

func (s *DeleteSourceResponse) SetHeaders(v map[string]*string) *DeleteSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSourceResponse) SetStatusCode(v int32) *DeleteSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSourceResponse) SetBody(v *DeleteSourceResponseBody) *DeleteSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
