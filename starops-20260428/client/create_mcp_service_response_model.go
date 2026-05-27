// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcpServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcpServiceResponseBody) *CreateMcpServiceResponse
	GetBody() *CreateMcpServiceResponseBody
}

type CreateMcpServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcpServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateMcpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcpServiceResponse) GetBody() *CreateMcpServiceResponseBody {
	return s.Body
}

func (s *CreateMcpServiceResponse) SetHeaders(v map[string]*string) *CreateMcpServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateMcpServiceResponse) SetStatusCode(v int32) *CreateMcpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcpServiceResponse) SetBody(v *CreateMcpServiceResponseBody) *CreateMcpServiceResponse {
	s.Body = v
	return s
}

func (s *CreateMcpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
