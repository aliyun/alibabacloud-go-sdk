// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewClusterResponse
	GetStatusCode() *int32
	SetBody(v *RenewClusterResponseBody) *RenewClusterResponse
	GetBody() *RenewClusterResponseBody
}

type RenewClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewClusterResponse) GoString() string {
	return s.String()
}

func (s *RenewClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewClusterResponse) GetBody() *RenewClusterResponseBody {
	return s.Body
}

func (s *RenewClusterResponse) SetHeaders(v map[string]*string) *RenewClusterResponse {
	s.Headers = v
	return s
}

func (s *RenewClusterResponse) SetStatusCode(v int32) *RenewClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewClusterResponse) SetBody(v *RenewClusterResponseBody) *RenewClusterResponse {
	s.Body = v
	return s
}

func (s *RenewClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
