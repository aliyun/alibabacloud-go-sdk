// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdleInstanceCullerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIdleInstanceCullerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIdleInstanceCullerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIdleInstanceCullerResponseBody) *DeleteIdleInstanceCullerResponse
	GetBody() *DeleteIdleInstanceCullerResponseBody
}

type DeleteIdleInstanceCullerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIdleInstanceCullerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIdleInstanceCullerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdleInstanceCullerResponse) GoString() string {
	return s.String()
}

func (s *DeleteIdleInstanceCullerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIdleInstanceCullerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIdleInstanceCullerResponse) GetBody() *DeleteIdleInstanceCullerResponseBody {
	return s.Body
}

func (s *DeleteIdleInstanceCullerResponse) SetHeaders(v map[string]*string) *DeleteIdleInstanceCullerResponse {
	s.Headers = v
	return s
}

func (s *DeleteIdleInstanceCullerResponse) SetStatusCode(v int32) *DeleteIdleInstanceCullerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponse) SetBody(v *DeleteIdleInstanceCullerResponseBody) *DeleteIdleInstanceCullerResponse {
	s.Body = v
	return s
}

func (s *DeleteIdleInstanceCullerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
