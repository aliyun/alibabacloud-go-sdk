// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInterceptionTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInterceptionTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInterceptionTargetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInterceptionTargetResponseBody) *DeleteInterceptionTargetResponse
	GetBody() *DeleteInterceptionTargetResponseBody
}

type DeleteInterceptionTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInterceptionTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInterceptionTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterceptionTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteInterceptionTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInterceptionTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInterceptionTargetResponse) GetBody() *DeleteInterceptionTargetResponseBody {
	return s.Body
}

func (s *DeleteInterceptionTargetResponse) SetHeaders(v map[string]*string) *DeleteInterceptionTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteInterceptionTargetResponse) SetStatusCode(v int32) *DeleteInterceptionTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInterceptionTargetResponse) SetBody(v *DeleteInterceptionTargetResponseBody) *DeleteInterceptionTargetResponse {
	s.Body = v
	return s
}

func (s *DeleteInterceptionTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
