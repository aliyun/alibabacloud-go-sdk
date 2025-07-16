// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssistantCapabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssistantCapabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssistantCapabilityResponse
	GetStatusCode() *int32
	SetBody(v *GetAssistantCapabilityResponseBody) *GetAssistantCapabilityResponse
	GetBody() *GetAssistantCapabilityResponseBody
}

type GetAssistantCapabilityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssistantCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssistantCapabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssistantCapabilityResponse) GoString() string {
	return s.String()
}

func (s *GetAssistantCapabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssistantCapabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssistantCapabilityResponse) GetBody() *GetAssistantCapabilityResponseBody {
	return s.Body
}

func (s *GetAssistantCapabilityResponse) SetHeaders(v map[string]*string) *GetAssistantCapabilityResponse {
	s.Headers = v
	return s
}

func (s *GetAssistantCapabilityResponse) SetStatusCode(v int32) *GetAssistantCapabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssistantCapabilityResponse) SetBody(v *GetAssistantCapabilityResponseBody) *GetAssistantCapabilityResponse {
	s.Body = v
	return s
}

func (s *GetAssistantCapabilityResponse) Validate() error {
	return dara.Validate(s)
}
