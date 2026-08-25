// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserProvisioningEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserProvisioningEventResponse
	GetStatusCode() *int32
	SetBody(v *GetUserProvisioningEventResponseBody) *GetUserProvisioningEventResponse
	GetBody() *GetUserProvisioningEventResponseBody
}

type GetUserProvisioningEventResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningEventResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserProvisioningEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserProvisioningEventResponse) GetBody() *GetUserProvisioningEventResponseBody {
	return s.Body
}

func (s *GetUserProvisioningEventResponse) SetHeaders(v map[string]*string) *GetUserProvisioningEventResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningEventResponse) SetStatusCode(v int32) *GetUserProvisioningEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningEventResponse) SetBody(v *GetUserProvisioningEventResponseBody) *GetUserProvisioningEventResponse {
	s.Body = v
	return s
}

func (s *GetUserProvisioningEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
