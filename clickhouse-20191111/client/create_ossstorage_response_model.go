// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOSSStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOSSStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOSSStorageResponse
	GetStatusCode() *int32
	SetBody(v *CreateOSSStorageResponseBody) *CreateOSSStorageResponse
	GetBody() *CreateOSSStorageResponseBody
}

type CreateOSSStorageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOSSStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOSSStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOSSStorageResponse) GoString() string {
	return s.String()
}

func (s *CreateOSSStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOSSStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOSSStorageResponse) GetBody() *CreateOSSStorageResponseBody {
	return s.Body
}

func (s *CreateOSSStorageResponse) SetHeaders(v map[string]*string) *CreateOSSStorageResponse {
	s.Headers = v
	return s
}

func (s *CreateOSSStorageResponse) SetStatusCode(v int32) *CreateOSSStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOSSStorageResponse) SetBody(v *CreateOSSStorageResponseBody) *CreateOSSStorageResponse {
	s.Body = v
	return s
}

func (s *CreateOSSStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
