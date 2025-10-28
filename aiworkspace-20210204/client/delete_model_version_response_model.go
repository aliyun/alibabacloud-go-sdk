// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelVersionResponseBody) *DeleteModelVersionResponse
	GetBody() *DeleteModelVersionResponseBody
}

type DeleteModelVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelVersionResponse) GetBody() *DeleteModelVersionResponseBody {
	return s.Body
}

func (s *DeleteModelVersionResponse) SetHeaders(v map[string]*string) *DeleteModelVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelVersionResponse) SetStatusCode(v int32) *DeleteModelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelVersionResponse) SetBody(v *DeleteModelVersionResponseBody) *DeleteModelVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteModelVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
