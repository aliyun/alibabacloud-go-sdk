// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserProvisioningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserProvisioningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserProvisioningResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserProvisioningResponseBody) *DeleteUserProvisioningResponse
	GetBody() *DeleteUserProvisioningResponseBody
}

type DeleteUserProvisioningResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserProvisioningResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserProvisioningResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserProvisioningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserProvisioningResponse) GetBody() *DeleteUserProvisioningResponseBody {
	return s.Body
}

func (s *DeleteUserProvisioningResponse) SetHeaders(v map[string]*string) *DeleteUserProvisioningResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserProvisioningResponse) SetStatusCode(v int32) *DeleteUserProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserProvisioningResponse) SetBody(v *DeleteUserProvisioningResponseBody) *DeleteUserProvisioningResponse {
	s.Body = v
	return s
}

func (s *DeleteUserProvisioningResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
