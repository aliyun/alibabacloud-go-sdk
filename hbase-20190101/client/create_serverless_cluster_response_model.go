// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerlessClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServerlessClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServerlessClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateServerlessClusterResponseBody) *CreateServerlessClusterResponse
	GetBody() *CreateServerlessClusterResponseBody
}

type CreateServerlessClusterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServerlessClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServerlessClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServerlessClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateServerlessClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServerlessClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServerlessClusterResponse) GetBody() *CreateServerlessClusterResponseBody {
	return s.Body
}

func (s *CreateServerlessClusterResponse) SetHeaders(v map[string]*string) *CreateServerlessClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateServerlessClusterResponse) SetStatusCode(v int32) *CreateServerlessClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerlessClusterResponse) SetBody(v *CreateServerlessClusterResponseBody) *CreateServerlessClusterResponse {
	s.Body = v
	return s
}

func (s *CreateServerlessClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
