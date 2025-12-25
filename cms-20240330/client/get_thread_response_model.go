// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetThreadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetThreadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetThreadResponse
	GetStatusCode() *int32
	SetBody(v *GetThreadResponseBody) *GetThreadResponse
	GetBody() *GetThreadResponseBody
}

type GetThreadResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetThreadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetThreadResponse) String() string {
	return dara.Prettify(s)
}

func (s GetThreadResponse) GoString() string {
	return s.String()
}

func (s *GetThreadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetThreadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetThreadResponse) GetBody() *GetThreadResponseBody {
	return s.Body
}

func (s *GetThreadResponse) SetHeaders(v map[string]*string) *GetThreadResponse {
	s.Headers = v
	return s
}

func (s *GetThreadResponse) SetStatusCode(v int32) *GetThreadResponse {
	s.StatusCode = &v
	return s
}

func (s *GetThreadResponse) SetBody(v *GetThreadResponseBody) *GetThreadResponse {
	s.Body = v
	return s
}

func (s *GetThreadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
