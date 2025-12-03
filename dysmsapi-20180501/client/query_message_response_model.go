// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMessageResponse
	GetStatusCode() *int32
	SetBody(v *QueryMessageResponseBody) *QueryMessageResponse
	GetBody() *QueryMessageResponseBody
}

type QueryMessageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageResponse) GoString() string {
	return s.String()
}

func (s *QueryMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMessageResponse) GetBody() *QueryMessageResponseBody {
	return s.Body
}

func (s *QueryMessageResponse) SetHeaders(v map[string]*string) *QueryMessageResponse {
	s.Headers = v
	return s
}

func (s *QueryMessageResponse) SetStatusCode(v int32) *QueryMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMessageResponse) SetBody(v *QueryMessageResponseBody) *QueryMessageResponse {
	s.Body = v
	return s
}

func (s *QueryMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
