// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaCoverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaCoverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaCoverResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaCoverResponseBody) *UpdateMediaCoverResponse
	GetBody() *UpdateMediaCoverResponseBody
}

type UpdateMediaCoverResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaCoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaCoverResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaCoverResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaCoverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaCoverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaCoverResponse) GetBody() *UpdateMediaCoverResponseBody {
	return s.Body
}

func (s *UpdateMediaCoverResponse) SetHeaders(v map[string]*string) *UpdateMediaCoverResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaCoverResponse) SetStatusCode(v int32) *UpdateMediaCoverResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaCoverResponse) SetBody(v *UpdateMediaCoverResponseBody) *UpdateMediaCoverResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaCoverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
