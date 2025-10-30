// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBizEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBizEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBizEntityResponseBody) *DeleteBizEntityResponse
	GetBody() *DeleteBizEntityResponseBody
}

type DeleteBizEntityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBizEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBizEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteBizEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBizEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBizEntityResponse) GetBody() *DeleteBizEntityResponseBody {
	return s.Body
}

func (s *DeleteBizEntityResponse) SetHeaders(v map[string]*string) *DeleteBizEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteBizEntityResponse) SetStatusCode(v int32) *DeleteBizEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBizEntityResponse) SetBody(v *DeleteBizEntityResponseBody) *DeleteBizEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteBizEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
