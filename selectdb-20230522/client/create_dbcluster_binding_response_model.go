// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBClusterBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBClusterBindingResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBClusterBindingResponseBody) *CreateDBClusterBindingResponse
	GetBody() *CreateDBClusterBindingResponseBody
}

type CreateDBClusterBindingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBClusterBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBClusterBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterBindingResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBClusterBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBClusterBindingResponse) GetBody() *CreateDBClusterBindingResponseBody {
	return s.Body
}

func (s *CreateDBClusterBindingResponse) SetHeaders(v map[string]*string) *CreateDBClusterBindingResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterBindingResponse) SetStatusCode(v int32) *CreateDBClusterBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBClusterBindingResponse) SetBody(v *CreateDBClusterBindingResponseBody) *CreateDBClusterBindingResponse {
	s.Body = v
	return s
}

func (s *CreateDBClusterBindingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
