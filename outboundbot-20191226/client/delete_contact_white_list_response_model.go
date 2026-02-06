// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactWhiteListResponseBody) *DeleteContactWhiteListResponse
	GetBody() *DeleteContactWhiteListResponseBody
}

type DeleteContactWhiteListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactWhiteListResponse) GetBody() *DeleteContactWhiteListResponseBody {
	return s.Body
}

func (s *DeleteContactWhiteListResponse) SetHeaders(v map[string]*string) *DeleteContactWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactWhiteListResponse) SetStatusCode(v int32) *DeleteContactWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactWhiteListResponse) SetBody(v *DeleteContactWhiteListResponseBody) *DeleteContactWhiteListResponse {
	s.Body = v
	return s
}

func (s *DeleteContactWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
