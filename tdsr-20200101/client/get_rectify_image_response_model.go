// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRectifyImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRectifyImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRectifyImageResponse
	GetStatusCode() *int32
	SetBody(v *GetRectifyImageResponseBody) *GetRectifyImageResponse
	GetBody() *GetRectifyImageResponseBody
}

type GetRectifyImageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRectifyImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRectifyImageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyImageResponse) GoString() string {
	return s.String()
}

func (s *GetRectifyImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRectifyImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRectifyImageResponse) GetBody() *GetRectifyImageResponseBody {
	return s.Body
}

func (s *GetRectifyImageResponse) SetHeaders(v map[string]*string) *GetRectifyImageResponse {
	s.Headers = v
	return s
}

func (s *GetRectifyImageResponse) SetStatusCode(v int32) *GetRectifyImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRectifyImageResponse) SetBody(v *GetRectifyImageResponseBody) *GetRectifyImageResponse {
	s.Body = v
	return s
}

func (s *GetRectifyImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
