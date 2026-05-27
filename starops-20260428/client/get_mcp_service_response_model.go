// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpServiceResponseBody) *GetMcpServiceResponse
	GetBody() *GetMcpServiceResponseBody
}

type GetMcpServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServiceResponse) GoString() string {
	return s.String()
}

func (s *GetMcpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpServiceResponse) GetBody() *GetMcpServiceResponseBody {
	return s.Body
}

func (s *GetMcpServiceResponse) SetHeaders(v map[string]*string) *GetMcpServiceResponse {
	s.Headers = v
	return s
}

func (s *GetMcpServiceResponse) SetStatusCode(v int32) *GetMcpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpServiceResponse) SetBody(v *GetMcpServiceResponseBody) *GetMcpServiceResponse {
	s.Body = v
	return s
}

func (s *GetMcpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
