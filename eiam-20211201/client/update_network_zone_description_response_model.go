// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkZoneDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNetworkZoneDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNetworkZoneDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNetworkZoneDescriptionResponseBody) *UpdateNetworkZoneDescriptionResponse
	GetBody() *UpdateNetworkZoneDescriptionResponseBody
}

type UpdateNetworkZoneDescriptionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkZoneDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkZoneDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkZoneDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkZoneDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNetworkZoneDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNetworkZoneDescriptionResponse) GetBody() *UpdateNetworkZoneDescriptionResponseBody {
	return s.Body
}

func (s *UpdateNetworkZoneDescriptionResponse) SetHeaders(v map[string]*string) *UpdateNetworkZoneDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkZoneDescriptionResponse) SetStatusCode(v int32) *UpdateNetworkZoneDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkZoneDescriptionResponse) SetBody(v *UpdateNetworkZoneDescriptionResponseBody) *UpdateNetworkZoneDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateNetworkZoneDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
