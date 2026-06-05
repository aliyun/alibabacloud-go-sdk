// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHiveResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHiveResponseBody) *DeleteHiveResponse
	GetBody() *DeleteHiveResponseBody
}

type DeleteHiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHiveResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHiveResponse) GoString() string {
	return s.String()
}

func (s *DeleteHiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHiveResponse) GetBody() *DeleteHiveResponseBody {
	return s.Body
}

func (s *DeleteHiveResponse) SetHeaders(v map[string]*string) *DeleteHiveResponse {
	s.Headers = v
	return s
}

func (s *DeleteHiveResponse) SetStatusCode(v int32) *DeleteHiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHiveResponse) SetBody(v *DeleteHiveResponseBody) *DeleteHiveResponse {
	s.Body = v
	return s
}

func (s *DeleteHiveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
