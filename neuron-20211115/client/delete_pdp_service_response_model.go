// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePdpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePdpServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeletePdpServiceResponseBody) *DeletePdpServiceResponse
	GetBody() *DeletePdpServiceResponseBody
}

type DeletePdpServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePdpServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePdpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpServiceResponse) GoString() string {
	return s.String()
}

func (s *DeletePdpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePdpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePdpServiceResponse) GetBody() *DeletePdpServiceResponseBody {
	return s.Body
}

func (s *DeletePdpServiceResponse) SetHeaders(v map[string]*string) *DeletePdpServiceResponse {
	s.Headers = v
	return s
}

func (s *DeletePdpServiceResponse) SetStatusCode(v int32) *DeletePdpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePdpServiceResponse) SetBody(v *DeletePdpServiceResponseBody) *DeletePdpServiceResponse {
	s.Body = v
	return s
}

func (s *DeletePdpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
