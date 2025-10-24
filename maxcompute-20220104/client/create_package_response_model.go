// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePackageResponse
	GetStatusCode() *int32
	SetBody(v *CreatePackageResponseBody) *CreatePackageResponse
	GetBody() *CreatePackageResponseBody
}

type CreatePackageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePackageResponse) GoString() string {
	return s.String()
}

func (s *CreatePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePackageResponse) GetBody() *CreatePackageResponseBody {
	return s.Body
}

func (s *CreatePackageResponse) SetHeaders(v map[string]*string) *CreatePackageResponse {
	s.Headers = v
	return s
}

func (s *CreatePackageResponse) SetStatusCode(v int32) *CreatePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePackageResponse) SetBody(v *CreatePackageResponseBody) *CreatePackageResponse {
	s.Body = v
	return s
}

func (s *CreatePackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
