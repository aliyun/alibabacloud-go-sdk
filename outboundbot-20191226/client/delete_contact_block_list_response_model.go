// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactBlockListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactBlockListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactBlockListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactBlockListResponseBody) *DeleteContactBlockListResponse
	GetBody() *DeleteContactBlockListResponseBody
}

type DeleteContactBlockListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactBlockListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactBlockListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactBlockListResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactBlockListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactBlockListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactBlockListResponse) GetBody() *DeleteContactBlockListResponseBody {
	return s.Body
}

func (s *DeleteContactBlockListResponse) SetHeaders(v map[string]*string) *DeleteContactBlockListResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactBlockListResponse) SetStatusCode(v int32) *DeleteContactBlockListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactBlockListResponse) SetBody(v *DeleteContactBlockListResponseBody) *DeleteContactBlockListResponse {
	s.Body = v
	return s
}

func (s *DeleteContactBlockListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
