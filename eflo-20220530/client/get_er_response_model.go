// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErResponse
	GetStatusCode() *int32
	SetBody(v *GetErResponseBody) *GetErResponse
	GetBody() *GetErResponseBody
}

type GetErResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErResponse) GoString() string {
	return s.String()
}

func (s *GetErResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErResponse) GetBody() *GetErResponseBody {
	return s.Body
}

func (s *GetErResponse) SetHeaders(v map[string]*string) *GetErResponse {
	s.Headers = v
	return s
}

func (s *GetErResponse) SetStatusCode(v int32) *GetErResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErResponse) SetBody(v *GetErResponseBody) *GetErResponse {
	s.Body = v
	return s
}

func (s *GetErResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
