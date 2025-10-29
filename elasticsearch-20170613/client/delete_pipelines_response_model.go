// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePipelinesResponse
	GetStatusCode() *int32
	SetBody(v *DeletePipelinesResponseBody) *DeletePipelinesResponse
	GetBody() *DeletePipelinesResponseBody
}

type DeletePipelinesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelinesResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePipelinesResponse) GetBody() *DeletePipelinesResponseBody {
	return s.Body
}

func (s *DeletePipelinesResponse) SetHeaders(v map[string]*string) *DeletePipelinesResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelinesResponse) SetStatusCode(v int32) *DeletePipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelinesResponse) SetBody(v *DeletePipelinesResponseBody) *DeletePipelinesResponse {
	s.Body = v
	return s
}

func (s *DeletePipelinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
