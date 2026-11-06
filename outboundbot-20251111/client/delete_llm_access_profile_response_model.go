// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLlmAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLlmAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLlmAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLlmAccessProfileResponseBody) *DeleteLlmAccessProfileResponse
	GetBody() *DeleteLlmAccessProfileResponseBody
}

type DeleteLlmAccessProfileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLlmAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLlmAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLlmAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteLlmAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLlmAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLlmAccessProfileResponse) GetBody() *DeleteLlmAccessProfileResponseBody {
	return s.Body
}

func (s *DeleteLlmAccessProfileResponse) SetHeaders(v map[string]*string) *DeleteLlmAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteLlmAccessProfileResponse) SetStatusCode(v int32) *DeleteLlmAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLlmAccessProfileResponse) SetBody(v *DeleteLlmAccessProfileResponseBody) *DeleteLlmAccessProfileResponse {
	s.Body = v
	return s
}

func (s *DeleteLlmAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
