// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePerspectiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePerspectiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePerspectiveResponse
	GetStatusCode() *int32
	SetBody(v *DeletePerspectiveResponseBody) *DeletePerspectiveResponse
	GetBody() *DeletePerspectiveResponseBody
}

type DeletePerspectiveResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePerspectiveResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *DeletePerspectiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePerspectiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePerspectiveResponse) GetBody() *DeletePerspectiveResponseBody {
	return s.Body
}

func (s *DeletePerspectiveResponse) SetHeaders(v map[string]*string) *DeletePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *DeletePerspectiveResponse) SetStatusCode(v int32) *DeletePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePerspectiveResponse) SetBody(v *DeletePerspectiveResponseBody) *DeletePerspectiveResponse {
	s.Body = v
	return s
}

func (s *DeletePerspectiveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
