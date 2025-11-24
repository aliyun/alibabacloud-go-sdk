// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegisteredServiceNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegisteredServiceNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegisteredServiceNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *GetRegisteredServiceNamespacesResponseBody) *GetRegisteredServiceNamespacesResponse
	GetBody() *GetRegisteredServiceNamespacesResponseBody
}

type GetRegisteredServiceNamespacesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegisteredServiceNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegisteredServiceNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegisteredServiceNamespacesResponse) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegisteredServiceNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegisteredServiceNamespacesResponse) GetBody() *GetRegisteredServiceNamespacesResponseBody {
	return s.Body
}

func (s *GetRegisteredServiceNamespacesResponse) SetHeaders(v map[string]*string) *GetRegisteredServiceNamespacesResponse {
	s.Headers = v
	return s
}

func (s *GetRegisteredServiceNamespacesResponse) SetStatusCode(v int32) *GetRegisteredServiceNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegisteredServiceNamespacesResponse) SetBody(v *GetRegisteredServiceNamespacesResponseBody) *GetRegisteredServiceNamespacesResponse {
	s.Body = v
	return s
}

func (s *GetRegisteredServiceNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
