// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRectVerticalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RectVerticalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RectVerticalResponse
	GetStatusCode() *int32
	SetBody(v *RectVerticalResponseBody) *RectVerticalResponse
	GetBody() *RectVerticalResponseBody
}

type RectVerticalResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RectVerticalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RectVerticalResponse) String() string {
	return dara.Prettify(s)
}

func (s RectVerticalResponse) GoString() string {
	return s.String()
}

func (s *RectVerticalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RectVerticalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RectVerticalResponse) GetBody() *RectVerticalResponseBody {
	return s.Body
}

func (s *RectVerticalResponse) SetHeaders(v map[string]*string) *RectVerticalResponse {
	s.Headers = v
	return s
}

func (s *RectVerticalResponse) SetStatusCode(v int32) *RectVerticalResponse {
	s.StatusCode = &v
	return s
}

func (s *RectVerticalResponse) SetBody(v *RectVerticalResponseBody) *RectVerticalResponse {
	s.Body = v
	return s
}

func (s *RectVerticalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
