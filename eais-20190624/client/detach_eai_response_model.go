// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEaiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachEaiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachEaiResponse
	GetStatusCode() *int32
	SetBody(v *DetachEaiResponseBody) *DetachEaiResponse
	GetBody() *DetachEaiResponseBody
}

type DetachEaiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachEaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachEaiResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachEaiResponse) GoString() string {
	return s.String()
}

func (s *DetachEaiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachEaiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachEaiResponse) GetBody() *DetachEaiResponseBody {
	return s.Body
}

func (s *DetachEaiResponse) SetHeaders(v map[string]*string) *DetachEaiResponse {
	s.Headers = v
	return s
}

func (s *DetachEaiResponse) SetStatusCode(v int32) *DetachEaiResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachEaiResponse) SetBody(v *DetachEaiResponseBody) *DetachEaiResponse {
	s.Body = v
	return s
}

func (s *DetachEaiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
