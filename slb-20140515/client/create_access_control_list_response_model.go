// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessControlListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessControlListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessControlListResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessControlListResponseBody) *CreateAccessControlListResponse
	GetBody() *CreateAccessControlListResponseBody
}

type CreateAccessControlListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessControlListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessControlListResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessControlListResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessControlListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessControlListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessControlListResponse) GetBody() *CreateAccessControlListResponseBody {
	return s.Body
}

func (s *CreateAccessControlListResponse) SetHeaders(v map[string]*string) *CreateAccessControlListResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessControlListResponse) SetStatusCode(v int32) *CreateAccessControlListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessControlListResponse) SetBody(v *CreateAccessControlListResponseBody) *CreateAccessControlListResponse {
	s.Body = v
	return s
}

func (s *CreateAccessControlListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
