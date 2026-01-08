// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsByIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactsByIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactsByIdsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactsByIdsResponseBody) *DeleteContactsByIdsResponse
	GetBody() *DeleteContactsByIdsResponseBody
}

type DeleteContactsByIdsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactsByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactsByIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsByIdsResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactsByIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactsByIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactsByIdsResponse) GetBody() *DeleteContactsByIdsResponseBody {
	return s.Body
}

func (s *DeleteContactsByIdsResponse) SetHeaders(v map[string]*string) *DeleteContactsByIdsResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactsByIdsResponse) SetStatusCode(v int32) *DeleteContactsByIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactsByIdsResponse) SetBody(v *DeleteContactsByIdsResponseBody) *DeleteContactsByIdsResponse {
	s.Body = v
	return s
}

func (s *DeleteContactsByIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
