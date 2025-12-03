// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iXpackRelateDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *XpackRelateDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *XpackRelateDBResponse
	GetStatusCode() *int32
	SetBody(v *XpackRelateDBResponseBody) *XpackRelateDBResponse
	GetBody() *XpackRelateDBResponseBody
}

type XpackRelateDBResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *XpackRelateDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s XpackRelateDBResponse) String() string {
	return dara.Prettify(s)
}

func (s XpackRelateDBResponse) GoString() string {
	return s.String()
}

func (s *XpackRelateDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *XpackRelateDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *XpackRelateDBResponse) GetBody() *XpackRelateDBResponseBody {
	return s.Body
}

func (s *XpackRelateDBResponse) SetHeaders(v map[string]*string) *XpackRelateDBResponse {
	s.Headers = v
	return s
}

func (s *XpackRelateDBResponse) SetStatusCode(v int32) *XpackRelateDBResponse {
	s.StatusCode = &v
	return s
}

func (s *XpackRelateDBResponse) SetBody(v *XpackRelateDBResponseBody) *XpackRelateDBResponse {
	s.Body = v
	return s
}

func (s *XpackRelateDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
