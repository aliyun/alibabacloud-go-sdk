// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePxfsUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePxfsUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePxfsUserResponse
	GetStatusCode() *int32
	SetBody(v *DeletePxfsUserResponseBody) *DeletePxfsUserResponse
	GetBody() *DeletePxfsUserResponseBody
}

type DeletePxfsUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePxfsUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePxfsUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsUserResponse) GoString() string {
	return s.String()
}

func (s *DeletePxfsUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePxfsUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePxfsUserResponse) GetBody() *DeletePxfsUserResponseBody {
	return s.Body
}

func (s *DeletePxfsUserResponse) SetHeaders(v map[string]*string) *DeletePxfsUserResponse {
	s.Headers = v
	return s
}

func (s *DeletePxfsUserResponse) SetStatusCode(v int32) *DeletePxfsUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePxfsUserResponse) SetBody(v *DeletePxfsUserResponseBody) *DeletePxfsUserResponse {
	s.Body = v
	return s
}

func (s *DeletePxfsUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
