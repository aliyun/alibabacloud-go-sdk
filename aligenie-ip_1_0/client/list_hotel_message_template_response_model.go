// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelMessageTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelMessageTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelMessageTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelMessageTemplateResponseBody) *ListHotelMessageTemplateResponse
	GetBody() *ListHotelMessageTemplateResponseBody
}

type ListHotelMessageTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelMessageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelMessageTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelMessageTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListHotelMessageTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelMessageTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelMessageTemplateResponse) GetBody() *ListHotelMessageTemplateResponseBody {
	return s.Body
}

func (s *ListHotelMessageTemplateResponse) SetHeaders(v map[string]*string) *ListHotelMessageTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListHotelMessageTemplateResponse) SetStatusCode(v int32) *ListHotelMessageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelMessageTemplateResponse) SetBody(v *ListHotelMessageTemplateResponseBody) *ListHotelMessageTemplateResponse {
	s.Body = v
	return s
}

func (s *ListHotelMessageTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
