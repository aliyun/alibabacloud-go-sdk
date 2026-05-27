// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMcpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMcpServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMcpServiceResponseBody) *UpdateMcpServiceResponse
	GetBody() *UpdateMcpServiceResponseBody
}

type UpdateMcpServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMcpServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMcpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateMcpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMcpServiceResponse) GetBody() *UpdateMcpServiceResponseBody {
	return s.Body
}

func (s *UpdateMcpServiceResponse) SetHeaders(v map[string]*string) *UpdateMcpServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateMcpServiceResponse) SetStatusCode(v int32) *UpdateMcpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMcpServiceResponse) SetBody(v *UpdateMcpServiceResponseBody) *UpdateMcpServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateMcpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
