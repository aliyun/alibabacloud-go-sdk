// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactsResponseBody) *DeleteContactsResponse
	GetBody() *DeleteContactsResponseBody
}

type DeleteContactsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactsResponse) GetBody() *DeleteContactsResponseBody {
	return s.Body
}

func (s *DeleteContactsResponse) SetHeaders(v map[string]*string) *DeleteContactsResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactsResponse) SetStatusCode(v int32) *DeleteContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactsResponse) SetBody(v *DeleteContactsResponseBody) *DeleteContactsResponse {
	s.Body = v
	return s
}

func (s *DeleteContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
