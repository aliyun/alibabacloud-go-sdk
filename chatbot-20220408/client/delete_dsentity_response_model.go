// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDSEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDSEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDSEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDSEntityResponseBody) *DeleteDSEntityResponse
	GetBody() *DeleteDSEntityResponseBody
}

type DeleteDSEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDSEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDSEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDSEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDSEntityResponse) GetBody() *DeleteDSEntityResponseBody {
	return s.Body
}

func (s *DeleteDSEntityResponse) SetHeaders(v map[string]*string) *DeleteDSEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteDSEntityResponse) SetStatusCode(v int32) *DeleteDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDSEntityResponse) SetBody(v *DeleteDSEntityResponseBody) *DeleteDSEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteDSEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
