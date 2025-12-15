// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupCapabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceGroupCapabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceGroupCapabilityResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceGroupCapabilityResponseBody) *ListResourceGroupCapabilityResponse
	GetBody() *ListResourceGroupCapabilityResponseBody
}

type ListResourceGroupCapabilityResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupCapabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupCapabilityResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupCapabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceGroupCapabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceGroupCapabilityResponse) GetBody() *ListResourceGroupCapabilityResponseBody {
	return s.Body
}

func (s *ListResourceGroupCapabilityResponse) SetHeaders(v map[string]*string) *ListResourceGroupCapabilityResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupCapabilityResponse) SetStatusCode(v int32) *ListResourceGroupCapabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupCapabilityResponse) SetBody(v *ListResourceGroupCapabilityResponseBody) *ListResourceGroupCapabilityResponse {
	s.Body = v
	return s
}

func (s *ListResourceGroupCapabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
