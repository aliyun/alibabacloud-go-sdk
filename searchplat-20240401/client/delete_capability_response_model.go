// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCapabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCapabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCapabilityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCapabilityResponseBody) *DeleteCapabilityResponse
	GetBody() *DeleteCapabilityResponseBody
}

type DeleteCapabilityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCapabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCapabilityResponse) GoString() string {
	return s.String()
}

func (s *DeleteCapabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCapabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCapabilityResponse) GetBody() *DeleteCapabilityResponseBody {
	return s.Body
}

func (s *DeleteCapabilityResponse) SetHeaders(v map[string]*string) *DeleteCapabilityResponse {
	s.Headers = v
	return s
}

func (s *DeleteCapabilityResponse) SetStatusCode(v int32) *DeleteCapabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCapabilityResponse) SetBody(v *DeleteCapabilityResponseBody) *DeleteCapabilityResponse {
	s.Body = v
	return s
}

func (s *DeleteCapabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
