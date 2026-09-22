// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListZonesResponse
	GetStatusCode() *int32
	SetBody(v *ListZonesResponseBody) *ListZonesResponse
	GetBody() *ListZonesResponseBody
}

type ListZonesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListZonesResponse) GoString() string {
	return s.String()
}

func (s *ListZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListZonesResponse) GetBody() *ListZonesResponseBody {
	return s.Body
}

func (s *ListZonesResponse) SetHeaders(v map[string]*string) *ListZonesResponse {
	s.Headers = v
	return s
}

func (s *ListZonesResponse) SetStatusCode(v int32) *ListZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListZonesResponse) SetBody(v *ListZonesResponseBody) *ListZonesResponse {
	s.Body = v
	return s
}

func (s *ListZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
