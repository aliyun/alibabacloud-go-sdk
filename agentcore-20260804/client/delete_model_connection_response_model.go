// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelConnectionResponseBody) *DeleteModelConnectionResponse
	GetBody() *DeleteModelConnectionResponseBody
}

type DeleteModelConnectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelConnectionResponse) GetBody() *DeleteModelConnectionResponseBody {
	return s.Body
}

func (s *DeleteModelConnectionResponse) SetHeaders(v map[string]*string) *DeleteModelConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelConnectionResponse) SetStatusCode(v int32) *DeleteModelConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelConnectionResponse) SetBody(v *DeleteModelConnectionResponseBody) *DeleteModelConnectionResponse {
	s.Body = v
	return s
}

func (s *DeleteModelConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
