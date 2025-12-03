// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessControlListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessControlListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessControlListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessControlListResponseBody) *DeleteAccessControlListResponse
	GetBody() *DeleteAccessControlListResponseBody
}

type DeleteAccessControlListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessControlListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessControlListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessControlListResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessControlListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessControlListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessControlListResponse) GetBody() *DeleteAccessControlListResponseBody {
	return s.Body
}

func (s *DeleteAccessControlListResponse) SetHeaders(v map[string]*string) *DeleteAccessControlListResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessControlListResponse) SetStatusCode(v int32) *DeleteAccessControlListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessControlListResponse) SetBody(v *DeleteAccessControlListResponseBody) *DeleteAccessControlListResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessControlListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
