// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationSupportedProvisionProtocolTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationSupportedProvisionProtocolTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationSupportedProvisionProtocolTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationSupportedProvisionProtocolTypesResponseBody) *ListApplicationSupportedProvisionProtocolTypesResponse
	GetBody() *ListApplicationSupportedProvisionProtocolTypesResponseBody
}

type ListApplicationSupportedProvisionProtocolTypesResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationSupportedProvisionProtocolTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationSupportedProvisionProtocolTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationSupportedProvisionProtocolTypesResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponse) GetBody() *ListApplicationSupportedProvisionProtocolTypesResponseBody {
	return s.Body
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponse) SetHeaders(v map[string]*string) *ListApplicationSupportedProvisionProtocolTypesResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponse) SetStatusCode(v int32) *ListApplicationSupportedProvisionProtocolTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponse) SetBody(v *ListApplicationSupportedProvisionProtocolTypesResponseBody) *ListApplicationSupportedProvisionProtocolTypesResponse {
	s.Body = v
	return s
}

func (s *ListApplicationSupportedProvisionProtocolTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
