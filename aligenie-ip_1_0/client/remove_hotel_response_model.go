// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveHotelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveHotelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveHotelResponse
	GetStatusCode() *int32
	SetBody(v *RemoveHotelResponseBody) *RemoveHotelResponse
	GetBody() *RemoveHotelResponseBody
}

type RemoveHotelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveHotelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveHotelResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveHotelResponse) GoString() string {
	return s.String()
}

func (s *RemoveHotelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveHotelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveHotelResponse) GetBody() *RemoveHotelResponseBody {
	return s.Body
}

func (s *RemoveHotelResponse) SetHeaders(v map[string]*string) *RemoveHotelResponse {
	s.Headers = v
	return s
}

func (s *RemoveHotelResponse) SetStatusCode(v int32) *RemoveHotelResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveHotelResponse) SetBody(v *RemoveHotelResponseBody) *RemoveHotelResponse {
	s.Body = v
	return s
}

func (s *RemoveHotelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
