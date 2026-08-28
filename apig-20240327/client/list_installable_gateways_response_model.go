// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstallableGatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstallableGatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstallableGatewaysResponse
	GetStatusCode() *int32
	SetBody(v *ListInstallableGatewaysResponseBody) *ListInstallableGatewaysResponse
	GetBody() *ListInstallableGatewaysResponseBody
}

type ListInstallableGatewaysResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstallableGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstallableGatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstallableGatewaysResponse) GoString() string {
	return s.String()
}

func (s *ListInstallableGatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstallableGatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstallableGatewaysResponse) GetBody() *ListInstallableGatewaysResponseBody {
	return s.Body
}

func (s *ListInstallableGatewaysResponse) SetHeaders(v map[string]*string) *ListInstallableGatewaysResponse {
	s.Headers = v
	return s
}

func (s *ListInstallableGatewaysResponse) SetStatusCode(v int32) *ListInstallableGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstallableGatewaysResponse) SetBody(v *ListInstallableGatewaysResponseBody) *ListInstallableGatewaysResponse {
	s.Body = v
	return s
}

func (s *ListInstallableGatewaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
