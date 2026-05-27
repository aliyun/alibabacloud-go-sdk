// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcpServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcpServiceResponseBody) *DeleteMcpServiceResponse
	GetBody() *DeleteMcpServiceResponseBody
}

type DeleteMcpServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcpServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcpServiceResponse) GetBody() *DeleteMcpServiceResponseBody {
	return s.Body
}

func (s *DeleteMcpServiceResponse) SetHeaders(v map[string]*string) *DeleteMcpServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcpServiceResponse) SetStatusCode(v int32) *DeleteMcpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcpServiceResponse) SetBody(v *DeleteMcpServiceResponseBody) *DeleteMcpServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteMcpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
