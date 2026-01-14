// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceManagedControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceManagedControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceManagedControlResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceManagedControlResponseBody) *UpdateServiceManagedControlResponse
	GetBody() *UpdateServiceManagedControlResponseBody
}

type UpdateServiceManagedControlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceManagedControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceManagedControlResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceManagedControlResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceManagedControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceManagedControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceManagedControlResponse) GetBody() *UpdateServiceManagedControlResponseBody {
	return s.Body
}

func (s *UpdateServiceManagedControlResponse) SetHeaders(v map[string]*string) *UpdateServiceManagedControlResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceManagedControlResponse) SetStatusCode(v int32) *UpdateServiceManagedControlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceManagedControlResponse) SetBody(v *UpdateServiceManagedControlResponseBody) *UpdateServiceManagedControlResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceManagedControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
