// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProjectLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProjectLabelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProjectLabelResponseBody) *DeleteProjectLabelResponse
	GetBody() *DeleteProjectLabelResponseBody
}

type DeleteProjectLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProjectLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProjectLabelResponse) GetBody() *DeleteProjectLabelResponseBody {
	return s.Body
}

func (s *DeleteProjectLabelResponse) SetHeaders(v map[string]*string) *DeleteProjectLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectLabelResponse) SetStatusCode(v int32) *DeleteProjectLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectLabelResponse) SetBody(v *DeleteProjectLabelResponseBody) *DeleteProjectLabelResponse {
	s.Body = v
	return s
}

func (s *DeleteProjectLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
