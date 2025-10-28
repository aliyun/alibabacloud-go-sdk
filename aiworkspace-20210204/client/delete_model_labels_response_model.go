// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelLabelsResponseBody) *DeleteModelLabelsResponse
	GetBody() *DeleteModelLabelsResponseBody
}

type DeleteModelLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelLabelsResponse) GetBody() *DeleteModelLabelsResponseBody {
	return s.Body
}

func (s *DeleteModelLabelsResponse) SetHeaders(v map[string]*string) *DeleteModelLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelLabelsResponse) SetStatusCode(v int32) *DeleteModelLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelLabelsResponse) SetBody(v *DeleteModelLabelsResponseBody) *DeleteModelLabelsResponse {
	s.Body = v
	return s
}

func (s *DeleteModelLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
