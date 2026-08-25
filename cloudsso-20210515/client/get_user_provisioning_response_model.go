// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserProvisioningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserProvisioningResponse
	GetStatusCode() *int32
	SetBody(v *GetUserProvisioningResponseBody) *GetUserProvisioningResponse
	GetBody() *GetUserProvisioningResponseBody
}

type GetUserProvisioningResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserProvisioningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserProvisioningResponse) GetBody() *GetUserProvisioningResponseBody {
	return s.Body
}

func (s *GetUserProvisioningResponse) SetHeaders(v map[string]*string) *GetUserProvisioningResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningResponse) SetStatusCode(v int32) *GetUserProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningResponse) SetBody(v *GetUserProvisioningResponseBody) *GetUserProvisioningResponse {
	s.Body = v
	return s
}

func (s *GetUserProvisioningResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
