// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserProvisioningEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserProvisioningEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserProvisioningEventResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserProvisioningEventResponseBody) *DeleteUserProvisioningEventResponse
	GetBody() *DeleteUserProvisioningEventResponseBody
}

type DeleteUserProvisioningEventResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserProvisioningEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserProvisioningEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserProvisioningEventResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserProvisioningEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserProvisioningEventResponse) GetBody() *DeleteUserProvisioningEventResponseBody {
	return s.Body
}

func (s *DeleteUserProvisioningEventResponse) SetHeaders(v map[string]*string) *DeleteUserProvisioningEventResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserProvisioningEventResponse) SetStatusCode(v int32) *DeleteUserProvisioningEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserProvisioningEventResponse) SetBody(v *DeleteUserProvisioningEventResponseBody) *DeleteUserProvisioningEventResponse {
	s.Body = v
	return s
}

func (s *DeleteUserProvisioningEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
