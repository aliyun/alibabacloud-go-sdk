// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotlineNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHotlineNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHotlineNumberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHotlineNumberResponseBody) *DeleteHotlineNumberResponse
	GetBody() *DeleteHotlineNumberResponseBody
}

type DeleteHotlineNumberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHotlineNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHotlineNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotlineNumberResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotlineNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHotlineNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHotlineNumberResponse) GetBody() *DeleteHotlineNumberResponseBody {
	return s.Body
}

func (s *DeleteHotlineNumberResponse) SetHeaders(v map[string]*string) *DeleteHotlineNumberResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotlineNumberResponse) SetStatusCode(v int32) *DeleteHotlineNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotlineNumberResponse) SetBody(v *DeleteHotlineNumberResponseBody) *DeleteHotlineNumberResponse {
	s.Body = v
	return s
}

func (s *DeleteHotlineNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
